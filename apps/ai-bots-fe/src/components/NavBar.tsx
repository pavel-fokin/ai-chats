import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { useChats } from 'hooks';

import './Navbar.css';
import { Flex, Heading } from '@radix-ui/themes';

export function Navbar() {
    const chats = useChats();

    return (
        <Flex direction="column" gap="2" pl="4">
            <Heading as="h2">Chats</Heading>
            <NavigationMenu.Root orientation="vertical" className="NavigationMenuRoot">
            <NavigationMenu.List className="NavigationMenuList">
                {chats.map((chat) => (
                    <NavigationMenu.Item key={chat.id}>
                        <NavigationMenu.Link
                            href={`/chats/${chat.id}`}
                            className='NavigationMenuLink'>
                            {chat.id}
                        </NavigationMenu.Link>
                    </NavigationMenu.Item>
                ))}
            </NavigationMenu.List>
        </NavigationMenu.Root>
        </Flex>

    )
}