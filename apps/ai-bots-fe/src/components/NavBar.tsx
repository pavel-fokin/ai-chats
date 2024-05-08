import { useContext } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Button, Flex, Separator } from '@radix-ui/themes';
import { IconLogout, IconMessagePlus } from '@tabler/icons-react';

import { AuthContext } from 'contexts';
import { useAuth, useChats } from 'hooks';

import styles from './Navbar.module.css';

export function Navbar() {
    const navigate = useNavigate();
    const { chatId } = useParams<{ chatId: string }>();

    const { chats, createChat } = useChats();
    const { signOut } = useAuth();
    const { setIsAuthenticated } = useContext(AuthContext);

    const handleNewChat = async () => {
        createChat.mutateAsync().then((data) => {
            navigate(`/app/chats/${data.data.id}`);
        }).catch((error) => {
            console.error('Failed to create a chat', error);
        });
    }

    const handleSignOut = () => {
        signOut();
        setIsAuthenticated(false);
        navigate('/app/signin');
    }

    return (
        <Flex direction="column" gap="2" height="100%" justify="between">
            <Flex direction="column">
                <Button size="3" variant="ghost" m="3" onClick={handleNewChat}>
                    <IconMessagePlus size={16} />Start a new chat
                </Button>
                <NavigationMenu.Root orientation="vertical">
                    <NavigationMenu.List className={styles.NavigationMenuList}>
                        {chats.map((chat) => (
                            <NavigationMenu.Item key={chat.id}>
                                <NavigationMenu.Link
                                    href={`/app/chats/${chat.id}`}
                                    className={
                                        styles.NavigationMenuLink
                                        + ' '
                                        + (chatId === chat.id ? styles.NavigationMenuLinkActive : '')
                                    }
                                >
                                    {chat.id}
                                </NavigationMenu.Link>
                            </NavigationMenu.Item>
                        ))}
                    </NavigationMenu.List>
                </NavigationMenu.Root>
            </Flex>

            <Flex direction="column" gap="2">
                <Separator size="4" />
                <Button size="3" variant="ghost" m="3" onClick={handleSignOut}>
                    <IconLogout size={16} />Sign Out
                </Button>
            </Flex>
        </Flex>

    )
}