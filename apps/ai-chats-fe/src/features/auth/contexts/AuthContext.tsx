import {
  ReactNode,
  createContext,
  useEffect,
  useState,
  useContext,
} from 'react';

type AuthContextValue = {
  isAuthenticated: boolean;
  setIsAuthenticated: (isAuthenticated: boolean) => void;
};

export const AuthContext = createContext<AuthContextValue>(
  {} as AuthContextValue,
);

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => useContext(AuthContext);

export const AuthContextProvider = ({ children }: { children: ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(
    localStorage.getItem('accessToken') ? true : false,
  );

  useEffect(() => {
    if (isAuthenticated) {
      setAuthTimeout();
    }
  }, [isAuthenticated]);

  const setAuthTimeout = () => {
    const accessToken = localStorage.getItem('accessToken');
    setTimeout(
      () => {
        localStorage.removeItem('accessToken');
        setIsAuthenticated(false);
      },
      // Check if accessToken is not null before passing it to getExpirationTime
      accessToken ? getExpirationTime(accessToken) : 0,
    );
  };

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated,
        setIsAuthenticated,
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
