import React, {useState} from 'react'
import {Link, useParams} from 'react-router-dom'
import Modal from '../components/UI/modal/Modal';
import ModalConfirmAction from './ModalConfirmAction';
import Button from './UI/button/Button';
import Input from './UI/input/Input';

const AccountItem = (props) => {
    const params = useParams();
    const [modalDelete, setModalDelete] = useState(false);
    const [code, setCode] = useState("");

    const deleteAccount = () => {
        props.remove(props.account);
    }

    const sendCodeParsing = () => {
        props.sendCodeParsing(props.account.id);
    }

    const parsingApi = () => {
        props.parsingApi(props.account.id, code);
        setCode('');
    }

    const sendCodeSession = () => {
        props.sendCodeSession(props.account.id);
    }

    const createSession = () => {
        props.createSession(props.account.id, code);
        setCode('');
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
                    {props.account.api_id === 0
                        ?
                        <div className='actions__parsing'>
                            <Button style={{backgroundColor: "rgb(165, 165, 165)"}} onClick={sendCodeParsing}>
                                <i className="fas fa-paper-plane"></i> Отпправить код
                            </Button>
                            <Input 
                                value={code}
                                onChange={e => setCode(e.target.value)}
                                type='text' 
                                placeholder='Код'
                            />
                            <Button style={{backgroundColor: "rgb(81, 211, 113)"}} onClick={parsingApi}>
                                <i className="fas fa-user-check"></i> Спарсить api
                            </Button>
                        </div>
                        :
                        props.account.verify !== true
                            ?
                            <div className='actions__create'>
                                <Button style={{backgroundColor: "rgb(165, 165, 165)"}} onClick={sendCodeSession}>
                                    <i className="fas fa-paper-plane"></i> Отпправить код
                                </Button>
                                <Input 
                                    value={code}
                                    onChange={e => setCode(e.target.value)}
                                    type='text' 
                                    placeholder='Код'
                                />
                                <Button style={{backgroundColor: "rgb(81, 211, 113)"}} onClick={createSession}>
                                    <i className="fas fa-user-check"></i> Создать файл
                                </Button>
                            </div>
                            :
                            <div className='actions__status'>
                                {props.account.status_block === "clean"
                                    ? <h6 className="status-block no-block"><i className="fas fa-check"></i></h6>
                                    : <h6 className="status-block info-block" unblocking={props.account.status_block}><i className="fas fa-info-circle"></i></h6>
                                }

                                {props.account.launch &&
                                    props.account.status_block === 'clean' &&
                                        <h6 className="status-launch">&bull;</h6>
                                }
                            </div>
                    }

                    <button className='btn btn-danger btn-delete' onClick={deleteAccount}>
                        <i className="fas fa-trash-alt"></i>
                    </button>
                </div>
            </div>

            <Modal visible={modalDelete} setVisible={setModalDelete}>
                <ModalConfirmAction result={getModalAction}/>
            </Modal>
        </div>
    );
};

export default AccountItem;
