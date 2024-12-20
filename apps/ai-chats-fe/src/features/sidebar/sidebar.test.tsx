import { Theme } from '@radix-ui/themes';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { HttpResponse, http } from 'msw';
import { setupServer } from 'msw/node';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { AuthProvider } from 'features/auth';
import { Sidebar } from './sidebar';

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

function renderWithRouter(ui: JSX.Element, { route = '/' } = {}) {
  return render(
    <Theme>
      <AuthProvider>
        <QueryClientProvider client={queryClient}>
          <MemoryRouter initialEntries={[route]}>
            <Routes>
              <Route path="/" element={ui} />
              <Route
                path="/app/new-chat"
                element={<div>Start a new chat</div>}
              />
              <Route path="/app/chats/:chatId" element={<div>Chat</div>} />
            </Routes>
          </MemoryRouter>
        </QueryClientProvider>
      </AuthProvider>
    </Theme>,
  );
}

test('renders Sidebar component', async () => {
  renderWithRouter(<Sidebar />, { route: '/' });

  expect(screen.getByLabelText('Start a new chat')).toBeInTheDocument();
  expect(screen.getByLabelText('Sign out')).toBeInTheDocument();

  await waitFor(() => {
    expect(screen.getByText('Some chat')).toBeInTheDocument();
  });
});

test('navigates to chat on chat link click', async () => {
  renderWithRouter(<Sidebar />, { route: '/' });

  await waitFor(async () => {
    expect(screen.getByText('Some chat')).toBeInTheDocument();
    await userEvent.click(screen.getByText('Some chat'));
    expect(screen.getByText('Chat')).toBeInTheDocument();
  });
});

test('calls handleNewChat on new chat button click', async () => {
  renderWithRouter(<Sidebar />, { route: '/' });

  await waitFor(async () => {
    await userEvent.click(screen.getByLabelText('Start a new chat'));
    expect(screen.getByText('Start a new chat')).toBeInTheDocument();
  });
});

test('calls handleSignOut on sign out button click', async () => {
  renderWithRouter(<Sidebar />, { route: '/' });

  await waitFor(async () => {
    await userEvent.click(screen.getByLabelText('Sign out'));
    // expect(screen.getByText('Sign in')).toBeInTheDocument();
  });
});
