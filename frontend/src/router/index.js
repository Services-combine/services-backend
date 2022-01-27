import Login from "../components/Login";
import Inviting from "../pages/Inviting";
import Services from "../pages/Services";

export const privateRoutes = [
    {path: '/', element: Services},
    {path: '/inviting', element: Inviting},
]

export const publicRoutes = [
    {path: '/', element: Services},
]