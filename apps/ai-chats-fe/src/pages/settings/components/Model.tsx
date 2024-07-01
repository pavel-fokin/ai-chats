import { Button, Flex, Heading, Text } from '@radix-ui/themes';

import { useDeleteOllamaModel } from 'hooks';
import { OllamaModel } from 'types';

type ModelProps = {
  model: OllamaModel;
};

export const Model: React.FC<ModelProps> = ({ model }) => {
  const deleteModel = useDeleteOllamaModel();

  const handleDelete = (model: string) => {
    deleteModel.mutate(model);
  };

  return (
    <Flex direction="column" gap="2" key={model.id} width="100%">
      <Heading as="h2">
        {model.name}:{model.tag}
      </Heading>
      <Text>
        Meta Llama 3: The most capable openly available LLM to date 8B.
      </Text>
      <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
        <Button
          size="2"
          variant="soft"
          onClick={() => handleDelete(model.name)}
          loading={deleteModel.isPending}
        >
          Delete
        </Button>
      </Flex>
    </Flex>
  );
};
