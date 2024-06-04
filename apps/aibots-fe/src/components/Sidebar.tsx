import { useContext } from 'react';
import { useNavigate, NavLink } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Button, Flex, Separator } from '@radix-ui/themes';
import { ChatText as ChatIcon, SignOut as SignOutIcon } from "@phosphor-icons/react";

import { AuthContext, SidebarContext } from 'contexts';
import { useChats } from 'hooks';

import styles from './Sidebar.module.css';

export const Sidebar = () => {
    const navigate = useNavigate();

    const { toggleSidebar } = useContext(SidebarContext);

    const { chats, createChat } = useChats();
    const { signout } = useContext(AuthContext);

    const handleNewChat = async () => {
        createChat.mutate(void 0, {
            onSuccess: (data) => {
                navigate(`/app/chats/${data.data.id}`);
            },
        });
    }

    const handleSignOut = () => {
        signout();
    }

    return (
        <Flex direction="column" gap="2" height="100%" justify="between">
            <Flex direction="column">
                <Button size="3" variant="ghost" m="4" onClick={handleNewChat}>
                    <ChatIcon width={16} height={16} />Start a new chat
                </Button>
                <NavigationMenu.Root orientation="vertical">
                    <NavigationMenu.List className={styles.NavigationMenuList}>
                        {!!chats && chats.data.chats?.map((chat) => (
                            <NavigationMenu.Item key={chat.id}>
                                {/* <NavigationMenu.Link
                                    asChild
                                > */}
                                    <NavLink
                                        className={({isActive}) =>
                                            isActive ? styles.NavigationMenuLink + ' ' + styles.NavigationMenuLinkActive: styles.NavigationMenuLink
                                        }
                                        to={`/app/chats/${chat.id}`}
                                        onClick={toggleSidebar}
                                    >
                                        {chat.title}
                                    </NavLink>
                                {/* </NavigationMenu.Link> */}
                            </NavigationMenu.Item>
                        ))}
                    </NavigationMenu.List>
                </NavigationMenu.Root>
            </Flex>

            <Flex direction="column" gap="2">
                <Separator size="4" />
                <Button size="3" variant="ghost" m="3" onClick={handleSignOut}>
                    <SignOutIcon width={16} />Sign out
                </Button>
            </Flex>
        </Flex>

    )
}