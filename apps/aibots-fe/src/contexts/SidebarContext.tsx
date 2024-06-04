import { createContext, useState } from 'react';

type SidebarContextValue = {
  isOpen: boolean;
  toggleSidebar: () => void;
};

export const SidebarContext = createContext({} as SidebarContextValue);

export const SidebarContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleSidebar = () => {
    setIsOpen(!isOpen);
  };

  return (
    <SidebarContext.Provider value={{ isOpen, toggleSidebar }}>
      {children}
    </SidebarContext.Provider>
  );
};
