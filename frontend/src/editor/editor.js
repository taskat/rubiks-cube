import Tab from "./tab.js";
import * as monaco from "monaco-editor";


export default class Editor {
    constructor() {
        let map = new Map();
        map.set("config", new Tab(monaco.editor.createModel("Hello from config", "txt"), null));
        map.set("algo", new Tab(monaco.editor.createModel("Hello from algo", "txt"), null));
        this.tabs = map;
        this.editor = monaco.editor.create(document.getElementById("editor"), {
            model: this.tabs.get("config").model,
            theme: "vs-dark",
        });
        console.log("created editor");
        console.log(this.editor === undefined);
    }

    change_tab() {
        let button = document.getElementById("change_tab");
        console.log(this.constructor.name);
        console.log(this.editor === undefined);
        let currentState = this.editor.saveViewState();
        let nextTab = null;
        if (button.textContent == "Algorithm") {
            button.textContent = "Configuration";
            this.tabs.get("config").updateState(currentState);
            nextTab = this.tabs.get("algo");
        } else if (button.textContent == "Configuration") {
            button.textContent = "Algorithm";
            this. tabs.get("algo").updateState(currentState);
            nextTab = this.tabs.get("config");
        }
        this.editor.setModel(nextTab.get_model());
        this.editor.restoreViewState(nextTab.get_state());
        this.editor.focus();
    }
}
