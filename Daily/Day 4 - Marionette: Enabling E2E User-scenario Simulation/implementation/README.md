I found the idea of the platform very cool so I decided to start a project that attempts to implement something like Marionette. 
Link to repo when I remember should be here:


# Mini Marionette

A local E2E simulation platform.

## Core Idea: 

- Simulate passengers/drivers.
- Run booking flow automatically (matching)
- Isolate simulations
- Visualize state transitions
- Expose SDK/RESTful API

## Cohorts
Cohorts describe isolated groups where the simulation will happen. The matching logic and simulation should be only consider a specific cohort for each simulation. 
Example: 

``` 
cohort_a
    |- passengers
    |- drivers
    |- bookings

cohort_b
    |- passengers
    |- drivers
    |- bookings
```
Each entity gets its own ```cohort_id``` and the matching logic only matches within cohort

An ECS-style system will be used for entity behaviours for example:
- Different driver behaviours:
    ```
    always_accept
    always_reject
    random_delay
    cancel_after_accept
    ```
- Different passenger behaviours:
    ```
    spam_booking
    cancel_after_30s
    wait_forever
    ```


