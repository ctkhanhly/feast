//
// Copyright 2020 The Feast Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: feast/core/FeatureView.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FeatureView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User-specified specifications of this feature view.
	Spec *FeatureViewSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// System-populated metadata for this feature view.
	Meta *FeatureViewMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *FeatureView) Reset() {
	*x = FeatureView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_FeatureView_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureView) ProtoMessage() {}

func (x *FeatureView) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_FeatureView_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureView.ProtoReflect.Descriptor instead.
func (*FeatureView) Descriptor() ([]byte, []int) {
	return file_feast_core_FeatureView_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureView) GetSpec() *FeatureViewSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FeatureView) GetMeta() *FeatureViewMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// TODO(adchia): refactor common fields from this and ODFV into separate metadata proto
type FeatureViewSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the feature view. Must be unique. Not updated.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of Feast project that this feature view belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// List names of entities to associate with the Features defined in this
	// Feature View. Not updatable.
	Entities []string `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	// List of features specifications for each feature defined with this feature view.
	Features []*FeatureSpecV2 `protobuf:"bytes,4,rep,name=features,proto3" json:"features,omitempty"`
	// User defined metadata
	Tags map[string]string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Features in this feature view can only be retrieved from online serving
	// younger than ttl. Ttl is measured as the duration of time between
	// the feature's event timestamp and when the feature is retrieved
	// Feature values outside ttl will be returned as unset values and indicated to end user
	Ttl *durationpb.Duration `protobuf:"bytes,6,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// Batch/Offline DataSource where this view can retrieve offline feature data.
	BatchSource *DataSource `protobuf:"bytes,7,opt,name=batch_source,json=batchSource,proto3" json:"batch_source,omitempty"`
	// Streaming DataSource from where this view can consume "online" feature data.
	StreamSource *DataSource `protobuf:"bytes,9,opt,name=stream_source,json=streamSource,proto3" json:"stream_source,omitempty"`
	// Whether these features should be served online or not
	Online bool `protobuf:"varint,8,opt,name=online,proto3" json:"online,omitempty"`
}

func (x *FeatureViewSpec) Reset() {
	*x = FeatureViewSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_FeatureView_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureViewSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureViewSpec) ProtoMessage() {}

func (x *FeatureViewSpec) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_FeatureView_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureViewSpec.ProtoReflect.Descriptor instead.
func (*FeatureViewSpec) Descriptor() ([]byte, []int) {
	return file_feast_core_FeatureView_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureViewSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureViewSpec) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *FeatureViewSpec) GetEntities() []string {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *FeatureViewSpec) GetFeatures() []*FeatureSpecV2 {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *FeatureViewSpec) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *FeatureViewSpec) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *FeatureViewSpec) GetBatchSource() *DataSource {
	if x != nil {
		return x.BatchSource
	}
	return nil
}

func (x *FeatureViewSpec) GetStreamSource() *DataSource {
	if x != nil {
		return x.StreamSource
	}
	return nil
}

func (x *FeatureViewSpec) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

type FeatureViewMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time where this Feature View is created
	CreatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp,omitempty"`
	// Time where this Feature View is last updated
	LastUpdatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_updated_timestamp,json=lastUpdatedTimestamp,proto3" json:"last_updated_timestamp,omitempty"`
	// List of pairs (start_time, end_time) for which this feature view has been materialized.
	MaterializationIntervals []*MaterializationInterval `protobuf:"bytes,3,rep,name=materialization_intervals,json=materializationIntervals,proto3" json:"materialization_intervals,omitempty"`
}

func (x *FeatureViewMeta) Reset() {
	*x = FeatureViewMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_FeatureView_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureViewMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureViewMeta) ProtoMessage() {}

func (x *FeatureViewMeta) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_FeatureView_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureViewMeta.ProtoReflect.Descriptor instead.
func (*FeatureViewMeta) Descriptor() ([]byte, []int) {
	return file_feast_core_FeatureView_proto_rawDescGZIP(), []int{2}
}

func (x *FeatureViewMeta) GetCreatedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTimestamp
	}
	return nil
}

func (x *FeatureViewMeta) GetLastUpdatedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedTimestamp
	}
	return nil
}

func (x *FeatureViewMeta) GetMaterializationIntervals() []*MaterializationInterval {
	if x != nil {
		return x.MaterializationIntervals
	}
	return nil
}

type MaterializationInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *MaterializationInterval) Reset() {
	*x = MaterializationInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_FeatureView_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterializationInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterializationInterval) ProtoMessage() {}

func (x *MaterializationInterval) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_FeatureView_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterializationInterval.ProtoReflect.Descriptor instead.
func (*MaterializationInterval) Descriptor() ([]byte, []int) {
	return file_feast_core_FeatureView_proto_rawDescGZIP(), []int{3}
}

func (x *MaterializationInterval) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *MaterializationInterval) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_feast_core_FeatureView_proto protoreflect.FileDescriptor

var file_feast_core_FeatureView_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x66, 0x65, 0x61,
	0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0b, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x22, 0xc3, 0x03, 0x0a, 0x0f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x70, 0x65, 0x63, 0x56, 0x32, 0x52, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x12, 0x39, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e, 0x02, 0x0a, 0x0f, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x47, 0x0a,
	0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x50, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x60, 0x0a, 0x19, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x52, 0x18, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x55, 0x0a, 0x10, 0x66, 0x65, 0x61, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x10, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feast_core_FeatureView_proto_rawDescOnce sync.Once
	file_feast_core_FeatureView_proto_rawDescData = file_feast_core_FeatureView_proto_rawDesc
)

func file_feast_core_FeatureView_proto_rawDescGZIP() []byte {
	file_feast_core_FeatureView_proto_rawDescOnce.Do(func() {
		file_feast_core_FeatureView_proto_rawDescData = protoimpl.X.CompressGZIP(file_feast_core_FeatureView_proto_rawDescData)
	})
	return file_feast_core_FeatureView_proto_rawDescData
}

var file_feast_core_FeatureView_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_feast_core_FeatureView_proto_goTypes = []interface{}{
	(*FeatureView)(nil),             // 0: feast.core.FeatureView
	(*FeatureViewSpec)(nil),         // 1: feast.core.FeatureViewSpec
	(*FeatureViewMeta)(nil),         // 2: feast.core.FeatureViewMeta
	(*MaterializationInterval)(nil), // 3: feast.core.MaterializationInterval
	nil,                             // 4: feast.core.FeatureViewSpec.TagsEntry
	(*FeatureSpecV2)(nil),           // 5: feast.core.FeatureSpecV2
	(*durationpb.Duration)(nil),     // 6: google.protobuf.Duration
	(*DataSource)(nil),              // 7: feast.core.DataSource
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_feast_core_FeatureView_proto_depIdxs = []int32{
	1,  // 0: feast.core.FeatureView.spec:type_name -> feast.core.FeatureViewSpec
	2,  // 1: feast.core.FeatureView.meta:type_name -> feast.core.FeatureViewMeta
	5,  // 2: feast.core.FeatureViewSpec.features:type_name -> feast.core.FeatureSpecV2
	4,  // 3: feast.core.FeatureViewSpec.tags:type_name -> feast.core.FeatureViewSpec.TagsEntry
	6,  // 4: feast.core.FeatureViewSpec.ttl:type_name -> google.protobuf.Duration
	7,  // 5: feast.core.FeatureViewSpec.batch_source:type_name -> feast.core.DataSource
	7,  // 6: feast.core.FeatureViewSpec.stream_source:type_name -> feast.core.DataSource
	8,  // 7: feast.core.FeatureViewMeta.created_timestamp:type_name -> google.protobuf.Timestamp
	8,  // 8: feast.core.FeatureViewMeta.last_updated_timestamp:type_name -> google.protobuf.Timestamp
	3,  // 9: feast.core.FeatureViewMeta.materialization_intervals:type_name -> feast.core.MaterializationInterval
	8,  // 10: feast.core.MaterializationInterval.start_time:type_name -> google.protobuf.Timestamp
	8,  // 11: feast.core.MaterializationInterval.end_time:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_feast_core_FeatureView_proto_init() }
func file_feast_core_FeatureView_proto_init() {
	if File_feast_core_FeatureView_proto != nil {
		return
	}
	file_feast_core_DataSource_proto_init()
	file_feast_core_Feature_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_feast_core_FeatureView_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureView); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_FeatureView_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureViewSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_FeatureView_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureViewMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_FeatureView_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaterializationInterval); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feast_core_FeatureView_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feast_core_FeatureView_proto_goTypes,
		DependencyIndexes: file_feast_core_FeatureView_proto_depIdxs,
		MessageInfos:      file_feast_core_FeatureView_proto_msgTypes,
	}.Build()
	File_feast_core_FeatureView_proto = out.File
	file_feast_core_FeatureView_proto_rawDesc = nil
	file_feast_core_FeatureView_proto_goTypes = nil
	file_feast_core_FeatureView_proto_depIdxs = nil
}
