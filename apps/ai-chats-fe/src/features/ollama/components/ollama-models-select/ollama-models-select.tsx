import { Flex } from '@radix-ui/themes';
import { useEffect } from 'react';

import { OllamaModel } from 'types';

import { OllamaModelItem } from './ollama-model-item';

interface OllamaModelsSelectProps {
  models: OllamaModel[];
  selectedModel: OllamaModel | null;
  setSelectedModel: (model: OllamaModel) => void;
}

export const OllamaModelsSelect = ({
  models: ollamaModels,
  selectedModel,
  setSelectedModel,
}: OllamaModelsSelectProps): JSX.Element => {
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
        <OllamaModelItem
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
