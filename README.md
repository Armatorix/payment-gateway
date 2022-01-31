# payment-gateway

## NOTES

* codes for specific error triggers are not a good one, as the first one is Luhn correct, but the second and third are not.
* did not finish all endpoints as it requires quite a lot more craftwork - main concept and repo structure are here (except unit tests).

## DONE

* Fully working /Authorization endpoint
* openspec api specification base
* Pipeline for linter, build, tests
* Request fields validation
* Prometheus metrics
* Containerization
* E2E tests
* Whole structure for more craftwork

## TODO

* verify Luhn algorithm library
* add all endpoints to api spec, update api client, provide more e2e tests
* finish rest of the endpoints
* unit tests
* define currency format and add validation (ISO 4217 probably)

## NTH

* Cache for docker image building
* docker image with filesystem tracking for development - [air autoreload](https://github.com/cosmtrek/air)
* pipeline for e2e tests
* Postgres migration scripts
* System for api key/secret registration for merchants.

## Scaling

* Verify if [redis lock](https://github.com/bsm/redislock) works faster than pg serializable isolation level.
