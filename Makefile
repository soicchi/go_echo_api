.PHONY: go_fmt go_vet go_get go_tidy go_test ecr_push

path = ./cmd/dev/main.go
aws_region = ap-northeast-1

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

ecr_push:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${aws_account_id}.dkr.ecr.ap-northeast-1.amazonaws.com
	docker build --target prd -t ${aws_account_id}.dkr.ecr.${aws_region}.amazonaws.com/go-echo-api:latest .
	docker push ${aws_account_id}.dkr.ecr.ap-northeast-1.amazonaws.com/go-echo-api:latest
