package handle

import (
	"context"
	"growth/pb"
)

// 用户积分服务
type UgCoinServer struct {
	pb.UnimplementedUserCoinServer
}

// ListTasks 获取所有的积分任务列表
func (s *UgCoinServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksReply, error) {
	out := &pb.ListTasksReply{}
	return out, nil
}

// UserCoinInfo 获取用户的积分信息
func (s *UgCoinServer) UserCoinInfo(ctx context.Context, in *pb.UserCoinInfoRequest) (*pb.UserCoinInfoReply, error) {
	out := &pb.UserCoinInfoReply{}
	return out, nil
}

// UserDetails 获取用户的积分明细列表
func (s *UgCoinServer) UserDetails(ctx context.Context, in *pb.UserDetailsRequest) (*pb.UserDetailsReply, error) {
	out := &pb.UserDetailsReply{}
	return out, nil
}

// UserCoinChange 调整用户积分-奖励和惩罚都是这个接口
func (s *UgCoinServer) UserCoinChange(ctx context.Context, in *pb.UserCoinChangeRequest) (*pb.UserCoinChangeReply, error) {
	out := &pb.UserCoinChangeReply{}
	return out, nil
}
