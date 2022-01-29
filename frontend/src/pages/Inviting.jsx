import React, {useEffect, useState} from 'react'
import '../styles/Inviting.css';
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Modal from '../components/UI/modal/Modal';
import Loader from '../components/UI/loader/Loader';
import ModalForm from '../components/ModalForm';
import FolderList from '../components/FolderList';

const Inviting = () => {
    const [folders, setFolders] = useState([]);
    const [isError, setIsError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    const [modal, setModal] = useState(false);
    const timeout = 3000;

    useEffect(() => {
        fetchFolders();
    }, [])

    async function fetchFolders() {
        try {
            setIsLoading(true);
            const response = await InvitingService.fetchFolders();
            setFolders(response.data);
            setIsLoading(false);
        } catch (e) {
            setIsError('Ошибка при получении папок');
            setTimeout(() => {
                setIsError(null)
            }, timeout)
        }
    }
    
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

    const getModalData = (getData) => {
		if (getData.mode === "createFolder") {
			setModal(false);
			createFolder(getData.text);
		}
		
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

            {isLoading 
                ? <div style={{display: "flex", justifyContent: "center", marginTop: 50}}><Loader/></div>
                :
                folders.length !== 0
                    ? <FolderList folders={folders} />
                    : <h4 className='notification'>У вас пока нет папок</h4>
            }

            <Modal visible={modal} setVisible={setModal}>
                <ModalForm create={getModalData} title="Создание папки" buttonText="Создать" mode="createFolder"/>
            </Modal>
        </div>
	);
}

export default Inviting;
