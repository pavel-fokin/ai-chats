import { useContext } from 'react';

import { SidebarContext } from './sidebar-context';

export const useSidebarContext = () => {
  const context = useContext(SidebarContext);

  if (!context) {
    throw new Error(
      'useSidebarContext must be used within SidebarContextProvider',
    );
  }

  return context;
};
