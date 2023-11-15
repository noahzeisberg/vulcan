if (-not ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Warning "You need to have PowerShell running as administrator."
    Write-Warning "Please exit the installation and try to run again as administrator."
    exit 1
}

Write-Output "Starting Vulcan installation..."
Write-Output "Setting up variables..."
$fileName = "v.exe"
$folderPath = Join-Path $env:USERPROFILE ".vulcan"
$filePath = Join-Path $folderPath $fileName
$githubApiUrl = "https://api.github.com/repos/noahonfyre/vulcan/releases/latest"

Write-Output "Creating local file..."
New-Item -ItemType Directory -Path $folderPath -Force | Out-Null

Write-Output "Fetching API information..."
$releaseInfo = Invoke-RestMethod -Uri $githubApiUrl

Write-Output "Processing data..."
$fileDownloadUrl = $releaseInfo.assets | Where-Object { $_.name -eq $fileName } | Select-Object -ExpandProperty browser_download_url

Write-Output "Downloading file from GitHub..."
Invoke-WebRequest -Uri $fileDownloadUrl -OutFile $filePath

Write-Output "Checking for environment variable..."
if (-not ($folderPath -in $env:VULCAN)) {
    Write-Output "Adding directory to your environment variables..."
    [Environment]::SetEnvironmentVariable("VULCAN", $filePath, "Machine")
}

Write-Output "Excluding Vulcan directory from Windows Defender..."
if (Get-Command -ErrorAction SilentlyContinue Get-MpPreference) {
    $existingExclusions = Get-MpPreference | Select-Object -ExpandProperty ExclusionPath
    if ($existingExclusions -contains $folderPath) {
        Write-Output "Exclusion for $folderPath already exists. No changes made."
    }
    else {
        $existingExclusions += $folderPath
        Set-MpPreference -ExclusionPath $existingExclusions
        Write-Output "Exclusion for $folderPath added successfully."
    }
}
else {
    Write-Warning "Windows Defender is not installed or not available on this system."
}

Write-Output " "
Write-Output "Installation of Vulcan complete! Please follow the other steps in the README."
Write-Output "https://github.com/noahonfyre/vulcan#installation"
Write-Output " "