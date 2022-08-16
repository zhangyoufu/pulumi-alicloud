// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CDN Accelerated Domain resource. This resource is based on CDN's new version OpenAPI.
//
// For information about Cdn Domain New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/91176.html).
//
// > **NOTE:** Available in v1.34.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cdn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cdn.NewDomainNew(ctx, "domain", &cdn.DomainNewArgs{
//				CdnType:    pulumi.String("web"),
//				DomainName: pulumi.String("terraform.test.com"),
//				Scope:      pulumi.String("overseas"),
//				Sources: cdn.DomainNewSourceArray{
//					&cdn.DomainNewSourceArgs{
//						Content:  pulumi.String("1.1.1.1"),
//						Port:     pulumi.Int(80),
//						Priority: pulumi.Int(20),
//						Type:     pulumi.String("ipaddr"),
//						Weight:   pulumi.Int(10),
//					},
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
//
// ## Import
//
// CDN domain can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cdn/domainNew:DomainNew example xxxx.com
//
// ```
type DomainNew struct {
	pulumi.CustomResourceState

	// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
	CdnType pulumi.StringOutput `pulumi:"cdnType"`
	// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
	CertificateConfig DomainNewCertificateConfigOutput `pulumi:"certificateConfig"`
	// (Available in v1.90.0+) The CNAME of the CDN domain.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Resource group ID.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
	Scope pulumi.StringOutput `pulumi:"scope"`
	// The source address list of the accelerated domain. Defaults to null. See Block Sources.
	Sources DomainNewSourceArrayOutput `pulumi:"sources"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDomainNew registers a new resource with the given unique name, arguments, and options.
func NewDomainNew(ctx *pulumi.Context,
	name string, args *DomainNewArgs, opts ...pulumi.ResourceOption) (*DomainNew, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CdnType == nil {
		return nil, errors.New("invalid value for required argument 'CdnType'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	var resource DomainNew
	err := ctx.RegisterResource("alicloud:cdn/domainNew:DomainNew", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainNew gets an existing DomainNew resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainNew(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainNewState, opts ...pulumi.ResourceOption) (*DomainNew, error) {
	var resource DomainNew
	err := ctx.ReadResource("alicloud:cdn/domainNew:DomainNew", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainNew resources.
type domainNewState struct {
	// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
	CdnType *string `pulumi:"cdnType"`
	// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
	CertificateConfig *DomainNewCertificateConfig `pulumi:"certificateConfig"`
	// (Available in v1.90.0+) The CNAME of the CDN domain.
	Cname *string `pulumi:"cname"`
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName *string `pulumi:"domainName"`
	// Resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
	Scope *string `pulumi:"scope"`
	// The source address list of the accelerated domain. Defaults to null. See Block Sources.
	Sources []DomainNewSource `pulumi:"sources"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type DomainNewState struct {
	// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
	CdnType pulumi.StringPtrInput
	// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
	CertificateConfig DomainNewCertificateConfigPtrInput
	// (Available in v1.90.0+) The CNAME of the CDN domain.
	Cname pulumi.StringPtrInput
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringPtrInput
	// Resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
	Scope pulumi.StringPtrInput
	// The source address list of the accelerated domain. Defaults to null. See Block Sources.
	Sources DomainNewSourceArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (DomainNewState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNewState)(nil)).Elem()
}

type domainNewArgs struct {
	// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
	CdnType string `pulumi:"cdnType"`
	// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
	CertificateConfig *DomainNewCertificateConfig `pulumi:"certificateConfig"`
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName string `pulumi:"domainName"`
	// Resource group ID.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
	Scope *string `pulumi:"scope"`
	// The source address list of the accelerated domain. Defaults to null. See Block Sources.
	Sources []DomainNewSource `pulumi:"sources"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a DomainNew resource.
type DomainNewArgs struct {
	// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
	CdnType pulumi.StringInput
	// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
	CertificateConfig DomainNewCertificateConfigPtrInput
	// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
	DomainName pulumi.StringInput
	// Resource group ID.
	ResourceGroupId pulumi.StringPtrInput
	// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
	Scope pulumi.StringPtrInput
	// The source address list of the accelerated domain. Defaults to null. See Block Sources.
	Sources DomainNewSourceArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (DomainNewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNewArgs)(nil)).Elem()
}

type DomainNewInput interface {
	pulumi.Input

	ToDomainNewOutput() DomainNewOutput
	ToDomainNewOutputWithContext(ctx context.Context) DomainNewOutput
}

func (*DomainNew) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainNew)(nil)).Elem()
}

func (i *DomainNew) ToDomainNewOutput() DomainNewOutput {
	return i.ToDomainNewOutputWithContext(context.Background())
}

func (i *DomainNew) ToDomainNewOutputWithContext(ctx context.Context) DomainNewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNewOutput)
}

// DomainNewArrayInput is an input type that accepts DomainNewArray and DomainNewArrayOutput values.
// You can construct a concrete instance of `DomainNewArrayInput` via:
//
//	DomainNewArray{ DomainNewArgs{...} }
type DomainNewArrayInput interface {
	pulumi.Input

	ToDomainNewArrayOutput() DomainNewArrayOutput
	ToDomainNewArrayOutputWithContext(context.Context) DomainNewArrayOutput
}

type DomainNewArray []DomainNewInput

func (DomainNewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainNew)(nil)).Elem()
}

func (i DomainNewArray) ToDomainNewArrayOutput() DomainNewArrayOutput {
	return i.ToDomainNewArrayOutputWithContext(context.Background())
}

func (i DomainNewArray) ToDomainNewArrayOutputWithContext(ctx context.Context) DomainNewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNewArrayOutput)
}

// DomainNewMapInput is an input type that accepts DomainNewMap and DomainNewMapOutput values.
// You can construct a concrete instance of `DomainNewMapInput` via:
//
//	DomainNewMap{ "key": DomainNewArgs{...} }
type DomainNewMapInput interface {
	pulumi.Input

	ToDomainNewMapOutput() DomainNewMapOutput
	ToDomainNewMapOutputWithContext(context.Context) DomainNewMapOutput
}

type DomainNewMap map[string]DomainNewInput

func (DomainNewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainNew)(nil)).Elem()
}

func (i DomainNewMap) ToDomainNewMapOutput() DomainNewMapOutput {
	return i.ToDomainNewMapOutputWithContext(context.Background())
}

func (i DomainNewMap) ToDomainNewMapOutputWithContext(ctx context.Context) DomainNewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNewMapOutput)
}

type DomainNewOutput struct{ *pulumi.OutputState }

func (DomainNewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainNew)(nil)).Elem()
}

func (o DomainNewOutput) ToDomainNewOutput() DomainNewOutput {
	return o
}

func (o DomainNewOutput) ToDomainNewOutputWithContext(ctx context.Context) DomainNewOutput {
	return o
}

// Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
func (o DomainNewOutput) CdnType() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNew) pulumi.StringOutput { return v.CdnType }).(pulumi.StringOutput)
}

// Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
func (o DomainNewOutput) CertificateConfig() DomainNewCertificateConfigOutput {
	return o.ApplyT(func(v *DomainNew) DomainNewCertificateConfigOutput { return v.CertificateConfig }).(DomainNewCertificateConfigOutput)
}

// (Available in v1.90.0+) The CNAME of the CDN domain.
func (o DomainNewOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNew) pulumi.StringOutput { return v.Cname }).(pulumi.StringOutput)
}

// Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
func (o DomainNewOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNew) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Resource group ID.
func (o DomainNewOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNew) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
func (o DomainNewOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNew) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// The source address list of the accelerated domain. Defaults to null. See Block Sources.
func (o DomainNewOutput) Sources() DomainNewSourceArrayOutput {
	return o.ApplyT(func(v *DomainNew) DomainNewSourceArrayOutput { return v.Sources }).(DomainNewSourceArrayOutput)
}

// A mapping of tags to assign to the resource.
func (o DomainNewOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *DomainNew) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type DomainNewArrayOutput struct{ *pulumi.OutputState }

func (DomainNewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainNew)(nil)).Elem()
}

func (o DomainNewArrayOutput) ToDomainNewArrayOutput() DomainNewArrayOutput {
	return o
}

func (o DomainNewArrayOutput) ToDomainNewArrayOutputWithContext(ctx context.Context) DomainNewArrayOutput {
	return o
}

func (o DomainNewArrayOutput) Index(i pulumi.IntInput) DomainNewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainNew {
		return vs[0].([]*DomainNew)[vs[1].(int)]
	}).(DomainNewOutput)
}

type DomainNewMapOutput struct{ *pulumi.OutputState }

func (DomainNewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainNew)(nil)).Elem()
}

func (o DomainNewMapOutput) ToDomainNewMapOutput() DomainNewMapOutput {
	return o
}

func (o DomainNewMapOutput) ToDomainNewMapOutputWithContext(ctx context.Context) DomainNewMapOutput {
	return o
}

func (o DomainNewMapOutput) MapIndex(k pulumi.StringInput) DomainNewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainNew {
		return vs[0].(map[string]*DomainNew)[vs[1].(string)]
	}).(DomainNewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNewInput)(nil)).Elem(), &DomainNew{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNewArrayInput)(nil)).Elem(), DomainNewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNewMapInput)(nil)).Elem(), DomainNewMap{})
	pulumi.RegisterOutputType(DomainNewOutput{})
	pulumi.RegisterOutputType(DomainNewArrayOutput{})
	pulumi.RegisterOutputType(DomainNewMapOutput{})
}
