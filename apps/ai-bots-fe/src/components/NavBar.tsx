import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Flex, Button } from '@radix-ui/themes';
import { IconMessagePlus } from '@tabler/icons-react';
import { useNavigate, useParams } from 'react-router-dom';

import { useChats } from 'hooks';

import styles from './Navbar.module.css';

export function Navbar() {
    const { chats } = useChats();
    const { createChat } = useChats();
    const navigate = useNavigate();
    const { chatId } = useParams<{ chatId: string }>();

    const handleNewChat = async () => {
        createChat.mutateAsync().then((data) => {
            navigate(`/chats/${data.data.id}`);
        }).catch((error) => {
            console.error('Failed to create a chat', error);
        });
    }

    return (
        <Flex direction="column" gap="2">
            <Button size="3" variant="ghost" m="3" onClick={handleNewChat}>
                <IconMessagePlus size={16} />Start a new chat
            </Button>
            <NavigationMenu.Root orientation="vertical">
                <NavigationMenu.List className={styles.NavigationMenuList}>
                    {chats.map((chat) => (
                        <NavigationMenu.Item key={chat.id}>
                            <NavigationMenu.Link
                                href={`/chats/${chat.id}`}
                                className={styles.NavigationMenuLink + ' ' + (chatId === chat.id ? styles.NavigationMenuLinkActive : '')}
                            >
                                {chat.id}
                            </NavigationMenu.Link>
                        </NavigationMenu.Item>
                    ))}
                </NavigationMenu.List>
            </NavigationMenu.Root>
        </Flex>

    )
}