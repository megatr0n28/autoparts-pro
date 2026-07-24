import { useState } from "react";
import type { FormEvent } from "react";

import { useAuth } from "../auth/AuthContext";

export default function Register() {
  const { register } = useAuth();

  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function submit(e: FormEvent) {
    e.preventDefault();

    await register({
      first_name: firstName,
      last_name: lastName,
      email,
      password,
    });

    alert("Registration successful");
  }

  return (
    <form onSubmit={submit}>
      <h1>Register</h1>

      <input
        placeholder="First Name"
        value={firstName}
        onChange={(e) => setFirstName(e.target.value)}
      />

      <input
        placeholder="Last Name"
        value={lastName}
        onChange={(e) => setLastName(e.target.value)}
      />

      <input
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <button type="submit">Register</button>
    </form>
  );
}