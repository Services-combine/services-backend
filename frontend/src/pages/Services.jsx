import React, {useEffect, useContext} from 'react'
import {Link} from "react-router-dom"
import '../styles/Services.css';
import {Context} from "../index";
import Loader from '../components/UI/loader/Loader';
import { observer } from 'mobx-react-lite';

const Services = () => {
    const {store} = useContext(Context)

    useEffect(() => {

	}, [])

    if (store.isLoading) {
        <div style={{display: 'flex', justifyContent: 'center', marginTop: 50}}><Loader/></div>
    }

	return (
        <div className='services'>
            <h3>Сервисы</h3>
            
            <ul className="services__list btn-toolbar" role="toolbar">
                <li className="services__list-item">
                    <Link to="/inviting" className="services__list-item_link">
                        <h6 className="title">Инвайтинг & Рассылка</h6>
                    </Link>
                </li>
            </ul>
        </div>
	);
}

export default observer(Services);
