## JWT API
# API to manage user access.
Made with Go and GIN.


## FEATURES
- JWT authorization


## INSTALATION
1. Clone this repository: `git clone https://github.com/matheusjost/jwt-api.git`
2. Navigate to the project directory: `cd jwt-api.git`


## CONFIGURATION
1. Create a .env file on base dir:
```
PORT=http_port
sDB="host={dbhost} user={dbuser} password={dbpass} dbname={dbname} port={dbport} sslmode=disable"
SECRET_KEY=jwtscretkey
```
2. Install dependencies:
```
go mod tidy
```


## USAGE
1. Run main file: 
```
go run main.go
```
2. Access the API @ `http://localhost:9000`


## API ENDPOINTS
- POST `/signup`: Adds a user.
- POST `/login`: Login creating a JWT token, returns auth cookie.
- **REQUEST BODY**:
    - `login` (string, required): Your user login.
    - `password` (string, required): Your user password.

- GET `/validate`: Simple endpoint with a middleware validating your token.
- **REQUIRES AN AUTH COOKIE**:

## CONTRIBUTION
I'm still learning Golang basics so feel free to help me or give any advice!
You can fork the repository and make a pull request!

## ACKNOWLEDGMENTS
Special thanks to @codingwithrobby. This API was made inspired over his videos!