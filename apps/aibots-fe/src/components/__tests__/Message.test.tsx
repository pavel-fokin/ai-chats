import { render, screen } from '@testing-library/react';
import { Message } from '../Message';

test('renders message with sender and text', () => {
  const sender = 'John Doe';
  const text = 'Hello, world!';

  render(<Message sender={sender} text={text} />);

  const senderElement = screen.getByText(sender);
  const textElement = screen.getByText(text);

  expect(senderElement).toBeInTheDocument();
  expect(textElement).toBeInTheDocument();
});

test('renders fallback avatar when sender is not provided', () => {
  const text = 'Hello, world!';

  render(<Message sender="A" text={text} />);

  const fallbackAvatar = screen.getByText('A');
  const textElement = screen.getByText(text);

  expect(fallbackAvatar).toBeInTheDocument();
  expect(textElement).toBeInTheDocument();
});
