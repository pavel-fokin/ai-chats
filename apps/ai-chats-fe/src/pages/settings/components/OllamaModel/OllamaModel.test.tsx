import { render, screen, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import userEvent from '@testing-library/user-event';

import { OllamaModel } from './OllamaModel';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
    },
  },
});

const mockModel = {
  model: 'llama3:latest',
  description:
    'Meta Llama 3: The most capable openly available LLM to date 8B.',
  isPulling: false,
};

const renderComponent = (ui: JSX.Element) => {
  return render(
    <QueryClientProvider client={queryClient}>{ui}</QueryClientProvider>,
  );
};

test('renders model name and tag', () => {
  renderComponent(<OllamaModel model={mockModel} />);

  expect(screen.getByText('llama3:latest')).toBeInTheDocument();
  expect(
    screen.getByText(
      'Meta Llama 3: The most capable openly available LLM to date 8B.',
    ),
  ).toBeInTheDocument();
});

test('shows and hides delete dialog', async () => {
  renderComponent(<OllamaModel model={mockModel} />);

  expect(screen.queryByText('Delete model?')).toBeNull();

  const deleteButton = screen.getByRole('button', { name: 'Delete model' });
  userEvent.click(deleteButton);

  await waitFor(() => {
    expect(screen.getByText('Delete model?')).toBeInTheDocument();
  });

  const cancelButton = screen.getByRole('button', { name: 'Cancel' });
  userEvent.click(cancelButton);

  await waitFor(() => {
    expect(screen.queryByText('Delete model?')).toBeNull();
  });
});
