import React, {useEffect, useState} from 'react'
import '../styles/Inviting.css';
import {Link} from "react-router-dom"
import { useFetching } from '../hooks/useFetching';
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Modal from '../components/UI/modal/Modal';
import Loader from '../components/UI/loader/Loader';
import CreateFolderForm from '../components/CreateFolderForm';

const Inviting = () => {
    const [folders, setFolders] = useState([]);
    const [isError, setIsError] = useState(null);
    const [modal, setModal] = useState(false);
    const timeout = 3000;

    const [fetchFolders, isFetchLoading, fetchError] = useFetching(async () => {
        const response = await InvitingService.fetchFolders();
        setFolders(response.data);
    })

    useEffect(() => {
        fetchFolders();
    }, [])
    
    async function createFolder(folderName) {
        try {
            await InvitingService.createFolder(folderName);
            fetchFolders();
        } catch (e) {
            setIsError('Ошибка при создании папки');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }

    const modalCreateFolder = (newFolder) => {
        setModal(false);
        createFolder(newFolder.folderName);
	}

    return (
        <div>
            <div className='header'>
                <h3 className='logo'>Инвайтинг & Рассылка</h3>
                <Button onClick={() => setModal(true)}><i className="fas fa-plus"></i> Создать папку</Button>
            </div>

            <Modal visible={modal} setVisible={setModal}>
                <CreateFolderForm create={modalCreateFolder}/>
            </Modal>

            {fetchError &&
                <Error>Ошибка при получении папок</Error>
            }

            {isError &&
                <Error>{isError}</Error>
            }

            {isFetchLoading 
                ? <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
                : 
                <div className='folders btn-toolbar' role="toolbar">
                    {folders.map(folder => 
                        <Link to={`/inviting/${folder.id}`} key={folder.id} className='folder'>
                            <i className="fas fa-folder-open folder__icon"></i>
                            <h6 className="folder__name">{folder.name}</h6>
                        </Link>
                    )}
                </div> 
            }
        </div>
	);
}

export default Inviting;
