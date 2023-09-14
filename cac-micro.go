package smicro

import (
	"cac-micro/config/apollo"
)

type Cmd struct {
}

// NewCommands 初始化一个 服务
func (cmd *Cmd) NewCommands() *Cmd {
	return new(Cmd)
}

// SetApolloConfig 设置 配置中心信息
func (cmd *Cmd) SetApolloConfig(appId string) {
	apollo.GetApollo(appId)
}
