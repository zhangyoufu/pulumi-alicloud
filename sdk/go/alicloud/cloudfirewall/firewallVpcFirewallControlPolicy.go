// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Firewall Vpc Firewall Control Policy resource.
//
// For information about Cloud Firewall Vpc Firewall Control Policy and how to use it, see [What is Vpc Firewall Control Policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcontrolpolicy).
//
// > **NOTE:** Available since v1.194.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultAccount, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", &cen.InstanceArgs{
//				CenInstanceName: pulumi.String(name),
//				Description:     pulumi.String("example_value"),
//				Tags: pulumi.Map{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("acceptance test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudfirewall.NewFirewallVpcFirewallControlPolicy(ctx, "defaultFirewallVpcFirewallControlPolicy", &cloudfirewall.FirewallVpcFirewallControlPolicyArgs{
//				Order:           pulumi.Int(1),
//				Destination:     pulumi.String("127.0.0.2/32"),
//				ApplicationName: pulumi.String("ANY"),
//				Description:     pulumi.String("example_value"),
//				SourceType:      pulumi.String("net"),
//				DestPort:        pulumi.String("80/88"),
//				AclAction:       pulumi.String("accept"),
//				Lang:            pulumi.String("zh"),
//				DestinationType: pulumi.String("net"),
//				Source:          pulumi.String("127.0.0.1/32"),
//				DestPortType:    pulumi.String("port"),
//				Proto:           pulumi.String("TCP"),
//				Release:         pulumi.Bool(true),
//				MemberUid:       *pulumi.String(defaultAccount.Id),
//				VpcFirewallId:   defaultInstance.ID(),
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
// Cloud Firewall Vpc Firewall Control Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy example <vpc_firewall_id>:<acl_uuid>
// ```
type FirewallVpcFirewallControlPolicy struct {
	pulumi.CustomResourceState

	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction pulumi.StringOutput `pulumi:"aclAction"`
	// Access control over VPC firewalls strategy unique identifier.
	AclUuid pulumi.StringOutput `pulumi:"aclUuid"`
	// Policy specifies the application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// Access control over VPC firewalls description of the strategy information.
	Description pulumi.StringOutput `pulumi:"description"`
	// The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
	DestPort pulumi.StringOutput `pulumi:"destPort"`
	// Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
	DestPortGroup pulumi.StringPtrOutput `pulumi:"destPortGroup"`
	// Port Address Book port list.
	DestPortGroupPorts pulumi.StringArrayOutput `pulumi:"destPortGroupPorts"`
	// The type of the destination port in the access control policy. Valid values: `port`, `group`.
	DestPortType pulumi.StringOutput `pulumi:"destPortType"`
	// The destination address in the access control policy. Valid values:
	// - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
	// - If `destinationType` is set to `group`, the value of `destination` must be an address book.
	// - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Destination address book defined in the address list.
	DestinationGroupCidrs pulumi.StringArrayOutput `pulumi:"destinationGroupCidrs"`
	// The destination address book type in the access control policy.
	DestinationGroupType pulumi.StringOutput `pulumi:"destinationGroupType"`
	// The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
	DestinationType pulumi.StringOutput `pulumi:"destinationType"`
	// Control strategy of hits per second.
	HitTimes pulumi.IntOutput `pulumi:"hitTimes"`
	// The language of the content within the request and response. Valid values: `zh`, `en`.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The UID of the member account of the current Alibaba cloud account.
	MemberUid pulumi.StringOutput `pulumi:"memberUid"`
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order pulumi.IntOutput `pulumi:"order"`
	// The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto pulumi.StringOutput `pulumi:"proto"`
	// The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
	Release pulumi.BoolOutput `pulumi:"release"`
	// Access control over VPC firewalls strategy in the source address.
	Source pulumi.StringOutput `pulumi:"source"`
	// SOURCE address of the address list.
	SourceGroupCidrs pulumi.StringArrayOutput `pulumi:"sourceGroupCidrs"`
	// The source address type in the access control policy.
	SourceGroupType pulumi.StringOutput `pulumi:"sourceGroupType"`
	// The type of the source address in the access control policy. Valid values: `net`, `group`.
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// The ID of the VPC firewall instance. Valid values:
	// - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
	// - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
	VpcFirewallId pulumi.StringOutput `pulumi:"vpcFirewallId"`
}

// NewFirewallVpcFirewallControlPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallVpcFirewallControlPolicy(ctx *pulumi.Context,
	name string, args *FirewallVpcFirewallControlPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallVpcFirewallControlPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclAction == nil {
		return nil, errors.New("invalid value for required argument 'AclAction'")
	}
	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.DestinationType == nil {
		return nil, errors.New("invalid value for required argument 'DestinationType'")
	}
	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.Proto == nil {
		return nil, errors.New("invalid value for required argument 'Proto'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	if args.VpcFirewallId == nil {
		return nil, errors.New("invalid value for required argument 'VpcFirewallId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallVpcFirewallControlPolicy
	err := ctx.RegisterResource("alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVpcFirewallControlPolicy gets an existing FirewallVpcFirewallControlPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVpcFirewallControlPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVpcFirewallControlPolicyState, opts ...pulumi.ResourceOption) (*FirewallVpcFirewallControlPolicy, error) {
	var resource FirewallVpcFirewallControlPolicy
	err := ctx.ReadResource("alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVpcFirewallControlPolicy resources.
type firewallVpcFirewallControlPolicyState struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction *string `pulumi:"aclAction"`
	// Access control over VPC firewalls strategy unique identifier.
	AclUuid *string `pulumi:"aclUuid"`
	// Policy specifies the application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
	ApplicationName *string `pulumi:"applicationName"`
	// Access control over VPC firewalls description of the strategy information.
	Description *string `pulumi:"description"`
	// The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
	DestPort *string `pulumi:"destPort"`
	// Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
	DestPortGroup *string `pulumi:"destPortGroup"`
	// Port Address Book port list.
	DestPortGroupPorts []string `pulumi:"destPortGroupPorts"`
	// The type of the destination port in the access control policy. Valid values: `port`, `group`.
	DestPortType *string `pulumi:"destPortType"`
	// The destination address in the access control policy. Valid values:
	// - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
	// - If `destinationType` is set to `group`, the value of `destination` must be an address book.
	// - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
	Destination *string `pulumi:"destination"`
	// Destination address book defined in the address list.
	DestinationGroupCidrs []string `pulumi:"destinationGroupCidrs"`
	// The destination address book type in the access control policy.
	DestinationGroupType *string `pulumi:"destinationGroupType"`
	// The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
	DestinationType *string `pulumi:"destinationType"`
	// Control strategy of hits per second.
	HitTimes *int `pulumi:"hitTimes"`
	// The language of the content within the request and response. Valid values: `zh`, `en`.
	Lang *string `pulumi:"lang"`
	// The UID of the member account of the current Alibaba cloud account.
	MemberUid *string `pulumi:"memberUid"`
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order *int `pulumi:"order"`
	// The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto *string `pulumi:"proto"`
	// The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
	Release *bool `pulumi:"release"`
	// Access control over VPC firewalls strategy in the source address.
	Source *string `pulumi:"source"`
	// SOURCE address of the address list.
	SourceGroupCidrs []string `pulumi:"sourceGroupCidrs"`
	// The source address type in the access control policy.
	SourceGroupType *string `pulumi:"sourceGroupType"`
	// The type of the source address in the access control policy. Valid values: `net`, `group`.
	SourceType *string `pulumi:"sourceType"`
	// The ID of the VPC firewall instance. Valid values:
	// - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
	// - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
	VpcFirewallId *string `pulumi:"vpcFirewallId"`
}

type FirewallVpcFirewallControlPolicyState struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction pulumi.StringPtrInput
	// Access control over VPC firewalls strategy unique identifier.
	AclUuid pulumi.StringPtrInput
	// Policy specifies the application ID.
	ApplicationId pulumi.StringPtrInput
	// The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
	ApplicationName pulumi.StringPtrInput
	// Access control over VPC firewalls description of the strategy information.
	Description pulumi.StringPtrInput
	// The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
	DestPort pulumi.StringPtrInput
	// Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
	DestPortGroup pulumi.StringPtrInput
	// Port Address Book port list.
	DestPortGroupPorts pulumi.StringArrayInput
	// The type of the destination port in the access control policy. Valid values: `port`, `group`.
	DestPortType pulumi.StringPtrInput
	// The destination address in the access control policy. Valid values:
	// - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
	// - If `destinationType` is set to `group`, the value of `destination` must be an address book.
	// - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
	Destination pulumi.StringPtrInput
	// Destination address book defined in the address list.
	DestinationGroupCidrs pulumi.StringArrayInput
	// The destination address book type in the access control policy.
	DestinationGroupType pulumi.StringPtrInput
	// The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
	DestinationType pulumi.StringPtrInput
	// Control strategy of hits per second.
	HitTimes pulumi.IntPtrInput
	// The language of the content within the request and response. Valid values: `zh`, `en`.
	Lang pulumi.StringPtrInput
	// The UID of the member account of the current Alibaba cloud account.
	MemberUid pulumi.StringPtrInput
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order pulumi.IntPtrInput
	// The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto pulumi.StringPtrInput
	// The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
	Release pulumi.BoolPtrInput
	// Access control over VPC firewalls strategy in the source address.
	Source pulumi.StringPtrInput
	// SOURCE address of the address list.
	SourceGroupCidrs pulumi.StringArrayInput
	// The source address type in the access control policy.
	SourceGroupType pulumi.StringPtrInput
	// The type of the source address in the access control policy. Valid values: `net`, `group`.
	SourceType pulumi.StringPtrInput
	// The ID of the VPC firewall instance. Valid values:
	// - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
	// - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
	VpcFirewallId pulumi.StringPtrInput
}

func (FirewallVpcFirewallControlPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVpcFirewallControlPolicyState)(nil)).Elem()
}

type firewallVpcFirewallControlPolicyArgs struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction string `pulumi:"aclAction"`
	// The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
	ApplicationName string `pulumi:"applicationName"`
	// Access control over VPC firewalls description of the strategy information.
	Description string `pulumi:"description"`
	// The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
	DestPort *string `pulumi:"destPort"`
	// Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
	DestPortGroup *string `pulumi:"destPortGroup"`
	// The type of the destination port in the access control policy. Valid values: `port`, `group`.
	DestPortType *string `pulumi:"destPortType"`
	// The destination address in the access control policy. Valid values:
	// - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
	// - If `destinationType` is set to `group`, the value of `destination` must be an address book.
	// - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
	Destination string `pulumi:"destination"`
	// The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
	DestinationType string `pulumi:"destinationType"`
	// The language of the content within the request and response. Valid values: `zh`, `en`.
	Lang *string `pulumi:"lang"`
	// The UID of the member account of the current Alibaba cloud account.
	MemberUid *string `pulumi:"memberUid"`
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order int `pulumi:"order"`
	// The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto string `pulumi:"proto"`
	// The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
	Release *bool `pulumi:"release"`
	// Access control over VPC firewalls strategy in the source address.
	Source string `pulumi:"source"`
	// The type of the source address in the access control policy. Valid values: `net`, `group`.
	SourceType string `pulumi:"sourceType"`
	// The ID of the VPC firewall instance. Valid values:
	// - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
	// - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
	VpcFirewallId string `pulumi:"vpcFirewallId"`
}

// The set of arguments for constructing a FirewallVpcFirewallControlPolicy resource.
type FirewallVpcFirewallControlPolicyArgs struct {
	// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
	AclAction pulumi.StringInput
	// The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
	ApplicationName pulumi.StringInput
	// Access control over VPC firewalls description of the strategy information.
	Description pulumi.StringInput
	// The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
	DestPort pulumi.StringPtrInput
	// Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
	DestPortGroup pulumi.StringPtrInput
	// The type of the destination port in the access control policy. Valid values: `port`, `group`.
	DestPortType pulumi.StringPtrInput
	// The destination address in the access control policy. Valid values:
	// - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
	// - If `destinationType` is set to `group`, the value of `destination` must be an address book.
	// - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
	Destination pulumi.StringInput
	// The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
	DestinationType pulumi.StringInput
	// The language of the content within the request and response. Valid values: `zh`, `en`.
	Lang pulumi.StringPtrInput
	// The UID of the member account of the current Alibaba cloud account.
	MemberUid pulumi.StringPtrInput
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order pulumi.IntInput
	// The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
	Proto pulumi.StringInput
	// The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
	Release pulumi.BoolPtrInput
	// Access control over VPC firewalls strategy in the source address.
	Source pulumi.StringInput
	// The type of the source address in the access control policy. Valid values: `net`, `group`.
	SourceType pulumi.StringInput
	// The ID of the VPC firewall instance. Valid values:
	// - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
	// - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
	VpcFirewallId pulumi.StringInput
}

func (FirewallVpcFirewallControlPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVpcFirewallControlPolicyArgs)(nil)).Elem()
}

type FirewallVpcFirewallControlPolicyInput interface {
	pulumi.Input

	ToFirewallVpcFirewallControlPolicyOutput() FirewallVpcFirewallControlPolicyOutput
	ToFirewallVpcFirewallControlPolicyOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyOutput
}

func (*FirewallVpcFirewallControlPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVpcFirewallControlPolicy)(nil)).Elem()
}

func (i *FirewallVpcFirewallControlPolicy) ToFirewallVpcFirewallControlPolicyOutput() FirewallVpcFirewallControlPolicyOutput {
	return i.ToFirewallVpcFirewallControlPolicyOutputWithContext(context.Background())
}

func (i *FirewallVpcFirewallControlPolicy) ToFirewallVpcFirewallControlPolicyOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVpcFirewallControlPolicyOutput)
}

// FirewallVpcFirewallControlPolicyArrayInput is an input type that accepts FirewallVpcFirewallControlPolicyArray and FirewallVpcFirewallControlPolicyArrayOutput values.
// You can construct a concrete instance of `FirewallVpcFirewallControlPolicyArrayInput` via:
//
//	FirewallVpcFirewallControlPolicyArray{ FirewallVpcFirewallControlPolicyArgs{...} }
type FirewallVpcFirewallControlPolicyArrayInput interface {
	pulumi.Input

	ToFirewallVpcFirewallControlPolicyArrayOutput() FirewallVpcFirewallControlPolicyArrayOutput
	ToFirewallVpcFirewallControlPolicyArrayOutputWithContext(context.Context) FirewallVpcFirewallControlPolicyArrayOutput
}

type FirewallVpcFirewallControlPolicyArray []FirewallVpcFirewallControlPolicyInput

func (FirewallVpcFirewallControlPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVpcFirewallControlPolicy)(nil)).Elem()
}

func (i FirewallVpcFirewallControlPolicyArray) ToFirewallVpcFirewallControlPolicyArrayOutput() FirewallVpcFirewallControlPolicyArrayOutput {
	return i.ToFirewallVpcFirewallControlPolicyArrayOutputWithContext(context.Background())
}

func (i FirewallVpcFirewallControlPolicyArray) ToFirewallVpcFirewallControlPolicyArrayOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVpcFirewallControlPolicyArrayOutput)
}

// FirewallVpcFirewallControlPolicyMapInput is an input type that accepts FirewallVpcFirewallControlPolicyMap and FirewallVpcFirewallControlPolicyMapOutput values.
// You can construct a concrete instance of `FirewallVpcFirewallControlPolicyMapInput` via:
//
//	FirewallVpcFirewallControlPolicyMap{ "key": FirewallVpcFirewallControlPolicyArgs{...} }
type FirewallVpcFirewallControlPolicyMapInput interface {
	pulumi.Input

	ToFirewallVpcFirewallControlPolicyMapOutput() FirewallVpcFirewallControlPolicyMapOutput
	ToFirewallVpcFirewallControlPolicyMapOutputWithContext(context.Context) FirewallVpcFirewallControlPolicyMapOutput
}

type FirewallVpcFirewallControlPolicyMap map[string]FirewallVpcFirewallControlPolicyInput

func (FirewallVpcFirewallControlPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVpcFirewallControlPolicy)(nil)).Elem()
}

func (i FirewallVpcFirewallControlPolicyMap) ToFirewallVpcFirewallControlPolicyMapOutput() FirewallVpcFirewallControlPolicyMapOutput {
	return i.ToFirewallVpcFirewallControlPolicyMapOutputWithContext(context.Background())
}

func (i FirewallVpcFirewallControlPolicyMap) ToFirewallVpcFirewallControlPolicyMapOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVpcFirewallControlPolicyMapOutput)
}

type FirewallVpcFirewallControlPolicyOutput struct{ *pulumi.OutputState }

func (FirewallVpcFirewallControlPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVpcFirewallControlPolicy)(nil)).Elem()
}

func (o FirewallVpcFirewallControlPolicyOutput) ToFirewallVpcFirewallControlPolicyOutput() FirewallVpcFirewallControlPolicyOutput {
	return o
}

func (o FirewallVpcFirewallControlPolicyOutput) ToFirewallVpcFirewallControlPolicyOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyOutput {
	return o
}

// The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
func (o FirewallVpcFirewallControlPolicyOutput) AclAction() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.AclAction }).(pulumi.StringOutput)
}

// Access control over VPC firewalls strategy unique identifier.
func (o FirewallVpcFirewallControlPolicyOutput) AclUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.AclUuid }).(pulumi.StringOutput)
}

// Policy specifies the application ID.
func (o FirewallVpcFirewallControlPolicyOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
func (o FirewallVpcFirewallControlPolicyOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// Access control over VPC firewalls description of the strategy information.
func (o FirewallVpcFirewallControlPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The destination port in the access control policy. **Note:** If `destPortType` is set to `port`, you must specify this parameter.
func (o FirewallVpcFirewallControlPolicyOutput) DestPort() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.DestPort }).(pulumi.StringOutput)
}

// Access control policy in the access traffic of the destination port address book name. **Note:** If `destPortType` is set to `group`, you must specify this parameter.
func (o FirewallVpcFirewallControlPolicyOutput) DestPortGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringPtrOutput { return v.DestPortGroup }).(pulumi.StringPtrOutput)
}

// Port Address Book port list.
func (o FirewallVpcFirewallControlPolicyOutput) DestPortGroupPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringArrayOutput { return v.DestPortGroupPorts }).(pulumi.StringArrayOutput)
}

// The type of the destination port in the access control policy. Valid values: `port`, `group`.
func (o FirewallVpcFirewallControlPolicyOutput) DestPortType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.DestPortType }).(pulumi.StringOutput)
}

// The destination address in the access control policy. Valid values:
// - If `destinationType` is set to `net`, the value of `destination` must be a CIDR block.
// - If `destinationType` is set to `group`, the value of `destination` must be an address book.
// - If `destinationType` is set to `domain`, the value of `destination` must be a domain name.
func (o FirewallVpcFirewallControlPolicyOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Destination address book defined in the address list.
func (o FirewallVpcFirewallControlPolicyOutput) DestinationGroupCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringArrayOutput { return v.DestinationGroupCidrs }).(pulumi.StringArrayOutput)
}

// The destination address book type in the access control policy.
func (o FirewallVpcFirewallControlPolicyOutput) DestinationGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.DestinationGroupType }).(pulumi.StringOutput)
}

// The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
func (o FirewallVpcFirewallControlPolicyOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

// Control strategy of hits per second.
func (o FirewallVpcFirewallControlPolicyOutput) HitTimes() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.IntOutput { return v.HitTimes }).(pulumi.IntOutput)
}

// The language of the content within the request and response. Valid values: `zh`, `en`.
func (o FirewallVpcFirewallControlPolicyOutput) Lang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringPtrOutput { return v.Lang }).(pulumi.StringPtrOutput)
}

// The UID of the member account of the current Alibaba cloud account.
func (o FirewallVpcFirewallControlPolicyOutput) MemberUid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.MemberUid }).(pulumi.StringOutput)
}

// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
func (o FirewallVpcFirewallControlPolicyOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
func (o FirewallVpcFirewallControlPolicyOutput) Proto() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.Proto }).(pulumi.StringOutput)
}

// The enabled status of the access control policy. The policy is enabled by default after it is created.. Valid values:
func (o FirewallVpcFirewallControlPolicyOutput) Release() pulumi.BoolOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.BoolOutput { return v.Release }).(pulumi.BoolOutput)
}

// Access control over VPC firewalls strategy in the source address.
func (o FirewallVpcFirewallControlPolicyOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// SOURCE address of the address list.
func (o FirewallVpcFirewallControlPolicyOutput) SourceGroupCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringArrayOutput { return v.SourceGroupCidrs }).(pulumi.StringArrayOutput)
}

// The source address type in the access control policy.
func (o FirewallVpcFirewallControlPolicyOutput) SourceGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.SourceGroupType }).(pulumi.StringOutput)
}

// The type of the source address in the access control policy. Valid values: `net`, `group`.
func (o FirewallVpcFirewallControlPolicyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

// The ID of the VPC firewall instance. Valid values:
// - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
// - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
func (o FirewallVpcFirewallControlPolicyOutput) VpcFirewallId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVpcFirewallControlPolicy) pulumi.StringOutput { return v.VpcFirewallId }).(pulumi.StringOutput)
}

type FirewallVpcFirewallControlPolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallVpcFirewallControlPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVpcFirewallControlPolicy)(nil)).Elem()
}

func (o FirewallVpcFirewallControlPolicyArrayOutput) ToFirewallVpcFirewallControlPolicyArrayOutput() FirewallVpcFirewallControlPolicyArrayOutput {
	return o
}

func (o FirewallVpcFirewallControlPolicyArrayOutput) ToFirewallVpcFirewallControlPolicyArrayOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyArrayOutput {
	return o
}

func (o FirewallVpcFirewallControlPolicyArrayOutput) Index(i pulumi.IntInput) FirewallVpcFirewallControlPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallVpcFirewallControlPolicy {
		return vs[0].([]*FirewallVpcFirewallControlPolicy)[vs[1].(int)]
	}).(FirewallVpcFirewallControlPolicyOutput)
}

type FirewallVpcFirewallControlPolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallVpcFirewallControlPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVpcFirewallControlPolicy)(nil)).Elem()
}

func (o FirewallVpcFirewallControlPolicyMapOutput) ToFirewallVpcFirewallControlPolicyMapOutput() FirewallVpcFirewallControlPolicyMapOutput {
	return o
}

func (o FirewallVpcFirewallControlPolicyMapOutput) ToFirewallVpcFirewallControlPolicyMapOutputWithContext(ctx context.Context) FirewallVpcFirewallControlPolicyMapOutput {
	return o
}

func (o FirewallVpcFirewallControlPolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallVpcFirewallControlPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallVpcFirewallControlPolicy {
		return vs[0].(map[string]*FirewallVpcFirewallControlPolicy)[vs[1].(string)]
	}).(FirewallVpcFirewallControlPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVpcFirewallControlPolicyInput)(nil)).Elem(), &FirewallVpcFirewallControlPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVpcFirewallControlPolicyArrayInput)(nil)).Elem(), FirewallVpcFirewallControlPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVpcFirewallControlPolicyMapInput)(nil)).Elem(), FirewallVpcFirewallControlPolicyMap{})
	pulumi.RegisterOutputType(FirewallVpcFirewallControlPolicyOutput{})
	pulumi.RegisterOutputType(FirewallVpcFirewallControlPolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallVpcFirewallControlPolicyMapOutput{})
}
