import { Avatar, Flex, Text } from '@radix-ui/themes';
import Markdown from 'react-markdown';

import styles from './Message.module.css';

interface MessageProps {
  sender: string;
  text: string;
}

export const Message: React.FC<MessageProps> = ({ sender, text }) => {
  return (
    <Flex direction="column" gap="1" p="2">
      <Flex gap="2">
        <Avatar size="1" fallback="A" />
        <Text>{sender}</Text>
      </Flex>
      <Markdown className={styles.MarkdownPre}>{text}</Markdown>
    </Flex>
  );
};
