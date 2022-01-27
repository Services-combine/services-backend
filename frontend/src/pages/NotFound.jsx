import React from 'react'
import {Link} from "react-router-dom"
import '../styles/NotFound.css';

const NotFound = () => {
    return (
        <div className="not-foung-page">
            <div className="error mx-auto" data-text="404">404</div>
            <p className="lead text-gray-800 mb-5">Page Not Found</p>
            <p className="again text-gray-500 mb-0">
                Вернуться на <Link to="/" className='again-link'>главную</Link> страницу
            </p>
        </div>
    );
};

export default NotFound;
