import { useState } from 'react';

import { AIActionIcon, DeleteIcon } from '@/components/icons';
import { DropdownMenu } from '@/components/ui';
import { useGenerateChatTitle } from '@/hooks';

import { ChatDeleteDialog } from './chat-delete-dialog';
import styles from './chat-menu.module.css';

interface ChatMenuProps {
  chatId: string;
  trigger: React.ReactNode;
}

export const ChatMenu = ({ chatId, trigger }: ChatMenuProps) => {
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const generateChatTitle = useGenerateChatTitle();

  const handleGenerateChatTitle = () => {
    generateChatTitle.mutate(chatId);
  };
  const handleOpenDeleteDialog = () => {
    setIsDeleteDialogOpen(true);
  };

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger>
        <div>{trigger}</div>
      </DropdownMenu.Trigger>
      <DropdownMenu.Content className={styles.chatMenu__content} highContrast>
        <DropdownMenu.Item onClick={handleGenerateChatTitle}>
          <div className={styles.chatMenu__itemContainer}>
            <AIActionIcon size="16" /> Generate title
          </div>
        </DropdownMenu.Item>
        <DropdownMenu.Separator />
        <DropdownMenu.Item color="tomato" onClick={handleOpenDeleteDialog}>
          <div className={styles.chatMenu__itemContainer}>
            <DeleteIcon size="16" /> Delete chat
          </div>
        </DropdownMenu.Item>
      </DropdownMenu.Content>
      {/* Delete dialog */}
      <ChatDeleteDialog
        chatId={chatId}
        open={isDeleteDialogOpen}
        setOpen={setIsDeleteDialogOpen}
      />
    </DropdownMenu.Root>
  );
};
