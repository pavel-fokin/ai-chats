import { useContext } from 'react';

import { ChatMenuContext } from '../contexts/chat-menu-context';

export const useChatMenu = () => {
  const context = useContext(ChatMenuContext);
  if (!context) {
    throw new Error('useChatMenu must be used within a ChatMenuProvider');
  }
  return context;
};
