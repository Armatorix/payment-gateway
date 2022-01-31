# payment-gateway

# TODO:
* verify luhn algorithm library
* add all endpoints to api spec, update api client, provide more e2e tests
* change unauthorize/internal server error to unifed model 
* finish rest of the endpoints
* define curreny format and add validation (ISO 4217 probably)

# NTH:
* System for api key/secret registration for marchants.
* postgres migration scripts
* docker image with filesystem tracking for development (https://github.com/cosmtrek/air)

# Scaling
. Verify if [redis lock](https://github.com/bsm/redislock) works faster than pg serializable isolation level.

