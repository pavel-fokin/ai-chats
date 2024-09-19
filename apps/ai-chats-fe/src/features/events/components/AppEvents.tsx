import { useAppEvents } from '../hooks/useAppEvents';

export const AppEvents: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  useAppEvents();
  return <>{children}</>;
};
