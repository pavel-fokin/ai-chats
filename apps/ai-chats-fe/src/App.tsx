import { createBrowserRouter, RouterProvider, Outlet } from 'react-router-dom';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { Theme } from '@radix-ui/themes';

import '@radix-ui/themes/styles.css';

import { AuthRequired } from 'components';
import {
  Chat,
  Landing,
  Startpage,
  LogIn,
  SignUp,
  Settings,
  SignOut,
} from 'pages';
import { AuthContextProvider, SidebarContextProvider } from 'contexts';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Landing />,
  },
  {
    path: '/app/login',
    element: <LogIn />,
  },
  {
    path: '/app/signup',
    element: <SignUp />,
  },
  {
    path: '/app/signout',
    element: <SignOut />,
  },
  {
    path: '/app',
    element: (
      <AuthRequired>
        <SidebarContextProvider>
          <Outlet />
        </SidebarContextProvider>
      </AuthRequired>
    ),
    children: [
      {
        path: '',
        element: <Startpage />,
      },
      {
        path: 'chats/:chatId',
        element: <Chat />,
      },
      {
        path: 'settings',
        element: <Settings />,
      },
    ],
  },
]);

const queryClient = new QueryClient();

function App() {
  return (
    <Theme appearance="light" accentColor="gray" grayColor="slate">
      <QueryClientProvider client={queryClient}>
        <AuthContextProvider>
          <RouterProvider router={router} />
        </AuthContextProvider>
      </QueryClientProvider>
    </Theme>
  );
}

export default App;
