import { fireEvent, render, screen } from '@testing-library/react';
import { Theme } from '@radix-ui/themes';
import { vi } from 'vitest';

import { InputMessage } from './input-message';

const renderComponent = (children: React.ReactNode) => {
  return render(<Theme>{children}</Theme>);
};

test('updates input message text on change', () => {
  renderComponent(<InputMessage onSendMessage={() => {}} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLTextAreaElement;
  const newText = 'Hello, world!';

  fireEvent.change(inputElement, { target: { value: newText } });

  expect(inputElement.value).toBe(newText);
});

test('calls handleSend with correct message on send click', () => {
  const handleSendMock = vi.fn();
  const text = 'Hello, world!';

  renderComponent(<InputMessage onSendMessage={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLTextAreaElement;
  const sendButton = screen.getByRole('button');

  fireEvent.change(inputElement, { target: { value: text } });
  expect(inputElement.value).toBe('Hello, world!');

  fireEvent.click(sendButton);
  expect(handleSendMock).toHaveBeenCalledWith(text);
  expect(inputElement.value).toBe('');
});

test('calls handleSend with correct message on enter key press', () => {
  const handleSendMock = vi.fn();
  const text = 'Hello, world!';

  renderComponent(<InputMessage onSendMessage={handleSendMock} />);

  const inputElement = screen.getByRole('textbox') as HTMLTextAreaElement;

  fireEvent.change(inputElement, { target: { value: text } });
  expect(inputElement.value).toBe('Hello, world!');

  fireEvent.keyDown(inputElement, { key: 'Enter', code: 13, charCode: 13 });
  expect(handleSendMock).toHaveBeenCalledWith(text);
  expect(inputElement.value).toBe('');
});

test('does not call handleSend on empty message send click', () => {
  const handleSendMock = vi.fn();

  renderComponent(<InputMessage onSendMessage={handleSendMock} />);

  const sendButton = screen.getByRole('button');

  fireEvent.click(sendButton);

  expect(handleSendMock).not.toHaveBeenCalled();
});

test('does not call handleSend on whitespace message send click', () => {
  const handleSendMock = vi.fn();

  renderComponent(<InputMessage onSendMessage={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLTextAreaElement;
  const sendButton = screen.getByRole('button');

  fireEvent.change(inputElement, { target: { value: ' ' } });
  fireEvent.click(sendButton);

  expect(handleSendMock).not.toHaveBeenCalled();
});
