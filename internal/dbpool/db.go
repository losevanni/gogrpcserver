package dbpool

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB(connStr string)(*pgxpool.Pool,error){
	dbpool,err:=pgxpool.Connect(context.Background(),connStr)
	if err != nil{
		return nil, err
	}
	return dbpool,nil
}