// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Load Balancer Server Certificate is an ssl Certificate used by the listener of the protocol https.
//
// For information about slb and how to use it, see [What is Server Load Balancer](https://www.alibabacloud.com/help/doc-detail/27539.htm).
//
// For information about Server Certificate and how to use it, see [Configure Server Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).
//
// ## Example Usage
//
// * using server_certificate/private content as string example
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := slb.NewServerCertificate(ctx, "foo", &slb.ServerCertificateArgs{
//				PrivateKey:        pulumi.String(fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQDO0knDrlNdiys******ErVpjsckAaOW/JDG5PCSwkaMxk=\n-----END RSA PRIVATE KEY-----\n")),
//				ServerCertificate: pulumi.String(fmt.Sprintf("-----BEGIN CERTIFICATE-----\nMIIDRjCCAq+gAwIBAgI+OuMs******XTtI90EAxEG/bJJyOm5LqoiA=\n-----END CERTIFICATE-----\n")),
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
// * using server_certificate/private file example
//
// ```go
// package main
//
// import (
//
//	"fmt"
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/slb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := slb.NewServerCertificate(ctx, "foo", &slb.ServerCertificateArgs{
//				ServerCertificate: readFileOrPanic(fmt.Sprintf("%v/server_certificate.pem", path.Module)),
//				PrivateKey:        readFileOrPanic(fmt.Sprintf("%v/private_key.pem", path.Module)),
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
// Server Load balancer Server Certificate can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:slb/serverCertificate:ServerCertificate example abc123456
//
// ```
type ServerCertificate struct {
	pulumi.CustomResourceState

	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId pulumi.StringPtrOutput `pulumi:"alicloudCertifacteId"`
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName pulumi.StringPtrOutput `pulumi:"alicloudCertifacteName"`
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId pulumi.StringPtrOutput `pulumi:"alicloudCertificateId"`
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName pulumi.StringPtrOutput `pulumi:"alicloudCertificateName"`
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId pulumi.StringPtrOutput `pulumi:"alicloudCertificateRegionId"`
	// Name of the Server Certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewServerCertificate registers a new resource with the given unique name, arguments, and options.
func NewServerCertificate(ctx *pulumi.Context,
	name string, args *ServerCertificateArgs, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	if args == nil {
		args = &ServerCertificateArgs{}
	}

	var resource ServerCertificate
	err := ctx.RegisterResource("alicloud:slb/serverCertificate:ServerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerCertificate gets an existing ServerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCertificateState, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	var resource ServerCertificate
	err := ctx.ReadResource("alicloud:slb/serverCertificate:ServerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerCertificate resources.
type serverCertificateState struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId *string `pulumi:"alicloudCertifacteId"`
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName *string `pulumi:"alicloudCertifacteName"`
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId *string `pulumi:"alicloudCertificateId"`
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName *string `pulumi:"alicloudCertificateName"`
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId *string `pulumi:"alicloudCertificateRegionId"`
	// Name of the Server Certificate.
	Name *string `pulumi:"name"`
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey *string `pulumi:"privateKey"`
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate *string `pulumi:"serverCertificate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ServerCertificateState struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId pulumi.StringPtrInput
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName pulumi.StringPtrInput
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId pulumi.StringPtrInput
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName pulumi.StringPtrInput
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId pulumi.StringPtrInput
	// Name of the Server Certificate.
	Name pulumi.StringPtrInput
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey pulumi.StringPtrInput
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId pulumi.StringPtrInput
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ServerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateState)(nil)).Elem()
}

type serverCertificateArgs struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId *string `pulumi:"alicloudCertifacteId"`
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName *string `pulumi:"alicloudCertifacteName"`
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId *string `pulumi:"alicloudCertificateId"`
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName *string `pulumi:"alicloudCertificateName"`
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId *string `pulumi:"alicloudCertificateRegionId"`
	// Name of the Server Certificate.
	Name *string `pulumi:"name"`
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey *string `pulumi:"privateKey"`
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate *string `pulumi:"serverCertificate"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ServerCertificate resource.
type ServerCertificateArgs struct {
	// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
	AlicloudCertifacteId pulumi.StringPtrInput
	// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
	AlicloudCertifacteName pulumi.StringPtrInput
	// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateId pulumi.StringPtrInput
	// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateName pulumi.StringPtrInput
	// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
	AlicloudCertificateRegionId pulumi.StringPtrInput
	// Name of the Server Certificate.
	Name pulumi.StringPtrInput
	// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	PrivateKey pulumi.StringPtrInput
	// The Id of resource group which the slb server certificate belongs.
	ResourceGroupId pulumi.StringPtrInput
	// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
	ServerCertificate pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ServerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateArgs)(nil)).Elem()
}

type ServerCertificateInput interface {
	pulumi.Input

	ToServerCertificateOutput() ServerCertificateOutput
	ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput
}

func (*ServerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificate)(nil)).Elem()
}

func (i *ServerCertificate) ToServerCertificateOutput() ServerCertificateOutput {
	return i.ToServerCertificateOutputWithContext(context.Background())
}

func (i *ServerCertificate) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateOutput)
}

// ServerCertificateArrayInput is an input type that accepts ServerCertificateArray and ServerCertificateArrayOutput values.
// You can construct a concrete instance of `ServerCertificateArrayInput` via:
//
//	ServerCertificateArray{ ServerCertificateArgs{...} }
type ServerCertificateArrayInput interface {
	pulumi.Input

	ToServerCertificateArrayOutput() ServerCertificateArrayOutput
	ToServerCertificateArrayOutputWithContext(context.Context) ServerCertificateArrayOutput
}

type ServerCertificateArray []ServerCertificateInput

func (ServerCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerCertificate)(nil)).Elem()
}

func (i ServerCertificateArray) ToServerCertificateArrayOutput() ServerCertificateArrayOutput {
	return i.ToServerCertificateArrayOutputWithContext(context.Background())
}

func (i ServerCertificateArray) ToServerCertificateArrayOutputWithContext(ctx context.Context) ServerCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateArrayOutput)
}

// ServerCertificateMapInput is an input type that accepts ServerCertificateMap and ServerCertificateMapOutput values.
// You can construct a concrete instance of `ServerCertificateMapInput` via:
//
//	ServerCertificateMap{ "key": ServerCertificateArgs{...} }
type ServerCertificateMapInput interface {
	pulumi.Input

	ToServerCertificateMapOutput() ServerCertificateMapOutput
	ToServerCertificateMapOutputWithContext(context.Context) ServerCertificateMapOutput
}

type ServerCertificateMap map[string]ServerCertificateInput

func (ServerCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerCertificate)(nil)).Elem()
}

func (i ServerCertificateMap) ToServerCertificateMapOutput() ServerCertificateMapOutput {
	return i.ToServerCertificateMapOutputWithContext(context.Background())
}

func (i ServerCertificateMap) ToServerCertificateMapOutputWithContext(ctx context.Context) ServerCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateMapOutput)
}

type ServerCertificateOutput struct{ *pulumi.OutputState }

func (ServerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateOutput) ToServerCertificateOutput() ServerCertificateOutput {
	return o
}

func (o ServerCertificateOutput) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return o
}

// Deprecated: Field 'alicloud_certifacte_id' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_id' replaces it.
func (o ServerCertificateOutput) AlicloudCertifacteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.AlicloudCertifacteId }).(pulumi.StringPtrOutput)
}

// Deprecated: Field 'alicloud_certifacte_name' has been deprecated from provider version 1.68.0. Use 'alicloud_certificate_name' replaces it.
func (o ServerCertificateOutput) AlicloudCertifacteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.AlicloudCertifacteName }).(pulumi.StringPtrOutput)
}

// an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.
func (o ServerCertificateOutput) AlicloudCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.AlicloudCertificateId }).(pulumi.StringPtrOutput)
}

// the name of the certificate specified by `alicloudCertificateId`.but it is not supported on the international site of alibaba cloud now.
func (o ServerCertificateOutput) AlicloudCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.AlicloudCertificateName }).(pulumi.StringPtrOutput)
}

// the region of the certificate specified by `alicloudCertificateId`. but it is not supported on the international site of alibaba cloud now.
func (o ServerCertificateOutput) AlicloudCertificateRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.AlicloudCertificateRegionId }).(pulumi.StringPtrOutput)
}

// Name of the Server Certificate.
func (o ServerCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the content of privat key of the ssl certificate specified by `serverCertificate`. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
func (o ServerCertificateOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the slb server certificate belongs.
func (o ServerCertificateOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// the content of the ssl certificate. where `alicloudCertificateId` is null, it is required, otherwise it is ignored.
func (o ServerCertificateOutput) ServerCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.ServerCertificate }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o ServerCertificateOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type ServerCertificateArrayOutput struct{ *pulumi.OutputState }

func (ServerCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateArrayOutput) ToServerCertificateArrayOutput() ServerCertificateArrayOutput {
	return o
}

func (o ServerCertificateArrayOutput) ToServerCertificateArrayOutputWithContext(ctx context.Context) ServerCertificateArrayOutput {
	return o
}

func (o ServerCertificateArrayOutput) Index(i pulumi.IntInput) ServerCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerCertificate {
		return vs[0].([]*ServerCertificate)[vs[1].(int)]
	}).(ServerCertificateOutput)
}

type ServerCertificateMapOutput struct{ *pulumi.OutputState }

func (ServerCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateMapOutput) ToServerCertificateMapOutput() ServerCertificateMapOutput {
	return o
}

func (o ServerCertificateMapOutput) ToServerCertificateMapOutputWithContext(ctx context.Context) ServerCertificateMapOutput {
	return o
}

func (o ServerCertificateMapOutput) MapIndex(k pulumi.StringInput) ServerCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerCertificate {
		return vs[0].(map[string]*ServerCertificate)[vs[1].(string)]
	}).(ServerCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateInput)(nil)).Elem(), &ServerCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateArrayInput)(nil)).Elem(), ServerCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateMapInput)(nil)).Elem(), ServerCertificateMap{})
	pulumi.RegisterOutputType(ServerCertificateOutput{})
	pulumi.RegisterOutputType(ServerCertificateArrayOutput{})
	pulumi.RegisterOutputType(ServerCertificateMapOutput{})
}
