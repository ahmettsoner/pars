* Install Chocolatey
```pwsh
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```


* Install Go
```pwsh
# Set version (adjust if needed)
$goVersion = "1.22.0"

# Download URL
$goUrl = "https://go.dev/dl/go$goVersion.windows-amd64.msi"

# Target path for the MSI installer
$installerPath = "$env:TEMP\go-installer.msi"

# Download the MSI
Invoke-WebRequest -Uri $goUrl -OutFile $installerPath

# Install silently
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet /norestart" -Wait

# Clean up
Remove-Item $installerPath

# Confirm installation
go version
```


* Install Nodejs
```pwsh
# Set version (adjust as needed)
$nodeVersion = "22.16.0"

# Download URL for Windows x64 MSI installer
$nodeUrl = "https://nodejs.org/dist/v$nodeVersion/node-v$nodeVersion-x64.msi"

# Target path for the installer
$installerPath = "$env:TEMP\nodejs-installer.msi"

# Download the installer
Invoke-WebRequest -Uri $nodeUrl -OutFile $installerPath

# Run installer silently
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet /norestart" -Wait

# Clean up
Remove-Item $installerPath

# Check Node.js version
node -v
npm -v
```


* Install Cmake
```pwsh
# Set version (adjust as needed)
$cmakeVersion = "4.0.3"

# Download URL for Windows x64 installer (.msi)
$cmakeUrl = "https://github.com/Kitware/CMake/releases/download/v$cmakeVersion/cmake-$cmakeVersion-windows-x86_64.msi"

# Target path for the installer
$installerPath = "$env:TEMP\cmake-installer.msi"

# Download the installer
Invoke-WebRequest -Uri $cmakeUrl -OutFile $installerPath

# Run installer silently, add to PATH for all users
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet ADD_CMAKE_TO_PATH=System" -Wait

# Clean up
Remove-Item $installerPath

# Verify installation
cmake --version
```

* Make
```pwsh
```

* Install MinGW
```pwsh
```

* Install Git
```pwsh
# Set version (adjust if needed)
$gitVersion = "2.45.1"

# Download URL for 64-bit Git installer
$gitUrl = "https://github.com/git-for-windows/git/releases/download/v$gitVersion.windows.1/Git-$gitVersion-64-bit.exe"

# Path to save installer
$installerPath = "$env:TEMP\Git-Installer.exe"

# Download Git installer
Invoke-WebRequest -Uri $gitUrl -OutFile $installerPath

# Install silently (default options)
Start-Process -FilePath $installerPath -ArgumentList "/VERYSILENT" -Wait

# Clean up
Remove-Item $installerPath

# Verify Git version
git --version

```

* Instal OpenJDK
```pwsh
# Java versiyonu (Temurin 17 LTS örneği)
# Java MSI URL'si (senin verdiğin bağlantı)
$jdkUrl = "https://github.com/adoptium/temurin17-binaries/releases/download/jdk-17.0.9%2B9.1/OpenJDK17U-jdk_x64_windows_hotspot_17.0.9_9.msi"

# Geçici dosya yolu
$installerPath = "$env:TEMP\temurin17.msi"

# MSI dosyasını indir
Invoke-WebRequest -Uri $jdkUrl -OutFile $installerPath

# Java'yı sessizce kur
Start-Process -FilePath "msiexec.exe" -ArgumentList "/i `"$installerPath`" /quiet /norestart" -Wait

# Geçici dosyayı temizle
Remove-Item $installerPath

# Java kurulum yolu (yine senin verdiğin .msi varsayılan yolu)
$javaHome = "C:\Program Files\Eclipse Adoptium\jdk-17.0.9.9-hotspot"

# JAVA_HOME değişkenini kullanıcı seviyesinde ayarla
[Environment]::SetEnvironmentVariable("JAVA_HOME", $javaHome, "User")

# PATH'e ekle (eğer yoksa)
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = "$javaHome\bin"

if ($currentPath -notlike "*$newPath*") {
    $updatedPath = "$currentPath;$newPath"
    [Environment]::SetEnvironmentVariable("Path", $updatedPath, "User")
    Write-Host "`n✅ JAVA_HOME ve PATH (User) seviyesinde ayarlandı. Yeni terminal açarak tekrar deneyin: 'java -version'"
} else {
    Write-Host "`nℹ PATH zaten ayarlanmış (User)."
}


# Versiyon kontrolü
java -version
javac -version
```

* Git Release Manager
```pwsh
npm install -g git-release-manager@0.0.14
```


* Add Agent
```pwsh
# Konfigürasyon ayarları
$jenkinsUrl   = "http://192.168.200.1:8080"
$secret       = "eaec0e226f7e6f22242cb34a1117606eb4d3093082829df49619447ec0fe49f3"
$agentName    = "windows"
$workDir      = "C:\Users\automation\Desktop\Agent"
$agentJarUrl  = "$jenkinsUrl/jnlpJars/agent.jar"
$agentJarPath = "$env:TEMP\agent.jar"

# agent.jar dosyasını indir
Invoke-WebRequest -Uri $agentJarUrl -OutFile $agentJarPath -UseBasicParsing



$taskName    = "StartJenkinsAgent"
# Java agent argümanları
$arguments = "-jar `"$agentJarPath`" -url $jenkinsUrl -secret $secret -name $agentName -webSocket -workDir `"$workDir`""

# Scheduled Task oluşturma
$Action = New-ScheduledTaskAction -Execute "java.exe" -Argument $arguments
$Trigger = New-ScheduledTaskTrigger -AtStartup
$Principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -RunLevel Highest

# Task register et (varsa önce sil)
if (Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue) {
    Unregister-ScheduledTask -TaskName $taskName -Confirm:$false
}

Register-ScheduledTask -TaskName $taskName -Action $Action -Trigger $Trigger -Principal $Principal

# Task'ı hemen başlat
Start-ScheduledTask -TaskName $taskName
```



* Install WIX
```pwsh
# WiX CLI MSI URL'si
$wixMsiUrl = "https://github.com/wixtoolset/wix/releases/download/v6.0.1/wix-cli-x64.msi"

# Geçici dosya yolu
$tempMsiPath = "$env:TEMP\wix-cli-x64.msi"

# MSI dosyasını indir
Invoke-WebRequest -Uri $wixMsiUrl -OutFile $tempMsiPath

# MSI'yı sessiz (silent) kur
Start-Process msiexec.exe -ArgumentList "/i `"$tempMsiPath`" /quiet /norestart" -Wait

# İndirilen MSI dosyasını sil (isteğe bağlı)
Remove-Item $tempMsiPath

# WiX CLI yolunu kontrol etmek için ekleyebilirsin (örneğin PATH değişkenine)
# Normalde kurulum PATH'e ekler, yoksa
$env:PATH = "C:\Program Files\WiX Toolset v6.0\bin;" + $env:PATH

```
