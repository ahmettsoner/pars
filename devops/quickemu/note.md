https://github.com/quickemu-project/quickemu/wiki/01-Installation

Debian

```
sudo apt-get install bash coreutils curl genisoimage grep jq mesa-utils ovmf pciutils procps python3 qemu sed socat spice-client-gtk swtpm-tools unzip usbutils util-linux uuidgen-runtime xdg-user-dirs xrandr zsync
```

Fedora

```
sudo dnf install bash coreutils curl edk2-tools genisoimage grep jq mesa-demos pciutils procps python3 qemu sed socat spice-gtk-tools swtpm unzip usbutils util-linux xdg-user-dirs xrandr zsync

```

```
sudo mkdir -p /home/jenkins/agent
sudo chown $(whoami):$(whoami) /home/jenkins/agent
sudo chown $jenkins:jenkins /home/jenkins/agent
sudo chmod 755 /home/jenkins/agent
```

```
sudo curl -s -o /home/jenkins/agent/agent.jar http://localhost:8080/jnlpJars/agent.jar
```

```
sudo nano /etc/systemd/system/jenkins-agent.service
```

```
[Unit]
Description=Jenkins Agent
After=network.target

[Service]
ExecStart=/usr/bin/java -jar /home/jenkins/agent/agent.jar \
  -url http://localhost:8080/ \
  -secret 922f5a63a260a9cfcf329121dc22ed208d38900b7c969fee9ea310791e816e71 \
  -name "fedora-local" \
  -webSocket \
  -workDir "/home/jenkins/agent"
User=jenkins
Group=jenkins
Restart=always
WorkingDirectory=/home/jenkins/agent

[Install]
WantedBy=multi-user.target

```

```
sudo systemctl daemon-reload
sudo systemctl enable jenkins-agent.service
sudo systemctl restart jenkins-agent.service
sudo systemctl status jenkins-agent.service
```

```
sudo chown $(whoami):$(whoami) /home/ahmetsoner/AS/prs
sudo chown -R jenkins:jenkins /home/ahmetsoner/AS/prs
sudo chmod -R 755 /home/ahmetsoner/AS/prs

```
