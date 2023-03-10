// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package dockerworker

import (
	"encoding/json"
	tcclient "github.com/taskcluster/taskcluster/v48/clients/client-go"
)

type (
	Artifact struct {
		Expires tcclient.Time `json:"expires,omitempty"`

		Path string `json:"path"`

		// Possible values:
		//   * "file"
		//   * "directory"
		Type string `json:"type"`
	}

	// Set of capabilities that must be enabled or made available to the task container Example: ```{ "capabilities": { "privileged": true }```
	Capabilities struct {

		// Allows devices from the host system to be attached to a task container similar to using `--device` in docker.
		Devices Devices `json:"devices,omitempty"`

		// Allows a task to run in a privileged container, similar to running docker with `--privileged`.  This only works for worker-types configured to enable it.
		//
		// Default:    false
		Privileged bool `json:"privileged,omitempty"`
	}

	// Allows devices from the host system to be attached to a task container similar to using `--device` in docker.
	Devices struct {

		// Mount /dev/shm from the host in the container.
		HostSharedMemory bool `json:"hostSharedMemory,omitempty"`

		// Mount /dev/kvm from the host in the container.
		Kvm bool `json:"kvm,omitempty"`

		// Audio loopback device created using snd-aloop
		LoopbackAudio bool `json:"loopbackAudio,omitempty"`

		// Video loopback device created using v4l2loopback.
		LoopbackVideo bool `json:"loopbackVideo,omitempty"`
	}

	// Image to use for the task.  Images can be specified as an image tag as used by a docker registry, or as an object declaring type and name/namespace
	DockerImageArtifact struct {
		Path string `json:"path"`

		TaskID string `json:"taskId"`

		// Possible values:
		//   * "task-image"
		Type string `json:"type"`
	}

	// Image to use for the task.  Images can be specified as an image tag as used by a docker registry, or as an object declaring type and name/namespace
	DockerImageName string

	// `.payload` field of the queue.
	DockerWorkerPayload struct {

		// Artifact upload map example: ```{"public/build.tar.gz": {"path": "/home/worker/build.tar.gz", "expires": "2016-05-28T16:12:56.693817Z", "type": "file"}}```
		Artifacts map[string]Artifact `json:"artifacts,omitempty"`

		// Caches are mounted within the docker container at the mount point specified. Example: ```{ "CACHE NAME": "/mount/path/in/container" }```
		//
		// Map entries:
		Cache map[string]string `json:"cache,omitempty"`

		// Set of capabilities that must be enabled or made available to the task container Example: ```{ "capabilities": { "privileged": true }```
		Capabilities Capabilities `json:"capabilities,omitempty"`

		// Example: `['/bin/bash', '-c', 'ls']`.
		//
		// Default:    []
		//
		// Array items:
		Command []string `json:"command,omitempty"`

		// Example: ```
		// {
		//   "PATH": '/borked/path'
		//   "ENV_NAME": "VALUE"
		// }
		// ```
		//
		// Map entries:
		Env map[string]string `json:"env,omitempty"`

		// Used to enable additional functionality.
		Features FeatureFlags `json:"features,omitempty"`

		// Image to use for the task.  Images can be specified as an image tag as used by a docker registry, or as an object declaring type and name/namespace
		//
		// One of:
		//   * DockerImageName
		//   * NamedDockerImage
		//   * IndexedDockerImage
		//   * DockerImageArtifact
		Image json.RawMessage `json:"image"`

		// Specifies a custom location for the livelog artifact
		Log string `json:"log,omitempty"`

		// Maximum time the task container can run in seconds.
		//
		// Mininum:    1
		// Maximum:    86400
		MaxRunTime int64 `json:"maxRunTime"`

		// By default docker-worker will fail a task with a non-zero exit status without retrying.  This payload property allows a task owner to define certain exit statuses that will be marked as a retriable exception.
		OnExitStatus ExitStatusHandling `json:"onExitStatus,omitempty"`

		// Maintained for backward compatibility, but no longer used
		SupersederURL string `json:"supersederUrl,omitempty"`
	}

	// By default docker-worker will fail a task with a non-zero exit status without retrying.  This payload property allows a task owner to define certain exit statuses that will be marked as a retriable exception.
	ExitStatusHandling struct {

		// If the task exists with a purge caches exit status, all caches associated with the task will be purged.
		//
		// Array items:
		PurgeCaches []int64 `json:"purgeCaches,omitempty"`

		// If the task exists with a retriable exit status, the task will be marked as an exception and a new run created.
		//
		// Array items:
		Retry []int64 `json:"retry,omitempty"`
	}

	// Used to enable additional functionality.
	FeatureFlags struct {

		// This allows you to use the Linux ptrace functionality inside the container; it is otherwise disallowed by Docker's security policy.
		AllowPtrace bool `json:"allowPtrace,omitempty"`

		Artifacts bool `json:"artifacts,omitempty"`

		// Useful if live logging is not interesting but the overalllog is later on
		BulkLog bool `json:"bulkLog,omitempty"`

		// Artifacts named chain-of-trust.json and chain-of-trust.json.sig should be generated which will include information for downstream tasks to build a level of trust for the artifacts produced by the task and the environment it ran in.
		ChainOfTrust bool `json:"chainOfTrust,omitempty"`

		// Runs docker-in-docker and binds `/var/run/docker.sock` into the container. Doesn't allow privileged mode, capabilities or host volume mounts.
		Dind bool `json:"dind,omitempty"`

		// Uploads docker images as artifacts
		DockerSave bool `json:"dockerSave,omitempty"`

		// This allows you to interactively run commands inside the container and attaches you to the stdin/stdout/stderr over a websocket. Can be used for SSH-like access to docker containers.
		Interactive bool `json:"interactive,omitempty"`

		// Logs are stored on the worker during the duration of tasks and available via http chunked streaming then uploaded to s3
		LocalLiveLog bool `json:"localLiveLog,omitempty"`

		// The auth proxy allows making requests to taskcluster/queue and taskcluster/scheduler directly from your task with the same scopes as set in the task. This can be used to make api calls via the [client](https://github.com/taskcluster/taskcluster-client) CURL, etc... Without embedding credentials in the task.
		TaskclusterProxy bool `json:"taskclusterProxy,omitempty"`
	}

	// Image to use for the task.  Images can be specified as an image tag as used by a docker registry, or as an object declaring type and name/namespace
	IndexedDockerImage struct {
		Namespace string `json:"namespace"`

		Path string `json:"path"`

		// Possible values:
		//   * "indexed-image"
		Type string `json:"type"`
	}

	// Image to use for the task.  Images can be specified as an image tag as used by a docker registry, or as an object declaring type and name/namespace
	NamedDockerImage struct {
		Name string `json:"name"`

		// Possible values:
		//   * "docker-image"
		Type string `json:"type"`
	}
)
