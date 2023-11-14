package remote

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type VjCrawler struct {
	userName string
	passWord string
	cookie   string
}

var VJCrawler VjCrawler

const getRankUrl = "https://vjudge.net/contest/rank/single/"
const loginUrl = "https://vjudge.net/user/login"
const checkLoginStatusUrl = "https://vjudge.net/user/checkLogInStatus"

func init() {
	//config := model.Config{}
	//middle.DB.Model(&model.Config{}).First(&config)
	//VJCrawler.userName = config.VjUserName
	//VJCrawler.passWord = config.VjPassWord
	//VJCrawler.cookie = config.VjCookie
	//if flag, err := VJCrawler.checkLoginStatus(); flag && err != nil {
	//	return
	//}
	//if cookie, err := VJCrawler.Login(); err != nil {
	//	VJCrawler.cookie = cookie
	//}
}

func (vj *VjCrawler) checkLoginStatus() (bool, error) {

	req, err := http.NewRequest("POST", checkLoginStatusUrl, nil)

	if err != nil {
		return false, err
	}

	req.Header.Add("cookie", vj.cookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	return string(body) == "true", nil
}

func (vj *VjCrawler) Login() (string, error) {
	if len(vj.cookie) > 0 {
		return vj.cookie, nil
	}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", "3553928717@qq.com")
	_ = writer.WriteField("password", "wsy16675060764")
	_ = writer.WriteField("captcha", "")
	err := writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", loginUrl, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if string(body) != "success" {
		return "", errors.New("login fail")
	}

	cookie := res.Header.Values("Set-Cookie")
	for _, s := range cookie {
		fmt.Println(s)
	}
	//fmt.Println(cookie)
	return string(body), nil
}
