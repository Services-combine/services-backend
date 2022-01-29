import React, {useContext} from 'react'
import {Link} from "react-router-dom"
import '../styles/Services.css';
import {Context} from "../index";
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';

const ListServices = () => {
    const {store} = useContext(Context)

	return (
        <div className='services'>
            <div className='header'>
                <h3 className='logo'>Сервисы</h3>
                <div className='header__btns'>
                    <Button> <i className="fas fa-cog"></i> Настройки</Button>
                    <Button onClick={() => store.logout()}><i className="fas fa-sign-out-alt"></i> Выйти</Button>
                </div>
            </div>

            <div className="services__list btn-toolbar" role="toolbar">
                <Link to="/inviting" className="services__list-item">
                    <h6 className="services__list-item__title">Инвайтинг & Рассылка</h6>
                </Link>
            </div>

            {store.isError &&
                <Error>{store.isError}</Error>
            }
        </div>
	);
}

export default ListServices;
