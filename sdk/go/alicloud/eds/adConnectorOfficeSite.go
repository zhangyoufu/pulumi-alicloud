// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECD Ad Connector Office Site resource.
//
// For information about ECD Ad Connector Office Site and how to use it, see [What is Ad Connector Office Site](https://www.alibabacloud.com/help/en/elastic-desktop-service/latest/createadconnectorofficesite).
//
// > **NOTE:** Available in v1.176.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cen"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultInstance, err := cen.NewInstance(ctx, "defaultInstance", &cen.InstanceArgs{
// 			CenInstanceName: pulumi.Any(_var.Name),
// 			ProtectionLevel: pulumi.String("REDUCED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eds.NewAdConnectorOfficeSite(ctx, "defaultAdConnectorOfficeSite", &eds.AdConnectorOfficeSiteArgs{
// 			AdConnectorOfficeSiteName: pulumi.Any(_var.Name),
// 			Bandwidth:                 pulumi.Int(100),
// 			CenId:                     defaultInstance.ID(),
// 			CidrBlock:                 pulumi.String("10.0.0.0/12"),
// 			DesktopAccessType:         pulumi.String("INTERNET"),
// 			DnsAddresses: pulumi.StringArray{
// 				pulumi.String("127.0.0.2"),
// 			},
// 			DomainName:           pulumi.String("example1234.com"),
// 			DomainPassword:       pulumi.String("YourPassword1234"),
// 			DomainUserName:       pulumi.String("Administrator"),
// 			EnableAdminAccess:    pulumi.Bool(true),
// 			EnableInternetAccess: pulumi.Bool(true),
// 			MfaEnabled:           pulumi.Bool(false),
// 			SubDomainDnsAddresses: pulumi.StringArray{
// 				pulumi.String("127.0.0.3"),
// 			},
// 			SubDomainName: pulumi.String("child.example1234.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ECD Ad Connector Office Site can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:eds/adConnectorOfficeSite:AdConnectorOfficeSite example <id>
// ```
type AdConnectorOfficeSite struct {
	pulumi.CustomResourceState

	// The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	AdConnectorOfficeSiteName pulumi.StringOutput `pulumi:"adConnectorOfficeSiteName"`
	// The ad hostname.
	AdHostname pulumi.StringPtrOutput `pulumi:"adHostname"`
	// The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// The ID of the CEN instance.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The cen owner id.
	CenOwnerId pulumi.StringPtrOutput `pulumi:"cenOwnerId"`
	// Workspace Corresponds to the Security Office Network of IPv4 Segment.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
	// - `INTERNET`: connects clients to cloud desktops only over the Internet.
	// - `VPC`: connects clients to cloud desktops only over a VPC.
	// - `ANY`: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method when you use a client to connect to the cloud desktop.
	DesktopAccessType pulumi.StringOutput `pulumi:"desktopAccessType"`
	// The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
	DnsAddresses pulumi.StringArrayOutput `pulumi:"dnsAddresses"`
	// The domain name of the enterprise AD system. You can register each domain name only once.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The password of the domain administrator. The password can be up to 64 characters in length.
	DomainPassword pulumi.StringPtrOutput `pulumi:"domainPassword"`
	// The username of the domain administrator. The username can be up to 64 characters in length.
	DomainUserName pulumi.StringPtrOutput `pulumi:"domainUserName"`
	// Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
	EnableAdminAccess pulumi.BoolOutput `pulumi:"enableAdminAccess"`
	// Specifies whether to enable Internet access.
	EnableInternetAccess pulumi.BoolOutput `pulumi:"enableInternetAccess"`
	// Specifies whether to enable multi-factor authentication (MFA).
	MfaEnabled pulumi.BoolPtrOutput `pulumi:"mfaEnabled"`
	// The protocol type. Valid values: `ASP`, `HDX`.
	ProtocolType pulumi.StringPtrOutput `pulumi:"protocolType"`
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification pulumi.IntPtrOutput `pulumi:"specification"`
	// The resource State.
	Status pulumi.StringOutput `pulumi:"status"`
	// The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
	SubDomainDnsAddresses pulumi.StringArrayOutput `pulumi:"subDomainDnsAddresses"`
	// The domain name of the enterprise AD subdomain.
	SubDomainName pulumi.StringPtrOutput `pulumi:"subDomainName"`
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
	VerifyCode pulumi.StringPtrOutput `pulumi:"verifyCode"`
}

// NewAdConnectorOfficeSite registers a new resource with the given unique name, arguments, and options.
func NewAdConnectorOfficeSite(ctx *pulumi.Context,
	name string, args *AdConnectorOfficeSiteArgs, opts ...pulumi.ResourceOption) (*AdConnectorOfficeSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdConnectorOfficeSiteName == nil {
		return nil, errors.New("invalid value for required argument 'AdConnectorOfficeSiteName'")
	}
	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.DnsAddresses == nil {
		return nil, errors.New("invalid value for required argument 'DnsAddresses'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource AdConnectorOfficeSite
	err := ctx.RegisterResource("alicloud:eds/adConnectorOfficeSite:AdConnectorOfficeSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdConnectorOfficeSite gets an existing AdConnectorOfficeSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdConnectorOfficeSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdConnectorOfficeSiteState, opts ...pulumi.ResourceOption) (*AdConnectorOfficeSite, error) {
	var resource AdConnectorOfficeSite
	err := ctx.ReadResource("alicloud:eds/adConnectorOfficeSite:AdConnectorOfficeSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdConnectorOfficeSite resources.
type adConnectorOfficeSiteState struct {
	// The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	AdConnectorOfficeSiteName *string `pulumi:"adConnectorOfficeSiteName"`
	// The ad hostname.
	AdHostname *string `pulumi:"adHostname"`
	// The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
	Bandwidth *int `pulumi:"bandwidth"`
	// The ID of the CEN instance.
	CenId *string `pulumi:"cenId"`
	// The cen owner id.
	CenOwnerId *string `pulumi:"cenOwnerId"`
	// Workspace Corresponds to the Security Office Network of IPv4 Segment.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
	// - `INTERNET`: connects clients to cloud desktops only over the Internet.
	// - `VPC`: connects clients to cloud desktops only over a VPC.
	// - `ANY`: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method when you use a client to connect to the cloud desktop.
	DesktopAccessType *string `pulumi:"desktopAccessType"`
	// The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
	DnsAddresses []string `pulumi:"dnsAddresses"`
	// The domain name of the enterprise AD system. You can register each domain name only once.
	DomainName *string `pulumi:"domainName"`
	// The password of the domain administrator. The password can be up to 64 characters in length.
	DomainPassword *string `pulumi:"domainPassword"`
	// The username of the domain administrator. The username can be up to 64 characters in length.
	DomainUserName *string `pulumi:"domainUserName"`
	// Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
	EnableAdminAccess *bool `pulumi:"enableAdminAccess"`
	// Specifies whether to enable Internet access.
	EnableInternetAccess *bool `pulumi:"enableInternetAccess"`
	// Specifies whether to enable multi-factor authentication (MFA).
	MfaEnabled *bool `pulumi:"mfaEnabled"`
	// The protocol type. Valid values: `ASP`, `HDX`.
	ProtocolType *string `pulumi:"protocolType"`
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification *int `pulumi:"specification"`
	// The resource State.
	Status *string `pulumi:"status"`
	// The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
	SubDomainDnsAddresses []string `pulumi:"subDomainDnsAddresses"`
	// The domain name of the enterprise AD subdomain.
	SubDomainName *string `pulumi:"subDomainName"`
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
	VerifyCode *string `pulumi:"verifyCode"`
}

type AdConnectorOfficeSiteState struct {
	// The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	AdConnectorOfficeSiteName pulumi.StringPtrInput
	// The ad hostname.
	AdHostname pulumi.StringPtrInput
	// The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
	Bandwidth pulumi.IntPtrInput
	// The ID of the CEN instance.
	CenId pulumi.StringPtrInput
	// The cen owner id.
	CenOwnerId pulumi.StringPtrInput
	// Workspace Corresponds to the Security Office Network of IPv4 Segment.
	CidrBlock pulumi.StringPtrInput
	// The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
	// - `INTERNET`: connects clients to cloud desktops only over the Internet.
	// - `VPC`: connects clients to cloud desktops only over a VPC.
	// - `ANY`: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method when you use a client to connect to the cloud desktop.
	DesktopAccessType pulumi.StringPtrInput
	// The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
	DnsAddresses pulumi.StringArrayInput
	// The domain name of the enterprise AD system. You can register each domain name only once.
	DomainName pulumi.StringPtrInput
	// The password of the domain administrator. The password can be up to 64 characters in length.
	DomainPassword pulumi.StringPtrInput
	// The username of the domain administrator. The username can be up to 64 characters in length.
	DomainUserName pulumi.StringPtrInput
	// Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
	EnableAdminAccess pulumi.BoolPtrInput
	// Specifies whether to enable Internet access.
	EnableInternetAccess pulumi.BoolPtrInput
	// Specifies whether to enable multi-factor authentication (MFA).
	MfaEnabled pulumi.BoolPtrInput
	// The protocol type. Valid values: `ASP`, `HDX`.
	ProtocolType pulumi.StringPtrInput
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification pulumi.IntPtrInput
	// The resource State.
	Status pulumi.StringPtrInput
	// The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
	SubDomainDnsAddresses pulumi.StringArrayInput
	// The domain name of the enterprise AD subdomain.
	SubDomainName pulumi.StringPtrInput
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
	VerifyCode pulumi.StringPtrInput
}

func (AdConnectorOfficeSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*adConnectorOfficeSiteState)(nil)).Elem()
}

type adConnectorOfficeSiteArgs struct {
	// The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	AdConnectorOfficeSiteName string `pulumi:"adConnectorOfficeSiteName"`
	// The ad hostname.
	AdHostname *string `pulumi:"adHostname"`
	// The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
	Bandwidth *int `pulumi:"bandwidth"`
	// The ID of the CEN instance.
	CenId string `pulumi:"cenId"`
	// The cen owner id.
	CenOwnerId *string `pulumi:"cenOwnerId"`
	// Workspace Corresponds to the Security Office Network of IPv4 Segment.
	CidrBlock string `pulumi:"cidrBlock"`
	// The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
	// - `INTERNET`: connects clients to cloud desktops only over the Internet.
	// - `VPC`: connects clients to cloud desktops only over a VPC.
	// - `ANY`: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method when you use a client to connect to the cloud desktop.
	DesktopAccessType *string `pulumi:"desktopAccessType"`
	// The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
	DnsAddresses []string `pulumi:"dnsAddresses"`
	// The domain name of the enterprise AD system. You can register each domain name only once.
	DomainName string `pulumi:"domainName"`
	// The password of the domain administrator. The password can be up to 64 characters in length.
	DomainPassword *string `pulumi:"domainPassword"`
	// The username of the domain administrator. The username can be up to 64 characters in length.
	DomainUserName *string `pulumi:"domainUserName"`
	// Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
	EnableAdminAccess *bool `pulumi:"enableAdminAccess"`
	// Specifies whether to enable Internet access.
	EnableInternetAccess *bool `pulumi:"enableInternetAccess"`
	// Specifies whether to enable multi-factor authentication (MFA).
	MfaEnabled *bool `pulumi:"mfaEnabled"`
	// The protocol type. Valid values: `ASP`, `HDX`.
	ProtocolType *string `pulumi:"protocolType"`
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification *int `pulumi:"specification"`
	// The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
	SubDomainDnsAddresses []string `pulumi:"subDomainDnsAddresses"`
	// The domain name of the enterprise AD subdomain.
	SubDomainName *string `pulumi:"subDomainName"`
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
	VerifyCode *string `pulumi:"verifyCode"`
}

// The set of arguments for constructing a AdConnectorOfficeSite resource.
type AdConnectorOfficeSiteArgs struct {
	// The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	AdConnectorOfficeSiteName pulumi.StringInput
	// The ad hostname.
	AdHostname pulumi.StringPtrInput
	// The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
	Bandwidth pulumi.IntPtrInput
	// The ID of the CEN instance.
	CenId pulumi.StringInput
	// The cen owner id.
	CenOwnerId pulumi.StringPtrInput
	// Workspace Corresponds to the Security Office Network of IPv4 Segment.
	CidrBlock pulumi.StringInput
	// The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
	// - `INTERNET`: connects clients to cloud desktops only over the Internet.
	// - `VPC`: connects clients to cloud desktops only over a VPC.
	// - `ANY`: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method when you use a client to connect to the cloud desktop.
	DesktopAccessType pulumi.StringPtrInput
	// The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
	DnsAddresses pulumi.StringArrayInput
	// The domain name of the enterprise AD system. You can register each domain name only once.
	DomainName pulumi.StringInput
	// The password of the domain administrator. The password can be up to 64 characters in length.
	DomainPassword pulumi.StringPtrInput
	// The username of the domain administrator. The username can be up to 64 characters in length.
	DomainUserName pulumi.StringPtrInput
	// Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
	EnableAdminAccess pulumi.BoolPtrInput
	// Specifies whether to enable Internet access.
	EnableInternetAccess pulumi.BoolPtrInput
	// Specifies whether to enable multi-factor authentication (MFA).
	MfaEnabled pulumi.BoolPtrInput
	// The protocol type. Valid values: `ASP`, `HDX`.
	ProtocolType pulumi.StringPtrInput
	// The AD Connector specifications. Valid values: `1`, `2`.
	Specification pulumi.IntPtrInput
	// The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
	SubDomainDnsAddresses pulumi.StringArrayInput
	// The domain name of the enterprise AD subdomain.
	SubDomainName pulumi.StringPtrInput
	// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
	VerifyCode pulumi.StringPtrInput
}

func (AdConnectorOfficeSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adConnectorOfficeSiteArgs)(nil)).Elem()
}

type AdConnectorOfficeSiteInput interface {
	pulumi.Input

	ToAdConnectorOfficeSiteOutput() AdConnectorOfficeSiteOutput
	ToAdConnectorOfficeSiteOutputWithContext(ctx context.Context) AdConnectorOfficeSiteOutput
}

func (*AdConnectorOfficeSite) ElementType() reflect.Type {
	return reflect.TypeOf((**AdConnectorOfficeSite)(nil)).Elem()
}

func (i *AdConnectorOfficeSite) ToAdConnectorOfficeSiteOutput() AdConnectorOfficeSiteOutput {
	return i.ToAdConnectorOfficeSiteOutputWithContext(context.Background())
}

func (i *AdConnectorOfficeSite) ToAdConnectorOfficeSiteOutputWithContext(ctx context.Context) AdConnectorOfficeSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdConnectorOfficeSiteOutput)
}

// AdConnectorOfficeSiteArrayInput is an input type that accepts AdConnectorOfficeSiteArray and AdConnectorOfficeSiteArrayOutput values.
// You can construct a concrete instance of `AdConnectorOfficeSiteArrayInput` via:
//
//          AdConnectorOfficeSiteArray{ AdConnectorOfficeSiteArgs{...} }
type AdConnectorOfficeSiteArrayInput interface {
	pulumi.Input

	ToAdConnectorOfficeSiteArrayOutput() AdConnectorOfficeSiteArrayOutput
	ToAdConnectorOfficeSiteArrayOutputWithContext(context.Context) AdConnectorOfficeSiteArrayOutput
}

type AdConnectorOfficeSiteArray []AdConnectorOfficeSiteInput

func (AdConnectorOfficeSiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdConnectorOfficeSite)(nil)).Elem()
}

func (i AdConnectorOfficeSiteArray) ToAdConnectorOfficeSiteArrayOutput() AdConnectorOfficeSiteArrayOutput {
	return i.ToAdConnectorOfficeSiteArrayOutputWithContext(context.Background())
}

func (i AdConnectorOfficeSiteArray) ToAdConnectorOfficeSiteArrayOutputWithContext(ctx context.Context) AdConnectorOfficeSiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdConnectorOfficeSiteArrayOutput)
}

// AdConnectorOfficeSiteMapInput is an input type that accepts AdConnectorOfficeSiteMap and AdConnectorOfficeSiteMapOutput values.
// You can construct a concrete instance of `AdConnectorOfficeSiteMapInput` via:
//
//          AdConnectorOfficeSiteMap{ "key": AdConnectorOfficeSiteArgs{...} }
type AdConnectorOfficeSiteMapInput interface {
	pulumi.Input

	ToAdConnectorOfficeSiteMapOutput() AdConnectorOfficeSiteMapOutput
	ToAdConnectorOfficeSiteMapOutputWithContext(context.Context) AdConnectorOfficeSiteMapOutput
}

type AdConnectorOfficeSiteMap map[string]AdConnectorOfficeSiteInput

func (AdConnectorOfficeSiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdConnectorOfficeSite)(nil)).Elem()
}

func (i AdConnectorOfficeSiteMap) ToAdConnectorOfficeSiteMapOutput() AdConnectorOfficeSiteMapOutput {
	return i.ToAdConnectorOfficeSiteMapOutputWithContext(context.Background())
}

func (i AdConnectorOfficeSiteMap) ToAdConnectorOfficeSiteMapOutputWithContext(ctx context.Context) AdConnectorOfficeSiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdConnectorOfficeSiteMapOutput)
}

type AdConnectorOfficeSiteOutput struct{ *pulumi.OutputState }

func (AdConnectorOfficeSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdConnectorOfficeSite)(nil)).Elem()
}

func (o AdConnectorOfficeSiteOutput) ToAdConnectorOfficeSiteOutput() AdConnectorOfficeSiteOutput {
	return o
}

func (o AdConnectorOfficeSiteOutput) ToAdConnectorOfficeSiteOutputWithContext(ctx context.Context) AdConnectorOfficeSiteOutput {
	return o
}

// The name of the workspace. The name must be 2 to 255 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
func (o AdConnectorOfficeSiteOutput) AdConnectorOfficeSiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringOutput { return v.AdConnectorOfficeSiteName }).(pulumi.StringOutput)
}

// The ad hostname.
func (o AdConnectorOfficeSiteOutput) AdHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.AdHostname }).(pulumi.StringPtrOutput)
}

// The maximum public bandwidth value. Valid values: 0 to 200. If you do not specify this parameter or you set this parameter to 0, Internet access is disabled.
func (o AdConnectorOfficeSiteOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.IntPtrOutput { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// The ID of the CEN instance.
func (o AdConnectorOfficeSiteOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The cen owner id.
func (o AdConnectorOfficeSiteOutput) CenOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.CenOwnerId }).(pulumi.StringPtrOutput)
}

// Workspace Corresponds to the Security Office Network of IPv4 Segment.
func (o AdConnectorOfficeSiteOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The method that you use to connect to cloud desktops. **Note:** The VPC connection method is provided by Alibaba Cloud PrivateLink. You are not charged for PrivateLink. When you set this parameter to VPC or Any, PrivateLink is automatically activated. Default value: `INTERNET`. Valid values:
// - `INTERNET`: connects clients to cloud desktops only over the Internet.
// - `VPC`: connects clients to cloud desktops only over a VPC.
// - `ANY`: connects clients to cloud desktops over the Internet or a VPC. You can select a connection method when you use a client to connect to the cloud desktop.
func (o AdConnectorOfficeSiteOutput) DesktopAccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringOutput { return v.DesktopAccessType }).(pulumi.StringOutput)
}

// The IP address N of the DNS server of the enterprise AD system. You can specify only one IP address.
func (o AdConnectorOfficeSiteOutput) DnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringArrayOutput { return v.DnsAddresses }).(pulumi.StringArrayOutput)
}

// The domain name of the enterprise AD system. You can register each domain name only once.
func (o AdConnectorOfficeSiteOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The password of the domain administrator. The password can be up to 64 characters in length.
func (o AdConnectorOfficeSiteOutput) DomainPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.DomainPassword }).(pulumi.StringPtrOutput)
}

// The username of the domain administrator. The username can be up to 64 characters in length.
func (o AdConnectorOfficeSiteOutput) DomainUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.DomainUserName }).(pulumi.StringPtrOutput)
}

// Specifies whether to grant the permissions of the local administrator to the desktop users. Default value: true.
func (o AdConnectorOfficeSiteOutput) EnableAdminAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.BoolOutput { return v.EnableAdminAccess }).(pulumi.BoolOutput)
}

// Specifies whether to enable Internet access.
func (o AdConnectorOfficeSiteOutput) EnableInternetAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.BoolOutput { return v.EnableInternetAccess }).(pulumi.BoolOutput)
}

// Specifies whether to enable multi-factor authentication (MFA).
func (o AdConnectorOfficeSiteOutput) MfaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.BoolPtrOutput { return v.MfaEnabled }).(pulumi.BoolPtrOutput)
}

// The protocol type. Valid values: `ASP`, `HDX`.
func (o AdConnectorOfficeSiteOutput) ProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.ProtocolType }).(pulumi.StringPtrOutput)
}

// The AD Connector specifications. Valid values: `1`, `2`.
func (o AdConnectorOfficeSiteOutput) Specification() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.IntPtrOutput { return v.Specification }).(pulumi.IntPtrOutput)
}

// The resource State.
func (o AdConnectorOfficeSiteOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The DNS address N of the enterprise AD subdomain. If you specify a value for the `subDomainName` parameter but you do not specify a value for this parameter, the DNS address of the subdomain is the same as the DNS address of the parent domain.
func (o AdConnectorOfficeSiteOutput) SubDomainDnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringArrayOutput { return v.SubDomainDnsAddresses }).(pulumi.StringArrayOutput)
}

// The domain name of the enterprise AD subdomain.
func (o AdConnectorOfficeSiteOutput) SubDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.SubDomainName }).(pulumi.StringPtrOutput)
}

// The verification code. If the CEN instance that you specify for the CenId parameter belongs to another Alibaba Cloud account, you must call the SendVerifyCode operation to obtain the verification code.
func (o AdConnectorOfficeSiteOutput) VerifyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdConnectorOfficeSite) pulumi.StringPtrOutput { return v.VerifyCode }).(pulumi.StringPtrOutput)
}

type AdConnectorOfficeSiteArrayOutput struct{ *pulumi.OutputState }

func (AdConnectorOfficeSiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AdConnectorOfficeSite)(nil)).Elem()
}

func (o AdConnectorOfficeSiteArrayOutput) ToAdConnectorOfficeSiteArrayOutput() AdConnectorOfficeSiteArrayOutput {
	return o
}

func (o AdConnectorOfficeSiteArrayOutput) ToAdConnectorOfficeSiteArrayOutputWithContext(ctx context.Context) AdConnectorOfficeSiteArrayOutput {
	return o
}

func (o AdConnectorOfficeSiteArrayOutput) Index(i pulumi.IntInput) AdConnectorOfficeSiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AdConnectorOfficeSite {
		return vs[0].([]*AdConnectorOfficeSite)[vs[1].(int)]
	}).(AdConnectorOfficeSiteOutput)
}

type AdConnectorOfficeSiteMapOutput struct{ *pulumi.OutputState }

func (AdConnectorOfficeSiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AdConnectorOfficeSite)(nil)).Elem()
}

func (o AdConnectorOfficeSiteMapOutput) ToAdConnectorOfficeSiteMapOutput() AdConnectorOfficeSiteMapOutput {
	return o
}

func (o AdConnectorOfficeSiteMapOutput) ToAdConnectorOfficeSiteMapOutputWithContext(ctx context.Context) AdConnectorOfficeSiteMapOutput {
	return o
}

func (o AdConnectorOfficeSiteMapOutput) MapIndex(k pulumi.StringInput) AdConnectorOfficeSiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AdConnectorOfficeSite {
		return vs[0].(map[string]*AdConnectorOfficeSite)[vs[1].(string)]
	}).(AdConnectorOfficeSiteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdConnectorOfficeSiteInput)(nil)).Elem(), &AdConnectorOfficeSite{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdConnectorOfficeSiteArrayInput)(nil)).Elem(), AdConnectorOfficeSiteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdConnectorOfficeSiteMapInput)(nil)).Elem(), AdConnectorOfficeSiteMap{})
	pulumi.RegisterOutputType(AdConnectorOfficeSiteOutput{})
	pulumi.RegisterOutputType(AdConnectorOfficeSiteArrayOutput{})
	pulumi.RegisterOutputType(AdConnectorOfficeSiteMapOutput{})
}
