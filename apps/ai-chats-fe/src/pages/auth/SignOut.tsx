import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import { useAuthContext } from 'features/auth';

export const SignOut = () => {
  const navigate = useNavigate();
  const { setIsAuthenticated } = useAuthContext();

  useEffect(() => {
    setIsAuthenticated(false);
    localStorage.removeItem('accessToken');
    navigate('/app/login');
  }, [navigate, setIsAuthenticated]);

  return null;
};
