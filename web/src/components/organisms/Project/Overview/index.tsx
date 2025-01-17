import ProjectOverviewMolecule from "@reearth-cms/components/molecules/ProjectOverview";
import ModelDeletionModal from "@reearth-cms/components/molecules/Schema/ModelDeletionModal";
import ModelFormModal from "@reearth-cms/components/molecules/Schema/ModelFormModal";

import useHooks from "./hooks";

const ProjectOverview: React.FC = () => {
  const {
    currentProject,
    models,
    isModelKeyAvailable,
    modelModalShown,
    selectedModel,
    modelDeletionModalShown,
    handleSchemaNavigation,
    handleContentNavigation,
    handleModelKeyCheck,
    handleModelModalOpen,
    handleModelModalReset,
    handleModelCreate,
    handleModelDeletionModalOpen,
    handleModelDeletionModalClose,
    handleModelUpdateModalOpen,
    handleModelDelete,
    handleModelUpdate,
  } = useHooks();

  return (
    <>
      <ProjectOverviewMolecule
        projectName={currentProject?.name}
        projectDescription={currentProject?.description}
        models={models}
        onSchemaNavigate={handleSchemaNavigation}
        onContentNavigate={handleContentNavigation}
        onModelModalOpen={handleModelModalOpen}
        onModelDeletionModalOpen={handleModelDeletionModalOpen}
        onModelUpdateModalOpen={handleModelUpdateModalOpen}
      />
      <ModelFormModal
        model={selectedModel}
        isKeyAvailable={isModelKeyAvailable}
        onModelKeyCheck={handleModelKeyCheck}
        open={modelModalShown}
        onClose={handleModelModalReset}
        onCreate={handleModelCreate}
        onUpdate={handleModelUpdate}
      />
      <ModelDeletionModal
        model={selectedModel}
        open={modelDeletionModalShown}
        onDelete={handleModelDelete}
        onClose={handleModelDeletionModalClose}
      />
    </>
  );
};

export default ProjectOverview;
