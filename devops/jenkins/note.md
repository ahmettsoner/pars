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

Windows Jenkins Agent

```
ssh Administrator@192.168.122.118
```

```
curl.exe -sO http://192.168.122.1:8080/jnlpJars/agent.jar
```

```
$Action = New-ScheduledTaskAction -Execute "java.exe" -Argument '-jar "C:\Users\Administrator\jenkins\agent.jar" -url http://192.168.122.1:8080/ -secret 7d417c2e9f05e250bb734bdd8e13050ff2185bad27d9fb12395db5f9e8e98e22 -name windows -webSocket -workDir "C:\Users\Administrator\jenkins"'
$Trigger = New-ScheduledTaskTrigger -AtStartup
$Principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -RunLevel Highest
Register-ScheduledTask -TaskName "StartJenkinsAgent" -Action $Action -Trigger $Trigger -Principal $Principal
Start-ScheduledTask -TaskName "StartJenkinsAgent"
```

Ubuntu Jenkins Agent

```
ssh ahmettsoner@192.168.122.112
```

```
curl -sO http://192.168.122.1:8080/jnlpJars/agent.jar
```

```
sudo nano /etc/systemd/system/jenkins-agent.service

```

```
[Unit]
Description=Jenkins Agent
After=network.target

[Service]
User=ahmettsoner
WorkingDirectory=/home/ahmettsoner
ExecStart=/usr/bin/java -jar /home/ahmettsoner/agent.jar -url http://192.168.122.1:8080/ -secret 29a236ebcb6f755cfc9c96540f7a1cc8894b232890d709d6555b050533304bcc -name ubuntu -webSocket -workDir /home/ahmettsoner
Restart=always

[Install]
WantedBy=multi-user.target
```

```
sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl enable jenkins-agent.service
sudo systemctl start jenkins-agent.service
```
