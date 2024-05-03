import { useParams } from "react-router-dom";

import { Stack } from "@mantine/core";

import { Message, InputMessage } from 'components';
import { useMessages } from 'hooks';

export function ExistingChat() {
    const { chatId } = useParams<{ chatId: string }>();

    if (!chatId) {
        throw new Error('Chat ID is required');
    }

    const messages = useMessages(chatId);

    return (
        <Stack p="md" gap="0" justify="flex-end">
            <Stack gap="xl" style={{ marginBottom: 'auto' }}>
                {messages.map((message, index) => (
                    <Message
                        key={index}
                        sender={message.sender}
                        text={message.text}
                    />
                ))}
            </Stack>
            <InputMessage />
        </Stack>
    );
}