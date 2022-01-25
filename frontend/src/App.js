import React from 'react'
import './styles/App.css';
import {BrowserRouter} from "react-router-dom"
import AppRoute from './components/AppRoute';

function App() {
	return (
		<BrowserRouter>
			<AppRoute/>
		</BrowserRouter>
	);
}

export default App;
