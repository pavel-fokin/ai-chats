

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { Theme } from '@radix-ui/themes';

import '@radix-ui/themes/styles.css';

import { AuthRequired } from "components";
import { Chat, EmptyState, Landing, Main, LogIn, SignUp } from 'pages';
import { AuthContextProvider } from "contexts";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Landing />,
  },
  {
    path: "/app/login",
    element: <LogIn />,
  },
  {
    path: "/app/signup",
    element: <SignUp />,
  },
  {
    path: "/app",
    element: <AuthRequired><Main /></AuthRequired>,
    children: [
      {
        path: "",
        element: <EmptyState />,
      },
      {
        path: "chats/:chatId",
        element: <Chat />,
      },
    ],
  },
]);

const queryClient = new QueryClient()

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
