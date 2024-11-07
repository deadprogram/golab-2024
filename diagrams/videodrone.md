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
        YuNet
    end
    subgraph drone
        dualshock4-->flightcontrol
        GoCV<-->videostream
    end
end
subgraph nvidia
    YuNet<-- CUDA -->GPU
end
subgraph DJI Tello
    flightcontrol<-- WiFi -->droneflight
    videostream<-- WiFi -->dronevideo
end
```
