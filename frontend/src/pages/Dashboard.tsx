import {
useAuth,
} from "../auth/AuthContext";


export default function Dashboard(){

const {
logout,
}=useAuth();


return (

<div>

<h1>
AutoParts Pro Dashboard
</h1>


<p>
Authenticated user
</p>


<button onClick={logout}>
Logout
</button>


</div>

);

}