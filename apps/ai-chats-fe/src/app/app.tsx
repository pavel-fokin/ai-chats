import { Outlet } from 'react-router-dom';

import { Aside, Page } from 'components/layout';
import { AuthRequired } from 'features/auth';
import { Sidebar, SidebarProvider } from 'features/sidebar';

import { AppProvider } from './context';

// Main app component.
export const App = () => {
  return (
    <AuthRequired>
      <AppProvider>
        <Page>
          <Aside>
            <SidebarProvider>
              <Sidebar />
            </SidebarProvider>
          </Aside>
          <Outlet />
        </Page>
      </AppProvider>
    </AuthRequired>
  );
};
