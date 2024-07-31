import { Sidebar } from 'features/sidebar';

import { Aside, Root, Main } from './';

type PageLayoutProps = {
  children: React.ReactNode;
};

export const PageLayout: React.FC<PageLayoutProps> = ({ children }) => {
  return (
    <Root>
      <Aside>
        <Sidebar />
      </Aside>
      <Main>{children}</Main>
    </Root>
  );
};
