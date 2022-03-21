package model

import (
	"etrisfpocdatamodel"
	"io"
)

type DBHandler interface {
	GetDevices() ([]*etrisfpocdatamodel.Device, int, error)
	AddDevice(d *etrisfpocdatamodel.Device) error
	QueryDevice(dname string) (*etrisfpocdatamodel.Device, error)
	DeleteDevice(device *etrisfpocdatamodel.Device) error
	IsExistDevice(dname string) bool
	AddEdge(r io.Reader) (*etrisfpocdatamodel.Edge, error)
	GetEdges() ([]*etrisfpocdatamodel.Edge, error)
	IsExistEdge(cid string) bool
	GetServices() ([]*etrisfpocdatamodel.Service, error)
	AddService(name string) error
	UpdateService(name, addr string) (*etrisfpocdatamodel.Service, error)
	GetAddr(sid string) (string, error)
	GetSID(name string) (string, error)
	IsExistService(name string) bool
}
