import React, {useState} from 'react'
import '../styles/Inviting.css';
import Button from './UI/button/Button';
import Select from './UI/select/Select';

const ModalFormSelect = ({create, foldersMove, defaultName}) => {
    const [path, setPath] = useState('');
    //console.log(foldersMove)

    const addInputSelect = (e) => {
		e.preventDefault()

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
                defaultName={defaultName}
                options={[
                    {value: "/", name: "/"},
                    {value: "Value 2", name: "Name 2"}
                ]}
                value={path} 
                onChange={folder => setPath(folder)}
            />
            <Button onClick={addInputSelect}>Сохранить</Button>
        </form>
	);
}

export default ModalFormSelect;
