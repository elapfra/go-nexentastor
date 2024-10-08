package ns

import (
	"strings"
	"time"
)

// ACLRuleSet - filesystem ACL rule set
type ACLRuleSet int64

const (
	// ACLReadOnly - apply read only set of rules to filesystem
	ACLReadOnly ACLRuleSet = iota

	// ACLReadWrite - apply full access set of rules to filesystem
	ACLReadWrite
)

// License - NexentaStor license
type License struct {
	Valid   bool   `json:"valid"`
	Expires string `json:"expires"`
}

// Filesystem - NexentaStor filesystem
type Filesystem struct {
	Path           string `json:"path"`
	MountPoint     string `json:"mountPoint"`
	SharedOverNfs  bool   `json:"sharedOverNfs"`
	SharedOverSmb  bool   `json:"sharedOverSmb"`
	BytesAvailable int64  `json:"bytesAvailable"`
	BytesUsed      int64  `json:"bytesUsed"`
}

// Service response - NexentaStor /rsf/clusters
type Service struct {
	Size        int64    `json:"size"`
	ServiceName string   `json:"serviceName"`
	Status      []Status `json:"status"`
}

// Health response - NexentaStor /rsf/clusters
type Health struct {
	ServicesHealth          string `json:"servicesHealth"`
	ClusterHealth           string `json:"clusterHealth"`
	NetworkHeartbeatsHealth string `json:"networkHeartbeatsHealth"`
	NodesHealth             string `json:"nodesHealth"`
}

// Status response - NexentaStor /rsf/clusters
type Status struct {
	Node      string `json:"node"`
	Unblocked bool   `json:"unblocked"`
	Status    string `json:"status"`
}

// Volume - NexentaStor volume
type Volume struct {
	Path 		string `json:"path"`
	BytesAvailable int64  `json:"bytesAvailable"`
	BytesUsed      int64  `json:"bytesUsed"`
	VolumeSize     int64  `json:"volumeSize"`
}

// VolumeGroup - NexentaStor volumeGroup
type VolumeGroup struct {
    Path           string `json:"path"`
    BytesAvailable int64  `json:"bytesAvailable"`
    BytesUsed      int64  `json:"bytesUsed"`
}

// LunMapping - NexentaStor lunmapping
type LunMapping struct {
	Id			string `json:"id"`
	Volume      string `json:"volume"`
	TargetGroup string `json:"targetGroup"`
	HostGroup   string `json:"hostGroup"`
	Lun 		int    `json:"lun"`
}

// RemoteInitiator - NexentaStor remote initiator for CHAP access
type RemoteInitiator struct {
	Name             string `json:"name"`
	ChapUser   		 string `json:"chapUser"`
	ChapSecretSet    bool   `json:"chapSecretSet"`
}

// ISCSITarget - NexentaStor iSCSI target
type ISCSITarget struct {
	Name 				string
	State 				string
	Authentication 		string
	Alias 				string
	ChapSecretSet 		bool
	ChapUser 			string
	Portals     		[]Portal
}

// LogicalUnit - NexentaStor logicalUnit
type LogicalUnit struct {
	Guid                   string `json:"guid"`
	Alias                  string `json:"alias"`
	Volume                 string `json:"volume"`
	VolSize                int    `json:"volSize"`
	BlockSize              int    `json:"blockSize"`
	WriteProtect           bool   `json:"writeProtect"`
	WritebackCacheDisabled bool   `json:"writebackCacheDisabled"`
	State                  string `json:"state"`
	AccessState            string `json:"accessState"`
	MappingCount           int    `json:"mappingCount"`
	ExposedOverIscsi       bool   `json:"exposedOverIscsi"`
	ExposedOverFC          bool   `json:"exposedOverFC"`
	Href                   string `json:"href"`
}

func (fs *Filesystem) String() string {
	return fs.Path
}

// GetDefaultSmbShareName - get default SMB share name (all slashes get replaced by underscore)
// Converts '/pool/dataset/fs' to 'pool_dataset_fs'
func (fs *Filesystem) GetDefaultSmbShareName() string {
	return strings.Replace(strings.TrimPrefix(fs.Path, "/"), "/", "_", -1)
}

// GetReferencedQuotaSize - get total referenced quota size
func (fs *Filesystem) GetReferencedQuotaSize() int64 {
	return fs.BytesAvailable + fs.BytesUsed
}

// Snapshot - NexentaStor snapshot
type Snapshot struct {
	Path         string    `json:"path"`
	Name         string    `json:"name"`
	Parent       string    `json:"parent"`
	Clones       []string  `json:"clones"`
	CreationTxg  string    `json:"creationTxg"`
	CreationTime time.Time `json:"creationTime"`
}

func (snapshot *Snapshot) String() string {
	return snapshot.Path
}

// RSFCluster - RSF cluster with a name
type RSFCluster struct {
	Name string `json:"clusterName"`
	Services []Service `json:"services"`
	Health   Health    `json:"health"`
}

// Pool - NS pool
type Pool struct {
	Name string `json:"poolName"`
}

type TargetGroup struct {
	Name string 		`json:"name"`
	Members []string 	`json:"members"`
}

// NEF request/response types

type nefAuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type nefAuthLoginResponse struct {
	Token string `json:"token"`
}

type nefStoragePoolsResponse struct {
	Data []Pool `json:"data"`
}

type nefStorageFilesystemsResponse struct {
	Data []Filesystem `json:"data"`
}

type nefStorageVolumesResponse struct {
    Data []Volume `json:"data"`
}

type nefSanLogicalUnitsResponse struct {
	Data []LogicalUnit `json:"data"`
}

type nefStorageVolumeGroupsResponse struct {
    Data []VolumeGroup `json:"data"`
}

type nefLunMappingsResponse struct {
	Data[]LunMapping `json:"data"`
}

type nefStorageSnapshotsResponse struct {
	Data []Snapshot `json:"data"`
}

type nefNasNfsRequest struct {
	Filesystem       string                            `json:"filesystem"`
	Anon             string                            `json:"anon"`
	SecurityContexts []nefNasNfsRequestSecurityContext `json:"securityContexts"`
}
type nefNasNfsRequestSecurityContext struct {
	SecurityModes []string 							`json:"securityModes"`
	ReadWriteList []NfsRuleList						`json:"readWriteList"`
	ReadOnlyList  []NfsRuleList						`json:"readOnlyList"`
}

type NfsRuleList struct {
	Etype 	string 		`json:"etype"`
	Entity 	string 		`json:"entity"`
	Mask	int 		`json:"mask"`
}

type Portal struct {
	Address string `json:"address"`
	Port 	int    `json:"port"`
}

type nefNasSmbResponse struct {
	ShareName string `json:"shareName"`
}

type nefStorageFilesystemsACLRequest struct {
	Type        string   `json:"type"`
	Principal   string   `json:"principal"`
	Flags       []string `json:"flags"`
	Permissions []string `json:"permissions"`
}

type nefRsfClustersResponse struct {
	Data []RSFCluster `json:"data"`
}

type nefJobStatusResponse struct {
	Links []nefJobStatusResponseLink `json:"links"`
}
type nefJobStatusResponseLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type nefHostGroup struct {
	Members 	[]string  `json:"members"`
	Name 		string 	  `json:"name"`
}

type nefHostGroupsResponse struct {
	Data 	[]nefHostGroup  `json:"data"`
}

type nefTargetGroupsResponse struct {
	Data 	[]TargetGroup  `json:"data"`
}

type nefTargetsResponse struct {
	Data 	[]ISCSITarget  `json:"data"`
}
