import { useState } from 'react';

import { Box, Flex, IconButton, TextField } from "@radix-ui/themes";
import { IconSend } from '@tabler/icons-react';

import * as types from 'types';

type InputMessageProps = {
    handleSend: (msg: types.Message) => void;
};

function InputMessage({ handleSend }: InputMessageProps) {
    const [inputMessage, setInputMessage] = useState<types.Message>({ sender: '', text: '' });

    const onInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setInputMessage({ ...inputMessage, text: event.target.value });
    }

    const onSendClick = async (e: React.FormEvent) => {
        e.preventDefault();
        if (inputMessage.text.trim() !== ''){
            handleSend({ sender: 'User', text: inputMessage.text });
            setInputMessage({ sender: '', text: '' });
        }
    };

    return (
        <form role="form" onSubmit={onSendClick} >
            <Flex gap="2" justify="center" p={{
                initial: '2',
                sm: '4',
            }}>
                <Box flexGrow="1">
                    <TextField.Root
                        value={inputMessage.text}
                        placeholder="Type a message"
                        onChange={onInputChange}
                        size="3"
                    />
                </Box>
                <IconButton size="3" onClick={onSendClick} highContrast>
                    <IconSend size={16} />
                </IconButton>
            </Flex>
        </form>
    );
}

export { InputMessage };
