import { useNavigate } from 'react-router-dom';

import { useAuthContext } from 'features/auth';

export const useSignOut = () => {
  const navigate = useNavigate();
  const { setIsAuthenticated } = useAuthContext();

  return () => {
    localStorage.removeItem('accessToken');
    setIsAuthenticated(false);
    navigate('/app/login');
  };
};
