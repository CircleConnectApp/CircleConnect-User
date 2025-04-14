# User Microservice API Endpoints

## User Profile Endpoints

### Get Current User Profile
`GET /users/me`

**Description**: Get the profile of the authenticated user.

**Headers**:
- `Authorization: Bearer <token>`

**Responses**:
- `200 OK`: Returns the user profile
- `404 Not Found`: If the user is not found

---

### Update Current User Profile
`PUT /users/me`

**Description**: Update the profile of the authenticated user.

**Headers**:
- `Authorization: Bearer <token>`

**Request Body**:
```json
{
  "name": "string",
  "bio": "string"
}
```

**Responses**:
- `200 OK`: Returns the updated user profile
- `404 Not Found`: If the user is not found

## Get User Communities
`GET /users/me/communities`  

**Description**: Get the list of communities the authenticated user is part of (implementation coming soon).  

**Headers**:  
`Authorization: Bearer <token>`  

**Responses**:  
- `200 OK`: Returns a placeholder message  

---

## Admin Endpoints
*Note: These endpoints require admin privileges*

### Get User by ID
`GET /users/:id`  

**Description**: Get the profile of a user by their ID.  

**Headers**:  
`Authorization: Bearer <token>`  

**Responses**:  
- `200 OK`: Returns the requested user profile  
- `404 Not Found`: If the user is not found  
- `403 Forbidden`: If the requester is not an admin  

---

### Change User Role
`PUT /admin/users/:id/role`  

**Description**: Change the role of a user.  

**Headers**:  
`Authorization: Bearer <token>`  

**Request Body**:
```json
{
  "role": "string"
}
```
**Responses**:
- `200 OK`: Returns the updated user profile
- `404 Not Found`: If the user is not found
- `403 Forbidden`: If the requester is not an admin
