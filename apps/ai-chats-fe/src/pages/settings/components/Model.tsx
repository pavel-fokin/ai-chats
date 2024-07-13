import { AlertDialog, Button, Flex, Heading, Text } from '@radix-ui/themes';

import { useDeleteOllamaModel } from 'hooks';
import { OllamaModel } from 'types';

type ModelProps = {
  model: OllamaModel;
};

export const Model: React.FC<ModelProps> = ({ model }) => {
  return (
    <Flex
      direction="column"
      gap="2"
      key={`${model.name}:${model.tag}`}
      width="100%"
    >
      <Heading as="h2">{`${model.name}:${model.tag}`}</Heading>
      <Text>
        Meta Llama 3: The most capable openly available LLM to date 8B.
      </Text>
      <Flex align="center" justify="end" flexGrow="1" mt="4" gap="4">
        <Delete model={model} />
      </Flex>
    </Flex>
  );
};

type DeleteProps = {
  model: OllamaModel;
};

const Delete: React.FC<DeleteProps> = ({ model }) => {
  const deleteModel = useDeleteOllamaModel();

  const handleDelete = (model: string) => {
    deleteModel.mutate(model);
  };

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button size="2" variant="soft" loading={deleteModel.isPending}>
          Delete
        </Button>
      </AlertDialog.Trigger>
      <AlertDialog.Content maxWidth="450px">
        <AlertDialog.Title>Delete model?</AlertDialog.Title>
        <AlertDialog.Description size="2">
          {`Are you sure? The model `}
          <Text weight="bold">{`${model.name}:${model.tag}`}</Text>
          {` will not be available locally.`}
        </AlertDialog.Description>

        <Flex gap="4" mt="4" align="center" justify="end">
          <AlertDialog.Cancel>
            <Button variant="ghost">Cancel</Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button
              variant="solid"
              color="tomato"
              onClick={() => handleDelete(model.name)}
            >
              Delete model
            </Button>
          </AlertDialog.Action>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  );
};
