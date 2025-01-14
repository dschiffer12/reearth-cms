import { Viewer as CesiumViewer } from "cesium";
import { ComponentProps, useEffect } from "react";

import ResiumViewer from "@reearth-cms/components/atoms/ResiumViewer";
import Settings from "@reearth-cms/components/molecules/Asset/Viewers/Settings";
import { compressedFileFormats } from "@reearth-cms/components/molecules/Common/Asset";
import { WorkspaceSettings } from "@reearth-cms/components/molecules/Workspace/types";
import { getExtension } from "@reearth-cms/utils/file";

import Cesium3dTileSetComponent from "./Cesium3dTileSetComponent";

type Props = {
  viewerProps?: ComponentProps<typeof ResiumViewer>;
  url: string;
  onGetViewer: (viewer: CesiumViewer | undefined) => void;
  setAssetUrl: (url: string) => void;
  workspaceSettings?: WorkspaceSettings;
};

const Geo3dViewer: React.FC<Props> = ({
  viewerProps,
  url,
  setAssetUrl,
  onGetViewer,
  workspaceSettings,
}) => {
  useEffect(() => {
    const assetExtension = getExtension(url);
    if (compressedFileFormats.includes(assetExtension)) {
      const nameRegex = /\.\w+$/;
      const base = url.replace(nameRegex, "");
      setAssetUrl(`${base}/tileset.json`);
    }
  }, [setAssetUrl, url]);

  return (
    <ResiumViewer {...viewerProps} onGetViewer={onGetViewer}>
      <Cesium3dTileSetComponent url={url} />
      <Settings workspaceSettings={workspaceSettings} />
    </ResiumViewer>
  );
};

export default Geo3dViewer;
