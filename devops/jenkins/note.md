```
docker exec -it jenkins-master cat /var/jenkins_home/secrets/initialAdminPassword
```

install git scm plugin

run on "Script Console"

```
System.setProperty("hudson.plugins.git.GitSCM.ALLOW_LOCAL_CHECKOUT", "true")
```
