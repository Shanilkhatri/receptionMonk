# ReceptionMonk

## Boilerplate Setup

### Before Starting

#### Two Factor Authentication
We need to introduce two factor authentication on the backend, This needs to work through JSON.
You can refer to online resources to understand but the basics are you generate server key and save it to the database, on the basis of that the QR the user verifies the code with the server and server serves the backup codes for download so user can login if they don't have access to the password.

Also to remember with the password reset the two factor keys should also reset. 

#### Header Auth Setup

We are authenticating based on headers, we need to cache the keys to redis so we can immediately authenticate to clients without much disk usage and have low latency.

The token needs to be generated on random seed and should depend on user's password, it needs to be hydrated at a given interval to ensure that if a user changes it's password the old token is invalidated and a new one takes its place.

#### CSRF Form Validator

CSRF is a simple mechanism where with each form render a random token is generated and the server validates the session based on that. You can read more about it here : https://brightsec.com/blog/csrf-token/

#### ORM Setup

We need to introduce ORM setup for golang, what does ORM mean ?
In very simple terms we need to have proper functions where we can quickly do functions like
find, or findById(id), or save(struct model{}) or update(struct model{})

This would help us to ease up in writing SQL code for simple requirements like fetching by ID or updating through the struct directly.
This can be achieved with the use of query builder functions to ease it up.
