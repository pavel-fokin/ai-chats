import { useState } from 'react';

import { DropdownMenu, Flex } from '@radix-ui/themes';

import { Button } from 'components';
import { AIActionIcon, DeleteIcon } from 'components/icons';
import { useChat, useGenerateChatTitle } from 'hooks';

import { DeleteDialog } from './DeleteDialog';

interface ChatMenuProps {
  chatId: string;
}

export const ChatMenu: React.FC<ChatMenuProps> = ({ chatId }) => {
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const chat = useChat(chatId);
  const generateChatTitle = useGenerateChatTitle();

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger>
        <Button variant="ghost" size="3" highContrast>
          <span
            style={{
              overflow: 'hidden',
              textOverflow: 'ellipsis',
              whiteSpace: 'nowrap',
              maxWidth: '192px',
            }}
          >
            {chat.data?.title || 'Chat'}
          </span>
          <DropdownMenu.TriggerIcon />
        </Button>
      </DropdownMenu.Trigger>
      <DropdownMenu.Content style={{ minWidth: '192px' }}>
        <DropdownMenu.Item onClick={() => generateChatTitle.mutate(chatId)}>
          <Flex
            direction="row"
            align="center"
            justify="start"
            gap="4"
            width="100%"
          >
            <AIActionIcon size="16" /> Generate title
          </Flex>
        </DropdownMenu.Item>
        <DropdownMenu.Separator />
        <DropdownMenu.Item
          color="tomato"
          onClick={() => setIsDeleteDialogOpen(true)}
        >
          <Flex
            direction="row"
            align="center"
            justify="start"
            gap="4"
            width="100%"
          >
            <DeleteIcon size="16" /> Delete
          </Flex>
        </DropdownMenu.Item>
      </DropdownMenu.Content>
      <DeleteDialog
        chatId={chatId}
        open={isDeleteDialogOpen}
        setOpen={setIsDeleteDialogOpen}
        onCancelClick={() => setIsDeleteDialogOpen(false)}
      />
    </DropdownMenu.Root>
  );
};
