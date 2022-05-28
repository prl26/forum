package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


var GVA_DB *gorm.DB
var	GVA_LOG *zap.Logger
var GVA_Redis *redis.Client