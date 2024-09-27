import { ChatMenu, ChatMenuProvider, NewChatIconButton } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { Header } from 'shared/components/layout';

import { ChatTitleButton } from './ChatTitleButton';

interface ChatHeaderProps {
  chatId: string;
}

export const ChatHeader = ({ chatId }: ChatHeaderProps): JSX.Element => {
  return (
    <Header>
        <OpenSidebarButton />
        <ChatMenuProvider>
          <ChatMenu
            chatId={chatId}
            trigger={<ChatTitleButton chatId={chatId} />}
          />
        </ChatMenuProvider>
        <NewChatIconButton />
    </Header>
  );
};
