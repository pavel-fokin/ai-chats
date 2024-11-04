import { useState } from 'react';

import { SidebarContext } from './sidebar-context';

export const SidebarProvider = ({
  children,
}: {
    children: React.ReactNode;
  }) => {
    const [isOpen, setIsOpen] = useState(false);

    const toggleSidebar = () => {
      setIsOpen(!isOpen);
    };

    const closeSidebar = () => {
      setIsOpen(false);
    };

    return (
      <SidebarContext.Provider value={{ isOpen, toggleSidebar, closeSidebar }}>
        {children}
      </SidebarContext.Provider>
    );
  };
