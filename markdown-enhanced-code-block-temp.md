```pwsh
# Set version (adjust as needed)
$cmakeVersion = "3.27.3"

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