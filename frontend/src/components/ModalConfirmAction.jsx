import React, {useState, useEffect} from 'react'
import '../styles/Inviting.css';
import Button from './UI/button/Button';

const ModalConfirmAction = ({result}) => {
    const [action, setAction] = useState(null);

    useEffect(() => {
        addResultAction();
    }, [action])

    const addResultAction = () => {
		const newAction = {
            action, id: Date.now()
        }
        result(newAction)
	}

    return (
        <div style={{textAlign: "center"}}>
            <h5>Вы уверены?</h5>
            <Button style={{marginRight: 10, marginTop: 10}} onClick={() => setAction(true)}>Да</Button>
            <Button style={{background: "rgb(233, 62, 62)", color: "#dedede"}} onClick={() => setAction(false)}>Отмена</Button>
        </div>
	);
}

export default ModalConfirmAction;
