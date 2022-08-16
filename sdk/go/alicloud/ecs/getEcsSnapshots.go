// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecs Snapshots of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.120.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ecs.GetEcsSnapshots(ctx, &ecs.GetEcsSnapshotsArgs{
//				Ids: []string{
//					"s-bp1fvuxxxxxxxx",
//				},
//				NameRegex: pulumi.StringRef("tf-test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstEcsSnapshotId", example.Snapshots[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetEcsSnapshots(ctx *pulumi.Context, args *GetEcsSnapshotsArgs, opts ...pulumi.InvokeOption) (*GetEcsSnapshotsResult, error) {
	var rv GetEcsSnapshotsResult
	err := ctx.Invoke("alicloud:ecs/getEcsSnapshots:getEcsSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsSnapshots.
type GetEcsSnapshotsArgs struct {
	// The category of the snapshot.
	Category *string `pulumi:"category"`
	// Specifies whether to check the validity of the request without actually making the request.
	DryRun *bool `pulumi:"dryRun"`
	// Whether the snapshot is encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// A list of Snapshot IDs.
	Ids []string `pulumi:"ids"`
	// The kms key id.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A regex string to filter results by Snapshot name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The resource group id.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The snapshot link id.
	SnapshotLinkId *string `pulumi:"snapshotLinkId"`
	// Snapshot Display Name.
	SnapshotName *string `pulumi:"snapshotName"`
	// Snapshot creation type.
	SnapshotType *string `pulumi:"snapshotType"`
	// Source disk attributes.
	SourceDiskType *string `pulumi:"sourceDiskType"`
	// The status of the snapshot.
	Status *string `pulumi:"status"`
	// The tags.
	Tags map[string]interface{} `pulumi:"tags"`
	Type *string                `pulumi:"type"`
	// A resource type that has a reference relationship.
	Usage *string `pulumi:"usage"`
}

// A collection of values returned by getEcsSnapshots.
type GetEcsSnapshotsResult struct {
	Category  *string `pulumi:"category"`
	DryRun    *bool   `pulumi:"dryRun"`
	Encrypted *bool   `pulumi:"encrypted"`
	// The provider-assigned unique ID for this managed resource.
	Id              string                    `pulumi:"id"`
	Ids             []string                  `pulumi:"ids"`
	KmsKeyId        *string                   `pulumi:"kmsKeyId"`
	NameRegex       *string                   `pulumi:"nameRegex"`
	Names           []string                  `pulumi:"names"`
	OutputFile      *string                   `pulumi:"outputFile"`
	ResourceGroupId *string                   `pulumi:"resourceGroupId"`
	SnapshotLinkId  *string                   `pulumi:"snapshotLinkId"`
	SnapshotName    *string                   `pulumi:"snapshotName"`
	SnapshotType    *string                   `pulumi:"snapshotType"`
	Snapshots       []GetEcsSnapshotsSnapshot `pulumi:"snapshots"`
	SourceDiskType  *string                   `pulumi:"sourceDiskType"`
	Status          *string                   `pulumi:"status"`
	Tags            map[string]interface{}    `pulumi:"tags"`
	Type            *string                   `pulumi:"type"`
	Usage           *string                   `pulumi:"usage"`
}

func GetEcsSnapshotsOutput(ctx *pulumi.Context, args GetEcsSnapshotsOutputArgs, opts ...pulumi.InvokeOption) GetEcsSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEcsSnapshotsResult, error) {
			args := v.(GetEcsSnapshotsArgs)
			r, err := GetEcsSnapshots(ctx, &args, opts...)
			var s GetEcsSnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEcsSnapshotsResultOutput)
}

// A collection of arguments for invoking getEcsSnapshots.
type GetEcsSnapshotsOutputArgs struct {
	// The category of the snapshot.
	Category pulumi.StringPtrInput `pulumi:"category"`
	// Specifies whether to check the validity of the request without actually making the request.
	DryRun pulumi.BoolPtrInput `pulumi:"dryRun"`
	// Whether the snapshot is encrypted.
	Encrypted pulumi.BoolPtrInput `pulumi:"encrypted"`
	// A list of Snapshot IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The kms key id.
	KmsKeyId pulumi.StringPtrInput `pulumi:"kmsKeyId"`
	// A regex string to filter results by Snapshot name.
	NameRegex  pulumi.StringPtrInput `pulumi:"nameRegex"`
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The resource group id.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The snapshot link id.
	SnapshotLinkId pulumi.StringPtrInput `pulumi:"snapshotLinkId"`
	// Snapshot Display Name.
	SnapshotName pulumi.StringPtrInput `pulumi:"snapshotName"`
	// Snapshot creation type.
	SnapshotType pulumi.StringPtrInput `pulumi:"snapshotType"`
	// Source disk attributes.
	SourceDiskType pulumi.StringPtrInput `pulumi:"sourceDiskType"`
	// The status of the snapshot.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The tags.
	Tags pulumi.MapInput       `pulumi:"tags"`
	Type pulumi.StringPtrInput `pulumi:"type"`
	// A resource type that has a reference relationship.
	Usage pulumi.StringPtrInput `pulumi:"usage"`
}

func (GetEcsSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsSnapshotsArgs)(nil)).Elem()
}

// A collection of values returned by getEcsSnapshots.
type GetEcsSnapshotsResultOutput struct{ *pulumi.OutputState }

func (GetEcsSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEcsSnapshotsResult)(nil)).Elem()
}

func (o GetEcsSnapshotsResultOutput) ToGetEcsSnapshotsResultOutput() GetEcsSnapshotsResultOutput {
	return o
}

func (o GetEcsSnapshotsResultOutput) ToGetEcsSnapshotsResultOutputWithContext(ctx context.Context) GetEcsSnapshotsResultOutput {
	return o
}

func (o GetEcsSnapshotsResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *bool { return v.DryRun }).(pulumi.BoolPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *bool { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEcsSnapshotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEcsSnapshotsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEcsSnapshotsResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetEcsSnapshotsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) SnapshotLinkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.SnapshotLinkId }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) SnapshotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.SnapshotType }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) Snapshots() GetEcsSnapshotsSnapshotArrayOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) []GetEcsSnapshotsSnapshot { return v.Snapshots }).(GetEcsSnapshotsSnapshotArrayOutput)
}

func (o GetEcsSnapshotsResultOutput) SourceDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.SourceDiskType }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetEcsSnapshotsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetEcsSnapshotsResultOutput) Usage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEcsSnapshotsResult) *string { return v.Usage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEcsSnapshotsResultOutput{})
}
