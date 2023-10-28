# golang-gorm-jwt

### User Model

```
type User struct {
	ID           uint           `gorm:"primary_key;autoIncrement" json:"id"`
	UserID       string         `json:"user_id"`
	FirstName    string         `gorm:"not null" json:"first_name"`
	LastName     string         `gorm:"not null" json:"last_name"`
	Email        string         `gorm:"unique, not null" json:"email"`
	Password     string         `gorm:"not null" json:"password"`
	UserType     string         `gorm:"default:'user'" json:"user_type"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
```

### API Endpoints

##### 1. Signup

> Endpoint: /signup
> Method: POST
> Request Payload:

```
{
   "first_name": "A2",
   "last_name": "G2",
   "email": "e2@gmail.com",
   "password": "user_password"
   "user_type": "admin/user"
}
```

> Response:

```
{
  "id": 2,
  "user_id": "VV3EC2Tnrynz8Eb9",
  "first_name": "A2",
  "last_name": "G2",
  "email": "e2@gmail.com",
  "password": "$2a$14$BdCpUYbT23rrfUxmz4gHUeB179L1KQVp5u0uJHM3CE3E1klFuqW82",
  "user_type": "user",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiJBMiIsIkxhc3ROYW1lIjoiRzIiLCJFbWFpbCI6ImUyQGdtYWlsLmNvbSIsIlVzZXJJRCI6IlZWM0VDMlRucnluejhFYjkiLCJVc2VyVHlwZSI6IiIsImV4cCI6MTY5ODU2NDQzOH0.t3RZ-lQIUGiIsVazS0jMqlAiC5UjV8YXydmeVWLcqvY",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiIiLCJMYXN0TmFtZSI6IiIsIkVtYWlsIjoiIiwiVXNlcklEIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTkwODI4Mzh9.oRF1FMbPdYExuAFed9SjZKMe9q1cNz-KEZ2aKFubBGg",
  "created_at": "2023-10-28T12:57:18+05:30",
  "updated_at": "2023-10-28T12:57:18+05:30",
  "deleted_at": null
}
```

<br>
<br>

##### 2. Login

> Endpoint: /login
> Method: POST
> Request Payload:

```
{
   "email": "e2@gmail.com",
   "password": "user_password"
}
```

> Response:

```
{
  "id": 2,
  "user_id": "VV3EC2Tnrynz8Eb9",
  "first_name": "A2",
  "last_name": "G2",
  "email": "e2@gmail.com",
  "password": "$2a$14$BdCpUYbT23rrfUxmz4gHUeB179L1KQVp5u0uJHM3CE3E1klFuqW82",
  "user_type": "user",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiJBMiIsIkxhc3ROYW1lIjoiRzIiLCJFbWFpbCI6ImUyQGdtYWlsLmNvbSIsIlVzZXJJRCI6IlZWM0VDMlRucnluejhFYjkiLCJVc2VyVHlwZSI6IiIsImV4cCI6MTY5ODU2NDQzOH0.t3RZ-lQIUGiIsVazS0jMqlAiC5UjV8YXydmeVWLcqvY",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiIiLCJMYXN0TmFtZSI6IiIsIkVtYWlsIjoiIiwiVXNlcklEIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTkwODI4Mzh9.oRF1FMbPdYExuAFed9SjZKMe9q1cNz-KEZ2aKFubBGg",
  "created_at": "2023-10-28T12:57:18+05:30",
  "updated_at": "2023-10-28T12:57:18+05:30",
  "deleted_at": null
}
```

<br>
<br>

##### 3. Refresh Token

> Endpoint: /refresh/token
> Method: POST
> Request Payload: Add refresh_token in header

> Response:

```
{
  "id": 2,
  "user_id": "VV3EC2Tnrynz8Eb9",
  "first_name": "A2",
  "last_name": "G2",
  "email": "e2@gmail.com",
  "password": "$2a$14$BdCpUYbT23rrfUxmz4gHUeB179L1KQVp5u0uJHM3CE3E1klFuqW82",
  "user_type": "user",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiJBMiIsIkxhc3ROYW1lIjoiRzIiLCJFbWFpbCI6ImUyQGdtYWlsLmNvbSIsIlVzZXJJRCI6IlZWM0VDMlRucnluejhFYjkiLCJVc2VyVHlwZSI6IiIsImV4cCI6MTY5ODU2NDQzOH0.t3RZ-lQIUGiIsVazS0jMqlAiC5UjV8YXydmeVWLcqvY",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiIiLCJMYXN0TmFtZSI6IiIsIkVtYWlsIjoiIiwiVXNlcklEIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTkwODI4Mzh9.oRF1FMbPdYExuAFed9SjZKMe9q1cNz-KEZ2aKFubBGg",
  "created_at": "2023-10-28T12:57:18+05:30",
  "updated_at": "2023-10-28T12:57:18+05:30",
  "deleted_at": null
}
```

<br>
<br>

##### 4. Get Users

> Endpoint: /users/:user_id(self user id)
> Method: POST
> Request Payload: Add refresh_token in header

> Response:

```
[
  {
    "id": 1,
    "user_id": "cjEmOMVVn0Tqd35m",
    "first_name": "A1",
    "last_name": "G1",
    "email": "e1@gmail.com",
    "password": "$2a$14$jZM6KCJ7Vqmhq5Lw8ucHye7C43Zxn1S545I9.o3lbOcV4NNJ9Fu2G",
    "user_type": "admin",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiJBMSIsIkxhc3ROYW1lIjoiRzEiLCJFbWFpbCI6ImUxQGdtYWlsLmNvbSIsIlVzZXJJRCI6ImNqRW1PTVZWbjBUcWQzNW0iLCJVc2VyVHlwZSI6ImFkbWluIiwiZXhwIjoxNjk4OTI2NjAzfQ._ieO4IemgbeEsauMh3lkh6aZI8CtkrOC0D1UE8f1FXk",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiIiLCJMYXN0TmFtZSI6IiIsIkVtYWlsIjoiIiwiVXNlcklEIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTkwODUwMDN9.5o0zERPP0JBaNLwT0cBZ_vh2Ck2yXo5sNRss4Gzriq8",
    "created_at": "2023-10-28T12:51:02+05:30",
    "updated_at": "2023-10-28T13:33:23.606+05:30",
    "deleted_at": null
  },
  {
    "id": 2,
    "user_id": "VV3EC2Tnrynz8Eb9",
    "first_name": "A2",
    "last_name": "G2",
    "email": "e2@gmail.com",
    "password": "$2a$14$BdCpUYbT23rrfUxmz4gHUeB179L1KQVp5u0uJHM3CE3E1klFuqW82",
    "user_type": "user",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiJBMiIsIkxhc3ROYW1lIjoiRzIiLCJFbWFpbCI6ImUyQGdtYWlsLmNvbSIsIlVzZXJJRCI6IlZWM0VDMlRucnluejhFYjkiLCJVc2VyVHlwZSI6InVzZXIiLCJleHAiOjE2OTg0NzkyNjF9.fe-BRbizyhKOwAOgzoctpSYwMUp1mO9RPOscZhorg8w",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiIiLCJMYXN0TmFtZSI6IiIsIkVtYWlsIjoiIiwiVXNlcklEIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTkwODM5Mzd9.b2L3C5Qw7rdXidCGPtPdKggEstwn3EUoKNAJwGntjlQ",
    "created_at": "2023-10-28T12:57:18+05:30",
    "updated_at": "2023-10-28T13:15:37.452+05:30",
    "deleted_at": null
  }
]
```

<br>
<br>

##### 4. Get User

> Endpoint: /user/:user_id
> Method: POST
> Request Payload: Add refresh_token in header

> Response:

```
{
    "id": 1,
    "user_id": "cjEmOMVVn0Tqd35m",
    "first_name": "A1",
    "last_name": "G1",
    "email": "e1@gmail.com",
    "password": "$2a$14$jZM6KCJ7Vqmhq5Lw8ucHye7C43Zxn1S545I9.o3lbOcV4NNJ9Fu2G",
    "user_type": "admin",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiJBMSIsIkxhc3ROYW1lIjoiRzEiLCJFbWFpbCI6ImUxQGdtYWlsLmNvbSIsIlVzZXJJRCI6ImNqRW1PTVZWbjBUcWQzNW0iLCJVc2VyVHlwZSI6ImFkbWluIiwiZXhwIjoxNjk4OTI2NjAzfQ._ieO4IemgbeEsauMh3lkh6aZI8CtkrOC0D1UE8f1FXk",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGaXJzdE5hbWUiOiIiLCJMYXN0TmFtZSI6IiIsIkVtYWlsIjoiIiwiVXNlcklEIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTkwODUwMDN9.5o0zERPP0JBaNLwT0cBZ_vh2Ck2yXo5sNRss4Gzriq8",
    "created_at": "2023-10-28T12:51:02+05:30",
    "updated_at": "2023-10-28T13:33:23.606+05:30",
    "deleted_at": null
  }
```
