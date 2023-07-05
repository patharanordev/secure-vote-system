# **Vote Item's APIs**

## **Add Vote Item**

```sh
curl --location 'http://localhost:9323/api/v1/vote-item' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjNlMTE1MmE4LTc0YjUtNDUyMi05ZjE0LWE4MzAxMmUwMGFiOCIsIm5hbWUiOiJQYXRoYXJhTm9yIiwiYWRtaW4iOmZhbHNlLCJpc3MiOiJTZWNWb3RlU3lzIiwic3ViIjoiU2VjVm90ZVN5c19DdXN0b21BdXRoIiwiYXVkIjpbImdlbmVyYWxfdXNlciJdLCJleHAiOjE2ODg1NzIyMzgsIm5iZiI6MTY4ODQ4NTgzOCwiaWF0IjoxNjg4NDg1ODM4LCJqdGkiOiIxIn0.2zBydhR_x2fXchABOxGgcrAHzNBCYYmse_uocFzdvDc' \
--header 'Content-Type: application/json' \
--data '{
    "itemName": "item 2",
    "itemDescription": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}'
```

Response :

```json
{
    "status": 201,
    "data": "Vote item created.",
    "error": null
}
```

## **Update Vote Item**

```sh
curl --location --request PATCH 'http://localhost:9323/api/v1/vote-item' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjNlMTE1MmE4LTc0YjUtNDUyMi05ZjE0LWE4MzAxMmUwMGFiOCIsIm5hbWUiOiJQYXRoYXJhTm9yIiwiYWRtaW4iOmZhbHNlLCJpc3MiOiJTZWNWb3RlU3lzIiwic3ViIjoiU2VjVm90ZVN5c19DdXN0b21BdXRoIiwiYXVkIjpbImdlbmVyYWxfdXNlciJdLCJleHAiOjE2ODg1NzIyMzgsIm5iZiI6MTY4ODQ4NTgzOCwiaWF0IjoxNjg4NDg1ODM4LCJqdGkiOiIxIn0.2zBydhR_x2fXchABOxGgcrAHzNBCYYmse_uocFzdvDc' \
--header 'Content-Type: application/json' \
--data '{
    "id": "d2ffde25-b36a-47b3-8084-19e197b87ce0",
    "itemName": "item 2",
    "itemDescription": "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
    "voteCount": 2
}'
```

Response :

```json
{
    "status": 200,
    "data": "Vote item updated.",
    "error": null
}
```

## **Delete Vote Item**

```sh
curl --location --request DELETE 'http://localhost:9323/api/v1/vote-item' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjNlMTE1MmE4LTc0YjUtNDUyMi05ZjE0LWE4MzAxMmUwMGFiOCIsIm5hbWUiOiJQYXRoYXJhTm9yIiwiYWRtaW4iOmZhbHNlLCJpc3MiOiJTZWNWb3RlU3lzIiwic3ViIjoiU2VjVm90ZVN5c19DdXN0b21BdXRoIiwiYXVkIjpbImdlbmVyYWxfdXNlciJdLCJleHAiOjE2ODg1NzIyMzgsIm5iZiI6MTY4ODQ4NTgzOCwiaWF0IjoxNjg4NDg1ODM4LCJqdGkiOiIxIn0.2zBydhR_x2fXchABOxGgcrAHzNBCYYmse_uocFzdvDc' \
--header 'Content-Type: application/json' \
--data '{
    "id": "d2ffde25-b36a-47b3-8084-19e197b87ce0"
}'
```

Response :

```json
{
    "status": 200,
    "data": "Vote item deleted.",
    "error": null
}
```
