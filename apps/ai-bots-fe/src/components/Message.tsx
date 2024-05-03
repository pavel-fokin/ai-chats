import { Avatar, Group, Stack, Text } from "@mantine/core";
import Markdown from 'react-markdown';

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

export { Message };