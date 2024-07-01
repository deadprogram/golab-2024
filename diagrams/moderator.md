```mermaid
flowchart LR
subgraph mqtt broker
    discuss
end
subgraph moderator
    controller
end
subgraph Adafruit Macropad
    customkeys[tinygo-keyboard] -- USB-HID --> controller
end
moderator -- publish --> discuss
```
