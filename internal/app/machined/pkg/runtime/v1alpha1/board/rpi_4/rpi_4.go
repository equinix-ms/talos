// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package rpi4

import (
	"io/ioutil"

	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
	"github.com/talos-systems/talos/pkg/copy"
	"github.com/talos-systems/talos/pkg/machinery/constants"
)

var configTxt = []byte(`arm_64bit=1
enable_uart=1
kernel=u-boot.bin
`)

// RPi4 represents the Raspberry Pi 4 Model B.
//
// Reference: https://www.raspberrypi.org/products/raspberry-pi-4-model-b/
type RPi4 struct{}

// Name implements the runtime.Board.
func (l *RPi4) Name() string {
	return constants.BoardRPi4
}

// Install implements the runtime.Board.
func (l *RPi4) Install(disk string) (err error) {
	err = copy.Dir("/usr/install/raspberrypi-firmware/boot", "/boot/EFI")
	if err != nil {
		return err
	}

	err = copy.File("/usr/install/u-boot/rpi_4/u-boot.bin", "/boot/EFI/u-boot.bin")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("/boot/EFI/config.txt", configTxt, 0o600)
}

// PartitionOptions implements the runtime.Board.
func (l *RPi4) PartitionOptions() *runtime.PartitionOptions {
	return nil
}