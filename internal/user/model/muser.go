package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct{
	IDX int64
	ID string
	PASSWD string
}

func createUser(ctx context.Context, dbpool *pgxpool.Pool,user User)(bool,error){
	query:="INSERT INTO users (id,passwd) VALUES ($1,$2) RETURNING idx"
	// 자동 생성된 idx 받아 오기
	err:=dbpool.QueryRow(ctx,query,user.ID,user.PASSWD).Scan(&user.IDX)
	fmt.Print("createUser, ",err)
	if err != nil{
		return false,err
	}
	return true,nil
}
func loginUser(ctx context.Context, dbpool *pgxpool.Pool,user User)(bool,error){
	query:=""
	// 자동 생성된 idx 받아 오기
	err:=dbpool.QueryRow(ctx,query,user.ID,user.PASSWD).Scan(&user.IDX)
	fmt.Print("createUser, ",err)
	if err != nil{
		return false,err
	}
	return true,nil
}
func infoUser(ctx context.Context, dbpool *pgxpool.Pool,user User)(bool,error){
	query:=""
	// 자동 생성된 idx 받아 오기
	err:=dbpool.QueryRow(ctx,query,user.ID,user.PASSWD).Scan(&user.IDX)
	fmt.Print("createUser, ",err)
	if err != nil{
		return false,err
	}
	return true,nil
}