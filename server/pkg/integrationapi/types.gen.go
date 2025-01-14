// Package integrationapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package integrationapi

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/project"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for AssetArchiveExtractionStatus.
const (
	Done       AssetArchiveExtractionStatus = "done"
	Failed     AssetArchiveExtractionStatus = "failed"
	InProgress AssetArchiveExtractionStatus = "in_progress"
	Pending    AssetArchiveExtractionStatus = "pending"
)

// Defines values for AssetPreviewType.
const (
	Geo        AssetPreviewType = "geo"
	Geo3dTiles AssetPreviewType = "geo_3d_Tiles"
	GeoMvt     AssetPreviewType = "geo_mvt"
	Image      AssetPreviewType = "image"
	ImageSvg   AssetPreviewType = "image_svg"
	Model3d    AssetPreviewType = "model_3d"
	Unknown    AssetPreviewType = "unknown"
)

// Defines values for AssetEmbedding.
const (
	All   AssetEmbedding = "all"
	False AssetEmbedding = "false"
	True  AssetEmbedding = "true"
)

// Defines values for CommentAuthorType.
const (
	Integrtaion CommentAuthorType = "integrtaion"
	User        CommentAuthorType = "user"
)

// Defines values for RefOrVersionRef.
const (
	RefOrVersionRefLatest RefOrVersionRef = "latest"
	RefOrVersionRefPublic RefOrVersionRef = "public"
)

// Defines values for ValueType.
const (
	ValueTypeAsset     ValueType = "asset"
	ValueTypeBool      ValueType = "bool"
	ValueTypeCheckbox  ValueType = "checkbox"
	ValueTypeDate      ValueType = "date"
	ValueTypeGroup     ValueType = "group"
	ValueTypeInteger   ValueType = "integer"
	ValueTypeMarkdown  ValueType = "markdown"
	ValueTypeReference ValueType = "reference"
	ValueTypeRichText  ValueType = "richText"
	ValueTypeSelect    ValueType = "select"
	ValueTypeTag       ValueType = "tag"
	ValueTypeText      ValueType = "text"
	ValueTypeTextArea  ValueType = "textArea"
	ValueTypeUrl       ValueType = "url"
)

// Defines values for RefParam.
const (
	RefParamLatest RefParam = "latest"
	RefParamPublic RefParam = "public"
)

// Defines values for SortDirParam.
const (
	SortDirParamAsc  SortDirParam = "asc"
	SortDirParamDesc SortDirParam = "desc"
)

// Defines values for SortParam.
const (
	SortParamCreatedAt SortParam = "createdAt"
	SortParamUpdatedAt SortParam = "updatedAt"
)

// Defines values for ItemGetParamsRef.
const (
	ItemGetParamsRefLatest ItemGetParamsRef = "latest"
	ItemGetParamsRefPublic ItemGetParamsRef = "public"
)

// Defines values for ItemFilterParamsSort.
const (
	ItemFilterParamsSortCreatedAt ItemFilterParamsSort = "createdAt"
	ItemFilterParamsSortUpdatedAt ItemFilterParamsSort = "updatedAt"
)

// Defines values for ItemFilterParamsDir.
const (
	ItemFilterParamsDirAsc  ItemFilterParamsDir = "asc"
	ItemFilterParamsDirDesc ItemFilterParamsDir = "desc"
)

// Defines values for ItemFilterParamsRef.
const (
	ItemFilterParamsRefLatest ItemFilterParamsRef = "latest"
	ItemFilterParamsRefPublic ItemFilterParamsRef = "public"
)

// Defines values for ItemFilterWithProjectParamsSort.
const (
	ItemFilterWithProjectParamsSortCreatedAt ItemFilterWithProjectParamsSort = "createdAt"
	ItemFilterWithProjectParamsSortUpdatedAt ItemFilterWithProjectParamsSort = "updatedAt"
)

// Defines values for ItemFilterWithProjectParamsDir.
const (
	ItemFilterWithProjectParamsDirAsc  ItemFilterWithProjectParamsDir = "asc"
	ItemFilterWithProjectParamsDirDesc ItemFilterWithProjectParamsDir = "desc"
)

// Defines values for ItemFilterWithProjectParamsRef.
const (
	Latest ItemFilterWithProjectParamsRef = "latest"
	Public ItemFilterWithProjectParamsRef = "public"
)

// Defines values for AssetFilterParamsSort.
const (
	AssetFilterParamsSortCreatedAt AssetFilterParamsSort = "createdAt"
	AssetFilterParamsSortUpdatedAt AssetFilterParamsSort = "updatedAt"
)

// Defines values for AssetFilterParamsDir.
const (
	AssetFilterParamsDirAsc  AssetFilterParamsDir = "asc"
	AssetFilterParamsDirDesc AssetFilterParamsDir = "desc"
)

// Asset defines model for asset.
type Asset struct {
	ArchiveExtractionStatus *AssetArchiveExtractionStatus `json:"archiveExtractionStatus,omitempty"`
	ContentType             *string                       `json:"contentType,omitempty"`
	CreatedAt               time.Time                     `json:"createdAt"`
	File                    *File                         `json:"file,omitempty"`
	Id                      id.AssetID                    `json:"id"`
	Name                    *string                       `json:"name,omitempty"`
	PreviewType             *AssetPreviewType             `json:"previewType,omitempty"`
	ProjectId               id.ProjectID                  `json:"projectId"`
	TotalSize               *float32                      `json:"totalSize,omitempty"`
	UpdatedAt               time.Time                     `json:"updatedAt"`
	Url                     string                        `json:"url"`
}

// AssetArchiveExtractionStatus defines model for Asset.ArchiveExtractionStatus.
type AssetArchiveExtractionStatus string

// AssetPreviewType defines model for Asset.PreviewType.
type AssetPreviewType string

// AssetEmbedding defines model for assetEmbedding.
type AssetEmbedding string

// Comment defines model for comment.
type Comment struct {
	AuthorId   *any               `json:"authorId,omitempty"`
	AuthorType *CommentAuthorType `json:"authorType,omitempty"`
	Content    *string            `json:"content,omitempty"`
	CreatedAt  *time.Time         `json:"createdAt,omitempty"`
	Id         *id.CommentID      `json:"id,omitempty"`
}

// CommentAuthorType defines model for Comment.AuthorType.
type CommentAuthorType string

// Field defines model for field.
type Field struct {
	Group *id.ItemGroupID `json:"group,omitempty"`
	Id    *id.FieldID     `json:"id,omitempty"`
	Key   *string         `json:"key,omitempty"`
	Type  *ValueType      `json:"type,omitempty"`
	Value *interface{}    `json:"value,omitempty"`
}

// File defines model for file.
type File struct {
	Children    *[]File  `json:"children,omitempty"`
	ContentType *string  `json:"contentType,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Path        *string  `json:"path,omitempty"`
	Size        *float32 `json:"size,omitempty"`
}

// Item defines model for item.
type Item struct {
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	Fields         *[]Field   `json:"fields,omitempty"`
	Id             *id.ItemID `json:"id,omitempty"`
	IsMetadata     *bool      `json:"isMetadata,omitempty"`
	MetadataItemId *id.ItemID `json:"metadataItemId,omitempty"`
	ModelId        *string    `json:"modelId,omitempty"`
	OriginalItemId *id.ItemID `json:"originalItemId,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

// Model defines model for model.
type Model struct {
	CreatedAt        *time.Time    `json:"createdAt,omitempty"`
	Description      *string       `json:"description,omitempty"`
	Id               *id.ModelID   `json:"id,omitempty"`
	Key              *string       `json:"key,omitempty"`
	LastModified     *time.Time    `json:"lastModified,omitempty"`
	MetadataSchemaId *id.SchemaID  `json:"metadataSchemaId,omitempty"`
	Name             *string       `json:"name,omitempty"`
	ProjectId        *id.ProjectID `json:"projectId,omitempty"`
	Public           *bool         `json:"public,omitempty"`
	SchemaId         *id.SchemaID  `json:"schemaId,omitempty"`
	UpdatedAt        *time.Time    `json:"updatedAt,omitempty"`
}

// RefOrVersion defines model for refOrVersion.
type RefOrVersion struct {
	Ref     *RefOrVersionRef    `json:"ref,omitempty"`
	Version *openapi_types.UUID `json:"version,omitempty"`
}

// RefOrVersionRef defines model for RefOrVersion.Ref.
type RefOrVersionRef string

// Schema defines model for schema.
type Schema struct {
	TitleField *id.FieldID    `json:"TitleField,omitempty"`
	CreatedAt  *time.Time     `json:"createdAt,omitempty"`
	Fields     *[]SchemaField `json:"fields,omitempty"`
	Id         *id.SchemaID   `json:"id,omitempty"`
	ProjectId  *id.ProjectID  `json:"projectId,omitempty"`
}

// SchemaField defines model for schemaField.
type SchemaField struct {
	Id       *id.FieldID `json:"id,omitempty"`
	Key      *string     `json:"key,omitempty"`
	Required *bool       `json:"required,omitempty"`
	Type     *ValueType  `json:"type,omitempty"`
}

// TagResponse defines model for tagResponse.
type TagResponse struct {
	Color *string   `json:"color,omitempty"`
	Id    *id.TagID `json:"id,omitempty"`
	Name  *string   `json:"name,omitempty"`
}

// ValueType defines model for valueType.
type ValueType string

// Version defines model for version.
type Version struct {
	Parents *[]openapi_types.UUID `json:"parents,omitempty"`
	Refs    *[]openapi_types.UUID `json:"refs,omitempty"`
	Version *openapi_types.UUID   `json:"version,omitempty"`
}

// VersionedItem defines model for versionedItem.
type VersionedItem struct {
	CreatedAt       *time.Time            `json:"createdAt,omitempty"`
	Fields          *[]Field              `json:"fields,omitempty"`
	Id              *id.ItemID            `json:"id,omitempty"`
	IsMetadata      *bool                 `json:"isMetadata,omitempty"`
	MetadataFields  *[]Field              `json:"metadataFields,omitempty"`
	ModelId         *string               `json:"modelId,omitempty"`
	Parents         *[]openapi_types.UUID `json:"parents,omitempty"`
	ReferencedItems *[]VersionedItem      `json:"referencedItems,omitempty"`
	Refs            *[]string             `json:"refs,omitempty"`
	UpdatedAt       *time.Time            `json:"updatedAt,omitempty"`
	Version         *openapi_types.UUID   `json:"version,omitempty"`
}

// AssetIdParam defines model for assetIdParam.
type AssetIdParam = id.AssetID

// AssetParam defines model for assetParam.
type AssetParam = AssetEmbedding

// CommentIdParam defines model for commentIdParam.
type CommentIdParam = id.CommentID

// ItemIdParam defines model for itemIdParam.
type ItemIdParam = id.ItemID

// ModelIdOrKeyParam defines model for modelIdOrKeyParam.
type ModelIdOrKeyParam = model.IDOrKey

// ModelIdParam defines model for modelIdParam.
type ModelIdParam = id.ModelID

// PageParam defines model for pageParam.
type PageParam = int

// PerPageParam defines model for perPageParam.
type PerPageParam = int

// ProjectIdOrAliasParam defines model for projectIdOrAliasParam.
type ProjectIdOrAliasParam = project.IDOrAlias

// ProjectIdParam defines model for projectIdParam.
type ProjectIdParam = id.ProjectID

// RefParam defines model for refParam.
type RefParam string

// SortDirParam defines model for sortDirParam.
type SortDirParam string

// SortParam defines model for sortParam.
type SortParam string

// AssetCommentCreateJSONBody defines parameters for AssetCommentCreate.
type AssetCommentCreateJSONBody struct {
	Content *string `json:"content,omitempty"`
}

// AssetCommentUpdateJSONBody defines parameters for AssetCommentUpdate.
type AssetCommentUpdateJSONBody struct {
	Content *string `json:"content,omitempty"`
}

// ItemGetParams defines parameters for ItemGet.
type ItemGetParams struct {
	// Ref Used to select a ref or ver
	Ref *ItemGetParamsRef `form:"ref,omitempty" json:"ref,omitempty"`

	// Asset Specifies whether asset data are embedded in the results
	Asset *AssetParam `form:"asset,omitempty" json:"asset,omitempty"`
}

// ItemGetParamsRef defines parameters for ItemGet.
type ItemGetParamsRef string

// ItemUpdateJSONBody defines parameters for ItemUpdate.
type ItemUpdateJSONBody struct {
	Asset          *AssetEmbedding `json:"asset,omitempty"`
	Fields         *[]Field        `json:"fields,omitempty"`
	MetadataFields *[]Field        `json:"metadataFields,omitempty"`
}

// ItemCommentCreateJSONBody defines parameters for ItemCommentCreate.
type ItemCommentCreateJSONBody struct {
	Content *string `json:"content,omitempty"`
}

// ItemCommentUpdateJSONBody defines parameters for ItemCommentUpdate.
type ItemCommentUpdateJSONBody struct {
	Content *string `json:"content,omitempty"`
}

// ItemFilterParams defines parameters for ItemFilter.
type ItemFilterParams struct {
	// Sort Used to define the order of the response list
	Sort *ItemFilterParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Dir Used to define the order direction of the response list, will be ignored if the order is not presented
	Dir *ItemFilterParamsDir `form:"dir,omitempty" json:"dir,omitempty"`

	// Page Used to select the page
	Page *PageParam `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Used to select the page
	PerPage *PerPageParam `form:"perPage,omitempty" json:"perPage,omitempty"`

	// Ref Used to select a ref or ver
	Ref *ItemFilterParamsRef `form:"ref,omitempty" json:"ref,omitempty"`

	// Asset Specifies whether asset data are embedded in the results
	Asset *AssetParam `form:"asset,omitempty" json:"asset,omitempty"`
}

// ItemFilterParamsSort defines parameters for ItemFilter.
type ItemFilterParamsSort string

// ItemFilterParamsDir defines parameters for ItemFilter.
type ItemFilterParamsDir string

// ItemFilterParamsRef defines parameters for ItemFilter.
type ItemFilterParamsRef string

// ItemCreateJSONBody defines parameters for ItemCreate.
type ItemCreateJSONBody struct {
	Fields         *[]Field `json:"fields,omitempty"`
	MetadataFields *[]Field `json:"metadataFields,omitempty"`
}

// ModelFilterParams defines parameters for ModelFilter.
type ModelFilterParams struct {
	// Page Used to select the page
	Page *PageParam `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Used to select the page
	PerPage *PerPageParam `form:"perPage,omitempty" json:"perPage,omitempty"`
}

// ItemFilterWithProjectParams defines parameters for ItemFilterWithProject.
type ItemFilterWithProjectParams struct {
	// Sort Used to define the order of the response list
	Sort *ItemFilterWithProjectParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Dir Used to define the order direction of the response list, will be ignored if the order is not presented
	Dir *ItemFilterWithProjectParamsDir `form:"dir,omitempty" json:"dir,omitempty"`

	// Page Used to select the page
	Page *PageParam `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Used to select the page
	PerPage *PerPageParam `form:"perPage,omitempty" json:"perPage,omitempty"`

	// Ref Used to select a ref or ver
	Ref *ItemFilterWithProjectParamsRef `form:"ref,omitempty" json:"ref,omitempty"`

	// Asset Specifies whether asset data are embedded in the results
	Asset *AssetParam `form:"asset,omitempty" json:"asset,omitempty"`
}

// ItemFilterWithProjectParamsSort defines parameters for ItemFilterWithProject.
type ItemFilterWithProjectParamsSort string

// ItemFilterWithProjectParamsDir defines parameters for ItemFilterWithProject.
type ItemFilterWithProjectParamsDir string

// ItemFilterWithProjectParamsRef defines parameters for ItemFilterWithProject.
type ItemFilterWithProjectParamsRef string

// ItemCreateWithProjectJSONBody defines parameters for ItemCreateWithProject.
type ItemCreateWithProjectJSONBody struct {
	Fields         *[]Field `json:"fields,omitempty"`
	MetadataFields *[]Field `json:"metadataFields,omitempty"`
}

// AssetFilterParams defines parameters for AssetFilter.
type AssetFilterParams struct {
	// Sort Used to define the order of the response list
	Sort *AssetFilterParamsSort `form:"sort,omitempty" json:"sort,omitempty"`

	// Dir Used to define the order direction of the response list, will be ignored if the order is not presented
	Dir *AssetFilterParamsDir `form:"dir,omitempty" json:"dir,omitempty"`

	// Page Used to select the page
	Page *PageParam `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Used to select the page
	PerPage *PerPageParam `form:"perPage,omitempty" json:"perPage,omitempty"`
}

// AssetFilterParamsSort defines parameters for AssetFilter.
type AssetFilterParamsSort string

// AssetFilterParamsDir defines parameters for AssetFilter.
type AssetFilterParamsDir string

// AssetCreateJSONBody defines parameters for AssetCreate.
type AssetCreateJSONBody struct {
	SkipDecompression *bool   `json:"skipDecompression"`
	Url               *string `json:"url,omitempty"`
}

// AssetCreateMultipartBody defines parameters for AssetCreate.
type AssetCreateMultipartBody struct {
	File              *openapi_types.File `json:"file,omitempty"`
	SkipDecompression *bool               `json:"skipDecompression,omitempty"`
}

// AssetCommentCreateJSONRequestBody defines body for AssetCommentCreate for application/json ContentType.
type AssetCommentCreateJSONRequestBody AssetCommentCreateJSONBody

// AssetCommentUpdateJSONRequestBody defines body for AssetCommentUpdate for application/json ContentType.
type AssetCommentUpdateJSONRequestBody AssetCommentUpdateJSONBody

// ItemUpdateJSONRequestBody defines body for ItemUpdate for application/json ContentType.
type ItemUpdateJSONRequestBody ItemUpdateJSONBody

// ItemCommentCreateJSONRequestBody defines body for ItemCommentCreate for application/json ContentType.
type ItemCommentCreateJSONRequestBody ItemCommentCreateJSONBody

// ItemCommentUpdateJSONRequestBody defines body for ItemCommentUpdate for application/json ContentType.
type ItemCommentUpdateJSONRequestBody ItemCommentUpdateJSONBody

// ItemCreateJSONRequestBody defines body for ItemCreate for application/json ContentType.
type ItemCreateJSONRequestBody ItemCreateJSONBody

// ItemCreateWithProjectJSONRequestBody defines body for ItemCreateWithProject for application/json ContentType.
type ItemCreateWithProjectJSONRequestBody ItemCreateWithProjectJSONBody

// AssetCreateJSONRequestBody defines body for AssetCreate for application/json ContentType.
type AssetCreateJSONRequestBody AssetCreateJSONBody

// AssetCreateMultipartRequestBody defines body for AssetCreate for multipart/form-data ContentType.
type AssetCreateMultipartRequestBody AssetCreateMultipartBody
