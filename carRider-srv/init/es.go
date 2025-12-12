package init

import (
	"carRider-srv/global"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
)

func Es() {
	var err error
	global.Es, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			global.AppCfg.Es.Address,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("es 连接成功")
}
