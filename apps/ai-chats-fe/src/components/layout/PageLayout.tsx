import { Sidebar } from 'features/sidebar';

import { Aside } from './Aside';
import { Root } from './Root';

interface PageLayoutProps {
  children: React.ReactNode;
}

/**
 * Page layout component.
 * @param {React.ReactNode} children - The children to be rendered inside the page layout
 * @returns {React.ReactElement} - The page layout component
 */
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
