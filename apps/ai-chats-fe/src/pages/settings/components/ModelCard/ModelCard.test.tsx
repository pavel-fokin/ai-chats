import { render, screen } from '@testing-library/react';

import { ModelCard } from './ModelCard';

const mockModel = {
  model: 'llama3',
  description:
    'Meta Llama 3: The most capable openly available LLM to date 8B.',
};

test('renders model card', () => {
  render(<ModelCard model={mockModel} />);

  expect(screen.getByText('llama3')).toBeInTheDocument();
  expect(
    screen.getByText(
      'Meta Llama 3: The most capable openly available LLM to date 8B.',
    ),
  ).toBeInTheDocument();
});
