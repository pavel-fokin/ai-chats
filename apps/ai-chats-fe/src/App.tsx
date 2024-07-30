import {
  createBrowserRouter,
  Navigate,
  Outlet,
  RouterProvider,
} from 'react-router-dom';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import { Theme } from '@radix-ui/themes';

import '@radix-ui/themes/styles.css';

import { AuthRequired } from 'components';
import { AuthContextProvider, SidebarContextProvider } from 'contexts';
import {
  ChatPage,
  LandingPage,
  LogIn,
  NewChatPage,
  OllamaSettings,
  SignOut,
  SignUp,
} from 'pages';

const router = createBrowserRouter([
  {
    path: '/',
    element: <LandingPage />,
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
        element: <Navigate to="chats/new" />,
      },
      {
        path: 'chats/new',
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
