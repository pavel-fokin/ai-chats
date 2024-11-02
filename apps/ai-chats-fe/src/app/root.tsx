import { RouterProvider } from 'react-router-dom';

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Theme } from '@radix-ui/themes';

import '@radix-ui/themes/styles.css';
import 'styles/spacing.css';

import { AuthProvider } from 'features/auth';

import { Router } from './router';

// App root react component.
export const AppRoot = () => {
  const queryClient = new QueryClient();

  return (
    <Theme appearance="light" accentColor="gray" grayColor="slate">
      <AuthProvider>
        <QueryClientProvider client={queryClient}>
          <RouterProvider router={Router} />
        </QueryClientProvider>
      </AuthProvider>
    </Theme>
  );
};
