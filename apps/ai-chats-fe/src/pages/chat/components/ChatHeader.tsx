import { ChatMenu, NewChatIconButton } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { Header } from 'shared/components/layout';

interface ChatHeaderProps {
  chatId: string;
}

export const ChatHeader = ({ chatId }: ChatHeaderProps) => {
  return (
    <Header>
      <OpenSidebarButton />
      <ChatMenu chatId={chatId} />
      <NewChatIconButton />
    </Header>
  );
};
