import React, {useEffect, useState} from 'react'
import '../styles/Inviting.css';
import {Link} from "react-router-dom"
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Modal from '../components/UI/modal/Modal';
import CreateFolderForm from '../components/CreateFolderForm';

const Inviting = () => {
    const [folders, setFolders] = useState([]);
    const [isError, setIsError] = useState(null);
    const [modal, setModal] = useState(false);
    const timeout = 3000;

    useEffect(() => {
        getFolders();
    }, [])

    async function getFolders() {
        try {
            const response = await InvitingService.fetchMainFolders();
            setFolders(response.data);
        } catch (e) {
            setIsError('Ошибка при получении папок');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }
    
    async function addFolderDB(folderName) {
        try {
            const response = await InvitingService.createFolder(folderName);
            getFolders();
        } catch (e) {
            setIsError('Ошибка при создании папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }

    const createFolder = (newFolder) => {
        addFolderDB(newFolder.folderName);
		setModal(false);
	}

    return (
        <div>
            <div className='header'>
                <h3 className='logo'>Инвайтинг & Рассылка</h3>
                <Button onClick={() => setModal(true)}><i className="fas fa-plus"></i> Создать папку</Button>
            </div>

            {isError &&
                <Error>{isError}</Error>
            }

            <Modal visible={modal} setVisible={setModal}>
                <CreateFolderForm create={createFolder}/>
            </Modal>

            <div className='folders btn-toolbar' role="toolbar">
                {folders.map(folder => 
                    <Link to={'/inviting/' + folder.id} key={folder.id} className='folder'>
                        <i className="fas fa-folder-open folder__icon"></i>
                        <h6 className="folder__name">{folder.name}</h6>
                    </Link>
                )}
            </div>
        </div>
	);
}

export default Inviting;
