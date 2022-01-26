import axios from "axios";

export default class IndexService {
    static async index() {
        const headers = { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer my-token',
            'Access-Control-Allow-Origin': 'X-Uid, X-Authentication'
        };

        const response = await axios.get('http://127.0.0.1:8000/api/user/', {headers})
        return response;
    }
}
