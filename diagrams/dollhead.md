```mermaid
flowchart LR
subgraph Microcontroller
    USB
    GPIO
    PWM
end
GPIO --> WS2812Head[WS2812 Head LEDs]
GPIO --> WS2812Collar[WS2812 Collar LEDs]
PWM --> Servo
Computer <--> USB
```
