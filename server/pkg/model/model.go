package model

import (
	"errors"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/key"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

type Model struct {
	id          ID
	project     ProjectID
	schema      SchemaID
	name        string
	description string
	key         key.Key
	public      bool
	updatedAt   time.Time
}

func (p *Model) ID() ID {
	return p.id
}

func (p *Model) Schema() SchemaID {
	return p.schema
}

func (p *Model) Project() ProjectID {
	return p.project
}

func (p *Model) Name() string {
	return p.name
}

func (p *Model) SetName(name string) {
	p.name = name
}

func (p *Model) Description() string {
	return p.description
}

func (p *Model) SetDescription(description string) {
	p.description = description
}

func (p *Model) Key() key.Key {
	return p.key
}

func (p *Model) SetKey(key key.Key) error {
	if !key.IsValid() {
		return ErrInvalidKey
	}
	p.key = key
	return nil
}

func (p *Model) Public() bool {
	return p.public
}

func (p *Model) SetPublic(public bool) {
	p.public = public
}

func (p *Model) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Model) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

func (p *Model) CreatedAt() time.Time {
	return p.id.Timestamp()
}