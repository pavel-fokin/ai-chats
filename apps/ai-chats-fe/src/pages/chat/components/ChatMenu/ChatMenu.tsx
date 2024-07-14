import { Button, DropdownMenu, Flex } from '@radix-ui/themes';

import { ConfigurationIcon, DeleteIcon } from 'components/ui/icons';
import { useChat } from 'hooks';

interface ChatMenuProps {
  chatId?: string;
  onDeleteClick: () => void;
}

export const ChatMenu: React.FC<ChatMenuProps> = ({
  chatId,
  onDeleteClick,
}) => {
  const chat = useChat(chatId);

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
      <DropdownMenu.Content style={{ minWidth: '128px' }}>
        <DropdownMenu.Item shortcut="">
          <Flex direction="row" align="center" justify="between" width="100%">
            Configure <ConfigurationIcon size="16" />
          </Flex>
        </DropdownMenu.Item>
        <DropdownMenu.Separator />
        <DropdownMenu.Item color="tomato" onClick={onDeleteClick}>
          <Flex direction="row" align="center" justify="between" width="100%">
            Delete <DeleteIcon size="16" />
          </Flex>
        </DropdownMenu.Item>
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  );
};
