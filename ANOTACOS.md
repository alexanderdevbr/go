https://github.com/devfullcycle/go-intensivo-jul

Aula 1) https://www.youtube.com/watch?v=fAQ0cPduuIk
Aula 2) https://www.youtube.com/watch?v=mMEGjhgwmak
Aula 3) https://www.youtube.com/watch?v=aGqXxkGDJmw

go run main.go

go mod init github.com/alexander/projeto-go

GOOS=windows go build main.go

PS C:\Users\alexa\Workspace\Projetos-Go> go test ./...
?       github.com/alexander    [no test files]
ok      github.com/alexander/internal/entity    0.243s


go get github.com/stretchr/testify/assert
go get -t github.com/stretchr/testify/assert
go mod tidy

Resolvendo problema "panic: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub"
go env -w CGO_ENABLED=1