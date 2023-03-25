# go-shorty

This is a simple HTTP service written in Go that provides the functionality of a URL shortener service. The service can be accessed at three URLs: /ping for checking its status, /set for setting a long URL and generating a shortened URL, and /l/[shortcode] for fetching and redirecting to the long URL associated with the given shortcode. The service can be easily deployed using Docker Compose and integrated with a Redis database for persistent storage.

### Developing
```bash
echo 'SERVICE_ADDR=' > .env
echo 'REDIS_ADDR=' > .env

make docker-up
```

### Using
```bash
curl -X POST -H 'Content-Type: application/json' -d '{"url": "https://www.google.com"}' 0.0.0.0:8081/set

# and go to `0.0.0.0:8081/l/ef7efc`
```
