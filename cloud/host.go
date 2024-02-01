package cloud

import (
	"fmt"
	"strconv"
	"strings"
)

// Host is the API payload based on the legacy xmlrpc backend.
type Host struct {
	ID                    int            `json:"id" yaml:"id"`
	Name                  string         `json:"name" yaml:"name"`
	State                 int            `json:"state" yaml:"state"`
	PrevState             int            `json:"prev_state" yaml:"prev_state"`
	InformationManagerMAD string         `json:"information_manager_mad" yaml:"information_manager_mad"`
	InstanceMAD           string         `json:"instance_mad" yaml:"instance_mad"`
	ClusterID             int            `json:"cluster_id" yaml:"cluster_id"`
	Cluster               string         `json:"cluster" yaml:"cluster"`
	HostShare             HostShare      `json:"host_share" yaml:"host_share"`
	Instances             []int          `json:"instances" yaml:"instances"`
	Template              Template       `json:"template" yaml:"template"`
	Monitoring            HostMonitoring `json:"monitoring" yaml:"monitoring"`
}

// HostShare is the API payload based on the legacy xmlrpc backend.
type HostShare struct {
	MemUsage         int        `json:"mem_usage" yaml:"mem_usage"`
	CPUUsage         int        `json:"cpuusage" yaml:"cpuusage"`
	TotalMem         int        `json:"total_mem" yaml:"total_mem"`
	TotalCPU         int        `json:"total_cpu" yaml:"total_cpu"`
	MaxMem           int        `json:"max_mem" yaml:"max_mem"`
	MaxCPU           int        `json:"max_cpu" yaml:"max_cpu"`
	RunningInstances int        `json:"running_instances" yaml:"running_instances"`
	InstancessThread int        `json:"instances_thread" yaml:"instances_thread"`
	Datastores       Datastores `json:"datastores" yaml:"datastores"`
	PCIDevices       []PCI      `json:"pcidevices" yaml:"pcidevices"`
	NumaNodes        []Node     `json:"numa_nodes" yaml:"numa_nodes"`
}

// PCI is the API payload based on the legacy xmlrpc backend.
type PCI struct {
	Address      string `json:"address" yaml:"address"`
	Bus          string `json:"bus" yaml:"bus"`
	Class        string `json:"class" yaml:"class"`
	ClassName    string `json:"class_name" yaml:"class_name"`
	Device       string `json:"device" yaml:"device"`
	DeviceName   string `json:"device_name" yaml:"device_name"`
	Domain       string `json:"domain" yaml:"domain"`
	Function     string `json:"function" yaml:"function"`
	NumaNode     string `json:"numa_node" yaml:"numa_node"`
	ShortAddress string `json:"short_address" yaml:"short_address"`
	Slot         string `json:"slot" yaml:"slot"`
	Type         string `json:"type" yaml:"type"`
	Vendor       string `json:"vendor" yaml:"vendor"`
	VendorName   string `json:"vendor_name" yaml:"vendor_name"`
	InstanceID   int    `json:"instance_id" yaml:"instance_id"`
}

// Node is the API payload based on the legacy xmlrpc backend.
type Node struct {
	Core     []Core     `json:"core" yaml:"core"`
	Hugepage []HugePage `json:"hugepage" yaml:"hugepage"`
	Memory   Memory     `json:"memory" yaml:"memory"`
	NodeID   int        `json:"node_id" yaml:"node_id"`
}

// Capacity is the API payload based on the legacy xmlrpc backend.
type Capacity struct {
	FreeCPU    int `json:"free_cpu" yaml:"free_cpu"`
	FreeMemory int `json:"free_memory" yaml:"free_memory"`
	UsedCPU    int `json:"used_cpu" yaml:"used_cpu"`
	UsedMemory int `json:"used_memory" yaml:"used_memory"`
}

// Core is the API payload based on the legacy xmlrpc backend.
type Core struct {
	CPUs      string `json:"cpus" yaml:"cpus"`
	Dedicated bool   `json:"dedicated" yaml:"dedicated"`
	Free      int    `json:"free" yaml:"free"`
	ID        int    `json:"id" yaml:"id"`
}

// Datastores is the API payload based on the legacy xmlrpc backend.
type Datastores struct {
	DiskUsage int `json:"disk_usage" yaml:"disk_usage"`
	FreeDisk  int `json:"free_disk" yaml:"free_disk"`
	MaxDisk   int `json:"max_disk" yaml:"max_disk"`
	UsedDisk  int `json:"used_disk" yaml:"used_disk"`
}

// HugePage is the API payload based on the legacy xmlrpc backend.
type HugePage struct {
	Free  int `json:"free" yaml:"free"`
	Pages int `json:"pages" yaml:"pages"`
	Size  int `json:"size" yaml:"size"`
	Usage int `json:"usage" yaml:"usage"`
}

// Memory is the API payload based on the legacy xmlrpc backend.
type Memory struct {
	Distance string `json:"distance" yaml:"distance"`
	Free     int    `json:"free" yaml:"free"`
	Total    int    `json:"total" yaml:"total"`
	Usage    int    `json:"usage" yaml:"usage"`
	Used     int    `json:"used" yaml:"used"`
}

// HostMonitoring is the API payload based on the legacy xmlrpc backend.
type HostMonitoring struct {
	Timestamp int      `json:"timestamp" yaml:"timestamp"`
	ID        int      `json:"id" yaml:"id"`
	Capacity  Capacity `json:"capacity" yaml:"capacity"`
	System    System   `json:"system" yaml:"system"`
}

// System is the API payload based on the legacy xmlrpc backend.
type System struct {
	NetRX int `json:"netrx" yaml:"netrx"`
	NetTX int `json:"nettx" yaml:"nettx"`
}

// HostTemplate is the API payload based on the legacy xmlrpc backend.
type HostTemplate struct {
	Arch           string                 `json:"arch" yaml:"arch"`
	CGroupsVersion string                 `json:"cgroups_version" yaml:"cgroups_version"`
	CPUSpeed       int                    `json:"cpuspeed" yaml:"cpuspeed"`
	Hostname       string                 `json:"hostname" yaml:"hostname"`
	Hypervisor     string                 `json:"hypervisor" yaml:"hypervisor"`
	IMMAD          string                 `json:"im_mad" yaml:"im_mad"`
	KVMCPUModel    string                 `json:"kvm_cpumodel" yaml:"kvm_cpumodel"`
	KVMCPUModels   []string               `json:"kvm_cpumodels" yaml:"kvm_cpumodels"`
	KVMMachines    []string               `json:"kvm_machines" yaml:"kvm_machines"`
	ModelName      string                 `json:"modelname" yaml:"modelname"`
	ReservedCPU    string                 `json:"reserved_cpu" yaml:"reserved_cpu"`
	ReservedMem    string                 `json:"reserved_mem" yaml:"reserved_mem"`
	TotalWilds     int                    `json:"total_wilds" yaml:"total_wilds"`
	Version        string                 `json:"version" yaml:"version"`
	Instance       []HostTemplateInstance `json:"instance" yaml:"instance"`
	InstanceMAD    string                 `json:"instance_mad" yaml:"instance_mad"`
	Wilds          string                 `json:"wilds" yaml:"wilds"`
}

// HostTemplateInstance is the API payload based on the legacy xmlrpc backend.
type HostTemplateInstance struct {
	DeployID string `json:"deploy_id" yaml:"deploy_id"`
	ID       int    `json:"id" yaml:"id"`
	Monitor  string `json:"monitor" yaml:"monitor"`
}

// ParseTemplate returns a structured subset of the nested key x value pair map.
func (h *Host) ParseTemplate() (*HostTemplate, error) { // nolint:gocognit
	var t HostTemplate

	for key, value := range h.Template {
		switch v := value.(type) {
		case string:
			switch key {
			case "ARCH":
				t.Arch = v
			case "CGROUPS_VERSION":
				t.CGroupsVersion = v
			case "CPUSPEED":
				i, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				t.CPUSpeed = i
			case "HOSTNAME":
				t.Hostname = v
			case "HYPERVISOR":
				t.Hypervisor = v
			case "IM_MAD":
				t.IMMAD = v
			case "KVM_CPU_MODEL":
				t.KVMCPUModel = v
			case "KVM_CPU_MODELS":
				t.KVMCPUModels = strings.Fields(v)
			case "KVM_MACHINES":
				t.KVMMachines = strings.Fields(v)
			case "MODELNAME":
				t.ModelName = v
			case "RESERVED_CPU":
				t.ReservedCPU = v
			case "RESERVED_MEM":
				t.ReservedMem = v
			case "TOTAL_WILDS":
				i, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				t.TotalWilds = i
			case "VERSION":
				t.Version = v
			case "VM_MAD":
				t.InstanceMAD = v
			case "WILDS":
				t.Wilds = v
			}

		case map[string]any:
			if key == "VM" {
				template, err := newHostTemplateInstance(v)
				if err != nil {
					return nil, err
				}
				t.Instance = []HostTemplateInstance{*template}
			}
		case []map[string]any:
			if key == "VM" {
				for _, m := range v {
					template, err := newHostTemplateInstance(m)
					if err != nil {
						return nil, err
					}
					t.Instance = append(t.Instance, *template)
				}
			}
		}
	}

	return &t, nil
}

func newHostTemplateInstance(m map[string]any) (*HostTemplateInstance, error) {
	var t HostTemplateInstance

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "DEPLOY_ID":
				t.DeployID = v
			case "ID":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("invalid ID value %q: %w", v, err)
				}
				t.ID = n
			case "MONITOR":
				t.Monitor = v
			}
		}
	}

	return &t, nil
}
