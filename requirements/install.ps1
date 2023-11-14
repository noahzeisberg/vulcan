$fileName = "v.exe"
$folderPath = Join-Path $env:USERPROFILE ".vulcan"
$filePath = Join-Path $folderPath $fileName
$githubApiUrl = "https://api.github.com/repos/noahonfyre/vulcan/releases/latest"

New-Item -ItemType Directory -Path $folderPath -Force | Out-Null

$releaseInfo = Invoke-RestMethod -Uri $githubApiUrl

$fileDownloadUrl = $releaseInfo.assets | Where-Object { $_.name -eq $fileName } | Select-Object -ExpandProperty browser_download_url

Invoke-WebRequest -Uri $fileDownloadUrl -OutFile $filePath

$env:Path += ";$folderPath"

Write-Output "Installation of Vulcan complete! Please follow the other steps in the README."