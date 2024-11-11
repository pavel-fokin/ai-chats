import { IconButton } from '@/components/ui';
import { MenuIcon } from '@/components/icons';

import { ChatMenu } from '../chat-menu';

import styles from './chat-header-desktop.module.css';

interface ChatHeaderDesktopProps {
  chatId: string;
  title: string;
}

export const ChatHeaderDesktop = ({
  chatId,
  title,
}: ChatHeaderDesktopProps) => {
  return (
    <div className={styles.chatHeaderDesktop__container}>
      <h1 className={styles.chatHeaderDesktop__title}>{title}</h1>
      <ChatMenu chatId={chatId} trigger={<ChatMenuButton />} />
    </div>
  );
};

const ChatMenuButton = () => {
  return (
    <IconButton
      aria-label="Chat menu button"
      variant="ghost"
      size="2"
      highContrast
    >
      {' '}
      <MenuIcon size="24" />{' '}
    </IconButton>
  );
};
