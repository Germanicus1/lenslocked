# Readme.md

A project to learn webdevelopment with [GO](https://go.dev/). Following a course by [Jon Calhoun](https://www.calhoun.io/).

## Log

Just a note so I remember where I am, bacause progress is not stored on the course site.

### Section 17

- [x] 17.7 - Requiring a User via Middleware
- [x] 17.8 - Accessing the Current User in Templates
- [x] 17.9 - Request-Scoped Valuesew

### Section 18

- [x] 18.2 - SMTP Servicesls with SMTP
- [x] 18.5 - Building an email servicemailt DB Migration

### Section 19

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
- [x] 19.15 Password Reset Exercises

### Section 20

- [x] 20.1 Inspecting Errors
- [x] 20.2 Inspecting Wrapped Errors
- [x] 20.3 Designing the Alert Banner
- [x] 20.4 Dynamic Alerts
- [x] 20.5 Removing Alerts with JavaScript
- [x] 20.5 Detecting Existing Emails
- [x] 20.7 [Accepting Errors in Templates](https://courses.calhoun.io/lessons/les_wdv2_accept_errors_in_tpls)
- [x] 20.8 [Public vs Internal Errors](https://courses.calhoun.io/lessons/les_wdv2_pub_vs_int_errs)
- [x] 20.9 [Creating Public Errors](https://courses.calhoun.io/lessons/les_wdv2_create_pub_errs)
- [x] 20.10 [Using Public Errors](https://courses.calhoun.io/lessons/les_wdv2_using_pub_errs)
- [x] 20.11 [Better Error Handling Exercises](https://courses.calhoun.io/lessons/les_wdv2_better_err_exercises)

### Section 21

- [x] [Galleries Overview](https://courses.calhoun.io/lessons/les_wdv2_galleries_overview)
- [x] [Gallery Model and Migration](https://courses.calhoun.io/lessons/les_wdv2_gallery_model_and_migration)
- [ ] [Creating Gallery Records](https://courses.calhoun.io/lessons/les_wdv2_create_gallery_records)
- [ ] Querying for Galleries by ID
- [ ] Querying Galleries by UserID
- [ ] Updating Gallery Records
- [ ] Deleting Gallery Records
- [ ] New Gallery Handler
- [ ] views.Template Name Bug
- [ ] New Gallery Template
- [ ] Gallery Routing and CSRF Bug Fixes
- [ ] Create Gallery Handler
- [ ] Edit Gallery Handler
- [ ] Edit Gallery Template`
- [ ] Update Gallery Handler
- [ ] Gallery Index Handler
- [ ] Discovering and Fixing a Gallery Index Bug
- [ ] Gallery Index Template Continued
- [ ] Show Gallery Handler
- [ ] Show Gallery Template and a Tailwind Update
- [ ] Extracting Common Gallery Code
- [ ] Extra Gallery Checks with Functional Options
- [ ] Delete Gallery Handler
- [ ] Gallery Exercises


## Gallery functionallity
- [ ] Create a new gallary with title
- [ ] Upload images to a gallary
- [ ] Delete images from a gallary
- [ ] Update the title of the gallary
- [ ] View the gallary so we can share it with others
- [ ] Delete a gallary
- [ ] View a list of gallaries we are allowed to edit

### Views

This might give us the following views to create:

| Functionality | Page |
|---|---|
Create a new gallary| new
Edit a gallary| edit
View a gallary|show
View a list of all gallaries|index

### Controllers

We will also need controller (aka http handlers) to support those views.

Controller|Description
|-|-|
New and Create|Render and process a new-gallery form
Edit and Update|Render and process a form to edit a gallary
Show|Render a gallary
Delete|Delete a gallary

### Models

We need to persist the data in our models package. To this we need to support the following:

- Creating a gallery
- Updating a gallery
- Querying for a gallery by ID
- Querying for all galleries with a user ID
- Deleting a gallery

**For later in a separate handler:**
- Creating an image for a gallery
- Deleting an image from a gallery


### Data structures

**Galleries will need:**
- ID
- Title
- UserID (owner)
- Associateds images

