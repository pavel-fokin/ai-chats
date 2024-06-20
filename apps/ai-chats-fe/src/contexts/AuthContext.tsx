import { ReactNode, createContext, useEffect, useState } from 'react';

import { useAuth } from 'hooks';

type AuthContextValue = {
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (username: string, password: string) => Promise<boolean>;
  signup: (username: string, password: string) => Promise<boolean>;
  signout: () => void;
};

export const AuthContext = createContext<AuthContextValue>(
  {} as AuthContextValue,
);

// AuthContextProvider component that wraps the application and provides the AuthContext.
export const AuthContextProvider = ({ children }: { children: ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(
    localStorage.getItem('accessToken') ? true : false,
  );

  const { isLoading, logIn, signUp, signOut } = useAuth();

  useEffect(() => {
    if (isAuthenticated) {
      setAuthTimeout();
    }
  }, [isAuthenticated]);

  const setAuthTimeout = () => {
    const accessToken = localStorage.getItem('accessToken');
    setTimeout(
      () => {
        signout();
      },
      // Check if accessToken is not null before passing it to getExpirationTime
      accessToken ? getExpirationTime(accessToken) : 0,
    );
  };

  const login = async (username: string, password: string) => {
    const isLoggedIn = await logIn(username, password);
    setIsAuthenticated(isLoggedIn);
    return isLoggedIn;
  };

  const signup = async (username: string, password: string) => {
    const signedUp = await signUp(username, password);
    setIsAuthenticated(signedUp);
    return signedUp;
  };

  const signout = () => {
    signOut();
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated,
        isLoading,
        login,
        signup,
        signout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

const getExpirationTime = (token: string) => {
  const decodedToken = JSON.parse(atob(token.split('.')[1])); // Decode JWT payload
  const expiresAt = decodedToken.exp * 1000; // Convert to milliseconds
  return expiresAt - Date.now(); // Time until expiration
};
