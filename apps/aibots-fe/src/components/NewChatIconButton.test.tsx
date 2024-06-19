import { render, screen } from '@testing-library/react';
import { NewChatIconButton } from 'components';

test('renders the NewChatIconButton component', () => {
  render(<NewChatIconButton />);
  const buttonElement = screen.getByRole('button');
  expect(buttonElement).toBeInTheDocument();
});
