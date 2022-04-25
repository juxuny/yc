syntax = "proto3";

package {{ .PackageName }};

option go_package = "./;{{ .PackageName }}";

service {{ .ServiceStruct }} {
  // @group: devops
  rpc Health(HealthRequest) returns (HealthResponse) {}
}

message HealthRequest {}

message HealthResponse {
  string currentTime = 1;
}
