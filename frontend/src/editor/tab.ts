import { Model } from "./model";
import { State } from "./state";

export default class Tab {
    model: Model;
    state: State;
    constructor(model: Model, state: State) {
        this.model = model;
        this.state = state;
    }

    get_model(): Model {
        return this.model;
    }

    get_state(): State {
        return this.state;
    }

    updateState(state: State) {
        this.state = state;
    }
}
