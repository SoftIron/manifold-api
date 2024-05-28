package instance

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/softiron/manifold-api/internal/api"
)

// ParseTemplate return a structured Template based on the given map.
func ParseTemplate(m map[string]any) (*Template, error) { // nolint:gocognit
	var dst Template

	for key, value := range m {
		switch v := value.(type) {
		case []string:
			switch key {
			case "SECURITY_GROUP_RULE":
				dst.SecurityGroupRule = slices.Clone(v)
			case "VMGROUP":
				dst.InstanceGroup = slices.Clone(v)
			}
		case string:
			switch key {
			case "AUTOMATIC_DS_REQUIREMENTS":
				dst.AutomaticDSRequirements = v
			case "AUTOMATIC_NIC_REQUIREMENTS":
				dst.AutomaticNICRequirements = v
			case "AUTOMATIC_REQUIREMENTS":
				dst.AutomaticRequirements = v
			case "CPU":
				n, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return nil, fmt.Errorf("invalid CPU value %q: %w", v, err)
				}
				dst.CPU = float32(n)
			case "CPU_COST":
				dst.CPUCost = v
			case "CLONING_TEMPLATE_ID":
				dst.CloningTemplateID = v
			case "CREATED_BY":
				dst.CreatedBy = v
			case "DISK_COST":
				dst.DiskCost = v
			case "EMULATOR":
				dst.Emulator = v
			case "FEATURES":
				dst.Features = v
			case "HYPERV_OPTIONS":
				dst.HypervOptions = v
			case "IMPORTED":
				dst.Imported = v
			case "INPUT":
				dst.Input = v
			case "MEMORY":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid MEMORY value %q: %w", v, err)
				}
				dst.Memory = n
			case "MEMORY_COST":
				dst.MemoryCost = v
			case "MEMORY_MAX":
				dst.MemoryMax = v
			case "NIC_DEFAULT":
				dst.NICDefault = v
			case "NUMA_NODE":
				dst.NumaNode = v
			case "PCI":
				dst.PCI = v
			case "RAW":
				dst.Raw = v
			case "SPICE_OPTIONS":
				dst.SpiceOptions = v
			case "SUBMIT_ON_HOLD":
				dst.SubmitOnHold = v
			case "TEMPLATE_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid TEMPLATE_ID value %q: %w", v, err)
				}
				dst.TemplateID = n
			case "TM_MAD_SYSTEM":
				dst.TmMADSystem = v
			case "TOPOLOGY":
				dst.Topology = v
			case "VCPU":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid VCPU value %q: %w", v, err)
				}
				dst.VCPU = n
			case "VCPU_MAX":
				dst.VCPUMax = v
			case "VMID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid VMID value %q: %w", v, err)
				}
				dst.InstanceID = n
			case "VROUTER_ID":
				dst.VRouterID = v
			case "VROUTER_KEEPALIVED_ID":
				dst.VRouterKeepalivedID = v
			case "VROUTER_KEEPALIVED_PASSWORD":
				dst.VRouterKeepalivedPassword = v
			}
		case map[string]any:
			switch key {
			case "CONTEXT":
				c, err := newInstanceContext(v)
				if err != nil {
					return nil, err
				}
				dst.Context = *c
			case "CPU_MODEL":
				m := newInstanceCPUModel(v)
				dst.CPUModel = *m
			case "DISK":
				d, err := newInstanceDisk(v)
				if err != nil {
					return nil, err
				}
				dst.Disk = []Disk{*d}
			case "GRAPHICS":
				g, err := newInstanceGraphics(v)
				if err != nil {
					return nil, err
				}
				dst.Graphics = *g
			case "NIC":
				n, err := newInstanceNIC(v)
				if err != nil {
					return nil, err
				}
				dst.NIC = []NIC{*n}
			case "NIC_ALIAS":
				n := newInstanceNICAlias(v)
				dst.NICAlias = []NICAlias{*n}
			case "OS":
				o := newInstanceOS(v)
				dst.OS = *o
			case "SCHED_ACTION":
				s := newInstanceSchedAction(v)
				dst.SchedAction = []SchedAction{*s}
			case "SNAPSHOT":
				s, err := newInstanceSnapshot(v)
				if err != nil {
					return nil, err
				}
				dst.Snapshot = []Snapshot{*s}
			}
		case []map[string]any:
			switch key {
			case "DISK":
				for _, m := range v {
					d, err := newInstanceDisk(m)
					if err != nil {
						return nil, err
					}
					dst.Disk = append(dst.Disk, *d)
				}
			case "NIC":
				for _, m := range v {
					n, err := newInstanceNIC(m)
					if err != nil {
						return nil, err
					}
					dst.NIC = append(dst.NIC, *n)
				}
			case "NIC_ALIAS":
				for _, m := range v {
					n := newInstanceNICAlias(m)
					dst.NICAlias = append(dst.NICAlias, *n)
				}
			case "SCHED_ACTION":
				for _, m := range v {
					s := newInstanceSchedAction(m)
					dst.SchedAction = append(dst.SchedAction, *s)
				}
			case "SNAPSHOT":
				for _, m := range v {
					s, err := newInstanceSnapshot(m)
					if err != nil {
						return nil, err
					}
					dst.Snapshot = append(dst.Snapshot, *s)
				}
			}
		}
	}

	return &dst, nil
}

// ParseUserTemplate return a structured Template based on the given map.
//
//nolint:gocritic
func ParseUserTemplate(m map[string]any) (*UserTemplate, error) { // nolint:gocognit
	var dst UserTemplate

	for key, value := range m {
		switch v := value.(type) {
		case string:
			switch key {
			case "SCHED_REQUIREMENTS":
				dst.SchedRequirements = v
			case "SCHED_DS_REQUIREMENTS":
				dst.SchedDSRequirements = v
			case "LOGO":
				dst.Logo = v
			case "INFO":
				dst.Info = v
			case "ERROR":
				dst.Error = v
			case "SNAPSHOT_SCHEDULE":
				dst.SnapshotSchedule = v
			case "HYPERVISOR":
				dst.Hypervisor = v
			case "DESCRIPTION":
				dst.Description = v
			}
		}
	}

	return &dst, nil
}

func newInstanceContext(m map[string]any) (*Context, error) {
	var dst Context

	for key, value := range m {
		err := dst.Set(key, value)
		if err != nil {
			return nil, err
		}
	}

	return &dst, nil
}

// Set sets a key:value for the given context
// The key must match known fields and named with capital letters with underscores
// e.g. DISK_ID, SSH_PUBLIC_KEY, etc.
func (c *Context) Set(key string, value any) error {
	return setContextValue(c, key, value)
}

func setContextValue(dst *Context, key string, value any) error {
	if v, ok := value.(string); ok {
		switch key {
		case "DISK_ID":
			n, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("invalid DISK_ID value %q: %w", v, err)
			}
			dst.DiskID = n
		case "FIRMWARE":
			dst.Firmware = v
		case "GUESTOS":
			dst.GuestOS = v
		case "NETWORK":
			b, err := api.Str2Bool(v)
			if err != nil {
				return fmt.Errorf("invalid NETWORK value %q: %w", v, err)
			}
			dst.Network = b
		case "SSH_PUBLIC_KEY":
			dst.SSHPublicKey = v
		case "TARGET":
			dst.Target = v
		case "PROJECT_NAME":
			dst.ProjectName = v
		default:
			return fmt.Errorf("unknown key: %s", key)
		}
	}

	return nil
}

func newInstanceCPUModel(m map[string]any) *CPUModel {
	var dst CPUModel

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key { // nolint:gocritic
			case "MODEL":
				dst.Model = v
			}
		}
	}

	return &dst
}

func newInstanceGraphics(m map[string]any) (*Graphics, error) {
	var dst Graphics

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "KEYMAP":
				dst.Keymap = v
			case "LISTEN":
				dst.Listen = v
			case "PORT":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid PORT value %q: %w", v, err)
				}
				dst.Port = n
			case "TYPE":
				dst.Type = v
			case "PASSWD":
				dst.Password = v
			}
		}
	}

	return &dst, nil
}

func newInstanceOS(m map[string]any) *OS {
	var dst OS

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ARCH":
				dst.Arch = v
			case "BOOT":
				dst.Boot = v
			case "UUID":
				dst.UUID = v
			case "SD_DISK_BUS":
				dst.DiskBus = v
			}
		}
	}

	return &dst
}

func newInstanceDisk(m map[string]any) (*Disk, error) { // nolint:gocognit
	var dst Disk

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ALLOW_ORPHANS":
				dst.AllowOrphans = v
			case "CLONE":
				b, err := api.Str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid CLONE value %q: %w", v, err)
				}
				dst.Clone = b
			case "CLONE_TARGET":
				dst.CloneTarget = v
			case "CLUSTER_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid CLUSTER_ID value %q: %w", v, err)
				}
				dst.ClusterID = n
			case "DATASTORE":
				dst.Datastore = v
			case "DATASTORE_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DATASTORE_ID value %q: %w", v, err)
				}
				dst.DatastoreID = n
			case "DEV_PREFIX":
				dst.DevPrefix = v
			case "DISK_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DISK_ID value %q: %w", v, err)
				}
				dst.DiskID = n
			case "DISK_SNAPSHOT_TOTAL_SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid DISK_SNAPSHOT_TOTAL_SIZE value %q: %w", v, err)
				}
				dst.DiskSnapshotTotalSize = n
			case "DISK_TYPE":
				dst.DiskType = v
			case "DRIVER":
				dst.Driver = v
			case "FORMAT":
				dst.Format = v
			case "IMAGE":
				dst.Image = v
			case "IMAGE_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid IMAGE_ID value %q: %w", v, err)
				}
				dst.ImageID = n
			case "IMAGE_STATE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid IMAGE_STATE value %q: %w", v, err)
				}
				dst.ImageState = n
			case "IMAGE_UNAME":
				dst.ImageUserName = v
			case "LN_TARGET":
				dst.LnTarget = v
			case "ORIGINAL_SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid ORIGINAL_SIZE value %q: %w", v, err)
				}
				dst.OriginalSize = n
			case "PERSISTENT":
				b, err := api.Str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid PERSISTENT value %q: %w", v, err)
				}
				dst.Persistent = b
			case "POOL_NAME":
				dst.PoolName = v
			case "READONLY":
				b, err := api.Str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid READONLY value %q: %w", v, err)
				}
				dst.Readonly = b
			case "SAVE":
				b, err := api.Str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SAVE value %q: %w", v, err)
				}
				dst.Save = b
			case "SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SIZE value %q: %w", v, err)
				}
				dst.Size = n
			case "SOURCE":
				dst.Source = v
			case "TARGET":
				dst.Target = v
			case "TM_MAD":
				dst.TmMAD = v
			case "TM_MAD_SYSTEM":
				dst.TmMADSystem = v
			case "TYPE":
				dst.Type = v
			}
		}
	}

	return &dst, nil
}

func newInstanceNIC(m map[string]any) (*NIC, error) {
	var dst NIC

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			// "AR_ID": "0",
			// "BRIDGE": "mgmt0",
			// "BRIDGE_TYPE": "linux",
			// "CLUSTER_ID": "0",
			// "GATEWAY": "192.168.8.1",
			// "NAME": "NIC0",
			// "NETWORK": "Infrastructure Management Network",
			// "TARGET": "one-9-0",
			// "VLAN_ID": "12",
			// "VN_MAD": "802.1Q"
			case "NIC_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid NIC_ID value %q: %w", v, err)
				}
				dst.ID = n
			case "NETWORK_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid NETWORK_ID value %q: %w", v, err)
				}
				dst.NetworkID = n
			case "IP":
				dst.IP = v
			case "MAC":
				dst.MAC = v
			case "MODEL":
				dst.Model = v
			case "VIRTIO_QUEUES":
				dst.VirtioQueues = v
			case "PhyDev":
				dst.Phydev = v
			case "SECURITY_GROUPS":
				dst.SecurityGroups = v
			}
		}
	}

	return &dst, nil
}

func newInstanceNICAlias(m map[string]any) *NICAlias {
	var dst NICAlias

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ALIAS_ID":
				dst.AliasID = v
			case "PARENT":
				dst.Parent = v
			case "PARENT_ID":
				dst.ParentID = v
			}
		}
	}

	return &dst
}

func newInstanceSchedAction(m map[string]any) *SchedAction {
	var dst SchedAction

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ACTION":
				dst.Action = v
			case "ARGS":
				dst.Args = v
			case "DAYS":
				dst.Days = v
			case "END_TYPE":
				dst.EndType = v
			case "END_VALUE":
				dst.EndValue = v
			case "ID":
				dst.ID = v
			case "REPEAT":
				dst.Repeat = v
			case "TIME":
				dst.Time = v
			case "WARNING":
				dst.Warning = v
			}
		}
	}

	return &dst
}

func newInstanceSnapshot(m map[string]any) (*Snapshot, error) {
	var dst Snapshot

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "ACTION":
				dst.Action = v
			case "ACTIVE":
				b, err := api.Str2Bool(v)
				if err != nil {
					return nil, fmt.Errorf("invalid ACTIVE value %q: %w", v, err)
				}
				dst.Active = b
			case "HYPERVISOR_ID":
				dst.HypervisorID = v
			case "NAME":
				dst.Name = v
			case "SNAPSHOT_ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SNAPSHOT_ID value %q: %w", v, err)
				}
				dst.SnapshotID = n
			case "SYSTEM_DISK_SIZE":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid SYSTEM_DISK_SIZE value %q: %w", v, err)
				}
				dst.SystemDiskSize = n
			case "TIME":
				t, err := strconv.ParseInt(v, 10, 0)
				if err != nil {
					return nil, fmt.Errorf("invalid TIME value %q: %w", v, err)
				}
				dst.Time = time.Unix(t, 0) // TODO: we don't print as json so this should be time.Time (not our wrapper type)
			}
		}
	}

	return &dst, nil
}
