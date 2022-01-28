import Services from "../pages/Services";
import Inviting from "../pages/Inviting";
import Folder from "../pages/Folder";

export const privateRoutes = [
    {path: '/', element: Services},
    {path: '/inviting', element: Inviting},
    {path: '/inviting/:folderID', element: Folder},
]

export const publicRoutes = [
    {path: '/', element: Services},
]