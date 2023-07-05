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
            "id": "851ec324-92f7-499b-8d6a-330a34e024aa",
            "itemDescription": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            "itemName": "item 1",
            "voteCount": 0
        }
    ],
    "error": null
}
```

## **Clear Vote List**

```sh
curl --location --request DELETE 'http://localhost:9323/api/v1/votes' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjNlMTE1MmE4LTc0YjUtNDUyMi05ZjE0LWE4MzAxMmUwMGFiOCIsIm5hbWUiOiJQYXRoYXJhTm9yIiwiYWRtaW4iOmZhbHNlLCJpc3MiOiJTZWNWb3RlU3lzIiwic3ViIjoiU2VjVm90ZVN5c19DdXN0b21BdXRoIiwiYXVkIjpbImdlbmVyYWxfdXNlciJdLCJleHAiOjE2ODg1NzIyMzgsIm5iZiI6MTY4ODQ4NTgzOCwiaWF0IjoxNjg4NDg1ODM4LCJqdGkiOiIxIn0.2zBydhR_x2fXchABOxGgcrAHzNBCYYmse_uocFzdvDc'
```

Response :

```json
{
    "status": 200,
    "data": "Vote list is cleared.",
    "error": null
}
```
