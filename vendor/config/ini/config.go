package ini

import (
	"github.com/Unknwon/goconfig"
	"strings"
)

//获取值类型
const (
	CFG_BOOL = iota
	CFG_FLOAT64
	CFG_INT
	CFG_INT64
	CFG_STRING
)

type Cfg struct {
	filename string  //.ini 配置文件
	sep string //key.value连接符

	instance goconfig.ConfigFile //配置文件实例
}


//加载配置文件
func LoadConfigFile(file string , sep string) (Cfg , error) {
	cfg := Cfg{}
	config , err := goconfig.LoadConfigFile(file)
	if err != nil {
		return cfg , err
	}
	cfg.filename = file
	cfg.sep = sep
	cfg.instance = *config
	return cfg , nil
}

//根据键名获取值[返回string][失败返回nil]
func (c Cfg) Get(key string) interface{} {
	return c.GetValue(key , CFG_STRING , nil)
}

//根据键名获取值[必须]
func (c Cfg) GetMust(key string , def string) string {
	val := c.GetValue(key , CFG_STRING , def)

	if val , ok := val.(string); !ok {
		return def
	} else {
		return val
	}
}

//根据键名获取值[返回bool][失败返回nil]
func (c Cfg) GetBool(key string) interface{} {
	return c.GetValue(key , CFG_BOOL , nil)
}

//根据键名获取值[必须]
func (c Cfg) GetBoolMust(key string , def bool) bool {
	val := c.GetValue(key , CFG_BOOL , def)

	if val , ok := val.(bool); !ok {
		return def
	} else {
		return val
	}
}

//根据键名获取值[返回float64][失败返回nil]
func (c Cfg) GetFloat64(key string) interface{} {
	return c.GetValue(key , CFG_FLOAT64 , nil)
}

//根据键名获取值[必须]
func (c Cfg) GetFloat64Must(key string , def float64) float64 {
	val := c.GetValue(key , CFG_FLOAT64 , def)

	if val , ok := val.(float64); !ok {
		return def
	} else {
		return val
	}
}

//根据键名获取值[返回int][失败返回nil]
func (c Cfg) GetInt(key string) interface{} {
	return c.GetValue(key , CFG_INT , nil)
}

//根据键名获取值[必须]
func (c Cfg) GetIntMust(key string , def int) int {
	val := c.GetValue(key , CFG_INT , def)

	if val , ok := val.(int); !ok {
		return def
	} else {
		return val
	}
}

//根据键名获取值[返回int64][失败返回nil]
func (c Cfg) GetInt64(key string) interface{} {
	return c.GetValue(key , CFG_INT64 , nil)
}

//根据键名获取值[必须]
func (c Cfg) GetInt64Must(key string , def int64) int64 {
	val := c.GetValue(key , CFG_INT64 , def)

	if val , ok := val.(int64); !ok {
		return def
	} else {
		return val
	}
}

//获取配置值
func (c Cfg) GetValue(key string , flag int , def interface{}) interface{} {
	keysplit := strings.Split(key , c.sep)
	if len(keysplit) < 2 {
		return def
	}
	switch flag {
		case CFG_BOOL:
			if value , err := c.instance.Bool(keysplit[0] , keysplit[1]);err != nil {
				return def
			} else {
				return value
			}
		case CFG_FLOAT64:
			if value , err := c.instance.Float64(keysplit[0] , keysplit[1]);err != nil {
				return def
			} else {
				return value
			}
		case CFG_INT:
			if value , err := c.instance.Int(keysplit[0] , keysplit[1]);err != nil {
				return def
			} else {
				return value
			}
		case CFG_INT64:
			if value , err := c.instance.Int64(keysplit[0] , keysplit[1]);err != nil {
				return def
			} else {
				return value
			}
		case CFG_STRING:
			if value , err := c.instance.GetValue(keysplit[0] , keysplit[1]);err != nil {
				return def
			} else {
				return value
			}
	}
	return def
}