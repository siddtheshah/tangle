// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/profile.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Profile_Gender int32

const (
	Profile_UNKNOWN Profile_Gender = 0 // Should not exist
	Profile_FEMALE  Profile_Gender = 1
	Profile_MALE    Profile_Gender = 2 // Tangle doesn't support non-bipartite (and thus non-CIS) matching, currently.
)

// Enum value maps for Profile_Gender.
var (
	Profile_Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "FEMALE",
		2: "MALE",
	}
	Profile_Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"FEMALE":  1,
		"MALE":    2,
	}
)

func (x Profile_Gender) Enum() *Profile_Gender {
	p := new(Profile_Gender)
	*p = x
	return p
}

func (x Profile_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Profile_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_profile_proto_enumTypes[0].Descriptor()
}

func (Profile_Gender) Type() protoreflect.EnumType {
	return &file_proto_profile_proto_enumTypes[0]
}

func (x Profile_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Profile_Gender.Descriptor instead.
func (Profile_Gender) EnumDescriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{0, 0}
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A pair of IDs generated via a UUID function. We should ensure a mechanism to avoid purposeful collisions.
	PrimaryId uint64 `protobuf:"varint,1,opt,name=primary_id,json=primaryId,proto3" json:"primary_id,omitempty"`
	// These IDs can be concatenated to get a combined ID, which should have sufficient sparsity to avoid random collisions.
	SecondaryId uint64 `protobuf:"varint,5,opt,name=secondary_id,json=secondaryId,proto3" json:"secondary_id,omitempty"`
	// Subject's name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Subject's bio.
	Bio string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	// Profile image CID on IPFS.
	ImageCid string `protobuf:"bytes,4,opt,name=image_cid,json=imageCid,proto3" json:"image_cid,omitempty"`
	// Subject's gender.
	Gender Profile_Gender `protobuf:"varint,6,opt,name=gender,proto3,enum=Profile_Gender" json:"gender,omitempty"`
	// Subject's occupation.
	Occupation string `protobuf:"bytes,7,opt,name=occupation,proto3" json:"occupation,omitempty"`
	// Subject's general place of residence.
	Residence string `protobuf:"bytes,8,opt,name=residence,proto3" json:"residence,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_proto_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetPrimaryId() uint64 {
	if x != nil {
		return x.PrimaryId
	}
	return 0
}

func (x *Profile) GetSecondaryId() uint64 {
	if x != nil {
		return x.SecondaryId
	}
	return 0
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *Profile) GetImageCid() string {
	if x != nil {
		return x.ImageCid
	}
	return ""
}

func (x *Profile) GetGender() Profile_Gender {
	if x != nil {
		return x.Gender
	}
	return Profile_UNKNOWN
}

func (x *Profile) GetOccupation() string {
	if x != nil {
		return x.Occupation
	}
	return ""
}

func (x *Profile) GetResidence() string {
	if x != nil {
		return x.Residence
	}
	return ""
}

var File_proto_profile_proto protoreflect.FileDescriptor

var file_proto_profile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x43, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2b, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_profile_proto_rawDescOnce sync.Once
	file_proto_profile_proto_rawDescData = file_proto_profile_proto_rawDesc
)

func file_proto_profile_proto_rawDescGZIP() []byte {
	file_proto_profile_proto_rawDescOnce.Do(func() {
		file_proto_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_profile_proto_rawDescData)
	})
	return file_proto_profile_proto_rawDescData
}

var file_proto_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_profile_proto_goTypes = []interface{}{
	(Profile_Gender)(0), // 0: Profile.Gender
	(*Profile)(nil),     // 1: Profile
}
var file_proto_profile_proto_depIdxs = []int32{
	0, // 0: Profile.gender:type_name -> Profile.Gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_profile_proto_init() }
func file_proto_profile_proto_init() {
	if File_proto_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
			RawDescriptor: file_proto_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_profile_proto_goTypes,
		DependencyIndexes: file_proto_profile_proto_depIdxs,
		EnumInfos:         file_proto_profile_proto_enumTypes,
		MessageInfos:      file_proto_profile_proto_msgTypes,
	}.Build()
	File_proto_profile_proto = out.File
	file_proto_profile_proto_rawDesc = nil
	file_proto_profile_proto_goTypes = nil
	file_proto_profile_proto_depIdxs = nil
}
