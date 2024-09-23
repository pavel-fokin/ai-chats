import {
  createBrowserRouter,
  Navigate,
  Outlet,
  RouterProvider,
} from 'react-router-dom';

import { Theme } from '@radix-ui/themes';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import '@radix-ui/themes/styles.css';

import { AuthRequired } from 'components';
import { AuthContextProvider } from 'contexts';
import { AppEvents } from 'features/events';
import { SidebarContextProvider } from 'features/sidebar';
import {
  ChatPage,
  Landing,
  LogIn,
  NewChatPage,
  OllamaSettings,
  SignOut,
  SignUp,
} from 'pages';

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
        <AppEvents>
          <SidebarContextProvider>
            <Outlet />
          </SidebarContextProvider>
        </AppEvents>
      </AuthRequired>
    ),
    children: [
      {
        path: '',
        element: <Navigate to="new-chat" />,
      },
      {
        path: 'new-chat',
        element: <NewChatPage />,
      },
      {
        path: 'chats/:chatId',
        element: <ChatPage />,
      },
      {
        path: 'settings',
        element: <OllamaSettings />,
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
