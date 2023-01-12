package helper

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func SaleRequest(baseUrl, username, password string, req []byte) (int, interface{}, error) {

	url := fmt.Sprintf("%v/sale", baseUrl)
	code, resp, err := MakeRequests(http.MethodPost, url, username, password, req)
	if err != nil {
		return code, nil, errors.Wrap(err, "error while doing OneC sale request")
	}

	return code, resp, nil
}

func RefundRequest(baseUrl, username, password string, req []byte) (int, interface{}, error) {

	url := fmt.Sprintf("%v/refund", baseUrl)
	code, resp, err := MakeRequests(http.MethodPost, url, username, password, req)
	if err != nil {
		return code, nil, errors.Wrap(err, "error while doing OneC refund request")
	}

	return code, resp, nil
}

func PaymentRequest(baseUrl, username, password string, req []byte) (int, interface{}, error) {

	url := fmt.Sprintf("%v/payment", baseUrl)
	code, resp, err := MakeRequests(http.MethodPost, url, username, password, req)
	if err != nil {
		return code, nil, errors.Wrap(err, "error while doing OneC payment request")
	}

	return code, resp, nil
}
