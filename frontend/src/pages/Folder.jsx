import React, {useEffect, useState, useRef} from 'react'
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
import ModalFormInput from '../components/ModalFormInput';
import ModalFormSelect from '../components/ModalFormSelect';
import ModalFormTextarea from '../components/ModalFormTextarea';
import ModalFormCreateAccount from '../components/ModalFormCreateAccount';
import ModalLaunch from '../components/ModalLaunch';

const Folder = () => {
    const params = useParams();
	const navigate = useNavigate()
    const [accounts, setAccounts] = useState([]);
    const [folders, setFolders] = useState([]);
    const [countAccounts, setCountAccounts] = useState({});
    const [dataFolder, setDataFolder] = useState({});
    const [foldersMove, setFoldersMove] = useState({});
    const [foldersHash, setFoldersHash] = useState({});
	const [totalCount, setTotalCount] = useState(0);
	const [limit, setLimit] = useState(20);
	const [skip, setSkip] = useState(0);
    const [isError, setIsError] = useState(null);
	const [isLoading, setIsLoading] = useState(false);
    const [modalCreateFolder, setModalCreateFolder] = useState(false);
    const [modalRename, setModalRename] = useState(false);
	const [modalChat, setModalChat] = useState(false);
	const [modalMessage, setModalMessage] = useState(false);
	const [modalUsernames, setModalUsernames] = useState(false);
	const [modalGroups, setModalGroups] = useState(false);
	const [modalMove, setModalMove] = useState(false);
	const [modalCreateAccount, setModaleCreateAccount] = useState(false);
	const [modalLaunch, setModalLaunch] = useState(false);
    const timeout = 3000;
	const lastElement = useRef();
	const observer = useRef();
	const [reload, setReload] = useState(false);

	useEffect(() => {
		if (isLoading) return;
		if (observer.current) observer.current.disconnect();
		var cb = function(entries, observer) {
			if (entries[0].isIntersecting && skip + limit < totalCount) {
				setSkip(skip + limit);
			}
		};
		observer.current = new IntersectionObserver(cb);
		observer.current.observe(lastElement.current)
	}, [isLoading])

    useEffect(() => {
		if (!reload)
	        fetchDataFolder();
    }, [params.folderID, skip])


	async function reloadData() {
		try {
			setReload(true);
			setSkip(0);
			setAccounts([]);
			setIsLoading(true);
			const response = await InvitingService.fetchDataFolder(params.folderID, limit, 0);

			if (response.data.accounts != null)
				setAccounts(response.data.accounts);
			else
				setAccounts([]);

			distributionData(response);
			setReload(false);
		} catch (e) {
			setIsError('Ошибка при получении данных папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
		}
	}

	async function fetchDataFolder() {
		try {
			setIsLoading(true);
			const response = await InvitingService.fetchDataFolder(params.folderID, limit, skip);

			if (response.data.accounts != null)
				setAccounts([...accounts, ...response.data.accounts]);
			else
				setAccounts([]);
			
			distributionData(response);
		} catch (e) {
			setIsError('Ошибка при получении данных папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
		}
	}

	const distributionData = (response) => {
		try {
			if (response.data.folders != null)
				setFolders(response.data.folders);
			else
				setFolders([]);

			setDataFolder(response.data.folder);
			setCountAccounts(response.data.countAccounts);
			setTotalCount(response.data.countAccounts.all);
			setFoldersMove(response.data.foldersMove);
			setFoldersHash(response.data.pathHash);

			setIsLoading(false);
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
            reloadData();
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
            reloadData();
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
			dataFolder.chat = chatName;
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
			dataFolder.message = message;
		} catch (e) {
			setIsError('Ошибка при изменении сообщения');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function changeUsernames(usernames) {
		try {
			const noDupUsernames = new Set(usernames.split("\n"));
			await InvitingService.changeUsernames(params.folderID, [...noDupUsernames]);
			dataFolder.usernames = [...noDupUsernames];
		} catch (e) {
			setIsError('Ошибка при добавлении usernames');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function changeGroups(groups) {
		try {
			const noDupGroups = new Set(groups.split("\n"));
			await InvitingService.changeGroups(params.folderID, [...noDupGroups]);
			dataFolder.groups = [...noDupGroups];
		} catch (e) {
			setIsError('Ошибка при добавлении групп');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function moveFolder(path) {
		try {
			await InvitingService.moveFolder(params.folderID, path);
			reloadData();
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

	async function createAccount(name, phone) {
		try {
			await InvitingService.createAccount(params.folderID, name, phone);
			reloadData();
		} catch (e) {
			setIsError('Ошибка при создании аккаунта');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function deleteAccount(account) {
        try {
			await InvitingService.deleteAccount(params.folderID, account.id);
			reloadData();
		} catch (e) {
			setIsError('Ошибка при удалении аккаунта');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
    }

	async function geterateInterval() {
		try {
			await InvitingService.geterateInterval(params.folderID);
			reloadData();
		} catch (e) {
			setIsError('Ошибка при генерации интервалов');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function sendCodeParsing(accountID) {
		try {
			await InvitingService.sendCodeParsing(params.folderID, accountID);
		} catch (e) {
			setIsError('Ошибка при отправке кода');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function parsingApi(accountID, code) {
		try {
			await InvitingService.parsingApi(params.folderID, accountID, code);
			reloadData();
		} catch (e) {
			setIsError('Ошибка при парсинге API');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function sendCodeSession(accountID) {
		try {
			await InvitingService.sendCodeSession(params.folderID, accountID);
			reloadData();
		} catch (e) {
			setIsError('Ошибка при отправке кода');
			setTimeout(() => {
				setIsError(null)
			}, timeout)
		}
	}

	async function createSession(accountID, code) {
		try {
			await InvitingService.createSession(params.folderID, accountID, code);
			reloadData();
		} catch (e) {
			setIsError('Ошибка при создании .session файла');
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
		else if (getInput.mode === "changeUsernames") {
			setModalUsernames(false);
			changeUsernames(getInput.text);
		}
		else if (getInput.mode === "changeGroups") {
			setModalGroups(false);
			changeGroups(getInput.text);
		}
		else if (getInput.mode === "createAccount") {
			setModaleCreateAccount(false);
			createAccount(getInput.name, getInput.phone);
		}
	}

	const getModalSelect = (getSelect) => {
		setModalMove(false);
		if (getSelect.path !== "")
			moveFolder(getSelect.path);
	}

	const getModalLaunch = () => {
		setModalLaunch(false);
		fetchDataFolder();
	}

    return (
        <div>
            <div className='header'>
                <div className='path'>
                    <Link to='/inviting' className='path__item-link'>Главная</Link>

					{Object.keys(foldersHash).length !== 0 &&
						Object.entries(foldersHash).map(([key, item]) => (
							<div className='path__item' key={item.value}>
								<b>/</b>
								<Link to={`/inviting/${item.value}`} className='path__item-link'>
									{item.name}
								</Link>
							</div>
						))
					}
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
					<Button className="btn-action" onClick={() => setModalLaunch(true)}>
						<i className="fas fa-play"></i> Запустить
					</Button>
				}

                <Button className='btn-action' onClick={() => setModalMessage(true)}>
					<i className="fas fa-comment-dots"></i> Сообщение
				</Button>
                <Button className='btn-action' onClick={() => setModalGroups(true)}>
					<i className="fas fa-users"></i> Группы
				</Button>
                <Button className='btn-action' onClick={() => setModalUsernames(true)}>
					<i className="fas fa-file-signature"></i> Username
				</Button>
                <Button className='btn-action' onClick={() => setModalChat(true)}>
					<i className="fas fa-user-friends"></i> Чат
				</Button>

                <Button className='btn-action' onClick={() => setModalCreateFolder(true)}>
					<i className="fas fa-folder-plus"></i> Папка
				</Button>
                <Button className='btn-action' onClick={() => setModaleCreateAccount(true)}>
					<i className="fas fa-user-plus"></i> Аккаунт
				</Button>

                <Button className='btn-action' onClick={() => setModalMove(true)}>
					<i className="fas fa-angle-double-right"></i> Переместить
				</Button>
                <Button className='btn-action' onClick={() => setModalRename(true)}>
					<i className="fas fa-signature"></i> Переименовать
				</Button>
                <Button className='btn-action' onClick={geterateInterval}>
					<i className="fas fa-random"></i> Сгенерировать
				</Button>
				<Button className='btn-action' onClick={reloadData}>
					<i className="fas fa-redo-alt"></i>
				</Button>
            </div>

            {isError &&
                <Error>{isError}</Error>
            }

			<FolderList folders={folders} />
			{accounts.length
				? <AccountList remove={deleteAccount} accounts={accounts} sendCodeParsing={sendCodeParsing} parsingApi={parsingApi} sendCodeSession={sendCodeSession} createSession={createSession} />
				: <h4 className='notification'>У вас пока нет аккаунтов</h4>
			}
			<div ref={lastElement} style={{height: 20}}></div>

            {isLoading &&
                <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
            }

			<Modal visible={modalCreateFolder} setVisible={setModalCreateFolder}>
                <ModalFormInput create={getModalInput} title="Создание папки" buttonText="Создать" mode="createFolder"/>
            </Modal>

			{dataFolder.name
				?
				<Modal visible={modalRename} setVisible={setModalRename}>
					<ModalFormInput create={getModalInput} title="Переименование папки" buttonText="Сохранить" mode="renameFolder" defaultData={dataFolder.name}/>
				</Modal>
				:
				<Modal visible={modalRename} setVisible={setModalRename}>
					<ModalFormInput create={getModalInput} title="Переименование папки" buttonText="Сохранить" mode="renameFolder" defaultData=""/>
				</Modal>
			}

			{dataFolder.chat
				?
				<Modal visible={modalChat} setVisible={setModalChat}>
					<ModalFormInput create={getModalInput} title="Изменение чата" buttonText="Сохранить" mode="changeChat" defaultData={dataFolder.chat}/>
				</Modal>
				:
				<Modal visible={modalChat} setVisible={setModalChat}>
					<ModalFormInput create={getModalInput} title="Изменение чата" buttonText="Сохранить" mode="changeChat" defaultData=""/>
				</Modal>
			}

			{dataFolder.message
				?
				<Modal visible={modalMessage} setVisible={setModalMessage}>
					<ModalFormTextarea create={getModalInput} title="Изменение сообщения" buttonText="Сохранить" mode="changeMessage" placeholderText="Введите сообщение" defaultData={dataFolder.message}/>
				</Modal>
				:
				<Modal visible={modalMessage} setVisible={setModalMessage}>
					<ModalFormTextarea create={getModalInput} title="Изменение сообщения" buttonText="Сохранить" mode="changeMessage" placeholderText="Введите сообщение" defaultData=""/>
				</Modal>
			}

			{dataFolder.usernames
				?
				<Modal visible={modalUsernames} setVisible={setModalUsernames}>
                	<ModalFormTextarea create={getModalInput} title="Добавление usernames" buttonText="Сохранить" mode="changeUsernames" placeholderText="Введите пользователей" defaultData={dataFolder.usernames}/>
            	</Modal>
				:
				<Modal visible={modalUsernames} setVisible={setModalUsernames}>
                	<ModalFormTextarea create={getModalInput} title="Добавление usernames" buttonText="Сохранить" mode="changeUsernames" placeholderText="Введите пользователей" defaultData=""/>
            	</Modal>
			}

			{dataFolder.groups
				? 
				<Modal visible={modalGroups} setVisible={setModalGroups}>
					<ModalFormTextarea create={getModalInput} title="Добавление групп" buttonText="Сохранить" mode="changeGroups" placeholderText="Введите группы" defaultData={dataFolder.groups}/>
				</Modal>
				:
				<Modal visible={modalGroups} setVisible={setModalGroups}>
					<ModalFormTextarea create={getModalInput} title="Добавление групп" buttonText="Сохранить" mode="changeGroups" placeholderText="Введите группы" defaultData=""/>
				</Modal>
			}

			{dataFolder.name_path &&
				Object.keys(foldersMove).length !== 0
					?
					<Modal visible={modalMove} setVisible={setModalMove}>
						<ModalFormSelect create={getModalSelect} optionsData={foldersMove} defaultName={dataFolder.name_path}/>
					</Modal>
					:
					<Modal visible={modalMove} setVisible={setModalMove}>
						<ModalFormSelect create={getModalSelect} optionsData={[]} defaultName={dataFolder.name_path}/>
					</Modal>
			}

			<Modal visible={modalCreateAccount} setVisible={setModaleCreateAccount}>
                <ModalFormCreateAccount create={getModalInput} mode="createAccount"/>
            </Modal>

			<Modal visible={modalLaunch} setVisible={setModalLaunch}>
                <ModalLaunch create={getModalLaunch}/>
            </Modal>
        </div>
	);
}

export default Folder;
