:: this bat file is supposed to build the application from source code, build the documents and
:: run the StickersApi app after all of these are done.
:: created by ALiwoto (woto@kaizoku.cyou)
@echo off
TITLE Building StickersApi binary file
go mod tidy
go build -o StickersApi.exe
cd .\docs
.\Build.bat && cd .. && .\StickersApi.exe
