import React, {useState} from 'react'
import '../styles/Inviting.css';
import Button from './UI/button/Button';
import Textarea from './UI/textarea/Textarea';

const ModalFormTextarea = ({create, mode, title, buttonText, placeholderText, defaultData}) => {
    const [text, setText] = useState('');
    if (defaultData !== undefined && text === '' && defaultData !== '' && defaultData !== text) {
        console.log(defaultData);
        setText(defaultData);
    }

    const addTextareaText = (e) => {
		e.preventDefault()

		const newTextarea = {
            text, id: Date.now(), mode
        }
        create(newTextarea)
		setText('')
	}

    return (
        <form>
            <h5>{title}</h5>
            <Textarea 
                value={text} 
                onChange={e => setText(e.target.value)}
                placeholder={placeholderText}
            />
            <Button onClick={addTextareaText}>{buttonText}</Button>
        </form>
	);
}

export default ModalFormTextarea;
