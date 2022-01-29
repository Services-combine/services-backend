import React, {useEffect, useState} from 'react'
import '../styles/Folder.css';
import {Link, useParams, useNavigate} from "react-router-dom"
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Loader from '../components/UI/loader/Loader';
import Modal from '../components/UI/modal/Modal';
import CountAccounts from '../components/CountAccounts';
import FolderList from '../components/FolderList';
import AccountList from '../components/AccountList';
import ModalForm from '../components/ModalForm';
import ModalMoveFolder from '../components/ModalMoveFolder';

const Folder = () => {
    const params = useParams();
	const navigate = useNavigate()
    const [accounts, setAccounts] = useState([]);
    const [folders, setFolders] = useState([]);
    const [countAccounts, setCountAccounts] = useState({});
    const [dataFolder, setDataFolder] = useState({});
    const [foldersMove, setFoldersMove] = useState({});
    const [foldersHash, setFoldersHash] = useState({});
    const [isError, setIsError] = useState(null);
	const [isLoading, setIsLoading] = useState(false);
    const [modalCreateFolder, setModalCreateFolder] = useState(false);
    const [modalRename, setModalRename] = useState(false);
	const [modalChat, setModalChat] = useState(false);
	const [modalMessage, setModalMessage] = useState(false);
	const [modalMove, setModalMove] = useState(false);
    const timeout = 3000;

    useEffect(() => {
        fetchDataFolder();
    }, [params.folderID])

	async function fetchDataFolder() {
		try {
			setIsLoading(true);
			const response = await InvitingService.fetchDataFolder(params.folderID);
			//console.log(response.data.pathHash)

			if (response.data.folders != null)
				setFolders(response.data.folders);
			else
				setFolders([]);

			if (response.data.accounts != null)
				setAccounts(response.data.accounts);
			else
				setAccounts([]);

			setIsLoading(false);

			setCountAccounts(response.data.countAccounts);
			setDataFolder(response.data.folder);
			setFoldersMove(response.data.foldersMove);
			setFoldersHash(response.data.foldersHash);
		} catch (e) {
			setIsError('Ошибка при получении данных папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
		}
	}

    async function createFolder(folderName) {
        try {
            await InvitingService.createFolderInFolder(params.folderID, folderName);
            fetchDataFolder();
        } catch (e) {
            setIsError('Ошибка при создании папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }

    async function renameFolder(folderName) {
        try {
            await InvitingService.renameFolder(params.folderID, folderName);
            fetchDataFolder();
        } catch (e) {
            setIsError('Ошибка при переименовывании папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }

	async function changeChat(chatName) {
		try {
			await InvitingService.changeChat(params.folderID, chatName);
		} catch (e) {
			setIsError('Ошибка при изменении чата');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function changeMessage(message) {
		try {
			await InvitingService.addMessage(params.folderID, message);
		} catch (e) {
			setIsError('Ошибка при изменении сообщения');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function moveFolder(path) {
		try {
			await InvitingService.moveFolder(params.folderID, path);
		} catch (e) {
			setIsError('Ошибка при перемещении папки');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function deleteFolder() {
		try {
			const response = await InvitingService.deleteFolder(params.folderID);

			if (response.data === "/")
				navigate('/inviting')
			else 
				navigate(`/inviting/${response.data}`)
		} catch (e) {
			setIsError('Ошибка при удалении папки');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	const getModalInput = (getInput) => {
		if (getInput.mode === "createFolder") {
			setModalCreateFolder(false);
			createFolder(getInput.text);
		} 
		else if (getInput.mode === "renameFolder") {
			setModalRename(false);
			renameFolder(getInput.text);
		}
        else if (getInput.mode === "changeChat") {
			setModalChat(false);
			changeChat(getInput.text);
		}
		else if (getInput.mode === "changeMessage") {
			setModalMessage(false);
			changeMessage(getInput.text);
		}
	}

	const getModalSelect = (getSelect) => {
		setModalMove(false);
		//moveFolder(getSelect.path);
		console.log(getSelect);
	}

    return (
        <div>
            <div className='header'>
                <div className='path'>
                    <Link to='/inviting' className='path__item'>Главная</Link>
                </div>
				<div className='header__btns'>
					<CountAccounts all={countAccounts.all} clean={countAccounts.clean} block={countAccounts.block} />
				</div>
            </div>

            <div className='menu btn-toolbar' role="toolbar">
				{accounts.length === 0 &&
					folders.length === 0 && 
						<Button style={{background: "rgb(233, 62, 62)", color: "#dedede"}} className="delete" onClick={deleteFolder}>
							<i className="fas fa-trash-alt"></i>
						</Button>
				}

				{accounts.length !== 0 &&
					<Button className="btn-action">
						<i className="fas fa-play"></i> Запустить
					</Button>
				}

                <Button className='btn-action' onClick={() => setModalMessage(true)}><i className="fas fa-comment-dots"></i> Сообщение</Button>
                <Button className='btn-action'><i className="fas fa-users"></i> Группы</Button>
                <Button className='btn-action'><i className="fas fa-file-signature"></i> Username</Button>
                <Button className='btn-action' onClick={() => setModalChat(true)}><i className="fas fa-user-friends"></i> Чат</Button>

                <Button className='btn-action' onClick={() => setModalCreateFolder(true)}><i className="fas fa-folder-plus"></i> Папка</Button>
                <Button className='btn-action'><i className="fas fa-user-plus"></i> Аккаунт</Button>

                <Button className='btn-action' onClick={() => setModalMove(true)}><i className="fas fa-angle-double-right"></i> Переместить</Button>
                <Button className='btn-action' onClick={() => setModalRename(true)}><i className="fas fa-signature"></i> Переименовать</Button>
                <Button className='btn-action'><i className="fas fa-random"></i> Сгенерировать</Button>
            </div>

            {isError &&
                <Error>{isError}</Error>
            }

            {isLoading
                ? <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
                :
                <>
					<FolderList folders={folders} />
					{accounts.length
						? <AccountList accounts={accounts} />
						: <h4 className='notification'>У вас пока нет аккаунтов</h4>
					}
                </>
            }

			<Modal visible={modalCreateFolder} setVisible={setModalCreateFolder}>
                <ModalForm create={getModalInput} title="Создание папки" buttonText="Создать" mode="createFolder"/>
            </Modal>

			<Modal visible={modalRename} setVisible={setModalRename}>
                <ModalForm create={getModalInput} title="Переименование папки" buttonText="Сохранить" mode="renameFolder"/>
            </Modal>

			<Modal visible={modalChat} setVisible={setModalChat}>
                <ModalForm create={getModalInput} title="Изменение чата" buttonText="Сохранить" mode="changeChat"/>
            </Modal>

			<Modal visible={modalMessage} setVisible={setModalMessage}>
                <ModalForm create={getModalInput} title="Изменение сообщения" buttonText="Сохранить" mode="changeMessage"/>
            </Modal>

			<Modal visible={modalMove} setVisible={setModalMove}>
                <ModalMoveFolder create={getModalSelect} foldersMove={foldersMove} defaultPath={dataFolder.path}/>
            </Modal>
        </div>
	);
}

export default Folder;
