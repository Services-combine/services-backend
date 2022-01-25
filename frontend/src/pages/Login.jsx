import React, {useState} from 'react'
import '../styles/Login.css';
import Button from '../components/UI/button/Button';
import Input from '../components/UI/input/Input';
import Loader from '../components/UI/loader/Loader';
import LoginService from '../API/LoginService';
import { useFetching } from '../hooks/useFetching';
import Error from '../components/UI/error/Error';


const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const [fetchLogin, isLoginLoading, loginError] = useFetching(async () => {
		const response = await LoginService.loginUser(username, password);
		//const resultAuth = response
        console.log(response);
	})

    const login = async (e) => {
        e.preventDefault();
        fetchLogin()
    }

    return (
        <div className='login'>
            <form onSubmit={login} className="form-login">
                <h3 className='title'>Авторизация</h3>
                <div className="form-input">
                    <Input onChange={e => setUsername(e.target.value)} type='text' placeholder='Введите логин' />
                    <Input onChange={e => setPassword(e.target.value)} type='password' placeholder='Введите пароль' />
                </div>
                
                <Button>Войти</Button>
            </form>

            {loginError &&
				<Error>Произошла ошибка: {loginError}</Error>
			}

            {isLoginLoading &&
				<div style={{display: 'flex', justifyContent: 'center', marginTop: 50}}><Loader/></div>
			}
        </div>
    );
};

export default Login;
