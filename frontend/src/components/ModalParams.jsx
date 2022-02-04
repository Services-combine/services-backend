import React from 'react'
import '../styles/Folder.css';

const ModalParams = ({props}) => {

    return (
        <div>
            <h5>Показатели</h5>
            <div className='params'>
                <h6>Всего аккаунтов - {props.all}</h6>
                <h6>Чистых аккаунтов - {props.clean}</h6>
                <h6>Заблокированных аккаунтов - {props.block}</h6>
            </div>
        </div>
	);
}

export default ModalParams;
