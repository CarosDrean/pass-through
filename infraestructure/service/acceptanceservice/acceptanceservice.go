package acceptanceservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerrror"
	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"
	authorization "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-oauth.git"

	"pass-trougth/model"
)

const (
	_isResetTokenDefault = false
	_reintentDefault     = 1
)

const (
	_headerAuthorization    = "Authorization"
	_headerContentType      = "Content-Type"
	_headerContentTypeValue = "application/json"
	_headerSubscriptionKey  = "Ocp-Apim-Subscription-Key"
)

type AcceptanceService struct {
	dataForToken *strings.Reader

	config model.ServiceConfig
	auth   authorization.Auth
}

func New(config model.ServiceConfig) AcceptanceService {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", config.User)
	data.Set("password", config.Password)

	return AcceptanceService{
		dataForToken: strings.NewReader(data.Encode()),
		config:       config,
		auth: authorization.New(authorization.Config{
			Url:      config.UrlAuth,
			User:     config.UserBasic,
			Password: config.PasswordBasic,
			Headers:  map[string]string{_headerSubscriptionKey: config.SubscriptionKey},
		})}
}

func (a AcceptanceService) Create(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error) {
	output := model.AcceptanceServiceOutput(input)

	body, err := a.doRequest(ctx, http.MethodPost, a.config.Url, headers, output, _isResetTokenDefault, _reintentDefault)
	if err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Create().doRequest(): %w", err)
	}

	var response model.Message
	if err := json.Unmarshal(body, &response); err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Create().json.Unmarshal(): %w", err)
	}

	return response, nil
}

func (a AcceptanceService) Update(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error) {
	output := model.AcceptanceServiceOutput(input)

	body, err := a.doRequest(ctx, http.MethodPut, a.config.Url, headers, output, _isResetTokenDefault, _reintentDefault)
	if err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Update().doRequest(): %w", err)
	}

	var response model.Message
	if err := json.Unmarshal(body, &response); err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Update().json.Unmarshal(): %w", err)
	}

	return response, nil
}

func (a AcceptanceService) Delete(ctx tracingModel.Context, input model.RequestDeleteInput, headers map[string]string) (model.Message, error) {
	output := model.RequestDeleteOutput(input)

	body, err := a.doRequest(ctx, http.MethodDelete, a.config.Url, headers, output, _isResetTokenDefault, _reintentDefault)
	if err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Delete().doRequest(): %w", err)
	}

	var response model.Message
	if err := json.Unmarshal(body, &response); err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Delete().json.Unmarshal(): %w", err)
	}

	return response, nil
}

func (a AcceptanceService) Retrieve(ctx tracingModel.Context, inputRequest model.RetrieveInputRequest, headers map[string]string) (model.RetrieveOutputResponse, error) {
	outputRequest := model.RetrieveOutputRequest{}
	for _, serviceInput := range inputRequest.Documents {
		outputRequest.Documents = append(outputRequest.Documents, model.FieldRetrieveOutputRequest(serviceInput))
	}

	body, err := a.doRequest(ctx, http.MethodPost, a.config.UrlRetrieve, headers, outputRequest, _isResetTokenDefault, _reintentDefault)
	if err != nil {
		return model.RetrieveOutputResponse{}, fmt.Errorf("service.acceptance.Retrieve().doRequest(): %w", err)
	}

	var inputResponse model.RetrieveInputResponse
	if err := json.Unmarshal(body, &inputResponse); err != nil {
		return model.RetrieveOutputResponse{}, fmt.Errorf("service.acceptance.Retrieve().json.Unmarshal(): %w", err)
	}

	outputResponse := model.RetrieveOutputResponse{}
	for _, serviceInput := range inputResponse.Data {
		outputResponse.Data = append(outputResponse.Data, model.FieldRetrieveOutputResponse(serviceInput))
	}

	return outputResponse, nil
}

func (a AcceptanceService) doRequest(ctx tracingModel.Context, method, url string, headers map[string]string, payloadData any, isResetToken bool, numberRetries int) ([]byte, error) {
	dataTokenBuffer := &bytes.Buffer{}
	_, err := dataTokenBuffer.ReadFrom(a.dataForToken)
	if err != nil {
		return nil, fmt.Errorf("dataTokenBuffer.ReadFrom(): %w", err)
	}

	token, err := a.auth.GetToken(dataTokenBuffer.Bytes(), isResetToken)
	if err != nil {
		return nil, fmt.Errorf("auth.GetToken(): %w", err)
	}

	payload, err := json.Marshal(payloadData)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(): %w", err)
	}

	timeInit := time.Now()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	ctx.AddInfoWithGenericKey(tracingModel.TypeInfoProcessRequestService, tracingModel.InfoProcess{
		Name:    "AcceptanceService",
		Time:    timeInit,
		Where:   "service.acceptanceservice.doRequest()",
		Payload: payload,
		Error:   err,
	})
	if err != nil {
		return nil, fmt.Errorf("NewRequest(): %w", err)
	}

	req.Header.Add(_headerContentType, _headerContentTypeValue)
	req.Header.Add(_headerSubscriptionKey, a.config.SubscriptionKey)
	req.Header.Add(_headerAuthorization, "BEARER "+token)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading Body: %v", err)
	}

	ctx.AddInfoWithGenericKey(tracingModel.TypeInfoProcessResponseService, tracingModel.InfoProcess{
		Name:        "AcceptanceService",
		Time:        time.Now(),
		ElapsedNano: time.Since(timeInit).Nanoseconds(),
		Where:       "service.acceptanceservice.doRequest()",
		Response:    string(bodyBytes),
	})

	// the token may expire, for which the need arises to retry if the response status is 401
	if res.StatusCode == http.StatusUnauthorized && numberRetries > 0 {
		numberRetries--
		body, err := a.doRequest(ctx, method, url, headers, payloadData, true, numberRetries)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		customErr := customerror.NewError()
		customErr.SetStatusHTTP(res.StatusCode)
		customErr.SetAPIMessage(customerror.StatusCodeMessage(bodyBytes))

		return nil, customErr
	}

	return bodyBytes, nil
}
