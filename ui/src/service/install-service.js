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
        store.installed = await api.get('/install').then(data => data.installed);
    }

    return store.installed;
}

function postUser(username, password) {
    return api.post('/install', {username, password});
}

function clearStore() {
    store.installed = null;
}
