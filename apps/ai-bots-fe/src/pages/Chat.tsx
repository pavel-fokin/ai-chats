import { useParams } from "react-router-dom";

import { Flex } from "@radix-ui/themes";

import { InputMessage, Message } from 'components';
import { useMessages } from 'hooks';
import * as types from 'types';

export function Chat() {
    const { chatId } = useParams<{ chatId: string }>();

    if (!chatId) {
        throw new Error('Chat ID is required');
    }

    const { messages, sendMessage } = useMessages(chatId);

    const handleSend = async (msg: types.Message) => {
        sendMessage.mutate({ sender: 'human', text: msg.text });
    }

    return (
        <Flex direction="column-reverse" flexGrow="1">
            <Flex direction="column" gap="6">
                {messages.map((message, index) => (
                    <Message
                        key={index}
                        sender={message.sender}
                        text={message.text}
                    />
                ))}
                <InputMessage handleSend={handleSend} />
            </Flex>
        </Flex>
    );
}