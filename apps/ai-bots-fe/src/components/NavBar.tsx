import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { useChats } from 'hooks';

import './Navbar.css';
import { Flex } from '@radix-ui/themes';
import { IconMessagePlus } from '@tabler/icons-react';

export function Navbar() {
    const chats = useChats();

    return (
        <Flex direction="column" gap="2" pl="4">
            <NavigationMenu.Root orientation="vertical" className="NavigationMenuRoot">
                <NavigationMenu.List className="NavigationMenuList">
                    <NavigationMenu.Item >
                        <NavigationMenu.Link
                            href="/"
                            className='NavigationMenuLink'
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