# Coin Chameleon

A Simple Binance Kline Crawler

## Usage

Build and run your Oracle Database

Install [Go](https://golang.org/doc/install)

Download [Oracle Instant Client](https://www.oracle.com/database/technologies/instant-client/downloads.html)

Put instant client under: `~/instantclient_19_8`

Install dependencies with go mod

```sh
go mod download
```

Create and Provide environment data in `.env`

Run app

```sh
make dev
```

## References

- [How to Connect a Go Program to Oracle Database using godror](https://blogs.oracle.com/developers/how-to-connect-a-go-program-to-oracle-database-using-goracle)
- [Quick and Easy Setup – Oracle Xe and ORDS Using Docker](https://learncodeshare.net/2019/04/22/quick-and-easy-setup-oracle-xe-and-ords-using-docker/)
- [Oracle Database container images](https://github.com/oracle/docker-images/tree/main/OracleDatabase/SingleInstance#running-sqlplus-in-a-container)

## License

MIT
