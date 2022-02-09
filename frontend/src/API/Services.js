import $api from ".";

export default class Services {
    static async fetchData() {
        return $api.get('/user/')
    }

    static async saveSettings(countInviting, countMailing) {
        return $api.post('/user/save-settings', {countInviting: Number(countInviting), countMailing: Number(countMailing)})
    }
}