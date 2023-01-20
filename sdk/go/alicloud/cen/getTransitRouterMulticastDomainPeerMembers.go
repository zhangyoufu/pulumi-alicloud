// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides Cen Transit Router Multicast Domain Peer Member available to the user.[What is Transit Router Multicast Domain Peer Member](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-registertransitroutermulticastgroupmembers)
//
// > **NOTE:** Available in 1.195.0+
//
// ## Example Usage
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
//			_default, err := cen.GetTransitRouterMulticastDomainPeerMembers(ctx, &cen.GetTransitRouterMulticastDomainPeerMembersArgs{
//				TransitRouterMulticastDomainId: "tr-mcast-domain-2d9oq455uk533zfrxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("alicloudCenTransitRouterMulticastDomainPeerMemberExampleId", _default.Members[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetTransitRouterMulticastDomainPeerMembers(ctx *pulumi.Context, args *GetTransitRouterMulticastDomainPeerMembersArgs, opts ...pulumi.InvokeOption) (*GetTransitRouterMulticastDomainPeerMembersResult, error) {
	var rv GetTransitRouterMulticastDomainPeerMembersResult
	err := ctx.Invoke("alicloud:cen/getTransitRouterMulticastDomainPeerMembers:getTransitRouterMulticastDomainPeerMembers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTransitRouterMulticastDomainPeerMembers.
type GetTransitRouterMulticastDomainPeerMembersArgs struct {
	// A list of Cen Transit Router Multicast Domain Peer Member IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomains []string `pulumi:"peerTransitRouterMulticastDomains"`
	// The ID of the resource associated with the multicast resource.
	ResourceId *string `pulumi:"resourceId"`
	// The type of the multicast resource. Valid values:
	// * VPC: queries multicast resources by VPC.
	// * TR: queries multicast resources that are also deployed in a different region.
	ResourceType *string `pulumi:"resourceType"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId string `pulumi:"transitRouterMulticastDomainId"`
}

// A collection of values returned by getTransitRouterMulticastDomainPeerMembers.
type GetTransitRouterMulticastDomainPeerMembersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// A list of Transit Router Multicast Domain Peer Member Entries. Each element contains the following attributes:
	Members                           []GetTransitRouterMulticastDomainPeerMembersMember `pulumi:"members"`
	OutputFile                        *string                                            `pulumi:"outputFile"`
	PeerTransitRouterMulticastDomains []string                                           `pulumi:"peerTransitRouterMulticastDomains"`
	ResourceId                        *string                                            `pulumi:"resourceId"`
	ResourceType                      *string                                            `pulumi:"resourceType"`
	TransitRouterAttachmentId         *string                                            `pulumi:"transitRouterAttachmentId"`
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId string `pulumi:"transitRouterMulticastDomainId"`
}

func GetTransitRouterMulticastDomainPeerMembersOutput(ctx *pulumi.Context, args GetTransitRouterMulticastDomainPeerMembersOutputArgs, opts ...pulumi.InvokeOption) GetTransitRouterMulticastDomainPeerMembersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTransitRouterMulticastDomainPeerMembersResult, error) {
			args := v.(GetTransitRouterMulticastDomainPeerMembersArgs)
			r, err := GetTransitRouterMulticastDomainPeerMembers(ctx, &args, opts...)
			var s GetTransitRouterMulticastDomainPeerMembersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTransitRouterMulticastDomainPeerMembersResultOutput)
}

// A collection of arguments for invoking getTransitRouterMulticastDomainPeerMembers.
type GetTransitRouterMulticastDomainPeerMembersOutputArgs struct {
	// A list of Cen Transit Router Multicast Domain Peer Member IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The IDs of the inter-region multicast domains.
	PeerTransitRouterMulticastDomains pulumi.StringArrayInput `pulumi:"peerTransitRouterMulticastDomains"`
	// The ID of the resource associated with the multicast resource.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// The type of the multicast resource. Valid values:
	// * VPC: queries multicast resources by VPC.
	// * TR: queries multicast resources that are also deployed in a different region.
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId pulumi.StringPtrInput `pulumi:"transitRouterAttachmentId"`
	// The ID of the multicast domain to which the multicast member belongs.
	TransitRouterMulticastDomainId pulumi.StringInput `pulumi:"transitRouterMulticastDomainId"`
}

func (GetTransitRouterMulticastDomainPeerMembersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterMulticastDomainPeerMembersArgs)(nil)).Elem()
}

// A collection of values returned by getTransitRouterMulticastDomainPeerMembers.
type GetTransitRouterMulticastDomainPeerMembersResultOutput struct{ *pulumi.OutputState }

func (GetTransitRouterMulticastDomainPeerMembersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTransitRouterMulticastDomainPeerMembersResult)(nil)).Elem()
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) ToGetTransitRouterMulticastDomainPeerMembersResultOutput() GetTransitRouterMulticastDomainPeerMembersResultOutput {
	return o
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) ToGetTransitRouterMulticastDomainPeerMembersResultOutputWithContext(ctx context.Context) GetTransitRouterMulticastDomainPeerMembersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of Transit Router Multicast Domain Peer Member Entries. Each element contains the following attributes:
func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) Members() GetTransitRouterMulticastDomainPeerMembersMemberArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) []GetTransitRouterMulticastDomainPeerMembersMember {
		return v.Members
	}).(GetTransitRouterMulticastDomainPeerMembersMemberArrayOutput)
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) PeerTransitRouterMulticastDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) []string {
		return v.PeerTransitRouterMulticastDomains
	}).(pulumi.StringArrayOutput)
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) TransitRouterAttachmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) *string { return v.TransitRouterAttachmentId }).(pulumi.StringPtrOutput)
}

// The ID of the multicast domain to which the multicast member belongs.
func (o GetTransitRouterMulticastDomainPeerMembersResultOutput) TransitRouterMulticastDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTransitRouterMulticastDomainPeerMembersResult) string {
		return v.TransitRouterMulticastDomainId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTransitRouterMulticastDomainPeerMembersResultOutput{})
}
