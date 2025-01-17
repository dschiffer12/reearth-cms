package http

import (
	"context"

	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
)

type TaskController struct {
	usecase interfaces.Asset
}

type NotifyInput struct {
	Type    string                         `json:"type"`
	AssetID string                         `json:"assetId"`
	Status  *asset.ArchiveExtractionStatus `json:"status"`
	Task    *NotifyInputTask               `json:"-"`
}

type NotifyInputTask struct {
	TaskID string
	Status string
}

func NewTaskController(uc interfaces.Asset) *TaskController {
	return &TaskController{usecase: uc}
}

func (tc *TaskController) Notify(ctx context.Context, input NotifyInput) error {
	if input.Task != nil && input.Task.Status == "EXPIRED" {
		return tc.usecase.RetryDecompression(ctx, input.Task.TaskID)
	}

	aID, err := id.AssetIDFrom(input.AssetID)
	if err != nil {
		return err
	}

	_, err = tc.usecase.UpdateFiles(ctx, aID, input.Status, adapter.Operator(ctx))
	if err != nil {
		return err
	}

	return nil
}
