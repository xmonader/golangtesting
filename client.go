package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TodosClient struct {
	address string
	client  *http.Client
}

// didn't allow here to pass the client, maybe we want a custom one
// either make it exported or pass it to the New function or use Options pattern
func NewClient(address string) *TodosClient {
	return &TodosClient{address: address, client: http.DefaultClient}

}

// Providing some other ways to provide a client to be injected
func NewClientWithHTTPClient(address string, cl *http.Client) *TodosClient {
	return &TodosClient{address: address, client: cl}
}

type Option = func(c *TodosClient)

func WithHTTPClient(cl *http.Client) Option {
	return func(c *TodosClient) {
		c.client = cl
	}
}

func NewClientWithOptions(address string, opts ...Option) *TodosClient {
	cl := &TodosClient{address: address}
	for _, opt := range opts {
		opt(cl)
	}
	return cl
}

func (c *TodosClient) ListTodos() (map[int]Todo, error) {
	req, err := http.NewRequest(http.MethodGet, c.address, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result map[int]Todo
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *TodosClient) GetTodo(id int) (Todo, error) {
	res, err := c.client.Get(fmt.Sprintf("%s/%d", c.address, id))
	if err != nil {
		return Todo{}, err
	}

	defer res.Body.Close()
	var result Todo
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *TodosClient) NewTodo(title string) (Todo, error) {

	postBody, _ := json.Marshal(map[string]string{
		"title": title,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := c.client.Post(c.address, "application/json", responseBody)
	if err != nil {
		return Todo{}, err
	}
	defer resp.Body.Close()

	var result Todo
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (c *TodosClient) UpdateTodo(id int, title string) (Todo, error) {
	putBody, _ := json.Marshal(map[string]string{
		"title": title,
	})

	bodyBytes := bytes.NewBuffer(putBody)
	//Leverage Go's HTTP Post function to make request
	req, err := http.NewRequest(http.MethodPut, c.address, bodyBytes)
	if err != nil {
		return Todo{}, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return Todo{}, err
	}
	defer res.Body.Close()

	var result Todo
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
