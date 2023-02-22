/*
 * @Author: caiyan1
 * @Date: 2023-02-17 17:37:50
 * @LastEditors: caiyan1
 * @LastEditTime: 2023-02-21 20:02:37
 * @Description:
 */
package atom

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cccaiyan/upipe-go-toolkit/model"
	"github.com/cccaiyan/upipe-go-toolkit/utils"
)

func GetInput(key string) interface{} {
	outputDir := os.Getenv("PLUGIN_PARAMS")
	atomAlias := os.Args[1]
	inputsString,_ := utils.Sh([]string{
		fmt.Sprintf("grep INPUTS= %s/%s_inputs.txt", outputDir, atomAlias),
	})
	inputsString = strings.Replace(inputsString, "INPUTS=", "", 1)
	if inputsString == "" {
		fmt.Printf("%s 插件无原子输入参数\n", atomAlias)
		return ""
	}

	inputs := new([]model.InputParams)
	json.Unmarshal([]byte(inputsString), inputs)
	for _, ip := range *inputs {
		if ip.Key == strings.ToUpper(key) {
			switch ip.ParamType{
			case "string", "select","code":
				return ip.Value.(string)
			case "boolean" :
				value,err := strconv.ParseBool(ip.Value.(string))
				if err!= nil {
					fmt.Printf("%s 参数非boolean类型\n", ip.Key)
				}
				return value
			case "list":
				slice := strings.Split(ip.Value.(string), ",")
				return slice
			case "mapList":
				slice := ip.Value.([]interface{})
				if len(slice) == 0 {
					return slice
				}
				mapInputs := map[string]string{}
				for _, imlp := range slice {
					mp := imlp.(map[string]interface{})
					mapInputs[mp["name"].(string)] = mp["info"].(string)
				}
				return mapInputs
			default:
				fmt.Printf("%s 参数类型无法自动解析", ip.ParamType)
			}
			
		}
	}
	return ""
}

