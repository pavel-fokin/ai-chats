import { OnlyDesktop, OnlyMobile } from '@/components/layout';
import { ChatMenu, NewChatIconButton } from '@/features/chat';
import { OpenSidebarButton } from '@/features/sidebar';
import { useChat } from '@/hooks';

import { ChatTitleMenuButton } from './chat-title-menu-button';
import { MenuButton } from './menu-button';
import { ChatMenuProvider } from '../chat-menu/chat-menu-provider';

import styles from './chat-header.module.css';

interface ChatHeaderProps {
  chatId: string;
}

// Chat header component.
export const ChatHeader = ({ chatId }: ChatHeaderProps): JSX.Element => {
  const chat = useChat(chatId);

  const title = chat.data?.title || 'Chat';

  return (
    <>
      <OnlyMobile>
        <div className={styles.chatHeaderMobile}>
          <OpenSidebarButton />
          <ChatMenuProvider>
            <ChatMenu
              chatId={chatId}
            trigger={<ChatTitleMenuButton title={title} />}
            />
          </ChatMenuProvider>
          <NewChatIconButton />
        </div>
      </OnlyMobile>
      <OnlyDesktop>
        <div className={styles.chatHeaderDesktop}>
          <span className={styles.chatHeaderTitle}>{title}</span>
          <ChatMenuProvider>
            <ChatMenu chatId={chatId} trigger={<MenuButton />} />
          </ChatMenuProvider>
        </div>
      </OnlyDesktop>
    </>
  );
};
