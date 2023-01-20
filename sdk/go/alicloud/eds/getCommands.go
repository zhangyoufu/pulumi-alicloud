// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Ecd Commands of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.146.0+.
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
//				CidrBlock:         pulumi.String("172.16.0.0/12"),
//				DesktopAccessType: pulumi.String("Internet"),
//				OfficeSiteName:    pulumi.String("your_office_site_name"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultBundles, err := eds.GetBundles(ctx, &eds.GetBundlesArgs{
//				BundleType: pulumi.StringRef("SYSTEM"),
//				NameRegex:  pulumi.StringRef("windows"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultEcdPolicyGroup, err := eds.NewEcdPolicyGroup(ctx, "defaultEcdPolicyGroup", &eds.EcdPolicyGroupArgs{
//				PolicyGroupName: pulumi.String("your_policy_group_name"),
//				Clipboard:       pulumi.String("readwrite"),
//				LocalDrive:      pulumi.String("read"),
//				AuthorizeAccessPolicyRules: eds.EcdPolicyGroupAuthorizeAccessPolicyRuleArray{
//					&eds.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs{
//						Description: pulumi.String("example_value"),
//						CidrIp:      pulumi.String("1.2.3.4/24"),
//					},
//				},
//				AuthorizeSecurityPolicyRules: eds.EcdPolicyGroupAuthorizeSecurityPolicyRuleArray{
//					&eds.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs{
//						Type:        pulumi.String("inflow"),
//						Policy:      pulumi.String("accept"),
//						Description: pulumi.String("example_value"),
//						PortRange:   pulumi.String("80/80"),
//						IpProtocol:  pulumi.String("TCP"),
//						Priority:    pulumi.String("1"),
//						CidrIp:      pulumi.String("0.0.0.0/0"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultDesktop, err := eds.NewDesktop(ctx, "defaultDesktop", &eds.DesktopArgs{
//				OfficeSiteId:  defaultSimpleOfficeSite.ID(),
//				PolicyGroupId: defaultEcdPolicyGroup.ID(),
//				BundleId:      *pulumi.String(defaultBundles.Bundles[0].Id),
//				DesktopName:   pulumi.Any(_var.Name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eds.NewCommand(ctx, "defaultCommand", &eds.CommandArgs{
//				CommandContent: pulumi.String("ipconfig"),
//				CommandType:    pulumi.String("RunPowerShellScript"),
//				DesktopId:      defaultDesktop.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ids, err := eds.GetCommands(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ecdCommandId1", ids.Commands[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetCommands(ctx *pulumi.Context, args *GetCommandsArgs, opts ...pulumi.InvokeOption) (*GetCommandsResult, error) {
	var rv GetCommandsResult
	err := ctx.Invoke("alicloud:eds/getCommands:getCommands", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCommands.
type GetCommandsArgs struct {
	// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
	CommandType *string `pulumi:"commandType"`
	// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The desktop id of the Desktop.
	DesktopId *string `pulumi:"desktopId"`
	// A list of Command IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getCommands.
type GetCommandsResult struct {
	CommandType     *string              `pulumi:"commandType"`
	Commands        []GetCommandsCommand `pulumi:"commands"`
	ContentEncoding *string              `pulumi:"contentEncoding"`
	DesktopId       *string              `pulumi:"desktopId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	Status     *string  `pulumi:"status"`
}

func GetCommandsOutput(ctx *pulumi.Context, args GetCommandsOutputArgs, opts ...pulumi.InvokeOption) GetCommandsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCommandsResult, error) {
			args := v.(GetCommandsArgs)
			r, err := GetCommands(ctx, &args, opts...)
			var s GetCommandsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCommandsResultOutput)
}

// A collection of arguments for invoking getCommands.
type GetCommandsOutputArgs struct {
	// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
	CommandType pulumi.StringPtrInput `pulumi:"commandType"`
	// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
	ContentEncoding pulumi.StringPtrInput `pulumi:"contentEncoding"`
	// The desktop id of the Desktop.
	DesktopId pulumi.StringPtrInput `pulumi:"desktopId"`
	// A list of Command IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetCommandsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCommandsArgs)(nil)).Elem()
}

// A collection of values returned by getCommands.
type GetCommandsResultOutput struct{ *pulumi.OutputState }

func (GetCommandsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCommandsResult)(nil)).Elem()
}

func (o GetCommandsResultOutput) ToGetCommandsResultOutput() GetCommandsResultOutput {
	return o
}

func (o GetCommandsResultOutput) ToGetCommandsResultOutputWithContext(ctx context.Context) GetCommandsResultOutput {
	return o
}

func (o GetCommandsResultOutput) CommandType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommandsResult) *string { return v.CommandType }).(pulumi.StringPtrOutput)
}

func (o GetCommandsResultOutput) Commands() GetCommandsCommandArrayOutput {
	return o.ApplyT(func(v GetCommandsResult) []GetCommandsCommand { return v.Commands }).(GetCommandsCommandArrayOutput)
}

func (o GetCommandsResultOutput) ContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommandsResult) *string { return v.ContentEncoding }).(pulumi.StringPtrOutput)
}

func (o GetCommandsResultOutput) DesktopId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommandsResult) *string { return v.DesktopId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCommandsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCommandsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCommandsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCommandsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCommandsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommandsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCommandsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCommandsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCommandsResultOutput{})
}
