package main

import(
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
)

type UserGroup struct{
	Name string `json:"name"`
	Server    string `json:"server"`
	Port int	`json:"port"`
	Password  string `json:"password"`
	Cipher    string `json:"cipher"`
	Key       string `json:"key"`
	Keygen    int	 `json:"key_gen"`
	UDPTimeout	  time.Duration	 `json:"time_out"`
}

type Config struct {
	UserGroups []*UserGroup `json:"user_groups"`
	UDPTimeout time.Duration `json:"time_out"`
}


/**
 * @msg: 读取json 文件数据，按照参数 v 的格式解析
 * @param: fileName
 * @return: 
 */
func LoadJsonData(fileName string, v interface{}) error{
	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		return err;
	}

	dataJson := []byte(data)

	if err = json.Unmarshal(dataJson, v); err != nil{
		return err
	}
	return nil
}


func main(){
	/**
	 * 异常捕获
	 */
	defer func() {
		if err := recover(); err != nil{
			fmt.Println("发生异常：",err)
		}
	}()
	var MConfig = &Config{}
	if err := LoadJsonData("config.json", MConfig); err != nil{
		fmt.Println("读取配置错误")
		panic(err)
	}
	fmt.Println("获取到的数据为：", *MConfig.UserGroups[0])
}