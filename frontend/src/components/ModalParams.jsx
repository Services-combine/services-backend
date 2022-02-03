import React from 'react'
import '../styles/Folder.css';

const ModalParams = ({params}) => {

    return (
        <div>
            <h5>Показатели</h5>
            <div className='params'>
                <h6>Всего аккаунтов - {params.all}</h6>
                <h6>Чистых аккаунтов - {params.clean}</h6>
                <h6>Заблокированных аккаунтов - {params.block}</h6>
            </div>
        </div>
	);
}

export default ModalParams;
