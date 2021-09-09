package config

import "github.com/spf13/viper"

// LoadConfigN 加载配置文件，如果出错直接抛出致命异常。适用于必须读取的配置文件。
func LoadConfigN(confPtr interface{}, path string, name string) interface{} {
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.AddConfigPath(path)
	v.SetConfigName(name)
	err := v.ReadInConfig()

	if err != nil {
		panic("Err: " + err.Error())
	}
	err = v.Unmarshal(confPtr)
	if err != nil {
		panic("Err: " + err.Error())
	}
	return confPtr
}
