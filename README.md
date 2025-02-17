To build on windows -> CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/cql.exe

Queries
1. CREATE column1, column2, ... AS "filename";
2. SHOW *; -> It shows all the filenames registered
3. 