package voldriver

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/rata"
)

const (
	ActivateRoute     = "activate"
	CreateRoute       = "create"
	GetRoute          = "get"
	ListRoute         = "list"
	MountRoute        = "mount"
	PathRoute         = "path"
	RemoveRoute       = "remove"
	UnmountRoute      = "unmount"
	CapabilitiesRoute = "capabilities"
)

var Routes = rata.Routes{
	{Path: "/Plugin.Activate", Method: "POST", Name: ActivateRoute},
	{Path: "/VolumeDriver.Create", Method: "POST", Name: CreateRoute},
	{Path: "/VolumeDriver.Get", Method: "POST", Name: GetRoute},
	{Path: "/VolumeDriver.List", Method: "POST", Name: ListRoute},
	{Path: "/VolumeDriver.Mount", Method: "POST", Name: MountRoute},
	{Path: "/VolumeDriver.Path", Method: "POST", Name: PathRoute},
	{Path: "/VolumeDriver.Remove", Method: "POST", Name: RemoveRoute},
	{Path: "/VolumeDriver.Unmount", Method: "POST", Name: UnmountRoute},
	{Path: "/VolumeDriver.Capabilities", Method: "POST", Name: CapabilitiesRoute},
}

//go:generate counterfeiter -o voldriverfakes/fake_env.go . Env
type Env interface {
	Logger() lager.Logger
	Context() context.Context
}

//go:generate counterfeiter -o voldriverfakes/fake_matchable_driver_client.go . MatchableDriver
type MatchableDriver interface {
	Matches(lager.Logger, string, *TLSConfig) bool

	Driver
}

//go:generate counterfeiter -o voldriverfakes/fake_plugin_client.go . Plugin
type Plugin interface {
	// Eventually this method will have List, Mount, Unmount and Matches methods
	// allowing LocalClient and Purger to interact with Plugin without having
	// to know if they are a docker volume driver or a CSI plugin.
	//
	// However, in order to do a step-wise refactor we are initially introducing
	// the interface with a method that allows LocalClient and Purger to get at the
	// underlying Voldriver
	GetVoldriver() Driver

	Matches(lager.Logger, string, *TLSConfig) bool
}

type driverWrapper struct {
	Driver Driver
}

func (dw *driverWrapper) GetVoldriver() Driver {
	return dw.Driver
}

func (dw *driverWrapper) Matches(logger lager.Logger, address string, tlsConfig *TLSConfig) bool {
	logger = logger.Session("matches")
	logger.Info("start")
	defer logger.Info("end")

	matchableDriver, ok := dw.Driver.(MatchableDriver)
	logger.Info("is-matchable", lager.Data{"matchable": ok})
	if ok {
		return matchableDriver.Matches(logger, address, tlsConfig)
	}
	return false
}

func NewVoldriverPlugin(driver Driver) Plugin {
	return &driverWrapper{
		Driver: driver,
	}
}

//go:generate counterfeiter -o voldriverfakes/fake_driver_client.go . Driver
type Driver interface {
	Activate(env Env) ActivateResponse
	Get(env Env, getRequest GetRequest) GetResponse
	List(env Env) ListResponse
	Mount(env Env, mountRequest MountRequest) MountResponse
	Path(env Env, pathRequest PathRequest) PathResponse
	Unmount(env Env, unmountRequest UnmountRequest) ErrorResponse
	Capabilities(env Env) CapabilitiesResponse

	Provisioner
}

//go:generate counterfeiter -o voldriverfakes/fake_provisioner.go . Provisioner
type Provisioner interface {
	Create(env Env, createRequest CreateRequest) ErrorResponse
	Remove(env Env, removeRequest RemoveRequest) ErrorResponse
}

type ActivateResponse struct {
	Err        string
	Implements []string
}

type CreateRequest struct {
	Name string
	Opts map[string]interface{}
}

type MountRequest struct {
	Name string
}

type MountResponse struct {
	Err        string
	Mountpoint string
}

type ListResponse struct {
	Volumes []VolumeInfo
	Err     string
}

type PathRequest struct {
	Name string
}

type PathResponse struct {
	Err        string
	Mountpoint string
}

type UnmountRequest struct {
	Name string
}

type RemoveRequest struct {
	Name string
}

type ErrorResponse struct {
	Err string
}

type GetRequest struct {
	Name string
}

type GetResponse struct {
	Volume VolumeInfo
	Err    string
}

type CapabilitiesResponse struct {
	Capabilities CapabilityInfo
}

type VolumeInfo struct {
	Name       string
	Mountpoint string
	MountCount int
}

type CapabilityInfo struct {
	Scope string
}

type Error struct {
	Description string `json:"description"`
}

func (e Error) Error() string {
	return e.Description
}

type TLSConfig struct {
	InsecureSkipVerify bool   `json:"InsecureSkipVerify"`
	CAFile             string `json:"CAFile"`
	CertFile           string `json:"CertFile"`
	KeyFile            string `json:"KeyFile"`
}

type DriverSpec struct {
	Name      string     `json:"Name"`
	Address   string     `json:"Addr"`
	TLSConfig *TLSConfig `json:"TLSConfig"`
}
