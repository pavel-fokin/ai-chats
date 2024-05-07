import { Flex, Avatar, Text } from "@radix-ui/themes";
// import { IconRobotFace } from '@tabler/icons-react';
import Markdown from 'react-markdown';

const Message = (props: { sender: string; text: string }) => {
    return (
        <Flex direction="column" gap="1">
            <Flex>
                <Avatar size="1" fallback="A"/>
                <Text>{props.sender}</Text>
            </Flex>
            <Markdown>{props.text}</Markdown>
        </Flex>
    );
};

export { Message };