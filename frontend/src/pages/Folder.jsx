import React, {useEffect, useState} from 'react'
import '../styles/Folder.css';
import {Link, useParams} from "react-router-dom"
import { useFetching } from '../hooks/useFetching';
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Loader from '../components/UI/loader/Loader';
import Modal from '../components/UI/modal/Modal';
import CreateFolderForm from '../components/CreateFolderForm';
import RenameFolderForm from '../components/RenameFolderForm';

const Folder = () => {
    const params = useParams();
    const [accounts, setAccounts] = useState([]);
    const [folders, setFolders] = useState([]);
    const [countAccounts, setCountAccounts] = useState({});
    const [dataFolder, setDataFolder] = useState({});
    const [foldersMove, setFoldersMove] = useState({});
    const [foldersHash, setFoldersHash] = useState({});
    const [isError, setIsError] = useState(null);
    const [modalCreateFolder, setModalCreateFolder] = useState(false);
    const [modalRename, setModalRename] = useState(false);
    const timeout = 3000;

    const [fetchDataFolder, isFetchLoading, fetchError] = useFetching(async () => {
        const response = await InvitingService.fetchDataFolder(params.folderID)

        if (response.data.accounts != null)
            setAccounts(response.data.accounts);
        if (response.data.folders != null)
            setFolders(response.data.folders);

        setCountAccounts(response.data.countAccounts);
        setDataFolder(response.data.folder);
        setFoldersMove(response.data.foldersMove);
        setFoldersHash(response.data.foldersHash);
    })

    useEffect(() => {
        fetchDataFolder();
    }, [])

    const deleteAccount = (accountID) => {
        console.log(accountID);
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

    const getCreateFolder = (newFolder) => {
        setModalCreateFolder(false);
        createFolder(newFolder.folderName);
	}

    const getRenameFolder = (newName) => {
        setModalRename(false);
        renameFolder(newName.folderName);
	}

    return (
        <div>
            <div className='header'>
                <div className='path'>
                    <Link to='/inviting' className='path__item'>Главная</Link>
                </div>
                <div className="count-accounts">
                    <h6><i className="fas fa-user-alt"></i> - {countAccounts.all}</h6>
                    <h6><i className="fas fa-check"></i> - {countAccounts.clean}</h6>
                    <h6><i className="fas fa-info-circle"></i> - {countAccounts.block}</h6>
                </div>
            </div>

            <div className='menu btn-toolbar' role="toolbar">
                <Button className='btn-action'><i className="fas fa-comment-dots"></i> Сообщение</Button>
                <Button className='btn-action'><i className="fas fa-users"></i> Группы</Button>
                <Button className='btn-action'><i className="fas fa-file-signature"></i> Username</Button>
                <Button className='btn-action'><i className="fas fa-user-friends"></i> Чат</Button>

                <Button className='btn-action' onClick={() => setModalCreateFolder(true)}><i className="fas fa-folder-plus"></i> Папка</Button>
                <Button className='btn-action'><i className="fas fa-user-plus"></i> Аккаунт</Button>

                <Button className='btn-action'><i className="fas fa-angle-double-right"></i> Переместить</Button>
                <Button className='btn-action' onClick={() => setModalRename(true)}><i className="fas fa-signature"></i> Переименовать</Button>
                <Button className='btn-action'><i className="fas fa-random"></i> Сгенерировать</Button>
            </div>

            <Modal visible={modalCreateFolder} setVisible={setModalCreateFolder}>
                <CreateFolderForm create={getCreateFolder}/>
            </Modal>

            <Modal visible={modalRename} setVisible={setModalRename}>
                <RenameFolderForm create={getRenameFolder}/>
            </Modal>

            {isError &&
                <Error>{isError}</Error>
            }

            {fetchError &&
                <Error>Ошибка при получении данных папки</Error>
            }

            {isFetchLoading
                ? <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
                :
                <>
                    <div className='folders btn-toolbar' role="toolbar">
                        {folders.map(folder => 
                            <Link to={`/inviting/${folder.id}`} key={folder.id} className='folder'>
                                <i className="fas fa-folder-open folder__icon"></i>
                                <h6 className="folder__name">{folder.name}</h6>
                            </Link>
                        )}
                    </div>

                    <div className='accounts'>
                        {accounts.map(account => 
                            <div key={account.id} className='alert alert-secondary'>
                                <Link to={`/inviting/${params.folderID}/${account.id}`} className='open-account'>
                                    {account.name} (+{account.phone})
                                </Link>

                                <div className="actions">
                                    <button className='btn btn-danger btn-delete' onClick={deleteAccount(account.id)}>
                                        <i className="fas fa-trash-alt"></i>
                                    </button>
                                </div>
                            </div>
                        )}
                    </div>
                </>
            }
        </div>
	);
}

export default Folder;
