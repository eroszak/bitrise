package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/bitrise/bitrise"
	"github.com/codegangsta/cli"
)

func normalize(c *cli.Context) {
	// Input validation
	bitriseConfigPath, err := GetBitriseConfigFilePath(c)
	if err != nil {
		log.Fatalf("Failed to get bitrise config path, err: %s", err)
	}
	if bitriseConfigPath == "" {
		log.Fatal("No bitrise config path defined!")
	}

	// Config validation
	bitriseConfig, warnings, err := CreateBitriseConfigFromCLIParams(c)
	for _, warning := range warnings {
		log.Warnf("warning: %s", warning)
	}
	if err != nil {
		log.Fatalf("Failed to create bitrise config, err: %s", err)
	}

	// Normalize
	if err := bitrise.RemoveConfigRedundantFieldsAndFillStepOutputs(&bitriseConfig); err != nil {
		log.Fatal("Failed to remove redundant fields:", err)
	}
	if err := bitrise.SaveConfigToFile(bitriseConfigPath, bitriseConfig); err != nil {
		log.Fatal("Failed to save config to file:", err)
	}

	log.Info("Redundant fields removed")
}
