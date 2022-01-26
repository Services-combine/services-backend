import axios from "axios";

export default class LoginService {
    static async loginUser(username, password) {
        const article = {
            username: username,
            password: password
        };
        const headers = { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer my-token',
            'Access-Control-Allow-Origin': 'X-Uid, X-Authentication'
        };

        const response = await axios.post('http://127.0.0.1:8000/api/login', article, {headers})
        return response;
    }
}
