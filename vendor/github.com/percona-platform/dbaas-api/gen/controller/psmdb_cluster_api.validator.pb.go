// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/psmdb_cluster_api.proto

package controllerv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListPSMDBClustersRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	return nil
}
func (this *ListPSMDBClustersResponse) Validate() error {
	for _, item := range this.Clusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Clusters", err)
			}
		}
	}
	return nil
}
func (this *ListPSMDBClustersResponse_Cluster) Validate() error {
	if this.Operation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operation", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *GetPSMDBClusterCredentialsRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *GetPSMDBClusterCredentialsResponse) Validate() error {
	if this.Credentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credentials", err)
		}
	}
	return nil
}
func (this *CreatePSMDBClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *CreatePSMDBClusterResponse) Validate() error {
	return nil
}
func (this *UpdatePSMDBClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *UpdatePSMDBClusterRequest_UpdatePSMDBClusterParams) Validate() error {
	if this.Replicaset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Replicaset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", err)
		}
	}
	return nil
}
func (this *UpdatePSMDBClusterRequest_UpdatePSMDBClusterParams_ReplicaSet) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdatePSMDBClusterResponse) Validate() error {
	return nil
}
func (this *DeletePSMDBClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *DeletePSMDBClusterResponse) Validate() error {
	return nil
}
func (this *RestartPSMDBClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *RestartPSMDBClusterResponse) Validate() error {
	return nil
}
