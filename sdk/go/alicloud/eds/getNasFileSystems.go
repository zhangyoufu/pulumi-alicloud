// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecd Nas File Systems of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.141.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultSimpleOfficeSite, err := eds.NewSimpleOfficeSite(ctx, "defaultSimpleOfficeSite", &eds.SimpleOfficeSiteArgs{
//				CidrBlock:            pulumi.String("172.16.0.0/12"),
//				DesktopAccessType:    pulumi.String("Internet"),
//				OfficeSiteName:       pulumi.String("your_office_site_name"),
//				EnableInternetAccess: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			defaultNasFileSystem, err := eds.NewNasFileSystem(ctx, "defaultNasFileSystem", &eds.NasFileSystemArgs{
//				Description:       pulumi.String("your_description"),
//				OfficeSiteId:      defaultSimpleOfficeSite.ID(),
//				NasFileSystemName: pulumi.String("your_nas_file_system_name"),
//			})
//			if err != nil {
//				return err
//			}
//			ids, err := eds.GetNasFileSystems(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecdNasFileSystemId1", ids.Systems[0].Id)
//			nameRegex := defaultNasFileSystem.NasFileSystemName.ApplyT(func(nasFileSystemName *string) (eds.GetNasFileSystemsResult, error) {
//				return eds.GetNasFileSystemsOutput(ctx, eds.GetNasFileSystemsOutputArgs{
//					NameRegex: nasFileSystemName,
//				}, nil), nil
//			}).(eds.GetNasFileSystemsResultOutput)
//			ctx.Export("ecdNasFileSystemId2", nameRegex.ApplyT(func(nameRegex eds.GetNasFileSystemsResult) (*string, error) {
//				return &nameRegex.Systems[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
func GetNasFileSystems(ctx *pulumi.Context, args *GetNasFileSystemsArgs, opts ...pulumi.InvokeOption) (*GetNasFileSystemsResult, error) {
	var rv GetNasFileSystemsResult
	err := ctx.Invoke("alicloud:eds/getNasFileSystems:getNasFileSystems", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNasFileSystems.
type GetNasFileSystemsArgs struct {
	// A list of Nas File System IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Nas File System name.
	NameRegex *string `pulumi:"nameRegex"`
	// The ID of office site.
	OfficeSiteId *string `pulumi:"officeSiteId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getNasFileSystems.
type GetNasFileSystemsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string                    `pulumi:"id"`
	Ids          []string                  `pulumi:"ids"`
	NameRegex    *string                   `pulumi:"nameRegex"`
	Names        []string                  `pulumi:"names"`
	OfficeSiteId *string                   `pulumi:"officeSiteId"`
	OutputFile   *string                   `pulumi:"outputFile"`
	Status       *string                   `pulumi:"status"`
	Systems      []GetNasFileSystemsSystem `pulumi:"systems"`
}

func GetNasFileSystemsOutput(ctx *pulumi.Context, args GetNasFileSystemsOutputArgs, opts ...pulumi.InvokeOption) GetNasFileSystemsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNasFileSystemsResult, error) {
			args := v.(GetNasFileSystemsArgs)
			r, err := GetNasFileSystems(ctx, &args, opts...)
			var s GetNasFileSystemsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNasFileSystemsResultOutput)
}

// A collection of arguments for invoking getNasFileSystems.
type GetNasFileSystemsOutputArgs struct {
	// A list of Nas File System IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Nas File System name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The ID of office site.
	OfficeSiteId pulumi.StringPtrInput `pulumi:"officeSiteId"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of nas file system. Valid values: `Pending`, `Running`, `Stopped`,`Deleting`, `Deleted`, `Invalid`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetNasFileSystemsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNasFileSystemsArgs)(nil)).Elem()
}

// A collection of values returned by getNasFileSystems.
type GetNasFileSystemsResultOutput struct{ *pulumi.OutputState }

func (GetNasFileSystemsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNasFileSystemsResult)(nil)).Elem()
}

func (o GetNasFileSystemsResultOutput) ToGetNasFileSystemsResultOutput() GetNasFileSystemsResultOutput {
	return o
}

func (o GetNasFileSystemsResultOutput) ToGetNasFileSystemsResultOutputWithContext(ctx context.Context) GetNasFileSystemsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNasFileSystemsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNasFileSystemsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetNasFileSystemsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetNasFileSystemsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetNasFileSystemsResultOutput) OfficeSiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) *string { return v.OfficeSiteId }).(pulumi.StringPtrOutput)
}

func (o GetNasFileSystemsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetNasFileSystemsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetNasFileSystemsResultOutput) Systems() GetNasFileSystemsSystemArrayOutput {
	return o.ApplyT(func(v GetNasFileSystemsResult) []GetNasFileSystemsSystem { return v.Systems }).(GetNasFileSystemsSystemArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNasFileSystemsResultOutput{})
}
