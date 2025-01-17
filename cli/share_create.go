package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/bitrise/tools"
	"github.com/codegangsta/cli"
)

func create(c *cli.Context) {
	// Input validation
	tag := c.String(TagKey)
	if tag == "" {
		log.Fatalln("No step tag specified")
	}

	gitURI := c.String(GitKey)
	if gitURI == "" {
		log.Fatalln("No step url specified")
	}

	stepID := c.String(StepIDKey)

	if err := tools.StepmanShareCreate(tag, gitURI, stepID); err != nil {
		log.Fatalf("Bitrise share create failed, err: %s", err)
	}
}
