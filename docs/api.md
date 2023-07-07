# **Vote APIs**

## **Account**

***Step#1.1 : Create user or Sign-up***

```sh
curl --location 'http://localhost:9323/signup' \
--header 'Content-Type: application/json' \
--data '{
    "username": "PatharaNor",
    "password": "1234567890",
    "isAdmin": true
}'
```

Response :

```json
{
    "status": 201,
    "data": "Account created.",
    "error": null
}
```

![create-user](assets/create-user.png)

***Step#1.2 : Login***

```sh
curl --location 'http://localhost:9323/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "PatharaNor1",
    "password": "1234567890"
}'
```

Response :

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA3Y2MzMmI1LTRlNzMtNGEyZS1iYWI1LThkZTM5OWE0MWRmNSIsIm5hbWUiOiJKb24gU25vdyIsImFkbWluIjp0cnVlLCJpc3MiOiJ0ZXN0Iiwic3ViIjoic29tZWJvZHkiLCJhdWQiOlsic29tZWJvZHlfZWxzZSJdLCJleHAiOjE2ODgyMDMwNTgsIm5iZiI6MTY4ODExNjY1OCwiaWF0IjoxNjg4MTE2NjU4LCJqdGkiOiIxIn0.8Tp7n6MBitnkczEi-KkbP0ZVTiXQKxbt1z-8CB4UVGo"
}
```

***Optional#1.3 : Edit account***

```sh
curl --location 'http://localhost:9323/api/v1/account' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZiNGEyNzY5LTRmZjMtNGVjNi1iNTg2LTY2OThjM2ZlNGE5ZCIsIm5hbWUiOiJQYXRoYXJhTm9yIiwiYWRtaW4iOmZhbHNlLCJpc3MiOiJTZWNWb3RlU3lzIiwic3ViIjoiU2VjVm90ZVN5c19DdXN0b21BdXRoIiwiYXVkIjpbImdlbmVyYWxfdXNlciJdLCJleHAiOjE2ODg0ODI1OTYsIm5iZiI6MTY4ODM5NjE5NiwiaWF0IjoxNjg4Mzk2MTk2LCJqdGkiOiIxIn0.uxBd6HSPT2kKRY_u9QZ8GJt2ZtLHV3OpRjWEVFb9A0s' \
--header 'Content-Type: application/json' \
--data '{
    "username": "Bom Ja",
    "isAdmin": false
}'
```

Response :

```json
{
    "status": 200,
    "data": "Account updated.",
    "error": null
}
```

***Optional#1.4 : Delete account***

```sh
curl --location --request DELETE 'http://localhost:9323/api/v1/account' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImU3ZTRmMjc3LWEzMjctNGY5Ny04OGUzLTViMzlhMTAzZTAxOSIsIm5hbWUiOiJQYXRoYXJhTm9yMSIsImFkbWluIjpmYWxzZSwiaXNzIjoiU2VjVm90ZVN5cyIsInN1YiI6IlNlY1ZvdGVTeXNfQ3VzdG9tQXV0aCIsImF1ZCI6WyJnZW5lcmFsX3VzZXIiXSwiZXhwIjoxNjg4NDc4NTk2LCJuYmYiOjE2ODgzOTIxOTYsImlhdCI6MTY4ODM5MjE5NiwianRpIjoiMSJ9.kPhnHwj53XWdhoWTNacfU3tId9zu9rKxeZKeSv0E8bo'
```

Response :

```json
{
    "status": 200,
    "data": "Account deleted.",
    "error": null
}
```

## **CRUD Vote Item**

More detail [here](./api-vote-item.md).

## **Vote List**

More detail [here](./api-vote-list.md).
