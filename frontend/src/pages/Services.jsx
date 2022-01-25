import React, {useEffect} from 'react'
import {Link, useNavigate} from "react-router-dom"
import '../styles/Services.css';
import Loader from '../components/UI/loader/Loader';
import { useFetching } from '../hooks/useFetching';
import IndexService from '../API/IndexService';
import Error from '../components/UI/error/Error';

const Services = () => {
    const [fetchIndex, isIndexLoading, indexError] = useFetching(async () => {
		const response = await IndexService.index();
        console.log(response);
		//const resultCheckToken = response.data
	})

    useEffect(() => {
		fetchIndex()
	}, [])

    let navigate = useNavigate();

	return (
        <div className='services'>
            <h3>Сервисы</h3>
            
            <ul className="services__list btn-toolbar" role="toolbar">
                <li className="services__list-item">
                    <Link to="/inviting" className="services__list-item_link">
                        <h6 className="title">Инвайтинг & Рассылка</h6>
                    </Link>
                </li>
            </ul>

            {indexError &&
				//<Error>Произошла ошибка: {indexError}</Error>
                navigate("/login")
			}

            {isIndexLoading &&
				<div style={{display: 'flex', justifyContent: 'center', marginTop: 50}}><Loader/></div>
			}
        </div>
	);
}

export default Services;
