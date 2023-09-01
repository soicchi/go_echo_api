.PHONY: go_fmt go_vet go_get go_tidy go_test

path = ./...

go_fmt:
	docker compose run --rm  api go fmt ${path}

go_vet:
	docker compose run --rm api go vet ${path}

go_get:
	docker compose run --rm api go get ${pkg}

go_tidy:
	docker compose run --rm api go mod tidy

go_test:
	docker compose run --rm api go test -v -cover ${path}
