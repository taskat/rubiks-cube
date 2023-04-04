import Simulator from "./three-app";
import Editor from "./editor/editor";

import "./styles/styles.css";
import Client from "./httpclient/client";
import { Error } from "./error/error";

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

  addError(error: Error) {
    const errorList = document.getElementById("errorlist");
    const errorItem = document.createElement("tr");
    errorItem.innerHTML = `
    <td>${error.level}</td>
    <td>${error.text}</td>
    <td>Line ${error.pos.line}, Column ${error.pos.column}</td>
    <td>${error.file}</td>
    `;
    errorList?.appendChild(errorItem);
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
    this.httpClient.postConfig(content).then((response) => {
      this.simulator.colorPalette = response.colorPalette;
      this.updateErrors(response.errors);
      this.simulator.recreateUiPieces();
    });
  }

  showErrorTable() {
    const errortable = document.getElementById("errortable");
    if (errortable) {
      errortable.hidden = false;
    }
    const noerror = document.getElementById("noerror");
    if (noerror) {
      noerror.hidden = true;
    }
  }

  showNoError() {
    const errortable = document.getElementById("errortable");
    if (errortable) {
      errortable.hidden = true;
    }
    const noerror = document.getElementById("noerror");
    if (noerror) {
      noerror.hidden = false;
    }
  }

  solve() {
    console.log("solve");
  }

  updateErrors(errors: Error[]) {
    if (errors.length === 0) {
      this.showNoError();
    } else {
      this.showErrorTable();
    }
    const errorList = document.getElementById("errorlist");
    if (errorList) {
      errorList.innerHTML = "";
    }
    errors.forEach((error: Error) => this.addError(error));
  }
}


function main() {
  const app = new App();
  app.start();
}

main();
