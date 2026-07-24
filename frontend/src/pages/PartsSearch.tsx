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


interface Part {

  id: string;

  name: string;

  part_number: string;

  description?: string;

}



export default function PartsSearch() {


  const [vehicles, setVehicles] =
    useState<Vehicle[]>([]);


  const [vehicleId, setVehicleId] =
    useState<string>("");


  const [query, setQuery] =
    useState<string>("");


  const [parts, setParts] =
    useState<Part[]>([]);


  const [loading, setLoading] =
    useState(false);


  const [error, setError] =
    useState<string>("");



  //
  // Load customer vehicles
  //
  useEffect(() => {


    async function loadVehicles() {

      try {

        const response =
          await api.get(
            "/vehicles"
          );


        setVehicles(
          response.data
        );


      } catch (err) {

        console.error(
          "Failed loading vehicles",
          err
        );


        setError(
          "Unable to load vehicles"
        );

      }

    }


    loadVehicles();


  }, []);



  //
  // Search compatible parts
  //
  async function searchParts() {


    if (!vehicleId) {

      setError(
        "Please select a vehicle"
      );

      return;

    }


    if (!query) {

      setError(
        "Please enter a search term"
      );

      return;

    }


    try {

      setLoading(true);

      setError("");


      const response =
        await api.get(
          "/parts/search",
          {
            params: {

              vehicle_id:
                vehicleId,

              query:
                query,

            },
          }
        );


      setParts(
        response.data.results
      );


    } catch (err) {

      console.error(
        "Parts search failed",
        err
      );


      setError(
        "Unable to search parts"
      );


    } finally {

      setLoading(false);

    }

  }



  return (

    <div>

      <h1>
        Parts Search
      </h1>


      {
        error && (

          <p>
            {error}
          </p>

        )
      }



      <div>


        <label>
          Vehicle
        </label>


        <br />


        <select

          value={vehicleId}

          onChange={
            (e) =>
              setVehicleId(
                e.target.value
              )
          }

        >

          <option value="">
            Select Vehicle
          </option>


          {
            vehicles.map(
              (vehicle) => (

                <option

                  key={
                    vehicle.id
                  }

                  value={
                    vehicle.id
                  }

                >

                  {vehicle.year}
                  {" "}
                  {vehicle.make}
                  {" "}
                  {vehicle.model}


                </option>

              )
            )
          }


        </select>


      </div>



      <br />



      <div>


        <label>
          Search Parts
        </label>


        <br />


        <input

          value={query}

          onChange={
            (e) =>
              setQuery(
                e.target.value
              )
          }

          placeholder="Example: brake pad"

        />


        <button

          onClick={
            searchParts
          }

          disabled={
            loading
          }

        >

          {
            loading
              ? "Searching..."
              : "Search"
          }


        </button>


      </div>



      <hr />



      <h2>
        Results
      </h2>



      {
        parts.length === 0 && (

          <p>
            No parts found
          </p>

        )
      }



      {
        parts.map(
          (part) => (

            <div

              key={
                part.id
              }

            >

              <h3>
                {part.name}
              </h3>


              <p>
                Part Number:
                {" "}
                {part.part_number}
              </p>


              {
                part.description && (

                  <p>
                    {part.description}
                  </p>

                )
              }


            </div>

          )
        )
      }


    </div>

  );

}