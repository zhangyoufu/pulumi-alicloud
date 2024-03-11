// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package threatdetection

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Threat Detection Web Lock Config resource.
//
// For information about Threat Detection Web Lock Config and how to use it, see [What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifyweblockstart).
//
// > **NOTE:** Available in v1.195.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/threatdetection"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultAssets, err := threatdetection.GetAssets(ctx, &threatdetection.GetAssetsArgs{
//				MachineTypes: pulumi.StringRef("ecs"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = threatdetection.NewWebLockConfig(ctx, "defaultWebLockConfig", &threatdetection.WebLockConfigArgs{
//				InclusiveFileType: pulumi.String("php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg"),
//				Uuid:              *pulumi.String(defaultAssets.Ids[0]),
//				Mode:              pulumi.String("whitelist"),
//				LocalBackupDir:    pulumi.String("/usr/local/aegis/bak"),
//				Dir:               pulumi.String("/tmp/"),
//				DefenceMode:       pulumi.String("audit"),
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
// Threat Detection Web Lock Config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:threatdetection/webLockConfig:WebLockConfig example <id>
// ```
type WebLockConfig struct {
	pulumi.CustomResourceState

	// Protection mode. Value:-**block**: Intercept-**audit**: Alarm
	DefenceMode pulumi.StringOutput `pulumi:"defenceMode"`
	// Specify the protection directory.
	Dir pulumi.StringOutput `pulumi:"dir"`
	// Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveDir pulumi.StringPtrOutput `pulumi:"exclusiveDir"`
	// Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFile pulumi.StringPtrOutput `pulumi:"exclusiveFile"`
	// Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFileType pulumi.StringPtrOutput `pulumi:"exclusiveFileType"`
	// Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
	InclusiveFileType pulumi.StringPtrOutput `pulumi:"inclusiveFileType"`
	// The local backup path is used to protect the safe backup of the Directory.
	LocalBackupDir pulumi.StringOutput `pulumi:"localBackupDir"`
	// Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewWebLockConfig registers a new resource with the given unique name, arguments, and options.
func NewWebLockConfig(ctx *pulumi.Context,
	name string, args *WebLockConfigArgs, opts ...pulumi.ResourceOption) (*WebLockConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefenceMode == nil {
		return nil, errors.New("invalid value for required argument 'DefenceMode'")
	}
	if args.Dir == nil {
		return nil, errors.New("invalid value for required argument 'Dir'")
	}
	if args.LocalBackupDir == nil {
		return nil, errors.New("invalid value for required argument 'LocalBackupDir'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Uuid == nil {
		return nil, errors.New("invalid value for required argument 'Uuid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebLockConfig
	err := ctx.RegisterResource("alicloud:threatdetection/webLockConfig:WebLockConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebLockConfig gets an existing WebLockConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebLockConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebLockConfigState, opts ...pulumi.ResourceOption) (*WebLockConfig, error) {
	var resource WebLockConfig
	err := ctx.ReadResource("alicloud:threatdetection/webLockConfig:WebLockConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebLockConfig resources.
type webLockConfigState struct {
	// Protection mode. Value:-**block**: Intercept-**audit**: Alarm
	DefenceMode *string `pulumi:"defenceMode"`
	// Specify the protection directory.
	Dir *string `pulumi:"dir"`
	// Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveDir *string `pulumi:"exclusiveDir"`
	// Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFile *string `pulumi:"exclusiveFile"`
	// Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFileType *string `pulumi:"exclusiveFileType"`
	// Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
	InclusiveFileType *string `pulumi:"inclusiveFileType"`
	// The local backup path is used to protect the safe backup of the Directory.
	LocalBackupDir *string `pulumi:"localBackupDir"`
	// Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
	Mode *string `pulumi:"mode"`
	// Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
	Uuid *string `pulumi:"uuid"`
}

type WebLockConfigState struct {
	// Protection mode. Value:-**block**: Intercept-**audit**: Alarm
	DefenceMode pulumi.StringPtrInput
	// Specify the protection directory.
	Dir pulumi.StringPtrInput
	// Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveDir pulumi.StringPtrInput
	// Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFile pulumi.StringPtrInput
	// Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFileType pulumi.StringPtrInput
	// Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
	InclusiveFileType pulumi.StringPtrInput
	// The local backup path is used to protect the safe backup of the Directory.
	LocalBackupDir pulumi.StringPtrInput
	// Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
	Mode pulumi.StringPtrInput
	// Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
	Uuid pulumi.StringPtrInput
}

func (WebLockConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*webLockConfigState)(nil)).Elem()
}

type webLockConfigArgs struct {
	// Protection mode. Value:-**block**: Intercept-**audit**: Alarm
	DefenceMode string `pulumi:"defenceMode"`
	// Specify the protection directory.
	Dir string `pulumi:"dir"`
	// Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveDir *string `pulumi:"exclusiveDir"`
	// Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFile *string `pulumi:"exclusiveFile"`
	// Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFileType *string `pulumi:"exclusiveFileType"`
	// Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
	InclusiveFileType *string `pulumi:"inclusiveFileType"`
	// The local backup path is used to protect the safe backup of the Directory.
	LocalBackupDir string `pulumi:"localBackupDir"`
	// Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
	Mode string `pulumi:"mode"`
	// Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
	Uuid string `pulumi:"uuid"`
}

// The set of arguments for constructing a WebLockConfig resource.
type WebLockConfigArgs struct {
	// Protection mode. Value:-**block**: Intercept-**audit**: Alarm
	DefenceMode pulumi.StringInput
	// Specify the protection directory.
	Dir pulumi.StringInput
	// Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveDir pulumi.StringPtrInput
	// Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFile pulumi.StringPtrInput
	// Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
	ExclusiveFileType pulumi.StringPtrInput
	// Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
	InclusiveFileType pulumi.StringPtrInput
	// The local backup path is used to protect the safe backup of the Directory.
	LocalBackupDir pulumi.StringInput
	// Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
	Mode pulumi.StringInput
	// Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
	Uuid pulumi.StringInput
}

func (WebLockConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webLockConfigArgs)(nil)).Elem()
}

type WebLockConfigInput interface {
	pulumi.Input

	ToWebLockConfigOutput() WebLockConfigOutput
	ToWebLockConfigOutputWithContext(ctx context.Context) WebLockConfigOutput
}

func (*WebLockConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**WebLockConfig)(nil)).Elem()
}

func (i *WebLockConfig) ToWebLockConfigOutput() WebLockConfigOutput {
	return i.ToWebLockConfigOutputWithContext(context.Background())
}

func (i *WebLockConfig) ToWebLockConfigOutputWithContext(ctx context.Context) WebLockConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebLockConfigOutput)
}

// WebLockConfigArrayInput is an input type that accepts WebLockConfigArray and WebLockConfigArrayOutput values.
// You can construct a concrete instance of `WebLockConfigArrayInput` via:
//
//	WebLockConfigArray{ WebLockConfigArgs{...} }
type WebLockConfigArrayInput interface {
	pulumi.Input

	ToWebLockConfigArrayOutput() WebLockConfigArrayOutput
	ToWebLockConfigArrayOutputWithContext(context.Context) WebLockConfigArrayOutput
}

type WebLockConfigArray []WebLockConfigInput

func (WebLockConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebLockConfig)(nil)).Elem()
}

func (i WebLockConfigArray) ToWebLockConfigArrayOutput() WebLockConfigArrayOutput {
	return i.ToWebLockConfigArrayOutputWithContext(context.Background())
}

func (i WebLockConfigArray) ToWebLockConfigArrayOutputWithContext(ctx context.Context) WebLockConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebLockConfigArrayOutput)
}

// WebLockConfigMapInput is an input type that accepts WebLockConfigMap and WebLockConfigMapOutput values.
// You can construct a concrete instance of `WebLockConfigMapInput` via:
//
//	WebLockConfigMap{ "key": WebLockConfigArgs{...} }
type WebLockConfigMapInput interface {
	pulumi.Input

	ToWebLockConfigMapOutput() WebLockConfigMapOutput
	ToWebLockConfigMapOutputWithContext(context.Context) WebLockConfigMapOutput
}

type WebLockConfigMap map[string]WebLockConfigInput

func (WebLockConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebLockConfig)(nil)).Elem()
}

func (i WebLockConfigMap) ToWebLockConfigMapOutput() WebLockConfigMapOutput {
	return i.ToWebLockConfigMapOutputWithContext(context.Background())
}

func (i WebLockConfigMap) ToWebLockConfigMapOutputWithContext(ctx context.Context) WebLockConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebLockConfigMapOutput)
}

type WebLockConfigOutput struct{ *pulumi.OutputState }

func (WebLockConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebLockConfig)(nil)).Elem()
}

func (o WebLockConfigOutput) ToWebLockConfigOutput() WebLockConfigOutput {
	return o
}

func (o WebLockConfigOutput) ToWebLockConfigOutputWithContext(ctx context.Context) WebLockConfigOutput {
	return o
}

// Protection mode. Value:-**block**: Intercept-**audit**: Alarm
func (o WebLockConfigOutput) DefenceMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringOutput { return v.DefenceMode }).(pulumi.StringOutput)
}

// Specify the protection directory.
func (o WebLockConfigOutput) Dir() pulumi.StringOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringOutput { return v.Dir }).(pulumi.StringOutput)
}

// Specify a directory address that does not require Web tamper protection (I. E. Excluded directories).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
func (o WebLockConfigOutput) ExclusiveDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringPtrOutput { return v.ExclusiveDir }).(pulumi.StringPtrOutput)
}

// Specify files that do not need to enable tamper protection for web pages (that is, exclude files).> The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
func (o WebLockConfigOutput) ExclusiveFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringPtrOutput { return v.ExclusiveFile }).(pulumi.StringPtrOutput)
}

// Specify the type of file that does not require Web tamper protection (that is, the type of excluded file). When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png > The protection Mode **Mode** is set to **blacklist**, you need to configure this parameter.
func (o WebLockConfigOutput) ExclusiveFileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringPtrOutput { return v.ExclusiveFileType }).(pulumi.StringPtrOutput)
}

// Specify the type of file that requires tamper protection. When there are multiple file types, use semicolons (;) separation. Value:-php-jsp-asp-aspx-js-cgi-html-htm-xml-shtml-shtm-jpg-gif-png> The protection Mode **Mode** is set to **whitelist**, you need to configure this parameter.
func (o WebLockConfigOutput) InclusiveFileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringPtrOutput { return v.InclusiveFileType }).(pulumi.StringPtrOutput)
}

// The local backup path is used to protect the safe backup of the Directory.
func (o WebLockConfigOutput) LocalBackupDir() pulumi.StringOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringOutput { return v.LocalBackupDir }).(pulumi.StringOutput)
}

// Specify the protected directory mode. Value:-**whitelist**: whitelist mode, which protects the added protected directories and file types.-**blacklist**: blacklist mode, which protects all unexcluded subdirectories, file types, and specified files under the added protection directory.
func (o WebLockConfigOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Specify the UUID of the server to which you want to add a protection directory.> You can call the DescribeCloudCenterInstances interface to obtain the UUID of the server.
func (o WebLockConfigOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *WebLockConfig) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type WebLockConfigArrayOutput struct{ *pulumi.OutputState }

func (WebLockConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebLockConfig)(nil)).Elem()
}

func (o WebLockConfigArrayOutput) ToWebLockConfigArrayOutput() WebLockConfigArrayOutput {
	return o
}

func (o WebLockConfigArrayOutput) ToWebLockConfigArrayOutputWithContext(ctx context.Context) WebLockConfigArrayOutput {
	return o
}

func (o WebLockConfigArrayOutput) Index(i pulumi.IntInput) WebLockConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebLockConfig {
		return vs[0].([]*WebLockConfig)[vs[1].(int)]
	}).(WebLockConfigOutput)
}

type WebLockConfigMapOutput struct{ *pulumi.OutputState }

func (WebLockConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebLockConfig)(nil)).Elem()
}

func (o WebLockConfigMapOutput) ToWebLockConfigMapOutput() WebLockConfigMapOutput {
	return o
}

func (o WebLockConfigMapOutput) ToWebLockConfigMapOutputWithContext(ctx context.Context) WebLockConfigMapOutput {
	return o
}

func (o WebLockConfigMapOutput) MapIndex(k pulumi.StringInput) WebLockConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebLockConfig {
		return vs[0].(map[string]*WebLockConfig)[vs[1].(string)]
	}).(WebLockConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebLockConfigInput)(nil)).Elem(), &WebLockConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebLockConfigArrayInput)(nil)).Elem(), WebLockConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebLockConfigMapInput)(nil)).Elem(), WebLockConfigMap{})
	pulumi.RegisterOutputType(WebLockConfigOutput{})
	pulumi.RegisterOutputType(WebLockConfigArrayOutput{})
	pulumi.RegisterOutputType(WebLockConfigMapOutput{})
}
