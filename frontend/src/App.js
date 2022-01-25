import React, { useEffect, useState } from 'react'
import './styles/App.css';
import {BrowserRouter, Link} from "react-router-dom"
import AppRoute from './components/AppRoute';
import { AuthContext } from './context';
import MyButton from './components/UI/button/MyButton';

function App() {
	const [isAuth, setIsAuth] = useState(false); 
	const [isLoading, setIsLoading] = useState(true);

	useEffect(() => {
		if (localStorage.getItem('auth')) {
			setIsAuth(true)
		}
		setIsLoading(false);
	})

	const logout = () => {
		setIsAuth(false);
		localStorage.removeItem('auth')
	}

	return (
		<AuthContext.Provider value={{
			isAuth,
			setIsAuth,
			isLoading
		}}>
			<BrowserRouter>
				<div className='navbar'>
					<MyButton onClick={logout}>Выйти</MyButton>
					<div className='navbar__lings'>
						<Link to='/about'>О сайте</Link>
						<Link to='/posts'>Посты</Link>
					</div>
				</div>

				<AppRoute/>
			</BrowserRouter>
		</AuthContext.Provider>
	)
}

export default App;
