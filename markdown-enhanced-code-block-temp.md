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