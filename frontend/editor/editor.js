require.config({ paths: { "vs": "https://unpkg.com/monaco-editor@0.23.0/min/vs" } });
import Tab from "./tab.js";

export default class Editor {
    constructor() {
        require(["vs/editor/editor.main"], function () {
            let map = new Map();
            map.set("config", new Tab(monaco.editor.createModel("Hello from config", "txt"), null));
            map.set("algo", new Tab(monaco.editor.createModel("Hello from algo", "txt"), null));
            this.tabs = map;
            this.editor = monaco.editor.create(document.getElementById("editor"), {
                model: tabs.get("config").model,
                theme: "vs-dark",
            });
        });
    }

    change_tab() {
        let button = document.getElementById("change_tab");
        let currentState = editor.saveViewState();
        let nextTab = null;
        if (button.textContent == "Algorithm") {
            button.textContent = "Configuration";
            tabs.get("config").updateState(currentState);
            nextTab = tabs.get("algo");
        } else if (button.textContent == "Configuration") {
            button.textContent = "Algorithm";
            tabs.get("algo").updateState(currentState);
            nextTab = tabs.get("config");
        }
        editor.setModel(nextTab.get_model());
        editor.restoreViewState(nextTab.get_state());
        editor.focus();
    }
}
