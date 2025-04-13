# go-backend-member-vetautet

## How to run 

If your configuration is low, you should only start one docker service, which is
```bash
~ ./environment/start.sh
```

On the contrary, if your configuration supports adding a redis sentinel docker cluster, then start it

```bash
~ ./environment/cluster-redis/start.sh
```
Note: please replace your server or local ip -> ip-address -> here when starting the sentinel service, such as `../cluster-redis/docker-compose-cluster.sh` and `./cluster-redis/sentinel.conf`

```go
ip-address is xxx.xxxx.xxx.xx
```

And docker will start with a single node of Redis, and comment the sentinel service in the run.go file

```go
package initialize

import (
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// load configuration
	LoadConfig()
	InitLogger()
	InitMysql()
	InitMysqlC()
	InitRedis() // open this service...
	InitRedisSentinel()// this commented out
	InitKafka()
	InitServiceInterface()
	r := InitRouter()
	return r
}
```

Then start go api by

```go
~ make dev
```

And try

```go
~ curl http://localhost:8002/v1/2024/ticket/item/1
```

Thanks!