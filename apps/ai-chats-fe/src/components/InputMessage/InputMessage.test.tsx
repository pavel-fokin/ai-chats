import { fireEvent, render, screen } from '@testing-library/react';
import { Theme } from '@radix-ui/themes';
import { vi } from 'vitest';

import { InputMessage } from 'components';

const renderComponent = (children: React.ReactNode) => {
  return render(<Theme>{children}</Theme>);
};

test('updates input message text on change', () => {
  renderComponent(<InputMessage handleSend={() => {}} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLInputElement;
  const newText = 'Hello, world!';

  fireEvent.change(inputElement, { target: { value: newText } });

  expect(inputElement.value).toBe(newText);
});

test('calls handleSend with correct message on send click', () => {
  const handleSendMock = vi.fn();
  const text = 'Hello, world!';

  renderComponent(<InputMessage handleSend={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLTextAreaElement;
  const sendButton = screen.getByRole('button');

  fireEvent.change(inputElement, { target: { value: text } });
  fireEvent.click(sendButton);

  expect(handleSendMock).toHaveBeenCalledWith(text);
  expect(inputElement.value).toBe('Hello, world!');
});

test('calls handleSend with correct message on enter key press', () => {
  const handleSendMock = vi.fn();
  const text = 'Hello, world!';

  renderComponent(<InputMessage handleSend={handleSendMock} />);

  const inputElement = screen.getByRole('textbox') as HTMLTextAreaElement;

  fireEvent.change(inputElement, { target: { value: text } });
  fireEvent.keyDown(inputElement, { key: 'Enter', code: 13, charCode: 13 });

  expect(handleSendMock).toHaveBeenCalledWith(text);
  expect(inputElement.value).toBe('Hello, world!');
});

test('does not call handleSend on empty message send click', () => {
  const handleSendMock = vi.fn();

  renderComponent(<InputMessage handleSend={handleSendMock} />);

  const sendButton = screen.getByRole('button');

  fireEvent.click(sendButton);

  expect(handleSendMock).not.toHaveBeenCalled();
});

test('does not call handleSend on whitespace message send click', () => {
  const handleSendMock = vi.fn();

  renderComponent(<InputMessage handleSend={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLTextAreaElement;
  const sendButton = screen.getByRole('button');

  fireEvent.change(inputElement, { target: { value: ' ' } });
  fireEvent.click(sendButton);

  expect(handleSendMock).not.toHaveBeenCalled();
});
