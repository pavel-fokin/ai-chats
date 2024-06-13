import { Button, DropdownMenu, Flex } from '@radix-ui/themes';
import { useNavigate } from 'react-router-dom';

import { DeleteIcon, ConfigurationIcon } from 'components/ui/icons';
import { useChat, useDeleteChat } from 'hooks';

type ChatMenuProps = {
  chatId: string;
};

export const ChatMenu = ({ chatId }: ChatMenuProps) => {
  const navigate = useNavigate();

  const { data: chat } = useChat(chatId);
  const deleteChat = useDeleteChat(chatId);

  const handleDelete = () => {
    deleteChat.mutate(void 0, {
      onSuccess: () => {
        navigate('/app');
      },
    });
  };

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
            {chat?.title || 'Chat'}
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
        <DropdownMenu.Item color="tomato" onClick={handleDelete}>
          <Flex direction="row" align="center" justify="between" width="100%">
            Delete <DeleteIcon size="16" />
          </Flex>
        </DropdownMenu.Item>
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  );
};
