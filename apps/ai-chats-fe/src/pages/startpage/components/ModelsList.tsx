import { useEffect } from 'react';
import { Box, Flex, Heading, Code } from '@radix-ui/themes';

import { OllamaModel } from 'types';

import { Model } from './Model';

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
    <>
      <Box mb="4">
        <Heading as="h2" size="6" weight="bold">
          Choose a model to chat <Code variant="ghost">[*_*]</Code>
        </Heading>
      </Box>
      <Flex direction="column" gap="3">
        {ollamaModels.map((model) => (
          <Model
            key={`${model.name}}`}
            model={model}
            isSelected={
              selectedModel ? selectedModel.name === model.name : false
            }
            onClick={onModelClick}
          />
        ))}
      </Flex>
    </>
  );
};
