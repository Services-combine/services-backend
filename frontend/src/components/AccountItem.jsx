import React from 'react'
import Button from './UI/button/Button';
import {Link, useNavigate, useParams} from 'react-router-dom'

const AccountItem = (props) => {
    const navigate = useNavigate()
    const params = useParams();
    
    return (
        <div className='alert alert-secondary'>
            <Link to={`/inviting/${params.folderID}/${props.account.id}`} className='open-account'>
                {props.index+1}. {props.account.name} (+{props.account.phone})
            </Link>

            <div className="actions">
                <button className='btn btn-danger btn-delete'>
                    <i className="fas fa-trash-alt"></i>
                </button>
            </div>
        </div>
    );
};

export default AccountItem;
