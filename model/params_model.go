/*
 * @Author: caiyan1
 * @Date: 2023-02-20 14:21:30
 * @LastEditors: caiyan1
 * @LastEditTime: 2023-02-21 20:01:02
 * @Description:
 */
package model

type InputParams struct {
	Key string `json:"name"`
	Value interface{} `json:"value"`
	ParamType string `json:"param_type"`
}


type InputMapListParam struct {
	Name string `json:"name"`
	Info string `json:"info"`
}