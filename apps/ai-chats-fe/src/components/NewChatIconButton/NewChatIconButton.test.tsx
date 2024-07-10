import { render, screen } from '@testing-library/react';
import { MemoryRouter, Route, Routes } from 'react-router-dom';
import userEvent from '@testing-library/user-event';

import { SidebarContextProvider } from 'contexts';
import { NewChatIconButton } from 'components';

const renderWithRouter = (ui: JSX.Element) => {
  return render(
    <SidebarContextProvider>
      <MemoryRouter initialEntries={['/']}>
        <Routes>
          <Route path="/" element={ui} />
          <Route path="/app" element={<div>App</div>} />
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

test('navigates to /app on button click', async () => {
  renderWithRouter(<NewChatIconButton />);
  const buttonElement = screen.getByRole('button');
  await userEvent.click(buttonElement);
  expect(screen.getByText('App')).toBeInTheDocument();
});
