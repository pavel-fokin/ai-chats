import { Outlet } from 'react-router-dom';

import { Aside, Page } from '@/components/layout';
import { AuthRequired } from '@/features/auth';
import { Sidebar, SidebarProvider } from '@/features/sidebar';

import { AppProvider } from './app-context';

// Main app component.
export const AppMain = () => {
  return (
    <AuthRequired>
      <AppProvider>
        <Page>
          <SidebarProvider>
            <Aside>
              <Sidebar />
            </Aside>
            <Outlet />
          </SidebarProvider>
        </Page>
      </AppProvider>
    </AuthRequired>
  );
};
