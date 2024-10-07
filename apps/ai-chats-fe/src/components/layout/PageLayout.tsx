import { Sidebar } from 'features/sidebar';

import { Aside } from './Aside';
import { Root } from './Root';

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
