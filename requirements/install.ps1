$filePath = Join-Path $env:USERPROFILE ".vulcan" "v.exe"
$githubApiUrl = "https://api.github.com/repos/noahonfyre/vulcan/releases/latest"

New-Item -ItemType Directory -Path $folderPath -Force

$releaseInfo = Invoke-RestMethod -Uri $githubApiUrl
$fileDownloadUrl = $releaseInfo.assets | Where-Object { $_.name -eq $fileName } | Select-Object -ExpandProperty browser_download_url
Invoke-WebRequest -Uri $fileDownloadUrl -OutFile $filePath