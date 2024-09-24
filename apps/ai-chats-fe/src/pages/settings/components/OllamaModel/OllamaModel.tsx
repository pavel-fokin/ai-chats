import { Flex, Heading, Text } from '@radix-ui/themes';

import { Progress } from 'shared/components';
import * as types from 'types';

import { DeleteOllamaModelDialog } from './DeleteOllamaModelDialog';

interface OllamaModelProps {
  model: types.OllamaModel;
}

export const OllamaModel: React.FC<OllamaModelProps> = ({ model }) => {
  return (
    <Flex direction="column" gap="2" key={`${model.model}`} width="100%">
      <Heading as="h2" size="3">{`${model.model}`}</Heading>
      <Text>{model.description}</Text>
      {model.isPulling ? (
        <Progress mt="4" mb="8" />
      ) : (
        <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
          <DeleteOllamaModelDialog model={model} />
        </Flex>
      )}
    </Flex>
  );
};
