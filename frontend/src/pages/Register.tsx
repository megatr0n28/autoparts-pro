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

import type {
  AxiosError,
} from "axios";


export default function Register() {

  const {
    register,
  } = useAuth();


  const navigate =
    useNavigate();


  const [firstName, setFirstName] =
    useState("");

  const [lastName, setLastName] =
    useState("");

  const [email, setEmail] =
    useState("");

  const [password, setPassword] =
    useState("");


  const [error, setError] =
    useState("");


  const [loading, setLoading] =
    useState(false);


  async function submit(
    e: FormEvent
  ) {

    e.preventDefault();

    setError("");

    setLoading(true);


    try {

      await register({
        first_name: firstName,
        last_name: lastName,
        email,
        password,
      });


      navigate("/login");


    } catch (err) {

      const error =
        err as AxiosError<{
          error?: string;
        }>;


      setError(
        error.response?.data?.error ??
        "Registration failed"
      );

    } finally {

      setLoading(false);

    }

  }


  return (

    <div>

      <h1>
        Create Account
      </h1>


      <form
        onSubmit={submit}
      >


        <div>

          <input
            type="text"
            placeholder="First Name"
            value={firstName}
            onChange={
              e => setFirstName(
                e.target.value
              )
            }
            required
          />

        </div>


        <div>

          <input
            type="text"
            placeholder="Last Name"
            value={lastName}
            onChange={
              e => setLastName(
                e.target.value
              )
            }
            required
          />

        </div>


        <div>

          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={
              e => setEmail(
                e.target.value
              )
            }
            required
          />

        </div>


        <div>

          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={
              e => setPassword(
                e.target.value
              )
            }
            required
            minLength={8}
          />

        </div>


        {
          error && (

            <p>
              {error}
            </p>

          )
        }


        <button
          type="submit"
          disabled={loading}
        >

          {
            loading
              ? "Creating..."
              : "Register"
          }

        </button>


      </form>


      <p>

        Already have an account?

        <button
          type="button"
          onClick={() => navigate("/login")}
        >
          Login
        </button>

      </p>


    </div>

  );
}