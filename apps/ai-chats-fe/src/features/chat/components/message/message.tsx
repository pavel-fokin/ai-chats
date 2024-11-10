import Markdown from 'react-markdown';

import { Avatar } from '@/components/ui/avatar';

import styles from './message.module.css';

interface MessageProps {
  sender: string;
  text: string;
}

// Message component.
export const Message = ({ sender, text }: MessageProps) => {
  const fallback = sender.charAt(0);

  return (
    <article className={styles.message__container}>
      <header className={styles.message__headerContainer}>
        <Avatar
          aria-label={`Avatar for ${sender}`}
          size="1"
          fallback={fallback}
        />
        <h2>{sender}</h2>
      </header>
      <Markdown className={styles.message__markdownPre}>{text}</Markdown>
    </article>
  );
};
