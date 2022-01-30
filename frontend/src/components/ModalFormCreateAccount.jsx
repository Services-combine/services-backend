import React, {useState} from 'react'
import '../styles/Inviting.css';
import Input from './UI/input/Input';
import Button from './UI/button/Button';

const ModalFormCreateAccount = ({create, mode}) => {
    const [name, setName] = useState('');
    const [phone, setPhone] = useState('');

    const addInputText = (e) => {
		e.preventDefault()

		const newInput = {
            name, phone, id: Date.now(), mode
        }
        create(newInput)
		setName('')
        setPhone('')
	}

    return (
        <form>
            <h5>Создание аккаунта</h5>
            <Input 
                value={name} 
                onChange={e => setName(e.target.value)}
                type='text' 
                placeholder='Введите название' 
            />
            <Input 
                value={phone} 
                onChange={e => setPhone(e.target.value)}
                type='text' 
                placeholder='Введите номер телефона' 
            />
            <Button onClick={addInputText}>Создать</Button>
        </form>
	);
}

export default ModalFormCreateAccount;
