/*
 * @Author: caiyan1
 * @Date: 2022-11-02 18:02:49
 * @LastEditors: caiyan1
 * @LastEditTime: 2023-02-20 11:41:00
 * @Description:
 */
package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	"github.com/cccaiyan/upipe-go-toolkit/model"
	"github.com/go-resty/resty/v2"
)

/**
 * @description: 查找文件
 * @param {*} regexPath
 * @param {string} targetPath
 * @return {*}
 * @Author: caiyan1
 */



func FindFiles(regexPath, targetPath string) (string, error) {
	fmt.Println("AntPathMatcher:", regexPath)

	_, err := Sh([]string{
		"cd " + targetPath,
		fmt.Sprintf("chmod 777 %s/common.jar", model.ToolsInstallDir),
		fmt.Sprintf("java -jar %s/common.jar \"%s\" \"%s\" \"%s\" ", model.ToolsInstallDir, model.FindFiles, targetPath, regexPath),
	})
	if err != nil {
		fmt.Println("java ant style path error", err)
		return "", err
	}
	key := fmt.Sprintf("%s_%s", targetPath, regexPath)
	message, err := Sh([]string{
		"cd " + targetPath,

		fmt.Sprintf("cat %s.txt", strings.ToUpper(CalcMd5(key))),
	})
	if err != nil {
		return "", err
	}
	return message, nil
}

func Sh(cmd []string) (stdout string, err error) {
	if cmd == nil {
		return "", nil
	}

	exe := exec.Command("sh", "-c", strings.Join(cmd, " && "))
	result, err := exe.Output()
	return string(result), err
}

// CalcMd5
//
//	@Description: 计算字符串的md5
//	@param message
//	@return string
func CalcMd5(message string) string {
	data := []byte(message)
	return fmt.Sprintf("%x", md5.Sum(data))
}

/**
 * @description: 保留两位小数除法
 * @param {*} molecular
 * @param {float64} denominator
 * @return {*}
 * @Author: caiyan1
 */
func Division(molecular, denominator float64) float64 {
	if denominator == 0.0 {
		return 0.0
	}
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", molecular/denominator), 64)
	return result
}

/**
 * @description: 模糊查找当前工作空间下的多个匹配的文件夹
 * @param {string} regesPath
 * @return {*}
 * @Author: caiyan1
 */
func FindFolders(regexPath, rootPath string) ([]string, error) {
	folders, err := Sh([]string{
		"cd " + rootPath,
		"find . -path '" + regexPath + "'",
	})
	if err != nil {
		return nil, err
	}
	folders = strings.TrimSpace(folders)
	return strings.Split(folders, "\n"), nil
}

func Tar(rootPath, dest, foleders string) error {
	_, err := Sh([]string{
		"cd " + rootPath,
		fmt.Sprintf("tar -zcvf %s %s", dest, foleders),
	})
	return err
}

/**
 * @description: 判断obj是否在target中，target支持的类型arrary,slice,map
 * @param {interface{}} obj
 * @param {interface{}} target
 * @return {*}
 * @Author: caiyan1
 */
func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

func PostClient(uri string, headers map[string]string, body interface{}, res interface{}) *resty.Response {
	// Create a Resty Client
	client := resty.New()
	fmt.Println("postClient")
	fmt.Println(uri)
	data, _ := json.Marshal(body)
	fmt.Println(string(data))
	response, err := client.R().SetHeader("Content-Type", "application/json").SetHeaders(headers).SetBody(body).SetResult(&res).Post(uri)
	defer fmt.Println(response)
	if err != nil {
		panic(err)
	}
	return response
}
