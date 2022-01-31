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

.PHONY: compose
compose:
	docker-compose up --force-recreate -V