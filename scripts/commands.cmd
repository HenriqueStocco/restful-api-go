@REM Requests via poweshell
powershell -command "Invoke-WebRequest -Uri http://localhost:8080 -UseBasicParsing -Method GET"

@REM Hot Reload with Bun + Nodemon
@REM Direct
powershell bun x nodemon --exec go run .\cmd\api\main.go --signal SIGTERM

@REM or add package.json file and add scripts
@REM "scripts": { "go": "nodemon --exec go run .\cmd\api\main.go --signal SIGTERM" }
@REM run: bun run go