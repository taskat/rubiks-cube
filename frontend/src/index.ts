import Simulator from "./three-app";
import Editor from "./editor/editor";

import "./styles/styles.css";

const main = async () => {
  document.querySelector("#check")?.addEventListener("click", check);
  document.querySelector("#solve")?.addEventListener("click", solve);
  let editor = new Editor();
  const changeTab = () => editor.change_tab();
  document.querySelector("#change_tab")?.addEventListener("click", changeTab);
  document.addEventListener('keyup', (e) => keyboard_shortcuts(e, changeTab), false);
  const simulator = new Simulator();
  simulator.init();
}

function keyboard_shortcuts(e: KeyboardEvent, changeTab: () => void) {
  if (e.ctrlKey && e.key === "Enter") {
    solve();
  } else if (e.shiftKey && e.key === "Enter") {
    check();
  } else if (e.shiftKey && e.key === "Tab") {
    changeTab();
  }
}

function check() {
  console.log("check");
}

function solve() {
  console.log("solve");
}

main()
