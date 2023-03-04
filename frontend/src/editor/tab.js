export default class Tab {
    constructor(model, state) {
        this.model = model;
        this.state = state;
    }

    get_model() {
        return this.model;
    }

    get_state() {
        return this.state;
    }

    updateState(state) {
        this.state = state;
    }
}
