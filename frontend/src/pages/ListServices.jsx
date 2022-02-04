import React, {useContext, useEffect, useState} from 'react'
import {Link} from "react-router-dom"
import '../styles/Services.css';
import {Context} from "../index";
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Loader from '../components/UI/loader/Loader';
import Modal from '../components/UI/modal/Modal';
import ModalSettings from '../components/ModalSetings';
import Services from '../API/Services';

const ListServices = () => {
    const {store} = useContext(Context)
    const [modalSettings, setModalSettings] = useState(false);
    const [isError, setIsError] = useState(null);
	const [isLoading, setIsLoading] = useState(false);
    const timeout = 3000;

    useEffect(() => {
        fetchData();
    }, [])

    async function fetchData() {
        try {
			setIsLoading(true);
			//const response = await Services.fetchData(localStorage.getItem('userID'));
            //console.log(response);

			setIsLoading(false);
		} catch (e) {
			setIsError('Ошибка при получении данных');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
		}
    }

    async function saveSettings(settings) {
        try {
            await Services.saveSettings(localStorage.getItem('userID'), settings.countInviting, settings.countMailing);
        } catch (e) {
            setIsError('Ошибка при сохранении настроек');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }

    const getModalSettings = (settings) => {
        setModalSettings(false);
        saveSettings(settings);
    }

	return (
        <div className='services'>
            <div className='header'>
                <h3 className='logo'>Сервисы</h3>
                <div className='header__btns'>
                    <Button onClick={() => setModalSettings(true)}>
                        <i className="fas fa-cog"></i> Настройки
                    </Button>
                    <Button onClick={() => store.logout()}>
                        <i className="fas fa-sign-out-alt"></i> Выйти
                    </Button>
                </div>
            </div>

            <div className="services__list btn-toolbar" role="toolbar">
                <Link to="/inviting" className="services__list-item">
                    <h6 className="services__list-item__title">Инвайтинг & Рассылка</h6>
                </Link>
            </div>

            {isLoading &&
                <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
            }

            {store.isError &&
                <Error>{store.isError}</Error>
            }

            {isError &&
                <Error>{isError}</Error>
            }

            <Modal visible={modalSettings} setVisible={setModalSettings}>
                <ModalSettings save={getModalSettings} defaultCountInviting="5" defaultCountMailing="3"/>
            </Modal>
        </div>
	);
}

export default ListServices;
