import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

import { NewChatIconButton } from 'components';
import { SidebarContextProvider } from 'features/sidebar';

const renderWithRouter = (ui: JSX.Element) => {
  return render(
    <SidebarContextProvider>
      <MemoryRouter initialEntries={['/']}>
        <Routes>
          <Route path="/" element={ui} />
          <Route path="/app/new-chat" element={<div>Start a new chat</div>} />
        </Routes>
      </MemoryRouter>
    </SidebarContextProvider>,
  );
};

test('renders the NewChatIconButton component', () => {
  renderWithRouter(<NewChatIconButton />);
  const buttonElement = screen.getByRole('button');
  expect(buttonElement).toBeInTheDocument();
});

test('navigates to /app/new-chat on button click', async () => {
  renderWithRouter(<NewChatIconButton />);
  const buttonElement = screen.getByRole('button');
  await userEvent.click(buttonElement);
  expect(screen.getByText('Start a new chat')).toBeInTheDocument();
});
