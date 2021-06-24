// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: google/maps/routes/v1/route_service.proto

package routes

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_google_maps_routes_v1_route_service_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_route_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x98, 0x03, 0x0a, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x88, 0x01, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x76, 0x31, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x30, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x22, 0x16, 0x2f, 0x76, 0x31, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x1a, 0x61, 0xca,
	0x41, 0x1e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0xd2, 0x41, 0x3d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x42, 0xa8, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x4d, 0x52, 0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70,
	0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_google_maps_routes_v1_route_service_proto_goTypes = []interface{}{
	(*ComputeRoutesRequest)(nil),      // 0: google.maps.routes.v1.ComputeRoutesRequest
	(*ComputeRouteMatrixRequest)(nil), // 1: google.maps.routes.v1.ComputeRouteMatrixRequest
	(*ComputeRoutesResponse)(nil),     // 2: google.maps.routes.v1.ComputeRoutesResponse
	(*RouteMatrixElement)(nil),        // 3: google.maps.routes.v1.RouteMatrixElement
}
var file_google_maps_routes_v1_route_service_proto_depIdxs = []int32{
	0, // 0: google.maps.routes.v1.RoutesPreferred.ComputeRoutes:input_type -> google.maps.routes.v1.ComputeRoutesRequest
	1, // 1: google.maps.routes.v1.RoutesPreferred.ComputeRouteMatrix:input_type -> google.maps.routes.v1.ComputeRouteMatrixRequest
	2, // 2: google.maps.routes.v1.RoutesPreferred.ComputeRoutes:output_type -> google.maps.routes.v1.ComputeRoutesResponse
	3, // 3: google.maps.routes.v1.RoutesPreferred.ComputeRouteMatrix:output_type -> google.maps.routes.v1.RouteMatrixElement
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_route_service_proto_init() }
func file_google_maps_routes_v1_route_service_proto_init() {
	if File_google_maps_routes_v1_route_service_proto != nil {
		return
	}
	file_google_maps_routes_v1_compute_custom_routes_request_proto_init()
	file_google_maps_routes_v1_compute_custom_routes_response_proto_init()
	file_google_maps_routes_v1_compute_route_matrix_request_proto_init()
	file_google_maps_routes_v1_compute_routes_request_proto_init()
	file_google_maps_routes_v1_compute_routes_response_proto_init()
	file_google_maps_routes_v1_route_matrix_element_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routes_v1_route_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_maps_routes_v1_route_service_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_route_service_proto_depIdxs,
	}.Build()
	File_google_maps_routes_v1_route_service_proto = out.File
	file_google_maps_routes_v1_route_service_proto_rawDesc = nil
	file_google_maps_routes_v1_route_service_proto_goTypes = nil
	file_google_maps_routes_v1_route_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoutesPreferredClient is the client API for RoutesPreferred service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutesPreferredClient interface {
	// Returns the primary route along with optional alternate routes, given a set
	// of terminal and intermediate waypoints.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using URL parameter
	// `$fields` or `fields`, or by using an HTTP/gRPC header `X-Goog-FieldMask`
	// (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See detailed documentation about
	// [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of Route-level duration, distance, and polyline (an example
	// production setup):
	//   `X-Goog-FieldMask:
	//   routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline`
	//
	// Google discourage the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need
	// in your production job ensures stable latency performance. We might add
	// more response fields in the future, and those new fields might require
	// extra computation time. If you select all fields, or if you select all
	// fields at the top level, then you might experience performance degradation
	// because any new field we add will be automatically included in the
	// response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRoutes(ctx context.Context, in *ComputeRoutesRequest, opts ...grpc.CallOption) (*ComputeRoutesResponse, error)
	// Takes in a list of origins and destinations and returns a stream containing
	// route information for each combination of origin and destination.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of route durations, distances, element status, condition, and
	//   element indices (an example production setup):
	//   `X-Goog-FieldMask:
	//   originIndex,destinationIndex,status,condition,distanceMeters,duration`
	//
	// It is critical that you include `status` in your field mask as otherwise
	// all messages will appear to be OK. Google discourages the use of the
	// wildcard (`*`) response field mask, because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRouteMatrix(ctx context.Context, in *ComputeRouteMatrixRequest, opts ...grpc.CallOption) (RoutesPreferred_ComputeRouteMatrixClient, error)
}

type routesPreferredClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutesPreferredClient(cc grpc.ClientConnInterface) RoutesPreferredClient {
	return &routesPreferredClient{cc}
}

func (c *routesPreferredClient) ComputeRoutes(ctx context.Context, in *ComputeRoutesRequest, opts ...grpc.CallOption) (*ComputeRoutesResponse, error) {
	out := new(ComputeRoutesResponse)
	err := c.cc.Invoke(ctx, "/google.maps.routes.v1.RoutesPreferred/ComputeRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routesPreferredClient) ComputeRouteMatrix(ctx context.Context, in *ComputeRouteMatrixRequest, opts ...grpc.CallOption) (RoutesPreferred_ComputeRouteMatrixClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoutesPreferred_serviceDesc.Streams[0], "/google.maps.routes.v1.RoutesPreferred/ComputeRouteMatrix", opts...)
	if err != nil {
		return nil, err
	}
	x := &routesPreferredComputeRouteMatrixClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutesPreferred_ComputeRouteMatrixClient interface {
	Recv() (*RouteMatrixElement, error)
	grpc.ClientStream
}

type routesPreferredComputeRouteMatrixClient struct {
	grpc.ClientStream
}

func (x *routesPreferredComputeRouteMatrixClient) Recv() (*RouteMatrixElement, error) {
	m := new(RouteMatrixElement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoutesPreferredServer is the server API for RoutesPreferred service.
type RoutesPreferredServer interface {
	// Returns the primary route along with optional alternate routes, given a set
	// of terminal and intermediate waypoints.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using URL parameter
	// `$fields` or `fields`, or by using an HTTP/gRPC header `X-Goog-FieldMask`
	// (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See detailed documentation about
	// [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of Route-level duration, distance, and polyline (an example
	// production setup):
	//   `X-Goog-FieldMask:
	//   routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline`
	//
	// Google discourage the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need
	// in your production job ensures stable latency performance. We might add
	// more response fields in the future, and those new fields might require
	// extra computation time. If you select all fields, or if you select all
	// fields at the top level, then you might experience performance degradation
	// because any new field we add will be automatically included in the
	// response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRoutes(context.Context, *ComputeRoutesRequest) (*ComputeRoutesResponse, error)
	// Takes in a list of origins and destinations and returns a stream containing
	// route information for each combination of origin and destination.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of route durations, distances, element status, condition, and
	//   element indices (an example production setup):
	//   `X-Goog-FieldMask:
	//   originIndex,destinationIndex,status,condition,distanceMeters,duration`
	//
	// It is critical that you include `status` in your field mask as otherwise
	// all messages will appear to be OK. Google discourages the use of the
	// wildcard (`*`) response field mask, because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRouteMatrix(*ComputeRouteMatrixRequest, RoutesPreferred_ComputeRouteMatrixServer) error
}

// UnimplementedRoutesPreferredServer can be embedded to have forward compatible implementations.
type UnimplementedRoutesPreferredServer struct {
}

func (*UnimplementedRoutesPreferredServer) ComputeRoutes(context.Context, *ComputeRoutesRequest) (*ComputeRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeRoutes not implemented")
}
func (*UnimplementedRoutesPreferredServer) ComputeRouteMatrix(*ComputeRouteMatrixRequest, RoutesPreferred_ComputeRouteMatrixServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeRouteMatrix not implemented")
}

func RegisterRoutesPreferredServer(s *grpc.Server, srv RoutesPreferredServer) {
	s.RegisterService(&_RoutesPreferred_serviceDesc, srv)
}

func _RoutesPreferred_ComputeRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesPreferredServer).ComputeRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.routes.v1.RoutesPreferred/ComputeRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesPreferredServer).ComputeRoutes(ctx, req.(*ComputeRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutesPreferred_ComputeRouteMatrix_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ComputeRouteMatrixRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutesPreferredServer).ComputeRouteMatrix(m, &routesPreferredComputeRouteMatrixServer{stream})
}

type RoutesPreferred_ComputeRouteMatrixServer interface {
	Send(*RouteMatrixElement) error
	grpc.ServerStream
}

type routesPreferredComputeRouteMatrixServer struct {
	grpc.ServerStream
}

func (x *routesPreferredComputeRouteMatrixServer) Send(m *RouteMatrixElement) error {
	return x.ServerStream.SendMsg(m)
}

var _RoutesPreferred_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.routes.v1.RoutesPreferred",
	HandlerType: (*RoutesPreferredServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeRoutes",
			Handler:    _RoutesPreferred_ComputeRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComputeRouteMatrix",
			Handler:       _RoutesPreferred_ComputeRouteMatrix_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/maps/routes/v1/route_service.proto",
}
