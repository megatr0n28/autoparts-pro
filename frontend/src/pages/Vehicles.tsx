import {
  useEffect,
  useState,
} from "react";

import api from "../api/client";


interface Vehicle {

  id: string;

  year: number;

  make: string;

  model: string;

}



export default function Vehicles() {


  const [vehicles, setVehicles] =
    useState<Vehicle[]>([]);



  useEffect(() => {


    async function loadVehicles() {


      const response =
        await api.get(
          "/vehicles"
        );


      setVehicles(
        response.data
      );


    }


    loadVehicles();


  }, []);



  return (

    <div>

      <h1>
        Vehicles
      </h1>


      {
        vehicles.length === 0

        ?

        <p>
          No vehicles found
        </p>

        :

        vehicles.map(
          vehicle => (

            <div
              key={vehicle.id}
            >

              {vehicle.year}
              {" "}
              {vehicle.make}
              {" "}
              {vehicle.model}

            </div>

          )

        )

      }


    </div>

  );

}