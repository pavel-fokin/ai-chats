

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import {
  MantineProvider,
  createTheme,
} from '@mantine/core';

import '@mantine/core/styles.css';

import { Main } from './pages/Main';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Main />,
    children: [
      {
        path: "/chat",
        element: <Main />,
      },
      {
        path: "/chat/:chatId",
        element: <Main />,
      },
    ],
  },
]);

const theme = createTheme({
  /** Put your mantine theme override here */
});

const queryClient = new QueryClient()

function App() {
  return (
    <MantineProvider theme={theme}>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </MantineProvider>
  );
}

export default App;
