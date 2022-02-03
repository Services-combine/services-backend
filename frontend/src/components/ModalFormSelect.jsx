import React, {useEffect, useState} from 'react'
import '../styles/Inviting.css';
import Button from './UI/button/Button';
import Select from './UI/select/Select';

const ModalFormSelect = ({create, optionsData, defaultName}) => {
    const [path, setPath] = useState('');
    const [listOptions, setListOptions] = useState([]);

    useEffect(() => {
        if (Object.keys(optionsData).length !== 0) {
            //console.log(optionsData);
            //console.log(listOptions);
            for (var option in optionsData) {
                //console.log(option, optionsData[option]);
                setListOptions([...listOptions, {"value": optionsData[option], "name": option}])
            }
            //console.log(listOptions, optionsData);
        }
    }, [optionsData])

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
                defaultName="Выберите папку"
                options={listOptions}
                value={path} 
                onChange={folder => setPath(folder)}
            />
            <Button onClick={addInputSelect}>Сохранить</Button>
        </form>
	);
}

export default ModalFormSelect;
