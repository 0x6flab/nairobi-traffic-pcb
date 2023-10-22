package api

import (
	"errors"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
)

var (
	errMissingToken = errors.New("missing token")
	errMissingID    = errors.New("missing device ID")
)

type createDevicesReq struct {
	token   string
	Devices []devicemanager.Device `json:"devices"`
}

func (req createDevicesReq) validate() error {
	if req.token == "" {
		return errMissingToken
	}

	return nil
}

type getDeviceReq struct {
	token string
	id    string
}

func (req getDeviceReq) validate() error {
	if req.token == "" {
		return errMissingToken
	}
	if req.id == "" {
		return errMissingID
	}

	return nil
}

type getDevicesReq struct {
	token string
	page  devicemanager.PageMetadata
}

func (req getDevicesReq) validate() error {
	if req.token == "" {
		return errMissingToken
	}

	return nil
}

type updateDeviceReq struct {
	token  string
	Device devicemanager.Device `json:",inline"`
}

func (req updateDeviceReq) validate() error {
	if req.token == "" {
		return errMissingToken
	}
	if req.Device.ID == "" {
		return errMissingID
	}

	return nil
}

type deleteDeviceReq struct {
	token string
	id    string
}

func (req deleteDeviceReq) validate() error {
	if req.token == "" {
		return errMissingToken
	}
	if req.id == "" {
		return errMissingID
	}

	return nil
}
