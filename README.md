# Google Oauth 2 Golang
## Libraries
- [gin](https://github.com/gin-gonic/gin)

## Getting Started
### Run Server
Run server
```shell
go run main.go
```

## API
### Register
#### POST http://localhost:8000/api/auth/register
```shell
curl --location 'http://localhost:8000/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "nur@gmail.com",
    "name": "jon",
    "password": "jono12345",
    "passwordConfirm": "jono12345"
}'
```

### Login
#### POST http://localhost:8000/api/auth/login
```shell
curl --location 'http://localhost:8000/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "nur@gmail.com",
    "password": "jono12345"
}'
```

### Logout
#### GET http://localhost:8000/api/auth/logout
```shell
curl --location 'http://localhost:8000/api/auth/logout'
```

### Profile Google Me
#### GET http://localhost:8000/api/users/me
```shell
curl --location 'http://localhost:8000/api/users/me' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTg5NDY1NzUsImlhdCI6MTY5ODk0Mjk3NSwibmJmIjoxNjk4OTQyOTc1LCJzdWIiOiI5OTJiYTY1My0xOTBiLTQ0YTgtYTAzYi1lZjQ0ZGZkN2ZmNTkifQ.IVdntjW3WVn3TpRnZJhFgx7F7vyTWXKOxQlgUwcgeKI'
```

### Callback Google Oauth
#### Get http://localhost:8000/api/sessions/oauth/google
