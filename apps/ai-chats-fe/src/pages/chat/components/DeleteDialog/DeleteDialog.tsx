import { AlertDialog, Button, Flex, Strong, Text } from '@radix-ui/themes';
import { useNavigate } from 'react-router-dom';

import { useChat, useDeleteChat } from 'hooks';

interface DeleteDialogProps {
  chatId: string;
  open: boolean;
  onCancelClick: () => void;
}

export const DeleteDialog: React.FC<DeleteDialogProps> = ({
  chatId,
  open,
  onCancelClick,
}) => {
  const chat = useChat(chatId);
  const deleteChat = useDeleteChat();
  const navigate = useNavigate();

  const handleDelete = () => {
    deleteChat.mutate(chatId, {
      onSuccess: () => {
        navigate('/app');
      },
    });
  };

  return (
    <AlertDialog.Root open={open}>
      <AlertDialog.Content maxWidth="450px">
        <AlertDialog.Title>Delete chat?</AlertDialog.Title>
        <AlertDialog.Description size="2">
          <Flex direction="column" gap="2">
            <Text>
              The chat "<Strong>{`${chat.data?.title}`}</Strong>" will be
              deleted.
            </Text>
          </Flex>
        </AlertDialog.Description>
        <Flex gap="4" mt="4" align="center" justify="end">
          <AlertDialog.Cancel>
            <Button variant="ghost" onClick={onCancelClick}>
              Cancel
            </Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button variant="solid" color="tomato" onClick={handleDelete}>
              Delete chat
            </Button>
          </AlertDialog.Action>
        </Flex>
      </AlertDialog.Content>
    </AlertDialog.Root>
  );
};