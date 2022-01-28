import React, {useState} from 'react'
import '../styles/Inviting.css';
import Input from '../components/UI/input/Input';
import Button from '../components/UI/button/Button';

const CreateFolderForm = ({create}) => {
    const [folderName, setFolderName] = useState('');

    const addNewFolder = (e) => {
		e.preventDefault()

		const newFolder = {
            folderName, id: Date.now()
        }
        create(newFolder)
		setFolderName('')
	}

    return (
        <form>
            <h5>Создание папки</h5>
            <Input 
                value={folderName} 
                onChange={e => setFolderName(e.target.value)}
                type='text' 
                placeholder='Введите название' 
            />
            <Button onClick={addNewFolder}>Создать папку</Button>
        </form>
	);
}

export default CreateFolderForm;
