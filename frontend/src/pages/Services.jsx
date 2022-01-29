import React, {useContext} from 'react'
import { observer } from 'mobx-react-lite';
import {Context} from "../index";
import Login from './Login';
import ListServices from './ListServices';

const Services = () => {
    const {store} = useContext(Context)

	return (
        store.isAuth
            ?
            <ListServices/>
            : 
            <Login/>

	);
}

export default observer(Services);
