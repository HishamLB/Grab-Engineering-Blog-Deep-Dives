# [Article Link](https://engineering.grab.com/quotas-service)

Author(s): [NAME](LINK)

[Implementation](./implementation/README.md)

# Notes:

Migrating to a microservice architecture requires thinking about problems that don't exist for a monolothic service. For example, load balancing, monitoring, rate-limiting and so on.
Quotos is meant to be a scalable API request rate limiting solution that mitigates service abuse and service failures. 
Global rate limiting means that regardless of the service instance a client calls, it will be subjected to one global API quota. While local rate limiting is for each service
Global rate limiting is difficult to implement in a huge microsservice environment.
Quotas guarantess SLA (service level agreement) which means it basically throttles certain requests made ot services to avoid cascading failures.
It uses an asynchronous processing pipeline to avoid any delays. 
Quotas' architecture revolves around Kafka and an S3 bucket for analysis/logging.
When a request comes in from a service, the Quotas middleware will intercept the request and call the client SDK for a rate limiting decision based on the API and client information. Kafka is used as the stream to update the client SDK. That means if the previous rate limiting decision is true and the Kafka stream now says false, the in-memory cache will be updated to say false. 
Quotas uses DataDog and Pagerduty for monitoring and alerting (to slack).
Quotas supports gRPC and REST. It also uses sliding window algorithm on 1 and 5 second levels.
Benchmark shows an average added delay of only 200 ms.

# Comments/Analysis:
Internal API rate-limiting is something that I've heard being important when it comes to the frontend but I've never looked at it closer between backend services. The part with redis expiry key Garbage Collection is quite cool I wish they went more in depth about that. 
The way caching is used is quite clever and also very impressive considering this is happening asynchronously. Golang seems quite convenient for concurrency.

