package service

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

func raw() ([]Data, error) {
	r, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var payLoad Payload
	err = json.Unmarshal(r, &payLoad.Data)

	if err != nil {
		return nil, err
	}
	return payLoad.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()

	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetById(idx int) (any, error) {
	data, err := raw()

	if err != nil {
		return nil, err
	}

	if idx > len(data) {
		res := make([]string, 0)
		return res, nil
	}
	return data[idx], nil

}
