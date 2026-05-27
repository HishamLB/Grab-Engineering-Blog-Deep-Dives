# [Article Link](https://engineering.grab.com/marionette-enabling-e2e-user-scenario-simulation)

Author(s): Anish Jha, Biju Joseph Jacob, Phuc Lam Ngyuen, Vineet Nair, Yiwei Yeo
[Implementation](./implementation/README.md)

# Notes:
Grab has a lot of microservices and if one breaks the whole app may get unstable. This is why they concduct E2E testing. 
It is getting difficult to write E2E tests properly for all the microservices Grab powers.
These difficulties include: 

- Availability: Getting all the microservices together for E2E is difficult as each dev team works independently and is responsible only for their microservices.
- Data or resource set up
- Access and Authentication
- Resource and time intensive

Approach: 
Create a platform that simulates user scenarios before any new versions of microservices are released. The platform should set up the data required for these simulations. 

Solution: Marionette
Devs set up data as well as configurations to mimic real-world behaviour. They can interface with Marionette through UI, SDK and even RESTful API. 
For independent simulations actors are grouped into "cohorts" that only exist for that specific simulation.


# Comments/Analysis:
This is actually very impressive. Marionette seems like such a powerful tool and probably very important for QA. I would imagine it saves devs and testers a lot of time writing boilerplate code and the sort. 

