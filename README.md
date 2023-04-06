# Sports API in Go

This application provides a RESTful API that allows you to manage a database of sports teams and players. The API is built using Go programming language and SQLite database.

## TOC
- [Sports API in Go](#sports-api-in-go)
  * [Features](#features)
  * [Getting Started](#getting-started)
  * [API Endpoints](#api-endpoints)
  * [Configuration](#configuration)
  * [Contributing](#contributing)


## Features

-   Create your own endpoint `/api/<your-name>` that returns a JSON object with your name and city.
-   Add a SQLite database to the application.
-   Tables: Team (id, name, city, foundedYear) and Player (id, name, team_id, jerseyNumber, birthYear).
-   Seed data for teams and players through the seed data function.
-   Endpoint `/api/team` that returns a list of all teams.
-   Endpoint `/api/team/<team-id>` that provides information about a team, including a list of all its players.
-   Endpoint `/api/player/<player-id>` that provides information about a player, including all their details as well as the name of the team they belong to.
-   Configuration files `config.yml` and `configProduction.yml` with different connection strings for development and production environments, respectively.

## Getting Started

1.  Clone this repository.
2.  Install the necessary dependencies by running `go mod download`.
3.  Create a `config.yml` file and populate it with your database connection details.
4.  Seed the database by running the `go run main.go --seed` command.
5.  Start the application by running `go run main.go`.

## API Endpoints

The following endpoints are available:

-   `GET /api/<your-name>` - returns a JSON object with your name and city.
-   `GET /api/team` - returns a list of all teams.
-   `GET /api/team/<team-id>` - returns information about a team, including a list of all its players.
-   `GET /api/player/<player-id>` - returns information about a player, including all their details as well as the name of the team they belong to.

## Configuration

You can customize the application's configuration by modifying the `config.yml` file. In addition, you can create a `configProduction.yml` file to specify different settings for the production environment.

## Contributing

Contributions to this project are welcome. If you find any bugs or have suggestions for new features, please open an issue or submit a pull request.
