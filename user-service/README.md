# Requirements

## Functional
- User authentication
- Password reset
- Use of Redis for caching/token storage

## Non-functional
- DB access time should be less than 3 seconds
- Disable sign in, if there is more than 2 invalid attempts when signing in
- The system should be able to handle a few thousand users
