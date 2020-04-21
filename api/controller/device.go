package controller

import (
	"gim/internal/logic/model"
	"gim/internal/logic/service"
	"gim/pkg/imerror"
)

func init() {
	g := Engine.Group("/device")
	g.POST("", handler(DeviceController{}.Register))
}

type DeviceController struct{}

// Register 设备注册
func (DeviceController) Register(c *context) {
	var device model.Device
	if c.ShouldBindJSON(&device) != nil {
		return
	}

	if device.AppId == 0 || device.Type == 0 || device.Brand == "" || device.Model == "" ||
		device.SystemVersion == "" || device.SDKVersion == "" {
		c.response(nil, imerror.ErrBadRequest)
		return
	}

	id, err := service.DeviceService.Register(c, device)
	c.response(map[string]interface{}{"device_id": id}, err)
}
