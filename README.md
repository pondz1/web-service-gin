# web-service-gin

#### GET "/getuser" <br/>

```
output
  
 {
    "data": [
        {
            "id": "60fd52b0fbbd2bcefcb0da0c",
            "lineid": "asd2",
            "displayName": "Betty",
            "points": 115,
            "pictureUrl": "https://upload.wikimedia.org/wikipedia/commons/thumb/1/14/Gatto_europeo4.jpg/250px-Gatto_europeo4.jpg"
        }
    ]
}

```
#### GET "/userid/:id"  <br/>
```
output
{
    "data": {
        "id": "60fd52b0fbbd2bcefcb0da0c",
        "lineid": "asd2",
        "displayName": "Betty",
        "points": 115,
        "pictureUrl": "https://upload.wikimedia.org/wikipedia/commons/thumb/1/14/Gatto_europeo4.jpg/250px-Gatto_europeo4.jpg"
    },
    "message": "found"
}
```
#### GET "/getAllProducts" <br/>

```
output
{
    "data": [
        {
            "_id": "60fd5da694c11162837905bc",
            "p_point": 100,
            "p_name": "100"
        },
        {
            "_id": "60fd5dd11a2a265f9699271a",
            "p_point": 100,
            "p_name": "100"
        }
}
```
#### GET "/getProduct/:id" <br/>
```
out
{
    "data": {
        "_id": "60fd5e15829835390f46a4b2",
        "p_point": 100,
        "p_name": "p2"
    },
    "message": "found"
}
```
#### GET "/getRewardByUserId/:id"<br/>
```
out
{
    "data": [
        {
            "productName": "100",
            "usePoint": 100
        },
        {
            "productName": "100",
            "usePoint": 100
        },
        {
            "productName": "100",
            "usePoint": 100
        },
        {
            "productName": "100",
            "usePoint": 100
        }
    ],
    "message": "ok"
}
```
#### GET "/getReceipt" <br/>
```
out
{
    "data": [
        {
            "id": "60fe69d14b14c2b2784f2eb5",
            "status": "complete",
            "lineid": "asd3",
            "picturePath": "/image/359582.jpg"
        },
        {
            "id": "60fe6a204b14c2b2784f2eb6",
            "status": "pending",
            "lineid": "asd3",
            "picturePath": "/image/359582.jpg"
        }
    ],
    "message": "ok"
}
```
#### GET "/getReceipt/:id" <br/>
```
get Receipt By UserID
out
{
    "data": [
        {
            "id": "60fd8f24a5d7b0c8f23861f3",
            "status": "complete",
            "lineid": "asd2",
            "picturePath": "/image/359582.jpg"
        },
        {
            "id": "60fd8ff10c2348c376e1bdec",
            "status": "complete",
            "lineid": "asd2",
            "picturePath": "/image/359582.jpg"
        }
    ],
    "message": "ok"
}
```
#### POST "/insertUser" <br/>
```
input
{
    "lineid": "asd5",
    "name": "Betty",
    "points": 0,
    "pictureurl": "https://upload.wikimedia.org/wikipedia/commons/thumb/1/14/Gatto_europeo4.jpg/250px-Gatto_europeo4.jpg"
}
out
{
    "id": "60fe6af44b14c2b2784f2eb7",
    "message": "inserted"
}
```
#### POST "/updateUser" <br/>
```
input
{
    "lineid": "asd3",
    "name": "Betty3",
    "pictureUrl": "https://upload.wikimedia.org/wikipedia/commons/thumb/1/14/Gatto_europeo4.jpg/250px-Gatto_europeo4.jpg"
}


```
#### POST "/addUserPoint/:receiptID" <br/>
```
ex http://localhost:8080/addUserPoint/60fe69d14b14c2b2784f2eb5
receiptID for pending status only
input
{
    "lineid": "asd3",
    "points": 100
}
```
#### POST "/insertProduct" <br/>
```
input
{
    "p_point": 100,
    "p_name": "p2"
}

out
{
    "id": "60fe6e50549f8d43b440bea1",
    "message": "inserted!"
}
```
#### POST "/insertReward" <br/>
 ```
 (แลกของ)
 input
 {
    "r_lineid": "asd2",
    "r_product": {
        "_id": "60fd5da694c11162837905bc",
        "p_point": 100,
        "p_name": "100"
    }
}
 ```
#### POST "/uploadReceipt" <br/>
```
from-data
Key upload Type File
Key lineid Type Text
```
