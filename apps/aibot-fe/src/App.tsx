import '@mantine/core/styles.css';

import React, { useState, useRef } from 'react';

import {
  ActionIcon,
  Avatar,
  Container,
  Group,
  MantineProvider,
  ScrollArea,
  Stack,
  Text,
  Textarea,
  createTheme,
} from '@mantine/core';
import { IconSend } from '@tabler/icons-react';
import Markdown from 'react-markdown';

import * as api from './api';

const theme = createTheme({
  /** Put your mantine theme override here */
});

const Message = (props: { sender: string; text: string }) => {
  return (
    <Stack gap="2px">
      <Group gap="4px">
        <Avatar size="sm" radius="xl" />
        <Text fw={500}>{props.sender}</Text>
      </Group>
      <Text component='span'><Markdown>{props.text}</Markdown></Text>
    </Stack>
  );
};

function App() {
  const [inputMessage, setInputMessage] = useState('');
  const [messages, setMessages] = useState<string[]>([]);

  const viewport = useRef<HTMLDivElement>(null);

  const scrollToBottom = () =>
    viewport.current!.scrollTo({ top: viewport.current!.scrollHeight, behavior: 'smooth' });

  const onSendClick = async () => {
    if (inputMessage) {
      const response = await api.SendMessage(inputMessage);
      setMessages([...messages, inputMessage, response.data.text]);
      setInputMessage('');
      scrollToBottom();
    }
  };

  const onInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputMessage(event.target.value);
  };

  return (
    <MantineProvider theme={theme}>
      <Container size="xs" style={{ paddingInline: '0' }}>
        <Stack p="md" gap="xl" justify="flex-end" style={{ height: '100vh' }}>
          <ScrollArea type="scroll" viewportRef={viewport}>
            <Stack gap="md" style={{ marginBottom: 'auto' }}>
            {messages.map((message, index) => (
              <Message
                key={index}
                sender={index % 2 === 0 ? 'You' : 'Bot'}
                text={message}
              />
            ))}
            </Stack>
          </ScrollArea>
          <Group gap="xs">
            <Textarea
              autosize
              style={{ flexGrow: '1' }}
              onChange={onInputChange}
            />
            <ActionIcon size="input-sm" onClick={onSendClick}>
              <IconSend size={16} />
            </ActionIcon>
          </Group>
        </Stack>
      </Container>
    </MantineProvider>
  );
}

export default App;
