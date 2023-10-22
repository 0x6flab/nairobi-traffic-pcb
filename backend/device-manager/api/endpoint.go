package api

import (
	"context"

	devicemanager "github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager"
	"github.com/go-kit/kit/endpoint"
)

func createDevicesEndpoint(svc devicemanager.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createDevicesReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.CreateDevice(ctx, req.token, req.Devices...); err != nil {
			return createDevicesResp{created: false}, err
		}

		return createDevicesResp{created: true}, nil
	}
}

func getDeviceEndpoint(svc devicemanager.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getDeviceReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		device, err := svc.GetDevice(ctx, req.token, req.id)
		if err != nil {
			return getDeviceResp{}, err
		}

		return getDeviceResp{Device: device}, nil
	}
}

func getDevicesEndpoint(svc devicemanager.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getDevicesReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		devices, err := svc.GetDevices(ctx, req.token, req.page)
		if err != nil {
			return getDevicesResp{}, err
		}

		return getDevicesResp{devices}, nil
	}
}

func updateDeviceEndpoint(svc devicemanager.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateDeviceReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.UpdateDevice(ctx, req.token, req.Device); err != nil {
			return updateDeviceResp{updated: false}, err
		}

		return updateDeviceResp{updated: true}, nil
	}
}

func deleteDeviceEndpoint(svc devicemanager.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteDeviceReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		if err := svc.DeleteDevice(ctx, req.token, req.id); err != nil {
			return deleteDeviceResp{deleted: false}, err
		}

		return deleteDeviceResp{deleted: true}, nil
	}
}
