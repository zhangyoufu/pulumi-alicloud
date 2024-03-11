// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Cloud Firewall Address Books of the current Alibaba Cloud user.
//
// > **NOTE:** Available since v1.178.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudfirewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			_, err := cloudfirewall.NewAddressBook(ctx, "default", &cloudfirewall.AddressBookArgs{
//				GroupName:     pulumi.String(name),
//				GroupType:     pulumi.String("ip"),
//				Description:   pulumi.String("tf-description"),
//				AutoAddTagEcs: pulumi.Int(0),
//				AddressLists: pulumi.StringArray{
//					pulumi.String("10.21.0.0/16"),
//					pulumi.String("10.168.0.0/16"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ids := cloudfirewall.GetAddressBooksOutput(ctx, cloudfirewall.GetAddressBooksOutputArgs{
//				Ids: pulumi.StringArray{
//					_default.ID(),
//				},
//			}, nil)
//			ctx.Export("cloudFirewallAddressBookId1", ids.ApplyT(func(ids cloudfirewall.GetAddressBooksResult) (*string, error) {
//				return &ids.Books[0].Id, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAddressBooks(ctx *pulumi.Context, args *GetAddressBooksArgs, opts ...pulumi.InvokeOption) (*GetAddressBooksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAddressBooksResult
	err := ctx.Invoke("alicloud:cloudfirewall/getAddressBooks:getAddressBooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddressBooks.
type GetAddressBooksArgs struct {
	// The type of the Address Book. Valid values: `ip`, `ipv6`, `domain`, `port`, `tag`.
	// **NOTE:** From version 1.213.1, `groupType` can be set to `ipv6`, `domain`, `port`.
	GroupType *string `pulumi:"groupType"`
	// A list of Address Book IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results Address Book name.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAddressBooks.
type GetAddressBooksResult struct {
	// A list of Cloud Firewall Address Books. Each element contains the following attributes:
	Books []GetAddressBooksBook `pulumi:"books"`
	// The type of the Address Book.
	GroupType *string `pulumi:"groupType"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// A list of Address Book names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}

func GetAddressBooksOutput(ctx *pulumi.Context, args GetAddressBooksOutputArgs, opts ...pulumi.InvokeOption) GetAddressBooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddressBooksResult, error) {
			args := v.(GetAddressBooksArgs)
			r, err := GetAddressBooks(ctx, &args, opts...)
			var s GetAddressBooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddressBooksResultOutput)
}

// A collection of arguments for invoking getAddressBooks.
type GetAddressBooksOutputArgs struct {
	// The type of the Address Book. Valid values: `ip`, `ipv6`, `domain`, `port`, `tag`.
	// **NOTE:** From version 1.213.1, `groupType` can be set to `ipv6`, `domain`, `port`.
	GroupType pulumi.StringPtrInput `pulumi:"groupType"`
	// A list of Address Book IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results Address Book name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results (after running `pulumi preview`).
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAddressBooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressBooksArgs)(nil)).Elem()
}

// A collection of values returned by getAddressBooks.
type GetAddressBooksResultOutput struct{ *pulumi.OutputState }

func (GetAddressBooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddressBooksResult)(nil)).Elem()
}

func (o GetAddressBooksResultOutput) ToGetAddressBooksResultOutput() GetAddressBooksResultOutput {
	return o
}

func (o GetAddressBooksResultOutput) ToGetAddressBooksResultOutputWithContext(ctx context.Context) GetAddressBooksResultOutput {
	return o
}

// A list of Cloud Firewall Address Books. Each element contains the following attributes:
func (o GetAddressBooksResultOutput) Books() GetAddressBooksBookArrayOutput {
	return o.ApplyT(func(v GetAddressBooksResult) []GetAddressBooksBook { return v.Books }).(GetAddressBooksBookArrayOutput)
}

// The type of the Address Book.
func (o GetAddressBooksResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressBooksResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddressBooksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddressBooksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAddressBooksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddressBooksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAddressBooksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressBooksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// A list of Address Book names.
func (o GetAddressBooksResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddressBooksResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAddressBooksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddressBooksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddressBooksResultOutput{})
}
