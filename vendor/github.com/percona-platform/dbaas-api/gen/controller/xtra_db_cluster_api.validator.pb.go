// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/xtra_db_cluster_api.proto

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

func (this *ListXtraDBClustersRequest) Validate() error {
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
func (this *ListXtraDBClustersResponse) Validate() error {
	for _, item := range this.Clusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Clusters", err)
			}
		}
	}
	return nil
}
func (this *ListXtraDBClustersResponse_Cluster) Validate() error {
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
func (this *GetXtraDBClusterCredentialsRequest) Validate() error {
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
func (this *GetXtraDBClusterCredentialsResponse) Validate() error {
	if this.Credentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credentials", err)
		}
	}
	return nil
}
func (this *CreateXtraDBClusterRequest) Validate() error {
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
func (this *CreateXtraDBClusterResponse) Validate() error {
	return nil
}
func (this *UpdateXtraDBClusterRequest) Validate() error {
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
func (this *UpdateXtraDBClusterRequest_UpdateXtraDBClusterParams) Validate() error {
	if this.Pxc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pxc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pxc", err)
		}
	}
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	return nil
}
func (this *UpdateXtraDBClusterRequest_UpdateXtraDBClusterParams_PXC) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdateXtraDBClusterRequest_UpdateXtraDBClusterParams_ProxySQL) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdateXtraDBClusterResponse) Validate() error {
	return nil
}
func (this *DeleteXtraDBClusterRequest) Validate() error {
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
func (this *DeleteXtraDBClusterResponse) Validate() error {
	return nil
}
func (this *RestartXtraDBClusterRequest) Validate() error {
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
func (this *RestartXtraDBClusterResponse) Validate() error {
	return nil
}
