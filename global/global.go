package global

import (
	"growth/pb"
)

var (
	ClientCoin  pb.UserCoinClient
	ClientGrade pb.UserGradeClient

	AllowOrigin = map[string]bool{
		"http://a.site.com": true,
		"http://b.site.com": true,
		"http://web.com":    true,
	}
)
