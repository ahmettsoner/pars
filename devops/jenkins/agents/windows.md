```
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
