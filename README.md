## About me

My name is Kristo Godari. Here you can find my background: https://www.linkedin.com/in/kristo-godari/

I'm new to go, basically this is my first go project. Coming from a java world, it was challenging and fun to write this project in go. I did not know the sintax, paddigram , best practices and go way of doing thing. I searched in the internet, hopfully i have done a good job. 

## How to run the project
Clone repository, then run `docker-compose build` and after run `docker-compose up`. By default all tests will run.

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

