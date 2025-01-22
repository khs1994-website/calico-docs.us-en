// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.5.0
// source: api.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FlowReceipt is a response from the server to a client after publishing a Flow.
type FlowReceipt struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowReceipt) Reset() {
	*x = FlowReceipt{}
	mi := &file_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowReceipt) ProtoMessage() {}

func (x *FlowReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowReceipt.ProtoReflect.Descriptor instead.
func (*FlowReceipt) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

// FlowRequest defines a message to request a particular selection of aggregated Flow objects.
type FlowRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// StartTimeGt specifies the beginning of a time window with which to filter Flows. Flows
	// will be returned only if their start time occurs after the requested time.
	StartTimeGt int64 `protobuf:"varint,1,opt,name=start_time_gt,json=startTimeGt,proto3" json:"start_time_gt,omitempty"`
	// StartTimeLt specifies the end of a time window with which to filter flows. Flows will
	// be returned only if their start time occurs before the requested time.
	StartTimeLt int64 `protobuf:"varint,2,opt,name=start_time_lt,json=startTimeLt,proto3" json:"start_time_lt,omitempty"`
	// PageNumber specifies the page number to return. It requires that PageSize is also specified in order
	// to determine page boundaries. Note that pages may change over time as new flow data is collected or expired.
	// Querying the same page at different points in time may return different results.
	PageNumber int64 `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// PageSize configures the maximum number of results to return as part of this query.
	PageSize      int64 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowRequest) Reset() {
	*x = FlowRequest{}
	mi := &file_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowRequest) ProtoMessage() {}

func (x *FlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowRequest.ProtoReflect.Descriptor instead.
func (*FlowRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *FlowRequest) GetStartTimeGt() int64 {
	if x != nil {
		return x.StartTimeGt
	}
	return 0
}

func (x *FlowRequest) GetStartTimeLt() int64 {
	if x != nil {
		return x.StartTimeLt
	}
	return 0
}

func (x *FlowRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *FlowRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// FlowUpdate wraps a Flow with additional metadata.
type FlowUpdate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Flow contains the actual flow being sent.
	Flow          *Flow `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowUpdate) Reset() {
	*x = FlowUpdate{}
	mi := &file_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowUpdate) ProtoMessage() {}

func (x *FlowUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowUpdate.ProtoReflect.Descriptor instead.
func (*FlowUpdate) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *FlowUpdate) GetFlow() *Flow {
	if x != nil {
		return x.Flow
	}
	return nil
}

// FlowKey includes the identifying fields for a Flow.
// - Source: Name, namespace, type, and labels.
// - Destination: Name, namespace, type, labels and port
// - Action taken on the connection.
// - Reporter (i.e., measured at source or destination).
// - Protocol of the connection (TCP, UDP, etc.).
type FlowKey struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// SourceName is the name of the source for this Flow. It represents one or more
	// source pods that share a GenerateName.
	SourceName string `protobuf:"bytes,1,opt,name=source_name,json=sourceName,proto3" json:"source_name,omitempty"`
	// SourceNamespace is the namespace of the source pods for this flow.
	SourceNamespace string `protobuf:"bytes,2,opt,name=source_namespace,json=sourceNamespace,proto3" json:"source_namespace,omitempty"`
	// SourceType is the type of the source, used to contextualize the source
	// name and namespace fields.
	//
	// This can be one of:
	//
	// - wep: WorkloadEndpoint (i.e., Pod)
	// - hep: HostEndpoint
	// - ns: NetworkSet
	// - pub/pvt: External network (source name omitted)
	SourceType string `protobuf:"bytes,3,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	// DestName is the name of the destination for this Flow. It represents one or more
	// destination pods that share a GenerateName.
	DestName string `protobuf:"bytes,4,opt,name=dest_name,json=destName,proto3" json:"dest_name,omitempty"`
	// DestNamespace is the namespace of the destination pods for this flow.
	DestNamespace string `protobuf:"bytes,5,opt,name=dest_namespace,json=destNamespace,proto3" json:"dest_namespace,omitempty"`
	// DestType is the type of the destination, used to contextualize the dest
	// name and namespace fields.
	//
	// This can be one of:
	//
	// - wep: WorkloadEndpoint (i.e., Pod)
	// - hep: HostEndpoint
	// - ns: NetworkSet
	// - pub/pvt: External network (dest name omitted)
	DestType string `protobuf:"bytes,6,opt,name=dest_type,json=destType,proto3" json:"dest_type,omitempty"`
	// DestPort is the destination port on the specified protocol accessed by this flow.
	DestPort int64 `protobuf:"varint,7,opt,name=dest_port,json=destPort,proto3" json:"dest_port,omitempty"`
	// DestServiceName is the name of the destination service, if any.
	DestServiceName string `protobuf:"bytes,8,opt,name=dest_service_name,json=destServiceName,proto3" json:"dest_service_name,omitempty"`
	// DestServiceNamespace is the namespace of the destination service, if any.
	DestServiceNamespace string `protobuf:"bytes,9,opt,name=dest_service_namespace,json=destServiceNamespace,proto3" json:"dest_service_namespace,omitempty"`
	// DestServicePortName is the name of the port on the destination service, if any.
	DestServicePortName string `protobuf:"bytes,10,opt,name=dest_service_port_name,json=destServicePortName,proto3" json:"dest_service_port_name,omitempty"`
	// DestServicePort is the port number on the destination service.
	DestServicePort int64 `protobuf:"varint,11,opt,name=dest_service_port,json=destServicePort,proto3" json:"dest_service_port,omitempty"`
	// Proto is the L4 protocol for this flow. Either TCP or UDP.
	Proto string `protobuf:"bytes,12,opt,name=proto,proto3" json:"proto,omitempty"`
	// Reporter is either "src" or "dst", depending on whether this flow was generated
	// at the initating or terminating end of the connection attempt.
	Reporter string `protobuf:"bytes,13,opt,name=reporter,proto3" json:"reporter,omitempty"`
	// Action is the ultimate action taken on the flow. Either Allow or Drop.
	Action string `protobuf:"bytes,14,opt,name=action,proto3" json:"action,omitempty"`
	// Policies includes an entry for each policy rule that took an action on the connections
	// aggregated into this flow.
	Policies      *FlowLogPolicy `protobuf:"bytes,15,opt,name=policies,proto3" json:"policies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowKey) Reset() {
	*x = FlowKey{}
	mi := &file_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowKey) ProtoMessage() {}

func (x *FlowKey) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowKey.ProtoReflect.Descriptor instead.
func (*FlowKey) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *FlowKey) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *FlowKey) GetSourceNamespace() string {
	if x != nil {
		return x.SourceNamespace
	}
	return ""
}

func (x *FlowKey) GetSourceType() string {
	if x != nil {
		return x.SourceType
	}
	return ""
}

func (x *FlowKey) GetDestName() string {
	if x != nil {
		return x.DestName
	}
	return ""
}

func (x *FlowKey) GetDestNamespace() string {
	if x != nil {
		return x.DestNamespace
	}
	return ""
}

func (x *FlowKey) GetDestType() string {
	if x != nil {
		return x.DestType
	}
	return ""
}

func (x *FlowKey) GetDestPort() int64 {
	if x != nil {
		return x.DestPort
	}
	return 0
}

func (x *FlowKey) GetDestServiceName() string {
	if x != nil {
		return x.DestServiceName
	}
	return ""
}

func (x *FlowKey) GetDestServiceNamespace() string {
	if x != nil {
		return x.DestServiceNamespace
	}
	return ""
}

func (x *FlowKey) GetDestServicePortName() string {
	if x != nil {
		return x.DestServicePortName
	}
	return ""
}

func (x *FlowKey) GetDestServicePort() int64 {
	if x != nil {
		return x.DestServicePort
	}
	return 0
}

func (x *FlowKey) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *FlowKey) GetReporter() string {
	if x != nil {
		return x.Reporter
	}
	return ""
}

func (x *FlowKey) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FlowKey) GetPolicies() *FlowLogPolicy {
	if x != nil {
		return x.Policies
	}
	return nil
}

// Flow is a message representing statistics gathered about connections that share common fields,
// aggregated across either time, nodes, or both.
type Flow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Key includes the identifying fields for this flow.
	Key *FlowKey `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// StartTime is the start time for this flow. It is represented as the number of
	// seconds since the UNIX epoch.
	StartTime int64 `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// EndTime is the end time for this flow. It is always exactly one aggregation
	// interval after the start time.
	EndTime int64 `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// SourceLabels contains the intersection of labels that appear on all source
	// pods that contributed to this flow.
	SourceLabels []string `protobuf:"bytes,4,rep,name=source_labels,json=sourceLabels,proto3" json:"source_labels,omitempty"`
	// SourceLabels contains the intersection of labels that appear on all destination
	// pods that contributed to this flow.
	DestLabels []string `protobuf:"bytes,5,rep,name=dest_labels,json=destLabels,proto3" json:"dest_labels,omitempty"`
	// Statistics.
	PacketsIn  int64 `protobuf:"varint,6,opt,name=packets_in,json=packetsIn,proto3" json:"packets_in,omitempty"`
	PacketsOut int64 `protobuf:"varint,7,opt,name=packets_out,json=packetsOut,proto3" json:"packets_out,omitempty"`
	BytesIn    int64 `protobuf:"varint,8,opt,name=bytes_in,json=bytesIn,proto3" json:"bytes_in,omitempty"`
	BytesOut   int64 `protobuf:"varint,9,opt,name=bytes_out,json=bytesOut,proto3" json:"bytes_out,omitempty"`
	// NumConnectionsStarted tracks the total number of new connections recorded for this Flow. It counts each
	// connection attempt that matches the FlowKey that was made between this Flow's StartTime and EndTime.
	NumConnectionsStarted int64 `protobuf:"varint,10,opt,name=num_connections_started,json=numConnectionsStarted,proto3" json:"num_connections_started,omitempty"`
	// NumConnectionsCompleted tracks the total number of completed TCP connections recorded for this Flow. It counts each
	// connection that matches the FlowKey that was completed between this Flow's StartTime and EndTime.
	NumConnectionsCompleted int64 `protobuf:"varint,11,opt,name=num_connections_completed,json=numConnectionsCompleted,proto3" json:"num_connections_completed,omitempty"`
	// NumConnectionsLive tracks the total number of still active connections recorded for this Flow. It counts each
	// connection that matches the FlowKey that was active at this Flow's EndTime.
	NumConnectionsLive int64 `protobuf:"varint,12,opt,name=num_connections_live,json=numConnectionsLive,proto3" json:"num_connections_live,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Flow) Reset() {
	*x = Flow{}
	mi := &file_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Flow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flow) ProtoMessage() {}

func (x *Flow) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flow.ProtoReflect.Descriptor instead.
func (*Flow) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Flow) GetKey() *FlowKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Flow) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Flow) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Flow) GetSourceLabels() []string {
	if x != nil {
		return x.SourceLabels
	}
	return nil
}

func (x *Flow) GetDestLabels() []string {
	if x != nil {
		return x.DestLabels
	}
	return nil
}

func (x *Flow) GetPacketsIn() int64 {
	if x != nil {
		return x.PacketsIn
	}
	return 0
}

func (x *Flow) GetPacketsOut() int64 {
	if x != nil {
		return x.PacketsOut
	}
	return 0
}

func (x *Flow) GetBytesIn() int64 {
	if x != nil {
		return x.BytesIn
	}
	return 0
}

func (x *Flow) GetBytesOut() int64 {
	if x != nil {
		return x.BytesOut
	}
	return 0
}

func (x *Flow) GetNumConnectionsStarted() int64 {
	if x != nil {
		return x.NumConnectionsStarted
	}
	return 0
}

func (x *Flow) GetNumConnectionsCompleted() int64 {
	if x != nil {
		return x.NumConnectionsCompleted
	}
	return 0
}

func (x *Flow) GetNumConnectionsLive() int64 {
	if x != nil {
		return x.NumConnectionsLive
	}
	return 0
}

type FlowLogPolicy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// AllPolicies is a list of strings containing policy rule information.
	AllPolicies   []string `protobuf:"bytes,1,rep,name=all_policies,json=allPolicies,proto3" json:"all_policies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlowLogPolicy) Reset() {
	*x = FlowLogPolicy{}
	mi := &file_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlowLogPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowLogPolicy) ProtoMessage() {}

func (x *FlowLogPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowLogPolicy.ProtoReflect.Descriptor instead.
func (*FlowLogPolicy) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *FlowLogPolicy) GetAllPolicies() []string {
	if x != nil {
		return x.AllPolicies
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x65, 0x6c,
	0x69, 0x78, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x67, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x47, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2d, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x77, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77,
	0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xb3, 0x04, 0x0a, 0x07, 0x46, 0x6c, 0x6f, 0x77, 0x4b,
	0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x33, 0x0a, 0x16, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x64, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x64, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66,
	0x65, 0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0xc6, 0x03, 0x0a,
	0x04, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x20, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4b,
	0x65, 0x79, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x73,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x49, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12,
	0x36, 0x0a, 0x17, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x15, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x6e, 0x75, 0x6d, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6e, 0x75, 0x6d, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4c, 0x69, 0x76, 0x65, 0x22, 0x32, 0x0a, 0x0d, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c,
	0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x32, 0x34, 0x0a, 0x07, 0x46, 0x6c, 0x6f,
	0x77, 0x41, 0x50, 0x49, 0x12, 0x29, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x66,
	0x65, 0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x30, 0x01, 0x32,
	0x45, 0x0a, 0x0d, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x66, 0x65,
	0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x12,
	0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []any{
	(*FlowReceipt)(nil),   // 0: felix.FlowReceipt
	(*FlowRequest)(nil),   // 1: felix.FlowRequest
	(*FlowUpdate)(nil),    // 2: felix.FlowUpdate
	(*FlowKey)(nil),       // 3: felix.FlowKey
	(*Flow)(nil),          // 4: felix.Flow
	(*FlowLogPolicy)(nil), // 5: felix.FlowLogPolicy
}
var file_api_proto_depIdxs = []int32{
	4, // 0: felix.FlowUpdate.flow:type_name -> felix.Flow
	5, // 1: felix.FlowKey.policies:type_name -> felix.FlowLogPolicy
	3, // 2: felix.Flow.Key:type_name -> felix.FlowKey
	1, // 3: felix.FlowAPI.List:input_type -> felix.FlowRequest
	2, // 4: felix.FlowCollector.Connect:input_type -> felix.FlowUpdate
	4, // 5: felix.FlowAPI.List:output_type -> felix.Flow
	0, // 6: felix.FlowCollector.Connect:output_type -> felix.FlowReceipt
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
