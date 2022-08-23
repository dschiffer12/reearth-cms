package schema

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Schema struct {
	id        ID
	workspace id.WorkspaceID
	fields    []*Field
}

func (s *Schema) ID() ID {
	return s.id
}

func (s *Schema) Workspace() id.WorkspaceID {
	return s.workspace
}

func (s *Schema) SetWorkspace(workspace id.WorkspaceID) {
	s.workspace = workspace
}

func (s *Schema) HasField(f FieldID) bool {
	if s == nil {
		return false
	}
	return lo.SomeBy(s.fields, func(g *Field) bool { return g.ID().Equal(f) })
}

func (s *Schema) AddField(f Field) {
	if s.fields == nil {
		s.fields = []*Field{}
	}
	if s.HasField(f.ID()) {
		return
	}
	s.fields = append(s.fields, &f)
}

func (s *Schema) Field(fId FieldID) *Field {
	if s == nil || s.fields == nil {
		return nil
	}
	f, found := lo.Find(s.fields, func(f *Field) bool { return f.id.Equal(fId) })
	if found {
		return f
	}
	return nil
}

func (s *Schema) Fields() []*Field {
	if s == nil {
		return nil
	}
	return slices.Clone(s.fields)
}

func (s *Schema) RemoveField(fid FieldID) {
	if s == nil {
		return
	}

	for i, field := range s.fields {
		if field.id.Equal(fid) {
			s.fields = slices.Delete(s.fields, i, i+1)
			return
		}
	}
}