# ratelimiter4e -- Rate Limiter for Education

## Setup
1. `git clone https://github.com/shhivam/ratelimiting4e.git`
2. `cd ratelimiting4e`
3. `make db-up` will run the redis in a docker container
4. `make run` will run the Gin server on port 8080

## Test the rate limiting

This project has a `GET /healthcheck` endpoint which is rate limited 10 requests per minute but you can easily change it in `ratelimiter/healthcheck.go` file.
For the purpose of this exercise, I just hardcoded the rate limiting values. 

The requests will return a 429 Status code if you send more than 10 requests in a minute to `GET /healthcheck`

