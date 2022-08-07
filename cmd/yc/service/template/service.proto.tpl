syntax = "proto3";

package {{ .PackageName }};

option go_package = "./;{{ .PackageName }}";

// @level: 100
// @version: 1.0.0
service {{ .ServiceStruct }} {
  // @group: devops
  // @ignore-auth
  // @desc: health check
  rpc Health(HealthRequest) returns (HealthResponse) {}
}

message HealthRequest {}

message HealthResponse {
  string currentTime = 1;
}
