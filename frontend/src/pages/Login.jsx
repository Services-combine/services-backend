import React, {useState, useContext} from 'react'
import '../styles/Login.css';
import {Context} from "../index";
import {useNavigate} from "react-router-dom"
import Button from '../components/UI/button/Button';
import Input from '../components/UI/input/Input';
import { observer } from 'mobx-react-lite';


const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const {store} = useContext(Context)

    let navigate = useNavigate();

    const login = async (e) => {
        e.preventDefault();
        store.login(username, password)
    }

    if (store.isAuth) {
        navigate("/")
    }

    return (
        <div className='login'>
            <form className="form-login">
                <h3 className='title'>Авторизация</h3>
                <div className="form-input">
                    <Input onChange={e => setUsername(e.target.value)} type='text' placeholder='Введите логин' />
                    <Input onChange={e => setPassword(e.target.value)} type='password' placeholder='Введите пароль' />
                </div>
                
                <Button onClick={login}>Войти</Button>
            </form>
        </div>
    );
};

export default observer(Login);
