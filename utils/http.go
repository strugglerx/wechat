package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func ContextApp(a App) App {
	return a
}

// ReduceUrl 合并url
func ReduceUrl(uri string, params Query) (string, error) {
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	parsedQuerys, err := url.ParseQuery(parsedUri.RawQuery)
	if err != nil {
		return "", err
	}
	for k, v := range params {
		parsedQuerys.Add(k, v)
	}
	if len(parsedQuerys) > 0 {
		return strings.Join([]string{strings.Replace(parsedUri.String(), "?"+parsedUri.RawQuery, "", -1), parsedQuerys.Encode()}, "?"), nil
	}
	return parsedUri.String(), nil
}

// checkTokenExpired 判断access_token 是否过期
func checkTokenExpired(responseString string, m App) bool {
	if value, ok := expiredToken[gjson.Get(responseString, "errcode").String()]; ok {
		m.GetAccessToken(true)
		return value
	}
	return false
}

// ExtractAppidAndAccessToken 提取appid 和 accessToken
func ExtractAppidAndAccessToken(appidAndAccessToken ...string) (ContextToken, error) {
	if len(appidAndAccessToken) == 2 {
		return ContextToken{
			Appid: appidAndAccessToken[0],
			Token: appidAndAccessToken[1],
		}, nil
	}
	return ContextToken{}, errors.New("appidAndAccessToken length must 2")
}

// FetchSource 获取资源
func FetchSource(uri string) []byte {
	uri2Url, _ := url.ParseRequestURI(uri)
	result := make([]byte, 0)
	request := http.Request{
		Method: "GET",
		Header: http.Header{
			"User-Agent": []string{"Mozilla/5.0 (Linux; Android 8.0; MI 6 Build/OPR1.170623.027; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/6.2 TBS/044304 Mobile Safari/537.36 MicroMessenger/6.7.3.1340(0x26070331) NetType/WIFI Language/zh_CN Process/tools"},
		},
		URL: uri2Url,
	}
	client := http.Client{}
	img, err := client.Do(&request)
	if err != nil {
		return result
	}
	imgBuffer, _ := ioutil.ReadAll(img.Body)
	return imgBuffer
}

//Get
/**
 * @param path string,params param,app App,domain string
 * @author struggler
 * @description client get 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return []byte,error
 **/
func Get(path string, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := Query{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case Query:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	var responseByte []byte
	response, err := http.Get(uri)
	defer response.Body.Close()
	if err != nil {
		return responseByte, err
	}

	responseByte, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return responseByte, err
	}
	responseString := string(responseByte)
	if strings.Contains(response.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(responseString, *m) {
		return Get(path, extends...)
	}
	return responseByte, nil
}

//PostBody
/**
 * @param path string,params param,app App,domain string
 * @author struggler
 * @description client postJson 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostBody(path string, body []byte, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := Query{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case Query:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	response, err := http.Post(uri, "", bytes.NewReader(body))
	defer response.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte(""), err
	}
	if strings.Contains(response.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(string(responseByte), *m) {
		return PostBody(path, body, extends...)
	}
	return responseByte, nil
}

//PostBufferFile
/**
 * @param path,name string, file multipart.File,fileName *multipart.fileName,params param,app App,domain string
 * @author struggler
 * @description client postBufferFile 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostBufferFile(path, name string, file io.Reader, fileName string, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := Query{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case Query:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(name, fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	result, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if strings.Contains(response.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(string(result), *m) {
		return PostBufferFile(path, name, file, fileName, extends...)
	}
	return result, err
}

//PostPathFile
/**
 * @param path,name string, file io.Reader,filePath string,params param,app App,domain string
 * @author struggler
 * @description client postFilePath 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostPathFile(path, name string, file io.Reader, filePath string, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := Query{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case Query:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(name, filePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	//_ = writer.WriteField("username", username)
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseByte := responseBody.Bytes()
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(string(responseByte), *m) {
		return PostPathFile(path, name, file, filePath, extends...)
	}
	return responseByte, err
}

func PostBufferFileWithField(path, name string, file io.Reader, fileName string, fields map[string]string, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := Query{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case Query:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range fields {
		writer.WriteField(k, v)
	}
	part, err := writer.CreateFormFile(name, fileName)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//request.Header.Set("header", "header")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	result := &bytes.Buffer{}
	_, err = result.ReadFrom(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if strings.Contains(response.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(result.String(), *m) {
		return PostBufferFileWithField(path, name, file, fileName, fields, extends...)
	}
	return result.Bytes(), nil
}
