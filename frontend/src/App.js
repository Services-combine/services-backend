import React, {useContext, useEffect} from 'react'
import {Context} from "./index";
import './styles/App.css';
import AppRoute from './components/AppRoute';

function App() {
	const {store} = useContext(Context);

	useEffect(() => {
		if (localStorage.getItem('token')) {
			store.checkAuth()
		}
	}, [])

	return (
		<div className="container">
			<AppRoute/>
		</div>
	);
}

export default App;
