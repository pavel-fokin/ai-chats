import { useEffect } from 'react';

import { Badge, Box, Card, Flex, Heading, Text } from '@radix-ui/themes';

import { OllamaModel } from 'types';

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
          Choose a model to chat 🤖
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

interface ModelProps {
  model: OllamaModel;
  isSelected: boolean;
  onClick: (model: OllamaModel) => void;
}

const Model: React.FC<ModelProps> = ({ model, isSelected, onClick }) => {
  return (
    <Card asChild onClick={() => onClick(model)}>
      <button>
        <Flex direction="column" gap="2" width="100%">
          <Flex align="center" justify="between">
            <Heading as="h3" size="4">
              {model.name}
            </Heading>
            {isSelected && <Badge color="jade">Choosen</Badge>}
          </Flex>
          <Text>
            Meta Llama 3: The most capable openly available LLM to date 8B.
          </Text>
        </Flex>
      </button>
    </Card>
  );
};
