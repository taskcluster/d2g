package d2g

import (
	"encoding/json"
	"fmt"

	"github.com/taskcluster/d2g/genericworker"
)

func (dia *DockerImageArtifact) PrepareCommands() []string {
	return []string{
		"IMAGE_NAME=$(podman load -i image.tar | sed -n 's/.*: //p')",
	}
}

func (dia *DockerImageArtifact) FileMounts() ([]genericworker.FileMount, error) {
	artifactContent := genericworker.ArtifactContent{
		Artifact: dia.Path,
		Sha256:   "", // We could add this as an optional property to docker worker schema
		TaskID:   dia.TaskID,
	}
	raw, err := json.MarshalIndent(&artifactContent, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("cannot marshal artifact content %#v into json: %w", artifactContent, err)
	}
	return []genericworker.FileMount{
		{
			Content: json.RawMessage(raw),
			File:    "image.tar",
		},
	}, nil
}

func (dia *DockerImageArtifact) String() (string, error) {
	return `"${IMAGE_NAME}"`, nil
}
