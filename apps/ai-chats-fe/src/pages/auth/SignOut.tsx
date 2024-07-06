import { useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import { AuthContext } from 'contexts';

export const SignOut = () => {
  const navigate = useNavigate();
  const { setIsAuthenticated } = useContext(AuthContext);

  useEffect(() => {
    setIsAuthenticated(false);
    localStorage.removeItem('accessToken');
    navigate('/app/login');
  }, []);

  return null;
};
