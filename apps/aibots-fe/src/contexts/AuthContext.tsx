import { createContext, useState, ReactNode } from 'react';

import { useAuth } from 'hooks';

type AuthContextValue = {
    isAuthenticated: boolean;
    isLoading: boolean;
    setIsAuthenticated: (isAuthenticated: boolean) => void;
    login: (username: string, password: string) => Promise<boolean>;
    signup: (username: string, password: string) => Promise<boolean>;
    signout: () => void;
}

export const AuthContext = createContext<AuthContextValue>({} as AuthContextValue);

// AuthContextProvider
export const AuthContextProvider = ({ children }: { children: ReactNode }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(
        localStorage.getItem("accessToken") ? true : false
    );

    const { isLoading, logIn, signUp, signOut } = useAuth();

    const login = async (username: string, password: string) => {
        const isLoggedIn = await logIn(username, password);
        setIsAuthenticated(isLoggedIn);
        return isLoggedIn;
    }

    const signup = async (username: string, password: string) => {
        const signedUp = await signUp(username, password);
        setIsAuthenticated(signedUp);
        return signedUp;
    }

    const signout = () => {
        signOut();
        setIsAuthenticated(false);
    }

    return (
        <AuthContext.Provider value={{
            isAuthenticated,
            isLoading,
            setIsAuthenticated,
            login,
            signup,
            signout
        }}>
            {children}
        </AuthContext.Provider>
    );
};