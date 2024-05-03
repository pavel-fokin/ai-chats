import { useState } from 'react';

import { ActionIcon, Group, Textarea } from "@mantine/core";
import { IconSend } from '@tabler/icons-react';

import * as types from 'types';

function InputMessage() {
    const [inputMessage, setInputMessage] = useState<types.Message>({ sender: '', text: '' });

    const onInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setInputMessage({
            sender: 'You',
            text: event.target.value
        });
    };

    const onSendClick = async () => {
        if (inputMessage) {
            setInputMessage({ sender: '', text: '' });
        }
    };

    return (
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
    );
    }

export { InputMessage };