## REST

### POST

```shell
⟩ curl -X POST http://localhost:8088/v1/echo -d '{ "value": "hey" }'
```
```json
{"value":"hey"}
```

### GET

```shell
⟩ curl -X GET 'http://localhost:8088/v1/greet/hail?name=hodor'
```
```json
{"msg":"hail, hodor"}
```

## GraphQL

### Query

```shell
⟩ curl -X POST -H 'Content-Type:application/graphql' \
  http://localhost:3009/graphql \
  -d '{ Greet(type: "hail", name: "hodor") { msg } }'
```
```json
{"data":{"Greet":{"msg":"hail, hodor"}}}
```

### Mutate

NOTE: this just hangs

```shell
⟩ curl -X POST http://localhost:3009/graphql Content-Type:application/json -d \
'{
  "query": "mutation ($body: echosvcEchoMessageInput!) { Echo(body:$body) { value } }",
  "variables": { "body": { "value": "hey now" } }
}'
```
