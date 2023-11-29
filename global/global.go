package global

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	ESClient *elasticsearch.Client
)
