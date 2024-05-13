# Readme.md

A project to learn webdevelopment with [GO](https://go.dev/). Following a course by [Jon Calhoun](https://www.calhoun.io/).

## Log
Just a note so I remember where I am. Progress is not stored on the course site.
- [x] 17.7 - Requiring a User via Middleware
- [x] 17.8 - Accessing the Current User in Templates
- [x] 17.9 - Request-Scoped Valuesew
- [x] 18.2 - SMTP Servicesls with SMTP
- [x] 18.5 - Building an email servicemailt DB Migration
- [x] 19.2 Password Reset Service Stubs
- [x] 19.3 Password Reset Service Stubs
- [x] 19.4 Forgot Password HTTP Handler
- [x] 19.5 Asynchronous Emails
- [x] 19.6 Forgot Password HTML Template
- [x] 19.7 Initializing Services with ENV Vars
- [x] 19.8 Check Your Email HTML Template
- [x] 19.9 Reset Password HTTP Handlers
- [x] 19.10 Reset Password HTML Template
- [x] 19.11 Update Password Function
- [x] 19.12 PasswordReset Creation
- [x] 19.13 Implementing Consume
- [ ] 19.15 Password Reset Exercises

## Exercise 19.15.1 Magic Link

- [x] Create a migration, e.g. a table for the magic links. This is nearly identical to the session password_resets table
- [x] Stub magic link creation and consumation (models/magic-link.go)
- [x] Add the magic link service to the users controller
- [ ] Alter the forgot password for and handler to include the magic link selection from the user. Default = false
