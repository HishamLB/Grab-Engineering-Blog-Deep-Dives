# [Article Link](https://engineering.grab.com/feature-toggles-ab-testing)

Author(s): Roman Atachiants 

[Implementation](./implementation/README.md)

# Notes:
Feature toggles lets development continue while algorithms/new features are being worked on.
This is very helpful in A/B Testing and also specific deployment plans.
A centralized platform is used in Grab to toggle features, manage roll outs, and set configurations.

Minimizing network calls was priority so instead of asking the server everytime a flag is needed they do a preload config where the SDK downloads all feature rules periodically from S3 and store them locally in memory.

The API has a very simple interface with 2 functions:

```
    GetVariables()
    Track()
```
Facets are contexts about the request like Passenger ID, Driver ID etc.
They are metadata used to decide which experiment group someone belongs to.

# Comments/Analysis:

Impressive, I'm surprised the API interface is this simple but it makes sense for the use case.
The way Grab translated business rules here is quite interesting too. Instead of repeated network calls, in-memory can be a way more efficient take. 

