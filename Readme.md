# URL shortening service 

A simple project developing a URL shortening service, primarily backend, using Golang. 

## Motivation

To built a backend service using Go and it's frameworks, and to try out creating an containerised application off of that service to gain hands-on experience.

## Architecture 
![Picture](/images/Architecture.jpg)

## Tech stack used
- Go
- Redis (for database)
- Fiber (for web services)
- Docker (for containerisation)

## Requirements
- [Go](https://go.dev/dl/)
- [Redis](https://github.com/go-redis/redis) 
- [Fiber](https://github.com/gofiber/fiber) 
- [Docker](https://docs.docker.com/desktop/install/windows-install/) 


## Local setup

1. Clone the repository.

```bash
git clone https://github.com/PMohanJ/go-url-shortener.git
```

2. Make sure you've all the [Requirements](#requirements)/dependencies installed in your system.

3. Enter project directory.

```bash 
cd go-url-shortener
```

4. Now initiate the docker-compose file to start the application.
```bash
docker-compose up
``` 

## Usage

After the container is up and running, you can make requests to the application at following port `http://localhost:3000/`

- Send a request with following syntax:

```bash
curl localhost:3000/api/v1 -H "Content-Type":"application/json" \
-d '{"url": "https://www.youtube.com/watch?v=R_fZjGm2OrM"}'
```

Response for above request looks as below

```json
{
    "url":"https://www.youtube.com/watch?v=R_fZjGm2OrM",
    "short":"localhost:3000/dd2970",
    "expiry":24,
    "rate_limit":6,
    "rate_limit_reset":12
}
```

- The application does even support `custom url`, for example:

```bash 
curl localhost:3000/api/v1 -H "Content-Type":"application/json" \
-d '{"url": "https://www.youtube.com/watch?v=R_fZjGm2OrM", "short":"123@Mine"}'
```

Response to above command

```json
{
    "url":"https://www.youtube.com/watch?v=R_fZjGm2OrM",
    "short":"localhost:3000/123@Mine",
    "expiry":24,
    "rate_limit":5,
    "rate_limit_reset":8
}
```

- You can also set the `custom expiry time` for the URL, which actaully defaults to 24hr if not set:

```bash
curl localhost:3000/api/v1 -H "Content-Type":"application/json" \
 -d '{"url": "https://www.youtube.com/watch?v=R_fZjGm2OrM", "expiry":10}'
``` 

Response to above command

```json
{
    "url":"https://www.youtube.com/watch?v=R_fZjGm2OrM",
    "short":"localhost:3000/325800",
    "expiry":10,
    "rate_limit":4,
    "rate_limit_reset":6
}
```

So, whenever you use the shortened URL, you'll will be redirected to original URL.

