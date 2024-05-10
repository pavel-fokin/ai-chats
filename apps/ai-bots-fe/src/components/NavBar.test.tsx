import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthContextProvider } from "contexts";
import { Navbar } from './Navbar';

const server = setupServer(
    http.get('/api/chats', () => {
        return HttpResponse.json({ data: { chats: [{ id: 'chatId' }] } });
    }),
    http.post('/api/chats', () => {
        return HttpResponse.json({ data: { id: 'newChatId' } });
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

function renderWithRouter(ui: JSX.Element, { route = '/' } = {}) {
    return render(
        <AuthContextProvider>
            <QueryClientProvider client={queryClient}>
                <MemoryRouter initialEntries={[route]}>
                    <Routes>
                        <Route path="/app" element={ui} />
                        <Route path="/app/chats/:chatId" element={<div>chatId</div>} />
                        <Route path="/app/signin" element={<div>Sign in</div>} />
                    </Routes>
                </MemoryRouter>
            </QueryClientProvider>
        </AuthContextProvider>
    );
}

test('renders Navbar component', async () => {
    renderWithRouter(<Navbar />, { route: '/app' });

    expect(screen.getByText('Start a new chat')).toBeInTheDocument();
    expect(screen.getByText('Sign out')).toBeInTheDocument();

    await waitFor(() => {
        expect(screen.getByText('chatId')).toBeInTheDocument();
    });
});

test('navigates to chat on chat link click', async () => {
    renderWithRouter(<Navbar />, { route: '/app' });

    await waitFor(async () => {
        expect(screen.getByText('chatId')).toBeInTheDocument();
        await userEvent.click(screen.getByText('chatId'));
        expect(screen.getByText('chatId')).toBeInTheDocument();
    });
});

test('calls handleNewChat on new chat button click', async () => {
    renderWithRouter(<Navbar />, { route: '/app' });

    await waitFor(async () => {
        await userEvent.click(screen.getByText('Start a new chat'));
        expect(screen.getByText('chatId')).toBeInTheDocument();
    });
});

test('calls handleSignOut on sign out button click', async () => {
    renderWithRouter(<Navbar />, { route: '/app' });

    await waitFor(async () => {
        await userEvent.click(screen.getByText('Sign out'));
        expect(screen.getByText('Sign in')).toBeInTheDocument();
    });
});