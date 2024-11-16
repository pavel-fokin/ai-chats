import { renderWithRouter } from '@/pages/auth/login/LogIn.test';

import { OllamaLibrary } from './ollama-library';

test('renders OllamaLibrary component', async () => {
  renderWithRouter(<OllamaLibrary modelCards={[]} />);
});

