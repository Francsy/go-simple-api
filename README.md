# go-simple-api

Playing around with Go. This repository is a first approach to learn the language structure and fundamentals. A fake database simple api with an endpoint and authentication middleware.

### Download go-simple-api:

1. Open the terminal.
2. Clone this repository the :
```
git clone https://github.com/Francsy/go-simple-api.git
```

### Execute go-simple-api:
1. Go to the project directory:
```
cd go-simple-api
```
2. Install dependencies: 
```
go mod tidy
```
3. Execute the api:
```
go run cmd/api/main.go 
```
4. You can make requests to `http://localhost:8000`:


## ENDPOINT:

### `/account/coins (GET)`

Introduce username as a query params and the Authorization key as a header

Example: `http://localhost:8000/coins?username=alex` (`123ABC` as Authorization header) will return a 200 code with the account balance.