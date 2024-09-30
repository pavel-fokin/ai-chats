import { IconButton } from 'shared/components';
import { MenuIcon } from 'shared/components/icons';

export const MenuButton = () => {
  return (
    <IconButton
      aria-label="Chat menu button"
      variant="ghost"
      size="2"
      highContrast
    >
      {' '}
      <MenuIcon size="24" />{' '}
    </IconButton>
  );
};
