import { Flex } from '@radix-ui/themes';
import { useEffect } from 'react';

import { OllamaModel } from 'types';

import { ModelItem } from './ModelItem';

interface ModelsListProps {
  models: OllamaModel[];
  selectedModel: OllamaModel | null;
  setSelectedModel: (model: OllamaModel) => void;
}

export const ModelsList: React.FC<ModelsListProps> = ({
  models: ollamaModels,
  selectedModel,
  setSelectedModel,
}) => {
  const onModelClick = (model: OllamaModel) => {
    setSelectedModel(model);
  };

  // Set the selected model to the first model in the list.
  useEffect(() => {
    if (!selectedModel && ollamaModels.length > 0) {
      setSelectedModel(ollamaModels[0]);
    }
  }, [ollamaModels, selectedModel, setSelectedModel]);

  return (
    <Flex direction="column" gap="3">
      {ollamaModels.map((model) => (
        <ModelItem
          key={`${model.model}}`}
          model={model}
          isSelected={
            selectedModel ? selectedModel.model === model.model : false
          }
          onClick={onModelClick}
        />
      ))}
    </Flex>
  );
};
