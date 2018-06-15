# sql-exporter

[![Build Status](https://travis-ci.org/inloop/sql-exporter.svg?branch=master)](https://travis-ci.org/inloop/sql-exporter)

Command line utility to export data from SQL database

# Installation

Basically download binary, make it executable and that's it.

```
# macos example
# releases: https://github.com/inloop/sql-exporter/releases
curl -L https://github.com/inloop/sql-exporter/releases/download/0.1.4/sql-exporter-darwin-amd64 > /usr/local/bin/sql-exporter
chmod +x /usr/local/bin/sql-exporter
```

# Usage

```
sql-exporter dump -u postgres://username:password@hostname/database?sslmode=required -q "SELECT * FROM table" -o output.csv

# more info by running `sql-exporter dump -h`

USAGE:
   sql-exporter dump [command options]

OPTIONS:
   -u value, --db-url value  database connection string ({type}://{username}:{password}@{host}/{db}) [$DATABASE_URL]
   -q value, --query value   SQL query as source for csv data [$QUERY]
   -o value, --output value  output file name (if not specified, stdout is used) [$OUTPUT]
```

You can also use pipes by omitting `-o` attribute:

```
sql-exporter dump -u postgres://username:password@hostname/database?sslmode=required -q "SELECT * FROM table" | grep "blah" > output.csv
```

## Docker usage

```
docker run --rm inloopx/sql-exporter sql-exporter dump -u postgres://username:password@hostname/database?sslmode=required -q "SELECT * FROM table" > output.csv
```
