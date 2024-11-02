import { createContext, useState } from 'react';

import { ModelResponse } from 'types';

interface ChatContextValue {
  modelResponse: ModelResponse | null;
  setModelResponse: (modelResponse: ModelResponse | null) => void;
}

export const ChatContext = createContext<ChatContextValue | null>(null);

export const ChatContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [modelResponse, setModelResponse] = useState<ModelResponse | null>(
    null,
  );

  return (
    <ChatContext.Provider value={{ modelResponse, setModelResponse }}>
      {children}
    </ChatContext.Provider>
  );
};
