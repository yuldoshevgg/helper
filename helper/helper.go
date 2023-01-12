package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			oldsize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
			}
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func MakeRequests(method, url, username, password string, reqBody []byte) (int, interface{}, error) {
	var response interface{}
	client := &http.Client{}

	fmt.Println("Make request body: ", string(reqBody))

	reqReader := bytes.NewReader(reqBody)
	req, err := http.NewRequest(method, url, reqReader)
	if err != nil {
		return http.StatusInternalServerError, "", errors.Wrap(err, "error while making request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, "", errors.Wrap(err, "error while doing request")
	}

	if resp == nil {
		err := errors.New("http response nil")
		return http.StatusInternalServerError, response, errors.Wrap(err, "error while connecting 1C")
	}

	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error while reading body")
	}

	fmt.Println("\n\nRESP: \n\n", string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error while unmarshaling body")
	}

	return resp.StatusCode, response, nil
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func IsIntegral(val float64) bool {
	return val == float64(int(val))
}

func HumanReadableNumber(val float64) (hrnumber float64) {

	return
}
