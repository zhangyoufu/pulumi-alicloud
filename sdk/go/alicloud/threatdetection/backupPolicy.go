// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Threat Detection Backup Policy resource.
//
// For information about Threat Detection Backup Policy and how to use it, see [What is Backup Policy](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := threatdetection.GetAssets(ctx, &threatdetection.GetAssetsArgs{
//				MachineTypes: pulumi.StringRef("ecs"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = threatdetection.NewBackupPolicy(ctx, "default", &threatdetection.BackupPolicyArgs{
//				BackupPolicyName: pulumi.String("tf-example-name"),
//				Policy:           pulumi.String("{\"Exclude\":[\"/bin/\",\"/usr/bin/\",\"/sbin/\",\"/boot/\",\"/proc/\",\"/sys/\",\"/srv/\",\"/lib/\",\"/selinux/\",\"/usr/sbin/\",\"/run/\",\"/lib32/\",\"/lib64/\",\"/lost+found/\",\"/var/lib/kubelet/\",\"/var/lib/ntp/proc\",\"/var/lib/container\"],\"ExcludeSystemPath\":true,\"Include\":[],\"IsDefault\":1,\"Retention\":7,\"Schedule\":\"I|1668703620|PT24H\",\"Source\":[],\"SpeedLimiter\":\"\",\"UseVss\":true}"),
//				PolicyVersion:    pulumi.String("2.0.0"),
//				UuidLists: pulumi.StringArray{
//					pulumi.String(_default.Ids[0]),
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
// Threat Detection Backup Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:threatdetection/backupPolicy:BackupPolicy example <id>
// ```
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Protection of the Name of the Policy.
	BackupPolicyName pulumi.StringOutput `pulumi:"backupPolicyName"`
	// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
	PolicyRegionId pulumi.StringPtrOutput `pulumi:"policyRegionId"`
	// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
	PolicyVersion pulumi.StringOutput `pulumi:"policyVersion"`
	// The status of the Backup Policy instance.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specify the Protection of Server UUID List.
	UuidLists pulumi.StringArrayOutput `pulumi:"uuidLists"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'BackupPolicyName'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.PolicyVersion == nil {
		return nil, errors.New("invalid value for required argument 'PolicyVersion'")
	}
	if args.UuidLists == nil {
		return nil, errors.New("invalid value for required argument 'UuidLists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupPolicy
	err := ctx.RegisterResource("alicloud:threatdetection/backupPolicy:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("alicloud:threatdetection/backupPolicy:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Protection of the Name of the Policy.
	BackupPolicyName *string `pulumi:"backupPolicyName"`
	// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
	Policy *string `pulumi:"policy"`
	// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
	PolicyRegionId *string `pulumi:"policyRegionId"`
	// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
	PolicyVersion *string `pulumi:"policyVersion"`
	// The status of the Backup Policy instance.
	Status *string `pulumi:"status"`
	// Specify the Protection of Server UUID List.
	UuidLists []string `pulumi:"uuidLists"`
}

type BackupPolicyState struct {
	// Protection of the Name of the Policy.
	BackupPolicyName pulumi.StringPtrInput
	// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
	Policy pulumi.StringPtrInput
	// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
	PolicyRegionId pulumi.StringPtrInput
	// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
	PolicyVersion pulumi.StringPtrInput
	// The status of the Backup Policy instance.
	Status pulumi.StringPtrInput
	// Specify the Protection of Server UUID List.
	UuidLists pulumi.StringArrayInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// Protection of the Name of the Policy.
	BackupPolicyName string `pulumi:"backupPolicyName"`
	// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
	Policy string `pulumi:"policy"`
	// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
	PolicyRegionId *string `pulumi:"policyRegionId"`
	// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
	PolicyVersion string `pulumi:"policyVersion"`
	// Specify the Protection of Server UUID List.
	UuidLists []string `pulumi:"uuidLists"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// Protection of the Name of the Policy.
	BackupPolicyName pulumi.StringInput
	// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
	Policy pulumi.StringInput
	// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
	PolicyRegionId pulumi.StringPtrInput
	// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
	PolicyVersion pulumi.StringInput
	// Specify the Protection of Server UUID List.
	UuidLists pulumi.StringArrayInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

// BackupPolicyArrayInput is an input type that accepts BackupPolicyArray and BackupPolicyArrayOutput values.
// You can construct a concrete instance of `BackupPolicyArrayInput` via:
//
//	BackupPolicyArray{ BackupPolicyArgs{...} }
type BackupPolicyArrayInput interface {
	pulumi.Input

	ToBackupPolicyArrayOutput() BackupPolicyArrayOutput
	ToBackupPolicyArrayOutputWithContext(context.Context) BackupPolicyArrayOutput
}

type BackupPolicyArray []BackupPolicyInput

func (BackupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return i.ToBackupPolicyArrayOutputWithContext(context.Background())
}

func (i BackupPolicyArray) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyArrayOutput)
}

// BackupPolicyMapInput is an input type that accepts BackupPolicyMap and BackupPolicyMapOutput values.
// You can construct a concrete instance of `BackupPolicyMapInput` via:
//
//	BackupPolicyMap{ "key": BackupPolicyArgs{...} }
type BackupPolicyMapInput interface {
	pulumi.Input

	ToBackupPolicyMapOutput() BackupPolicyMapOutput
	ToBackupPolicyMapOutputWithContext(context.Context) BackupPolicyMapOutput
}

type BackupPolicyMap map[string]BackupPolicyInput

func (BackupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (i BackupPolicyMap) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return i.ToBackupPolicyMapOutputWithContext(context.Background())
}

func (i BackupPolicyMap) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMapOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

// Protection of the Name of the Policy.
func (o BackupPolicyOutput) BackupPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.BackupPolicyName }).(pulumi.StringOutput)
}

// The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
func (o BackupPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
func (o BackupPolicyOutput) PolicyRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringPtrOutput { return v.PolicyRegionId }).(pulumi.StringPtrOutput)
}

// Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
func (o BackupPolicyOutput) PolicyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.PolicyVersion }).(pulumi.StringOutput)
}

// The status of the Backup Policy instance.
func (o BackupPolicyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specify the Protection of Server UUID List.
func (o BackupPolicyOutput) UuidLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupPolicy) pulumi.StringArrayOutput { return v.UuidLists }).(pulumi.StringArrayOutput)
}

type BackupPolicyArrayOutput struct{ *pulumi.OutputState }

func (BackupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutput() BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) ToBackupPolicyArrayOutputWithContext(ctx context.Context) BackupPolicyArrayOutput {
	return o
}

func (o BackupPolicyArrayOutput) Index(i pulumi.IntInput) BackupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].([]*BackupPolicy)[vs[1].(int)]
	}).(BackupPolicyOutput)
}

type BackupPolicyMapOutput struct{ *pulumi.OutputState }

func (BackupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPolicy)(nil)).Elem()
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutput() BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) ToBackupPolicyMapOutputWithContext(ctx context.Context) BackupPolicyMapOutput {
	return o
}

func (o BackupPolicyMapOutput) MapIndex(k pulumi.StringInput) BackupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPolicy {
		return vs[0].(map[string]*BackupPolicy)[vs[1].(string)]
	}).(BackupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyInput)(nil)).Elem(), &BackupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyArrayInput)(nil)).Elem(), BackupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyMapInput)(nil)).Elem(), BackupPolicyMap{})
	pulumi.RegisterOutputType(BackupPolicyOutput{})
	pulumi.RegisterOutputType(BackupPolicyArrayOutput{})
	pulumi.RegisterOutputType(BackupPolicyMapOutput{})
}
