import {makeAutoObservable} from "mobx";
import axios from "axios";
import { API_URL } from "../API";
import AuthService from "../API/AuthService";

export default class Store {
    isAuth = false;
    isLoading = false;
    isError = null;
    timeout = 3000;

    constructor() {
        makeAutoObservable(this);
    }

    setAuth(bool) {
        this.isAuth = bool;
    }

    setLoading(bool) {
        this.isLoading = bool;
    }

    setError(error) {
        this.isError = error;
    }

    async login(username, password) {
        try {
            const response = await AuthService.login(username, password);
            localStorage.setItem('token', response.data.accessToken);
            this.setAuth(true);
        } catch (e) {
            this.setError(e.response?.data?.message);
            setTimeout(() => {
                this.setError(null)
            }, this.timeout)
        }
    }

    async logout() {
        try {
            const response = await AuthService.logout()
            localStorage.removeItem('token');
            this.setAuth(false);
        } catch (e) {
            this.setError(e.response?.data?.message);
            setTimeout(() => {
                this.setError(null)
            }, this.timeout)
        }
    }

    async checkAuth() {
        this.setLoading(true);
        try {
            const response = await axios.get(`${API_URL}/refresh`, {withCredentials: true})
            localStorage.setItem('token', response.data.accessToken);
            this.setAuth(true);
        } catch (e) {
            this.setError(e.response?.data?.message);
            setTimeout(() => {
                this.setError(null)
            }, this.timeout)
        } finally {
            this.setLoading(false);
        }
    }
}
