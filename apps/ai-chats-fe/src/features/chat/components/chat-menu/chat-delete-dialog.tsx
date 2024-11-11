import { useNavigate } from 'react-router-dom';

import { useChat, useDeleteChat, useInvalidateChats } from '@/hooks';
import { AlertDialog, Button } from '@/components/ui';

import styles from './chat-delete-dialog.module.css';

interface ChatDeleteDialogProps {
  chatId: string;
  open: boolean;
  setOpen: (open: boolean) => void;
}

export const ChatDeleteDialog = ({
  chatId,
  open,
  setOpen,
}: ChatDeleteDialogProps) => {
  const chat = useChat(chatId);
  const deleteChat = useDeleteChat();
  const invalidateChats = useInvalidateChats();
  const navigate = useNavigate();

  const handleDelete = () => {
    deleteChat.mutate(chatId, {
      onSuccess: () => {
        invalidateChats();
        navigate('/app');
      },
    });
  };

  const handleCancel = () => {
    setOpen(false);
  };

  return (
    <AlertDialog.Root open={open} onOpenChange={setOpen}>
      <AlertDialog.Content maxWidth="450px">
        <AlertDialog.Title>Delete chat?</AlertDialog.Title>
        <AlertDialog.Description size="2">
          <p>
            This will delete <strong>{`${chat.data?.title}`}</strong>.
          </p>
        </AlertDialog.Description>
        <div className={styles.chatDeleteDialog__buttons}>
          <AlertDialog.Cancel>
            <Button variant="ghost" onClick={handleCancel}>
              Cancel
            </Button>
          </AlertDialog.Cancel>
          <AlertDialog.Action>
            <Button variant="solid" color="tomato" onClick={handleDelete}>
              Delete chat
            </Button>
          </AlertDialog.Action>
        </div>
      </AlertDialog.Content>
    </AlertDialog.Root>
  );
};
