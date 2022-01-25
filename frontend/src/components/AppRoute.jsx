import React, { useContext } from 'react'
import {Routes, Route} from "react-router-dom"
import { AuthContext } from '../context';
import ErrorPage from '../pages/ErrorPage';
import Login from '../pages/Login';
import Posts from '../pages/Posts';
import { privateRoutes, publicRoutes } from '../router';
import Loader from './UI/loader/Loader';

const AppRoute = () => {
    const {isAuth, isLoading} = useContext(AuthContext);

    if (isLoading) {
        return <Loader/>
    }

    return (
        isAuth
            ?
            <Routes>
                {privateRoutes.map(route =>
                    <Route 
                        path={route.path}
                        element={<route.element/>}
                        key={route.path}
                    />
                )}
                <Route  path="*" element={<ErrorPage />} />
            </Routes>
            :
            <Routes>
                {publicRoutes.map(route =>
                    <Route 
                        path={route.path}
                        element={<route.element/>}
                        key={route.path}
                    />
                )}
                <Route  path="*" element={<Login />} />
            </Routes>

    );
};

export default AppRoute;
