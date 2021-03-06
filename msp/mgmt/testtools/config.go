/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package msptesttools

import (
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/msp/mgmt"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
)

func getConfigPath(dir string) (string, error) {
	// Try to read the dir
	if _, err := os.Stat(dir); err != nil {
		cfg := os.Getenv("PEER_CFG_PATH")
		if cfg != "" {
			dir = filepath.Join(cfg, dir)
		} else {
			dir = filepath.Join(os.Getenv("GOPATH"), "/src/github.com/hyperledger/fabric/msp/sampleconfig/")
		}
		if _, err := os.Stat(dir); err != nil {
			return "", err
		}
	}
	return dir, nil
}

// LoadTestMSPSetup sets up the local MSP
// and a chain MSP for the default chain
func LoadMSPSetupForTesting(dir string) error {
	var err error
	if dir, err = getConfigPath(dir); err != nil {
		return err
	}
	conf, err := msp.GetLocalMspConfig(dir, "DEFAULT")
	if err != nil {
		return err
	}

	err = mgmt.GetLocalMSP().Setup(conf)
	if err != nil {
		return err
	}

	fakeConfig := []*mspprotos.MSPConfig{conf}

	err = mgmt.GetManagerForChain(util.GetTestChainID()).Setup(fakeConfig)
	if err != nil {
		return err
	}

	return nil
}
