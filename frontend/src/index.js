import React from "react"
import ReactDOM from "react-dom"
import { ThreeAppActionsProvider } from "./context"
import App from "./App"
import * as serviceWorkerRegistration from "./serviceWorkerRegistration"
import threeApp from "./three-app"
import "./styles/styles.css"
import Editor from "./editor/editor.js"

const main = async () => {
  document.querySelector("#check").addEventListener("click", check);
  document.querySelector("#solve").addEventListener("click", solve);
  let editor = new Editor();
  document.querySelector("#change_tab").addEventListener("click", () => editor.change_tab());
  document.addEventListener('keyup', (e) => keyboard_shortcuts(e, editor), false);
  const threeAppActions = threeApp()

  ReactDOM.render(
    <React.StrictMode>
      <ThreeAppActionsProvider threeAppActions={threeAppActions}>
        <App />
      </ThreeAppActionsProvider>
    </React.StrictMode>,
    document.getElementById("react-container")
  )

  // If you want your app to work offline and load faster, you can change
  // unregister() to register() below. Note this comes with some pitfalls.
  // Learn more about service workers: https://cra.link/PWA
  serviceWorkerRegistration.register()

  threeAppActions.init()
}

function keyboard_shortcuts(e, editor) {
  if (e.ctrlKey && e.keyCode === 13) {
    solve();
  } else if (e.shiftKey && e.keyCode === 13) {
    check();
  } else if (e.shiftKey && e.keyCode === 9) {
    editor.change_tab();
  }
}

function check() {
  console.log("check");
}

function solve() {
  console.log("solve");
}

main()
