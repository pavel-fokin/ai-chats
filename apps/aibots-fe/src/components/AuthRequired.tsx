import { useContext } from 'react';
import { Navigate } from 'react-router-dom';

import { AuthContext } from 'contexts';

type AuthRequiredProps = {
  children: React.ReactNode;
};

export const AuthRequired = ({ children }: AuthRequiredProps) => {
  const { isAuthenticated } = useContext(AuthContext);
  if (!isAuthenticated) {
    return <Navigate to="/app/login" />;
  }
  return children;
};
