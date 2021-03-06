import config from 'config';
import fetch from '../utils/fetch';

const fetchTimeout = 3000;

export default {
    get(path) {
        let url = `${config.apiUrl}${path}`;
        console.log(`GET ${url}`)

        const requestOptions = {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        };

        return fetch(url, requestOptions, fetchTimeout)
            .then(handleResponse);
    },

    post(path, data) {
        let url = `${config.apiUrl}${path}`;
        console.log(`POST ${url}`)

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data),
        };

        return fetch(url, requestOptions, fetchTimeout)
            .then(handleResponse);
    }
}

function handleResponse(response) {
    return response.text().then(text => {
        let data = null;
        try {
            data = text && JSON.parse(text);
        } catch (e) {
        }

        if (!response.ok) {
            if (response.status === 401) {
                // auto logout if 401 response returned from api
                localStorage.removeItem('user');
                location.reload();
            }

            const error = data || text || response.statusText;
            return Promise.reject(error);
        }

        return data;
    });
}
