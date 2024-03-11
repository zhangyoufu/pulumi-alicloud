// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the PolarDB Global Database Networks of the current Alibaba Cloud user.
//
// > **NOTE:** Available since v1.181.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			this, err := polardb.GetNodeClasses(ctx, &polardb.GetNodeClassesArgs{
//				DbType:    pulumi.StringRef("MySQL"),
//				DbVersion: pulumi.StringRef("8.0"),
//				PayType:   "PostPaid",
//				Category:  pulumi.StringRef("Normal"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      *pulumi.String(this.Classes[0].ZoneId),
//				VswitchName: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			cluster, err := polardb.NewCluster(ctx, "cluster", &polardb.ClusterArgs{
//				DbType:      pulumi.String("MySQL"),
//				DbVersion:   pulumi.String("8.0"),
//				PayType:     pulumi.String("PostPaid"),
//				DbNodeCount: pulumi.Int(2),
//				DbNodeClass: *pulumi.String(this.Classes[0].SupportedEngines[0].AvailableResources[0].DbNodeClass),
//				VswitchId:   defaultSwitch.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultGlobalDatabaseNetwork, err := polardb.NewGlobalDatabaseNetwork(ctx, "defaultGlobalDatabaseNetwork", &polardb.GlobalDatabaseNetworkArgs{
//				DbClusterId: cluster.ID(),
//				Description: cluster.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ids := polardb.GetGlobalDatabaseNetworksOutput(ctx, polardb.GetGlobalDatabaseNetworksOutputArgs{
//				Ids: pulumi.StringArray{
//					defaultGlobalDatabaseNetwork.ID(),
//				},
//			}, nil)
//			ctx.Export("polardbGlobalDatabaseNetworkId1", ids.ApplyT(func(ids polardb.GetGlobalDatabaseNetworksResult) (*string, error) {
//				return &ids.Networks[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			description := polardb.GetGlobalDatabaseNetworksOutput(ctx, polardb.GetGlobalDatabaseNetworksOutputArgs{
//				Description: defaultGlobalDatabaseNetwork.Description,
//			}, nil)
//			ctx.Export("polardbGlobalDatabaseNetworkId2", description.ApplyT(func(description polardb.GetGlobalDatabaseNetworksResult) (*string, error) {
//				return &description.Networks[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetGlobalDatabaseNetworks(ctx *pulumi.Context, args *GetGlobalDatabaseNetworksArgs, opts ...pulumi.InvokeOption) (*GetGlobalDatabaseNetworksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGlobalDatabaseNetworksResult
	err := ctx.Invoke("alicloud:polardb/getGlobalDatabaseNetworks:getGlobalDatabaseNetworks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGlobalDatabaseNetworks.
type GetGlobalDatabaseNetworksArgs struct {
	// The ID of the cluster.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The description of the Global Database Network.
	Description *string `pulumi:"description"`
	// The ID of the Global Database Network.
	GdnId *string `pulumi:"gdnId"`
	// A list of Global Database Network IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
	PageNumber *int    `pulumi:"pageNumber"`
	PageSize   *int    `pulumi:"pageSize"`
	// The status of the Global Database Network. Valid values:
	Status *string `pulumi:"status"`
}

// A collection of values returned by getGlobalDatabaseNetworks.
type GetGlobalDatabaseNetworksResult struct {
	// The ID of the PolarDB cluster.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The description of the Global Database Network.
	Description *string `pulumi:"description"`
	// The ID of the Global Database Network.
	GdnId *string `pulumi:"gdnId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// A list of PolarDB Global Database Networks. Each element contains the following attributes:
	Networks   []GetGlobalDatabaseNetworksNetwork `pulumi:"networks"`
	OutputFile *string                            `pulumi:"outputFile"`
	PageNumber *int                               `pulumi:"pageNumber"`
	PageSize   *int                               `pulumi:"pageSize"`
	// The status of the Global Database Network.
	Status *string `pulumi:"status"`
}

func GetGlobalDatabaseNetworksOutput(ctx *pulumi.Context, args GetGlobalDatabaseNetworksOutputArgs, opts ...pulumi.InvokeOption) GetGlobalDatabaseNetworksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalDatabaseNetworksResult, error) {
			args := v.(GetGlobalDatabaseNetworksArgs)
			r, err := GetGlobalDatabaseNetworks(ctx, &args, opts...)
			var s GetGlobalDatabaseNetworksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalDatabaseNetworksResultOutput)
}

// A collection of arguments for invoking getGlobalDatabaseNetworks.
type GetGlobalDatabaseNetworksOutputArgs struct {
	// The ID of the cluster.
	DbClusterId pulumi.StringPtrInput `pulumi:"dbClusterId"`
	// The description of the Global Database Network.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The ID of the Global Database Network.
	GdnId pulumi.StringPtrInput `pulumi:"gdnId"`
	// A list of Global Database Network IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	PageNumber pulumi.IntPtrInput    `pulumi:"pageNumber"`
	PageSize   pulumi.IntPtrInput    `pulumi:"pageSize"`
	// The status of the Global Database Network. Valid values:
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetGlobalDatabaseNetworksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalDatabaseNetworksArgs)(nil)).Elem()
}

// A collection of values returned by getGlobalDatabaseNetworks.
type GetGlobalDatabaseNetworksResultOutput struct{ *pulumi.OutputState }

func (GetGlobalDatabaseNetworksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalDatabaseNetworksResult)(nil)).Elem()
}

func (o GetGlobalDatabaseNetworksResultOutput) ToGetGlobalDatabaseNetworksResultOutput() GetGlobalDatabaseNetworksResultOutput {
	return o
}

func (o GetGlobalDatabaseNetworksResultOutput) ToGetGlobalDatabaseNetworksResultOutputWithContext(ctx context.Context) GetGlobalDatabaseNetworksResultOutput {
	return o
}

// The ID of the PolarDB cluster.
func (o GetGlobalDatabaseNetworksResultOutput) DbClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *string { return v.DbClusterId }).(pulumi.StringPtrOutput)
}

// The description of the Global Database Network.
func (o GetGlobalDatabaseNetworksResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Global Database Network.
func (o GetGlobalDatabaseNetworksResultOutput) GdnId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *string { return v.GdnId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGlobalDatabaseNetworksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGlobalDatabaseNetworksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of PolarDB Global Database Networks. Each element contains the following attributes:
func (o GetGlobalDatabaseNetworksResultOutput) Networks() GetGlobalDatabaseNetworksNetworkArrayOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) []GetGlobalDatabaseNetworksNetwork { return v.Networks }).(GetGlobalDatabaseNetworksNetworkArrayOutput)
}

func (o GetGlobalDatabaseNetworksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetGlobalDatabaseNetworksResultOutput) PageNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *int { return v.PageNumber }).(pulumi.IntPtrOutput)
}

func (o GetGlobalDatabaseNetworksResultOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

// The status of the Global Database Network.
func (o GetGlobalDatabaseNetworksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGlobalDatabaseNetworksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalDatabaseNetworksResultOutput{})
}
