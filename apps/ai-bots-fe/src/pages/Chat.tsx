import { useParams } from "react-router-dom";

import { Flex } from "@radix-ui/themes";

import { InputMessage, Message } from 'components';
import { useMessages } from 'hooks';

export function Chat() {
    const { chatId } = useParams<{ chatId: string }>();

    if (!chatId) {
        return (
            <InputMessage />
        );
    }

    const messages = useMessages(chatId);

    return (
        <Flex direction="column" gap="6">
            {messages.map((message, index) => (
                <Message
                    key={index}
                    sender={message.sender}
                    text={message.text}
                />
            ))}
            <InputMessage />
        </Flex>
    );
}