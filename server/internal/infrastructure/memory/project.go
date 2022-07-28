package memory

import (
	"context"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/rerror"
	"github.com/reearth/reearth-cms/server/pkg/util"
	"github.com/samber/lo"
)

type Project struct {
	data *util.SyncMap[id.ProjectID, *project.Project]
	f    repo.WorkspaceFilter
	now  *util.TimeNow
	err  error
}

func NewProject() repo.Project {
	return &Project{
		data: &util.SyncMap[id.ProjectID, *project.Project]{},
		now:  &util.TimeNow{},
	}
}

func (r *Project) Filtered(f repo.WorkspaceFilter) repo.Project {
	return &Project{
		data: r.data,
		f:    r.f.Merge(f),
		now:  &util.TimeNow{},
	}
}

func (r *Project) FindByWorkspace(_ context.Context, wid id.WorkspaceID, _ *usecase.Pagination) (project.List, *usecase.PageInfo, error) {
	if r.err != nil {
		return nil, nil, r.err
	}

	// TODO: implement pagination

	if !r.f.CanRead(wid) {
		return nil, nil, nil
	}

	result := project.List(r.data.FindAll(func(_ id.ProjectID, v *project.Project) bool {
		return v.Workspace() == wid
	})).SortByID()

	var startCursor, endCursor *usecase.Cursor
	if len(result) > 0 {
		startCursor = lo.ToPtr(usecase.Cursor(result[0].ID().String()))
		endCursor = lo.ToPtr(usecase.Cursor(result[len(result)-1].ID().String()))
	}

	return result, usecase.NewPageInfo(
		len(result),
		startCursor,
		endCursor,
		true,
		true,
	), nil
}

func (r *Project) FindByIDs(_ context.Context, ids id.ProjectIDList) (project.List, error) {
	if r.err != nil {
		return nil, r.err
	}

	result := r.data.FindAll(func(k id.ProjectID, v *project.Project) bool {
		return ids.Has(k) && r.f.CanRead(v.Workspace())
	})

	return project.List(result).SortByID(), nil
}

func (r *Project) FindByID(_ context.Context, pid id.ProjectID) (*project.Project, error) {
	if r.err != nil {
		return nil, r.err
	}

	p := r.data.Find(func(k id.ProjectID, v *project.Project) bool {
		return k == pid && r.f.CanRead(v.Workspace())
	})

	if p != nil {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Project) FindByPublicName(_ context.Context, name string) (*project.Project, error) {
	if r.err != nil {
		return nil, r.err
	}

	if name == "" {
		return nil, nil
	}

	p := r.data.Find(func(_ id.ProjectID, v *project.Project) bool {
		return v.Alias() == name && r.f.CanRead(v.Workspace())
	})

	if p != nil {
		return p, nil
	}
	return nil, rerror.ErrNotFound
}

func (r *Project) CountByWorkspace(_ context.Context, workspace id.WorkspaceID) (c int, err error) {
	if r.err != nil {
		return 0, r.err
	}

	if !r.f.CanRead(workspace) {
		return 0, nil
	}

	return r.data.CountAll(func(_ id.ProjectID, v *project.Project) bool {
		return v.Workspace() == workspace
	}), nil
}

func (r *Project) Save(_ context.Context, p *project.Project) error {
	if r.err != nil {
		return r.err
	}

	if !r.f.CanWrite(p.Workspace()) {
		return repo.ErrOperationDenied
	}

	p.SetUpdatedAt(r.now.Now())
	r.data.Store(p.ID(), p)
	return nil
}

func (r *Project) Remove(_ context.Context, id id.ProjectID) error {
	if r.err != nil {
		return r.err
	}

	if p, ok := r.data.Load(id); ok && r.f.CanWrite(p.Workspace()) {
		r.data.Delete(id)
		return nil
	}
	return rerror.ErrNotFound
}

func MockProjectNow(r repo.Project, t time.Time) func() {
	return r.(*Project).now.Mock(t)
}

func SetProjectError(r repo.Project, err error) {
	r.(*Project).err = err
}