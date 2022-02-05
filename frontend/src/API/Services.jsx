import $api from ".";

export default class Services {
    static async fetchData(userID) {
        return $api.post('/user/', {id: userID})
    }

    static async saveSettings(userID, countInviting, countMailing) {
        return $api.post('/user/save-settings', {id: userID, countInviting: Number(countInviting), countMailing: Number(countMailing)})
    }
}