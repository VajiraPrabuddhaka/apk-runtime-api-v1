# apk-runtime-api-v1

Use the following command to generate the server code using openAPI spec.

This used [oapi-codegen](https://github.com/deepmap/oapi-codegen) tool to generate code.

Server code

```
oapi-codegen -generate chi-server -old-config-style -o openapi_server.gen.go -package server oas.yaml
```

Types

```
oapi-codegen -generate types -old-config-style -o openapi_types.gen.go -package server oas.yaml
```
