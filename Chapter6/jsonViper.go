package main

// viper!!!
//  다양한 형태들의 포맷이 정해진 설정 파일을 사용할 때 파싱 작업을 처리해줌 유용하다.

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigStructure struct {
	MacPass     string `mapstructure:"macos"`
	LinuxPass   string `mapstructure:"linux"`
	WindowsPass string `mapstructure:"windows"`
	PostHost    string `mapstructure:"postgres"`
	MySQLHost   string `mapstructure:"mysql"`
	MongoHost   string `mapstructure:"mongodb"`
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

var CONFIG = ".config.json"

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Using default file", CONFIG)
	} else {
		CONFIG = os.Args[1]
	}

	viper.SetConfigType("json")                              //json 파일의 설정 파일을 설정할거다 말함
	viper.SetConfigFile(CONFIG)                              //CONFIG 파일을 찾음
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed()) //사용한 config 파일을 읽는다.
	viper.ReadInConfig()                                     // config 파일을 읽는다

	if viper.IsSet("macos") { //mapstructure를 통해 간단하게 접근 가능
		fmt.Println("macos:", viper.Get("macos"))
	} else {
		fmt.Println("macos not set!")
	}

	if viper.IsSet("active") {
		value := viper.GetBool("active")
		if value {
			postgres := viper.Get("postgres")
			mysql := viper.Get("mysql")
			mongo := viper.Get("mongodb")
			fmt.Println("P:", postgres, "My:", mysql, "Mo:", mongo)
		}
	} else {
		fmt.Println("active is not set!")
	}

	if !viper.IsSet("DoesNotExist") {
		fmt.Println("DoesNotExist is not set!")
	}

	var t ConfigStructure
	err := viper.Unmarshal(&t) //구조체를 json 형태로 unmarshal한다.
	if err != nil {
		fmt.Println(err)
		return
	}
	PrettyPrint(t)
}
