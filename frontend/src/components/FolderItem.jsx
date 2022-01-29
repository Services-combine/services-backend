import React from 'react'
import {Link} from 'react-router-dom'

const FolderItem = (props) => {
    return (
        <div>
            <Link to={`/inviting/${props.folder.id}`} className='folder'>
                <i className="fas fa-folder-open folder__icon"></i>
                <h6 className="folder__name">{props.folder.name}</h6>
            </Link>
        </div>
    );
};

export default FolderItem;
