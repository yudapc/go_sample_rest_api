# go_sample_rest_api
Golang sample Rest API

How to run:
```
$ git clone git@github.com:yudapc/go_sample_rest_api.git
$ cd go_sample_rest_api
$ go run main.go

Go to browser http://localhost:3000/users
```

Resources:
```
GET /users
GET /users/:id
POST /users
PUT /users/:id
DELETE /users/:id
```

Sample Login:
```
POST /login

parameters:
 - email
 - password
```

Parameters (create user):
```
POST /users 
 - email
 - first_name
 - last_name
 - password
```

Parameters edit user:
```
PUT /users/:id
 - first_name
 - last_name
 - password
```
