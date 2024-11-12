```mermaid
flowchart LR
    subgraph engine
        Capture
        Runtime[WASM Runtime]
        Capture--frame-->Runtime
        Capture<-->OpenCV
        Runtime<-->OpenCV
    end
    subgraph processors
        Runtime--frame-->tinygo-processor.wasm
        Runtime--frame-->rust-processor.wasm
        Runtime--frame-->c-processor.wasm
    end
```
