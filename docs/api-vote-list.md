# **Vote Item's APIs**

## **Get Vote List**

```sh
curl --location 'http://localhost:9323/api/v1/votes' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjNlMTE1MmE4LTc0YjUtNDUyMi05ZjE0LWE4MzAxMmUwMGFiOCIsIm5hbWUiOiJQYXRoYXJhTm9yIiwiYWRtaW4iOmZhbHNlLCJpc3MiOiJTZWNWb3RlU3lzIiwic3ViIjoiU2VjVm90ZVN5c19DdXN0b21BdXRoIiwiYXVkIjpbImdlbmVyYWxfdXNlciJdLCJleHAiOjE2ODg1NzIyMzgsIm5iZiI6MTY4ODQ4NTgzOCwiaWF0IjoxNjg4NDg1ODM4LCJqdGkiOiIxIn0.2zBydhR_x2fXchABOxGgcrAHzNBCYYmse_uocFzdvDc'
```

Response :

```json
{
    "status": 200,
    "data": [
        {
            "id": "e6087d49-23b3-443e-b12c-70dacc7c208c",
            "itemDescription": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            "itemName": "item 1",
            "userId": "173655f5-abcd-4125-908e-75cb4809d865",
            "voteCount": 0
        },
        {
            "id": "fb074920-7a0d-47ab-82fd-858e0e9b246d",
            "itemDescription": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            "itemName": "item 2",
            "userId": "173655f5-abcd-4125-908e-75cb4809d865",
            "voteCount": 0
        },
        {
            "id": "bdb06ef8-3cb9-489a-8e32-51b515cd2801",
            "itemDescription": "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
            "itemName": "item b",
            "userId": "9348b95a-7c51-4bcc-9564-c0614c6831a8",
            "voteCount": 0
        }
    ],
    "error": null
}
```

## **Clear Vote List**

> For admin ***ONLY***.

***In case : You have no permission***

```sh
curl --location --request DELETE 'http://localhost:9323/api/v1/votes' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjFmMmU3NzhlLThlNWItNGNjZC04ZGQ1LWExNjA2YjU5NmRiYyIsIm5hbWUiOiJQYXRoYXJhQWRtaW4iLCJhZG1pbiI6dHJ1ZSwiaXNzIjoiU2VjVm90ZVN5cyIsInN1YiI6IlNlY1ZvdGVTeXNfQ3VzdG9tQXV0aCIsImF1ZCI6WyJnZW5lcmFsX3VzZXIiXSwiZXhwIjoxNjg4OTc1NzQ2LCJuYmYiOjE2ODg4ODkzNDYsImlhdCI6MTY4ODg4OTM0NiwianRpIjoiMSJ9.1n-SvfqHT-8R69wH3xMXMUbYe0H1WRzqsIAsgykAUmU'
```

Response :

```json
{
    "status": 403,
    "data": null,
    "error": "Your role cannot access to the resource."
}
```

***In case : You are admin***

```sh
curl --location --request DELETE 'http://localhost:9323/api/v1/votes' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjFmMmU3NzhlLThlNWItNGNjZC04ZGQ1LWExNjA2YjU5NmRiYyIsIm5hbWUiOiJQYXRoYXJhQWRtaW4iLCJhZG1pbiI6dHJ1ZSwiaXNzIjoiU2VjVm90ZVN5cyIsInN1YiI6IlNlY1ZvdGVTeXNfQ3VzdG9tQXV0aCIsImF1ZCI6WyJnZW5lcmFsX3VzZXIiXSwiZXhwIjoxNjg4OTc1NzQ2LCJuYmYiOjE2ODg4ODkzNDYsImlhdCI6MTY4ODg4OTM0NiwianRpIjoiMSJ9.1n-SvfqHT-8R69wH3xMXMUbYe0H1WRzqsIAsgykAUmU'
```

Response :

```json
{
    "status": 200,
    "data": "Vote list is cleared.",
    "error": null
}
```
