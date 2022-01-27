import React, {useContext} from 'react'
import {Link} from "react-router-dom"
import '../styles/Services.css';
import {Context} from "../index";
import Button from '../components/UI/button/Button';

const Header = () => {
    const {store} = useContext(Context)

	return (
        <div className='header'>
            <Link to="/" className='logo'>Сервисы</Link>
            <Button onClick={() => store.logout()}>Выйти</Button>
        </div>
	);
}

export default Header;