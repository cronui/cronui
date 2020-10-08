import api from './api';

export const installService = {
    isInstalled,
    postUser,
    clearStore,
};

const store = {
    installed: null
};

async function isInstalled() {
    if (store.installed === null) {
        try {
            store.installed = await api.get('/install').then(data => data.installed);
        } catch (e) {

        }
    }

    return store.installed;
}

function postUser(username, password) {
    store.installed = null;
    return api.post('/install', {username, password});
}

function clearStore() {
    store.installed = null;
}
