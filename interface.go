package iotmakerDockerInterface

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.docker"
)

type DockerSystem interface {
	ImageGarbageCollector() (err error)
	ContainerStatisticsOneShotByName(
		name string,
	) (
		err error,
		ret types.Stats,
	)
	ContainerStatisticsOneShot(
		id string,
	) (
		err error,
		statsRet types.Stats,
	)
	ContainerListQuiet() (
		err error,
		list []types.Container,
	)
	ContainerListAll() (
		err error,
		list []types.Container,
	)
	ContainerInspectByName(
		name string,
	) (
		err error,
		inspect types.ContainerJSON,
	)
	ContainerInspect(
		id string,
	) (
		err error,
		inspect types.ContainerJSON,
	)
	NetworkCreate(
		name string,
		drive iotmakerDocker.NetworkDrive,
		scope,
		subnet,
		gateway string,
	) (err error, id string, networkGenerator *iotmakerDocker.NextNetworkAutoConfiguration)
	NetworkRemove(id string) error
	NetworkVerifyName(name string) (error, bool)
	ImagePull(name string, channel *chan iotmakerDocker.ContainerPullStatusSendToChannel) (err error, imageId string, imageName string)
	ImageWaitPull(name string) error
	ImageList() (
		err error,
		list []types.ImageSummary,
	)
	ImageFindIdByName(
		name string,
	) (
		err error,
		ID string,
	)
	ImageVerifyVolume(id, path string) (error, bool)
	ImageListExposedVolumes(
		id string,
	) (
		err error,
		list []string,
	)
	ImageListExposedVolumesByName(name string) (error, []string)
	ImageListExposedPorts(
		id string,
	) (
		err error,
		portList []nat.Port,
	)
	ImageListExposedPortsByName(
		name string,
	) (
		err error,
		portList []nat.Port,
	)
	ContainerStart(
		id string,
	) (
		err error,
	)
	ContainerCreateAndStart(
		imageName,
		containerName string,
		restart iotmakerDocker.RestartPolicy,
		portExposedList nat.PortMap,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
	) (
		err error,
		containerID string,
	)
	ContainerCreate(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		portExposedList nat.PortMap,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
	) (
		err error,
		containerID string,
	)
	ContainerCreateAndChangeExposedPort(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
		currentPort,
		changeToPort []nat.Port,
	) (
		err error,
		containerID string,
	)
	FileMakeAbsolutePath(
		filePath string,
	) (
		err error,
		fileAbsolutePath string,
	)
	ImageMountNatPortListChangeExposed(imageId string, currentPortList, changeToPortList []nat.Port) (error, nat.PortMap)
	ImageMountNatPortList(imageId string) (error, nat.PortMap)
	ClientCreate() (
		err error,
	)
	ContextCreate()
	Init() error
	ImageBuildFromFolder(
		folderPath string,
		tags []string,
		channel *chan iotmakerDocker.ContainerPullStatusSendToChannel,
	) (
		err error,
	)
	ContainerFindIdByName(
		name string,
	) (
		err error,
		ID string,
	)
	ContainerNetworkInspect(
		id string,
	) (
		err error,
		netDataList iotmakerDocker.ContainerNetworkDataList,
	)
	ContainerStopAndRemove(
		id string,
		removeVolumes,
		removeLinks,
		force bool,
	) (
		err error,
	)
	ImageRemove(id string) error
	ImageBuildFromRemoteServer(
		server,
		imageName string,
		tags []string,
		channel *chan iotmakerDocker.ContainerPullStatusSendToChannel,
	) (
		err error,
	)
	ImageRemoveByName(name string) error
	ContainerWaitStatusNotRunning(
		id string,
	) (
		err error,
	)
	ContainerRemove(
		id string,
		removeVolumes,
		removeLinks,
		force bool,
	) (
		err error,
	)
	VolumeRemove(name string) (err error)
	VolumeFindByName(name string) (err error, volume types.Volume)
	VolumeList() (err error, volList []types.Volume)
	NetworkList() (err error, netList []types.NetworkResource)
	NetworkFindIdByName(name string) (err error, id string)
	NetworkRemoveByName(name string) error
	ContainerStop(
		id string,
	) (
		err error,
	)
	ContainerInspectJSonByName(
		name string,
	) (
		err error,
		inspect []byte,
	)
	ContainerInspectJSon(
		id string,
	) (
		err error,
		inspect []byte,
	)
	ContainerCreateWithoutExposePortsAndStart(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
	) (
		err error,
		containerID string,
	)
	ContainerCreateExposePortsAutomaticallyAndStart(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
	) (
		err error,
		containerID string,
	)
	ContainerCreateWithoutExposePorts(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
	) (
		err error,
		containerID string,
	)
	ContainerCreateChangeExposedPortAndStart(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
		currentPort,
		changeToPort []nat.Port,
	) (
		err error,
		containerID string,
	)
	NetworkConnect(networkID, containerID string, config *network.EndpointSettings) (err error)
	ContainerCreateAndExposePortsAutomatically(
		imageName,
		containerName string,
		restartPolicy iotmakerDocker.RestartPolicy,
		mountVolumes []mount.Mount,
		containerNetwork *network.NetworkingConfig,
	) (
		err error,
		containerID string,
	)
	NetworkFindNetworkTypeBridgePublic() (err error, inspect types.NetworkResource)
	NetworkInspect(
		id string,
	) (err error, inspect types.NetworkResource)
	ContainerFindIdByNameContains(
		containsName string,
	) (
		err error,
		ID string,
	)
	ImageMountNatPortListChangeExposedWithIpAddress(imageId, ipAddress string, currentPortList, changeToPortList []nat.Port) (error, nat.PortMap)
	AdjustImageName(
		imageName string,
	) string
}
