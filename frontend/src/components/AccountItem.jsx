import React, {useState} from 'react'
import {Link, useParams} from 'react-router-dom'
import Modal from '../components/UI/modal/Modal';
import ModalConfirmAction from './ModalConfirmAction';

const AccountItem = (props) => {
    const params = useParams();
    const [modalDelete, setModalDelete] = useState(false);

    const deleteAccount = () => {
        props.remove(props.account);
    }

    const getModalAction = (getAction) => {
        setModalDelete(false);
        if (getAction.action) {
            deleteAccount();
        }
    }
    
    return (
        <div>
            <div className='alert alert-secondary'>
                <Link to={`/inviting/${params.folderID}/${props.account.id}`} className='account__link'>
                    {props.index+1}. {props.account.name} (+{props.account.phone})
                </Link>

                <div className="actions">
                    <button className='btn btn-danger btn-delete' onClick={() => setModalDelete(true)}>
                        <i className="fas fa-trash-alt"></i>
                    </button>
                    {props.account.launch &&
                        <h6 className="status-launch">&bull;</h6>
                    }
                </div>
            </div>

            <Modal visible={modalDelete} setVisible={setModalDelete}>
                <ModalConfirmAction result={getModalAction}/>
            </Modal>
        </div>
    );
};

export default AccountItem;
