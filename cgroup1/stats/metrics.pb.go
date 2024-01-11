// Code generated by protoc-gen-go.
// source: github.com/containerd/cgroups/cgroup1/stats/metrics.proto
// DO NOT EDIT!

/*
Package github_com_containerd_cgroups_cgroup1_stats is a generated protocol buffer package.

It is generated from these files:

	github.com/containerd/cgroups/cgroup1/stats/metrics.proto

It has these top-level messages:

	Metrics
	HugetlbStat
	PidsStat
	CPUStat
	CPUUsage
	Throttle
	CFSConfig
	MemoryStat
	MemoryEntry
	MemoryOomControl
	BlkIOStat
	BlkIOEntry
	RdmaStat
	RdmaEntry
	NetworkStat
	CgroupStats
*/
package stats

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Metrics struct {
	Hugetlb          []*HugetlbStat    `protobuf:"bytes,1,rep,name=hugetlb" json:"hugetlb,omitempty"`
	Pids             *PidsStat         `protobuf:"bytes,2,opt,name=pids" json:"pids,omitempty"`
	CPU              *CPUStat          `protobuf:"bytes,3,opt,name=cpu" json:"cpu,omitempty"`
	Memory           *MemoryStat       `protobuf:"bytes,4,opt,name=memory" json:"memory,omitempty"`
	Blkio            *BlkIOStat        `protobuf:"bytes,5,opt,name=blkio" json:"blkio,omitempty"`
	Rdma             *RdmaStat         `protobuf:"bytes,6,opt,name=rdma" json:"rdma,omitempty"`
	Network          []*NetworkStat    `protobuf:"bytes,7,rep,name=network" json:"network,omitempty"`
	CgroupStats      *CgroupStats      `protobuf:"bytes,8,opt,name=cgroup_stats" json:"cgroup_stats,omitempty"`
	MemoryOomControl *MemoryOomControl `protobuf:"bytes,9,opt,name=memory_oom_control" json:"memory_oom_control,omitempty"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}

func (m *Metrics) GetHugetlb() []*HugetlbStat {
	if m != nil {
		return m.Hugetlb
	}
	return nil
}

func (m *Metrics) GetPids() *PidsStat {
	if m != nil {
		return m.Pids
	}
	return nil
}

func (m *Metrics) GetCPU() *CPUStat {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *Metrics) GetMemory() *MemoryStat {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Metrics) GetBlkio() *BlkIOStat {
	if m != nil {
		return m.Blkio
	}
	return nil
}

func (m *Metrics) GetRdma() *RdmaStat {
	if m != nil {
		return m.Rdma
	}
	return nil
}

func (m *Metrics) GetNetwork() []*NetworkStat {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Metrics) GetCgroupStats() *CgroupStats {
	if m != nil {
		return m.CgroupStats
	}
	return nil
}

func (m *Metrics) GetMemoryOomControl() *MemoryOomControl {
	if m != nil {
		return m.MemoryOomControl
	}
	return nil
}

type HugetlbStat struct {
	Usage    uint64 `protobuf:"varint,1,opt,name=usage" json:"usage,omitempty"`
	Max      uint64 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Failcnt  uint64 `protobuf:"varint,3,opt,name=failcnt" json:"failcnt,omitempty"`
	Pagesize string `protobuf:"bytes,4,opt,name=pagesize" json:"pagesize,omitempty"`
}

func (m *HugetlbStat) Reset()         { *m = HugetlbStat{} }
func (m *HugetlbStat) String() string { return proto.CompactTextString(m) }
func (*HugetlbStat) ProtoMessage()    {}

type PidsStat struct {
	Current uint64 `protobuf:"varint,1,opt,name=current" json:"current,omitempty"`
	Limit   uint64 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *PidsStat) Reset()         { *m = PidsStat{} }
func (m *PidsStat) String() string { return proto.CompactTextString(m) }
func (*PidsStat) ProtoMessage()    {}

type CPUStat struct {
	Usage      *CPUUsage  `protobuf:"bytes,1,opt,name=usage" json:"usage,omitempty"`
	Throttling *Throttle  `protobuf:"bytes,2,opt,name=throttling" json:"throttling,omitempty"`
	CfsConfig  *CFSConfig `protobuf:"bytes,3,opt,name=cfs_config" json:"cfs_config,omitempty"`
}

func (m *CPUStat) Reset()         { *m = CPUStat{} }
func (m *CPUStat) String() string { return proto.CompactTextString(m) }
func (*CPUStat) ProtoMessage()    {}

func (m *CPUStat) GetUsage() *CPUUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *CPUStat) GetThrottling() *Throttle {
	if m != nil {
		return m.Throttling
	}
	return nil
}

func (m *CPUStat) GetCfsConfig() *CFSConfig {
	if m != nil {
		return m.CfsConfig
	}
	return nil
}

type CPUUsage struct {
	// values in nanoseconds
	Total  uint64   `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	Kernel uint64   `protobuf:"varint,2,opt,name=kernel" json:"kernel,omitempty"`
	User   uint64   `protobuf:"varint,3,opt,name=user" json:"user,omitempty"`
	PerCPU []uint64 `protobuf:"varint,4,rep,name=per_cpu" json:"per_cpu,omitempty"`
}

func (m *CPUUsage) Reset()         { *m = CPUUsage{} }
func (m *CPUUsage) String() string { return proto.CompactTextString(m) }
func (*CPUUsage) ProtoMessage()    {}

type Throttle struct {
	Periods          uint64 `protobuf:"varint,1,opt,name=periods" json:"periods,omitempty"`
	ThrottledPeriods uint64 `protobuf:"varint,2,opt,name=throttled_periods" json:"throttled_periods,omitempty"`
	ThrottledTime    uint64 `protobuf:"varint,3,opt,name=throttled_time" json:"throttled_time,omitempty"`
}

func (m *Throttle) Reset()         { *m = Throttle{} }
func (m *Throttle) String() string { return proto.CompactTextString(m) }
func (*Throttle) ProtoMessage()    {}

type CFSConfig struct {
	BurstUs  uint64 `protobuf:"varint,1,opt,name=burst_us" json:"burst_us,omitempty"`
	PeriodUs uint64 `protobuf:"varint,2,opt,name=period_us" json:"period_us,omitempty"`
	QuotaUs  int64  `protobuf:"varint,3,opt,name=quota_us" json:"quota_us,omitempty"`
}

func (m *CFSConfig) Reset()         { *m = CFSConfig{} }
func (m *CFSConfig) String() string { return proto.CompactTextString(m) }
func (*CFSConfig) ProtoMessage()    {}

type MemoryStat struct {
	Cache                   uint64       `protobuf:"varint,1,opt,name=cache" json:"cache,omitempty"`
	RSS                     uint64       `protobuf:"varint,2,opt,name=rss" json:"rss,omitempty"`
	RSSHuge                 uint64       `protobuf:"varint,3,opt,name=rss_huge" json:"rss_huge,omitempty"`
	MappedFile              uint64       `protobuf:"varint,4,opt,name=mapped_file" json:"mapped_file,omitempty"`
	Dirty                   uint64       `protobuf:"varint,5,opt,name=dirty" json:"dirty,omitempty"`
	Writeback               uint64       `protobuf:"varint,6,opt,name=writeback" json:"writeback,omitempty"`
	PgPgIn                  uint64       `protobuf:"varint,7,opt,name=pg_pg_in" json:"pg_pg_in,omitempty"`
	PgPgOut                 uint64       `protobuf:"varint,8,opt,name=pg_pg_out" json:"pg_pg_out,omitempty"`
	PgFault                 uint64       `protobuf:"varint,9,opt,name=pg_fault" json:"pg_fault,omitempty"`
	PgMajFault              uint64       `protobuf:"varint,10,opt,name=pg_maj_fault" json:"pg_maj_fault,omitempty"`
	InactiveAnon            uint64       `protobuf:"varint,11,opt,name=inactive_anon" json:"inactive_anon,omitempty"`
	ActiveAnon              uint64       `protobuf:"varint,12,opt,name=active_anon" json:"active_anon,omitempty"`
	InactiveFile            uint64       `protobuf:"varint,13,opt,name=inactive_file" json:"inactive_file,omitempty"`
	ActiveFile              uint64       `protobuf:"varint,14,opt,name=active_file" json:"active_file,omitempty"`
	Unevictable             uint64       `protobuf:"varint,15,opt,name=unevictable" json:"unevictable,omitempty"`
	HierarchicalMemoryLimit uint64       `protobuf:"varint,16,opt,name=hierarchical_memory_limit" json:"hierarchical_memory_limit,omitempty"`
	HierarchicalSwapLimit   uint64       `protobuf:"varint,17,opt,name=hierarchical_swap_limit" json:"hierarchical_swap_limit,omitempty"`
	TotalCache              uint64       `protobuf:"varint,18,opt,name=total_cache" json:"total_cache,omitempty"`
	TotalRSS                uint64       `protobuf:"varint,19,opt,name=total_rss" json:"total_rss,omitempty"`
	TotalRSSHuge            uint64       `protobuf:"varint,20,opt,name=total_rss_huge" json:"total_rss_huge,omitempty"`
	TotalMappedFile         uint64       `protobuf:"varint,21,opt,name=total_mapped_file" json:"total_mapped_file,omitempty"`
	TotalDirty              uint64       `protobuf:"varint,22,opt,name=total_dirty" json:"total_dirty,omitempty"`
	TotalWriteback          uint64       `protobuf:"varint,23,opt,name=total_writeback" json:"total_writeback,omitempty"`
	TotalPgPgIn             uint64       `protobuf:"varint,24,opt,name=total_pg_pg_in" json:"total_pg_pg_in,omitempty"`
	TotalPgPgOut            uint64       `protobuf:"varint,25,opt,name=total_pg_pg_out" json:"total_pg_pg_out,omitempty"`
	TotalPgFault            uint64       `protobuf:"varint,26,opt,name=total_pg_fault" json:"total_pg_fault,omitempty"`
	TotalPgMajFault         uint64       `protobuf:"varint,27,opt,name=total_pg_maj_fault" json:"total_pg_maj_fault,omitempty"`
	TotalInactiveAnon       uint64       `protobuf:"varint,28,opt,name=total_inactive_anon" json:"total_inactive_anon,omitempty"`
	TotalActiveAnon         uint64       `protobuf:"varint,29,opt,name=total_active_anon" json:"total_active_anon,omitempty"`
	TotalInactiveFile       uint64       `protobuf:"varint,30,opt,name=total_inactive_file" json:"total_inactive_file,omitempty"`
	TotalActiveFile         uint64       `protobuf:"varint,31,opt,name=total_active_file" json:"total_active_file,omitempty"`
	TotalUnevictable        uint64       `protobuf:"varint,32,opt,name=total_unevictable" json:"total_unevictable,omitempty"`
	Usage                   *MemoryEntry `protobuf:"bytes,33,opt,name=usage" json:"usage,omitempty"`
	Swap                    *MemoryEntry `protobuf:"bytes,34,opt,name=swap" json:"swap,omitempty"`
	Kernel                  *MemoryEntry `protobuf:"bytes,35,opt,name=kernel" json:"kernel,omitempty"`
	KernelTCP               *MemoryEntry `protobuf:"bytes,36,opt,name=kernel_tcp" json:"kernel_tcp,omitempty"`
}

func (m *MemoryStat) Reset()         { *m = MemoryStat{} }
func (m *MemoryStat) String() string { return proto.CompactTextString(m) }
func (*MemoryStat) ProtoMessage()    {}

func (m *MemoryStat) GetUsage() *MemoryEntry {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *MemoryStat) GetSwap() *MemoryEntry {
	if m != nil {
		return m.Swap
	}
	return nil
}

func (m *MemoryStat) GetKernel() *MemoryEntry {
	if m != nil {
		return m.Kernel
	}
	return nil
}

func (m *MemoryStat) GetKernelTCP() *MemoryEntry {
	if m != nil {
		return m.KernelTCP
	}
	return nil
}

type MemoryEntry struct {
	Limit   uint64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Usage   uint64 `protobuf:"varint,2,opt,name=usage" json:"usage,omitempty"`
	Max     uint64 `protobuf:"varint,3,opt,name=max" json:"max,omitempty"`
	Failcnt uint64 `protobuf:"varint,4,opt,name=failcnt" json:"failcnt,omitempty"`
}

func (m *MemoryEntry) Reset()         { *m = MemoryEntry{} }
func (m *MemoryEntry) String() string { return proto.CompactTextString(m) }
func (*MemoryEntry) ProtoMessage()    {}

type MemoryOomControl struct {
	OomKillDisable uint64 `protobuf:"varint,1,opt,name=oom_kill_disable" json:"oom_kill_disable,omitempty"`
	UnderOom       uint64 `protobuf:"varint,2,opt,name=under_oom" json:"under_oom,omitempty"`
	OomKill        uint64 `protobuf:"varint,3,opt,name=oom_kill" json:"oom_kill,omitempty"`
}

func (m *MemoryOomControl) Reset()         { *m = MemoryOomControl{} }
func (m *MemoryOomControl) String() string { return proto.CompactTextString(m) }
func (*MemoryOomControl) ProtoMessage()    {}

type BlkIOStat struct {
	IoServiceBytesRecursive []*BlkIOEntry `protobuf:"bytes,1,rep,name=io_service_bytes_recursive" json:"io_service_bytes_recursive,omitempty"`
	IoServicedRecursive     []*BlkIOEntry `protobuf:"bytes,2,rep,name=io_serviced_recursive" json:"io_serviced_recursive,omitempty"`
	IoQueuedRecursive       []*BlkIOEntry `protobuf:"bytes,3,rep,name=io_queued_recursive" json:"io_queued_recursive,omitempty"`
	IoServiceTimeRecursive  []*BlkIOEntry `protobuf:"bytes,4,rep,name=io_service_time_recursive" json:"io_service_time_recursive,omitempty"`
	IoWaitTimeRecursive     []*BlkIOEntry `protobuf:"bytes,5,rep,name=io_wait_time_recursive" json:"io_wait_time_recursive,omitempty"`
	IoMergedRecursive       []*BlkIOEntry `protobuf:"bytes,6,rep,name=io_merged_recursive" json:"io_merged_recursive,omitempty"`
	IoTimeRecursive         []*BlkIOEntry `protobuf:"bytes,7,rep,name=io_time_recursive" json:"io_time_recursive,omitempty"`
	SectorsRecursive        []*BlkIOEntry `protobuf:"bytes,8,rep,name=sectors_recursive" json:"sectors_recursive,omitempty"`
}

func (m *BlkIOStat) Reset()         { *m = BlkIOStat{} }
func (m *BlkIOStat) String() string { return proto.CompactTextString(m) }
func (*BlkIOStat) ProtoMessage()    {}

func (m *BlkIOStat) GetIoServiceBytesRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoServiceBytesRecursive
	}
	return nil
}

func (m *BlkIOStat) GetIoServicedRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoServicedRecursive
	}
	return nil
}

func (m *BlkIOStat) GetIoQueuedRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoQueuedRecursive
	}
	return nil
}

func (m *BlkIOStat) GetIoServiceTimeRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoServiceTimeRecursive
	}
	return nil
}

func (m *BlkIOStat) GetIoWaitTimeRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoWaitTimeRecursive
	}
	return nil
}

func (m *BlkIOStat) GetIoMergedRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoMergedRecursive
	}
	return nil
}

func (m *BlkIOStat) GetIoTimeRecursive() []*BlkIOEntry {
	if m != nil {
		return m.IoTimeRecursive
	}
	return nil
}

func (m *BlkIOStat) GetSectorsRecursive() []*BlkIOEntry {
	if m != nil {
		return m.SectorsRecursive
	}
	return nil
}

type BlkIOEntry struct {
	Op     string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Device string `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
	Major  uint64 `protobuf:"varint,3,opt,name=major" json:"major,omitempty"`
	Minor  uint64 `protobuf:"varint,4,opt,name=minor" json:"minor,omitempty"`
	Value  uint64 `protobuf:"varint,5,opt,name=value" json:"value,omitempty"`
}

func (m *BlkIOEntry) Reset()         { *m = BlkIOEntry{} }
func (m *BlkIOEntry) String() string { return proto.CompactTextString(m) }
func (*BlkIOEntry) ProtoMessage()    {}

type RdmaStat struct {
	Current []*RdmaEntry `protobuf:"bytes,1,rep,name=current" json:"current,omitempty"`
	Limit   []*RdmaEntry `protobuf:"bytes,2,rep,name=limit" json:"limit,omitempty"`
}

func (m *RdmaStat) Reset()         { *m = RdmaStat{} }
func (m *RdmaStat) String() string { return proto.CompactTextString(m) }
func (*RdmaStat) ProtoMessage()    {}

func (m *RdmaStat) GetCurrent() []*RdmaEntry {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *RdmaStat) GetLimit() []*RdmaEntry {
	if m != nil {
		return m.Limit
	}
	return nil
}

type RdmaEntry struct {
	Device     string `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	HcaHandles uint32 `protobuf:"varint,2,opt,name=hca_handles" json:"hca_handles,omitempty"`
	HcaObjects uint32 `protobuf:"varint,3,opt,name=hca_objects" json:"hca_objects,omitempty"`
}

func (m *RdmaEntry) Reset()         { *m = RdmaEntry{} }
func (m *RdmaEntry) String() string { return proto.CompactTextString(m) }
func (*RdmaEntry) ProtoMessage()    {}

type NetworkStat struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RxBytes   uint64 `protobuf:"varint,2,opt,name=rx_bytes" json:"rx_bytes,omitempty"`
	RxPackets uint64 `protobuf:"varint,3,opt,name=rx_packets" json:"rx_packets,omitempty"`
	RxErrors  uint64 `protobuf:"varint,4,opt,name=rx_errors" json:"rx_errors,omitempty"`
	RxDropped uint64 `protobuf:"varint,5,opt,name=rx_dropped" json:"rx_dropped,omitempty"`
	TxBytes   uint64 `protobuf:"varint,6,opt,name=tx_bytes" json:"tx_bytes,omitempty"`
	TxPackets uint64 `protobuf:"varint,7,opt,name=tx_packets" json:"tx_packets,omitempty"`
	TxErrors  uint64 `protobuf:"varint,8,opt,name=tx_errors" json:"tx_errors,omitempty"`
	TxDropped uint64 `protobuf:"varint,9,opt,name=tx_dropped" json:"tx_dropped,omitempty"`
}

func (m *NetworkStat) Reset()         { *m = NetworkStat{} }
func (m *NetworkStat) String() string { return proto.CompactTextString(m) }
func (*NetworkStat) ProtoMessage()    {}

// CgroupStats exports per-cgroup statistics.
type CgroupStats struct {
	// number of tasks sleeping
	NrSleeping uint64 `protobuf:"varint,1,opt,name=nr_sleeping" json:"nr_sleeping,omitempty"`
	// number of tasks running
	NrRunning uint64 `protobuf:"varint,2,opt,name=nr_running" json:"nr_running,omitempty"`
	// number of tasks in stopped state
	NrStopped uint64 `protobuf:"varint,3,opt,name=nr_stopped" json:"nr_stopped,omitempty"`
	// number of tasks in uninterruptible state
	NrUninterruptible uint64 `protobuf:"varint,4,opt,name=nr_uninterruptible" json:"nr_uninterruptible,omitempty"`
	// number of tasks waiting on IO
	NrIoWait uint64 `protobuf:"varint,5,opt,name=nr_io_wait" json:"nr_io_wait,omitempty"`
}

func (m *CgroupStats) Reset()         { *m = CgroupStats{} }
func (m *CgroupStats) String() string { return proto.CompactTextString(m) }
func (*CgroupStats) ProtoMessage()    {}
