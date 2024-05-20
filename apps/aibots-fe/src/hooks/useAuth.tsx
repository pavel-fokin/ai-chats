import { useState } from "react";

import { postLogIn, postSignUp } from "api";

// useAuth custom hook
export const useAuth = () => {
    const [isLoading, setIsLoading] = useState(false);
    const [accessToken, setAccessToken] = useState(localStorage.getItem("accessToken"));

    const logIn = async (username: string, password: string) => {
        setIsLoading(true);
        try {
            const { accessToken: token } = await postLogIn(username, password);
            if (!token) {
                return false;
            }

            setAccessToken(token);
            localStorage.setItem("accessToken", token);

            setIsLoading(false);
            return true;
        } catch (error) {
            console.error(error);
            setIsLoading(false);
            return false;
        }
    };

    const signUp = async (username: string, password: string) => {
        setIsLoading(true);
        try {
            const { accessToken: token } = await postSignUp(username, password);
            if (!token) {
                return false;
            }

            setAccessToken(token);
            localStorage.setItem("accessToken", token);

            setIsLoading(false);
            return true;
        } catch (error) {
            console.error(error);
            setIsLoading(false);
            return false;
        }
    }

    const signOut = () => {
        setAccessToken(null);
        localStorage.removeItem("accessToken");
    }

    return {
        accessToken,
        isLoading,
        logIn,
        signUp,
        signOut,
    };
};