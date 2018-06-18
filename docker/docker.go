package docker

import (
	"log"
	"os/exec"
	"strings"

	"github.com/EverLoSa/dbaas/model"
)

// const for docker commands
const (
	PortCMD  = "docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}'"
	IPCMD    = "docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'"
	DockerRM = "docker rm -f"
)

// NewContainer creates a new docker container in the host
func NewContainer(db *model.DataBaseInfo, image string) error {

	// create container
	var err error
	db.ContainerInfo.Image = "mysql:latest"
	cmdStr := "docker run -e MYSQL_ROOT_PASSWORD=changeMe -p 3306 -d " + image + ":latest"

	// create container and get container id
	if db.ContainerInfo.ContainerID, err = executeDockerCommand(cmdStr); err != nil {
		log.Println("An error ocurred executing a command to create a container")
	}

	// get container port
	if db.ContainerInfo.Port, err = executeDockerCommand(strings.Join([]string{PortCMD, db.ContainerInfo.ContainerID}, " ")); err != nil {
		log.Println("An error ocurred getting container port")
	}

	return err

}

// DeleteContainer deletes a container from the host
func DeleteContainer(containerID string) error {
	var err error
	rmCMD := strings.Join([]string{DockerRM, containerID}, " ")
	if _, err := executeDockerCommand(rmCMD); err != nil {
		log.Println("An error ocurred deleting the container")
	}
	return err
}

func executeDockerCommand(cmd string) (output string, err error) {
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	return string(out[:]), err
}
