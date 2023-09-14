package apollo

import (
	"cac-micro/consts"
	"encoding/json"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

type ApolloConfig interface {
	GetConfig()
}

type ApolloChangeListener struct {
}

//OnChange 增加变更监控

//OnNewestChange 监控最新变更

func (l ApolloChangeListener) OnChange(event *storage.ChangeEvent) {
	// 部分更新，会将更新的部分，新旧数据都传过来
	fmt.Println(event)
}

func (l ApolloChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	// 将最新的数据 全部都传过来

}

func GetApollo(appID string) {
	c := &config.AppConfig{
		AppID:            appID,
		Cluster:          consts.ApolloCluster,
		IP:               consts.ApolloIP,
		NamespaceName:    consts.ApolloNamespaceName,
		IsBackupConfig:   consts.ApolloIsBackupConfig,
		Secret:           consts.ApolloSecret,
		BackupConfigPath: consts.ApolloBackupConfigPath,
	}

	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("初始化Apollo配置成功")

	client.AddChangeListener(ApolloChangeListener{})

	if err != nil {
		fmt.Println(err.Error())
	}
}
