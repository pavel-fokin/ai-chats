import { useSignOut } from 'features/auth';

import { SignOutIcon } from 'shared/components/icons';

import styles from './SignOutButton.module.css';

export const SignOutButton = () => {
  const signOut = useSignOut();

  const handleSignOut = () => {
    signOut();
  };

  return (
    <a
      aria-label="Sign out"
      onClick={handleSignOut}
      className={styles.signOutButton}
    >
      <div className={styles.signOutButton__inner}>
        <SignOutIcon size={24} />
        Sign Out
      </div>
    </a>
  );
};
