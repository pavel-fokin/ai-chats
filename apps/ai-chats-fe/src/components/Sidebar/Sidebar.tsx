import { useContext } from 'react';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Flex, Heading, Text } from '@radix-ui/themes';

import { Tooltip, IconButton } from 'components';
import { ChatIcon, SettingsIcon, SignOutIcon } from 'components/ui/icons';
import { SidebarContext } from 'contexts';
import { useChats } from 'hooks';

import 'styles/styles.css';
import styles from './Sidebar.module.css';

interface LinkProps {
  to: string;
  children: React.ReactNode;
}

const Link: React.FC<LinkProps> = ({ to, children, ...props }) => {
  const { pathname } = useLocation();
  const isActive = to === pathname;

  const { toggleSidebar } = useContext(SidebarContext);

  const classNames = isActive
    ? `${styles.NavigationMenuLink} ${styles.NavigationMenuLinkActive}`
    : styles.NavigationMenuLink;

  return (
    <NavigationMenu.Link asChild active={isActive}>
      <NavLink
        to={to}
        className={classNames}
        onClick={toggleSidebar}
        {...props}
      >
        {children}
      </NavLink>
    </NavigationMenu.Link>
  );
};

export const Sidebar = () => {
  const navigate = useNavigate();
  const chats = useChats();

  const handleNewChat = () => {
    navigate('/app');
  };

  return (
    <Flex direction="column" gap="2" height="100%" justify="between">
      <Flex direction="column" flexGrow="1">
        <Flex
          className="mobile-hidden"
          align="center"
          justify="between"
          gap="2"
          px="2"
          pb={{
            initial: '4',
            sm: '5',
          }}
        >
          <Heading as="h2" align="center" size="5" weight="bold">
            AI Chats
          </Heading>
          <Tooltip content="Start a new chat">
            <IconButton
              variant="ghost"
              size="3"
              m="2"
              highContrast
              onClick={handleNewChat}
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
              <Tooltip  key={`key-${chat.id}`} content={chat.title}>
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
