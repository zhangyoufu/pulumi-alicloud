// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Anti-DDoS Advanced instance resource. "Ddosbgp" is the short term of this product.
//
// > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.
//
// > **NOTE:** Available in 1.57.0+ .
//
// Deprecated: alicloud.dns.DdosBgpInstance has been deprecated in favor of alicloud.ddos.DdosBgpInstance
type DdosBgpInstance struct {
	pulumi.CustomResourceState

	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth pulumi.IntPtrOutput `pulumi:"baseBandwidth"`
	// IP count of the instance. Valid values: 100.
	IpCount pulumi.IntOutput `pulumi:"ipCount"`
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType pulumi.StringOutput `pulumi:"ipType"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDdosBgpInstance registers a new resource with the given unique name, arguments, and options.
func NewDdosBgpInstance(ctx *pulumi.Context,
	name string, args *DdosBgpInstanceArgs, opts ...pulumi.ResourceOption) (*DdosBgpInstance, error) {
	if args == nil || args.Bandwidth == nil {
		return nil, errors.New("missing required argument 'Bandwidth'")
	}
	if args == nil || args.IpCount == nil {
		return nil, errors.New("missing required argument 'IpCount'")
	}
	if args == nil || args.IpType == nil {
		return nil, errors.New("missing required argument 'IpType'")
	}
	if args == nil {
		args = &DdosBgpInstanceArgs{}
	}
	var resource DdosBgpInstance
	err := ctx.RegisterResource("alicloud:dns/ddosBgpInstance:DdosBgpInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosBgpInstance gets an existing DdosBgpInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosBgpInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosBgpInstanceState, opts ...pulumi.ResourceOption) (*DdosBgpInstance, error) {
	var resource DdosBgpInstance
	err := ctx.ReadResource("alicloud:dns/ddosBgpInstance:DdosBgpInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosBgpInstance resources.
type ddosBgpInstanceState struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth *int `pulumi:"baseBandwidth"`
	// IP count of the instance. Valid values: 100.
	IpCount *int `pulumi:"ipCount"`
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType *string `pulumi:"ipType"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name *string `pulumi:"name"`
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
	Type *string `pulumi:"type"`
}

type DdosBgpInstanceState struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth pulumi.IntPtrInput
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth pulumi.IntPtrInput
	// IP count of the instance. Valid values: 100.
	IpCount pulumi.IntPtrInput
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType pulumi.StringPtrInput
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringPtrInput
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
	Type pulumi.StringPtrInput
}

func (DdosBgpInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosBgpInstanceState)(nil)).Elem()
}

type ddosBgpInstanceArgs struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth int `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth *int `pulumi:"baseBandwidth"`
	// IP count of the instance. Valid values: 100.
	IpCount int `pulumi:"ipCount"`
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType string `pulumi:"ipType"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name *string `pulumi:"name"`
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DdosBgpInstance resource.
type DdosBgpInstanceArgs struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.
	Bandwidth pulumi.IntInput
	// Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to `20`.
	BaseBandwidth pulumi.IntPtrInput
	// IP count of the instance. Valid values: 100.
	IpCount pulumi.IntInput
	// IP version of the instance. Valid values: IPv4,IPv6.
	IpType pulumi.StringInput
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringPtrInput
	// The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Type of the instance. Valid values: Enterprise,Professional. Default to `Enterprise`
	Type pulumi.StringPtrInput
}

func (DdosBgpInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosBgpInstanceArgs)(nil)).Elem()
}
