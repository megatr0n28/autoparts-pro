import {
  useEffect,
  useState,
} from "react";

import api from "../api/client";


interface HealthResponse {

  status: string;

}



export default function Dashboard() {


  const [health, setHealth] =
    useState<HealthResponse | null>(
      null
    );


  useEffect(() => {


    async function loadHealth() {

      const response =
        await api.get(
          "/health"
        );


      setHealth(
        response.data
      );

    }


    loadHealth();


  }, []);



  return (

    <div>

      <h1>
        Dashboard
      </h1>


      <p>
        AutoParts Pro API Status
      </p>


      <strong>

        {
          health
            ? health.status
            : "Checking..."
        }

      </strong>


    </div>

  );

}