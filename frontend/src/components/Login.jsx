import React, {useState, useContext} from 'react'
import '../styles/Login.css';
import {Context} from "../index";
import Button from '../components/UI/button/Button';
import Input from '../components/UI/input/Input';
import Error from '../components/UI/error/Error';
import { observer } from 'mobx-react-lite';


const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const {store} = useContext(Context)

    return (
        <div className='login'>
            <div className="form-login">
                <h3>Аутентификация</h3>
                <div className="form-input">
                    <Input onChange={e => setUsername(e.target.value)} type='text' placeholder='Введите логин' />
                    <Input onChange={e => setPassword(e.target.value)} type='password' placeholder='Введите пароль' />
                </div>
                
                <Button onClick={() => store.login(username, password)}>Войти</Button>
            </div>

            {store.isError &&
                <Error>{store.isError}</Error>
            }
        </div>
    );
};

export default observer(Login);
