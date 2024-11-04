import { createContext } from 'react';

interface SidebarContextValue {
  isOpen: boolean;
  toggleSidebar: () => void;
  closeSidebar: () => void;
}

export const SidebarContext = createContext({} as SidebarContextValue);
