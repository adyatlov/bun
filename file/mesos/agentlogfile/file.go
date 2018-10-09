package agentlogfile

import "github.com/adyatlov/bun"

func init() {
	f := bun.FileType{
		Name:        "mesos-agent-log",
		Description: "Mesos agent jounrald log",
		ContentType: bun.CTJournal,
		Paths: []string{
			"dcos-mesos-slave.service",
			"dcos-mesos-slave-public.service",
		},
		DirTypes: []bun.DirType{bun.DTAgent, bun.DTPublicAgent},
	}
	bun.RegisterFileType(f)
}

// MsgFailedToUnmouint message appears in the Mesos agent logs when agent cannot
// unmount local persisten colume.
const MsgFailedToUnmouint = "Failed to remove rootfs mount point"
