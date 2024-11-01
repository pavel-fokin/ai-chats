import { fireEvent, render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';

import { TextArea } from './text-area';

test('renders TextArea component', () => {
  render(<TextArea placeholder="Enter text" value="" onChange={() => {}} />);

  const textAreaElement = screen.getByPlaceholderText('Enter text');
  expect(textAreaElement).toBeInTheDocument();
});

test('calls onChange callback when text is entered', () => {
  const handleChange = vi.fn();

  render(
    <TextArea placeholder="Enter text" value="" onChange={handleChange} />,
  );

  const textAreaElement = screen.getByPlaceholderText('Enter text');
  fireEvent.change(textAreaElement, { target: { value: 'Hello, World!' } });

  expect(handleChange).toHaveBeenCalledTimes(1);
  expect(handleChange).toHaveBeenCalledWith(expect.any(Object));
});

test('updates text value when text is entered', async () => {
  render(<TextArea placeholder="Enter text" value="" onChange={() => {}} />);

  const textAreaElement = screen.getByPlaceholderText('Enter text');
  fireEvent.change(textAreaElement, { target: { value: 'Hello, World!' } });
  userEvent.type(textAreaElement, 'Hello, World!');

  expect(textAreaElement).toHaveValue('Hello, World!');
});
