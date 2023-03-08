package d2g

import (
	"github.com/taskcluster/d2g/genericworker"
)

func (din *DockerImageName) PrepareCommands() [][]string {
	return [][]string{}
}

func (din *DockerImageName) FileMounts() ([]genericworker.FileMount, error) {
	return []genericworker.FileMount{}, nil
}

func (din *DockerImageName) String() (string, error) {
	return string(*din), nil
}
