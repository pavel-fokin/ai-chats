import { useContext } from 'react';
import { useNavigate, useParams, Link } from 'react-router-dom';

import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Button, Flex, Separator } from '@radix-ui/themes';
import { IconLogout, IconMessagePlus } from '@tabler/icons-react';

import { AuthContext } from 'contexts';
import { useChats } from 'hooks';

import styles from './Navbar.module.css';

type NavbarProps = {
    open: (open: boolean) => void;
}

export function Navbar({ open }: NavbarProps) {
    const navigate = useNavigate();
    const { chatId } = useParams<{ chatId: string }>();

    const { chats, createChat } = useChats();
    const { setIsAuthenticated, signout } = useContext(AuthContext);

    const handleNewChat = async () => {
        createChat.mutateAsync().then((data) => {
            navigate(`/app/chats/${data.data.id}`);
        }).catch((error) => {
            console.error('Failed to create a chat', error);
        });
    }

    const handleSignOut = () => {
        signout();
        setIsAuthenticated(false);
        navigate('/app/login');
    }

    return (
        <Flex direction="column" gap="2" height="100%" justify="between">
            <Flex direction="column">

                <Button size="3" variant="ghost" m="4" onClick={handleNewChat}>
                    <IconMessagePlus size={16} />Start a new chat
                </Button>
                <NavigationMenu.Root orientation="vertical">
                    <NavigationMenu.List className={styles.NavigationMenuList}>
                        {!!chats && chats.data.chats?.map((chat) => (
                            <NavigationMenu.Item key={chat.id}>
                                <NavigationMenu.Link
                                    asChild
                                    className={
                                        styles.NavigationMenuLink
                                        + ' '
                                        + (chatId === chat.id ? styles.NavigationMenuLinkActive : '')
                                    }
                                >
                                    <Link to={`/app/chats/${chat.id}`} onClick={() => open(false)}>
                                        {chat.title}
                                    </Link>
                                </NavigationMenu.Link>
                            </NavigationMenu.Item>
                        ))}
                    </NavigationMenu.List>
                </NavigationMenu.Root>
            </Flex>

            <Flex direction="column" gap="2">
                <Separator size="4" />
                <Button size="3" variant="ghost" m="3" onClick={handleSignOut}>
                    <IconLogout size={16} />Sign out
                </Button>
            </Flex>
        </Flex>

    )
}