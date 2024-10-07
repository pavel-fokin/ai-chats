import { IconButton } from 'components';
import { MenuIcon } from 'components/icons';

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
