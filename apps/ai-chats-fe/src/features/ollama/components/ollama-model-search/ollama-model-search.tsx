import { SearchIcon } from '@/components/icons';
import { TextField } from '@/components/ui';

import { useModelSearch } from './use-model-search';

export const OllamaModelSearch = () => {
  const { handleSearchChange } = useModelSearch();

  return (
    <TextField.Root
      id="model"
      size="3"
      placeholder="Search model"
      onChange={handleSearchChange}
    >
      <TextField.Slot>
        <SearchIcon />
      </TextField.Slot>
    </TextField.Root>
  );
};
