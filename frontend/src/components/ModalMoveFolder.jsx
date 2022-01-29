import React, {useState} from 'react'
import '../styles/Inviting.css';
import Button from './UI/button/Button';
import Select from './UI/select/Select';

const ModalForm = ({create, foldersMove, defaultPath}) => {
    const [path, setPath] = useState('');

    const addInputSelect = (e) => {
		e.preventDefault()
        console.log(defaultPath)

		const newSelect = {
            path, id: Date.now()
        }
        create(newSelect)
		setPath('')
	}

    return (
        <form>
            <h5>Перемещение папки</h5>
            <Select
                defaultValue="Name 0"
                options={[
                    {value: "Value 1", name: "Name 1"},
                    {value: "Value 2", name: "Name 2"}
                ]}
                value={path} 
                onChange={folder => setPath(folder)}
            />
            <Button onClick={addInputSelect}>Сохранить</Button>
        </form>
	);
}

export default ModalForm;
