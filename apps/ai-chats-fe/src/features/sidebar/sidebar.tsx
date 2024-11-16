import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Flex, Text } from '@radix-ui/themes';

import { OllamaLibraryIcon } from '@/components/icons';

import {
  ChatsList,
  Link,
  MenuList,
  SidebarHeader,
  SignOutButton,
} from './components';

import styles from './sidebar.module.css';

export const Sidebar = () => {
  return (
    <Flex direction="column" gap="2" height="100%" justify="between">
      <Flex direction="column" flexGrow="1" height="100%">
        <SidebarHeader />
        <NavigationMenu.Root
          className={styles.NavigationMenuRoot}
          orientation="vertical"
        >
          <div className={styles.sidebarScrollable}>
            <MenuList ariaLabel="Chats list">
              <ChatsList />
            </MenuList>
          </div>

          <MenuList ariaLabel="Settings list">
            <NavigationMenu.Item>
              <Link aria-label="Ollama library" to="/app/ollama-library">
                <Flex align="center" gap="3">
                  <OllamaLibraryIcon size={24} />{' '}
                  <Text size="2">Ollama Library</Text>
                </Flex>
              </Link>
            </NavigationMenu.Item>
            <NavigationMenu.Item>
              <SignOutButton />
            </NavigationMenu.Item>
          </MenuList>
        </NavigationMenu.Root>
      </Flex>
    </Flex>
  );
};
