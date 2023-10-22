package sdk

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	"github.com/goccy/go-json"
)

const (
	deviceEndpoint = "users"
)

type SDK interface {
	CreateDevice(device devicemanager.Device) (devicemanager.Device, error)
	GetDevice(id string) (devicemanager.Device, error)
	GetDevices() (devicemanager.Page, error)
	UpdateDevice(id string, device devicemanager.Device) (devicemanager.Device, error)
	DeleteDevice(id string) error
}

type sdk struct {
	devicemanagerURL string
	token            string
	client           *http.Client
}

func New(url, token string) SDK {
	return &sdk{
		devicemanagerURL: url,
		token:            token,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (s *sdk) CreateDevice(device devicemanager.Device) (devicemanager.Device, error) {
	var devices = struct {
		Devices []devicemanager.Device `json:"devices"`
	}{
		Devices: []devicemanager.Device{device},
	}

	data, err := json.Marshal(devices)
	if err != nil {
		return devicemanager.Device{}, fmt.Errorf("failed to marshal device: %w", err)
	}

	url := fmt.Sprintf("%s/%s", s.devicemanagerURL, deviceEndpoint)

	body, err := s.processRequest(http.MethodPost, url, data, http.StatusCreated)
	if err != nil {
		return devicemanager.Device{}, err
	}

	device = devicemanager.Device{}
	if err := json.Unmarshal(body, &device); err != nil {
		return devicemanager.Device{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return device, nil
}

func (s *sdk) GetDevice(id string) (devicemanager.Device, error) {
	url := fmt.Sprintf("%s/%s/%s", s.devicemanagerURL, deviceEndpoint, id)

	body, err := s.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return devicemanager.Device{}, err
	}

	device := devicemanager.Device{}
	if err := json.Unmarshal(body, &device); err != nil {
		return devicemanager.Device{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return device, nil

}

func (s *sdk) GetDevices() (devicemanager.Page, error) {
	url := fmt.Sprintf("%s/%s", s.devicemanagerURL, deviceEndpoint)

	body, err := s.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return devicemanager.Page{}, err
	}

	page := devicemanager.Page{}
	if err := json.Unmarshal(body, &page); err != nil {
		return devicemanager.Page{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return page, nil
}

func (s *sdk) UpdateDevice(id string, device devicemanager.Device) (devicemanager.Device, error) {
	data, err := json.Marshal(device)
	if err != nil {
		return devicemanager.Device{}, fmt.Errorf("failed to marshal device: %w", err)
	}

	url := fmt.Sprintf("%s/%s/%s", s.devicemanagerURL, deviceEndpoint, id)

	body, err := s.processRequest(http.MethodPut, url, data, http.StatusAccepted)
	if err != nil {
		return devicemanager.Device{}, err
	}

	device = devicemanager.Device{}
	if err := json.Unmarshal(body, &device); err != nil {
		return devicemanager.Device{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return device, nil
}

func (s *sdk) DeleteDevice(id string) error {
	url := fmt.Sprintf("%s/%s/%s", s.devicemanagerURL, deviceEndpoint, id)

	_, err := s.processRequest(http.MethodDelete, url, nil, http.StatusAccepted)
	if err != nil {
		return err
	}

	return nil
}

func (s *sdk) processRequest(method, url string, data []byte, expectedRespCode int) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return []byte{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	if s.token != "" {
		token := "Bearer " + s.token
		req.Header.Set("Authorization", token)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedRespCode {
		return []byte{}, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
