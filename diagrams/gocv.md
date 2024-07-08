```mermaid
flowchart LR
    Mat1[Mat]
    Mat2[Mat]
    subgraph videoio
        camera
        file
        stream
    end
    subgraph imgproc
        gaussianBlur
    end
    subgraph dnn
        detect
    end
    camera-->Mat1
    Mat1-->gaussianBlur
    gaussianBlur-->Mat2
    Mat2-->detect
```
