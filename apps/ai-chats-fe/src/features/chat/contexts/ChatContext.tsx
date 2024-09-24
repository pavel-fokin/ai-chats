import { createContext, useState, useContext } from 'react';

import { MessageChunk } from 'types';

interface ChatContextValue {
  messageChunk: MessageChunk | null;
  setMessageChunk: (messageChunk: MessageChunk | null) => void;
}

const ChatContext = createContext<ChatContextValue | null>(null);

// eslint-disable-next-line
export const useChatContext = () => {
  const context = useContext(ChatContext);
  if (!context) {
    throw new Error('useChatContext must be used within a ChatContextProvider');
  }
  return context;
};

export const ChatContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [messageChunk, setMessageChunk] = useState<MessageChunk | null>(null);

  return (
    <ChatContext.Provider value={{ messageChunk, setMessageChunk }}>
      {children}
    </ChatContext.Provider>
  );
};
