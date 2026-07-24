import {
  useEffect,
  useState,
} from "react";

import api from "../api/client";


interface Profile {

  user_id: string;

  role: string;

}



export default function Profile() {


  const [profile, setProfile] =
    useState<Profile | null>(
      null
    );


  const [error, setError] =
    useState("");



  useEffect(() => {


    async function load() {


      try {


        const response =
          await api.get(
            "/users/me"
          );


        setProfile(
          response.data
        );


      } catch {


        setError(
          "Unable to load profile"
        );


      }


    }


    load();


  }, []);



  if (error) {

    return <p>{error}</p>;

  }



  if (!profile) {

    return <p>Loading...</p>;

  }



  return (

    <div>

      <h1>
        Profile
      </h1>


      <p>
        User ID:
        {" "}
        {profile.user_id}
      </p>


      <p>
        Role:
        {" "}
        {profile.role}
      </p>


    </div>

  );

}