import $api from ".";

export default class AuthService {
    static async login(username, password) {
        const headers = { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer my-token',
            'Access-Control-Allow-Origin': 'X-Uid, X-Authentication'
        };
        return $api.post('/login', {username, password})
    }

    static async logout() {
        return $api.post('/logout')
    }
}
