import { createContext } from 'react';

interface ChatMenuContextValue {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const ChatMenuContext = createContext<ChatMenuContextValue>({
  isOpen: false,
  setIsOpen: () => {},
});

