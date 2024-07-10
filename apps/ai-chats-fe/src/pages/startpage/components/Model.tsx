import { Badge, Card, Flex, Heading, Text } from '@radix-ui/themes';

import { OllamaModel } from 'types';

interface ModelProps {
  model: OllamaModel;
  isSelected: boolean;
  onClick: (model: OllamaModel) => void;
}

export const Model: React.FC<ModelProps> = ({ model, isSelected, onClick }) => {
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
