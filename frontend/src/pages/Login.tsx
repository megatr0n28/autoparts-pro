import {
  useState,
} from "react";

import type {
  FormEvent,
} from "react";

import {
  useNavigate,
} from "react-router-dom";

import {
  useAuth,
} from "../auth/AuthContext";


export default function Login() {

  const {
    login,
  } = useAuth();


  const navigate =
    useNavigate();


  const [email, setEmail] =
    useState("");

  const [password, setPassword] =
    useState("");


  async function submit(
    e: FormEvent
  ) {

    e.preventDefault();


    await login(
      email,
      password
    );


    navigate("/");

  }


  return (

    <form onSubmit={submit}>

      <h1>
        Login
      </h1>


      <input
        placeholder="Email"
        value={email}
        onChange={
          e => setEmail(e.target.value)
        }
      />


      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={
          e => setPassword(e.target.value)
        }
      />


      <button type="submit">
        Login
      </button>

    </form>

  );
}