import { useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import { AuthContext } from 'contexts';

export const SignOut = () => {
  const navigate = useNavigate();
  const { signout } = useContext(AuthContext);

  useEffect(() => {
    signout();
    navigate('/app/login');
  }, []);

  return null;
};
