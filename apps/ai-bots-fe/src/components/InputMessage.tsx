import { useState } from 'react';

import { Box, Flex, IconButton, TextField } from "@radix-ui/themes";
import { IconSend } from '@tabler/icons-react';

import * as types from 'types';

function InputMessage() {
    const [inputMessage, setInputMessage] = useState<types.Message>({ sender: '', text: '' });

    const onInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setInputMessage({ ...inputMessage, text: event.target.value });
    }

    const onSendClick = async () => {
        if (inputMessage) {
            setInputMessage({ sender: '', text: '' });
        }
    };

    return (
        <Flex gap="2" justify="center">
            <Box flexGrow="1">
                <TextField.Root
                    value={inputMessage.text}
                    placeholder="Type a message"
                    style={{ resize: 'none' }}
                    onChange={onInputChange}
                    size="3"
                />
            </Box>
            <IconButton size="3" onClick={onSendClick} highContrast>
                <IconSend size={16} />
            </IconButton>
        </Flex>
    );
}

export { InputMessage };
