import { Theme } from '@radix-ui/themes';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

import { Tooltip } from './Tooltip';

test.skip('renders tooltip content when hovered', async () => {
  render(
    <Theme>
      <Tooltip content="This is a tooltip">
        <button>Hover me</button>
      </Tooltip>
    </Theme>
  );

  const button = screen.getByRole('button', { name: 'Hover me' });

  await userEvent.hover(button);

  await waitFor(() => {
    expect(screen.getByText('This is a tooltip')).toBeInTheDocument();
  });
});

test('does not render tooltip content when not hovered', () => {
  render(
    <Theme>
      <Tooltip content="This is a tooltip">
        <button>Hover me</button>
      </Tooltip>
    </Theme>
  );

  expect(screen.getByRole('button', { name: 'Hover me' })).toBeInTheDocument;
  expect(screen.queryByText('This is a tooltip')).not.toBeInTheDocument();
});
