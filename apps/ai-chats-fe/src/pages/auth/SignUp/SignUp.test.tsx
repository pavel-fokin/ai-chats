import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthContextProvider } from 'contexts';
import { SignUp } from 'pages';

import { generateToken } from 'utils/utilsTests';

const server = setupServer(
  http.post('/api/auth/signup', () => {
    return HttpResponse.json({ data: { accessToken: generateToken() } });
  })
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
    <AuthContextProvider>
      <QueryClientProvider client={queryClient}>
        <MemoryRouter initialEntries={[route]}>
          <Routes>
            <Route path="/app" element={<div>App</div>} />
            <Route path="/app/signup" element={ui} />
          </Routes>
        </MemoryRouter>
      </QueryClientProvider>
    </AuthContextProvider>
  );
}

test('renders Sign Up component', () => {
  renderWithRouter(<SignUp />, { route: '/app/signup' });

  expect(screen.getByRole('heading', { name: 'Sign up' })).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Your username')).toBeInTheDocument();
  expect(screen.getByPlaceholderText('Your password')).toBeInTheDocument();
  expect(
    screen.getByRole('button', { name: 'Create an account' })
  ).toBeInTheDocument();
  expect(screen.getByText('Already have an account?')).toBeInTheDocument();
});

test('calls signUp function and navigates to /app on successful sign up', async () => {
  renderWithRouter(<SignUp />, { route: '/app/signup' });

  const username = 'user';
  const password = 'password';

  const usernameInput = screen.getByPlaceholderText('Your username');
  const passwordInput = screen.getByPlaceholderText('Your password');
  const signUpButton = screen.getByRole('button', {
    name: 'Create an account',
  });

  await userEvent.type(usernameInput, username);
  await userEvent.type(passwordInput, password);

  await userEvent.click(signUpButton);

  await waitFor(() => {
    expect(screen.getByText('App')).toBeInTheDocument();
  });
});

test('displays validation errors on invalid input', async () => {
  renderWithRouter(<SignUp />, { route: '/app/signup' });

  const signUpButton = screen.getByRole('button', {
    name: 'Create an account',
  });

  await userEvent.click(signUpButton);

  await waitFor(() => {
    expect(screen.getByText('Username is required')).toBeInTheDocument();
    expect(
      screen.getByText('Password must be at least 6 characters')
    ).toBeInTheDocument();
  });
});

test('displays error message on unsuccessful sign in', async () => {});
