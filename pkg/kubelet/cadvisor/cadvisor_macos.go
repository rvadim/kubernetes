//+build darwin

package cadvisor

import (
	"time"

	"github.com/google/cadvisor/events"
	cadvisorapi "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
)

type cadvisorMacOS struct {
}

var _ Interface = new(cadvisorMacOS)

// New creates a new cAdvisor Interface for unsupported systems.
func New(imageFsInfoProvider ImageFsInfoProvider, rootPath string, cgroupsRoots []string, usingLegacyStats bool) (Interface, error) {
	return &cadvisorMacOS{}, nil
}

// Start is a mock implementation of Interface.Start.
func (c *cadvisorMacOS) Start() error {
	return nil
}

// ContainerInfo is a mock implementation of Interface.ContainerInfo.
func (c *cadvisorMacOS) ContainerInfo(name string, req *cadvisorapi.ContainerInfoRequest) (*cadvisorapi.ContainerInfo, error) {
	return &cadvisorapi.ContainerInfo{}, nil
}

// ContainerInfoV2 is a mock implementation of Interface.ContainerInfoV2.
func (c *cadvisorMacOS) ContainerInfoV2(name string, options cadvisorapiv2.RequestOptions) (map[string]cadvisorapiv2.ContainerInfo, error) {
	return map[string]cadvisorapiv2.ContainerInfo{}, nil
}

// GetRequestedContainersInfo is a fake implementation if Interface.GetRequestedContainersInfo
func (c *cadvisorMacOS) GetRequestedContainersInfo(containerName string, options cadvisorapiv2.RequestOptions) (map[string]*cadvisorapi.ContainerInfo, error) {
	return map[string]*cadvisorapi.ContainerInfo{}, nil
}

// SubcontainerInfo is a mock implementation of Interface.SubcontainerInfo.
func (c *cadvisorMacOS) SubcontainerInfo(name string, req *cadvisorapi.ContainerInfoRequest) (map[string]*cadvisorapi.ContainerInfo, error) {
	return map[string]*cadvisorapi.ContainerInfo{}, nil
}

// DockerContainer is a mock implementation of Interface.DockerContainer.
func (c *cadvisorMacOS) DockerContainer(name string, req *cadvisorapi.ContainerInfoRequest) (cadvisorapi.ContainerInfo, error) {
	return cadvisorapi.ContainerInfo{}, nil
}

// MachineInfo is a mock implementation of Interface.MachineInfo.
func (c *cadvisorMacOS) MachineInfo() (*cadvisorapi.MachineInfo, error) {
	return &cadvisorapi.MachineInfo{
		Timestamp:        time.Now(),
		NumCores:         4,
		NumPhysicalCores: 1,
		NumSockets:       1,
		CpuFrequency:     1000,
		MemoryCapacity:   1024 * 1024 * 1024,
		MemoryByType:     map[string]*cadvisorapi.MemoryInfo{},
		NVMInfo:          cadvisorapi.NVMInfo{},
		HugePages:        []cadvisorapi.HugePagesInfo{},
		MachineID:        "my-mac",
		SystemUUID:       "my-uuid",
		BootID:           "my-boot-id",
		Filesystems:      []cadvisorapi.FsInfo{},
		DiskMap:          map[string]cadvisorapi.DiskInfo{},
		NetworkDevices:   []cadvisorapi.NetInfo{},
		Topology:         []cadvisorapi.Node{},
		CloudProvider:    "JetInfraPlatform",
		InstanceType:     "my-instance-type",
		InstanceID:       "my-instance-id",
	}, nil
}

// VersionInfo is a mock implementation of Interface.VersionInfo.
func (c *cadvisorMacOS) VersionInfo() (*cadvisorapi.VersionInfo, error) {
	return &cadvisorapi.VersionInfo{}, nil
}

// ImagesFsInfo is a mock implementation of Interface.ImagesFsInfo.
func (c *cadvisorMacOS) ImagesFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}

// RootFsInfo is a mock implementation of Interface.RootFsInfo.
func (c *cadvisorMacOS) RootFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}

// WatchEvents is a mock implementation of Interface.WatchEvents.
func (c *cadvisorMacOS) WatchEvents(request *events.Request) (*events.EventChannel, error) {
	return &events.EventChannel{}, nil
}

// GetDirFsInfo is a mock implementation of Interface.GetDirFsInfo.
func (c *cadvisorMacOS) GetDirFsInfo(path string) (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}
