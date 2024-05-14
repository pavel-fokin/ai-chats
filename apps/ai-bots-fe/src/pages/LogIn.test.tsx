import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthContextProvider } from "contexts";
import { LogIn } from './LogIn';

const server = setupServer(
    http.post('/api/auth/login', () => {
        return HttpResponse.json({ data: { accessToken: 'accessToken' } });
    }),
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

test('renders Log In component', () => {
    render(
        <AuthContextProvider>
            <MemoryRouter initialEntries={['/app/login']}>
                <Routes>
                    <Route path="/app/login" element={
                        <LogIn />
                    } />
                </Routes>
            </MemoryRouter>
        </AuthContextProvider>
    );

    expect(screen.getByRole('heading', { name: 'Log in' })).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Your username')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Your password')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Log In' })).toBeInTheDocument();
    expect(screen.getByText("Don't have an account?")).toBeInTheDocument();
});

test('calls signIn function and navigates to /app on successful log in', async () => {
    render(
        <AuthContextProvider>
            <MemoryRouter initialEntries={['/app/login']}>
                <Routes>
                    <Route path="/app/login" element={
                        <LogIn />
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
    const signInButton = screen.getByRole('button', { name: 'Log In' });

    userEvent.type(usernameInput, username);
    userEvent.type(passwordInput, password);

    userEvent.click(signInButton);

    await waitFor(() => {
        expect(screen.getByText('App')).toBeInTheDocument();
    });
});

test('displays error message on unsuccessful log in', async () => {

});