# Pinger

Example of a simple web service.

Send request to
~~~
/pinger/ping
~~~

## How to run in a local environment

At the first launch / after editing Docker files:

~~~
make create
make up
~~~

Access to the command line inside the running docker container

~~~
make bash
~~~

Run tests

~~~
make test
~~~