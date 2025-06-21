* Install Chocolatey
```pwsh
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```


* Install Go
```pwsh
$goVersion = "1.22.0"
$goUrl = "https://go.dev/dl/go$goVersion.windows-amd64.msi"
$installerPath = "$env:TEMP\go-installer.msi"
Invoke-WebRequest -Uri $goUrl -OutFile $installerPath
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet /norestart" -Wait
#Remove-Item $installerPath
#go version
```


* Install Nodejs
```pwsh
$nodeVersion = "22.16.0"
$nodeUrl = "https://nodejs.org/dist/v$nodeVersion/node-v$nodeVersion-x64.msi"
$installerPath = "$env:TEMP\nodejs-installer.msi"
Invoke-WebRequest -Uri $nodeUrl -OutFile $installerPath
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet /norestart" -Wait
#Remove-Item $installerPath
#node -v
#npm -v
```



* Install Git
```pwsh
$gitVersion = "2.45.1"
$gitUrl = "https://github.com/git-for-windows/git/releases/download/v$gitVersion.windows.1/Git-$gitVersion-64-bit.exe"
$installerPath = "$env:TEMP\Git-Installer.exe"
Invoke-WebRequest -Uri $gitUrl -OutFile $installerPath
Start-Process -FilePath $installerPath -ArgumentList "/VERYSILENT" -Wait
#Remove-Item $installerPath
#git --version
```

* Instal OpenJDK
```pwsh
$jdkUrl = "https://github.com/adoptium/temurin17-binaries/releases/download/jdk-17.0.9%2B9.1/OpenJDK17U-jdk_x64_windows_hotspot_17.0.9_9.msi"
$installerPath = "$env:TEMP\temurin17.msi"
Invoke-WebRequest -Uri $jdkUrl -OutFile $installerPath
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet /norestart" -Wait
Remove-Item $installerPath
$javaHome = "C:\Program Files\Eclipse Adoptium\jdk-17.0.9.9-hotspot"
[Environment]::SetEnvironmentVariable("JAVA_HOME", $javaHome, "User")

$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = "$javaHome\bin"

if ($currentPath -notlike "*$newPath*") {
    $updatedPath = "$currentPath;$newPath"
    [Environment]::SetEnvironmentVariable("Path", $updatedPath, "User")
    Write-Host "`n✅ JAVA_HOME ve PATH (User) seviyesinde ayarlandı. Yeni terminal açarak tekrar deneyin: 'java -version'"
} else {
    Write-Host "`nℹ PATH zaten ayarlanmış (User)."
}

#java -version
#javac -version
```

* Install WIX
```pwsh
$wixMsiUrl = "https://github.com/wixtoolset/wix/releases/download/v6.0.1/wix-cli-x64.msi"
$tempMsiPath = "$env:TEMP\wix-cli-x64.msi"
Invoke-WebRequest -Uri $wixMsiUrl -OutFile $tempMsiPath
Start-Process msiexec.exe -ArgumentList "/i `"$tempMsiPath`" /quiet /norestart" -Wait
Remove-Item $tempMsiPath
$wixPath = "C:\Program Files\WiX Toolset v6.0\bin"
$env:PATH = $wixPath + ";" + $env:PATH

#!!! BU bölüm kontrollü olarak yeniden test edilecek 
$oldPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
$newPath = $oldPath + ";" + $wixPath
[Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
```


* Restart Server
```
Restart-Computer -Confirm
```

* Add Agent
```pwsh
$jenkinsUrl   = "http://192.168.200.1:8080"
$secret       = "7d417c2e9f05e250bb734bdd8e13050ff2185bad27d9fb12395db5f9e8e98e22"
$agentName    = "windows"
$workDir      = "C:\Users\automation\Desktop\Agent"
$agentJarUrl  = "$jenkinsUrl/jnlpJars/agent.jar"
$agentJarPath = "$env:TEMP\agent.jar"
Invoke-WebRequest -Uri $agentJarUrl -OutFile $agentJarPath -UseBasicParsing

$taskName    = "StartJenkinsAgent"
$arguments = "-jar `"$agentJarPath`" -url $jenkinsUrl -secret $secret -name $agentName -webSocket -workDir `"$workDir`""

$Action = New-ScheduledTaskAction -Execute "java.exe" -Argument $arguments
$Trigger = New-ScheduledTaskTrigger -AtStartup
$Principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -RunLevel Highest

if (Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue) {
    Unregister-ScheduledTask -TaskName $taskName -Confirm:$false
}

Register-ScheduledTask -TaskName $taskName -Action $Action -Trigger $Trigger -Principal $Principal

Start-ScheduledTask -TaskName $taskName
```




* Install Cmake
```pwsh
choco install cmake -y
[System.Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\Program Files\CMake\bin", [System.EnvironmentVariableTarget]::Machine)
```

* Make
```pwsh
choco install make -y
```

* Install MinGW
```pwsh
choco install mingw -y
```


* Git Release Manager
```pwsh
npm install -g git-release-manager@0.0.14
```

