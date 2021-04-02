# Banking API

Golang API that deals with common banking routines



## This API uses the following routes:

### accounts:

#### `POST/accounts` - create an account

#### `GET/accounts` - get a list of accounts

#### `GET/accounts/{account_id}` - get the balance of that specific account

### login

#### `POST/login` - Authenticate the user using CPF and SECRET

### transfers

#### `POST/transfers` - Makes a transfer

#### `GET/transfers` - Get the list of transfers made by current user

#### User has to be authenticated to use transfers

## How to run

### Go:

* Check if Go is installed: `go version`

* If it returns a version for Go, Go is installed.

* If Go is not installed, follow the [documentation](https://golang.org/doc/install) to install Go

### Docker:

* Check if Docker is present: `docker -v`

* If it returns a version for Docker, Docker is installed.

* If Docker is not installed, follow the [documentation](https://docs.docker.com/engine/install/ubuntu/) to install Docker

### Postgres:

* Now we need to get Postgres image into Docker, run: `docker pull postgres:latest`

* With the image that we just pulled we can create a container. There is a MakeFile that shortens the command for us.

* Run: `make postgres`

* After creating the container now we need to create the Database. Run: `make createdb` (dropdb to remove the database)

* To perform our migrations, run: `make migrateup`

* Finally, we can start the server by running: `make server`

#### The API should be running

## API client

Using an API client like [Insomnia](https://insomnia.rest/) we can test our API.

We need to configure our routes in Insomnia:

### POST/accounts

#### We need to specify the URL and the JSON:

* URL: `http://localhost:8080/accounts`

The JSON will have 3 fields:

* `"name":"firstname",`

* `"cpf":"123.456.789-10",`

* `"secret":"password"`

### GET/accounts

This one has 2 Queries for pagination: 

* page_id - number of the page (minimum = 1)

* page_size - number of elements per page (minimum = 1)

If we select page_id to be 1 and page_size to be 5, our URL would be:

* `http://localhost:8080/accounts/?page_id=1&page_size=5`

### GET/accounts/{account_id}

The only element of this is the URL, where you pass the account id. So lets say I want to retrieve information about account with ID = 2. My url would be
`http://localhost:8080/accounts/2`

### POST/login

* URL: `http://localhost:8080/login`

* The JSON for this will send the `cpf` and `secret`

* This will return a token to be used during transfers

### POST/transfers

* URL: `http://localhost:8080/transfers`

* The JSON will send the `account_origin_id`, `account_destination_id`, and the `amount`

* The header will be `Authorization` and the value will be `Bearer {token}`

* `{token}` should be the value received during `POST/login`

### GET/transfers

This has the same Query fields for pagination as `GET/accounts`

* URL: `http://localhost:8080/transfers?page_id=1&page_size=5`

* The header will be `Authorization` and the value will be `Bearer {token}`

* `{token}` should be the value received during `POST/login`
