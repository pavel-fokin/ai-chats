import { useParams } from 'react-router-dom';

import { Box, Flex } from '@radix-ui/themes';

import { NewChatIconButton } from 'components';
import { useMessages, useSendMessage } from 'hooks';
import {
  ChatContextProvider,
  ChatMenu,
  InputMessage,
  Message,
} from 'features/chat';
import { OpenSidebarButton } from 'features/sidebar';
import { Header, PageLayout } from 'layout';

import { MessageChunk } from './components';

export const Chat: React.FC = () => {
  const { chatId } = useParams<{ chatId: string }>();
  const messages = useMessages(chatId);
  const sendMessage = useSendMessage(chatId ?? '');

  const handleSend = async (text: string) => {
    sendMessage.mutate(text);
  };

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
        <Flex direction="column" height="100%" width="100%">
          <Box flexGrow="1" style={{ overflow: 'auto' }}>
            <Box height="100%" style={{ maxWidth: '688px', margin: '0 auto' }}>
              <Flex flexGrow="1" justify="end" direction="column" gap="2">
                {messages.data?.length !== 0 && (
                  <Box style={{ height: '64px' }}></Box>
                )}
                {messages.data?.map((message, index) => (
                  <Message
                    key={index}
                    sender={message.sender}
                    text={message.text}
                  />
                ))}
                <MessageChunk />
              </Flex>
            </Box>
          </Box>
          <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
            <InputMessage handleSend={handleSend} />
          </Box>
        </Flex>
      </ChatContextProvider>
    </PageLayout>
  );
};
