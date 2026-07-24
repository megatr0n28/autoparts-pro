import {
  createContext,
  useContext,
  useState,
} from "react";

import type { ReactNode } from "react";
import type { RegisterRequest } from "../types/auth";

import api from "../api/client";


interface AuthContextType {
  token: string | null;

  login(
    email: string,
    password: string
  ): Promise<void>;

  register(
    data: RegisterRequest
  ): Promise<void>;

  logout(): void;

  authenticated: boolean;
}


const AuthContext =
  createContext<AuthContextType | undefined>(
    undefined
  );


export function AuthProvider({
  children,
}: {
  children: ReactNode;
}) {

  const [token, setToken] =
    useState<string | null>(
      localStorage.getItem(
        "access_token"
      )
    );


  async function login(
    email: string,
    password: string
  ) {

    const response =
      await api.post(
        "/auth/login",
        {
          email,
          password,
        }
      );


    const accessToken =
      response.data.access_token;


    localStorage.setItem(
      "access_token",
      accessToken
    );


    setToken(accessToken);
  }


  async function register(
    data: RegisterRequest
  ): Promise<void> {
    await api.post("/auth/register", data);
  }


  function logout() {

    localStorage.removeItem(
      "access_token"
    );

    setToken(null);
  }


  return (
    <AuthContext.Provider
      value={{
        token,
        login,
        register,
        logout,
        authenticated:
          Boolean(token),
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}


export function useAuth() {

  const context =
    useContext(AuthContext);


  if (!context) {

    throw new Error(
      "useAuth must be inside AuthProvider"
    );

  }

  return context;
}