package aws

// Worker configures AWS-specific worker node settings.
type Worker struct {
	InstanceType string `json:"instance_type"`
}

// DefaultWorker provides default WorkerAWSSettings.
func DefaultWorker() Worker {
	return Worker{
		InstanceType: "",
	}
}
