import { Flex, Heading, Text } from '@radix-ui/themes';

import { Button } from 'components';
import * as types from 'types';

import { DeleteDialog } from './DeleteDialog';

interface OllamaModelProps {
  model: types.OllamaModel;
}

export const OllamaModel: React.FC<OllamaModelProps> = ({ model }) => {
  return (
    <Flex direction="column" gap="2" key={`${model.model}`} width="100%">
      <Heading as="h2" size="3">{`${model.model}`}</Heading>
      <Text>{model.description}</Text>
      <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
        <DeleteDialog model={model} />
        <Button variant="soft" onClick={() => {}}>
          Chat
        </Button>
      </Flex>
    </Flex>
  );
};
