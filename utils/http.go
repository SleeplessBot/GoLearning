package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// make http json request and parse json response into resp
// resp should be pointer type
func HttpJsonRequest(method string, reqUrl string, queries map[string]string, req interface{}, headers map[string]string, resp interface{}) (err error) {
	var httpReq *http.Request
	if req == nil {
		httpReq, err = http.NewRequest(method, reqUrl, nil)
	} else {
		var json_data []byte
		json_data, err = json.Marshal(req)
		if err != nil {
			return nil
		}
		httpReq, err = http.NewRequest(method, reqUrl, bytes.NewBuffer(json_data))
	}
	if err != nil {
		return err
	}

	if len(queries) > 0 {
		urlValues := httpReq.URL.Query()
		for k, v := range queries {
			urlValues.Add(k, v)
		}
		httpReq.URL.RawQuery = urlValues.Encode()
	}

	httpReq.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		httpReq.Header.Add(k, v)
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code %d", httpResp.StatusCode)
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if resp != nil {
		err = json.Unmarshal(body, resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func HttpGet(reqUrl string, resp interface{}) error {
	return HttpJsonRequest("GET", reqUrl, nil, nil, nil, resp)
}

func HttpGetQ(reqUrl string, queries map[string]string, resp interface{}) error {
	return HttpJsonRequest("GET", reqUrl, queries, nil, nil, resp)
}

func HttpPostJson(reqUrl string, req interface{}, resp interface{}) error {
	return HttpJsonRequest("POST", reqUrl, nil, req, nil, resp)
}

// with headers
func HttpPostJsonH(reqUrl string, req interface{}, headers map[string]string, resp interface{}) error {
	return HttpJsonRequest("POST", reqUrl, nil, req, headers, resp)
}

func HttpPutJson(reqUrl string, req interface{}, resp interface{}) error {
	return HttpJsonRequest("PUT", reqUrl, nil, req, nil, resp)
}

// make http form request and parse json response into resp
// resp should be pointer type
func HttpFormRequest(method string, reqUrl string, params map[string]string, files map[string]io.Reader, resp interface{}) (err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for fileName, fileReader := range files {
		part, err := writer.CreateFormFile(fileName, fileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, fileReader)
		if err != nil {
			return err
		}
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	httpResp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code %d", httpResp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if resp != nil {
		err = json.Unmarshal(respBody, resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func HttpPostForm(reqUrl string, params map[string]string, files map[string]io.Reader, resp interface{}) error {
	return HttpFormRequest("POST", reqUrl, params, files, resp)
}

func HttpDownloadFile(reqUrl string, filePath string) error {
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code %d", httpResp.StatusCode)
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, httpResp.Body)
	return err
}

// download file to cas
func HttpDownloadFileC(reqUrl string) (string, error) {
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status code %d", httpResp.StatusCode)
	}

	return DefaultCas.Create(httpResp.Body)
}
