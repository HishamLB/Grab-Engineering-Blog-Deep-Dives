# [Article Link](https://engineering.grab.com/mockers)

Author(s): Mayan Gupta, Vineet Nair, Shivkumar Krishnan, Thuy NgyNguyen, Vishal Prakash

[Implementation](./implementation/README.md)

# Notes:
Having a lot of microservices means Grab has some difficulty testing:
- Little centralised management since each team is responsible for only its microservices.
- Teams use different programming languagees so it's hard to make a good test environment that covers everything.
- Maturity levels differs.

Testing in a staged environments has a few limitations
- Ownership is ambiguous as there is no centralised management so no one is sure about who fixes what.
- Cost of testing in staging is high because the environment constantly changes and configuration might be inconsistent.
- Hard to simulate negative cases because of dependencies on other microservices.

Solution:
Mockers, run tests locally. It is a Go SDK that simulates a staging environment on dev local boxes. Basically, lets you create mock servers for mimicking behaviours of microservice dependencies.
The way it works is you have one standard mock server per microservice and then it is available for all other teams for testing.
You set up the expectations for the mock server and then you can run the tests. Mock servers do not process the requests but instead just return the specific response.

You can also apply "chaos" which is simulating errors, CPU spikes and so on:
```mockServer.Inject().Error(500).ApplyTo(1).Build()```


# Comments/Analysis:
So basically black-box testing taken to another level. I think I'll have to do quite a lot more research on this topic because some of this was explained at a very high-level and I'm not sure how it would work in code and lower-level. 

