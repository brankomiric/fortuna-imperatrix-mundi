# fortuna-imperatrix-mundi

## Prerequisites
This project is created using [Go](https://go.dev) and a newer installation (e.g. 1.23) is recommended to run this project. [Fiber](https://gofiber.io/) lib is used for serving HTTP and [Sqlx](https://github.com/jmoiron/sqlx) for DB connections.

## DB Setup
To run an instance of MySQL you can use the provided [docker-compose file](docker-compose.yml). Running Docker installation is needed.
```bash
docker-compose up
```

## Running
Prior to running the project please:
- start the MySQL server
- create a *.env* file where you can set the application port and DB connection information. [.env.example](.env.example) is        provided for reference.
- run the DB migrations

### Starting the app
Using Go
```bash
go run cmd/fortuna-imperatrix-mundi/main.go
```

Or, you can build and then start the executable
```bash
go build cmd/fortuna-imperatrix-mundi/main.go
```

## Migrations
[Golang-migrate](https://github.com/golang-migrate/migrate) tool is used for working with migrations. You can use [this script](runMigrations.sh) for running migrations against the DB. Migrations are placed in the [db/migrations](db/migrations) directory. Currently, there are migrations for creating players, tournaments and bets tables and DistributePrizes stored procedure.
When executing the script, DB connection string is needed to be passed as parameter:
```bash
./runMigrations.sh "mysql://root:J88d44Jq5ekG@tcp(localhost:5433)/fortuna_imperatrix_mundi"
```

In the [scripts](db/scripts) section, a convenience script for populating players table can be found.

## Endpoints
*POST /tournaments/create*
Example req body:
```json
{
    "name": "texas-holdem-01",
    "start_date": "2025-03-11T10:00:00Z",
    "end_date": "2025-03-15T00:00:00Z"
}
```

*POST /players/bet*
Example req body:
```json
{
    "player_id": 1,
    "tournament_id": 1,
    "amount": 100
}
```

*POST /tournaments/prizes/distribute/:tournament_id*
Invokes the Stored Procedure

*GET /players/rank*
Returns list of players ranked by balance

---

### Qs:
*What did you learn and if you encountered any challenges, how did you overcome them?*

I kept the application simple, not diverging much from the requirements, while also trying to stick to good principles (e.g. including migration tools, concern separation in app code, working with abstract structures for easier test implementation later, etc.). There where no major challenges.

*What did you take in consideration to ensure the query and or stored procedure is efficient and handles edge cases?*

Table structures are very simple. No columns are indexed (save for primary and foreign keys) - which is fine as we are not querying on non-indexed columns atm. 

_PlaceBet_ function is the most complex as it: updates the player balance, inserts into bets table and updates the tournaments prize column. That is why these are ran in a transaction.

Before running the stored procedure a check is made to skip calling the procedure if tournament is still in progress. A possible improvement for early return would be to also check if tournament hasn't started yet. Inside the procedure a check is made to exit if prizes are distributed already.

Another edge case that is handled (although in a very simple way) is when players have placed equal bets - earliest placed bet is given advantage.

*If you used CTEs or Window Functions, what did you learn about their power and flexibility?*

I used a simple CTE in the _GetPlayersRankedByBalance_ function just for show as the same functionality could have been achieved with a SELECT and a ORDER BY. But I can imagine how CTEs can be useful for more complex queries. They can also greatly improve readability of complex queries. 

*How might you apply the technique in more complex scenarios?*

CTEs can improve complex query readability and allow subquery reusability. 