// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Nas Snapshots of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.152.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := nas.GetSnapshots(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nasSnapshotId1", ids.Snapshots[0].Id)
//			nameRegex, err := nas.GetSnapshots(ctx, &nas.GetSnapshotsArgs{
//				NameRegex: pulumi.StringRef("^my-Snapshot"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nasSnapshotId2", nameRegex.Snapshots[0].Id)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetSnapshots(ctx *pulumi.Context, args *GetSnapshotsArgs, opts ...pulumi.InvokeOption) (*GetSnapshotsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSnapshotsResult
	err := ctx.Invoke("alicloud:nas/getSnapshots:getSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsArgs struct {
	// The ID of the file system.
	FileSystemId *string `pulumi:"fileSystemId"`
	// A list of Snapshot IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Snapshot name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// The status of the snapshot.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getSnapshots.
type GetSnapshotsResult struct {
	FileSystemId *string `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id           string                 `pulumi:"id"`
	Ids          []string               `pulumi:"ids"`
	NameRegex    *string                `pulumi:"nameRegex"`
	Names        []string               `pulumi:"names"`
	OutputFile   *string                `pulumi:"outputFile"`
	SnapshotName *string                `pulumi:"snapshotName"`
	Snapshots    []GetSnapshotsSnapshot `pulumi:"snapshots"`
	Status       *string                `pulumi:"status"`
}

func GetSnapshotsOutput(ctx *pulumi.Context, args GetSnapshotsOutputArgs, opts ...pulumi.InvokeOption) GetSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSnapshotsResult, error) {
			args := v.(GetSnapshotsArgs)
			r, err := GetSnapshots(ctx, &args, opts...)
			var s GetSnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSnapshotsResultOutput)
}

// A collection of arguments for invoking getSnapshots.
type GetSnapshotsOutputArgs struct {
	// The ID of the file system.
	FileSystemId pulumi.StringPtrInput `pulumi:"fileSystemId"`
	// A list of Snapshot IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Snapshot name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the snapshot.
	SnapshotName pulumi.StringPtrInput `pulumi:"snapshotName"`
	// The status of the snapshot.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshots.
type GetSnapshotsResultOutput struct{ *pulumi.OutputState }

func (GetSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotsResult)(nil)).Elem()
}

func (o GetSnapshotsResultOutput) ToGetSnapshotsResultOutput() GetSnapshotsResultOutput {
	return o
}

func (o GetSnapshotsResultOutput) ToGetSnapshotsResultOutputWithContext(ctx context.Context) GetSnapshotsResultOutput {
	return o
}

func (o GetSnapshotsResultOutput) FileSystemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.FileSystemId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSnapshotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSnapshotsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSnapshotsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSnapshotsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetSnapshotsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSnapshotsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetSnapshotsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSnapshotsResultOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

func (o GetSnapshotsResultOutput) Snapshots() GetSnapshotsSnapshotArrayOutput {
	return o.ApplyT(func(v GetSnapshotsResult) []GetSnapshotsSnapshot { return v.Snapshots }).(GetSnapshotsSnapshotArrayOutput)
}

func (o GetSnapshotsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotsResultOutput{})
}
