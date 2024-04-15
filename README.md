# ShortLink - URL Shortening Service

## Overview

ShortLink is a simple URL shortening service that allows users to create shorter aliases for long URLs. It provides an easy way to share links while keeping them concise.

## Features

- **Shortening:** Generate short aliases for long URLs.
- **Redirection:** Redirect users from short aliases to the original long URLs.
- **Customization:** Optionally customize short aliases for easier sharing.
- **Analytics:** Track usage statistics such as click counts for each short alias.Also manage your all URLs.

## Usage

### Creating a Short Link

To create a short link, simply submit a long URL to shortLink's API endpoint. The service will return a shortened alias that you can use to redirect users.

Example API Request:

- POST api/url/create-short-url

```json
{
    "original_url":"https://github.com/gofiber/recipes/blob/master/auth-docker-postgres-jwt/handler/user.go",
    "user_id": "6d76dea9-3c08-4b27-a328-f337b4089d24",
    "expiration_date": "2024-04-16T13:45:20Z"
}
```

**Example Response:**
```json
{
    "status":  "success",
	"message": "shortUrl created successfully",
	"data": {"short_url": "http://{your-domain}/abc123"}
}
```

### Get All Your URLs

To get all your short link, simply send request to this API endpoint. This service will return a all the URLs you had created.

Example API Request:

- GET api/url/get-urls

**Example Response:**
```json
[
    {
    "data": [
        {
        "id": "f2300fe7-0bbe-4605-974e-e8e17a5adf20",
        "original_url": "{your-initial-URL}",
        "short_url": "{your-domain.com}/abc123",
        "visiter": 231
        }
    ],
    "message": "user URL's",
    "status": "success"
    },
    {
    "data": [
        {
        "id": "cs5lofe7-kc05-7349-974e-0eae7afddf20",
        "original_url": "{your-initial-URL}",
        "short_url": "{your-domain.com}/abc124",
        "visiter": 435
        }
    ],
    "message": "user URL's",
    "status": "success"
    }
]
```

>   - NOTE: You need to first login to use above service API endpoints and provide Autherisation token to header.
> - You recieved token from login or signup API endpoint in system
 ```
    Authorisation Bearer {token}
 ```

## Installation

To install shortLink on your own server, follow these steps:

1. Clone this repository.
2. Install redis and postgres on your system and run both database.
3. Configure the environment variables for your database connection, API keys, etc.
4. Run the application using `go run main.go`.

## Technologies Used

- **Web Framework:** [GoFiber](https://github.com/gofiber/fiber)
- **Database:** PostgreSQL for user and link storage.
- **Cache:** Redis for managing short URL.


# Todo List

- [ ] Customisation of link
- [ ] Gorm Error Handling
- Continue
