require.config({ paths: { "vs": "https://unpkg.com/monaco-editor@0.23.0/min/vs" } });
require(["vs/editor/editor.main"], function () {
    var editor = monaco.editor.create(document.getElementById("editor"), {
        value: [
            `function x() {`,
            `\tconsole.log("Hello world!");`,
            `}`
        ].join("\n"),
        language: "javascript",
        theme: "vs-dark",
        
    });
});