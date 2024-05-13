import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthContextProvider } from "contexts";
import { SignIn } from './SignIn';

const server = setupServer(
    http.post('/api/auth/signin', () => {
        return HttpResponse.json({ data: { accessToken: 'accessToken' } });
    }),
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

test('renders Sign In component', () => {
    render(
        <AuthContextProvider>
            <MemoryRouter initialEntries={['/app/signin']}>
                <Routes>
                    <Route path="/app/signin" element={
                        <SignIn />
                    } />
                </Routes>
            </MemoryRouter>
        </AuthContextProvider>
    );

    expect(screen.getByRole('heading', { name: 'Sign In' })).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Your username')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Your password')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Sign In' })).toBeInTheDocument();
    expect(screen.getByText("Don't have an account?")).toBeInTheDocument();
});

test('calls signIn function and navigates to /app on successful sign in', async () => {
    render(
        <AuthContextProvider>
            <MemoryRouter initialEntries={['/app/signin']}>
                <Routes>
                    <Route path="/app/signin" element={
                        <SignIn />
                    } />
                    <Route path="/app" element={<div>App</div>} />
                </Routes>
            </MemoryRouter>
        </AuthContextProvider>
    );

    const username = 'user';
    const password = 'password';

    const usernameInput = screen.getByPlaceholderText('Your username');
    const passwordInput = screen.getByPlaceholderText('Your password');
    const signInButton = screen.getByRole('button', { name: 'Sign In' });

    userEvent.type(usernameInput, username);
    userEvent.type(passwordInput, password);

    userEvent.click(signInButton);

    await waitFor(() => {
        expect(screen.getByText('App')).toBeInTheDocument();
    });
});

test('displays error message on unsuccessful sign in', async () => {

});