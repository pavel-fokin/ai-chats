import { useState } from 'react';

import { DropdownMenu } from '@radix-ui/themes';

import { ChatMenuContext } from './chat-menu-context';

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