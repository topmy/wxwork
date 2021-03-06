// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package verify 身份验证
package verify

import (
	"net/url"

	"github.com/fastwego/wxwork/corporation"
)

const (
	apiGetUserInfo = "/cgi-bin/user/getuserinfo"
)

/*
获取访问用户身份

该接口用于根据code获取成员信息

See: https://work.weixin.qq.com/api/doc/90000/90135/91023

GET https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE参数说明：
*/
func GetUserInfo(ctx *corporation.App, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetUserInfo + "?" + params.Encode())
}
