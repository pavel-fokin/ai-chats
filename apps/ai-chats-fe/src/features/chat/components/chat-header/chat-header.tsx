import { OnlyDesktop, OnlyMobile } from '@/components/layout';
import { ChatMenu, NewChatIconButton } from '@/features/chat';
import { OpenSidebarButton } from '@/features/sidebar';
import { useChat } from '@/hooks';

import { ChatTitleMenuButton } from './chat-title-menu-button';
import { MenuButton } from './menu-button';

import styles from './chat-header.module.css';

interface ChatHeaderProps {
  chatId: string;
}

/**
 * Chat header component.
 * @param {string} chatId - The ID of the chat.
 * @returns {JSX.Element} - The chat header component.
 */
export const ChatHeader = ({ chatId }: ChatHeaderProps): JSX.Element => {
  const chat = useChat(chatId);

  const title = chat.data?.title || 'Chat';

  return (
    <>
      <OnlyMobile>
        <div className={styles.chatHeaderMobile}>
          <OpenSidebarButton />
          <ChatMenu
            chatId={chatId}
            trigger={<ChatTitleMenuButton title={title} />}
          />
          <NewChatIconButton />
        </div>
      </OnlyMobile>
      <OnlyDesktop>
        <div className={styles.chatHeaderDesktop}>
          <span className={styles.chatHeaderTitle}>{title}</span>
          <ChatMenu chatId={chatId} trigger={<MenuButton />} />
        </div>
      </OnlyDesktop>
    </>
  );
};
