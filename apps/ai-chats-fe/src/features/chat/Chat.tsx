import { Box, Flex } from '@radix-ui/themes';

import { useMessages, useSendMessage } from 'hooks';

import { InputMessage, Message } from './components';
import { useChatContext } from './contexts';

interface ChatProps {
  chatId: string;
}

export const Chat: React.FC<ChatProps> = ({ chatId }) => {
  const { messageChunk } = useChatContext();
  const messages = useMessages(chatId);
  const sendMessage = useSendMessage(chatId);

  const handleSend = async (text: string) => {
    sendMessage.mutate(text);
  };

  return (
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
            {messageChunk?.text && (
              <Message sender={messageChunk.sender} text={messageChunk.text} />
            )}
          </Flex>
        </Box>
      </Box>
      <Box style={{ maxWidth: '688px', width: '100%', margin: '0 auto' }}>
        <InputMessage handleSend={handleSend} />
      </Box>
    </Flex>
  );
};
