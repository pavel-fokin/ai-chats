import { useContext } from 'react';
import { Navigate } from 'react-router-dom';

import { AuthContext } from 'contexts';

interface AuthRequiredProps {
  children: React.ReactNode;
}

export const AuthRequired: React.FC<AuthRequiredProps> = ({ children }) => {
  const { isAuthenticated } = useContext(AuthContext);
  if (!isAuthenticated) {
    return <Navigate to="/app/login" />;
  }
  return children;
};
