import "./wasm_exec.js"

const go = new Go();
WebAssembly.instantiateStreaming(fetch("../skilltreetool.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
