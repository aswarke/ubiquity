package block_device_utils

import (
	"github.com/IBM/ubiquity/utils"
	"github.com/IBM/ubiquity/logutil"
)

type Protocol int

const (
	SCSI Protocol = iota
	ISCSI
)

//go:generate counterfeiter -o ../fakes/fake_block_device_utils.go . BlockDeviceUtils
type BlockDeviceUtils interface {
	Rescan(protocol Protocol) error
	ReloadMultipath() error
	Discover(volumeWwn string) (string, error)
	Cleanup(mpath string) error
	CheckFs(mpath string) (bool, error)
	MakeFs(mpath string, fsType string) error
	MountFs(mpath string, mpoint string) error
	UmountFs(mpoint string) error
}

type impBlockDeviceUtils struct {
	logger logutil.Logger
	exec   utils.Executor
}

func NewBlockDeviceUtils() BlockDeviceUtils {
	blockDeviceUtils := impBlockDeviceUtils{exec: utils.NewExecutor()}
	blockDeviceUtils.logger = logutil.GetLogger()
	return &blockDeviceUtils
}

func NewBlockDeviceUtilsWithExecutor(executor utils.Executor) BlockDeviceUtils {
	blockDeviceUtils := impBlockDeviceUtils{exec: executor}
	blockDeviceUtils.logger = logutil.GetLogger()
	return &blockDeviceUtils
}