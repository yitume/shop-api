package trans

import (
	"github.com/yitume/shop-api/model/mysql"
)

// ReqCronList 你可以把ReqCronList嵌套到需要自行修改的结构体中
type ReqCronList struct {
	ReqPage
	mysql.Cron
}

// ReqCronCreate 你可以把ReqCronCreate或mysql.Cron嵌套到需要自行修改的结构体中
type ReqCronCreate = mysql.Cron

// ReqCronUpdate 你可以把ReqCronUpdate或mysql.Cron嵌套到需要自行修改的结构体中
type ReqCronUpdate = mysql.Cron

// ReqCronDelete 你可以把ReqCronDelete嵌套到需要自行修改的结构体中
type ReqCronDelete struct {
}
