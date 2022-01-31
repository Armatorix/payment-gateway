OPENAPI_PATH=./api/openapi-spec/api.yaml
GENERATE_PATH=./api-client

.PHONY: model-rebuild
model-rebuild: model-remove model-generate model-clean format

.PHONY: model-generate
model-generate:
	docker run --rm \
			-u "$$(id -u)" \
  			-v ${PWD}:/local openapitools/openapi-generator-cli:v5.4.0 \
		generate \
			-i /local/${OPENAPI_PATH} \
			-g go \
			-o /local/${GENERATE_PATH} \
			--package-name model

.PHONY: model-remove
model-remove:
	-rm -rf ${GENERATE_PATH}

.PHONY: model-clean
model-clean:
	cd ${GENERATE_PATH} && \
	rm -rf .openapi-generator \
		api \
		docs 

	cd ${GENERATE_PATH} && \
	rm -f git_push.sh \
		go.mod \
		go.sum \
		README.md \
		.travis.yml \
		.openapi-generator-ignore

.PHONY: formatter
format:
	gofumpt -w ./..

.PHONY: run
run:
	docker-compose up --build -d --always-recreate-deps --renew-anon-volumes
	bash ./scripts/wait_for_it.sh
.PHONY: stop
stop:
	docker-compose down

.PHONY: test-e2e
test-e2e: 
	$(MAKE) run
	go test --count=1 ./e2e
	$(MAKE) stop