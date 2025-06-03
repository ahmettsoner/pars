```
docker compose -f ./devops/opa/docker-compose.yaml  up -d
```

Build bundle

```
opa build -b devops/opa/bundle -o devops/opa/bundle/bundle.tar.gz
```
