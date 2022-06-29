package mock

import "github.com/pz2147/p-rpc-1/internal/model"

func UserMock() model.UserInfo {
	user := model.UserInfo{
		Uid:      "1234567890",
		Nickname: "paul",
		Pic:      "pic/pic/pic",
	}
	return user
}