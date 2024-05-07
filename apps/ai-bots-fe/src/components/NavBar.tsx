import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import { Flex } from '@radix-ui/themes';
import { IconMessagePlus } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';

import { useChats } from 'hooks';

import './Navbar.css';

export function Navbar() {
    const { chats } = useChats();
    const { createChat } = useChats();
    const navigate = useNavigate();

    const handleNewChat = async () => {
        createChat.mutateAsync().then((data) => {
            navigate(`/chats/${data.data.id}`);
        }).catch((error) => {
            console.error('Failed to create a chat', error);
        });
    }

    return (
        <Flex direction="column" gap="2" px="4">
            <NavigationMenu.Root orientation="vertical">
                <NavigationMenu.List className="NavigationMenuList">
                    <NavigationMenu.Item >
                        <NavigationMenu.Link
                            // href="/"
                            className='NavigationMenuLink'
                            onClick={handleNewChat}
                        >
                            <Flex gap="3">
                                <IconMessagePlus size={16} /> Start a new chat
                            </Flex>
                        </NavigationMenu.Link>
                    </NavigationMenu.Item>
                    {chats.map((chat) => (
                        <NavigationMenu.Item key={chat.id}>
                            <NavigationMenu.Link
                                href={`/chats/${chat.id}`}
                                className='NavigationMenuLink'
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