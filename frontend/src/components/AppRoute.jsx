import React from 'react'
import {Routes, Route} from "react-router-dom"
import ErrorPage from '../pages/ErrorPage';
import { privateRoutes, publicRoutes } from '../router';

const AppRoute = () => {
    return (
        <Routes>
            {privateRoutes.map(route =>
                <Route 
                    path={route.path}
                    element={<route.element/>}
                    key={route.path}
                />
            )}
            {publicRoutes.map(route =>
                <Route 
                    path={route.path}
                    element={<route.element/>}
                    key={route.path}
                />
            )}
            <Route  path="*" element={<ErrorPage />} />
        </Routes>

    );
};

export default AppRoute;
