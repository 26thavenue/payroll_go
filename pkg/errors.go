package pkg

import "errors"

var (
    ErrNoUserInContext = errors.New("no user found in context")
    ErrInvalidTokenClaims = errors.New("invalid token claims")
    ErrMissingIDClaim = errors.New("missing id claim in token")
    ErrMissingRoleClaim = errors.New("missing role claim in token")
    ErrInvalidIDType = errors.New("invalid id type in token")
    ErrInvalidRoleType = errors.New("invalid role type in token")
)

var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrUsernameTaken = errors.New("username already taken")
    ErrInvalidPassword = errors.New("invalid password")
)