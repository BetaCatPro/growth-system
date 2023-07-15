package handle

import (
	"context"
	"growth/pb"
)

// 用户等级服务
type UgGradeServer struct {
	pb.UnimplementedUserGradeServer
}

// ListGrades 获取所有的等级信息列表
func (s *UgGradeServer) ListGrades(ctx context.Context, in *pb.ListGradesRequest) (*pb.ListGradesReply, error) {
	out := &pb.ListGradesReply{}
	return out, nil
}

// ListGradePrivileges 获取等级的特权列表
func (s *UgGradeServer) ListGradePrivileges(ctx context.Context, in *pb.ListGradePrivilegesRequest) (*pb.ListGradePrivilegesReply, error) {
	out := &pb.ListGradePrivilegesReply{}
	return out, nil
}

// CheckUserPrivilege 检查用户是否有某个产品特权
func (s *UgGradeServer) CheckUserPrivilege(ctx context.Context, in *pb.CheckUserPrivilegeRequest) (*pb.CheckUserPrivilegeReply, error) {
	out := &pb.CheckUserPrivilegeReply{}
	return out, nil
}

// UserGradeInfo 获取用户的等级信息
func (s *UgGradeServer) UserGradeInfo(ctx context.Context, in *pb.UserGradeInfoRequest) (*pb.UserGradeInfoReply, error) {
	out := &pb.UserGradeInfoReply{}
	return out, nil
}

// UserGradeChange 调整用户的等级成长值
func (s *UgGradeServer) UserGradeChange(ctx context.Context, in *pb.UserGradeChangeRequest) (*pb.UserGradeChangeReply, error) {
	out := &pb.UserGradeChangeReply{}
	return out, nil
}
