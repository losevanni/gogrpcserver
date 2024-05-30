protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service.proto


# go mod init goserver
# go get google.golang.org/grpc
# go get google.golang.org/protobuf/cmd/protoc-gen-go
# go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go get github.com/jackc/pgx/v4
# go get github.com/jackc/pgx/v4/pgxpool@v4.18.3

sudo apt install postgresql postgresql-contrib
su -i -u postgres

psql 

# CREATE DATABASE goserverdb;
# CREATE USER gsdbuser WITH ENCRYPTED PASSWORD 'kms1030';
# GRANT ALL PRIVILEGES ON DATABASE goserverdb TO gsdbuser;

psql -U gsdbuser -d goserverdb


