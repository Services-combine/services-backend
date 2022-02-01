import React, {useState} from 'react'
import '../styles/Folder.css';
import {useParams} from "react-router-dom"
import Button from './UI/button/Button';
import InvitingService from '../API/InvitingService';
import Error from './UI/error/Error';

const ModalLaunch = ({create}) => {
    const params = useParams();
    const [isError, setIsError] = useState(null);
    const timeout = 3000;

    const addLaunchAction = () => {
        create()
	}

    async function launchInviting(e) {
        e.preventDefault()
        try {
			const response = await InvitingService.launchInviting(params.folderID);
            addLaunchAction();
		} catch (e) {
			setIsError(e.response?.data?.message);
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
    }

    async function launchMailingUsernames(e) {
        e.preventDefault()
        try {
			const response = await InvitingService.launchMailingUsernames(params.folderID);
            addLaunchAction();
		} catch (e) {
			setIsError(e.response?.data?.message);
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
    }

    async function launchMailingGroups(e) {
        e.preventDefault()
        try {
			const response = await InvitingService.launchMailingGroups(params.folderID);
            addLaunchAction();
		} catch (e) {
			setIsError(e.response?.data?.message);
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
    }

    return (
        <form className="launch__btns btn-toolbar" role="toolbar">
            <h5>Выберите действие</h5>
            <Button><i className="fas fa-user-lock"></i> Блокировка</Button>
            <Button onClick={launchInviting}><i className="fas fa-play"></i> Инвайтинг</Button>
            <Button onClick={launchMailingUsernames}><i className="fas fa-play"></i> Рассылка пользователям</Button>
            <Button onClick={launchMailingGroups}><i className="fas fa-play"></i> Рассылка в группы</Button>

            {isError &&
                <Error style={{width: '100%'}}>{isError}</Error>
            }
        </form>
	);
}

export default ModalLaunch;
