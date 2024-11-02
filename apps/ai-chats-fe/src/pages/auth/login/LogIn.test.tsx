import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthProvider } from 'features/auth';
import { LogIn } from 'pages';

import { generateToken } from 'utils/utilsTests';

const server = setupServer(
  http.post('/api/auth/login', () => {
    return HttpResponse.json({ data: { accessToken: generateToken() } });
  }),
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
    },
  },
});

// Render a component with required providers and routing.
export function renderWithRouter(ui: JSX.Element, { route = '/app' } = {}) {
  return render(
    <AuthProvider>
      <QueryClientProvider client={queryClient}>
        <MemoryRouter initialEntries={[route]}>
          <Routes>
            <Route path="/app" element={<div>App</div>} />
            <Route path="/app/login" element={ui} />
          </Routes>
        </MemoryRouter>
      </QueryClientProvider>
    </AuthProvider>,
  );
}

test('renders Log In component', () => {
  renderWithRouter(<LogIn />, { route: '/app/login' });

  expect(screen.getByRole('heading', { name: 'Log in' })).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Your username')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Your password')).toBeInTheDocument();
  expect(screen.getByRole('button', { name: 'Log in' })).toBeInTheDocument();
  expect(screen.getByText("Don't have an account?")).toBeInTheDocument();
});

test('calls logIn function and navigates to /app on successful log in', async () => {
  renderWithRouter(<LogIn />, { route: '/app/login' });

  const username = 'user';
  const password = 'password';

  const usernameInput = screen.getByPlaceholderText('Your username');
  const passwordInput = screen.getByPlaceholderText('Your password');
  const logInButton = screen.getByRole('button', { name: 'Log in' });

  await userEvent.type(usernameInput, username);
  await userEvent.type(passwordInput, password);

  await userEvent.click(logInButton);

  await waitFor(() => {
    expect(screen.getByText('App')).toBeInTheDocument();
  });
});

test('validation errors are displayed when submitting an empty form', async () => {
  renderWithRouter(<LogIn />, { route: '/app/login' });

  const logInButton = screen.getByRole('button', { name: 'Log in' });

  await userEvent.click(logInButton);

  await waitFor(() => {
    expect(screen.getByText('Username is required')).toBeInTheDocument();
    expect(
      screen.getByText('Password must be at least 6 characters'),
    ).toBeInTheDocument();
  });
});

test('displays error message on unsuccessful log in', async () => {
  server.use(
    http.post('/api/auth/login', () => {
      return HttpResponse.json(
        { errors: [{ message: 'Error' }] },
        { status: 400 },
      );
    }),
  );

  renderWithRouter(<LogIn />, { route: '/app/login' });

  const username = 'user';
  const password = 'password';

  const usernameInput = screen.getByPlaceholderText('Your username');
  const passwordInput = screen.getByPlaceholderText('Your password');
  const logInButton = screen.getByRole('button', { name: 'Log in' });

  await userEvent.type(usernameInput, username);
  await userEvent.type(passwordInput, password);

  await userEvent.click(logInButton);

  await waitFor(() => {
    expect(screen.getByText('Error')).toBeInTheDocument();
  });
});
