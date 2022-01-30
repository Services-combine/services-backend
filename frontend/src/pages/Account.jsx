import React, {useEffect, useState} from 'react'
import '../styles/Account.css';
import {Link, useParams, useNavigate} from "react-router-dom"
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Loader from '../components/UI/loader/Loader';
import Input from '../components/UI/input/Input';

const Folder = () => {
    const params = useParams();
    const navigate = useNavigate()
    const [account, setAccount] = useState({});
    const [name, setName] = useState('');
    const [interval, setInterval] = useState(0);
    const [folder, setFolder] = useState('');
    const [isError, setIsError] = useState(null);
	const [isLoading, setIsLoading] = useState(false);
    const timeout = 3000;

    useEffect(() => {
        fetchDataAccount();
    }, [])

	async function fetchDataAccount() {
		try {
			setIsLoading(true);
			const response = await InvitingService.fetchDataAccount(params.folderID, params.accountID);
            setAccount(response.data);
            setName(account.name);
            setInterval(account.interval);
            setFolder(account.folder_id);

			setIsLoading(false);
		} catch (e) {
			setIsError('Ошибка при получении данных аккаунта');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
		}
	}

    const randomInterval = (e) => {
        e.preventDefault()
        const min = 15;
        const max = 40;
        const rand = Math.floor(min + Math.random() * (max - min));
        setInterval(rand);
    }

    async function saveSettings() {
        try {
			await InvitingService.saveSettingsAccount(params.folderID, params.accountID, name, interval, folder);
		} catch (e) {
			setIsError('Ошибка при сохранении настроек');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
    }

    return (
        <div>
            <div className='header'>
                <div className='header__left'>
                    <Link to={`/inviting/${params.folderID}`} className='again'>
                        <i className="fas fa-arrow-left"></i>
                    </Link>
                    <h3 className='title'>Настройки аккаунта</h3>
                </div>
				<div className='header__btns'>
					
				</div>
            </div>

            {isError &&
                <Error>{isError}</Error>
            }

            {isLoading
                ? <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
                :
                <>
					<form className='settings-accounts'>
                        <h6 className="title">Название аккаунта</h6>
                        <Input 
                            value={name}
                            onChange={e => setName(e.target.value)}
                            type='text' 
                            placeholder='Введите название' 
                        />
                        
                        <h6 className="title">Номер телефона - +{account.phone}</h6>
                        <h6 className="title">Чат - {account.chat}</h6>

                        <h6 className="title">Интервал</h6>
                        <div className="interval">
                            <Input 
                                value={interval}
                                onChange={e => setInterval(e.target.value)}
                                type='text' 
                                placeholder='Введите интервал' 
                            />
                            <Button className='generate' onClick={randomInterval}>
                                <i className="fas fa-random"></i>
                            </Button>
                        </div>

                        <h6 className="title">Папка</h6>

                        <Button onClick={saveSettings}>
                            Сохранить
                        </Button>
                    </form>
                </>
            }
        </div>
	);
}

export default Folder;
