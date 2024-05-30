package service

import (
	"context"
	"log"

	pb "goserver/protos/pbuser"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConDB(){
	// return 
}
// type serviceDB struct{
// 	DB *pgxpool.Pool
// }
type pbjoin struct{
	pb.UnimplementedServiceJoinServer
	DB *pgxpool.Pool
}
type pblogin struct{
	pb.UnimplementedServiceLoginServer
	DB *pgxpool.Pool
}
type pbinfo struct{
	pb.UnimplementedServiceInfoServer
	DB *pgxpool.Pool
}
// func newsuser(dbpool *pgxpool.Pool) *serviceDB{
// 	return &serviceDB{DB:dbpool}
// }

func NewUserJoin(dbpool *pgxpool.Pool) *pbjoin{
	return &pbjoin{DB:dbpool}
}
func NewUserLogin(dbpool *pgxpool.Pool) *pblogin{
	return &pblogin{DB:dbpool}
}
func NewUserInfo(dbpool *pgxpool.Pool) *pbinfo{
	return &pbinfo{DB:dbpool}
}

func (s *pbjoin) userjoin(ctx context.Context ,req *pb.ReqJoin,res *pb.ResJoin)(*pb.ResJoin,error){
	log.Printf("userjoin req = %v ",req.GetId())
	// db 쿼리
	return &pb.ResJoin{State: true},nil
}

func (s *pblogin) userlogin(ctx context.Context ,req *pb.ReqLogin,res *pb.ResLogin)(*pb.ResLogin,error){
	log.Printf("userjoin req = %v ",req.GetId())
	// db 쿼리
	return &pb.ResLogin{State: true},nil
}

func (s *pbinfo) userinfo(ctx context.Context ,req *pb.ReqInfo,res *pb.ResInfo)(*pb.ResInfo,error){
	log.Printf("userjoin req = %v ",req.GetId())
	// db 쿼리
	return &pb.ResInfo{Id:"testid"},nil
}


