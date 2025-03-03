// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecompute

import (
	"context"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestStepTeardownInstance_impl(t *testing.T) {
	var _ multistep.Step = new(StepTeardownInstance)
}

func TestStepTeardownInstance(t *testing.T) {
	state := testState(t)
	step := new(StepTeardownInstance)
	defer step.Cleanup(state)

	config := state.Get("config").(*Config)
	driver := state.Get("driver").(*DriverMock)

	// run the step
	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	if driver.DeleteInstanceName != config.InstanceName {
		t.Fatal("should've deleted instance")
	}
	if driver.DeleteInstanceZone != config.Zone {
		t.Fatalf("bad zone: %#v", driver.DeleteInstanceZone)
	}

	// cleanup
	step.Cleanup(state)

	if driver.DeleteDiskName != config.InstanceName {
		t.Fatal("should've deleted disk")
	}
	if driver.DeleteDiskZone != config.Zone {
		t.Fatalf("bad zone: %#v", driver.DeleteDiskZone)
	}
}
