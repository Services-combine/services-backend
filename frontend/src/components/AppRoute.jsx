import React, {useContext} from 'react'
import {Routes, Route} from "react-router-dom"
import {Context} from "../index";
import NotFound from '../pages/NotFound';
import { privateRoutes, publicRoutes } from '../router';
import { observer } from 'mobx-react-lite';
import Loader from '../components/UI/loader/Loader';

const AppRoute = () => {
    const {store} = useContext(Context);

    if (store.isLoading) {
        return (
            <div style={{display: 'flex', justifyContent: 'center', marginTop: 50}}>
                <Loader/>
            </div>
        );
    }

    return (
        store.isAuth
            ?
            <Routes>
                {privateRoutes.map(route =>
                    <Route 
                        path={route.path}
                        element={<route.element/>}
                        key={route.path}
                    />
                )}
                <Route  path="*" element={<NotFound />} />
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
                <Route  path="*" element={<NotFound />} />
            </Routes>
    );
};

export default observer(AppRoute);
