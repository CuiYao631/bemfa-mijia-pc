package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
type T struct {
}

// PostHandle 发送post请求
func PostHandle(ctx context.Context, url string, msg []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	var resData Response
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(msg))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	req.Header.Set("x-envoy-internal", "true")
	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		//解析返回数据
		json.NewDecoder(resp.Body).Decode(&resData)
		return nil, err
	}
	bodys, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodys, nil
}

// GetHandle 发送Get请求
func GetHandle(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	var resData Response
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			return nil, errors.New("401")
		}
		//解析返回数据
		json.NewDecoder(resp.Body).Decode(&resData)
		return nil, err
	}
	bodys, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodys, nil
}

// PostForm 发送post from请求
func PostForm(ctx context.Context, urls string, formData map[string]string) ([]byte, error) {
	client := &http.Client{}
	var resData Response
	data := url.Values{}
	// 构建查询参数。
	for k, v := range formData {
		data.Add(k, v)
	}
	// 将查询参数转换为字节切片，并设置为请求体。
	reqBody := bytes.NewBufferString(data.Encode())

	// 创建一个请求。
	req, err := http.NewRequestWithContext(ctx, "POST", urls, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		//解析返回数据
		err := json.NewDecoder(resp.Body).Decode(&resData)
		if err != nil {
			return nil, err
		}
		errMsg := "请求失败"
		switch resData.Code {
		case 10002:
			errMsg = "请求参数错误"
		case 20102:
			errMsg = "用户不存在"
		case 20103:
			errMsg = "用户已存在"
		case 40000:
			errMsg = "未知错误"
		case 40007:
			errMsg = "请求次数过多"
		case 40008:
			errMsg = "密码错误"
		}

		return nil, errors.New(errMsg)
	}
	bodys, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodys, nil
}
