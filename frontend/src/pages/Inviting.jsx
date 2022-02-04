import React, {useEffect, useState} from 'react'
import '../styles/Inviting.css';
import { useNavigate } from 'react-router-dom';
import InvitingService from '../API/InvitingService';
import Error from '../components/UI/error/Error';
import Button from '../components/UI/button/Button';
import Modal from '../components/UI/modal/Modal';
import Loader from '../components/UI/loader/Loader';
import ModalFormInput from '../components/ModalFormInput';
import FolderList from '../components/FolderList';
import ModalParams from '../components/ModalParams';

const Inviting = () => {
    let navigate = useNavigate();
    const [folders, setFolders] = useState([]);
    const [countAccounts, setCountAccounts] = useState({});
    const [isError, setIsError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    const [modalCreateFolder, setModalCreateFolder] = useState(false);
    const [modalParams, setModalParams] = useState(false);
    const timeout = 3000;

    useEffect(() => {
        fetchFolders();
    }, [])

    async function fetchFolders() {
        try {
            setIsLoading(true);
            const response = await InvitingService.fetchFolders();

            if (response.data !== null) {
                setFolders(response.data.folders);
                setCountAccounts(response.data.countAccounts);
            }
            
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
			setModalCreateFolder(false);
			createFolder(getData.text);
		}
	}

    return (
        <div>
            <div className='header'>
                <h3 className='logo'>Инвайтинг & Рассылка</h3>
                <div className='header__btns'>
                    <Button onClick={() => navigate("/")}><i className="fas fa-home"></i> На главную</Button>
                    <Button onClick={() => setModalParams(true)}><i className="fas fa-chart-pie"></i> Показатели</Button>
                    <Button onClick={() => setModalCreateFolder(true)}><i className="fas fa-plus"></i> Создать папку</Button>
                </div>
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

            <Modal visible={modalCreateFolder} setVisible={setModalCreateFolder}>
                <ModalFormInput create={getModalData} title="Создание папки" buttonText="Создать" mode="createFolder"/>
            </Modal>

            <Modal visible={modalParams} setVisible={setModalParams}>
                <ModalParams props={countAccounts}/>
            </Modal>
        </div>
	);
}

export default Inviting;
