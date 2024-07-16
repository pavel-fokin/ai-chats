import { Flex, Avatar, Text } from '@radix-ui/themes';
import Markdown from 'react-markdown';

import styles from './Message.module.css';

const Message = (props: { sender: string; text: string }) => {
  return (
    <Flex direction="column" gap="1" p="2">
      <Flex gap="2">
        <Avatar size="1" fallback="A" />
        <Text>{props.sender}</Text>
      </Flex>
      <Markdown className={styles.MarkdownPre}>{props.text}</Markdown>
    </Flex>
  );
};

export { Message };
