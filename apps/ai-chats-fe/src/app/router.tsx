import { createBrowserRouter, Navigate } from 'react-router';

import { AuthRequired } from '@/features/auth';
import { ChatContextProvider } from '@/features/chat';
import {
  Chat,
  Landing,
  LogIn,
  NewChat,
  OllamaLibraryPage,
  SignUp,
} from '@/pages';

import { AppMain } from './app-main';

export const Router = createBrowserRouter(
  [
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
          <AppMain />
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
          path: 'ollama-library',
          element: <OllamaLibraryPage />,
        },
      ],
    },
  ],
  {
    future: {
      v7_relativeSplatPath: true,
      v7_fetcherPersist: true,
      v7_normalizeFormMethod: true,
      v7_partialHydration: true,
      v7_skipActionErrorRevalidation: true,
    },
  }
);
