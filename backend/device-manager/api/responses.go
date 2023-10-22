package api

import (
	"context"
	"net/http"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	"github.com/goccy/go-json"
)

var (
	_ response = (*createDevicesResp)(nil)
	_ response = (*getDeviceResp)(nil)
	_ response = (*getDevicesResp)(nil)
	_ response = (*updateDeviceResp)(nil)
	_ response = (*deleteDeviceResp)(nil)
)

type response interface {
	Code() int
	Empty() bool
}

func encodeResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	if ar, ok := resp.(response); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(ar.Code())

		if ar.Empty() {
			return nil
		}
	}

	return json.NewEncoder(w).Encode(resp)
}

type createDevicesResp struct {
	created bool
}

func (resp createDevicesResp) Code() int {
	if resp.created {
		return http.StatusCreated
	}

	return http.StatusBadRequest
}

func (resp createDevicesResp) Empty() bool {
	return true
}

type getDeviceResp struct {
	Device devicemanager.Device `json:",inline"`
}

func (resp getDeviceResp) Code() int {
	if resp.Device.ID != "" {
		return http.StatusOK
	}

	return http.StatusNotFound
}

func (resp getDeviceResp) Empty() bool {
	return resp.Device.ID == ""
}

type getDevicesResp struct {
	devicemanager.Page `json:",inline"`
}

func (resp getDevicesResp) Code() int {
	return http.StatusOK
}

func (resp getDevicesResp) Empty() bool {
	return false
}

type updateDeviceResp struct {
	updated bool
}

func (resp updateDeviceResp) Code() int {
	if resp.updated {
		return http.StatusAccepted
	}

	return http.StatusBadRequest
}

func (resp updateDeviceResp) Empty() bool {
	return true
}

type deleteDeviceResp struct {
	deleted bool
}

func (resp deleteDeviceResp) Code() int {
	if resp.deleted {
		return http.StatusAccepted
	}

	return http.StatusBadRequest
}

func (resp deleteDeviceResp) Empty() bool {
	return true
}
