```
echo '{
    Greet(type: "hi", name: "bob") { msg }
}' | http POST :3009/graphql 'Content-Type:application/graphql'
```

NOTE: this just hangs

```
echo '{
  "query": "mutation ($body: echosvcEchoMessageInput!) { Echo(body:$body) { value } }",
  "variables": { "body": { "value": "hey now" } }
}' | http POST :3009/graphql
```
