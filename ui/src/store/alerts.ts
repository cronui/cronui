import {Store} from './state';

interface Alert extends Object {
    type: string
    message: string
}

class AlertStore extends Store<Array<Alert>> {

    protected data(): Array<Alert> {
        return [];
    }

    success(message: string) {
        this.add('alert-success', message)
    }

    error(message: string) {
        this.add('alert-danger', message)
    }

    add(type: string, message: string) {
        let alert = {type, message};
        this.state.push(alert);
        setTimeout(() => this.remove(alert), 3000);
    }

    remove(alert: Alert) {
        const index = this.state.indexOf(alert);
        if (index > -1) {
            this.state.splice(index, 1);
        }
    }

    clear() {
        this.state = [];
    }

}

export const alertStore: AlertStore = new AlertStore()
