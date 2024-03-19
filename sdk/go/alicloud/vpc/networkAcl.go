// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Network Acl resource.
//
// For information about VPC Network Acl and how to use it, see [What is Network Acl](https://www.alibabacloud.com/help/en/ens/latest/createnetworkacl).
//
// > **NOTE:** Available since v1.43.0.
//
// ## Example Usage
//
// # Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_default, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef("VSwitch"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleNetwork, err := vpc.NewNetwork(ctx, "exampleNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSwitch, err := vpc.NewSwitch(ctx, "exampleSwitch", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       exampleNetwork.ID(),
//				ZoneId:      pulumi.String(_default.Zones[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewNetworkAcl(ctx, "exampleNetworkAcl", &vpc.NetworkAclArgs{
//				VpcId:          exampleNetwork.ID(),
//				NetworkAclName: pulumi.String(name),
//				Description:    pulumi.String(name),
//				IngressAclEntries: vpc.NetworkAclIngressAclEntryArray{
//					&vpc.NetworkAclIngressAclEntryArgs{
//						Description:         pulumi.String(fmt.Sprintf("%v-ingress", name)),
//						NetworkAclEntryName: pulumi.String(fmt.Sprintf("%v-ingress", name)),
//						SourceCidrIp:        pulumi.String("10.0.0.0/24"),
//						Policy:              pulumi.String("accept"),
//						Port:                pulumi.String("20/80"),
//						Protocol:            pulumi.String("tcp"),
//					},
//				},
//				EgressAclEntries: vpc.NetworkAclEgressAclEntryArray{
//					&vpc.NetworkAclEgressAclEntryArgs{
//						Description:         pulumi.String(fmt.Sprintf("%v-egress", name)),
//						NetworkAclEntryName: pulumi.String(fmt.Sprintf("%v-egress", name)),
//						DestinationCidrIp:   pulumi.String("10.0.0.0/24"),
//						Policy:              pulumi.String("accept"),
//						Port:                pulumi.String("20/80"),
//						Protocol:            pulumi.String("tcp"),
//					},
//				},
//				Resources: vpc.NetworkAclResourceArray{
//					&vpc.NetworkAclResourceArgs{
//						ResourceId:   exampleSwitch.ID(),
//						ResourceType: pulumi.String("VSwitch"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// VPC Network Acl can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/networkAcl:NetworkAcl example <id>
// ```
type NetworkAcl struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Out direction rule information. See `egressAclEntries` below.
	EgressAclEntries NetworkAclEgressAclEntryArrayOutput `pulumi:"egressAclEntries"`
	// Inward direction rule information. See `ingressAclEntries` below.
	IngressAclEntries NetworkAclIngressAclEntryArrayOutput `pulumi:"ingressAclEntries"`
	// Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
	NetworkAclName pulumi.StringOutput `pulumi:"networkAclName"`
	// The associated resource. See `resources` below.
	Resources NetworkAclResourceArrayOutput `pulumi:"resources"`
	// The status of the associated resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of this resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the associated VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewNetworkAcl(ctx *pulumi.Context,
	name string, args *NetworkAclArgs, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkAcl
	err := ctx.RegisterResource("alicloud:vpc/networkAcl:NetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAcl gets an existing NetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclState, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	var resource NetworkAcl
	err := ctx.ReadResource("alicloud:vpc/networkAcl:NetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAcl resources.
type networkAclState struct {
	// The creation time of the resource.
	CreateTime *string `pulumi:"createTime"`
	// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
	Description *string `pulumi:"description"`
	// Out direction rule information. See `egressAclEntries` below.
	EgressAclEntries []NetworkAclEgressAclEntry `pulumi:"egressAclEntries"`
	// Inward direction rule information. See `ingressAclEntries` below.
	IngressAclEntries []NetworkAclIngressAclEntry `pulumi:"ingressAclEntries"`
	// Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	Name *string `pulumi:"name"`
	// The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
	NetworkAclName *string `pulumi:"networkAclName"`
	// The associated resource. See `resources` below.
	Resources []NetworkAclResource `pulumi:"resources"`
	// The status of the associated resource.
	Status *string `pulumi:"status"`
	// The tags of this resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the associated VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId *string `pulumi:"vpcId"`
}

type NetworkAclState struct {
	// The creation time of the resource.
	CreateTime pulumi.StringPtrInput
	// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
	Description pulumi.StringPtrInput
	// Out direction rule information. See `egressAclEntries` below.
	EgressAclEntries NetworkAclEgressAclEntryArrayInput
	// Inward direction rule information. See `ingressAclEntries` below.
	IngressAclEntries NetworkAclIngressAclEntryArrayInput
	// Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	Name pulumi.StringPtrInput
	// The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
	NetworkAclName pulumi.StringPtrInput
	// The associated resource. See `resources` below.
	Resources NetworkAclResourceArrayInput
	// The status of the associated resource.
	Status pulumi.StringPtrInput
	// The tags of this resource.
	Tags pulumi.MapInput
	// The ID of the associated VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId pulumi.StringPtrInput
}

func (NetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclState)(nil)).Elem()
}

type networkAclArgs struct {
	// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
	Description *string `pulumi:"description"`
	// Out direction rule information. See `egressAclEntries` below.
	EgressAclEntries []NetworkAclEgressAclEntry `pulumi:"egressAclEntries"`
	// Inward direction rule information. See `ingressAclEntries` below.
	IngressAclEntries []NetworkAclIngressAclEntry `pulumi:"ingressAclEntries"`
	// Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	Name *string `pulumi:"name"`
	// The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
	NetworkAclName *string `pulumi:"networkAclName"`
	// The associated resource. See `resources` below.
	Resources []NetworkAclResource `pulumi:"resources"`
	// The tags of this resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the associated VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a NetworkAcl resource.
type NetworkAclArgs struct {
	// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
	Description pulumi.StringPtrInput
	// Out direction rule information. See `egressAclEntries` below.
	EgressAclEntries NetworkAclEgressAclEntryArrayInput
	// Inward direction rule information. See `ingressAclEntries` below.
	IngressAclEntries NetworkAclIngressAclEntryArrayInput
	// Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
	Name pulumi.StringPtrInput
	// The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
	NetworkAclName pulumi.StringPtrInput
	// The associated resource. See `resources` below.
	Resources NetworkAclResourceArrayInput
	// The tags of this resource.
	Tags pulumi.MapInput
	// The ID of the associated VPC.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	VpcId pulumi.StringInput
}

func (NetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclArgs)(nil)).Elem()
}

type NetworkAclInput interface {
	pulumi.Input

	ToNetworkAclOutput() NetworkAclOutput
	ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput
}

func (*NetworkAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (i *NetworkAcl) ToNetworkAclOutput() NetworkAclOutput {
	return i.ToNetworkAclOutputWithContext(context.Background())
}

func (i *NetworkAcl) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclOutput)
}

// NetworkAclArrayInput is an input type that accepts NetworkAclArray and NetworkAclArrayOutput values.
// You can construct a concrete instance of `NetworkAclArrayInput` via:
//
//	NetworkAclArray{ NetworkAclArgs{...} }
type NetworkAclArrayInput interface {
	pulumi.Input

	ToNetworkAclArrayOutput() NetworkAclArrayOutput
	ToNetworkAclArrayOutputWithContext(context.Context) NetworkAclArrayOutput
}

type NetworkAclArray []NetworkAclInput

func (NetworkAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclArray) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return i.ToNetworkAclArrayOutputWithContext(context.Background())
}

func (i NetworkAclArray) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclArrayOutput)
}

// NetworkAclMapInput is an input type that accepts NetworkAclMap and NetworkAclMapOutput values.
// You can construct a concrete instance of `NetworkAclMapInput` via:
//
//	NetworkAclMap{ "key": NetworkAclArgs{...} }
type NetworkAclMapInput interface {
	pulumi.Input

	ToNetworkAclMapOutput() NetworkAclMapOutput
	ToNetworkAclMapOutputWithContext(context.Context) NetworkAclMapOutput
}

type NetworkAclMap map[string]NetworkAclInput

func (NetworkAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclMap) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return i.ToNetworkAclMapOutputWithContext(context.Background())
}

func (i NetworkAclMap) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclMapOutput)
}

type NetworkAclOutput struct{ *pulumi.OutputState }

func (NetworkAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (o NetworkAclOutput) ToNetworkAclOutput() NetworkAclOutput {
	return o
}

func (o NetworkAclOutput) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return o
}

// The creation time of the resource.
func (o NetworkAclOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the network ACL.The description must be 1 to 256 characters in length and cannot start with http:// or https.
func (o NetworkAclOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Out direction rule information. See `egressAclEntries` below.
func (o NetworkAclOutput) EgressAclEntries() NetworkAclEgressAclEntryArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclEgressAclEntryArrayOutput { return v.EgressAclEntries }).(NetworkAclEgressAclEntryArrayOutput)
}

// Inward direction rule information. See `ingressAclEntries` below.
func (o NetworkAclOutput) IngressAclEntries() NetworkAclIngressAclEntryArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclIngressAclEntryArrayOutput { return v.IngressAclEntries }).(NetworkAclIngressAclEntryArrayOutput)
}

// Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
//
// Deprecated: Field 'name' has been deprecated from provider version 1.122.0. New field 'network_acl_name' instead.
func (o NetworkAclOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the network ACL.The name must be 1 to 128 characters in length and cannot start with http:// or https.
func (o NetworkAclOutput) NetworkAclName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.NetworkAclName }).(pulumi.StringOutput)
}

// The associated resource. See `resources` below.
func (o NetworkAclOutput) Resources() NetworkAclResourceArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclResourceArrayOutput { return v.Resources }).(NetworkAclResourceArrayOutput)
}

// The status of the associated resource.
func (o NetworkAclOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of this resource.
func (o NetworkAclOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// The ID of the associated VPC.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o NetworkAclOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type NetworkAclArrayOutput struct{ *pulumi.OutputState }

func (NetworkAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) Index(i pulumi.IntInput) NetworkAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].([]*NetworkAcl)[vs[1].(int)]
	}).(NetworkAclOutput)
}

type NetworkAclMapOutput struct{ *pulumi.OutputState }

func (NetworkAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) MapIndex(k pulumi.StringInput) NetworkAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].(map[string]*NetworkAcl)[vs[1].(string)]
	}).(NetworkAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclInput)(nil)).Elem(), &NetworkAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclArrayInput)(nil)).Elem(), NetworkAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclMapInput)(nil)).Elem(), NetworkAclMap{})
	pulumi.RegisterOutputType(NetworkAclOutput{})
	pulumi.RegisterOutputType(NetworkAclArrayOutput{})
	pulumi.RegisterOutputType(NetworkAclMapOutput{})
}
