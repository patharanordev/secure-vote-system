# **Secure Voting System**

## **Features**

- API Gateway
- Authentication System
- [Database](./docs/database.md)
- Voting APIs
- Web Application

## **Usage**

### **Start services**

The service, including :

- API Gateway
- Voting APIs
- Web application

```sh
docker-compose -f docker-compose.dev.yml down -v && \
docker-compose -f docker-compose.dev.yml up --build
```

### **Login**

Let's login first :

```sh
curl --location 'http://localhost:1323/login' \
--form 'username="jon"' \
--form 'password="shhh!"'
```

Response :

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA3Y2MzMmI1LTRlNzMtNGEyZS1iYWI1LThkZTM5OWE0MWRmNSIsIm5hbWUiOiJKb24gU25vdyIsImFkbWluIjp0cnVlLCJpc3MiOiJ0ZXN0Iiwic3ViIjoic29tZWJvZHkiLCJhdWQiOlsic29tZWJvZHlfZWxzZSJdLCJleHAiOjE2ODgyMDMwNTgsIm5iZiI6MTY4ODExNjY1OCwiaWF0IjoxNjg4MTE2NjU4LCJqdGkiOiIxIn0.8Tp7n6MBitnkczEi-KkbP0ZVTiXQKxbt1z-8CB4UVGo"
}
```

### **Calling APIs**

***Ex. Create user or Sign-up***

```sh
curl --location 'http://localhost:1323/signup' \
--header 'Content-Type: application/json' \
--data '{
    "username": "PatharaNor",
    "password": "1234567890",
    "isAdmin": true
}'
```

Response :

```text
Account created.
```

![create-user](assets/create-user.png)

***Ex. Voting list***

```sh
curl --location 'http://localhost:1323/api/v1/votes' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA3Y2MzMmI1LTRlNzMtNGEyZS1iYWI1LThkZTM5OWE0MWRmNSIsIm5hbWUiOiJKb24gU25vdyIsImFkbWluIjp0cnVlLCJpc3MiOiJ0ZXN0Iiwic3ViIjoic29tZWJvZHkiLCJhdWQiOlsic29tZWJvZHlfZWxzZSJdLCJleHAiOjE2ODgyMDMwNTgsIm5iZiI6MTY4ODExNjY1OCwiaWF0IjoxNjg4MTE2NjU4LCJqdGkiOiIxIn0.8Tp7n6MBitnkczEi-KkbP0ZVTiXQKxbt1z-8CB4UVGo'
```

Response :

- headers : `{ "server": "PatharaNor", "x-user-id": "07cc32b5-4e73-4a2e-bab5-8de399a41df5" }`
- body: `Vote list`
