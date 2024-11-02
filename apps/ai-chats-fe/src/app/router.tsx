import { createBrowserRouter, Navigate } from 'react-router-dom';

import { Chat, Landing, LogIn, NewChat, OllamaSettings, SignUp } from 'pages';

import { App } from './app';
import { AuthRequired } from 'features/auth';
import { ChatContextProvider } from 'features/chat';

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
        <App />
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
