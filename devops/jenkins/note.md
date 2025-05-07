```
docker compose -f ./devops/jenkins/docker-compose.yaml  up -d
```

```
docker exec -it jenkins-master cat /var/jenkins_home/secrets/initialAdminPassword
```

```
docker compose -f ./devops/jenkins/docker-compose.yaml exec jenkins bash

```

install git scm plugin

run on "Script Console"

```

System.setProperty("hudson.plugins.git.GitSCM.ALLOW_LOCAL_CHECKOUT", "true")

```
