// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OOS Patch Baseline resource.
//
// For information about OOS Patch Baseline and how to use it, see [What is Patch Baseline](https://www.alibabacloud.com/help/en/operation-orchestration-service/latest/patch-manager-overview).
//
// > **NOTE:** Available since v1.146.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/oos"
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
//			_, err := oos.NewPatchBaseline(ctx, "default", &oos.PatchBaselineArgs{
//				PatchBaselineName: pulumi.String(name),
//				OperationSystem:   pulumi.String("Windows"),
//				ApprovalRules:     pulumi.String("{\"PatchRules\":[{\"EnableNonSecurity\":true,\"PatchFilterGroup\":[{\"Values\":[\"*\"],\"Key\":\"Product\"},{\"Values\":[\"Security\",\"Bugfix\"],\"Key\":\"Classification\"},{\"Values\":[\"Critical\",\"Important\"],\"Key\":\"Severity\"}],\"ApproveAfterDays\":7,\"ComplianceLevel\":\"Unspecified\"}]}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OOS Patch Baseline can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:oos/patchBaseline:PatchBaseline example <id>
// ```
type PatchBaseline struct {
	pulumi.CustomResourceState

	// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
	ApprovalRules pulumi.StringOutput `pulumi:"approvalRules"`
	// Approved Patch.
	ApprovedPatches pulumi.StringArrayOutput `pulumi:"approvedPatches"`
	// ApprovedPatchesEnableNonSecurity.
	ApprovedPatchesEnableNonSecurity pulumi.BoolPtrOutput `pulumi:"approvedPatchesEnableNonSecurity"`
	// Creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Patches baseline description information.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
	OperationSystem pulumi.StringOutput `pulumi:"operationSystem"`
	// The name of the patch baseline.
	PatchBaselineName pulumi.StringOutput `pulumi:"patchBaselineName"`
	// Reject patches.
	RejectedPatches pulumi.StringArrayOutput `pulumi:"rejectedPatches"`
	// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
	RejectedPatchesAction pulumi.StringOutput `pulumi:"rejectedPatchesAction"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Source.
	Sources pulumi.StringArrayOutput `pulumi:"sources"`
	// Label.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPatchBaseline registers a new resource with the given unique name, arguments, and options.
func NewPatchBaseline(ctx *pulumi.Context,
	name string, args *PatchBaselineArgs, opts ...pulumi.ResourceOption) (*PatchBaseline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApprovalRules == nil {
		return nil, errors.New("invalid value for required argument 'ApprovalRules'")
	}
	if args.OperationSystem == nil {
		return nil, errors.New("invalid value for required argument 'OperationSystem'")
	}
	if args.PatchBaselineName == nil {
		return nil, errors.New("invalid value for required argument 'PatchBaselineName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PatchBaseline
	err := ctx.RegisterResource("alicloud:oos/patchBaseline:PatchBaseline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchBaseline gets an existing PatchBaseline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchBaseline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchBaselineState, opts ...pulumi.ResourceOption) (*PatchBaseline, error) {
	var resource PatchBaseline
	err := ctx.ReadResource("alicloud:oos/patchBaseline:PatchBaseline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchBaseline resources.
type patchBaselineState struct {
	// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
	ApprovalRules *string `pulumi:"approvalRules"`
	// Approved Patch.
	ApprovedPatches []string `pulumi:"approvedPatches"`
	// ApprovedPatchesEnableNonSecurity.
	ApprovedPatchesEnableNonSecurity *bool `pulumi:"approvedPatchesEnableNonSecurity"`
	// Creation time.
	CreateTime *string `pulumi:"createTime"`
	// Patches baseline description information.
	Description *string `pulumi:"description"`
	// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
	OperationSystem *string `pulumi:"operationSystem"`
	// The name of the patch baseline.
	PatchBaselineName *string `pulumi:"patchBaselineName"`
	// Reject patches.
	RejectedPatches []string `pulumi:"rejectedPatches"`
	// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
	RejectedPatchesAction *string `pulumi:"rejectedPatchesAction"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Source.
	Sources []string `pulumi:"sources"`
	// Label.
	Tags map[string]string `pulumi:"tags"`
}

type PatchBaselineState struct {
	// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
	ApprovalRules pulumi.StringPtrInput
	// Approved Patch.
	ApprovedPatches pulumi.StringArrayInput
	// ApprovedPatchesEnableNonSecurity.
	ApprovedPatchesEnableNonSecurity pulumi.BoolPtrInput
	// Creation time.
	CreateTime pulumi.StringPtrInput
	// Patches baseline description information.
	Description pulumi.StringPtrInput
	// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
	OperationSystem pulumi.StringPtrInput
	// The name of the patch baseline.
	PatchBaselineName pulumi.StringPtrInput
	// Reject patches.
	RejectedPatches pulumi.StringArrayInput
	// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
	RejectedPatchesAction pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// Source.
	Sources pulumi.StringArrayInput
	// Label.
	Tags pulumi.StringMapInput
}

func (PatchBaselineState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchBaselineState)(nil)).Elem()
}

type patchBaselineArgs struct {
	// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
	ApprovalRules string `pulumi:"approvalRules"`
	// Approved Patch.
	ApprovedPatches []string `pulumi:"approvedPatches"`
	// ApprovedPatchesEnableNonSecurity.
	ApprovedPatchesEnableNonSecurity *bool `pulumi:"approvedPatchesEnableNonSecurity"`
	// Patches baseline description information.
	Description *string `pulumi:"description"`
	// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
	OperationSystem string `pulumi:"operationSystem"`
	// The name of the patch baseline.
	PatchBaselineName string `pulumi:"patchBaselineName"`
	// Reject patches.
	RejectedPatches []string `pulumi:"rejectedPatches"`
	// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
	RejectedPatchesAction *string `pulumi:"rejectedPatchesAction"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Source.
	Sources []string `pulumi:"sources"`
	// Label.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PatchBaseline resource.
type PatchBaselineArgs struct {
	// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
	ApprovalRules pulumi.StringInput
	// Approved Patch.
	ApprovedPatches pulumi.StringArrayInput
	// ApprovedPatchesEnableNonSecurity.
	ApprovedPatchesEnableNonSecurity pulumi.BoolPtrInput
	// Patches baseline description information.
	Description pulumi.StringPtrInput
	// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
	OperationSystem pulumi.StringInput
	// The name of the patch baseline.
	PatchBaselineName pulumi.StringInput
	// Reject patches.
	RejectedPatches pulumi.StringArrayInput
	// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
	RejectedPatchesAction pulumi.StringPtrInput
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput
	// Source.
	Sources pulumi.StringArrayInput
	// Label.
	Tags pulumi.StringMapInput
}

func (PatchBaselineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchBaselineArgs)(nil)).Elem()
}

type PatchBaselineInput interface {
	pulumi.Input

	ToPatchBaselineOutput() PatchBaselineOutput
	ToPatchBaselineOutputWithContext(ctx context.Context) PatchBaselineOutput
}

func (*PatchBaseline) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchBaseline)(nil)).Elem()
}

func (i *PatchBaseline) ToPatchBaselineOutput() PatchBaselineOutput {
	return i.ToPatchBaselineOutputWithContext(context.Background())
}

func (i *PatchBaseline) ToPatchBaselineOutputWithContext(ctx context.Context) PatchBaselineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchBaselineOutput)
}

// PatchBaselineArrayInput is an input type that accepts PatchBaselineArray and PatchBaselineArrayOutput values.
// You can construct a concrete instance of `PatchBaselineArrayInput` via:
//
//	PatchBaselineArray{ PatchBaselineArgs{...} }
type PatchBaselineArrayInput interface {
	pulumi.Input

	ToPatchBaselineArrayOutput() PatchBaselineArrayOutput
	ToPatchBaselineArrayOutputWithContext(context.Context) PatchBaselineArrayOutput
}

type PatchBaselineArray []PatchBaselineInput

func (PatchBaselineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchBaseline)(nil)).Elem()
}

func (i PatchBaselineArray) ToPatchBaselineArrayOutput() PatchBaselineArrayOutput {
	return i.ToPatchBaselineArrayOutputWithContext(context.Background())
}

func (i PatchBaselineArray) ToPatchBaselineArrayOutputWithContext(ctx context.Context) PatchBaselineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchBaselineArrayOutput)
}

// PatchBaselineMapInput is an input type that accepts PatchBaselineMap and PatchBaselineMapOutput values.
// You can construct a concrete instance of `PatchBaselineMapInput` via:
//
//	PatchBaselineMap{ "key": PatchBaselineArgs{...} }
type PatchBaselineMapInput interface {
	pulumi.Input

	ToPatchBaselineMapOutput() PatchBaselineMapOutput
	ToPatchBaselineMapOutputWithContext(context.Context) PatchBaselineMapOutput
}

type PatchBaselineMap map[string]PatchBaselineInput

func (PatchBaselineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchBaseline)(nil)).Elem()
}

func (i PatchBaselineMap) ToPatchBaselineMapOutput() PatchBaselineMapOutput {
	return i.ToPatchBaselineMapOutputWithContext(context.Background())
}

func (i PatchBaselineMap) ToPatchBaselineMapOutputWithContext(ctx context.Context) PatchBaselineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchBaselineMapOutput)
}

type PatchBaselineOutput struct{ *pulumi.OutputState }

func (PatchBaselineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchBaseline)(nil)).Elem()
}

func (o PatchBaselineOutput) ToPatchBaselineOutput() PatchBaselineOutput {
	return o
}

func (o PatchBaselineOutput) ToPatchBaselineOutputWithContext(ctx context.Context) PatchBaselineOutput {
	return o
}

// Accept the rules. This value follows the json format. For more details, see the description of [ApprovalRules in the Request parameters table for details](https://www.alibabacloud.com/help/zh/operation-orchestration-service/latest/api-oos-2019-06-01-createpatchbaseline).
func (o PatchBaselineOutput) ApprovalRules() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.ApprovalRules }).(pulumi.StringOutput)
}

// Approved Patch.
func (o PatchBaselineOutput) ApprovedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringArrayOutput { return v.ApprovedPatches }).(pulumi.StringArrayOutput)
}

// ApprovedPatchesEnableNonSecurity.
func (o PatchBaselineOutput) ApprovedPatchesEnableNonSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.BoolPtrOutput { return v.ApprovedPatchesEnableNonSecurity }).(pulumi.BoolPtrOutput)
}

// Creation time.
func (o PatchBaselineOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Patches baseline description information.
func (o PatchBaselineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Operating system type. Valid values: `AliyunLinux`, `Anolis`, `CentOS`, `Debian`, `RedhatEnterpriseLinux`, `Ubuntu`, `Windows`, `AlmaLinux`.
func (o PatchBaselineOutput) OperationSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.OperationSystem }).(pulumi.StringOutput)
}

// The name of the patch baseline.
func (o PatchBaselineOutput) PatchBaselineName() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.PatchBaselineName }).(pulumi.StringOutput)
}

// Reject patches.
func (o PatchBaselineOutput) RejectedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringArrayOutput { return v.RejectedPatches }).(pulumi.StringArrayOutput)
}

// Rejected patches action. Valid values: `ALLOW_AS_DEPENDENCY`, `BLOCK`.
func (o PatchBaselineOutput) RejectedPatchesAction() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.RejectedPatchesAction }).(pulumi.StringOutput)
}

// The ID of the resource group.
func (o PatchBaselineOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Source.
func (o PatchBaselineOutput) Sources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringArrayOutput { return v.Sources }).(pulumi.StringArrayOutput)
}

// Label.
func (o PatchBaselineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PatchBaseline) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

type PatchBaselineArrayOutput struct{ *pulumi.OutputState }

func (PatchBaselineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchBaseline)(nil)).Elem()
}

func (o PatchBaselineArrayOutput) ToPatchBaselineArrayOutput() PatchBaselineArrayOutput {
	return o
}

func (o PatchBaselineArrayOutput) ToPatchBaselineArrayOutputWithContext(ctx context.Context) PatchBaselineArrayOutput {
	return o
}

func (o PatchBaselineArrayOutput) Index(i pulumi.IntInput) PatchBaselineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PatchBaseline {
		return vs[0].([]*PatchBaseline)[vs[1].(int)]
	}).(PatchBaselineOutput)
}

type PatchBaselineMapOutput struct{ *pulumi.OutputState }

func (PatchBaselineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchBaseline)(nil)).Elem()
}

func (o PatchBaselineMapOutput) ToPatchBaselineMapOutput() PatchBaselineMapOutput {
	return o
}

func (o PatchBaselineMapOutput) ToPatchBaselineMapOutputWithContext(ctx context.Context) PatchBaselineMapOutput {
	return o
}

func (o PatchBaselineMapOutput) MapIndex(k pulumi.StringInput) PatchBaselineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PatchBaseline {
		return vs[0].(map[string]*PatchBaseline)[vs[1].(string)]
	}).(PatchBaselineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PatchBaselineInput)(nil)).Elem(), &PatchBaseline{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchBaselineArrayInput)(nil)).Elem(), PatchBaselineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchBaselineMapInput)(nil)).Elem(), PatchBaselineMap{})
	pulumi.RegisterOutputType(PatchBaselineOutput{})
	pulumi.RegisterOutputType(PatchBaselineArrayOutput{})
	pulumi.RegisterOutputType(PatchBaselineMapOutput{})
}
