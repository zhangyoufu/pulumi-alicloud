// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates an EDAS ecs application on EDAS, see [What is EDAS Application](https://www.alibabacloud.com/help/en/edas/developer-reference/api-edas-2017-08-01-insertapplication). The application will be deployed when `groupId` and `warUrl` are given.
//
// > **NOTE:** Available since v1.82.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/edas"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
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
//			defaultRegions, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultCluster, err := edas.NewCluster(ctx, "defaultCluster", &edas.ClusterArgs{
//				ClusterName:     pulumi.String(name),
//				ClusterType:     pulumi.Int(2),
//				NetworkMode:     pulumi.Int(2),
//				LogicalRegionId: *pulumi.String(defaultRegions.Regions[0].Id),
//				VpcId:           defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = edas.NewApplication(ctx, "defaultApplication", &edas.ApplicationArgs{
//				ApplicationName: pulumi.String(name),
//				ClusterId:       defaultCluster.ID(),
//				PackageType:     pulumi.String("JAR"),
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
// EDAS application can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:edas/application:Application app app_Id
//
// ```
type Application struct {
	pulumi.CustomResourceState

	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId pulumi.IntPtrOutput `pulumi:"buildPackId"`
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The description of the application that you want to create.
	Descriotion pulumi.StringPtrOutput `pulumi:"descriotion"`
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayOutput `pulumi:"ecuInfos"`
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// The URL for health checking of the application.
	HealthCheckUrl pulumi.StringPtrOutput `pulumi:"healthCheckUrl"`
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId pulumi.StringPtrOutput `pulumi:"logicalRegionId"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion pulumi.StringPtrOutput `pulumi:"packageVersion"`
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl pulumi.StringPtrOutput `pulumi:"warUrl"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.PackageType == nil {
		return nil, errors.New("invalid value for required argument 'PackageType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("alicloud:edas/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("alicloud:edas/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName *string `pulumi:"applicationName"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId *int `pulumi:"buildPackId"`
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId *string `pulumi:"clusterId"`
	// The description of the application that you want to create.
	Descriotion *string `pulumi:"descriotion"`
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos []string `pulumi:"ecuInfos"`
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId *string `pulumi:"groupId"`
	// The URL for health checking of the application.
	HealthCheckUrl *string `pulumi:"healthCheckUrl"`
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId *string `pulumi:"logicalRegionId"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType *string `pulumi:"packageType"`
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion *string `pulumi:"packageVersion"`
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl *string `pulumi:"warUrl"`
}

type ApplicationState struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName pulumi.StringPtrInput
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId pulumi.IntPtrInput
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId pulumi.StringPtrInput
	// The description of the application that you want to create.
	Descriotion pulumi.StringPtrInput
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayInput
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId pulumi.StringPtrInput
	// The URL for health checking of the application.
	HealthCheckUrl pulumi.StringPtrInput
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId pulumi.StringPtrInput
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType pulumi.StringPtrInput
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion pulumi.StringPtrInput
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName string `pulumi:"applicationName"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId *int `pulumi:"buildPackId"`
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId string `pulumi:"clusterId"`
	// The description of the application that you want to create.
	Descriotion *string `pulumi:"descriotion"`
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos []string `pulumi:"ecuInfos"`
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId *string `pulumi:"groupId"`
	// The URL for health checking of the application.
	HealthCheckUrl *string `pulumi:"healthCheckUrl"`
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId *string `pulumi:"logicalRegionId"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType string `pulumi:"packageType"`
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion *string `pulumi:"packageVersion"`
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl *string `pulumi:"warUrl"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	ApplicationName pulumi.StringInput
	// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
	BuildPackId pulumi.IntPtrInput
	// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
	ClusterId pulumi.StringInput
	// The description of the application that you want to create.
	Descriotion pulumi.StringPtrInput
	// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
	EcuInfos pulumi.StringArrayInput
	// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
	GroupId pulumi.StringPtrInput
	// The URL for health checking of the application.
	HealthCheckUrl pulumi.StringPtrInput
	// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
	LogicalRegionId pulumi.StringPtrInput
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	PackageType pulumi.StringInput
	// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
	PackageVersion pulumi.StringPtrInput
	// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
	WarUrl pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

func (i *Application) ToOutput(ctx context.Context) pulumix.Output[*Application] {
	return pulumix.Output[*Application]{
		OutputState: i.ToApplicationOutputWithContext(ctx).OutputState,
	}
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

func (i ApplicationArray) ToOutput(ctx context.Context) pulumix.Output[[]*Application] {
	return pulumix.Output[[]*Application]{
		OutputState: i.ToApplicationArrayOutputWithContext(ctx).OutputState,
	}
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

func (i ApplicationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Application] {
	return pulumix.Output[map[string]*Application]{
		OutputState: i.ToApplicationMapOutputWithContext(ctx).OutputState,
	}
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToOutput(ctx context.Context) pulumix.Output[*Application] {
	return pulumix.Output[*Application]{
		OutputState: o.OutputState,
	}
}

// Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
func (o ApplicationOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
func (o ApplicationOutput) BuildPackId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.IntPtrOutput { return v.BuildPackId }).(pulumi.IntPtrOutput)
}

// The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
func (o ApplicationOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The description of the application that you want to create.
func (o ApplicationOutput) Descriotion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Descriotion }).(pulumi.StringPtrOutput)
}

// The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
func (o ApplicationOutput) EcuInfos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.EcuInfos }).(pulumi.StringArrayOutput)
}

// The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
func (o ApplicationOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The URL for health checking of the application.
func (o ApplicationOutput) HealthCheckUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.HealthCheckUrl }).(pulumi.StringPtrOutput)
}

// The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
func (o ApplicationOutput) LogicalRegionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.LogicalRegionId }).(pulumi.StringPtrOutput)
}

// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
func (o ApplicationOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
func (o ApplicationOutput) PackageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.PackageVersion }).(pulumi.StringPtrOutput)
}

// The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
func (o ApplicationOutput) WarUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.WarUrl }).(pulumi.StringPtrOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Application] {
	return pulumix.Output[[]*Application]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Application] {
	return pulumix.Output[map[string]*Application]{
		OutputState: o.OutputState,
	}
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}
