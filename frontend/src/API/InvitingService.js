import $api from ".";

export default class InvitingService {
    static async fetchFolders() {
        return $api.get('/user/inviting/')
    }

    static async createFolder(folderName) {
        return $api.post('/user/inviting/create-folder', {name: folderName})
    }

    static async fetchDataFolder(folderID, limit = 20, skip = 0) {
        return $api.post(`/user/inviting/${folderID}`, {limit: limit, skip: skip})
    }

    static async createFolderInFolder(folderID, folderName) {
        return $api.post(`/user/inviting/${folderID}/create-folder`, {name: folderName})
    }

    static async renameFolder(folderID, folderName) {
        return $api.post(`/user/inviting/${folderID}/rename`, {name: folderName})
    }

    static async changeChat(folderID, chatName) {
        return $api.post(`/user/inviting/${folderID}/change-chat`, {chat: chatName})
    }

    static async addMessage(folderID, message) {
        return $api.post(`/user/inviting/${folderID}/change-message`, {message: message})
    }

    static async changeUsernames(folderID, usernames) {
        return $api.post(`/user/inviting/${folderID}/change-usernames`, {usernames: usernames})
    }

    static async changeGroups(folderID, groups) {
        return $api.post(`/user/inviting/${folderID}/change-groups`, {groups: groups})
    }

    static async moveFolder(folderID, path) {
        return $api.post(`/user/inviting/${folderID}/move`, {path: path})
    }

    static async checkBlock(folderID) {
        return $api.get(`/user/inviting/${folderID}/check-block`)
    }

    static async launchInviting(folderID) {
        return $api.get(`/user/inviting/${folderID}/launch-inviting`)
    }

    static async launchMailingUsernames(folderID) {
        return $api.get(`/user/inviting/${folderID}/launch-mailing-usernames`)
    }

    static async launchMailingGroups(folderID) {
        return $api.get(`/user/inviting/${folderID}/launch-mailing-groups`)
    }

    static async deleteFolder(folderID) {
        return $api.get(`/user/inviting/${folderID}/delete`)
    }

    static async createAccount(folderID, name, phone) {
        return $api.post(`/user/inviting/${folderID}/create-account`, {name: name, phone: phone})
    }

    static async fetchDataAccount(folderID, accountID) {
        return $api.get(`/user/inviting/${folderID}/${accountID}`)
    }

    static async saveSettingsAccount(folderID, accountID, name, interval, folder) {
        return $api.post(`/user/inviting/${folderID}/${accountID}`, {name: name, interval: interval, folder_id: folder})
    }

    static async deleteAccount(folderID, accountID) {
        return $api.get(`/user/inviting/${folderID}/${accountID}/delete`)
    }

    static async geterateInterval(folderID) {
        return $api.get(`/user/inviting/${folderID}/generate-interval`)
    }

    static async sendCodeParsing(folderID, accountID) {
        return $api.get(`/user/inviting/${folderID}/${accountID}/login-api`)
    }

    static async parsingApi(folderID, accountID, code) {
        return $api.post(`/user/inviting/${folderID}/${accountID}/parsing-api`, {password: code})
    }

    static async sendCodeSession(folderID, accountID) {
        return $api.get(`/user/inviting/${folderID}/${accountID}/get-code-session`)
    }

    static async createSession(folderID, accountID, code) {
        return $api.post(`/user/inviting/${folderID}/${accountID}/create-session`, {password: code})
    }
}