import { Flex, Heading, Text, Button } from '@radix-ui/themes';

import { OllamaModel } from 'types';

interface ModelCardProps {
  model: OllamaModel;
}

export const ModelCard: React.FC<ModelCardProps> = ({ model }) => {
  return (
    <Flex direction="column" gap="2" key={`${model.model}`} width="100%">
      <Heading as="h2" size="3">{`${model.model}`}</Heading>
      <Text>{model.description}</Text>
      <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
        <Button variant="soft">Pull</Button>
      </Flex>
    </Flex>
  );
};
