import { ChatMenu, ChatMenuProvider, NewChatIconButton } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { OnlyDesktop, OnlyMobile } from 'shared/components';
import { useChat } from 'shared/hooks';
import { Header } from 'shared/components/layout';

import { ChatTitleMenuButton } from './ChatTitleMenuButton';
import { MenuButton } from './MenuButton';

import styles from './ChatHeader.module.css';

interface ChatHeaderProps {
  chatId: string;
}

export const ChatHeader = ({ chatId }: ChatHeaderProps): JSX.Element => {
  const chat = useChat(chatId);

  const title = chat.data?.title || 'Chat';

  return (
    <Header>
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
          <div>{title}</div>
          <ChatMenuProvider>
            <ChatMenu
              chatId={chatId}
              trigger={<MenuButton />}
            />
          </ChatMenuProvider>
        </div>
      </OnlyDesktop>
    </Header>
  );
};
