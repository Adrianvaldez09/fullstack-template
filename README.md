SERVER

    cd to server

    go mod init example.com/project-name

        to create go.mod file

    go build

        to create executionable file

    go run . OR go run main.go

        run program

FRONTEND

    npm create vite@latest frontend

        create frontend react file

FULLSTACK CONNECTION
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/sqlite
    go get github.com/gin-contrib/cors
    go get -u golang.org/x/crypto/bcrypt
    go get -u github.com/golang-jwt/jwt/v5

    TO SET CGO_ENABLED
    $env:CGO_ENABLED="1"; go build

    TO SET C++ PATH
    https://www.youtube.com/watch?v=DMWD7wfhgNY

sqlite task-name.db

.schema task-names


    
