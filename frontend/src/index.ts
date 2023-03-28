import Simulator from "./three-app";
import Editor from "./editor/editor";

import "./styles/styles.css";
import Client from "./httpclient/client";
import { Side } from "./cube/side";

class App {
  editor: Editor;
  simulator: Simulator;
  changeTab: () => void;
  httpClient: Client;
  constructor() {
    this.editor = new Editor();
    this.simulator = new Simulator();
    this.changeTab = this.editor.changeTab.bind(this.editor);
    this.httpClient = new Client();
  }
  async start() {
    document.querySelector("#check")?.addEventListener("click", this.check.bind(this));
    document.querySelector("#solve")?.addEventListener("click", this.solve.bind(this));
    document.querySelector("#change_tab")?.addEventListener("click", this.changeTab);
    document.addEventListener('keyup', (e) => this.keyboardShortcuts(e), false);
    this.simulator.init();
  }
  
  keyboardShortcuts(e: KeyboardEvent) {
    if (e.ctrlKey && e.key === "Enter") {
      this.solve();
    } else if (e.shiftKey && e.key === "Enter") {
      this.check();
    } else if (e.shiftKey && e.key === "Tab") {
      this.changeTab();
    }
  }
  
  check() {
    let content = this.editor.getText("config");
    this.httpClient.postConfig(content).then((data) => {
      const stringToColors: Map<string, string[][]> = new Map(Object.entries(data.sides));
      let sideToColors: Map<Side, string[][]> = new Map();
      stringToColors.forEach((colors, key) => {
        const side = Side[key as unknown as keyof typeof Side];
        sideToColors.set(side, colors);
      });
      this.simulator.colorPalette = sideToColors;
      this.simulator.recreateUiPieces();
    });
  }
  
  solve() {
    console.log("solve");
  }
}


function main() {
  const app = new App();
  app.start();
}

main();
