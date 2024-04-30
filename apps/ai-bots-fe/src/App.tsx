import '@mantine/core/styles.css';


import {
  AppShell,
  Burger,
  Group,
  NavLink,
  Text,
  MantineProvider,
  createTheme,
} from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { IconRobotFace, IconMessageChatbot } from '@tabler/icons-react';


import { Chat } from './pages/Chat';

const theme = createTheme({
  /** Put your mantine theme override here */
});


function App() {
  const [opened, { toggle }] = useDisclosure();

  return (
    <MantineProvider theme={theme}>
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
          <NavLink label="Bots" leftSection={<IconRobotFace size="1rem" stroke={1.5}/>}/>
          <NavLink label="Chats" leftSection={<IconMessageChatbot size="1rem" stroke={1.5}/>}/>
        </AppShell.Navbar>
        <AppShell.Main style={{display: 'flex', flexDirection: 'column-reverse'}}>
            <Chat />
        </AppShell.Main>
      </AppShell>
    </MantineProvider>
  );
}

export default App;
