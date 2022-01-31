# payment-gateway

## NOTES

* codes for specific error triggers are not good one, as first one is Luhn correct, but second and third are not.
* did not finished all endpoints as it requires quiet a lot more craft work - main concept and repo structure is here (except unit tests).

## DONE

* Fully working /Authorization endpoint
* openspec api specification base
* Pipeline for linter, build, tests
* Request fields validation
* Prometheus metrics
* Contenerization
* E2E tests
* Whole structure for more craft work

## TODO

* verify luhn algorithm library
* add all endpoints to api spec, update api client, provide more e2e tests
* finish rest of the endpoints
* unit tests
* define curreny format and add validation (ISO 4217 probably)

## NTH

* Cache for docker image building
* docker image with filesystem tracking for development - [air autoreload](https://github.com/cosmtrek/air)
* pipeline for e2e tests
* postgres migration scripts
* System for api key/secret registration for marchants.

## Scaling

* Verify if [redis lock](https://github.com/bsm/redislock) works faster than pg serializable isolation level.
