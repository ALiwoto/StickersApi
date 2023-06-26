@echo off
git pull && powershell -command "Stop-service -Force -name "StickersApi" -ErrorAction SilentlyContinue; go mod tidy; go build; Start-service -name "StickersApi""
:: Hail Hydra
