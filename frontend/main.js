import Editor from "./editor/editor.js";

function main() {
    document.querySelector("#check").addEventListener("click", check);
    document.querySelector("#solve").addEventListener("click", solve);

    let editor = new Editor();

    document.querySelector("#change_tab").addEventListener("click", editor.change_tab);

    document.addEventListener('keyup', (e) => keyboard_shortcuts(e, editor), false);
}

function keyboard_shortcuts(e, editor) {
    if (e.ctrlKey && e.keyCode == 13) {
        solve();
    } else if (e.shiftKey && e.keyCode == 13) {
        check();
    } else if (e.shiftKey && e.keyCode == 9) {
        editor.change_tab();
    }
}

function check() {
    console.log("check");
}

function solve() {
    console.log("solve");
}

main();