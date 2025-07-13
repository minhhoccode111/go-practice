# Go Web Authentication and Authorization

Concepts:

- Router
- Middleware
- Pagination
- Filter
- Soft-delete
- Authorization
- Authentication
- Use `PATCH` verb
- Database indexes
- Database migration
- Interface
- Password hashing
- Input validation + sanitization
- Use pointer to differentiate between "no provided" and "explicitly false" in Go

Features:

APIs:

```txt
POST   /auth/register
POST   /auth/login
GET    /users/all
    - admin can get all accounts
GET    /users/{id}
PUT    /users/{id}
PATCH  /users/{id}/status
    - users can deactivate their account
    - only admin can activate an account
PATCH  /users/{id}/password
DELETE /users/{id}
    - admin can hard-delete a user account
```

Database:

```txt
- Users:
    - id
    - role
idx - email
    - password
    - is_active
idx - last_name
idx - first_name
    - created_at
    - updated_at
```
