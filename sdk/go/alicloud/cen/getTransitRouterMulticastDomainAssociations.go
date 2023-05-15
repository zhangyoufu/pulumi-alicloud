// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cen Transit Router Multicast Domain Associations of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.195.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := cen.GetTransitRouterMulticastDomainAssociations(ctx, &cen.GetTransitRouterMulticastDomainAssociationsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				TransitRouterMulticastDomainId: "your_transit_router_multicast_domain_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("cenTransitRouterMulticastDomainId0", ids.Associations[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetTransitRouterMulticastDomainAssociations(ctx *pulumi.Context, args *GetTransitRouterMulticastDomainAssociationsArgs, opts ...pulumi.InvokeOption) (*GetTransitRouterMulticastDomainAssociationsResult, error) {
	var rv GetTransitRouterMulticastDomainAssociationsResult
	err := ctx.Invoke("alicloud:cen/getTransitRouterMulticastDomainAssociations:getTransitRouterMulticastDomainAssociations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransitRouterMulticastDomainAssociations.
type GetTransitRouterMulticastDomainAssociationsArgs struct {
	// A list of Transit Router Multicast Domain Association IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the resource associated with the multicast domain.
	ResourceId *string `pulumi:"resourceId"`
	// The type of resource associated with the multicast domain. Valid Value: `VPC`.
	ResourceType *string `pulumi:"resourceType"`
	// The status of the associated resource. Valid Value: `Associated`, `Associating`, `Dissociating`.
	Status *string `pulumi:"status"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the multicast domain.
	TransitRouterMulticastDomainId string `pulumi:"transitRouterMulticastDomainId"`
	// The ID of the vSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

// A collection of values returned by getTransitRouterMulticastDomainAssociations.
type GetTransitRouterMulticastDomainAssociationsResult struct {
	// A list of Cen Transit Router Multicast Domain Associations. Each element contains the following attributes:
	Associations []GetTransitRouterMulticastDomainAssociationsAssociation `pulumi:"associations"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the resource associated with the multicast domain.
	ResourceId *string `pulumi:"resourceId"`
	// The type of resource associated with the multicast domain.
	ResourceType *string `pulumi:"resourceType"`
	// The status of the associated resource.
	Status *string `pulumi:"status"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the multicast domain.
	TransitRouterMulticastDomainId string `pulumi:"transitRouterMulticastDomainId"`
	// The ID of the vSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

func GetTransitRouterMulticastDomainAssociationsOutput(ctx *pulumi.Context, args GetTransitRouterMulticastDomainAssociationsOutputArgs, opts ...pulumi.InvokeOption) GetTransitRouterMulticastDomainAssociationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTransitRouterMulticastDomainAssociationsResult, error) {
			args := v.(GetTransitRouterMulticastDomainAssociationsArgs)
			r, err := GetTransitRouterMulticastDomainAssociations(ctx, &args, opts...)
			var s GetTransitRouterMulticastDomainAssociationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTransitRouterMulticastDomainAssociationsResultOutput)
}

// A collection of arguments for invoking getTransitRouterMulticastDomainAssociations.
type GetTransitRouterMulticastDomainAssociationsOutputArgs struct {
	// A list of Transit Router Multicast Domain Association IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the resource associated with the multicast domain.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// The type of resource associated with the multicast domain. Valid Value: `VPC`.
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	// The status of the associated resource. Valid Value: `Associated`, `Associating`, `Dissociating`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId pulumi.StringPtrInput `pulumi:"transitRouterAttachmentId"`
	// The ID of the multicast domain.
	TransitRouterMulticastDomainId pulumi.StringInput `pulumi:"transitRouterMulticastDomainId"`
	// The ID of the vSwitch.
	VswitchId pulumi.StringPtrInput `pulumi:"vswitchId"`
}

func (GetTransitRouterMulticastDomainAssociationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterMulticastDomainAssociationsArgs)(nil)).Elem()
}

// A collection of values returned by getTransitRouterMulticastDomainAssociations.
type GetTransitRouterMulticastDomainAssociationsResultOutput struct{ *pulumi.OutputState }

func (GetTransitRouterMulticastDomainAssociationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterMulticastDomainAssociationsResult)(nil)).Elem()
}

func (o GetTransitRouterMulticastDomainAssociationsResultOutput) ToGetTransitRouterMulticastDomainAssociationsResultOutput() GetTransitRouterMulticastDomainAssociationsResultOutput {
	return o
}

func (o GetTransitRouterMulticastDomainAssociationsResultOutput) ToGetTransitRouterMulticastDomainAssociationsResultOutputWithContext(ctx context.Context) GetTransitRouterMulticastDomainAssociationsResultOutput {
	return o
}

// A list of Cen Transit Router Multicast Domain Associations. Each element contains the following attributes:
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) Associations() GetTransitRouterMulticastDomainAssociationsAssociationArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) []GetTransitRouterMulticastDomainAssociationsAssociation {
		return v.Associations
	}).(GetTransitRouterMulticastDomainAssociationsAssociationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTransitRouterMulticastDomainAssociationsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterMulticastDomainAssociationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ID of the resource associated with the multicast domain.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of resource associated with the multicast domain.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// The status of the associated resource.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The ID of the network instance connection.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) TransitRouterAttachmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) *string { return v.TransitRouterAttachmentId }).(pulumi.StringPtrOutput)
}

// The ID of the multicast domain.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) TransitRouterMulticastDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) string {
		return v.TransitRouterMulticastDomainId
	}).(pulumi.StringOutput)
}

// The ID of the vSwitch.
func (o GetTransitRouterMulticastDomainAssociationsResultOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainAssociationsResult) *string { return v.VswitchId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTransitRouterMulticastDomainAssociationsResultOutput{})
}
