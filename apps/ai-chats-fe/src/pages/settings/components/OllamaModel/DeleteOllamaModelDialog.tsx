import { AlertDialog, Flex, Strong, Text } from '@radix-ui/themes';

import { useDeleteOllamaModel } from 'hooks';
import { Button } from 'shared/components';
import { OllamaModel } from 'types';

interface DeleteOllamaModelDialogProps {
  model: OllamaModel;
}

export const DeleteOllamaModelDialog: React.FC<
  DeleteOllamaModelDialogProps
> = ({ model }) => {
  const deleteModel = useDeleteOllamaModel();

  const handleDelete = (model: string) => {
    deleteModel.mutate(model);
  };

  return (
    <AlertDialog.Root>
      <AlertDialog.Trigger>
        <Button
          size="2"
          variant="ghost"
          color="tomato"
          loading={deleteModel.isPending}
        >
          Delete
        </Button>
      </AlertDialog.Trigger>
      <AlertDialog.Content maxWidth="450px">
        <AlertDialog.Title>Delete model?</AlertDialog.Title>
        <AlertDialog.Description size="2">
          <Text>
            This will delete local model <Strong>{`${model.model}`}</Strong>.
          </Text>
        </AlertDialog.Description>

        <Flex gap="4" mt="4" align="center" justify="end">
          <AlertDialog.Cancel>
            <Button variant="ghost">Cancel</Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button
              variant="solid"
              color="tomato"
              onClick={() => handleDelete(model.model)}
            >
              Delete model
            </Button>
          </AlertDialog.Action>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  );
};
