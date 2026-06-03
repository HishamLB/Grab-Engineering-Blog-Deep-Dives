# [Article Link](https://engineering.grab.com/loki-dynamic-mock-server-http-tcp-testing)

Author(s): Thuy Ngyuen, Mayank Gupta, Vishal Prakash, Vineet Nair

[Implementation](./implementation/README.md)

# Notes:
Inspired by Grab's [Mockers](https://engineering.grab.com/mockers). Grab created a dynamic mock server for local box texting of mobile apps.
End-to-end testing of a mobile app is difficult due to high dependencies on backend services and even other app.
It works by turning the devs local box into a pseudo backend environment. 
A few engineering challenges were encountered to make mocking TCP work since: 
- TCP doesn't really follow HTTP-like request/response pattern but it's a running connection.
- Messages can be sent to apps without an incoming request
- Needed a way to delimit incoming requests to truncate appropriately. 

For http, the workflow involves setting expectation for the request/response, then storing that. After that, the request is made and the response is observed.
For TCP, it's a bit more involved, requiring scheduled events with time taken into consideration. 

# Comments/Analysis:

Mobile testing is very tedious for many reasons and that's even more the case at an enterprise-level. Just the idea of staging dependencies make this project worth it. TCP testing is something I am not exposed to but seems interesting enough to check out someday. 
