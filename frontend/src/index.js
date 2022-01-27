import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter} from "react-router-dom"
import { createContext } from 'react/cjs/react.production.min';
import App from './App';
import Store from './store/store';

const store = new Store()
export const Context = createContext({
	store,
})

ReactDOM.render(
	<Context.Provider value={{
		store
	}}>
		<BrowserRouter>
			<App/>
		</BrowserRouter>
	</Context.Provider>,
	document.getElementById('root')
);