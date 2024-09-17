import { Sidebar } from 'features/sidebar';
import { useAppEvents } from 'hooks/useAppEventsApi';

import { Aside, Root } from './';

type PageLayoutProps = {
  children: React.ReactNode;
};

export const PageLayout: React.FC<PageLayoutProps> = ({ children }) => {
  useAppEvents();

  return (
    <Root>
      <Aside>
        <Sidebar />
      </Aside>
      {children}
    </Root>
  );
};
