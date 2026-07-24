import {
BrowserRouter,
Routes,
Route,
} from "react-router-dom";


import ProtectedRoute from "./auth/ProtectedRoute";
import AppLayout from "./layout/AppLayout";


import Dashboard from "./pages/Dashboard";
import Profile from "./pages/Profile";
import Vehicles from "./pages/Vehicles";
import PartsSearch from "./pages/PartsSearch";
import Login from "./pages/Login";


export default function App(){

return (

<BrowserRouter>

<Routes>

<Route
  path="/login"
  element={<Login />}
/>

<Route
path="/"
element={
<ProtectedRoute>
<AppLayout />
</ProtectedRoute>
}
>

<Route
index
element={<Dashboard />}
/>


<Route
path="profile"
element={<Profile />}
/>


<Route
path="vehicles"
element={<Vehicles />}
/>


<Route
path="parts"
element={<PartsSearch />}
/>


</Route>


</Routes>

</BrowserRouter>

);

}