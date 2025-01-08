import { useEffect, useState } from 'react';

const defaultDelay = 1000;

// Debounces a value.
export const useDebounce = <T>(value: T, delay: number = defaultDelay): T => {
  const [debouncedValue, setDebouncedValue] = useState<T>(value);

  useEffect(() => {
    const timer = setTimeout(() => setDebouncedValue(value), delay);

    return () => {
      clearTimeout(timer);
    };
  }, [value, delay]);

  return debouncedValue;
};
