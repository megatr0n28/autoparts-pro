import {
NavLink,
} from "react-router-dom";


export default function Sidebar(){

return (

<aside
style={{
width:"200px",
padding:"20px",
}}
>


<nav>

<p>
<NavLink to="/">
Dashboard
</NavLink>
</p>


<p>
<NavLink to="/profile">
Profile
</NavLink>
</p>


<p>
<NavLink to="/vehicles">
Vehicles
</NavLink>
</p>


<p>
<NavLink to="/parts">
Parts Search
</NavLink>
</p>


</nav>


</aside>

);

}