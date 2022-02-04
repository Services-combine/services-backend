import $api from ".";

export default class Services {
    static async fetchData(userID) {
        return $api.post('/user/', {userID: userID})
    }

    static async saveSettings(userID, countInviting, countMailing) {
        return $api.post('/user/save-settings', {userID: userID, countInviting: countInviting, countMailing: countMailing})
    }
}