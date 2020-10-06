import {Store} from './state';

interface Alert extends Object {
    type: string
    message: string
}

class AlertStore extends Store<Alert> {

    protected data(): Alert {
        return {
            type: null,
            message: null,
        };
    }

    success(message: string) {
        this.state.type = 'alert-success';
        this.state.message = message;
    }

    error(message: string) {
        this.state.type = 'alert-danger';
        this.state.message = message;
    }

    clear() {
        this.state.type = null;
        this.state.message = null;
    }

}

export const alertStore: AlertStore = new AlertStore()
