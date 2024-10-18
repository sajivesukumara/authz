# AUTHZ  - Authorization service
The service will allow the following
* Create user
* Login user and return jwt
* Validate jwt

## JSON Web Token 
JWTs consist of a header, a payload and a signature separated by three dots (e.g. xxx.yyy.zzz). Each of these parts is base64 encoded. Here is an example token:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkzMTU5MzcsImlkIjoxfQ._AvSTQRwBS8unmXXJ6r-vqt2rftEz6ZYPsKp9YibmBs
```

The first piece of data in the tokne is the header. 
```
{
  "alg": "HS256",
  "typ": "JWT"
}
```

The JWT header contains a key id (kid) that identifies te public key which can be use tovalidate teh signature.
The header also indicates which algorithm was used to sign the token.
The public key and algorithm are then used to verify the signature of the JWT to confirm that the token payload has not been modified. 

he second portion of the JWT is the payload which contains the claims. Decoding the second string above results in the following claims:
```
{
  "exp": 1729315937,
  "id": 1
  "name": "admin"
}

```
Teh last part of the token string is the encryption key used 


## User Experience

Public Api
Per the requirement above, the service will expose a REST API to provide information about the privileges a user has on resources. The main consumer of such API is the UI to enable/disable certain GUI controls based on the user privileges, and to restrict the user for displaying information that he has no access for. The endpoint supports two APIs:
Following endpoint are defined

POST   /auth/signup     --> Adds a user to the DB with hashed password
POST   /auth/login      --> Authenticates user and provides the JWT Token
GET    /user/validate   --> validate user JSON Web Token (JWT)
GET    /user/profile    --> Returns user profile using the jwt token
                            Provide Information on User Privileges
GET    /user/access-controls  --> Returns user profile using the jwt token
                            Provide Information on User Privileges

