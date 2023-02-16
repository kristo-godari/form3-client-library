# Form 3 client library

## Introduction
This code is a project I completed for an interview practice assignment. 
A friend of mine had used it for their own interview preparation. 
My initial intention was to use it as a way to test myself for a potential GO-related job application. 
However, upon working through the exercise, I discovered that my interest in GO was not as strong as I had previously thought, so I ultimately decided not to apply.
Regardless, the project was an enjoyable learning experience and I had fun working with GO.

## Requirement
Requirement is detailed here: https://github.com/form3tech-oss/interview-accountapi

## How to run the project
Clone repository, then run `docker-compose build` and after run `docker-compose up`. By default, all tests will run.

## A little about project design
I have tried to build the project in such a way that can be extensible for the future. 

I have a base client, that contains common functionality(Form3BaseClient), and then an account client(Form3AccountClient) that uses the base client.
The idea is that if in the future we would like to implement a client also for Payment API, or Reports API we can do this quite easily,  by implementing a separate client that "extends" the base client.

I also thought that would be nice to use the Builder pattern, since those clients tend to have lots of parameters, and using a builder makes code much more readable.

Ps. Added some examples in [examples/examples.go](examples/examples.go) and small documentation in [docs/README.md](docs/README.md)

## Things to add for the future
For the future, if we would really want to lunch this to the public I would add:
 - Update models to contain all API fields
 - Authentication
 - Retry mechanism. (General one/Per endpoint)
 - Timeout support per endpoint  
 - Better error handling
 - Better documentation

