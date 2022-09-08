import { useParams } from "react-router-dom";

import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import AssetListHeader from "@reearth-cms/components/molecules/Asset/AssetList/AssetListHeader";
import AssetListTable from "@reearth-cms/components/molecules/Asset/AssetList/AssetListTable";

import useHooks from "./hooks";

const AssetList: React.FC = () => {
  const { workspaceId, projectId } = useParams();
  const {
    assetList,
    createAssets,
    assetsPerPage,
    navigate,
    handleSearchTerm,
    selection,
    setSelection,
    fileList,
    setFileList,
    uploading,
    setUploading,
    uploadModalVisibility,
    setUploadModalVisibility,
  } = useHooks(projectId);

  const displayUploadModal = () => {
    setUploadModalVisibility(true);
  };
  const hideUploadModal = () => {
    setUploadModalVisibility(false);
  };

  const handleUpload = () => {
    setUploading(true);
    createAssets(fileList)
      .catch(error => {
        // TODO: notification
        console.log(error);
      })
      .finally(() => {
        setUploading(false);
        setFileList([]);
        hideUploadModal();
      });
  };

  const handleEdit = (asset: Asset) => {
    navigate(`/workspaces/${workspaceId}/${projectId}/asset/${asset.id}`);
  };

  return (
    <>
      <AssetListHeader
        title="Asset"
        subTitle="This is a subtitle"
        fileList={fileList}
        uploading={uploading}
        uploadModalVisibility={uploadModalVisibility}
        displayUploadModal={displayUploadModal}
        hideUploadModal={hideUploadModal}
        handleUpload={handleUpload}
        setFileList={setFileList}
      />
      <AssetListTable
        assetList={assetList}
        assetsPerPage={assetsPerPage}
        handleEdit={handleEdit}
        handleSearchTerm={handleSearchTerm}
        selection={selection}
        setSelection={setSelection}
      />
    </>
  );
};

export default AssetList;