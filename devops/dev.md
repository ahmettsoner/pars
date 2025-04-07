### Build image for dev

```
docker build -t ahmetsoner/gh-build-image:latest ./devops -f ./devops/Dockerfile.dev
```

### Build image for prod

```
docker build -t ahmetsoner/gh-build-image:latest ./devops -f ./devops/Dockerfile
```

dev tag

```
act-gh --job tag-dev
```

build from tag

```
act push -e <(echo '{"ref": "refs/tags/v1.0.0-dev.29"}')
```
