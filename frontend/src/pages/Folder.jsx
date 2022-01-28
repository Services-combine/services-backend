import React, {useEffect, useState} from 'react'
import '../styles/Folder.css';
import {Link} from "react-router-dom"
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';

const Folder = () => {
    const [isError, setIsError] = useState(null);
    const timeout = 3000;

    useEffect(() => {
        
    }, [])

    async function getDataFolder() {
        try {
            
        } catch (e) {
            setIsError('Ошибка при получении данных папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }

    return (
        <div>
            <div className='header'>
                <h3 className='logo'>Folder</h3>
                <Button><i className="fas fa-plus"></i> Создать папку</Button>
            </div>

            <div className='menu'>
                <Button className='btn-action'><i className="fas fa-comment-dots"></i> Сообщение</Button>
                <Button className='btn-action'><i className="fas fa-users"></i> Группы</Button>
                <Button className='btn-action'><i className="fas fa-file-signature"></i> Username</Button>
                <Button className='btn-action'><i className="fas fa-user-friends"></i> Чат</Button>
            </div>

            {isError &&
                <Error>{isError}</Error>
            }

            <div className='folders btn-toolbar' role="toolbar">
                
            </div>
        </div>
	);
}

export default Folder;
