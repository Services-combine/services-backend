import Services from "../pages/Services";
import Inviting from "../pages/Inviting";
import Folder from "../pages/Folder";
import Account from "../pages/Account";

export const privateRoutes = [
    {path: '/', element: Services},
    {path: '/inviting', element: Inviting},
    {path: '/inviting/:folderID', element: Folder},
    {path: '/inviting/:folderID/:accountID', element: Account},
]

export const publicRoutes = [
    {path: '/', element: Services},
]