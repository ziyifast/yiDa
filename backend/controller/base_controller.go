package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/aobco/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"ziyi.com/yiDa/common/constant"
	"ziyi.com/yiDa/common/response"
)

type BaseController struct {
	Ctx iris.Context
}

func (c *BaseController) Param(paramName string) string {
	return c.Ctx.Params().Get(paramName)
}
func (c *BaseController) PostValue(paramName string) string {
	return c.Ctx.PostValue(paramName)
}

func (c *BaseController) ParamInt64Default(paramName string, defaultValue int64) int64 {
	return c.Ctx.Params().GetInt64Default(paramName, defaultValue)
}
func (c *BaseController) ParamInt64(paramName string) (int64, error) {
	return c.Ctx.Params().GetInt64(paramName)
}
func (c *BaseController) ParamInt32Default(paramName string, defaultValue int32) int32 {
	return c.Ctx.Params().GetInt32Default(paramName, defaultValue)
}
func (c *BaseController) ParamInt32(paramName string) (int32, error) {
	return c.Ctx.Params().GetInt32(paramName)
}
func (c *BaseController) ParamIntDefault(paramName string, defaultValue int) int {
	return c.Ctx.Params().GetIntDefault(paramName, defaultValue)
}
func (c *BaseController) ParamInt(paramName string) (int, error) {
	return c.Ctx.Params().GetInt(paramName)
}

func (c *BaseController) ReadParamJson(paramName string, result interface{}) error {
	paramJson := c.Ctx.FormValue(paramName)
	if err := json.Unmarshal([]byte(paramJson), result); err != nil {
		log.Errorf("unmarshal request param[%s] fail, param body [%v] error [%v]", paramName, paramJson, err)
		return err
	}
	return nil
}

func (c *BaseController) GetIntFormValue(key string, defaultValue int) int {
	value := c.Ctx.FormValue(key)
	if len(value) == 0 {
		return defaultValue
	}
	atoi, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return atoi
}

func (c *BaseController) ReadJSON(jsonObject interface{}) error {
	r := c.Ctx.Request()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("reading request body error %v", err)
		return err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, body, "", "\t")
	if err != nil {
		log.Warnf("indent json error %v", err)

	}
	if len(body) > 0 {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	if err := c.Ctx.ReadJSON(jsonObject); err != nil {
		log.Errorf("parsing request body error %v", err)
		return err
	}
	return nil
}

func (c *BaseController) ReadXML(jsonObject interface{}) error {
	r := c.Ctx.Request()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("reading request body error %v", err)
		return err
	}
	if len(body) > 0 {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	if err := c.Ctx.ReadXML(jsonObject); err != nil {
		log.Errorf("parsing request body error %v", err)
		return err
	}
	return nil
}
func (c *BaseController) ReadString() (string, error) {
	r := c.Ctx.Request()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("reading request body error %v", err)
		return "", err
	}
	content := string(body)
	if len(body) > 0 {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	return content, nil
}

func (c *BaseController) response(responseCode response.Code, msg string, content interface{}) mvc.Result {
	return commonResp(msg, http.StatusOK, responseCode, content)
}
func (c *BaseController) Ok(content interface{}) mvc.Result {
	return commonResp(response.SuccessMsg, http.StatusOK, response.ReturnCodeSuccess, content)
}

func (c *BaseController) OkWithMsg(str string, args ...interface{}) mvc.Result {
	return commonResp(response.SuccessMsg, http.StatusOK, response.ReturnCodeSuccess, nil)
}

func (c *BaseController) Failed(errMsg string, args ...interface{}) mvc.Result {
	return commonResp(errMsg, http.StatusInternalServerError, response.ReturnCodeError, nil)
}

func (c *BaseController) Forbidden() mvc.Result {
	return commonResp(http.StatusText(http.StatusForbidden), http.StatusForbidden, response.ReturnCodeError, nil)
}

func (c *BaseController) BadRequest() mvc.Result {
	return commonResp(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, response.ReturnCodeError, nil)
}

func (c *BaseController) BadRequestWithMsg(errMsg string) mvc.Result {
	return commonResp(errMsg, http.StatusBadRequest, response.ReturnCodeError, nil)
}

func (c *BaseController) SystemInternalError() mvc.Result {
	return commonResp(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, response.ReturnCodeError, nil)
}

func (c *BaseController) Unauthorized() mvc.Result {
	return commonResp(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, response.ReturnCodeError, nil)
}

func (c *BaseController) SystemInternalErrorWithMsg(errMsg string) mvc.Result {
	return commonResp(errMsg, http.StatusInternalServerError, response.ReturnCodeError, nil)
}

func (c *BaseController) NotFoundError() mvc.Result {
	return commonResp(http.StatusText(http.StatusNotFound), http.StatusNotFound, response.ReturnCodeError, nil)
}

func commonResp(errMsg string, httpCode int, returnCode response.Code, content interface{}) mvc.Response {
	payload := &response.JsonReponse{
		Code:    returnCode,
		Msg:     errMsg,
		Content: content,
	}

	contentDetail, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("marshal json response error %v", err)
	}

	return mvc.Response{
		Code:        httpCode,
		Content:     contentDetail,
		ContentType: constant.ContentTypeJson,
	}
}
