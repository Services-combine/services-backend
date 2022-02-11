import React, { useState } from 'react'
import '../styles/Header.css';

const Header = ({logo, children}) => {
    const [menuActive, setMenuActive] = useState(false);

    return (
        <header className='header'>
            <div className="container">
                <div className="header__body">
                    <h3 className="header__logo">{logo}</h3>

                    <div className={menuActive ? 'header__burger active' : 'header__burger'} onClick={() => setMenuActive(!menuActive)}>
                        <span></span>
                    </div>

                    <nav className={menuActive ? 'header__menu active' : 'header__menu'}>
                        <ul className="header__list">
                            {children}
                        </ul>
                    </nav>
                </div>
            </div>
        </header>
    );
};

export default Header;
