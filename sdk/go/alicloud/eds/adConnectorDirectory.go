// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECD Ad Connector Directory resource.
//
// For information about ECD Ad Connector Directory and how to use it, see [What is Ad Connector Directory](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createadconnectordirectory).
//
// > **NOTE:** Available since v1.174.0.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
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
//			_default, err := eds.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      pulumi.String(_default.Ids[0]),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInteger, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Min: 10000,
//				Max: 99999,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eds.NewAdConnectorDirectory(ctx, "default", &eds.AdConnectorDirectoryArgs{
//				DirectoryName:     pulumi.String(fmt.Sprintf("%v-%v", name, defaultInteger.Result)),
//				DesktopAccessType: pulumi.String("INTERNET"),
//				DnsAddresses: pulumi.StringArray{
//					pulumi.String("127.0.0.2"),
//				},
//				DomainName:        pulumi.String("corp.example.com"),
//				DomainPassword:    pulumi.String("Example1234"),
//				DomainUserName:    pulumi.String("sAMAccountName"),
//				EnableAdminAccess: pulumi.Bool(false),
//				MfaEnabled:        pulumi.Bool(false),
//				Specification:     pulumi.Int(1),
//				SubDomainDnsAddresses: pulumi.StringArray{
//					pulumi.String("127.0.0.3"),
//				},
//				SubDomainName: pulumi.String("child.example.com"),
//				VswitchIds: pulumi.StringArray{
//					defaultSwitch.ID(),
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
// ECD Ad Connector Directory can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:eds/adConnectorDirectory:AdConnectorDirectory example <id>
// ```
type AdConnectorDirectory struct {
	pulumi.CustomResourceState

	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType pulumi.StringOutput `pulumi:"desktopAccessType"`
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	DirectoryName pulumi.StringOutput `pulumi:"directoryName"`
	// The DNS address list.
	DnsAddresses pulumi.StringArrayOutput `pulumi:"dnsAddresses"`
	// The name of the domain.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The user password of the domain administrator. The maximum number of English characters is 64.
	DomainPassword pulumi.StringOutput `pulumi:"domainPassword"`
	// The username of the domain administrator. The maximum number of English characters is 64.
	DomainUserName pulumi.StringOutput `pulumi:"domainUserName"`
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableAdminAccess pulumi.BoolOutput `pulumi:"enableAdminAccess"`
	// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
	MfaEnabled pulumi.BoolOutput `pulumi:"mfaEnabled"`
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification pulumi.IntPtrOutput `pulumi:"specification"`
	// The status of directory.
	Status pulumi.StringOutput `pulumi:"status"`
	// The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
	SubDomainDnsAddresses pulumi.StringArrayOutput `pulumi:"subDomainDnsAddresses"`
	// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
	SubDomainName pulumi.StringPtrOutput `pulumi:"subDomainName"`
	// List of VSwitch IDs in the directory.
	VswitchIds pulumi.StringArrayOutput `pulumi:"vswitchIds"`
}

// NewAdConnectorDirectory registers a new resource with the given unique name, arguments, and options.
func NewAdConnectorDirectory(ctx *pulumi.Context,
	name string, args *AdConnectorDirectoryArgs, opts ...pulumi.ResourceOption) (*AdConnectorDirectory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryName == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryName'")
	}
	if args.DnsAddresses == nil {
		return nil, errors.New("invalid value for required argument 'DnsAddresses'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.DomainPassword == nil {
		return nil, errors.New("invalid value for required argument 'DomainPassword'")
	}
	if args.DomainUserName == nil {
		return nil, errors.New("invalid value for required argument 'DomainUserName'")
	}
	if args.VswitchIds == nil {
		return nil, errors.New("invalid value for required argument 'VswitchIds'")
	}
	if args.DomainPassword != nil {
		args.DomainPassword = pulumi.ToSecret(args.DomainPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"domainPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AdConnectorDirectory
	err := ctx.RegisterResource("alicloud:eds/adConnectorDirectory:AdConnectorDirectory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdConnectorDirectory gets an existing AdConnectorDirectory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdConnectorDirectory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdConnectorDirectoryState, opts ...pulumi.ResourceOption) (*AdConnectorDirectory, error) {
	var resource AdConnectorDirectory
	err := ctx.ReadResource("alicloud:eds/adConnectorDirectory:AdConnectorDirectory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdConnectorDirectory resources.
type adConnectorDirectoryState struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType *string `pulumi:"desktopAccessType"`
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	DirectoryName *string `pulumi:"directoryName"`
	// The DNS address list.
	DnsAddresses []string `pulumi:"dnsAddresses"`
	// The name of the domain.
	DomainName *string `pulumi:"domainName"`
	// The user password of the domain administrator. The maximum number of English characters is 64.
	DomainPassword *string `pulumi:"domainPassword"`
	// The username of the domain administrator. The maximum number of English characters is 64.
	DomainUserName *string `pulumi:"domainUserName"`
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableAdminAccess *bool `pulumi:"enableAdminAccess"`
	// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
	MfaEnabled *bool `pulumi:"mfaEnabled"`
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification *int `pulumi:"specification"`
	// The status of directory.
	Status *string `pulumi:"status"`
	// The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
	SubDomainDnsAddresses []string `pulumi:"subDomainDnsAddresses"`
	// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
	SubDomainName *string `pulumi:"subDomainName"`
	// List of VSwitch IDs in the directory.
	VswitchIds []string `pulumi:"vswitchIds"`
}

type AdConnectorDirectoryState struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType pulumi.StringPtrInput
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	DirectoryName pulumi.StringPtrInput
	// The DNS address list.
	DnsAddresses pulumi.StringArrayInput
	// The name of the domain.
	DomainName pulumi.StringPtrInput
	// The user password of the domain administrator. The maximum number of English characters is 64.
	DomainPassword pulumi.StringPtrInput
	// The username of the domain administrator. The maximum number of English characters is 64.
	DomainUserName pulumi.StringPtrInput
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableAdminAccess pulumi.BoolPtrInput
	// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
	MfaEnabled pulumi.BoolPtrInput
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification pulumi.IntPtrInput
	// The status of directory.
	Status pulumi.StringPtrInput
	// The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
	SubDomainDnsAddresses pulumi.StringArrayInput
	// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
	SubDomainName pulumi.StringPtrInput
	// List of VSwitch IDs in the directory.
	VswitchIds pulumi.StringArrayInput
}

func (AdConnectorDirectoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*adConnectorDirectoryState)(nil)).Elem()
}

type adConnectorDirectoryArgs struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType *string `pulumi:"desktopAccessType"`
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	DirectoryName string `pulumi:"directoryName"`
	// The DNS address list.
	DnsAddresses []string `pulumi:"dnsAddresses"`
	// The name of the domain.
	DomainName string `pulumi:"domainName"`
	// The user password of the domain administrator. The maximum number of English characters is 64.
	DomainPassword string `pulumi:"domainPassword"`
	// The username of the domain administrator. The maximum number of English characters is 64.
	DomainUserName string `pulumi:"domainUserName"`
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableAdminAccess *bool `pulumi:"enableAdminAccess"`
	// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
	MfaEnabled *bool `pulumi:"mfaEnabled"`
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification *int `pulumi:"specification"`
	// The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
	SubDomainDnsAddresses []string `pulumi:"subDomainDnsAddresses"`
	// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
	SubDomainName *string `pulumi:"subDomainName"`
	// List of VSwitch IDs in the directory.
	VswitchIds []string `pulumi:"vswitchIds"`
}

// The set of arguments for constructing a AdConnectorDirectory resource.
type AdConnectorDirectoryArgs struct {
	// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
	DesktopAccessType pulumi.StringPtrInput
	// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	DirectoryName pulumi.StringInput
	// The DNS address list.
	DnsAddresses pulumi.StringArrayInput
	// The name of the domain.
	DomainName pulumi.StringInput
	// The user password of the domain administrator. The maximum number of English characters is 64.
	DomainPassword pulumi.StringInput
	// The username of the domain administrator. The maximum number of English characters is 64.
	DomainUserName pulumi.StringInput
	// Whether to grant local administrator rights to users who use cloud desktops.
	EnableAdminAccess pulumi.BoolPtrInput
	// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
	MfaEnabled pulumi.BoolPtrInput
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification pulumi.IntPtrInput
	// The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
	SubDomainDnsAddresses pulumi.StringArrayInput
	// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
	SubDomainName pulumi.StringPtrInput
	// List of VSwitch IDs in the directory.
	VswitchIds pulumi.StringArrayInput
}

func (AdConnectorDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adConnectorDirectoryArgs)(nil)).Elem()
}

type AdConnectorDirectoryInput interface {
	pulumi.Input

	ToAdConnectorDirectoryOutput() AdConnectorDirectoryOutput
	ToAdConnectorDirectoryOutputWithContext(ctx context.Context) AdConnectorDirectoryOutput
}

func (*AdConnectorDirectory) ElementType() reflect.Type {
	return reflect.TypeOf((**AdConnectorDirectory)(nil)).Elem()
}

func (i *AdConnectorDirectory) ToAdConnectorDirectoryOutput() AdConnectorDirectoryOutput {
	return i.ToAdConnectorDirectoryOutputWithContext(context.Background())
}

func (i *AdConnectorDirectory) ToAdConnectorDirectoryOutputWithContext(ctx context.Context) AdConnectorDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdConnectorDirectoryOutput)
}

// AdConnectorDirectoryArrayInput is an input type that accepts AdConnectorDirectoryArray and AdConnectorDirectoryArrayOutput values.
// You can construct a concrete instance of `AdConnectorDirectoryArrayInput` via:
//
//	AdConnectorDirectoryArray{ AdConnectorDirectoryArgs{...} }
type AdConnectorDirectoryArrayInput interface {
	pulumi.Input

	ToAdConnectorDirectoryArrayOutput() AdConnectorDirectoryArrayOutput
	ToAdConnectorDirectoryArrayOutputWithContext(context.Context) AdConnectorDirectoryArrayOutput
}

type AdConnectorDirectoryArray []AdConnectorDirectoryInput

func (AdConnectorDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdConnectorDirectory)(nil)).Elem()
}

func (i AdConnectorDirectoryArray) ToAdConnectorDirectoryArrayOutput() AdConnectorDirectoryArrayOutput {
	return i.ToAdConnectorDirectoryArrayOutputWithContext(context.Background())
}

func (i AdConnectorDirectoryArray) ToAdConnectorDirectoryArrayOutputWithContext(ctx context.Context) AdConnectorDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdConnectorDirectoryArrayOutput)
}

// AdConnectorDirectoryMapInput is an input type that accepts AdConnectorDirectoryMap and AdConnectorDirectoryMapOutput values.
// You can construct a concrete instance of `AdConnectorDirectoryMapInput` via:
//
//	AdConnectorDirectoryMap{ "key": AdConnectorDirectoryArgs{...} }
type AdConnectorDirectoryMapInput interface {
	pulumi.Input

	ToAdConnectorDirectoryMapOutput() AdConnectorDirectoryMapOutput
	ToAdConnectorDirectoryMapOutputWithContext(context.Context) AdConnectorDirectoryMapOutput
}

type AdConnectorDirectoryMap map[string]AdConnectorDirectoryInput

func (AdConnectorDirectoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdConnectorDirectory)(nil)).Elem()
}

func (i AdConnectorDirectoryMap) ToAdConnectorDirectoryMapOutput() AdConnectorDirectoryMapOutput {
	return i.ToAdConnectorDirectoryMapOutputWithContext(context.Background())
}

func (i AdConnectorDirectoryMap) ToAdConnectorDirectoryMapOutputWithContext(ctx context.Context) AdConnectorDirectoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdConnectorDirectoryMapOutput)
}

type AdConnectorDirectoryOutput struct{ *pulumi.OutputState }

func (AdConnectorDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdConnectorDirectory)(nil)).Elem()
}

func (o AdConnectorDirectoryOutput) ToAdConnectorDirectoryOutput() AdConnectorDirectoryOutput {
	return o
}

func (o AdConnectorDirectoryOutput) ToAdConnectorDirectoryOutputWithContext(ctx context.Context) AdConnectorDirectoryOutput {
	return o
}

// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
func (o AdConnectorDirectoryOutput) DesktopAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringOutput { return v.DesktopAccessType }).(pulumi.StringOutput)
}

// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
func (o AdConnectorDirectoryOutput) DirectoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringOutput { return v.DirectoryName }).(pulumi.StringOutput)
}

// The DNS address list.
func (o AdConnectorDirectoryOutput) DnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringArrayOutput { return v.DnsAddresses }).(pulumi.StringArrayOutput)
}

// The name of the domain.
func (o AdConnectorDirectoryOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The user password of the domain administrator. The maximum number of English characters is 64.
func (o AdConnectorDirectoryOutput) DomainPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringOutput { return v.DomainPassword }).(pulumi.StringOutput)
}

// The username of the domain administrator. The maximum number of English characters is 64.
func (o AdConnectorDirectoryOutput) DomainUserName() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringOutput { return v.DomainUserName }).(pulumi.StringOutput)
}

// Whether to grant local administrator rights to users who use cloud desktops.
func (o AdConnectorDirectoryOutput) EnableAdminAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.BoolOutput { return v.EnableAdminAccess }).(pulumi.BoolOutput)
}

// Whether MFA authentication is enabled. After all AD users in this directory log on to the cloud desktop, enter the correct password and then enter the dynamic verification code generated by the MFA device. **NOTE:** The MFA device needs to be bound when logging in for the first time.
func (o AdConnectorDirectoryOutput) MfaEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.BoolOutput { return v.MfaEnabled }).(pulumi.BoolOutput)
}

// The AD Connector specifications. Valid values: `1`, `2`.
func (o AdConnectorDirectoryOutput) Specification() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.IntPtrOutput { return v.Specification }).(pulumi.IntPtrOutput)
}

// The status of directory.
func (o AdConnectorDirectoryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The Enterprise already has the DNS address of the AD subdomain. If `subDomainName` is set and this parameter is not set, the child Domain DNS is considered consistent with the parent domain.
func (o AdConnectorDirectoryOutput) SubDomainDnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringArrayOutput { return v.SubDomainDnsAddresses }).(pulumi.StringArrayOutput)
}

// The Enterprise already has a fully qualified domain name (FQDN) of an AD subdomain, with both a host name and a domain name.
func (o AdConnectorDirectoryOutput) SubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringPtrOutput { return v.SubDomainName }).(pulumi.StringPtrOutput)
}

// List of VSwitch IDs in the directory.
func (o AdConnectorDirectoryOutput) VswitchIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdConnectorDirectory) pulumi.StringArrayOutput { return v.VswitchIds }).(pulumi.StringArrayOutput)
}

type AdConnectorDirectoryArrayOutput struct{ *pulumi.OutputState }

func (AdConnectorDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdConnectorDirectory)(nil)).Elem()
}

func (o AdConnectorDirectoryArrayOutput) ToAdConnectorDirectoryArrayOutput() AdConnectorDirectoryArrayOutput {
	return o
}

func (o AdConnectorDirectoryArrayOutput) ToAdConnectorDirectoryArrayOutputWithContext(ctx context.Context) AdConnectorDirectoryArrayOutput {
	return o
}

func (o AdConnectorDirectoryArrayOutput) Index(i pulumi.IntInput) AdConnectorDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdConnectorDirectory {
		return vs[0].([]*AdConnectorDirectory)[vs[1].(int)]
	}).(AdConnectorDirectoryOutput)
}

type AdConnectorDirectoryMapOutput struct{ *pulumi.OutputState }

func (AdConnectorDirectoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdConnectorDirectory)(nil)).Elem()
}

func (o AdConnectorDirectoryMapOutput) ToAdConnectorDirectoryMapOutput() AdConnectorDirectoryMapOutput {
	return o
}

func (o AdConnectorDirectoryMapOutput) ToAdConnectorDirectoryMapOutputWithContext(ctx context.Context) AdConnectorDirectoryMapOutput {
	return o
}

func (o AdConnectorDirectoryMapOutput) MapIndex(k pulumi.StringInput) AdConnectorDirectoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdConnectorDirectory {
		return vs[0].(map[string]*AdConnectorDirectory)[vs[1].(string)]
	}).(AdConnectorDirectoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdConnectorDirectoryInput)(nil)).Elem(), &AdConnectorDirectory{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdConnectorDirectoryArrayInput)(nil)).Elem(), AdConnectorDirectoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdConnectorDirectoryMapInput)(nil)).Elem(), AdConnectorDirectoryMap{})
	pulumi.RegisterOutputType(AdConnectorDirectoryOutput{})
	pulumi.RegisterOutputType(AdConnectorDirectoryArrayOutput{})
	pulumi.RegisterOutputType(AdConnectorDirectoryMapOutput{})
}
