import React from 'react'
import FolderItem from './FolderItem';

const FolderList = ({folders}) => {
    return (
        <div className='folders btn-toolbar container' role="toolbar">
            {folders.map(folder => 
                <FolderItem folder={folder} key={folder.id} />
            )}
        </div>
    );
};

export default FolderList;
