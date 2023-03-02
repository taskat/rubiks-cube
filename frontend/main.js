import Editor from "./editor/editor.js";

function main() {
    document.querySelector("#check").addEventListener("click", check);
    document.querySelector("#solve").addEventListener("click", solve);

    let editor = new Editor();

    document.querySelector("#change_tab").addEventListener("click", editor.change_tab);

}

function check() {
    console.log("check");
}

function solve() {
    console.log("solve");
}

main();