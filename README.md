# openapi-go-test


```
npx openapi-generator generate -i agencyBaseCompiled.yml -g go-gin-server -o ./ --additional-properties=outputAsLibrary=true,interfaceOnly=true,apiPath=service/internal/api --git-host=github.com --git-repo-id=openapi-go-test --git-user-id=lisyaoran51 --model-package=model --api-package=api
```

```
npx openapi-generator help generate
npx openapi-generator config-help -g go-gin-server
```

```
go mod tidy
```