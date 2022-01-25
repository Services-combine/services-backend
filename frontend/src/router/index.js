import Inviting from "../pages/Inviting";
import Login from "../pages/Login";
import Services from "../pages/Services";

export const privateRoutes = [
    {path: '/', element: Services},
    {path: '/inviting', element: Inviting},
]

export const publicRoutes = [
    {path: '/login', element: Login},
]