package acceptanceservice

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerrror"
	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"
	"github.com/CarosDrean/go-json"

	"pass-trougth/infraestructure/service/auth"
	"pass-trougth/model"
)

const _descriptionDefault = "Aceptado"

const _queryParamForDelete = "?accion=cambioestado"

const (
	_isResetTokenDefault = false
	_retriesDefault      = 1
	_customTag           = "json-out"
)

const (
	_headerAuthorization    = "Authorization"
	_headerContentType      = "Content-Type"
	_headerContentTypeValue = "application/json"
	_headerSubscriptionKey  = "Ocp-Apim-Subscription-Key"
)

type AcceptanceService struct {
	config model.ServiceConfig
	auth   auth.Auth
}

func New(config model.ServiceConfig) AcceptanceService {
	return AcceptanceService{config: config, auth: auth.New(config.UrlAuth)}
}

func (a AcceptanceService) CreateOrUpdate(ctx tracingModel.Context, input model.AcceptanceService, headers map[string]string) (model.Message, error) {
	body, err := a.doRequest(ctx, http.MethodPost, a.config.Url, headers, input, _isResetTokenDefault, _retriesDefault)
	if err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.CreateOrUpdate().doRequest(): %w", err)
	}

	return model.Message{
		Description: string(body),
		Message:     _descriptionDefault,
	}, nil
}

func (a AcceptanceService) Delete(ctx tracingModel.Context, input model.RequestDelete, headers map[string]string) (model.Message, error) {
	body, err := a.doRequest(ctx, http.MethodPut, a.config.Url+_queryParamForDelete, headers, input, _isResetTokenDefault, _retriesDefault)
	if err != nil {
		return model.Message{}, fmt.Errorf("service.acceptance.Delete().doRequest(): %w", err)
	}

	return model.Message{
		Description: string(body),
		Message:     _descriptionDefault,
	}, nil
}

func (a AcceptanceService) Retrieve(ctx tracingModel.Context, inputRequest model.RetrieveRequest, headers map[string]string) (model.RetrieveResponse, error) {
	body, err := a.doRequest(ctx, http.MethodPost, a.config.UrlRetrieve, headers, inputRequest, _isResetTokenDefault, _retriesDefault)
	if err != nil {
		return model.RetrieveResponse{}, fmt.Errorf("service.acceptance.Retrieve().doRequest(): %w", err)
	}

	var response model.RetrieveResponse
	if err := json.UnmarshalWithCustomTag(body, &response, _customTag); err != nil {
		return model.RetrieveResponse{}, fmt.Errorf("service.acceptance.Retrieve().json.UnmarshalWithCustomTag(): %w", err)
	}

	return response, nil
}

func (a AcceptanceService) doRequest(ctx tracingModel.Context, method, url string, headers map[string]string, payloadData any, isResetToken bool, numberRetries int) ([]byte, error) {
	token, err := a.auth.GetToken(ctx, model.Auth{IsResetToken: isResetToken})
	if err != nil {
		return nil, fmt.Errorf("auth.GetToken(): %w", err)
	}

	payload, err := json.MarshalWithCustomTag(payloadData, _customTag)
	if err != nil {
		return nil, fmt.Errorf("json.MarshalWithCustomTag(): %w", err)
	}

	timeInit := time.Now()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	ctx.AddInfoWithGenericKey(tracingModel.TypeInfoProcessRequestService, tracingModel.InfoProcess{
		Name:    "AcceptanceService",
		Time:    timeInit,
		Where:   "service.acceptanceservice.doRequest()",
		Payload: string(payload),
		Error:   err,
	})
	if err != nil {
		return nil, fmt.Errorf("NewRequest(): %w", err)
	}

	req.Header.Add(_headerContentType, _headerContentTypeValue)
	req.Header.Add(_headerSubscriptionKey, a.config.SubscriptionKey)
	req.Header.Add(_headerAuthorization, "BEARER "+token.AccessToken)

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
