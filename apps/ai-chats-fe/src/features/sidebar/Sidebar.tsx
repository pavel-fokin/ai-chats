import { useNavigate } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Flex, Heading, Text } from '@radix-ui/themes';

import { IconButton, Tooltip } from 'components';
import { ChatIcon, SettingsIcon, SignOutIcon } from 'components/icons';
import { useChats } from 'hooks';

import { CloseSidebarButton, Link } from './components';
import styles from './Sidebar.module.css';

export const Sidebar = () => {
  const navigate = useNavigate();
  const chats = useChats();

  return (
    <Flex direction="column" gap="2" height="100%" justify="between">
      <Flex direction="column" flexGrow="1">
        <Flex
          align="center"
          justify="between"
          gap="2"
          px="2"
          pb={{
            initial: '4',
            sm: '5',
          }}
        >
          <CloseSidebarButton />
          <Heading as="h2" align="center" size="5" weight="bold">
            AI Chats
          </Heading>
          <Tooltip content="Start a new chat">
            <IconButton
              variant="ghost"
              size="3"
              m="2"
              highContrast
              onClick={() => navigate('/app/new-chat')}
              aria-label="New chat"
            >
              <ChatIcon size="24" weight="light" />
            </IconButton>
          </Tooltip>
        </Flex>
        <NavigationMenu.Root
          orientation="vertical"
          className={styles.NavigationMenuRoot}
        >
          <NavigationMenu.List className={styles.NavigationMenuList}>
            {chats.data?.map((chat) => (
              <Tooltip key={`key-${chat.id}`} content={chat.title}>
                <NavigationMenu.Item key={chat.id}>
                  <Link to={`/app/chats/${chat.id}`}>{chat.title}</Link>
                </NavigationMenu.Item>
              </Tooltip>
            ))}
          </NavigationMenu.List>

          <NavigationMenu.List className={styles.NavigationMenuList}>
            <NavigationMenu.Item>
              <Link aria-label="Ollama settings" to="/app/settings">
                <Flex align="center" gap="3">
                  <SettingsIcon size={24} />{' '}
                  <Text size="2">Ollama Settings</Text>
                </Flex>
              </Link>
            </NavigationMenu.Item>
            <NavigationMenu.Item>
              <Link aria-label="Sign out" to="/app/signout">
                <Flex align="center" gap="3">
                  <SignOutIcon size={24} /> <Text size="2">Sign Out</Text>
                </Flex>
              </Link>
            </NavigationMenu.Item>
          </NavigationMenu.List>
        </NavigationMenu.Root>
      </Flex>
    </Flex>
  );
};
