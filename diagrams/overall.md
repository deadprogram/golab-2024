```mermaid
flowchart LR
subgraph mqtt broker
    discuss
    responses
end
subgraph panelists
    panelist-1
    panelist-2
    panelist-3
end
subgraph moderator
    controller
end
moderator -- publish --> discuss
discuss-- subscribe -->panelists
panelists-- publish -->responses
responses-- subscribe -->panelists
```
