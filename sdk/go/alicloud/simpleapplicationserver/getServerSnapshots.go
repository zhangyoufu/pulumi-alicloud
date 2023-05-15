// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simpleapplicationserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Simple Application Server Snapshots of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.143.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/simpleapplicationserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ids, err := simpleapplicationserver.GetServerSnapshots(ctx, &simpleapplicationserver.GetServerSnapshotsArgs{
//				Ids: []string{
//					"example_id",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("simpleApplicationServerSnapshotId1", ids.Snapshots[0].Id)
//			nameRegex, err := simpleapplicationserver.GetServerSnapshots(ctx, &simpleapplicationserver.GetServerSnapshotsArgs{
//				NameRegex: pulumi.StringRef("^my-Snapshot"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("simpleApplicationServerSnapshotId2", nameRegex.Snapshots[0].Id)
//			diskIdConf, err := simpleapplicationserver.GetServerSnapshots(ctx, &simpleapplicationserver.GetServerSnapshotsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				DiskId: pulumi.StringRef("example_value"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("simpleApplicationServerSnapshotId3", diskIdConf.Snapshots[0].Id)
//			instanceIdConf, err := simpleapplicationserver.GetServerSnapshots(ctx, &simpleapplicationserver.GetServerSnapshotsArgs{
//				Ids: []string{
//					"example_id",
//				},
//				InstanceId: pulumi.StringRef("example_value"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("simpleApplicationServerSnapshotId4", instanceIdConf.Snapshots[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetServerSnapshots(ctx *pulumi.Context, args *GetServerSnapshotsArgs, opts ...pulumi.InvokeOption) (*GetServerSnapshotsResult, error) {
	var rv GetServerSnapshotsResult
	err := ctx.Invoke("alicloud:simpleapplicationserver/getServerSnapshots:getServerSnapshots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerSnapshots.
type GetServerSnapshotsArgs struct {
	// The ID of the source disk. This parameter has a value even after the source disk is released.
	DiskId *string `pulumi:"diskId"`
	// A list of Snapshot IDs.
	Ids []string `pulumi:"ids"`
	// The ID of the simple application server.
	InstanceId *string `pulumi:"instanceId"`
	// A regex string to filter results by Snapshot name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of the snapshots. Valid values: `Progressing`, `Accomplished` and `Failed`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getServerSnapshots.
type GetServerSnapshotsResult struct {
	DiskId *string `pulumi:"diskId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                       `pulumi:"id"`
	Ids        []string                     `pulumi:"ids"`
	InstanceId *string                      `pulumi:"instanceId"`
	NameRegex  *string                      `pulumi:"nameRegex"`
	Names      []string                     `pulumi:"names"`
	OutputFile *string                      `pulumi:"outputFile"`
	Snapshots  []GetServerSnapshotsSnapshot `pulumi:"snapshots"`
	Status     *string                      `pulumi:"status"`
}

func GetServerSnapshotsOutput(ctx *pulumi.Context, args GetServerSnapshotsOutputArgs, opts ...pulumi.InvokeOption) GetServerSnapshotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerSnapshotsResult, error) {
			args := v.(GetServerSnapshotsArgs)
			r, err := GetServerSnapshots(ctx, &args, opts...)
			var s GetServerSnapshotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerSnapshotsResultOutput)
}

// A collection of arguments for invoking getServerSnapshots.
type GetServerSnapshotsOutputArgs struct {
	// The ID of the source disk. This parameter has a value even after the source disk is released.
	DiskId pulumi.StringPtrInput `pulumi:"diskId"`
	// A list of Snapshot IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ID of the simple application server.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// A regex string to filter results by Snapshot name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the snapshots. Valid values: `Progressing`, `Accomplished` and `Failed`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetServerSnapshotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerSnapshotsArgs)(nil)).Elem()
}

// A collection of values returned by getServerSnapshots.
type GetServerSnapshotsResultOutput struct{ *pulumi.OutputState }

func (GetServerSnapshotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerSnapshotsResult)(nil)).Elem()
}

func (o GetServerSnapshotsResultOutput) ToGetServerSnapshotsResultOutput() GetServerSnapshotsResultOutput {
	return o
}

func (o GetServerSnapshotsResultOutput) ToGetServerSnapshotsResultOutputWithContext(ctx context.Context) GetServerSnapshotsResultOutput {
	return o
}

func (o GetServerSnapshotsResultOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerSnapshotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServerSnapshotsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetServerSnapshotsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetServerSnapshotsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetServerSnapshotsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetServerSnapshotsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetServerSnapshotsResultOutput) Snapshots() GetServerSnapshotsSnapshotArrayOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) []GetServerSnapshotsSnapshot { return v.Snapshots }).(GetServerSnapshotsSnapshotArrayOutput)
}

func (o GetServerSnapshotsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerSnapshotsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerSnapshotsResultOutput{})
}
