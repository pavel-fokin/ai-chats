import { Sidebar } from 'features/sidebar';

import { Aside, Root } from './';

interface PageLayoutProps {
  children: React.ReactNode;
}

export const PageLayout: React.FC<PageLayoutProps> = ({ children }) => {
  return (
    <Root>
      <Aside>
        <Sidebar />
      </Aside>
      {children}
    </Root>
  );
};
