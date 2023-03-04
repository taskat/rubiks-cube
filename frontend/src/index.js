import threeApp from "./three-app"
import Editor from "./editor/editor.js"

import "./styles/styles.css"

const main = async () => {
  document.querySelector("#check").addEventListener("click", check);
  document.querySelector("#solve").addEventListener("click", solve);
  let editor = new Editor();
  const changeTab = () => editor.change_tab();
  document.querySelector("#change_tab").addEventListener("click", changeTab);
  document.addEventListener('keyup', (e) => keyboard_shortcuts(e, changeTab), false);
  const threeAppActions = threeApp()

  threeAppActions.init()
}

function keyboard_shortcuts(e, changeTab) {
  if (e.ctrlKey && e.keyCode === 13) {
    solve();
  } else if (e.shiftKey && e.keyCode === 13) {
    check();
  } else if (e.shiftKey && e.keyCode === 9) {
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
