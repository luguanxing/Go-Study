package parser

import (
	"crawler/zhenai/model"
	"crawler/zhenai/types"
	"log"
	"regexp"
	"strings"
)

// 信息字段
// `<div class="des f-cl" data-v-3c42fade>淄博 | 19岁 | 高中及以下 | 未婚 | 165cm | 3001-5000元</div> <div class="actions" data-v-3c42fade>`
const infoRe = `<div class="des f-cl" data-v-[0-9a-zA-Z]+>(.*)</div> <div class="actions" data-v-[0-9a-zA-Z]+>`

// 个人简介
// <div class="m-content-box m-des" data-v-bff6f798=""><span data-v-bff6f798="">不想说过去，心酸，只想想未来。不想说太多过去的事。</span>
const descRe = `<div class="m-content-box m-des" data-v-[0-9a-zA-Z]+[^>]*><span data-v-[0-9a-zA-Z]+[^>]*>([^<]+)</span>`

// 预先编译，避免多用户时重复编译
var infoReCom = regexp.MustCompile(infoRe)
var descCom = regexp.MustCompile(descRe)

func ParseUser(bytes []byte, nameStr string) types.ParseResult {
	result := types.ParseResult{}
	infoMatch := infoReCom.FindSubmatch(bytes)
	infoStr := "";
	if infoMatch != nil {
		infoStr = string(infoMatch[1])
	}
	descMatch := descCom.FindSubmatch(bytes)
	descStr := "";
	if descMatch != nil {
		descStr = string(descMatch[1])
	}
	user := getUserFromInfo(infoStr)
	user.Name = nameStr
	user.Description = descStr
	log.Printf("已抓取用户\n名称[%s]\n个人资料[%s]\n个人简介[%s]", nameStr, infoStr, descStr)
	return result
}

func getUserFromInfo(info string) model.User {
	infos := strings.Split(info, "|")
	user := model.User{}
	user.Location = strings.TrimSpace(infos[0])
	user.Age = strings.TrimSpace(infos[1])
	user.Education = strings.TrimSpace(infos[2])
	user.Marriage = strings.TrimSpace(infos[3])
	user.Height = strings.TrimSpace(infos[4])
	user.Income = strings.TrimSpace(infos[5])
	return user
}