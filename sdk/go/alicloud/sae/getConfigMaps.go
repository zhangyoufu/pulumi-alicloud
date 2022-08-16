// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Sae Config Maps of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.130.0+.
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
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sae"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			configMapName := "examplename"
//			if param := cfg.Get("configMapName"); param != "" {
//				configMapName = param
//			}
//			exampleNamespace, err := sae.NewNamespace(ctx, "exampleNamespace", &sae.NamespaceArgs{
//				NamespaceId:          pulumi.String("cn-hangzhou:yourname"),
//				NamespaceName:        pulumi.String("example_value"),
//				NamespaceDescription: pulumi.String("your_description"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"env.home":  "/root",
//				"env.shell": "/bin/sh",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = sae.NewConfigMap(ctx, "exampleConfigMap", &sae.ConfigMapArgs{
//				Data:        pulumi.String(json0),
//				NamespaceId: exampleNamespace.NamespaceId,
//			})
//			if err != nil {
//				return err
//			}
//			nameRegex := sae.GetConfigMapsOutput(ctx, sae.GetConfigMapsOutputArgs{
//				NamespaceId: exampleNamespace.NamespaceId,
//				NameRegex:   pulumi.String("^example"),
//			}, nil)
//			ctx.Export("saeConfigMapId", nameRegex.ApplyT(func(nameRegex sae.GetConfigMapsResult) (string, error) {
//				return nameRegex.Maps[0].Id, nil
//			}).(pulumi.StringOutput))
//			return nil
//		})
//	}
//
// ```
func GetConfigMaps(ctx *pulumi.Context, args *GetConfigMapsArgs, opts ...pulumi.InvokeOption) (*GetConfigMapsResult, error) {
	var rv GetConfigMapsResult
	err := ctx.Invoke("alicloud:sae/getConfigMaps:getConfigMaps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfigMaps.
type GetConfigMapsArgs struct {
	// A list of Config Map IDs.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by Config Map name.
	NameRegex *string `pulumi:"nameRegex"`
	// The NamespaceId of Config Maps.
	NamespaceId string  `pulumi:"namespaceId"`
	OutputFile  *string `pulumi:"outputFile"`
}

// A collection of values returned by getConfigMaps.
type GetConfigMapsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string             `pulumi:"id"`
	Ids         []string           `pulumi:"ids"`
	Maps        []GetConfigMapsMap `pulumi:"maps"`
	NameRegex   *string            `pulumi:"nameRegex"`
	Names       []string           `pulumi:"names"`
	NamespaceId string             `pulumi:"namespaceId"`
	OutputFile  *string            `pulumi:"outputFile"`
}

func GetConfigMapsOutput(ctx *pulumi.Context, args GetConfigMapsOutputArgs, opts ...pulumi.InvokeOption) GetConfigMapsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetConfigMapsResult, error) {
			args := v.(GetConfigMapsArgs)
			r, err := GetConfigMaps(ctx, &args, opts...)
			var s GetConfigMapsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetConfigMapsResultOutput)
}

// A collection of arguments for invoking getConfigMaps.
type GetConfigMapsOutputArgs struct {
	// A list of Config Map IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A regex string to filter results by Config Map name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The NamespaceId of Config Maps.
	NamespaceId pulumi.StringInput    `pulumi:"namespaceId"`
	OutputFile  pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetConfigMapsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigMapsArgs)(nil)).Elem()
}

// A collection of values returned by getConfigMaps.
type GetConfigMapsResultOutput struct{ *pulumi.OutputState }

func (GetConfigMapsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigMapsResult)(nil)).Elem()
}

func (o GetConfigMapsResultOutput) ToGetConfigMapsResultOutput() GetConfigMapsResultOutput {
	return o
}

func (o GetConfigMapsResultOutput) ToGetConfigMapsResultOutputWithContext(ctx context.Context) GetConfigMapsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetConfigMapsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigMapsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetConfigMapsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetConfigMapsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetConfigMapsResultOutput) Maps() GetConfigMapsMapArrayOutput {
	return o.ApplyT(func(v GetConfigMapsResult) []GetConfigMapsMap { return v.Maps }).(GetConfigMapsMapArrayOutput)
}

func (o GetConfigMapsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConfigMapsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetConfigMapsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetConfigMapsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetConfigMapsResultOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigMapsResult) string { return v.NamespaceId }).(pulumi.StringOutput)
}

func (o GetConfigMapsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConfigMapsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConfigMapsResultOutput{})
}
