// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// CBWP Common Bandwidth Package can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:vpc/commonBandwithPackage:CommonBandwithPackage example <id>
// ```
type CommonBandwithPackage struct {
	pulumi.CustomResourceState

	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	BandwidthPackageName pulumi.StringOutput `pulumi:"bandwidthPackageName"`
	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Specifies whether to enable deletion protection. Valid values:
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
	Force pulumi.StringPtrOutput `pulumi:"force"`
	// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// The line type. Valid values:
	// - `BGP` All regions support BGP (Multi-ISP).
	// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	// - `ChinaTelecom`
	// - `ChinaUnicom`
	// - `ChinaMobile`
	// - `ChinaTelecom_L2`
	// - `ChinaUnicom_L2`
	// - `ChinaMobile_L2`
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
	Isp pulumi.StringOutput `pulumi:"isp"`
	// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The billing type of the Internet Shared Bandwidth instance. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
	//
	// > **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
	Ratio pulumi.IntOutput `pulumi:"ratio"`
	// The ID of the resource group to which you want to move the resource.
	//
	// > **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internetChargeType` is `PayBy95`.
	SecurityProtectionTypes pulumi.StringArrayOutput `pulumi:"securityProtectionTypes"`
	// The status of the Internet Shared Bandwidth instance. Default value: `Available`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tag of the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewCommonBandwithPackage registers a new resource with the given unique name, arguments, and options.
func NewCommonBandwithPackage(ctx *pulumi.Context,
	name string, args *CommonBandwithPackageArgs, opts ...pulumi.ResourceOption) (*CommonBandwithPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CommonBandwithPackage
	err := ctx.RegisterResource("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommonBandwithPackage gets an existing CommonBandwithPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommonBandwithPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommonBandwithPackageState, opts ...pulumi.ResourceOption) (*CommonBandwithPackage, error) {
	var resource CommonBandwithPackage
	err := ctx.ReadResource("alicloud:vpc/commonBandwithPackage:CommonBandwithPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommonBandwithPackage resources.
type commonBandwithPackageState struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
	Bandwidth *string `pulumi:"bandwidth"`
	// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// The creation time.
	CreateTime *string `pulumi:"createTime"`
	// Specifies whether to enable deletion protection. Valid values:
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
	Force *string `pulumi:"force"`
	// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The line type. Valid values:
	// - `BGP` All regions support BGP (Multi-ISP).
	// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	// - `ChinaTelecom`
	// - `ChinaUnicom`
	// - `ChinaMobile`
	// - `ChinaTelecom_L2`
	// - `ChinaUnicom_L2`
	// - `ChinaMobile_L2`
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
	Isp *string `pulumi:"isp"`
	// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name *string `pulumi:"name"`
	// The billing type of the Internet Shared Bandwidth instance. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType *string `pulumi:"paymentType"`
	// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
	//
	// > **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
	Ratio *int `pulumi:"ratio"`
	// The ID of the resource group to which you want to move the resource.
	//
	// > **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internetChargeType` is `PayBy95`.
	SecurityProtectionTypes []string `pulumi:"securityProtectionTypes"`
	// The status of the Internet Shared Bandwidth instance. Default value: `Available`.
	Status *string `pulumi:"status"`
	// The tag of the resource
	Tags map[string]string `pulumi:"tags"`
	// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Zone *string `pulumi:"zone"`
}

type CommonBandwithPackageState struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
	Bandwidth pulumi.StringPtrInput
	// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	BandwidthPackageName pulumi.StringPtrInput
	// The creation time.
	CreateTime pulumi.StringPtrInput
	// Specifies whether to enable deletion protection. Valid values:
	DeletionProtection pulumi.BoolPtrInput
	// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
	Force pulumi.StringPtrInput
	// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
	InternetChargeType pulumi.StringPtrInput
	// The line type. Valid values:
	// - `BGP` All regions support BGP (Multi-ISP).
	// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	// - `ChinaTelecom`
	// - `ChinaUnicom`
	// - `ChinaMobile`
	// - `ChinaTelecom_L2`
	// - `ChinaUnicom_L2`
	// - `ChinaMobile_L2`
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
	Isp pulumi.StringPtrInput
	// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name pulumi.StringPtrInput
	// The billing type of the Internet Shared Bandwidth instance. Valid values: `PayAsYouGo`, `Subscription`.
	PaymentType pulumi.StringPtrInput
	// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
	//
	// > **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
	Ratio pulumi.IntPtrInput
	// The ID of the resource group to which you want to move the resource.
	//
	// > **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
	ResourceGroupId pulumi.StringPtrInput
	// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internetChargeType` is `PayBy95`.
	SecurityProtectionTypes pulumi.StringArrayInput
	// The status of the Internet Shared Bandwidth instance. Default value: `Available`.
	Status pulumi.StringPtrInput
	// The tag of the resource
	Tags pulumi.StringMapInput
	// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Zone pulumi.StringPtrInput
}

func (CommonBandwithPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageState)(nil)).Elem()
}

type commonBandwithPackageArgs struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
	Bandwidth string `pulumi:"bandwidth"`
	// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// Specifies whether to enable deletion protection. Valid values:
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	Description *string `pulumi:"description"`
	// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
	Force *string `pulumi:"force"`
	// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The line type. Valid values:
	// - `BGP` All regions support BGP (Multi-ISP).
	// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	// - `ChinaTelecom`
	// - `ChinaUnicom`
	// - `ChinaMobile`
	// - `ChinaTelecom_L2`
	// - `ChinaUnicom_L2`
	// - `ChinaMobile_L2`
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
	Isp *string `pulumi:"isp"`
	// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name *string `pulumi:"name"`
	// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
	//
	// > **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
	Ratio *int `pulumi:"ratio"`
	// The ID of the resource group to which you want to move the resource.
	//
	// > **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internetChargeType` is `PayBy95`.
	SecurityProtectionTypes []string `pulumi:"securityProtectionTypes"`
	// The tag of the resource
	Tags map[string]string `pulumi:"tags"`
	// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a CommonBandwithPackage resource.
type CommonBandwithPackageArgs struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
	Bandwidth pulumi.StringInput
	// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
	BandwidthPackageName pulumi.StringPtrInput
	// Specifies whether to enable deletion protection. Valid values:
	DeletionProtection pulumi.BoolPtrInput
	// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	Description pulumi.StringPtrInput
	// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
	Force pulumi.StringPtrInput
	// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
	InternetChargeType pulumi.StringPtrInput
	// The line type. Valid values:
	// - `BGP` All regions support BGP (Multi-ISP).
	// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	// - `ChinaTelecom`
	// - `ChinaUnicom`
	// - `ChinaMobile`
	// - `ChinaTelecom_L2`
	// - `ChinaUnicom_L2`
	// - `ChinaMobile_L2`
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
	Isp pulumi.StringPtrInput
	// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
	//
	// Deprecated: Field 'name' has been deprecated since provider version 1.120.0. New field 'bandwidth_package_name' instead.
	Name pulumi.StringPtrInput
	// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
	//
	// > **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
	Ratio pulumi.IntPtrInput
	// The ID of the resource group to which you want to move the resource.
	//
	// > **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
	ResourceGroupId pulumi.StringPtrInput
	// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internetChargeType` is `PayBy95`.
	SecurityProtectionTypes pulumi.StringArrayInput
	// The tag of the resource
	Tags pulumi.StringMapInput
	// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
	//
	// The following arguments will be discarded. Please use new fields as soon as possible:
	Zone pulumi.StringPtrInput
}

func (CommonBandwithPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commonBandwithPackageArgs)(nil)).Elem()
}

type CommonBandwithPackageInput interface {
	pulumi.Input

	ToCommonBandwithPackageOutput() CommonBandwithPackageOutput
	ToCommonBandwithPackageOutputWithContext(ctx context.Context) CommonBandwithPackageOutput
}

func (*CommonBandwithPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonBandwithPackage)(nil)).Elem()
}

func (i *CommonBandwithPackage) ToCommonBandwithPackageOutput() CommonBandwithPackageOutput {
	return i.ToCommonBandwithPackageOutputWithContext(context.Background())
}

func (i *CommonBandwithPackage) ToCommonBandwithPackageOutputWithContext(ctx context.Context) CommonBandwithPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageOutput)
}

// CommonBandwithPackageArrayInput is an input type that accepts CommonBandwithPackageArray and CommonBandwithPackageArrayOutput values.
// You can construct a concrete instance of `CommonBandwithPackageArrayInput` via:
//
//	CommonBandwithPackageArray{ CommonBandwithPackageArgs{...} }
type CommonBandwithPackageArrayInput interface {
	pulumi.Input

	ToCommonBandwithPackageArrayOutput() CommonBandwithPackageArrayOutput
	ToCommonBandwithPackageArrayOutputWithContext(context.Context) CommonBandwithPackageArrayOutput
}

type CommonBandwithPackageArray []CommonBandwithPackageInput

func (CommonBandwithPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CommonBandwithPackage)(nil)).Elem()
}

func (i CommonBandwithPackageArray) ToCommonBandwithPackageArrayOutput() CommonBandwithPackageArrayOutput {
	return i.ToCommonBandwithPackageArrayOutputWithContext(context.Background())
}

func (i CommonBandwithPackageArray) ToCommonBandwithPackageArrayOutputWithContext(ctx context.Context) CommonBandwithPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageArrayOutput)
}

// CommonBandwithPackageMapInput is an input type that accepts CommonBandwithPackageMap and CommonBandwithPackageMapOutput values.
// You can construct a concrete instance of `CommonBandwithPackageMapInput` via:
//
//	CommonBandwithPackageMap{ "key": CommonBandwithPackageArgs{...} }
type CommonBandwithPackageMapInput interface {
	pulumi.Input

	ToCommonBandwithPackageMapOutput() CommonBandwithPackageMapOutput
	ToCommonBandwithPackageMapOutputWithContext(context.Context) CommonBandwithPackageMapOutput
}

type CommonBandwithPackageMap map[string]CommonBandwithPackageInput

func (CommonBandwithPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CommonBandwithPackage)(nil)).Elem()
}

func (i CommonBandwithPackageMap) ToCommonBandwithPackageMapOutput() CommonBandwithPackageMapOutput {
	return i.ToCommonBandwithPackageMapOutputWithContext(context.Background())
}

func (i CommonBandwithPackageMap) ToCommonBandwithPackageMapOutputWithContext(ctx context.Context) CommonBandwithPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonBandwithPackageMapOutput)
}

type CommonBandwithPackageOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonBandwithPackage)(nil)).Elem()
}

func (o CommonBandwithPackageOutput) ToCommonBandwithPackageOutput() CommonBandwithPackageOutput {
	return o
}

func (o CommonBandwithPackageOutput) ToCommonBandwithPackageOutputWithContext(ctx context.Context) CommonBandwithPackageOutput {
	return o
}

// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s. Valid values: `1` to `1000`. Default value: `1`.
func (o CommonBandwithPackageOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.Bandwidth }).(pulumi.StringOutput)
}

// The name of the Internet Shared Bandwidth instance. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-). The name must start with a letter.
func (o CommonBandwithPackageOutput) BandwidthPackageName() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.BandwidthPackageName }).(pulumi.StringOutput)
}

// The creation time.
func (o CommonBandwithPackageOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Specifies whether to enable deletion protection. Valid values:
func (o CommonBandwithPackageOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// The description of the Internet Shared Bandwidth instance. The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
func (o CommonBandwithPackageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to forcefully delete the Internet Shared Bandwidth instance. Valid values:
func (o CommonBandwithPackageOutput) Force() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringPtrOutput { return v.Force }).(pulumi.StringPtrOutput)
}

// The billing method of the Internet Shared Bandwidth instance. Set the value to `PayByTraffic`, which specifies the pay-by-data-transfer billing method.
func (o CommonBandwithPackageOutput) InternetChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringPtrOutput { return v.InternetChargeType }).(pulumi.StringPtrOutput)
}

// The line type. Valid values:
// - `BGP` All regions support BGP (Multi-ISP).
// - `BGP_PRO` BGP (Multi-ISP) Pro lines are available in the China (Hong Kong), Singapore, Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
//
// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
// - `ChinaTelecom`
// - `ChinaUnicom`
// - `ChinaMobile`
// - `ChinaTelecom_L2`
// - `ChinaUnicom_L2`
// - `ChinaMobile_L2`
//
// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to `BGP_FinanceCloud`.
func (o CommonBandwithPackageOutput) Isp() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.Isp }).(pulumi.StringOutput)
}

// . Field 'name' has been deprecated from provider version 1.120.0. New field 'bandwidth_package_name' instead.
//
// Deprecated: Field 'name' has been deprecated since provider version 1.120.0. New field 'bandwidth_package_name' instead.
func (o CommonBandwithPackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The billing type of the Internet Shared Bandwidth instance. Valid values: `PayAsYouGo`, `Subscription`.
func (o CommonBandwithPackageOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// The percentage of the minimum bandwidth commitment. Set the parameter to `20`.
//
// > **NOTE:**  This parameter is available only on the Alibaba Cloud China site.
func (o CommonBandwithPackageOutput) Ratio() pulumi.IntOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.IntOutput { return v.Ratio }).(pulumi.IntOutput)
}

// The ID of the resource group to which you want to move the resource.
//
// > **NOTE:**   You can use resource groups to facilitate resource grouping and permission management for an Alibaba Cloud. For more information, see [What is resource management?](https://www.alibabacloud.com/help/en/doc-detail/94475.html)
func (o CommonBandwithPackageOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to AntiDDoS_Enhanced, Anti-DDoS Pro(Premium) is used. It is valid when `internetChargeType` is `PayBy95`.
func (o CommonBandwithPackageOutput) SecurityProtectionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringArrayOutput { return v.SecurityProtectionTypes }).(pulumi.StringArrayOutput)
}

// The status of the Internet Shared Bandwidth instance. Default value: `Available`.
func (o CommonBandwithPackageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tag of the resource
func (o CommonBandwithPackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The zone of the Internet Shared Bandwidth instance. This parameter is required if you create an Internet Shared Bandwidth instance for a cloud box.
//
// The following arguments will be discarded. Please use new fields as soon as possible:
func (o CommonBandwithPackageOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonBandwithPackage) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

type CommonBandwithPackageArrayOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CommonBandwithPackage)(nil)).Elem()
}

func (o CommonBandwithPackageArrayOutput) ToCommonBandwithPackageArrayOutput() CommonBandwithPackageArrayOutput {
	return o
}

func (o CommonBandwithPackageArrayOutput) ToCommonBandwithPackageArrayOutputWithContext(ctx context.Context) CommonBandwithPackageArrayOutput {
	return o
}

func (o CommonBandwithPackageArrayOutput) Index(i pulumi.IntInput) CommonBandwithPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CommonBandwithPackage {
		return vs[0].([]*CommonBandwithPackage)[vs[1].(int)]
	}).(CommonBandwithPackageOutput)
}

type CommonBandwithPackageMapOutput struct{ *pulumi.OutputState }

func (CommonBandwithPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CommonBandwithPackage)(nil)).Elem()
}

func (o CommonBandwithPackageMapOutput) ToCommonBandwithPackageMapOutput() CommonBandwithPackageMapOutput {
	return o
}

func (o CommonBandwithPackageMapOutput) ToCommonBandwithPackageMapOutputWithContext(ctx context.Context) CommonBandwithPackageMapOutput {
	return o
}

func (o CommonBandwithPackageMapOutput) MapIndex(k pulumi.StringInput) CommonBandwithPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CommonBandwithPackage {
		return vs[0].(map[string]*CommonBandwithPackage)[vs[1].(string)]
	}).(CommonBandwithPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageInput)(nil)).Elem(), &CommonBandwithPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageArrayInput)(nil)).Elem(), CommonBandwithPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommonBandwithPackageMapInput)(nil)).Elem(), CommonBandwithPackageMap{})
	pulumi.RegisterOutputType(CommonBandwithPackageOutput{})
	pulumi.RegisterOutputType(CommonBandwithPackageArrayOutput{})
	pulumi.RegisterOutputType(CommonBandwithPackageMapOutput{})
}
