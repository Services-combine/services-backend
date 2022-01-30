import React from 'react'
import {Link, useParams} from 'react-router-dom'

const AccountItem = (props) => {
    const params = useParams();
    
    return (
        <div className='alert alert-secondary'>
            <Link to={`/inviting/${params.folderID}/${props.account.id}`} className='account__link'>
                {props.index+1}. {props.account.name} (+{props.account.phone})
            </Link>

            <div className="actions">
                <button className='btn btn-danger btn-delete' onClick={() => props.remove(props.account)}>
                    <i className="fas fa-trash-alt"></i>
                </button>
            </div>
        </div>
    );
};

export default AccountItem;
