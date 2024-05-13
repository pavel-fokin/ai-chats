import { useState } from "react";

import { SignIn, SignUp } from "api";

// useAuth custom hook
export const useAuth = () => {
    const [accessToken, setAccessToken] = useState(localStorage.getItem("accessToken"));

    const signIn = async (username: string, password: string) => {
        const { accessToken: token } = await SignIn(username, password);
        if (!token) {
            return false;
        }
        setAccessToken(token);
        localStorage.setItem("accessToken", token);
        return true;
    };

    const signUp = async (username: string, password: string) => {
        const { accessToken: token } = await SignUp(username, password);
        if (!token) {
            return false;
        }
        setAccessToken(token);
        localStorage.setItem("accessToken", token);
        return true;
    }

    const signOut = () => {
        setAccessToken(null);
        localStorage.removeItem("accessToken");
    }

    return {
        accessToken,
        signIn,
        signUp,
        signOut,
    };
};