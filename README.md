# go_sample_rest_api
Golang sample Rest API

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
