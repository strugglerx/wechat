package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

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
func checkTokenExpired(errcode int, m App) bool {
	if value, ok := expiredToken[strconv.Itoa(errcode)]; ok {
		m.GetAccessToken(true)
		return value
	}
	return false
}

func checkJsonResponse(resp *http.Response) bool {
	return strings.Contains(resp.Header.Get("Content-Type"), "application/json")
}

func checkCodeError(errcode int) bool {
	return errcode != 0
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
func Get(path string, extends ...interface{}) (Response, error) {
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
	var respByte Response
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if checkJsonResponse(resp) && m != nil {
		jsonResponse := JsonResponse{}
		err := respByte.Unmarshal(&jsonResponse)
		if err != nil {
			return nil, err
		}
		if checkTokenExpired(jsonResponse.Errcode, *m) {
			return Get(path, extends...)
		}
		if checkCodeError(jsonResponse.Errcode) {
			return nil, showError{errorCode: jsonResponse.Errcode, errorMsg: errors.New(jsonResponse.ErrMsg)}
		}
	}
	return respByte, nil
}

//PostBody
/**
 * @param path string,params param,app App,domain string
 * @author struggler
 * @description client postJson 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostBody(path string, body []byte, extends ...interface{}) (Response, error) {
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
	resp, err := http.Post(uri, "", bytes.NewReader(body))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var respByte Response
	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if checkJsonResponse(resp) && m != nil {
		jsonResponse := JsonResponse{}
		err := respByte.Unmarshal(&jsonResponse)
		if err != nil {
			return nil, err
		}
		if checkTokenExpired(jsonResponse.Errcode, *m) {
			return PostBody(path, body, extends...)
		}
		if checkCodeError(jsonResponse.Errcode) {
			return nil, showError{errorCode: jsonResponse.Errcode, errorMsg: errors.New(jsonResponse.ErrMsg)}
		}
	}
	return respByte, nil
}

//PostBufferFile
/**
 * @param path,name string, file multipart.File,fileName *multipart.fileName,params param,app App,domain string
 * @author struggler
 * @description client postBufferFile 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostBufferFile(path, name string, file io.Reader, fileName string, extends ...interface{}) (Response, error) {
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
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	var respByte Response
	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if checkJsonResponse(resp) && m != nil {
		jsonResponse := JsonResponse{}
		err := respByte.Unmarshal(&jsonResponse)
		if err != nil {
			return nil, err
		}
		if checkTokenExpired(jsonResponse.Errcode, *m) {
			return PostBufferFile(path, name, file, fileName, extends...)
		}
		if checkCodeError(jsonResponse.Errcode) {
			return nil, showError{errorCode: jsonResponse.Errcode, errorMsg: errors.New(jsonResponse.ErrMsg)}
		}
	}
	return respByte, err
}

//PostPathFile
/**
 * @param path,name string, file io.Reader,filePath string,params param,app App,domain string
 * @author struggler
 * @description client postFilePath 注意：当传递App的时候会自动获取token 并添加到 param里
 * @date 10:52 下午 2021/2/23
 * @return string,error
 **/
func PostPathFile(path, name string, file io.Reader, filePath string, extends ...interface{}) (Response, error) {
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
	var respByte Response
	respByte = responseBody.Bytes()
	if checkJsonResponse(resp) && m != nil {
		jsonResponse := JsonResponse{}
		err := respByte.Unmarshal(&jsonResponse)
		if err != nil {
			return nil, err
		}
		if checkTokenExpired(jsonResponse.Errcode, *m) {
			return PostPathFile(path, name, file, filePath, extends...)
		}
		if checkCodeError(jsonResponse.Errcode) {
			return nil, showError{errorCode: jsonResponse.Errcode, errorMsg: errors.New(jsonResponse.ErrMsg)}
		}
	}
	return respByte, err
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
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respByte Response

	respByte = buf.Bytes()
	if checkJsonResponse(resp) && m != nil {
		jsonResponse := JsonResponse{}
		err := respByte.Unmarshal(&jsonResponse)
		if err != nil {
			return nil, err
		}
		if checkTokenExpired(jsonResponse.Errcode, *m) {
			return PostBufferFileWithField(path, name, file, fileName, fields, extends...)
		}
		if checkCodeError(jsonResponse.Errcode) {
			return nil, showError{errorCode: jsonResponse.Errcode, errorMsg: errors.New(jsonResponse.ErrMsg)}
		}
	}
	return respByte, nil
}
