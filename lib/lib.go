package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
)

func GenerateSignatureOrder(MID string, Secret string, RefId string) string {
	formula := fmt.Sprintf("%v:%v:%v", MID, Secret, RefId)

	data := []byte(formula)
	hash := fmt.Sprintf("%x", md5.Sum(data))

	return hash
}

func GenerateSignatureDefault(MID string, Secret string) string {
	formula := fmt.Sprintf("%v:%v", MID, Secret)

	data := []byte(formula)
	hash := fmt.Sprintf("%x", md5.Sum(data))

	return hash
}

func ToIDR(nominal int) string {
	//konversi int ke format string rupiah
	nominalHuman := humanize.Comma(int64(nominal))
	result := fmt.Sprintf("Rp %v", nominalHuman)
	result = strings.ReplaceAll(result, ",", ".")
	return result
}

func IntToString(val int) string {
	s := strconv.Itoa(val)
	return s
}

func StringToInt(val string) int {
	data, err := strconv.Atoi(val)
	err = err
	return data
}

func SendPostJson(url string, data fiber.Map) (string, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		return "", err
	}
	defer resp.Body.Close()
	hasil, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	hasilString := string(hasil)
	return hasilString, nil
}

func SendGet(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	sb := string(body)
	return sb, nil
}
