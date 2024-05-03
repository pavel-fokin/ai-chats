import { useState } from 'react';

import {
    AppShell,
    Burger,
    Group,
    NavLink,
    Text,
} from '@mantine/core';
import { IconMessageChatbot } from '@tabler/icons-react';

import { Chat } from 'pages';
import { useChats } from 'hooks';

export function Main() {
    const [opened, setOpened] = useState(false);
    const toggle = () => setOpened(!opened);

    const chats = useChats();

    return (
        <AppShell
            header={{ height: 60 }}
            navbar={{ width: 300, breakpoint: 'sm', collapsed: { mobile: !opened } }}
            padding="md"
        >
            <AppShell.Header>
                <Group h="100%" px="md">
                    <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
                    <Text>AI Bots</Text>
                </Group>
            </AppShell.Header>
            <AppShell.Navbar p="md">
                <Group p="xs" gap="xs">
                    <IconMessageChatbot size="1.125rem" stroke={1.5} />
                    <Text fw={500}>Chats</Text>
                </Group>
                {chats.map((chat) => {
                    return <NavLink key={chat.id} label={chat.id} />;
                })}
            </AppShell.Navbar>
            <AppShell.Main style={{ display: 'flex', flexDirection: 'column-reverse' }}>
                <Chat />
            </AppShell.Main>
        </AppShell>
    )


}