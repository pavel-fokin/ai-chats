import { useParams } from 'react-router-dom';

import { NewChatIconButton } from 'components';
import { Chat, ChatContextProvider, ChatMenu } from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { Header, PageLayout } from 'layout';

export const ChatPage: React.FC = () => {
  const { chatId } = useParams<{ chatId: string }>();

  if (!chatId) {
    return <div>No chat selected</div>;
  }

  return (
    <PageLayout>
      <ChatContextProvider>
        <Header>
          <OpenSidebarButton />
          <ChatMenu chatId={chatId} />
          <NewChatIconButton />
        </Header>
        <Chat chatId={chatId} />
      </ChatContextProvider>
    </PageLayout>
  );
};
