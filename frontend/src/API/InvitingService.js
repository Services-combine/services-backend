import $api from ".";

export default class InvitingService {
    static async fetchMainFolders() {
        return $api.get('/user/inviting/')
    }

    static async createFolder(folderName) {
        return $api.post('/user/inviting/create-folder', {name: folderName})
    }
}