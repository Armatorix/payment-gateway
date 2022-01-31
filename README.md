# payment-gateway

# TODO:
* verify luhn algorithm library
* add all endpoints to api spec, update api client, provide more e2e tests
* finish rest of the endpoints
* define curreny format and add validation (ISO 4217 probably)

# NTH:
* System for api key/secret registration for marchants.
* Cache for docker image building
* postgres migration scripts
* docker image with filesystem tracking for development (https://github.com/cosmtrek/air)
* pipeline e2e

# Scaling
. Verify if [redis lock](https://github.com/bsm/redislock) works faster than pg serializable isolation level.

