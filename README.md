# openapi-go-test


```
npx openapi-generator generate -i agencyBaseCompiled.yml -g go-gin-server -o ./ --additional-properties=outputAsLibrary=true,interfaceOnly=true,apiPath=service/api,enumClassPrefix=I --git-host=github.com --git-repo-id=openapi-go-test --git-user-id=lisyaoran51 --model-package=model --api-package=api --package-name=api
```

```
npx openapi-generator help generate
npx openapi-generator config-help -g go-gin-server
```

```
go mod tidy
```

```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go get -u github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen 
go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen agencyBaseCompiled.yml
oapi-codegen -package api -generate gin agencyBaseCompiled.yml > api/api.go
oapi-codegen -package model -generate types agencyBaseCompiled.yml > model/model.go
oapi-codegen -package model --config=config.yml -generate gin,types agencyBaseCompiled.yml > api/api.go
oapi-codegen -package model --config=config.yml -generate server agencyBaseCompiled.yml > api/server.go
oapi-codegen -package api --config=config.yml agencyBaseCompiled.yml > api/api.go
oapi-codegen --config=model-config.yml agencyBaseCompiled.yml 
oapi-codegen --config=api-config.yml agencyBaseCompiled.yml 
oapi-codegen --config=client-config.yml agencyBaseCompiled.yml
oapi-codegen --config=client-model-config.yml agencyBaseCompiled.yml
```

```
export PATH=$PATH:$HOME/go/bin
```