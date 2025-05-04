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
act-gh push --job build-dev -e <(echo '{"ref": "refs/tags/v1.0.0-dev.40"}')
```

<!-- enc -->
```
act-gh --job translate --eventpath .notes/translate-event.json  -s T1BFTkFJX0FQSV9LRVk9c2stcHJvai1sV3lXNzY4ckhqNlZ5anR4OGt4U2hUNFdOM2VCN2QtelJJUDFvdGc3emE0Vk1lYU5jS0QxdzZSd2ItZG91NG5hRjlMT1lVd2U4OVQzQmxia0ZKSzV2eXBkOFByYnV6RmRYRWpTYk90eHNvVDd3SllpNG1IU0NsZThVTmhNVHRuQmxsamFMUS12R0xvQjRfUkNaV0N5bzdxLVoya0E=
```

```
Z2hwX21TRkVCVmRqMWRmZFhCZElFdWJ3MWtYUk9WanVFWTFuRnVtSA==
```

```
act-gh push -s OPENAI_API_KEY=$(cat .secrets | grep OPENAI_API_KEY | cut -d '=' -f2) -e <(echo '{"ref": "refs/heads/dev"}')
```

```
act-gh push --job build-dev-deb-package -e <(echo '{"ref": "refs/tags/v1.0.0-dev.40"}')
```
