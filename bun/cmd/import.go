package cmd

import (
	_ "github.com/adyatlov/bun/checks/dcosversion"
	_ "github.com/adyatlov/bun/checks/health"
	_ "github.com/adyatlov/bun/checks/marathon/deployments"
	_ "github.com/adyatlov/bun/checks/mesos/actormailboxes"
	_ "github.com/adyatlov/bun/checks/nodecount"
	_ "github.com/adyatlov/bun/checks/networking"
)
