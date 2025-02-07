// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package linux

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	"github.com/cilium/cilium/pkg/datapath/linux/probes"
	"github.com/cilium/cilium/pkg/option"
)

// CheckRequirements checks that minimum kernel requirements are met for
// configuring the BPF datapath. If not, fatally exits.
func CheckRequirements() {
	_, err := netlink.RuleList(netlink.FAMILY_V4)
	if errors.Is(err, unix.EAFNOSUPPORT) {
		log.WithError(err).Error("Policy routing:NOT OK. " +
			"Please enable kernel configuration item CONFIG_IP_MULTIPLE_TABLES")
	}

	if option.Config.EnableIPv6 {
		if _, err := os.Stat("/proc/net/if_inet6"); os.IsNotExist(err) {
			log.Fatalf("kernel: ipv6 is enabled in agent but ipv6 is either disabled or not compiled in the kernel")
		}
	}

	// bpftool checks
	if !option.Config.DryMode {
		probeManager := probes.NewProbeManager()

		// VTEP integration feature requires kernel 1m large instruction support
		if option.Config.EnableVTEP {
			if probes.HaveLargeInstructionLimit() != nil {
				log.Fatalf("VXLAN Tunnel Endpoint (VTEP) Integration: requires support for large programs (Linux 5.2.0 or newer)")
			}
		}
		if err := probeManager.SystemConfigProbes(); err != nil {
			errMsg := "BPF system config check: NOT OK."
			// TODO(vincentmli): revisit log when GH#14314 has been resolved
			// Warn missing required kernel config option
			log.WithError(err).Warn(errMsg)
		}
		if err := probes.CreateHeaderFiles(filepath.Join(option.Config.BpfDir, "include/bpf"), probes.ExecuteHeaderProbes()); err != nil {
			log.WithError(err).Fatal("failed to create header files with feature macros")
		}
	}
}
