import { useContext } from 'react';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Box, Button, Flex, Separator, Text, TextField, Heading } from '@radix-ui/themes';

import { ChatIcon, SettingsIcon, SignOutIcon, SearchIcon } from 'components/ui/icons';
import { AuthContext, SidebarContext } from 'contexts';
import { useChats } from 'hooks';

import 'styles/styles.css';
import styles from './Sidebar.module.css';

type LinkProps = {
  to: string;
  children: React.ReactNode;
};

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
        {...props}
        onClick={toggleSidebar}
      >
        {children}
      </NavLink>
    </NavigationMenu.Link>
  );
};

export const Sidebar = () => {
  const navigate = useNavigate();

  const { chats, createChat } = useChats();
  const { signout } = useContext(AuthContext);

  const handleNewChat = async () => {
    createChat.mutate(void 0, {
      onSuccess: ({ data }) => {
        navigate(`/app/chats/${data.chat.id}`);
      },
    });
  };

  const handleSignOut = () => {
    signout();
  };

  return (
    <Flex direction="column" gap="2" height="100%" justify="between">
      <Flex direction="column" flexGrow="1">
        <Box className="mobile-hidden" px="2" pb={{
          initial: '4',
          sm: '2',
        }}>
          <Heading as="h2" align="center" size='3' weight="regular">AI Chats</Heading>
        </Box>
        <Flex className="mobile-hidden" direction="column" flexGrow="1">
          <Button size="3" variant="ghost" m="4" onClick={handleNewChat}>
            <ChatIcon size={16} />
            New chat
          </Button>
        </Flex>
        <Box px="2" pb='4'>
          <TextField.Root size="3" placeholder='Find chat...'>
            <TextField.Slot>
              <SearchIcon size={16} />
            </TextField.Slot>
          </TextField.Root>
        </Box>


        <NavigationMenu.Root
          orientation="vertical"
          className={styles.NavigationMenuRoot}
        >
          <NavigationMenu.List className={styles.NavigationMenuList}>
            {!!chats &&
              chats?.map((chat) => (
                <NavigationMenu.Item key={chat.id}>
                  <Link to={`/app/chats/${chat.id}`}>{chat.title}</Link>
                </NavigationMenu.Item>
              ))}
          </NavigationMenu.List>
          <NavigationMenu.List className={styles.NavigationMenuList}>
            <NavigationMenu.Item>
              <Link to="/app/settings">
                <Flex align="center" gap="3">
                  <SettingsIcon size={24} /> <Text size="3">Models</Text>
                </Flex>
              </Link>
            </NavigationMenu.Item>
          </NavigationMenu.List>
        </NavigationMenu.Root>
      </Flex>

      <Flex direction="column" gap="2">
        <Separator size="4" />
        <Button size="3" variant="ghost" m="3" onClick={handleSignOut}>
          <SignOutIcon width={16} />
          Sign out
        </Button>
      </Flex>
    </Flex>
  );
};
