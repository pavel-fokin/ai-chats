import { createContext, useState, useContext } from 'react';

import { DropdownMenu } from '@radix-ui/themes';

interface ChatMenuContextValue {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

const ChatMenuContext = createContext<ChatMenuContextValue>({
  isOpen: false,
  setIsOpen: () => {},
});

// eslint-disable-next-line react-refresh/only-export-components
export const useChatMenu = () => {
  const context = useContext(ChatMenuContext);
  if (!context) {
    throw new Error('useChatMenu must be used within a ChatMenuProvider');
  }
  return context;
};

interface ChatMenuProviderProps {
  children: React.ReactNode;
}

export const ChatMenuProvider = ({ children }: ChatMenuProviderProps) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <DropdownMenu.Root open={isOpen} onOpenChange={setIsOpen}>
      <ChatMenuContext.Provider value={{ isOpen, setIsOpen }}>
        {children}
      </ChatMenuContext.Provider>
    </DropdownMenu.Root>
  );
};
