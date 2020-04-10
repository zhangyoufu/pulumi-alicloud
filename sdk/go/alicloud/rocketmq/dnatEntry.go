// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Sag DnatEntry resource. This topic describes how to add a DNAT entry to a Smart Access Gateway (SAG) instance to enable the DNAT function. By using the DNAT function, you can forward requests received by public IP addresses to Alibaba Cloud instances according to custom mapping rules.
//
// For information about Sag DnatEntry and how to use it, see [What is Sag DnatEntry](https://www.alibabacloud.com/help/doc-detail/124312.htm).
//
// > **NOTE:** Available in 1.63.0+
//
// > **NOTE:** Only the following regions suppor. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
type DnatEntry struct {
	pulumi.CustomResourceState

	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp pulumi.StringPtrOutput `pulumi:"externalIp"`
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort pulumi.StringOutput `pulumi:"externalPort"`
	// The destination private IP address.
	InternalIp pulumi.StringOutput `pulumi:"internalIp"`
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort pulumi.StringOutput `pulumi:"internalPort"`
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The ID of the SAG instance.
	SagId pulumi.StringOutput `pulumi:"sagId"`
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDnatEntry registers a new resource with the given unique name, arguments, and options.
func NewDnatEntry(ctx *pulumi.Context,
	name string, args *DnatEntryArgs, opts ...pulumi.ResourceOption) (*DnatEntry, error) {
	if args == nil || args.ExternalPort == nil {
		return nil, errors.New("missing required argument 'ExternalPort'")
	}
	if args == nil || args.InternalIp == nil {
		return nil, errors.New("missing required argument 'InternalIp'")
	}
	if args == nil || args.InternalPort == nil {
		return nil, errors.New("missing required argument 'InternalPort'")
	}
	if args == nil || args.IpProtocol == nil {
		return nil, errors.New("missing required argument 'IpProtocol'")
	}
	if args == nil || args.SagId == nil {
		return nil, errors.New("missing required argument 'SagId'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &DnatEntryArgs{}
	}
	var resource DnatEntry
	err := ctx.RegisterResource("alicloud:rocketmq/dnatEntry:DnatEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnatEntry gets an existing DnatEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnatEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnatEntryState, opts ...pulumi.ResourceOption) (*DnatEntry, error) {
	var resource DnatEntry
	err := ctx.ReadResource("alicloud:rocketmq/dnatEntry:DnatEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnatEntry resources.
type dnatEntryState struct {
	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp *string `pulumi:"externalIp"`
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort *string `pulumi:"externalPort"`
	// The destination private IP address.
	InternalIp *string `pulumi:"internalIp"`
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort *string `pulumi:"internalPort"`
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The ID of the SAG instance.
	SagId *string `pulumi:"sagId"`
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type *string `pulumi:"type"`
}

type DnatEntryState struct {
	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp pulumi.StringPtrInput
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort pulumi.StringPtrInput
	// The destination private IP address.
	InternalIp pulumi.StringPtrInput
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort pulumi.StringPtrInput
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol pulumi.StringPtrInput
	// The ID of the SAG instance.
	SagId pulumi.StringPtrInput
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type pulumi.StringPtrInput
}

func (DnatEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnatEntryState)(nil)).Elem()
}

type dnatEntryArgs struct {
	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp *string `pulumi:"externalIp"`
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort string `pulumi:"externalPort"`
	// The destination private IP address.
	InternalIp string `pulumi:"internalIp"`
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort string `pulumi:"internalPort"`
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol string `pulumi:"ipProtocol"`
	// The ID of the SAG instance.
	SagId string `pulumi:"sagId"`
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DnatEntry resource.
type DnatEntryArgs struct {
	// The external public IP address.when "type" is "Internet",automatically identify the external ip.
	ExternalIp pulumi.StringPtrInput
	// The public port.Value range: 1 to 65535 or "any".
	ExternalPort pulumi.StringInput
	// The destination private IP address.
	InternalIp pulumi.StringInput
	// The destination private port.Value range: 1 to 65535 or "any".
	InternalPort pulumi.StringInput
	// The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.
	IpProtocol pulumi.StringInput
	// The ID of the SAG instance.
	SagId pulumi.StringInput
	// The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses
	Type pulumi.StringInput
}

func (DnatEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnatEntryArgs)(nil)).Elem()
}
