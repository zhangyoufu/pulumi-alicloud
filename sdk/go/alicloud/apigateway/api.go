// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleGroup, err := apigateway.NewGroup(ctx, "exampleGroup", &apigateway.GroupArgs{
//				Description: pulumi.String("tf-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewApi(ctx, "exampleApi", &apigateway.ApiArgs{
//				GroupId:         exampleGroup.ID(),
//				Description:     pulumi.String("tf-example"),
//				AuthType:        pulumi.String("APP"),
//				ForceNonceCheck: pulumi.Bool(false),
//				RequestConfig: &apigateway.ApiRequestConfigArgs{
//					Protocol: pulumi.String("HTTP"),
//					Method:   pulumi.String("GET"),
//					Path:     pulumi.String("/example/path"),
//					Mode:     pulumi.String("MAPPING"),
//				},
//				ServiceType: pulumi.String("HTTP"),
//				HttpServiceConfig: &apigateway.ApiHttpServiceConfigArgs{
//					Address:  pulumi.String("http://apigateway-backend.alicloudapi.com:8080"),
//					Method:   pulumi.String("GET"),
//					Path:     pulumi.String("/web/cloudapi"),
//					Timeout:  pulumi.Int(12),
//					AoneName: pulumi.String("cloudapi-openapi"),
//				},
//				RequestParameters: apigateway.ApiRequestParameterArray{
//					&apigateway.ApiRequestParameterArgs{
//						Name:        pulumi.String("example"),
//						Type:        pulumi.String("STRING"),
//						Required:    pulumi.String("OPTIONAL"),
//						In:          pulumi.String("QUERY"),
//						InService:   pulumi.String("QUERY"),
//						NameService: pulumi.String("exampleservice"),
//					},
//				},
//				StageNames: pulumi.StringArray{
//					pulumi.String("RELEASE"),
//					pulumi.String("TEST"),
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
// Api gateway api can be imported using the id.Format to `<API Group Id>:<API Id>` e.g.
//
// ```sh
// $ pulumi import alicloud:apigateway/api:Api example "ab2351f2ce904edaa8d92a0510832b91:e4f728fca5a94148b023b99a3e5d0b62"
// ```
type Api struct {
	pulumi.CustomResourceState

	// The ID of the api of api gateway.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// constant_parameters defines the constant parameters of the api. See `constantParameters` below.
	ConstantParameters ApiConstantParameterArrayOutput `pulumi:"constantParameters"`
	// The description of the api. Defaults to null.
	Description pulumi.StringOutput `pulumi:"description"`
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'. See `fcServiceConfig` below.
	FcServiceConfig ApiFcServiceConfigPtrOutput `pulumi:"fcServiceConfig"`
	// Whether to prevent API replay attack. Default value: `false`.
	ForceNonceCheck pulumi.BoolOutput `pulumi:"forceNonceCheck"`
	// The api gateway that the api belongs to. Defaults to null.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// http_service_config defines the config when serviceType selected 'HTTP'. See `httpServiceConfig` below.
	HttpServiceConfig ApiHttpServiceConfigPtrOutput `pulumi:"httpServiceConfig"`
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'. See `httpVpcServiceConfig` below.
	HttpVpcServiceConfig ApiHttpVpcServiceConfigPtrOutput `pulumi:"httpVpcServiceConfig"`
	// http_service_config defines the config when serviceType selected 'MOCK'. See `mockServiceConfig` below.
	MockServiceConfig ApiMockServiceConfigPtrOutput `pulumi:"mockServiceConfig"`
	// The name of the api gateway api. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
	// Request_config defines how users can send requests to your API. See `requestConfig` below.
	RequestConfig ApiRequestConfigOutput `pulumi:"requestConfig"`
	// request_parameters defines the request parameters of the api. See `requestParameters` below.
	RequestParameters ApiRequestParameterArrayOutput `pulumi:"requestParameters"`
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
	StageNames pulumi.StringArrayOutput `pulumi:"stageNames"`
	// system_parameters defines the system parameters of the api. See `systemParameters` below.
	SystemParameters ApiSystemParameterArrayOutput `pulumi:"systemParameters"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthType == nil {
		return nil, errors.New("invalid value for required argument 'AuthType'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.RequestConfig == nil {
		return nil, errors.New("invalid value for required argument 'RequestConfig'")
	}
	if args.ServiceType == nil {
		return nil, errors.New("invalid value for required argument 'ServiceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("alicloud:apigateway/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("alicloud:apigateway/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// The ID of the api of api gateway.
	ApiId *string `pulumi:"apiId"`
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType *string `pulumi:"authType"`
	// constant_parameters defines the constant parameters of the api. See `constantParameters` below.
	ConstantParameters []ApiConstantParameter `pulumi:"constantParameters"`
	// The description of the api. Defaults to null.
	Description *string `pulumi:"description"`
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'. See `fcServiceConfig` below.
	FcServiceConfig *ApiFcServiceConfig `pulumi:"fcServiceConfig"`
	// Whether to prevent API replay attack. Default value: `false`.
	ForceNonceCheck *bool `pulumi:"forceNonceCheck"`
	// The api gateway that the api belongs to. Defaults to null.
	GroupId *string `pulumi:"groupId"`
	// http_service_config defines the config when serviceType selected 'HTTP'. See `httpServiceConfig` below.
	HttpServiceConfig *ApiHttpServiceConfig `pulumi:"httpServiceConfig"`
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'. See `httpVpcServiceConfig` below.
	HttpVpcServiceConfig *ApiHttpVpcServiceConfig `pulumi:"httpVpcServiceConfig"`
	// http_service_config defines the config when serviceType selected 'MOCK'. See `mockServiceConfig` below.
	MockServiceConfig *ApiMockServiceConfig `pulumi:"mockServiceConfig"`
	// The name of the api gateway api. Defaults to null.
	Name *string `pulumi:"name"`
	// Request_config defines how users can send requests to your API. See `requestConfig` below.
	RequestConfig *ApiRequestConfig `pulumi:"requestConfig"`
	// request_parameters defines the request parameters of the api. See `requestParameters` below.
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType *string `pulumi:"serviceType"`
	// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
	StageNames []string `pulumi:"stageNames"`
	// system_parameters defines the system parameters of the api. See `systemParameters` below.
	SystemParameters []ApiSystemParameter `pulumi:"systemParameters"`
}

type ApiState struct {
	// The ID of the api of api gateway.
	ApiId pulumi.StringPtrInput
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType pulumi.StringPtrInput
	// constant_parameters defines the constant parameters of the api. See `constantParameters` below.
	ConstantParameters ApiConstantParameterArrayInput
	// The description of the api. Defaults to null.
	Description pulumi.StringPtrInput
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'. See `fcServiceConfig` below.
	FcServiceConfig ApiFcServiceConfigPtrInput
	// Whether to prevent API replay attack. Default value: `false`.
	ForceNonceCheck pulumi.BoolPtrInput
	// The api gateway that the api belongs to. Defaults to null.
	GroupId pulumi.StringPtrInput
	// http_service_config defines the config when serviceType selected 'HTTP'. See `httpServiceConfig` below.
	HttpServiceConfig ApiHttpServiceConfigPtrInput
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'. See `httpVpcServiceConfig` below.
	HttpVpcServiceConfig ApiHttpVpcServiceConfigPtrInput
	// http_service_config defines the config when serviceType selected 'MOCK'. See `mockServiceConfig` below.
	MockServiceConfig ApiMockServiceConfigPtrInput
	// The name of the api gateway api. Defaults to null.
	Name pulumi.StringPtrInput
	// Request_config defines how users can send requests to your API. See `requestConfig` below.
	RequestConfig ApiRequestConfigPtrInput
	// request_parameters defines the request parameters of the api. See `requestParameters` below.
	RequestParameters ApiRequestParameterArrayInput
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType pulumi.StringPtrInput
	// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
	StageNames pulumi.StringArrayInput
	// system_parameters defines the system parameters of the api. See `systemParameters` below.
	SystemParameters ApiSystemParameterArrayInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType string `pulumi:"authType"`
	// constant_parameters defines the constant parameters of the api. See `constantParameters` below.
	ConstantParameters []ApiConstantParameter `pulumi:"constantParameters"`
	// The description of the api. Defaults to null.
	Description string `pulumi:"description"`
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'. See `fcServiceConfig` below.
	FcServiceConfig *ApiFcServiceConfig `pulumi:"fcServiceConfig"`
	// Whether to prevent API replay attack. Default value: `false`.
	ForceNonceCheck *bool `pulumi:"forceNonceCheck"`
	// The api gateway that the api belongs to. Defaults to null.
	GroupId string `pulumi:"groupId"`
	// http_service_config defines the config when serviceType selected 'HTTP'. See `httpServiceConfig` below.
	HttpServiceConfig *ApiHttpServiceConfig `pulumi:"httpServiceConfig"`
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'. See `httpVpcServiceConfig` below.
	HttpVpcServiceConfig *ApiHttpVpcServiceConfig `pulumi:"httpVpcServiceConfig"`
	// http_service_config defines the config when serviceType selected 'MOCK'. See `mockServiceConfig` below.
	MockServiceConfig *ApiMockServiceConfig `pulumi:"mockServiceConfig"`
	// The name of the api gateway api. Defaults to null.
	Name *string `pulumi:"name"`
	// Request_config defines how users can send requests to your API. See `requestConfig` below.
	RequestConfig ApiRequestConfig `pulumi:"requestConfig"`
	// request_parameters defines the request parameters of the api. See `requestParameters` below.
	RequestParameters []ApiRequestParameter `pulumi:"requestParameters"`
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType string `pulumi:"serviceType"`
	// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
	StageNames []string `pulumi:"stageNames"`
	// system_parameters defines the system parameters of the api. See `systemParameters` below.
	SystemParameters []ApiSystemParameter `pulumi:"systemParameters"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// The authorization Type including APP and ANONYMOUS. Defaults to null.
	AuthType pulumi.StringInput
	// constant_parameters defines the constant parameters of the api. See `constantParameters` below.
	ConstantParameters ApiConstantParameterArrayInput
	// The description of the api. Defaults to null.
	Description pulumi.StringInput
	// fc_service_config defines the config when serviceType selected 'FunctionCompute'. See `fcServiceConfig` below.
	FcServiceConfig ApiFcServiceConfigPtrInput
	// Whether to prevent API replay attack. Default value: `false`.
	ForceNonceCheck pulumi.BoolPtrInput
	// The api gateway that the api belongs to. Defaults to null.
	GroupId pulumi.StringInput
	// http_service_config defines the config when serviceType selected 'HTTP'. See `httpServiceConfig` below.
	HttpServiceConfig ApiHttpServiceConfigPtrInput
	// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'. See `httpVpcServiceConfig` below.
	HttpVpcServiceConfig ApiHttpVpcServiceConfigPtrInput
	// http_service_config defines the config when serviceType selected 'MOCK'. See `mockServiceConfig` below.
	MockServiceConfig ApiMockServiceConfigPtrInput
	// The name of the api gateway api. Defaults to null.
	Name pulumi.StringPtrInput
	// Request_config defines how users can send requests to your API. See `requestConfig` below.
	RequestConfig ApiRequestConfigInput
	// request_parameters defines the request parameters of the api. See `requestParameters` below.
	RequestParameters ApiRequestParameterArrayInput
	// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
	ServiceType pulumi.StringInput
	// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
	StageNames pulumi.StringArrayInput
	// system_parameters defines the system parameters of the api. See `systemParameters` below.
	SystemParameters ApiSystemParameterArrayInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

// ApiArrayInput is an input type that accepts ApiArray and ApiArrayOutput values.
// You can construct a concrete instance of `ApiArrayInput` via:
//
//	ApiArray{ ApiArgs{...} }
type ApiArrayInput interface {
	pulumi.Input

	ToApiArrayOutput() ApiArrayOutput
	ToApiArrayOutputWithContext(context.Context) ApiArrayOutput
}

type ApiArray []ApiInput

func (ApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (i ApiArray) ToApiArrayOutput() ApiArrayOutput {
	return i.ToApiArrayOutputWithContext(context.Background())
}

func (i ApiArray) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiArrayOutput)
}

// ApiMapInput is an input type that accepts ApiMap and ApiMapOutput values.
// You can construct a concrete instance of `ApiMapInput` via:
//
//	ApiMap{ "key": ApiArgs{...} }
type ApiMapInput interface {
	pulumi.Input

	ToApiMapOutput() ApiMapOutput
	ToApiMapOutputWithContext(context.Context) ApiMapOutput
}

type ApiMap map[string]ApiInput

func (ApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (i ApiMap) ToApiMapOutput() ApiMapOutput {
	return i.ToApiMapOutputWithContext(context.Background())
}

func (i ApiMap) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMapOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// The ID of the api of api gateway.
func (o ApiOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The authorization Type including APP and ANONYMOUS. Defaults to null.
func (o ApiOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// constant_parameters defines the constant parameters of the api. See `constantParameters` below.
func (o ApiOutput) ConstantParameters() ApiConstantParameterArrayOutput {
	return o.ApplyT(func(v *Api) ApiConstantParameterArrayOutput { return v.ConstantParameters }).(ApiConstantParameterArrayOutput)
}

// The description of the api. Defaults to null.
func (o ApiOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// fc_service_config defines the config when serviceType selected 'FunctionCompute'. See `fcServiceConfig` below.
func (o ApiOutput) FcServiceConfig() ApiFcServiceConfigPtrOutput {
	return o.ApplyT(func(v *Api) ApiFcServiceConfigPtrOutput { return v.FcServiceConfig }).(ApiFcServiceConfigPtrOutput)
}

// Whether to prevent API replay attack. Default value: `false`.
func (o ApiOutput) ForceNonceCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolOutput { return v.ForceNonceCheck }).(pulumi.BoolOutput)
}

// The api gateway that the api belongs to. Defaults to null.
func (o ApiOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// http_service_config defines the config when serviceType selected 'HTTP'. See `httpServiceConfig` below.
func (o ApiOutput) HttpServiceConfig() ApiHttpServiceConfigPtrOutput {
	return o.ApplyT(func(v *Api) ApiHttpServiceConfigPtrOutput { return v.HttpServiceConfig }).(ApiHttpServiceConfigPtrOutput)
}

// http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'. See `httpVpcServiceConfig` below.
func (o ApiOutput) HttpVpcServiceConfig() ApiHttpVpcServiceConfigPtrOutput {
	return o.ApplyT(func(v *Api) ApiHttpVpcServiceConfigPtrOutput { return v.HttpVpcServiceConfig }).(ApiHttpVpcServiceConfigPtrOutput)
}

// http_service_config defines the config when serviceType selected 'MOCK'. See `mockServiceConfig` below.
func (o ApiOutput) MockServiceConfig() ApiMockServiceConfigPtrOutput {
	return o.ApplyT(func(v *Api) ApiMockServiceConfigPtrOutput { return v.MockServiceConfig }).(ApiMockServiceConfigPtrOutput)
}

// The name of the api gateway api. Defaults to null.
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Request_config defines how users can send requests to your API. See `requestConfig` below.
func (o ApiOutput) RequestConfig() ApiRequestConfigOutput {
	return o.ApplyT(func(v *Api) ApiRequestConfigOutput { return v.RequestConfig }).(ApiRequestConfigOutput)
}

// request_parameters defines the request parameters of the api. See `requestParameters` below.
func (o ApiOutput) RequestParameters() ApiRequestParameterArrayOutput {
	return o.ApplyT(func(v *Api) ApiRequestParameterArrayOutput { return v.RequestParameters }).(ApiRequestParameterArrayOutput)
}

// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
func (o ApiOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
func (o ApiOutput) StageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Api) pulumi.StringArrayOutput { return v.StageNames }).(pulumi.StringArrayOutput)
}

// system_parameters defines the system parameters of the api. See `systemParameters` below.
func (o ApiOutput) SystemParameters() ApiSystemParameterArrayOutput {
	return o.ApplyT(func(v *Api) ApiSystemParameterArrayOutput { return v.SystemParameters }).(ApiSystemParameterArrayOutput)
}

type ApiArrayOutput struct{ *pulumi.OutputState }

func (ApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (o ApiArrayOutput) ToApiArrayOutput() ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) Index(i pulumi.IntInput) ApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Api {
		return vs[0].([]*Api)[vs[1].(int)]
	}).(ApiOutput)
}

type ApiMapOutput struct{ *pulumi.OutputState }

func (ApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (o ApiMapOutput) ToApiMapOutput() ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return o
}

func (o ApiMapOutput) MapIndex(k pulumi.StringInput) ApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Api {
		return vs[0].(map[string]*Api)[vs[1].(string)]
	}).(ApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiArrayInput)(nil)).Elem(), ApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiMapInput)(nil)).Elem(), ApiMap{})
	pulumi.RegisterOutputType(ApiOutput{})
	pulumi.RegisterOutputType(ApiArrayOutput{})
	pulumi.RegisterOutputType(ApiMapOutput{})
}
