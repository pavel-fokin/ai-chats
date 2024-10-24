import { Avatar, Flex, Text } from '@radix-ui/themes';
import Markdown from 'react-markdown';

import styles from './message.module.css';

interface MessageProps {
  sender: string;
  text: string;
}

/**
 * Message component.
 * @param {MessageProps} sender - The sender of the message.
 * @param {string} text - The text of the message.
 * @returns {JSX.Element} - The message component.
 */
export const Message = ({ sender, text }: MessageProps): JSX.Element => {
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
