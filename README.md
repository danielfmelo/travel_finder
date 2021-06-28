# TRAVEL_FINDER

Travel Finder is a simple application where the goal is to find the cheapest trip between two points doesn't matter how many hops. The solution implemented is based in the[Shortest path algorytm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm).

## Project structure

The project follows the Clean architecture where `deliveries` and `repositories` are the external interfaces, `services` are the use cases and `entity` is the business rules.

## Prerequisites

There are two ways to build the project: one is using Golang version >= 1.14 and the second is using Docker.


## How to use it

For your convenience you can use `make` commands to run the tests and the projects.

### Running

You can execute the app with the following commands:

```shell
make run database=database.csv
or
make docker-run database=database.csv
```

If the binary was already built, you can run using it:

```shell
./travel_finder database.csv
```

### Build

If you would like to only build and not run the application, you can use:

```shell 
make build
make docker-build
```

### Tests

The same way from other commands, you can run the tests using the make commands:

```shell
make tests
make docker-tests
```

You can also see the code coverage info typing:

```shell
make coverage
```

It will open a webpage with the coverage data.


## Rest API

The API communication is given through port number 8080. There are two routes for this API.


### Insert new route:

To insert a new route you should use the request below:

	POST /origins/:origin/destinations/:destination/values/:value

Sample:

```shell
curl -X POST -v http://localhost:8080/origins/gru/destinations/fln/values/10
```

Response codes:

* 201
* 400
* 500


### Find the cheapest route: 

To find the cheapest route between two points, you should use the request:

    GET /origins/:origin/destinations/:destination

Sample:

```shell
curl -v http://localhost:8080/origins/gru/destinations/orl
```

Response codes:
* 200
* 400
* 404
* 500

## Room for improvements

Some improvements points to be implemented in future versions:

* Create a proper API documentation using [swagger](https://swagger.io/);
* Refactor the package `command` to be able to implement unit test;
* Handle with the current routes instead of duplicate them;
* Use a database instead of a .csv file;
