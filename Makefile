
.PHONY: generate
generate: ## generate code from openapi spec
	docker run --rm \
  	-v ${PWD}:/local openapitools/openapi-generator-cli generate \
  	-i /local/open-api-spec.json \
		--skip-validate-spec \
  	-g go \
  	-o /local/client


.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
