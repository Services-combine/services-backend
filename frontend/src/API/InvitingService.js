import $api from ".";

export default class InvitingService {
    static async fetchFolders() {
        return $api.get('/user/inviting/')
    }

    static async createFolder(folderName) {
        return $api.post('/user/inviting/create-folder', {name: folderName})
    }

    static async fetchDataFolder(folderID) {
        return $api.get(`/user/inviting/${folderID}`)
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

    static async deleteFolder(folderID) {
        return $api.get(`/user/inviting/${folderID}/delete`)
    }
}