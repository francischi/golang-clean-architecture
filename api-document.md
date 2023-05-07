# api document
*this is an example for api document

prefix : ```/api```

---

##　member

### <span style="color:blue">/member

#### **method** ```POST```

#### **params**
**body** :

Content-Type : ```application/json```

| param     | required | description |
| :-----    | :---     | :----       |
| name      | v        |             |
| email     | v        |  XXX@YYY    |
| gender    | v        | MALE/FEMALE |
| password  | v        | length > 6  |


```
{
    "name":"frank",
    "email":"frank@gmail.com",
    "gender":"MALE",
    "password":"thisisfrank"
}
```

#### **result**

* success
```
{
    "error_msg": null,
    "result": null,
    "success": true,
}
```

* failed
```
{
    "error_msg": "account_existed",
    "result": null,
    "success": false
}
```

---
### <span style="color:blue">/member/login
#### **method** ```POST```
#### **params**
**body** :

Content-Type : ```application/json```
| param     | required | description |
| :-----    | :---     | :----       |
| account   | v        |  XXX@YYY    |
| password  | v        |  length > 6 |

```
{
    "account" : "frank@gmail.com",
    "password" : "thisisfrank3"
}
```

#### **result**
* success
```
{
    "error_msg": null,
    "result": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJtZW1iZXJJZCI6ImQ2OWIxNmU5LTYxZWUtNDZkMS1hOTllLTg5NGI3NjBmYTVhMyIsIm5hbWUiOiJmcmFuayIsInRpbWVTdGFtcCI6MTY4Mjg0MDM4MX0.AApPh8StLgI2HNV0AoBwuKZhvaxkWs0Q4LEn3ouFtFU",
    "success": true
}
```

* failed
```
{
    "error_msg": "password_not_match",
    "result": null,
    "success": false
}
```
---
### <span style="color:blue">/member/password
#### **method** ```PATCH```
#### **params**
**body** :

Content-Type : ```application/json```
| param        | required | description |
| :-----       | :---     | :----       |
| memberId     | v        | uuid    |
| oldPassword  | v        |  |
| newPassword  | v        | length > 6  |
| confirmPassword  | v    | length > 6  |
```
{
    "memberId":"d69b16e9-61ee-46d1-a99e-894b760fa5a3",
    "oldPassword":"thisisfrank3",
    "newPassword":"thisisfrank4",
    "confirmPassword":"thisisfrank4"
}
```

#### **result**
* success
```
{
    "error_msg": nul,
    "result": null,
    "success": true
}
```
* failed
```
{
    "error_msg": "old_password_not_match",
    "result": null,
    "success": false
}
```