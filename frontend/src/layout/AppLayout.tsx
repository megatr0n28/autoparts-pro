import {
  Outlet,
} from "react-router-dom";

import Sidebar from "../components/Sidebar";
import Navbar from "../components/Navbar";


export default function AppLayout() {

  return (
    <div>

      <Navbar />

      <div
        style={{
          display:"flex",
        }}
      >

        <Sidebar />

        <main
          style={{
            padding:"20px",
            flex:1,
          }}
        >
          <Outlet />
        </main>


      </div>

    </div>
  );
}