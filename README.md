# Go events Rest API

This repo contains the code for a simple Rest API developed using:

- The go programming language.
- Gin as a web framework.
- The JWT package for token authentication.

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go 1.20 or higher

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/efive-dev/go-events-rest-api.git
   cd go-events-rest-api
   ```

1. Install the dependencies.
1. Clone the repository:

   ```bash
   go run .
   ```

## Endpoints

The API consists of various endpoints:

| type of request | endpoint                     | description                                                                                                                           |
| --------------- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **GET**         | `/events`                    | Gets all events currently in the database.                                                                                            |
| **GET**         | `/events/{event_id}`         | Gets the event with id == event_id.                                                                                                   |
| **POST**        | `/signup`                    | Registers a user with an email and a password.                                                                                        |
| **LOGIN**       | `/login`                     | Allows a previously registered user to login inside the API. It also returns a JWT token for authentication.                          |
| **POST**        | `/events`                    | Creates a new event. It requires authentication.                                                                                      |
| **PUT**         | `/events/{event_id}`         | Allows to change the fields of an event given its id. It requires authentication and also only the creator of an event may delete it. |
| **DELETE**      | `/events/{event_id}`         | Deletes the event given its id. It requires authentication and also only the creator of an event may delete it.                       |
| **POST**        | `events/{event_id}/register` | Allows a user to register for a certain event. It requires authentication.                                                            |
| **DELETE**      | `events/{event_id}/register` | Allows a user to delete registration for a certain event. It requires authentication.                                                 |

```

```
