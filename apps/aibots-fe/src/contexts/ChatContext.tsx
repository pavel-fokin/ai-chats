import { createContext, useState } from 'react';

type ChatContextValue = {
    chatId: string;
    setChatId: (chatId: string) => void;
}

export const ChatContext = createContext({} as ChatContextValue);

export const ChatContextProvider = ({ children }: { children: React.ReactNode }) => {
    const [chatId, setChatId] = useState('');

    return (
        <ChatContext.Provider value={{ chatId, setChatId }}>
            {children}
        </ChatContext.Provider>
    );
}