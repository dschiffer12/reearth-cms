package mongogit

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	client *mongox.ClientCollection
}

func NewCollection(client *mongox.ClientCollection) *Collection {
	return &Collection{client: client}
}

func (c *Collection) Client() *mongox.ClientCollection {
	return c.client
}

func (c *Collection) FindOne(ctx context.Context, filter any, q version.Query, consumer mongox.Consumer) error {
	return c.client.FindOne(ctx, apply(q, filter), consumer)
}

func (c *Collection) Find(ctx context.Context, filter any, q version.Query, consumer mongox.Consumer) error {
	return c.client.Find(ctx, apply(q, filter), consumer)
}

func (c *Collection) Paginate(ctx context.Context, filter any, q version.Query, p *usecasex.Pagination, consumer mongox.Consumer) (*usecasex.PageInfo, error) {
	return c.client.Paginate(ctx, apply(q, filter), nil, p, consumer)
}

func (c *Collection) Count(ctx context.Context, filter any, q version.Query) (int64, error) {
	return c.client.Count(ctx, apply(q, filter))
}

func (c *Collection) SaveOne(ctx context.Context, id string, replacement any, vr *version.VersionOrRef) error {
	if archived, err := c.IsArchived(ctx, id); err != nil {
		return err
	} else if archived {
		return version.ErrArchived
	}

	actualVr := lo.FromPtrOr(vr, version.Latest.OrVersion())

	meta, err := c.meta(ctx, id, actualVr.Ref())
	if err != nil {
		return err
	}

	var newmeta Meta
	var refs []version.Ref
	actualVr.Match(nil, func(r version.Ref) { refs = []version.Ref{r} })
	if meta == nil {
		if !actualVr.IsRef(version.Latest) {
			return rerror.ErrNotFound // invalid dest
		}
		newmeta = Meta{
			Version: version.New(),
			Refs:    refs,
		}
	} else {
		newmeta = Meta{
			Version: version.New(),
			Parents: []version.Version{meta.Version},
			Refs:    refs,
		}
	}

	if err := version.MatchVersionOrRef(actualVr, nil, func(r version.Ref) error {
		return c.UpdateRef(ctx, id, r, nil)
	}); err != nil {
		return err
	}

	if _, err := c.client.Client().InsertOne(ctx, &Document[any]{
		Data: replacement,
		Meta: newmeta,
	}); err != nil {
		return rerror.ErrInternalBy(err)
	}

	return nil
}

func (c *Collection) UpdateRef(ctx context.Context, id string, ref version.Ref, dest *version.VersionOrRef) error {
	current, err := c.meta(ctx, id, ref.OrVersion().Ref())
	if err != nil && !errors.Is(err, rerror.ErrNotFound) {
		return err
	}

	if current != nil {
		if _, err := c.client.Client().UpdateOne(ctx, bson.M{
			"id":       id,
			versionKey: current.Version,
		}, bson.M{
			"$pull": bson.M{refsKey: ref},
		}); err != nil {
			return rerror.ErrInternalBy(err)
		}
	}

	if dest != nil {
		if _, err := c.client.Client().UpdateOne(ctx, apply(version.Eq(*dest), bson.M{
			"id": id,
		}), bson.M{
			"$push": bson.M{refsKey: ref},
		}); err != nil {
			return rerror.ErrInternalBy(err)
		}
	}

	return nil
}

func (c *Collection) IsArchived(ctx context.Context, id string) (bool, error) {
	cons := mongox.SliceConsumer[MetadataDocument]{}
	if err := c.client.FindOne(ctx, bson.M{
		"id":    id,
		metaKey: true,
	}, &cons); err != nil {
		if errors.Is(rerror.ErrNotFound, err) {
			return false, nil
		}
		return false, err
	}
	return cons.Result[0].Archived, nil
}

func (c *Collection) ArchiveOne(ctx context.Context, id string, archived bool) error {
	if !archived {
		_, err := c.client.Client().DeleteOne(ctx, bson.M{"id": id, metaKey: true})
		if err != nil {
			return rerror.ErrInternalBy(err)
		}
		return nil
	}

	_, err := c.client.Client().ReplaceOne(ctx, bson.M{"id": id, metaKey: true}, MetadataDocument{
		ID:       id,
		Meta:     true,
		Archived: archived,
	}, options.Replace().SetUpsert(true))
	if err != nil {
		return rerror.ErrInternalBy(err)
	}
	return nil
}

func (c *Collection) RemoveOne(ctx context.Context, id string) error {
	return c.client.RemoveAll(ctx, bson.M{"id": id})
}

func (c *Collection) Empty(ctx context.Context) error {
	return c.client.Client().Drop(ctx)
}

func (c *Collection) CreateIndexes(ctx context.Context, keys, uniqueKeys []string) error {
	indexes := append(
		[]mongo.IndexModel{
			{Keys: []string{"id", versionKey}, Options: options.Index().SetUnique(true)},
			{Keys: []string{"id", metaKey}, Options: options.Index().SetUnique(true)},
			{Keys: []string{refsKey}},
			{Keys: []string{parentsKey}},
		},
		append(
			util.Map(keys, func(k string) mongo.IndexModel {
				return mongo.IndexModel{Keys: []string{k}}
			}),
			util.Map(uniqueKeys, func(k string) mongo.IndexModel {
				return mongo.IndexModel{Keys: []string{k}, Options: options.Index().SetUnique(true)}
			})...,
		)...,
	)

	if _, err := c.client.Client().Indexes().CreateMany(ctx, indexes); err != nil {
		return rerror.ErrInternalBy(err)
	}
	return nil
}

func (c *Collection) meta(ctx context.Context, id string, v *version.VersionOrRef) (*Meta, error) {
	consumer := mongox.SliceConsumer[Meta]{}
	if err := c.client.FindOne(ctx, apply(version.Eq(lo.FromPtrOr(v, version.Latest.OrVersion())), bson.M{
		"id": id,
	}), &consumer); err != nil {
		if errors.Is(rerror.ErrNotFound, err) && (v == nil || v.IsRef(version.Latest)) {
			return nil, nil
		}
		return nil, err
	}
	return &consumer.Result[0], nil
}