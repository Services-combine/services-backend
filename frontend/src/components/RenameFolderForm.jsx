import React, {useState} from 'react'
import '../styles/Inviting.css';
import Input from '../components/UI/input/Input';
import Button from '../components/UI/button/Button';

const RenameFolderForm = ({create}) => {
    const [folderName, setFolderName] = useState('');

    const renameFolder = (e) => {
		e.preventDefault()

		const newFolderName = {
            folderName, id: Date.now()
        }
        create(newFolderName)
		setFolderName('')
	}

    return (
        <form>
            <h5>Переименовывание папки</h5>
            <Input 
                value={folderName} 
                onChange={e => setFolderName(e.target.value)}
                type='text' 
                placeholder='Введите название' 
            />
            <Button onClick={renameFolder}>Сохранить</Button>
        </form>
	);
}

export default RenameFolderForm;
