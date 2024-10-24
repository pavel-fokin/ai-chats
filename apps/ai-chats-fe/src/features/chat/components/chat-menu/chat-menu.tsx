import { useState } from 'react';

import { DropdownMenu, Flex } from '@radix-ui/themes';

import { AIActionIcon, DeleteIcon } from 'components/icons';
import { useGenerateChatTitle } from 'hooks';

import { useChatMenu } from '../../hooks';
import { DeleteChatDialog } from './delete-chat-dialog';

interface ChatMenuProps {
  chatId: string;
  trigger: React.ReactNode;
}

export const ChatMenu = ({ chatId, trigger }: ChatMenuProps): JSX.Element => {
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const { isOpen, setIsOpen } = useChatMenu();
  const generateChatTitle = useGenerateChatTitle();

  return (
    <DropdownMenu.Root open={isOpen} onOpenChange={setIsOpen}>
      <DropdownMenu.Trigger>
        <div>{trigger}</div>
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
            <DeleteIcon size="16" /> Delete chat
          </Flex>
        </DropdownMenu.Item>
      </DropdownMenu.Content>
      <DeleteChatDialog
        chatId={chatId}
        open={isDeleteDialogOpen}
        setOpen={setIsDeleteDialogOpen}
        onCancelClick={() => setIsDeleteDialogOpen(false)}
      />
    </DropdownMenu.Root>
  );
};
