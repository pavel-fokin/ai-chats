import { createContext } from 'react';

import { useAppEvents } from 'hooks';

// App context.
const AppContext = createContext(null);

// App context provider.
export const AppProvider = ({ children }: { children: React.ReactNode }) => {
  useAppEvents();

  return <AppContext.Provider value={null}>{children}</AppContext.Provider>;
};
