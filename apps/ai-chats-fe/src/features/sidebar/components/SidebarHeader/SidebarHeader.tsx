import { useNavigate } from 'react-router-dom';

import { Flex, Heading, Tooltip } from '@radix-ui/themes';

import { IconButton } from 'components';
import { ChatIcon } from 'components/icons';

import { CloseSidebarButton } from '../CloseSidebarButton';

export const SidebarHeader = () => {
  const navigate = useNavigate();

  return (
    <Flex
      align="center"
      justify="between"
      gap="2"
      pb={{
        initial: '4',
        sm: '5',
      }}
      px="2"
    >
      <CloseSidebarButton />
      <Heading as="h2" align="center" size="5" weight="bold">
        AI Chats
      </Heading>
      <Tooltip content="Start a new chat">
        <IconButton
          aria-label="Start a new chat"
          highContrast
          m="2"
          onClick={() => navigate('/app/new-chat')}
          size="3"
          variant="ghost"
        >
          <ChatIcon size="24" weight="light" />
        </IconButton>
      </Tooltip>
    </Flex>
  );
};
