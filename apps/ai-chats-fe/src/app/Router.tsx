import { createBrowserRouter, Navigate, Outlet } from 'react-router-dom';

import {
  Chat,
  Landing,
  LogIn,
  NewChat,
  OllamaSettings,
  SignUp,
} from 'pages';

import { AuthRequired } from 'features/auth';
import { ChatContextProvider } from 'features/chat';
import { AppEvents } from 'features/events';
import { SidebarContextProvider } from 'features/sidebar';

export const Router = createBrowserRouter([
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
        element: <NewChat />,
      },
      {
        path: 'chats/:chatId',
        element: (
          <ChatContextProvider>
            <Chat />
          </ChatContextProvider>
        ),
      },
      {
        path: 'settings',
        element: <OllamaSettings />,
      },
    ],
  },
]);
