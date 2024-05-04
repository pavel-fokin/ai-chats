import * as NavigationMenu from '@radix-ui/react-navigation-menu';

import { useChats } from 'hooks';

export function NavBar() {
    const chats = useChats();

    return (
        <NavigationMenu.Root orientation="vertical">
            <NavigationMenu.List>
                {chats.map((chat) => (
                    <NavigationMenu.Item key={chat.id}>
                        <NavigationMenu.Link href={`/chats/${chat.id}`}>{chat.id}</NavigationMenu.Link>
                    </NavigationMenu.Item>
                ))}
            </NavigationMenu.List>
        </NavigationMenu.Root>
    )
}