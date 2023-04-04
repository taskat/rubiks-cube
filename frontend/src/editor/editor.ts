import Tab from "./tab";
import * as monaco from "monaco-editor";
import { State } from "./state";


export default class Editor {
    tabs: Map<string, Tab>;
    editor: monaco.editor.IStandaloneCodeEditor;
    constructor() {
        let map = new Map();
        map.set("config", new Tab(monaco.editor.createModel("Hello from config", "txt"), null));
        map.set("algo", new Tab(monaco.editor.createModel("Hello from algo", "txt"), null));
        this.tabs = map;
        this.editor = monaco.editor.create(document.getElementById("editor") as HTMLElement, {
            model: this.tabs.get("config")?.model,
            theme: "vs-dark",
        });
    }

    changeTab() {
        let button = document.getElementById("change_tab") as HTMLButtonElement;
        let currentState: State = this.editor.saveViewState();
        let nextTab = null;
        if (button.textContent === "Algorithm") {
            button.textContent = "Configuration";
            this.tabs.get("config")?.updateState(currentState);
            nextTab = this.tabs.get("algo") as Tab;
        } else {
            button.textContent = "Algorithm";
            this.tabs.get("algo")?.updateState(currentState);
            nextTab = this.tabs.get("config") as Tab;
        }
        this.editor.setModel(nextTab.get_model());
        this.editor.restoreViewState(nextTab.get_state());
        this.editor.focus();
    }

    getText(tab: string): string {
        return this.tabs.get(tab)?.model?.getValue() as string;
    }
}
