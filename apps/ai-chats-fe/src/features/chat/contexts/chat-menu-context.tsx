import { createContext, useState } from 'react';

import { DropdownMenu } from '@radix-ui/themes';

interface ChatMenuContextValue {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const ChatMenuContext = createContext<ChatMenuContextValue>({
  isOpen: false,
  setIsOpen: () => {},
});

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
