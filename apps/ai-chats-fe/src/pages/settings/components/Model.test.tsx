import { render, screen, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import userEvent from '@testing-library/user-event';
import { Model } from './Model';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
    },
  },
});

const mockModel = {
  id: '1',
  name: 'llama3',
  tag: 'latest',
};

const renderComponent = (ui: JSX.Element) => {
  return render(
    <QueryClientProvider client={queryClient}>{ui}</QueryClientProvider>,
  );
};

test('renders model name and tag', () => {
  renderComponent(<Model model={mockModel} />);

  expect(screen.getByText('llama3:latest')).toBeInTheDocument();
});

test('displays model description', () => {
  renderComponent(<Model model={mockModel} />);

  expect(
    screen.getByText(
      'Meta Llama 3: The most capable openly available LLM to date 8B.',
    ),
  ).toBeInTheDocument();
});

test('shows and hides gdelete dialog', async () => {
  renderComponent(<Model model={mockModel} />);

  expect(screen.queryByText('Delete model - llama3:latest')).toBeNull();

  const deleteButton = screen.getByRole('button', { name: 'Delete' });
  userEvent.click(deleteButton);

  await waitFor(() => {
    expect(
      screen.getByText('Delete model - llama3:latest'),
    ).toBeInTheDocument();
  });

  const cancelButton = screen.getByRole('button', { name: 'Cancel' });
  userEvent.click(cancelButton);

  await waitFor(() => {
    expect(screen.queryByText('Delete model - llama3:latest')).toBeNull();
  });
});