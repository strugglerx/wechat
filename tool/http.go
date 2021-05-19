package tool

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// ContextApp
func ContextApp(a App) App{
	return  a
}

// ReduceUrl 合并url
func ReduceUrl(uri string, params MapStr) (string, error) {
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
//40001 获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口
//42001 access_token 超时，请检查 access_token 的有效期，请参考基础支持 - 获取 access_token 中，对 access_token 的详细机制说明
func checkTokenExpired(responseString string, m App) bool {
	if strings.Contains(expiredToken, gjson.Get(responseString, "errcode").String()) {
		m.GetAccessToken().UpdateTime = 0
		return true
	}
	return false
}

/**
 * @param path string,params param,app App,domain string
 * @author struggler
 * @description client get 注意：当传递mp的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func Get(path string, extends ...interface{}) (string, error) {
	var m *App
	domain := domain
	params := MapStr{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case MapStr:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	response, err := http.Get(uri)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	var responseByte []byte
	responseByte,err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	responseString := string(responseByte)
	if strings.Contains(response.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(responseString, *m) {
		return Get(path, extends...)
	}
	return responseString, nil
}

/**
 * @param path string,params param,app App,domain string
 * @author struggler
 * @description client postJson 注意：当传递mp的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostBody(path string, body []byte, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := MapStr{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case MapStr:
			params = a
		case Domain:
			domain = a
		case App:
			params["access_token"] = a.GetAccessToken().Token
			m = &a
		}
	}
	uri, _ := ReduceUrl(fmt.Sprintf("%s%s", domain, path), params)
	response, err := http.Post(uri,"",bytes.NewBuffer(body))
	defer response.Body.Close()
	if err != nil {
		return []byte(""), err
	}
	responseByte,err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte(""), err
	}
	if strings.Contains(response.Header.Get("Content-Type"), "application/json") && m != nil && checkTokenExpired(string(responseByte), *m) {
		return PostBody(path, body, extends...)
	}
	return responseByte, nil
}

/**
 * @param path,name string, file multipart.File,fileHeader *multipart.FileHeader,params param,app App,domain string
 * @author struggler
 * @description client postBufferFile 注意：当传递mp的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostBufferFile(path, name string, file multipart.File, fileHeader *multipart.FileHeader, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := MapStr{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case MapStr:
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
	part, err := writer.CreateFormFile(name, fileHeader.Filename)
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
		return PostBufferFile(path, name, file, fileHeader, extends...)
	}
	return responseByte, err
}

/**
 * @param path,name string, file io.Reader,filePath string,params param,app App,domain string
 * @author struggler
 * @description client postFilePath 注意：当传递mp的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostPathFile(path, name string, file io.Reader, filePath string, extends ...interface{}) ([]byte, error) {
	var m *App
	domain := domain
	params := MapStr{}
	for _, extend := range extends {
		switch a := extend.(type) {
		case MapStr:
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

func PostFileWithField(url, name string, file multipart.File, fileHeader *multipart.FileHeader, fields map[string]string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range fields {
		writer.WriteField(k, v)
	}
	part, err := writer.CreateFormFile(name, fileHeader.Filename)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//request.Header.Set("header", "header")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return body.Bytes(), err
	}
}

func PostFileNameWithField(url, name string, file io.Reader, fileName string, fields map[string]string) ([]byte, error) {
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

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//request.Header.Set("header", "header")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return body.Bytes(), err
	}
}