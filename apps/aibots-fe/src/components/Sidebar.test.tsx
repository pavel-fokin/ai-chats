import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthContextProvider } from 'contexts';
import { Sidebar } from 'components';

const server = setupServer(
  http.get('/api/chats', () => {
    return HttpResponse.json({
      data: {
        chats: [
          { id: 'someChatId', title: 'Some chat', createdAt: new Date() },
        ],
      },
    });
  }),
  http.post('/api/chats', () => {
    return HttpResponse.json({
      data: {
        chat: {
          id: 'newChatId',
          title: 'New chat',
          createdAt: new Date(),
        },
      },
    });
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

function renderWithRouter(ui: JSX.Element, { route = '/app' } = {}) {
  return render(
    <AuthContextProvider>
      <QueryClientProvider client={queryClient}>
        <MemoryRouter initialEntries={[route]}>
          <Routes>
            <Route path="/app" element={ui} />
            <Route path="/app/chats/:chatId" element={<div>Chat</div>} />
            <Route path="/app/login" element={<div>Sign in</div>} />
          </Routes>
        </MemoryRouter>
      </QueryClientProvider>
    </AuthContextProvider>,
  );
}

test('renders Navbar component', async () => {
  renderWithRouter(<Sidebar />, { route: '/app' });

  expect(screen.getByLabelText('New chat')).toBeInTheDocument();
  expect(screen.getByText('Sign out')).toBeInTheDocument();

  await waitFor(() => {
    expect(screen.getByText('Some chat')).toBeInTheDocument();
  });
});

test('navigates to chat on chat link click', async () => {
  renderWithRouter(<Sidebar />, { route: '/app' });

  await waitFor(async () => {
    expect(screen.getByText('Some chat')).toBeInTheDocument();
    await userEvent.click(screen.getByText('Some chat'));
    expect(screen.getByText('Chat')).toBeInTheDocument();
  });
});

test('calls handleNewChat on new chat button click', async () => {
  renderWithRouter(<Sidebar />, { route: '/app' });

  await waitFor(async () => {
    await userEvent.click(screen.getByLabelText('New chat'));
    expect(screen.getByText('Chat')).toBeInTheDocument();
  });
});

test('calls handleSignOut on sign out button click', async () => {
  renderWithRouter(<Sidebar />, { route: '/app' });

  await waitFor(async () => {
    await userEvent.click(screen.getByText('Sign out'));
    // expect(screen.getByText('Sign in')).toBeInTheDocument();
  });
});
