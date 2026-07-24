import {
  useAuth,
} from "../auth/AuthContext";


export default function Navbar(){

const {
 logout,
}=useAuth();


return (

<nav
style={{
padding:"15px",
borderBottom:"1px solid #ddd",
}}
>

<h2>
AutoParts Pro
</h2>


<button
onClick={logout}
>
Logout
</button>


</nav>

);

}