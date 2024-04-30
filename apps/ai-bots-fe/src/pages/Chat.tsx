import React, {useRef, useState} from "react";

import { Avatar, ActionIcon, Group, ScrollArea, Stack, Text, Textarea } from "@mantine/core";
import { IconSend } from '@tabler/icons-react';
import Markdown from 'react-markdown';

import type { Message } from '../api';
import * as api from '../api';

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

export function Chat() {
    const [inputMessage, setInputMessage] = useState<Message>({ sender: '', text: '' });
    const [messages, setMessages] = useState<Message[]>([]);

    const viewport = useRef<HTMLDivElement>(null);

    const scrollToBottom = () =>
        viewport.current!.scrollTo({ top: viewport.current!.scrollHeight, behavior: 'smooth' });

    const onSendClick = async () => {
        if (inputMessage) {
            const response = await api.SendMessage(inputMessage);
            setMessages([...messages, inputMessage, response.data]);
            setInputMessage({ sender: '', text: '' });
            scrollToBottom();
        }
    };

    const onInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setInputMessage({
            sender: 'You',
            text: event.target.value
        });
    };


    return (
        <Stack p="md" gap="0" justify="flex-end">
            <ScrollArea mb="xl" type="scroll" viewportRef={viewport}>
                <Stack gap="xl" style={{ marginBottom: 'auto' }}>
                    {messages.map((message, index) => (
                        <Message
                            key={index}
                            sender={message.sender}
                            text={message.text}
                        />
                    ))}
                </Stack>
            </ScrollArea>
            <Group gap="xs">
                <Textarea
                    autosize
                    style={{ flexGrow: '1' }}
                    onChange={onInputChange}
                    value={inputMessage.text}
                />
                <ActionIcon size="input-sm" onClick={onSendClick}>
                    <IconSend size={16} />
                </ActionIcon>
            </Group>
        </Stack>
    );
}