```mermaid
flowchart LR
subgraph videodrone
    subgraph joystick
        dualshock4
    end
    subgraph video
        GoCV<-->dnn
    end
    subgraph dnn
        YOLOv8
    end
    subgraph drone
        dualshock4-->flightcontrol
        GoCV<-->videostream
    end
end
subgraph nvidia
    YOLOv8<-- CUDA -->GPU
end
subgraph DJI Tello
    flightcontrol<-- WiFi -->droneflight
    videostream<-- WiFi -->dronevideo
end
```
