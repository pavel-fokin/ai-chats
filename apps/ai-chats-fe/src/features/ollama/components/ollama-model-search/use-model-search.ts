import { ChangeEvent, useState } from "react";

import { useDebounce } from "@/hooks";


export const useModelSearch = () => {
  const [searchValue, setSearchValue] = useState('');

  const debouncedSearchValue = useDebounce(searchValue);

  const handleSearchChange = (event: ChangeEvent<HTMLInputElement>) => {
    setSearchValue(event.target.value);
  };

  return {
    searchValue,
    debouncedSearchValue,
    handleSearchChange,
  };
};
