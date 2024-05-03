import { useState } from "react";

import { Stack } from "@mantine/core";

import * as types from 'types';
import { Message, InputMessage } from 'components';

export function NewChat() {
    const [messages] = useState<types.Message[]>([]);

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