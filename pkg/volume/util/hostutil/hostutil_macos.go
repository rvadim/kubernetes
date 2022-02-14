//go:build darwin
// +build darwin

package hostutil

import (
	"os"

	"k8s.io/mount-utils"
)

// HostUtil is an HostUtils implementation that allows compilation on
// unsupported platforms
type HostUtil struct{}

// NewHostUtil returns a struct that implements the HostUtils interface on
// unsupported platforms
func NewHostUtil() *HostUtil {
	return &HostUtil{}
}

// DeviceOpened always returns an error on unsupported platforms
func (hu *HostUtil) DeviceOpened(pathname string) (bool, error) {
	return false, nil
}

// PathIsDevice always returns an error on unsupported platforms
func (hu *HostUtil) PathIsDevice(pathname string) (bool, error) {
	return false, nil
}

// GetDeviceNameFromMount always returns an error on unsupported platforms
func (hu *HostUtil) GetDeviceNameFromMount(mounter mount.Interface, mountPath, pluginMountDir string) (string, error) {
	return getDeviceNameFromMount(mounter, mountPath, pluginMountDir)
}

// MakeRShared always returns an error on unsupported platforms
func (hu *HostUtil) MakeRShared(path string) error {
	return nil
}

// GetFileType always returns an error on unsupported platforms
func (hu *HostUtil) GetFileType(pathname string) (FileType, error) {
	return FileType("fake"), nil
}

// MakeFile always returns an error on unsupported platforms
func (hu *HostUtil) MakeFile(pathname string) error {
	return nil
}

// MakeDir always returns an error on unsupported platforms
func (hu *HostUtil) MakeDir(pathname string) error {
	return nil
}

// PathExists always returns an error on unsupported platforms
func (hu *HostUtil) PathExists(pathname string) (bool, error) {
	return false, nil
}

// EvalHostSymlinks always returns an error on unsupported platforms
func (hu *HostUtil) EvalHostSymlinks(pathname string) (string, error) {
	return "", nil
}

// GetOwner always returns an error on unsupported platforms
func (hu *HostUtil) GetOwner(pathname string) (int64, int64, error) {
	return -1, -1, nil
}

// GetSELinuxSupport always returns false on MacOS
func (hu *HostUtil) GetSELinuxSupport(pathname string) (bool, error) {
	return false, nil
}

//GetMode always returns an error on unsupported platforms
func (hu *HostUtil) GetMode(pathname string) (os.FileMode, error) {
	return 0, nil
}

func getDeviceNameFromMount(mounter mount.Interface, mountPath, pluginMountDir string) (string, error) {
	return "", nil
}
