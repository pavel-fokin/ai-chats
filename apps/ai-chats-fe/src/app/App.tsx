import { RouterProvider } from 'react-router-dom';

import { Theme } from '@radix-ui/themes';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import '@radix-ui/themes/styles.css';
import 'styles/spacing.css';

import { AuthContextProvider } from 'features/auth';

import { Router } from './Router';

const queryClient = new QueryClient();

export const App = () => {
  return (
    <Theme appearance="light" accentColor="gray" grayColor="slate">
      <QueryClientProvider client={queryClient}>
        <AuthContextProvider>
          <RouterProvider router={Router} />
        </AuthContextProvider>
      </QueryClientProvider>
    </Theme>
  );
};
