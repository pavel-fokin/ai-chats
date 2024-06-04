import { fireEvent, render, screen } from '@testing-library/react';
import { vi } from 'vitest';

import { InputMessage } from '../InputMessage';

test('updates input message text on change', () => {
  render(<InputMessage handleSend={() => {}} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLInputElement;
  const newText = 'Hello, world!';

  fireEvent.change(inputElement, { target: { value: newText } });

  expect(inputElement.value).toBe(newText);
});

test('calls handleSend with correct message on send click', () => {
  const handleSendMock = vi.fn();
  const sender = 'User';
  const text = 'Hello, world!';

  render(<InputMessage handleSend={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLInputElement;
  const sendButton = screen.getByRole('button');

  fireEvent.change(inputElement, { target: { value: text } });
  fireEvent.click(sendButton);

  expect(handleSendMock).toHaveBeenCalledWith({ sender, text });
  expect(inputElement.value).toBe('');
});

test('calls handleSend with correct message on enter key press', () => {
  const handleSendMock = vi.fn();
  const sender = 'User';
  const text = 'Hello, world!';

  render(<InputMessage handleSend={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLInputElement;
  const form = screen.getByRole('form');

  fireEvent.change(inputElement, { target: { value: text } });
  // fireEvent.keyDown(inputElement, { key: 'Enter', code: 13, charCode: 13 });
  fireEvent.submit(form);

  expect(handleSendMock).toHaveBeenCalledWith({ sender, text });
  expect(inputElement.value).toBe('');
});

test('does not call handleSend on empty message send click', () => {
  const handleSendMock = vi.fn();

  render(<InputMessage handleSend={handleSendMock} />);

  const sendButton = screen.getByRole('button');

  fireEvent.click(sendButton);

  expect(handleSendMock).not.toHaveBeenCalled();
});

test('does not call handleSend on whitespace message send click', () => {
  const handleSendMock = vi.fn();

  render(<InputMessage handleSend={handleSendMock} />);

  const inputElement = screen.getByPlaceholderText(
    'Type a message',
  ) as HTMLInputElement;
  const sendButton = screen.getByRole('button');

  fireEvent.change(inputElement, { target: { value: ' ' } });
  fireEvent.click(sendButton);

  expect(handleSendMock).not.toHaveBeenCalled();
});
