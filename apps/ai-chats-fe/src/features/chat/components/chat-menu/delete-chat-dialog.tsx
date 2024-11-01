import { AlertDialog, Flex, Strong, Text } from '@radix-ui/themes';
import { useNavigate } from 'react-router-dom';

import { useChat, useDeleteChat } from 'hooks';
import { Button } from 'components/ui';

interface DeleteChatDialogProps {
  chatId: string;
  open: boolean;
  setOpen: (open: boolean) => void;
  onCancelClick: () => void;
}

export const DeleteChatDialog: React.FC<DeleteChatDialogProps> = ({
  chatId,
  open,
  setOpen,
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
    <AlertDialog.Root open={open} onOpenChange={setOpen}>
      <AlertDialog.Content maxWidth="450px">
        <AlertDialog.Title>Delete chat?</AlertDialog.Title>
        <AlertDialog.Description size="2">
          <Flex direction="column" gap="2">
            <Text>
              This will delete <Strong>{`${chat.data?.title}`}</Strong>.
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
