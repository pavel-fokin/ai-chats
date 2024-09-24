import { Navigate } from 'react-router-dom';

import { useAuthContext } from 'features/auth';

interface AuthRequiredProps {
  children: React.ReactNode;
}

export const AuthRequired: React.FC<AuthRequiredProps> = ({ children }) => {
  const { isAuthenticated } = useAuthContext();
  if (!isAuthenticated) {
    return <Navigate to="/app/login" />;
  }
  return children;
};
