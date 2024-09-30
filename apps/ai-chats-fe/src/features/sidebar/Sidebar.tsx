import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Flex, Text } from '@radix-ui/themes';

import { SettingsIcon } from 'shared/components/icons';

import { Link, SidebarHeader, ChatsList, MenuList, SignOutButton } from './components';
import styles from './Sidebar.module.css';

export const Sidebar = () => {
  return (
    <Flex direction="column" gap="2" height="100%" justify="between">
      <Flex direction="column" flexGrow="1">
        <SidebarHeader />
        <NavigationMenu.Root
          className={styles.NavigationMenuRoot}
          orientation="vertical"
        >
          <MenuList>
            <ChatsList />
          </MenuList>

          <MenuList>
            <NavigationMenu.Item>
              <Link aria-label="Ollama settings" to="/app/settings">
                <Flex align="center" gap="3">
                  <SettingsIcon size={24} />{' '}
                  <Text size="2">Ollama Settings</Text>
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
