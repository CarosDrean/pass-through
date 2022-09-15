package auth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	customerror "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerrror"
	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"
	"github.com/CarosDrean/go-json"

	"pass-trougth/model"
)

type Auth struct {
	Url string
}

func New(url string) Auth {
	return Auth{Url: url}
}

func (a Auth) GetToken(ctx tracingModel.Context, auth model.Auth) (model.Token, error) {
	body, err := doRequest(ctx, http.MethodPost, a.Url, auth)
	if err != nil {
		return model.Token{}, fmt.Errorf("service.auth.doRequest(): %w", err)
	}

	var response model.Token
	if err := json.Unmarshal(body, &response); err != nil {
		return model.Token{}, fmt.Errorf("service.auth.json.Unmarshal(): %w", err)
	}

	return response, nil
}

func doRequest(ctx tracingModel.Context, method, url string, payloadData any) ([]byte, error) {
	payload, err := json.Marshal(payloadData)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(): %w", err)
	}

	timeInit := time.Now()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	ctx.AddInfoWithGenericKey(tracingModel.TypeInfoProcessRequestService, tracingModel.InfoProcess{
		Name:    "Auth",
		Time:    timeInit,
		Where:   "service.auth.doRequest()",
		Payload: string(payload),
		Error:   err,
	})
	if err != nil {
		return nil, fmt.Errorf("NewRequest(): %w", err)
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
		Name:        "Auth",
		Time:        time.Now(),
		ElapsedNano: time.Since(timeInit).Nanoseconds(),
		Where:       "service.auth.doRequest()",
		Response:    string(bodyBytes),
	})

	if res.StatusCode >= http.StatusMultipleChoices {
		customErr := customerror.NewError()
		customErr.SetStatusHTTP(res.StatusCode)
		customErr.SetAPIMessage(customerror.StatusCodeMessage(bodyBytes))

		return nil, customErr
	}

	return bodyBytes, nil
}
