$path = Get-Location

Set-Location $PSScriptRoot

    Set-Location "../services/1_user"

    go build -race -o service_exec.exe

    Start-Process ./service_exec.exe

Set-Location $path

# windows not allow run ps script:
# (admin start)Set-ExecutionPolicy RemoteSigned
