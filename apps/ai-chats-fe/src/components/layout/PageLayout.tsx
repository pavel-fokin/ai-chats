import { Root, Aside } from 'components/layout';
import { Sidebar } from 'components';

type PageLayoutProps = {
  children: React.ReactNode;
};

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
