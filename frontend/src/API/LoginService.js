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
            'My-Custom-Header': 'foobar',
            'Access-Control-Allow-Origin': 'X-Uid, X-Authentication'
        };

        const response = await axios.post('http://127.0.0.1:8000/login', article, {headers})
        return response;
    }
}
