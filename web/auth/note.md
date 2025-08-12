# Go Web Authentication and Authorization

Concepts:

- Router
- Filter
- Interface
- Middleware
- Pagination
- Soft-delete
- Authorization
- Authentication
- Use `PATCH` verb
- Database indexes
- Password hashing
- Database migration
- Input validation + sanitization
- Context for database query and middleware
- Concurrency for independent database query (get all users, count users)
- Use pointer to differentiate between "no provided" and "explicitly false" in Go

Features:

- Basic authentication and authorization

APIs:

```txt
       # public
POST   /auth/register
POST   /auth/login
GET    /users/all
       - filter email, status
       - pagination
       - TODO: only admin can get all users including inactive
GET    /users/{id}

       # auth
PATCH  /users/{id}
       - user update their profile, like email, bio, etc.
PATCH  /users/{id}/password
       - user update their password, must provide old password
PATCH  /users/{id}/status
       - user deactivate her account (soft-delete)
       - only admin can activate an account

       # authz
DELETE /users/{id}
       - admin hard-delete an account
PATCH  /users/{id}/role
       - admin make a user an admin
```

Database:

```txt
- Users:
    - id
    - role
idx - email     - unique
    - password
    - is_active
```

Todo (small details):

- use more context in middlewares and database query execution
