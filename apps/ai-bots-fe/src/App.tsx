

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { Theme } from '@radix-ui/themes';

import '@radix-ui/themes/styles.css';

import { Main, Chat } from 'pages';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Main />,
    children: [
      {
        path: "/",
        element: <Chat />,
      },
      {
        path: "/chats/:chatId",
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
        <RouterProvider router={router} />
      </QueryClientProvider>
    </Theme>
  );
}

export default App;
