import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthContextProvider } from "contexts";
import { SignUp } from '../SignUp';

import { generateToken } from './utils';

const server = setupServer(
    http.post('/api/auth/signup', () => {
        return HttpResponse.json({ data: { accessToken: generateToken() } });
    }),
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

test('renders Sign Up component', () => {
    render(
        <AuthContextProvider>
            <MemoryRouter initialEntries={['/app/signup']}>
                <Routes>
                    <Route path="/app/signup" element={
                        <SignUp />
                    } />
                </Routes>
            </MemoryRouter>
        </AuthContextProvider>
    );

    expect(screen.getByRole('heading', { name: 'Sign up' })).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Your username')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Your password')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Create an account' })).toBeInTheDocument();
    expect(screen.getByText('Already have an account?')).toBeInTheDocument();
});

test('calls signUp function and navigates to /app on successful sign up', async () => {
    render(
        <AuthContextProvider>
            <MemoryRouter initialEntries={['/app/signup']}>
                <Routes>
                    <Route path="/app/signup" element={
                        <SignUp />
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
    const signUpButton = screen.getByRole('button', { name: 'Create an account' });

    userEvent.type(usernameInput, username);
    userEvent.type(passwordInput, password);

    userEvent.click(signUpButton);

    await waitFor(() => {
        expect(screen.getByText('App')).toBeInTheDocument();
    });
});

test('displays error message on unsuccessful sign in', async () => {

});