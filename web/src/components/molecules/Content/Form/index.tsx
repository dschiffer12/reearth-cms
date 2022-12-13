import styled from "@emotion/styled";
import { useCallback, useEffect } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import InputNumber from "@reearth-cms/components/atoms/InputNumber";
import MarkdownInput from "@reearth-cms/components/atoms/Markdown";
import PageHeader from "@reearth-cms/components/atoms/PageHeader";
import Select from "@reearth-cms/components/atoms/Select";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import AssetItem from "@reearth-cms/components/molecules/Common/Form/AssetItem";
import MultiValueField from "@reearth-cms/components/molecules/Common/MultiValueField";
import MultiValueAsset from "@reearth-cms/components/molecules/Common/MultiValueField/MultiValueAsset";
import MultiValueSelect from "@reearth-cms/components/molecules/Common/MultiValueField/MultiValueSelect";
import FieldTitle from "@reearth-cms/components/molecules/Content/Form/FieldTitle";
import { ItemField } from "@reearth-cms/components/molecules/Content/types";
import { FieldType, Model } from "@reearth-cms/components/molecules/Schema/types";
import { useT } from "@reearth-cms/i18n";
import { validateURL } from "@reearth-cms/utils/regex";

export interface Props {
  itemId?: string;
  initialFormValues: any;
  loading: boolean;
  model?: Model;
  assetList: Asset[];
  fileList: UploadFile[];
  loadingAssets: boolean;
  uploading: boolean;
  uploadModalVisibility: boolean;
  uploadUrl: string;
  uploadType: UploadType;
  onUploadModalCancel: () => void;
  setUploadUrl: (url: string) => void;
  setUploadType: (type: UploadType) => void;
  onItemCreate: (data: { schemaId: string; fields: ItemField[] }) => Promise<void>;
  onItemUpdate: (data: { itemId: string; fields: ItemField[] }) => Promise<void>;
  onBack: (modelId?: string) => void;
  onAssetsCreate: (files: UploadFile[]) => Promise<(Asset | undefined)[]>;
  onAssetCreateFromUrl: (url: string) => Promise<Asset | undefined>;
  onAssetsReload: () => void;
  onAssetSearchTerm: (term?: string | undefined) => void;
  setFileList: (fileList: UploadFile<File>[]) => void;
  setUploadModalVisibility: (visible: boolean) => void;
  onNavigateToAsset: (asset: Asset) => void;
}

const ContentForm: React.FC<Props> = ({
  itemId,
  model,
  initialFormValues,
  loading,
  assetList,
  fileList,
  loadingAssets,
  uploading,
  uploadModalVisibility,
  uploadUrl,
  uploadType,
  onUploadModalCancel,
  setUploadUrl,
  setUploadType,
  onAssetsCreate,
  onAssetCreateFromUrl,
  onItemCreate,
  onItemUpdate,
  onBack,
  onAssetsReload,
  onAssetSearchTerm,
  setFileList,
  setUploadModalVisibility,
  onNavigateToAsset,
}) => {
  const t = useT();
  const { Option } = Select;
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue(initialFormValues);
  }, [form, initialFormValues]);

  const handleBack = useCallback(() => {
    onBack(model?.id);
  }, [onBack, model]);

  const handleSubmit = useCallback(async () => {
    try {
      const values = await form.validateFields();
      const fields: { schemaFieldId: string; type: FieldType; value: string }[] = [];
      for (const [key, value] of Object.entries(values)) {
        fields.push({
          value: (value || "") as string,
          schemaFieldId: key,
          type: model?.schema.fields.find(field => field.id === key)?.type as FieldType,
        });
      }
      if (!itemId) {
        await onItemCreate?.({ schemaId: model?.schema.id as string, fields });
      } else {
        await onItemUpdate?.({ itemId: itemId as string, fields });
      }
    } catch (info) {
      console.log("Validate Failed:", info);
    }
  }, [form, model?.schema.fields, model?.schema.id, itemId, onItemCreate, onItemUpdate]);

  return (
    <StyledForm form={form} layout="vertical" initialValues={initialFormValues}>
      <PageHeader
        title={model?.name}
        onBack={handleBack}
        extra={
          <Button htmlType="submit" onClick={handleSubmit} loading={loading}>
            {t("Save")}
          </Button>
        }
      />
      <FormItemsWrapper>
        {model?.schema.fields.map(field =>
          field.type === "TextArea" ? (
            <Form.Item
              extra={field.description}
              rules={[
                {
                  required: field.required,
                  message: t("Please input field!"),
                },
              ]}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  rows={3}
                  showCount
                  maxLength={field.typeProperty.maxLength ?? false}
                  FieldInput={TextArea}
                />
              ) : (
                <TextArea rows={3} showCount maxLength={field.typeProperty.maxLength ?? false} />
              )}
            </Form.Item>
          ) : field.type === "MarkdownText" ? (
            <Form.Item
              extra={field.description}
              rules={[
                {
                  required: field.required,
                  message: t("Please input field!"),
                },
              ]}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  maxLength={field.typeProperty.maxLength ?? false}
                  FieldInput={MarkdownInput}
                />
              ) : (
                <MarkdownInput maxLength={field.typeProperty.maxLength ?? false} />
              )}
            </Form.Item>
          ) : field.type === "Integer" ? (
            <Form.Item
              extra={field.description}
              rules={[
                {
                  required: field.required,
                  message: t("Please input field!"),
                },
              ]}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  type="number"
                  min={field.typeProperty.min}
                  max={field.typeProperty.max}
                  FieldInput={InputNumber}
                />
              ) : (
                <InputNumber
                  type="number"
                  min={field.typeProperty.min}
                  max={field.typeProperty.max}
                />
              )}
            </Form.Item>
          ) : field.type === "Asset" ? (
            <Form.Item
              extra={field.description}
              rules={[
                {
                  required: field.required,
                  message: t("Please input field!"),
                },
              ]}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueAsset
                  assetList={assetList}
                  fileList={fileList}
                  loadingAssets={loadingAssets}
                  uploading={uploading}
                  uploadModalVisibility={uploadModalVisibility}
                  uploadUrl={uploadUrl}
                  uploadType={uploadType}
                  onUploadModalCancel={onUploadModalCancel}
                  setUploadUrl={setUploadUrl}
                  setUploadType={setUploadType}
                  onAssetsCreate={onAssetsCreate}
                  onAssetCreateFromUrl={onAssetCreateFromUrl}
                  onAssetsReload={onAssetsReload}
                  onAssetSearchTerm={onAssetSearchTerm}
                  setFileList={setFileList}
                  setUploadModalVisibility={setUploadModalVisibility}
                  onNavigateToAsset={onNavigateToAsset}
                />
              ) : (
                <AssetItem
                  assetList={assetList}
                  fileList={fileList}
                  loadingAssets={loadingAssets}
                  uploading={uploading}
                  uploadModalVisibility={uploadModalVisibility}
                  uploadUrl={uploadUrl}
                  uploadType={uploadType}
                  onUploadModalCancel={onUploadModalCancel}
                  setUploadUrl={setUploadUrl}
                  setUploadType={setUploadType}
                  onAssetsCreate={onAssetsCreate}
                  onAssetCreateFromUrl={onAssetCreateFromUrl}
                  onAssetsReload={onAssetsReload}
                  onAssetSearchTerm={onAssetSearchTerm}
                  setFileList={setFileList}
                  setUploadModalVisibility={setUploadModalVisibility}
                  onNavigateToAsset={onNavigateToAsset}
                />
              )}
            </Form.Item>
          ) : field.type === "Select" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueSelect selectedValues={field.typeProperty?.values} />
              ) : (
                <Select>
                  {field.typeProperty?.values?.map((value: string) => (
                    <Option key={value} value={value}>
                      {value}
                    </Option>
                  ))}
                </Select>
              )}
            </Form.Item>
          ) : field.type === "URL" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}
              rules={[
                {
                  required: field.required,
                  message: t("Please input field!"),
                },
                {
                  message: "URL is not valid",
                  validator: async (_, value) => {
                    if (
                      Array.isArray(value) &&
                      value.some((valueItem: string) => !validateURL(valueItem))
                    )
                      return Promise.reject();
                    else if (!Array.isArray(value) && !validateURL(value) && value?.length > 0)
                      return Promise.reject();
                    return Promise.resolve();
                  },
                },
              ]}>
              {field.multiple ? (
                <MultiValueField
                  showCount={true}
                  maxLength={field.typeProperty.maxLength ?? 500}
                  FieldInput={Input}
                />
              ) : (
                <Input showCount={true} maxLength={field.typeProperty.maxLength ?? 500} />
              )}
            </Form.Item>
          ) : (
            <Form.Item
              extra={field.description}
              rules={[
                {
                  required: field.required,
                  message: t("Please input field!"),
                },
              ]}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  showCount={true}
                  maxLength={field.typeProperty.maxLength ?? 500}
                  FieldInput={Input}
                />
              ) : (
                <Input showCount={true} maxLength={field.typeProperty.maxLength ?? 500} />
              )}
            </Form.Item>
          ),
        )}
      </FormItemsWrapper>
    </StyledForm>
  );
};

const StyledForm = styled(Form)`
  padding: 16px;
  width: 100%;
  height: 100%;
  overflow-y: auto;
  background: #fff;
`;

const FormItemsWrapper = styled.div`
  width: 50%;
  @media (max-width: 1200px) {
    width: 100%;
  }
`;

export default ContentForm;
