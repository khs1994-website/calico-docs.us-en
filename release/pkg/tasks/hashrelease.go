// Copyright (c) 2024 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tasks

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/projectcalico/calico/release/internal/hashreleaseserver"
	"github.com/projectcalico/calico/release/internal/imagescanner"
	"github.com/projectcalico/calico/release/internal/pinnedversion"
	"github.com/projectcalico/calico/release/internal/slack"
	"github.com/projectcalico/calico/release/internal/utils"
)

// HashreleasePublished checks if the hashrelease has already been published.
func HashreleasePublished(cfg *hashreleaseserver.Config, hash string, ci bool) (bool, error) {
	if !cfg.Valid() {
		// Check if we're running in CI - if so, we should fail if this configuration is missing.
		// Otherwise, we should just log and continue.
		if ci {
			return false, fmt.Errorf("missing hashrelease server configuration")
		}
		logrus.Warn("Missing hashrelease server configuration, skipping remote hashrelease check")
		return false, nil
	}

	return hashreleaseserver.HasHashrelease(hash, cfg)
}

// HashreleaseSlackMessage sends a slack message to notify that a hashrelease has been published.
func HashreleaseSlackMessage(slackCfg *slack.Config, hashrel *hashreleaseserver.Hashrelease, imageScanResultsExists bool, ciURL, tmpDir string) error {
	var scanResultURL string
	if imageScanResultsExists {
		scanResultURL = imagescanner.RetrieveResultURL(tmpDir)
		if scanResultURL == "" {
			logrus.Warn("No image scan result URL found")
		}
	}
	slackMsg := slack.Message{
		Config: *slackCfg,
		Data: slack.MessageData{
			ReleaseName:        hashrel.Name,
			Product:            utils.DisplayProductName(),
			Stream:             hashrel.Stream,
			Version:            hashrel.ProductVersion,
			OperatorVersion:    hashrel.OperatorVersion,
			DocsURL:            hashrel.URL(),
			CIURL:              ciURL,
			ImageScanResultURL: scanResultURL,
		},
	}
	if err := slackMsg.SendSuccess(logrus.IsLevelEnabled(logrus.DebugLevel)); err != nil {
		logrus.WithError(err).Error("Failed to send slack message")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"name": hashrel.Name,
		"URL":  hashrel.URL(),
	}).Info("Sent hashrelease publish notification to slack")
	return nil
}

// ReformatHashrelease modifies the generated release output to match
// the "legacy" format our CI tooling expects. This should be temporary until
// we can update the tooling to expect the new format.
// Specifically, we need to do two things:
// - Copy the windows zip file to files/windows/calico-windows-<ver>.zip
// - Copy tigera-operator-<ver>.tgz to tigera-operator.tgz
// - Copy ocp.tgz to manifests/ocp.tgz
func ReformatHashrelease(hashreleaseOutputDir, tmpDir string) error {
	logrus.Info("Modifying hashrelease output to match legacy format")
	pinned, err := pinnedversion.RetrievePinnedVersion(tmpDir)
	if err != nil {
		return err
	}
	ver := pinned.Components[utils.Calico].Version

	// Copy the windows zip file to files/windows/calico-windows-<ver>.zip
	if err := os.MkdirAll(filepath.Join(hashreleaseOutputDir, "files", "windows"), 0o755); err != nil {
		return err
	}
	windowsZip := filepath.Join(hashreleaseOutputDir, fmt.Sprintf("calico-windows-%s.zip", ver))
	windowsZipDst := filepath.Join(hashreleaseOutputDir, "files", "windows", fmt.Sprintf("calico-windows-%s.zip", ver))
	if err := utils.CopyFile(windowsZip, windowsZipDst); err != nil {
		return err
	}

	// Copy the ocp.tgz to manifests/ocp.tgz
	ocpTarball := filepath.Join(hashreleaseOutputDir, "ocp.tgz")
	ocpTarballDst := filepath.Join(hashreleaseOutputDir, "manifests", "ocp.tgz")
	if err := utils.CopyFile(ocpTarball, ocpTarballDst); err != nil {
		return err
	}

	// Copy the operator tarball to tigera-operator.tgz
	helmChartVersion := ver
	operatorTarball := filepath.Join(hashreleaseOutputDir, fmt.Sprintf("tigera-operator-%s.tgz", helmChartVersion))
	operatorTarballDst := filepath.Join(hashreleaseOutputDir, "tigera-operator.tgz")
	if err := utils.CopyFile(operatorTarball, operatorTarballDst); err != nil {
		return err
	}
	return nil
}
