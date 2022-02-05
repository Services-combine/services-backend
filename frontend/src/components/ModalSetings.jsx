import React, {useEffect, useState} from 'react'
import '../styles/Folder.css';
import Input from './UI/input/Input';
import Button from './UI/button/Button';

const ModalSettings = ({save, defaultCountInviting, defaultCountMailing}) => {
    const [countInviting, setCountInviting] = useState('');
    const [countMailing, setCountMailing] = useState('');

    useEffect(() => {
        setCountInviting(defaultCountInviting);
        setCountMailing(defaultCountMailing);
    }, [defaultCountInviting, defaultCountMailing])

    const saveSettings = (e) => {
		e.preventDefault()

		const newSettings = {
            countInviting, countMailing, id: Date.now()
        }
        save(newSettings);
	}

    return (
        <div>
            <h5>Настройки</h5>

            <h6 className="title">Пользователей на инвайтинг</h6>
            <Input 
                value={countInviting} 
                onChange={e => setCountInviting(e.target.value)}
                type='number' 
                placeholder='Введите значение' 
            />
            <br/>
            <h6 className="title">Пользователей на рассылку</h6>
            <Input 
                value={countMailing} 
                onChange={e => setCountMailing(e.target.value)}
                type='number' 
                placeholder='Введите значение' 
            />
            <br/>

            <Button onClick={saveSettings}>Сохранить</Button>
        </div>
	);
}

export default ModalSettings;
