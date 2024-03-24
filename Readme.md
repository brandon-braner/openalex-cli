# OpenAlex Go Library

This code is a Go library for working with OpenAlex Snapshot data.

There is a current Python library for this but it can be slow. My hope is that this library will be faster.

## Installation
- Downlaod the OpenAlex snapshot following the instructions [here](https://docs.openalex.org/download-all-data/openalex-snapshot).
- Make sure you have Go installed on your machine. You can download it [here](https://go.dev/).

- TBD

## Usage

### Flattening Files for relational databases

Following the OpenAlex guide for importing data into a relational database, you can flatten the data using the following
you will need to flatten the files into a csv to import into a relational database. They instruct you that this can be slow
with their Python script and to try to paralize it. This library does that for you with Go.

Their instructions can be found [here](https://docs.openalex.org/download-all-data/upload-to-your-database/load-to-a-relational-database#step-2-convert-the-json-lines-files-to-csv)




# Commands

## Get a list of valid filenames to pass to other commands
```bash
openalex-cli get-valid-filenames
```

## Download Snapshot Data

## Database

### Create a database and the schemas

We currently only support Postgres databases. You can set one up at any of the cloud providers such as AWS, Google Cloud, or Azure. 
You can also run one locally using [Docker](https://hub.docker.com/_/postgres).

The SQL to create the database tables can be found in a [Github Repo](https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/create_table.sql).

Create an enviromental variable named OPEN_ALEX_CLI_DB for the database connection string. The format is `postgres://user:password@host:port/dbname`.

On a Mac or Linux machine you can set this in your shell by running the following command (replace the values with your own):
```bash
export OPEN_ALEX_CLI_DB="postgres://user:password@host:port/dbname"
```

On a Windows machine you can set this in your shell by running the following command:
```bash
set OPEN_ALEX_CLI_DB="postgres://user:password@host:port/dbname"
```







































