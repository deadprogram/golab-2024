```mermaid
flowchart LR
    subgraph wasmVision engine
        Capture
        Runtime[WASM Runtime]
        Capture--frame-->Runtime
        Capture<-->OpenCV
        Runtime<-->OpenCV
    end
    subgraph wasmVision processors
        Runtime--frame-->ollama.wasm
        ollama.wasm--frame-->mosaic.wasm
    end
    ollama.wasm--frame-->ollama
    subgraph ollama
        llava
    end
```
