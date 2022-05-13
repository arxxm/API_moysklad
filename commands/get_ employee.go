package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	URLEmployee = "https://online.moysklad.ru/api/remap/1.2/entity/employee"
	username    = "admin@arxxm"
	password    = "2245core"
)

type Token struct {
	Access_token string `json:"access_token"`
}

func Authorization(url, method string) (*Token, error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return &Token{}, fmt.Errorf("Got error: %s", err.Error())
	}

	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return &Token{}, fmt.Errorf("Got error: %s", err.Error())
	}
	defer response.Body.Close()

	sliceBytes := make([]byte, 100)
	for {
		sliceBytes = sliceBytes[:cap(sliceBytes)]
		n, err := response.Body.Read(sliceBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
		sliceBytes = sliceBytes[:n]
	}

	sliceBytes = bytes.Trim(sliceBytes, "\x00")

	s := string(sliceBytes)
	data := Token{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return &Token{}, err
	}
	// fmt.Println("Token is: ", data.Access_token)

	return &data, nil

}

func GetListEmplyees(access_token string) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", URLEmployee, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", access_token)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	sliceBytes := make([]byte, 10000)
	for {
		sliceBytes = sliceBytes[:cap(sliceBytes)]
		n, err := res.Body.Read(sliceBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
		sliceBytes = sliceBytes[:n]
	}

	// sliceBytes = bytes.Trim(sliceBytes, "\x00")

	s := string(sliceBytes)
	fmt.Println(s)

	return nil
}

func CreateNewEmployee(access_token string) error {

	var jsonStr = []byte(`{
            "firstName": "Дарья",
            "middleName": "Тригубчак",
            "lastName": "Алексеевна",
            "inn": "222490425273",
            "position": "Директор",
            "phone": "+7(777)888-7777",
            "description": "Новое описание"}`)

	req, err := http.NewRequest("POST", URLEmployee, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", access_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}

func ChangeEmployee(access_token string) error {
	s := fmt.Sprintf("%s%s", URLEmployee, "/97de4265-d2d5-11ec-0a80-06a20013f6ef")

	var jsonStr = []byte(`{
		"firstName": "Даша",
		"middleName": "Алексеевна",
		"lastName": "Тригубчак",
		"inn": "222490425273",
		"position": "Директор",
		"phone": "+7(999)878-7878",
		"description": "new descriprion"}`)

	req, err := http.NewRequest("PUT", s, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", access_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	return nil
}

func DeleteEmployee(access_token string) error {
	s := fmt.Sprintf("%s%s", URLEmployee, "/edf33e95-d058-11ec-0a80-08ab00a1ea60")

	req, err := http.NewRequest("DELETE", s, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", access_token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
