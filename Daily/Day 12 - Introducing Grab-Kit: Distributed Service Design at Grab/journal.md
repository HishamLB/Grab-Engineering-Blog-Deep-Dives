# [Article Link](https://engineering.grab.com/introducing-grab-kit)

Author(s): Karen Kue, Micheal Cartmell

[Implementation](./implementation/README.md)

# Notes:

Moving away from Monolithic to Microservices prompted Grab to create a framework to build Go microservices. This probvides abstractions of distributed systems and lets developers focus on business logic.
This saves time spent creating boilerplate code. 
Grab-kit maintains a ```.proto``` file (which is a configuration file that defines contracts for gRPC). This file will be used to generate code related to DTOs for the custom types in the ```.proto``` file. It also generates bindings so that Go DTOs can be converted to protobuf types.
Grab-kit uses a common middleware stack that deals with logging and stats.
Grab-kit also generates dashboards and graphs for monitoring the services and upstream dependencies.


# Comments/Analysis:
Distributed systems are quite hard to expect developers to work with. It's been figured out ages ago that abstractions for distributed systems are very valuable. A classic model is Google's MapReduce, which is a framework Google arrived to when they wanted to hide the complexity of distributed systems and let their developors focus on just writing code.
A lot of these blogs seem to focus on how to make microservices less of a hassle to work with and sort of standardize the way things are created downstream. Probably heaivly opinionated.

