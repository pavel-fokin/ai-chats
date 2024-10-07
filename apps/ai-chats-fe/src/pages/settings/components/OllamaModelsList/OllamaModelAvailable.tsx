import { Flex, Heading, Text } from '@radix-ui/themes';

import { OllamaModel } from 'types';

import { DeleteOllamaModelDialog } from './DeleteOllamaModelDialog';

interface OllamaModelAvailableProps {
  model: OllamaModel;
}

export const OllamaModelAvailable = ({
  model,
}: OllamaModelAvailableProps): JSX.Element => {
  return (
    <Flex direction="column" gap="2" key={`${model.model}`} width="100%">
      <Heading as="h2" size="3">{`${model.model}`}</Heading>
      <Text>{model.description}</Text>
      <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
        <DeleteOllamaModelDialog model={model} />
      </Flex>
    </Flex>
  );
};
