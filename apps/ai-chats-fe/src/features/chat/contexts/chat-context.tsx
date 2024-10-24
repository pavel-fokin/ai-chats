import { createContext, useState } from 'react';

import { MessageChunk } from 'types';

interface ChatContextValue {
  messageChunk: MessageChunk | null;
  setMessageChunk: (messageChunk: MessageChunk | null) => void;
}

export const ChatContext = createContext<ChatContextValue | null>(null);

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
