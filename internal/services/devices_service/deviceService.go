package devices_service

import (
	"applicant-backend/internal/entities"
	"errors"
)

type DeviceService struct {
}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (d *DeviceService) GetDevices() []entities.Device {
	return DUMMY_DEVICES
}

func (d *DeviceService) GetDeviceById(id string) (*entities.Device, error) {
	devices := d.GetDevices()
	for _, device := range devices {
		if device.Id == id {
			return &device, nil
		}
	}

	return nil, errors.New("Could not find Device with ID: " + id)
}
