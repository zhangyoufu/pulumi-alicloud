// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CR Chart Repository resource.
//
// For information about CR Chart Repository and how to use it, see [What is Chart Repository](https://www.alibabacloud.com/help/doc-detail/145318.htm).
//
// > **NOTE:** Available in v1.149.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultRegistryEnterpriseInstance, err := cr.NewRegistryEnterpriseInstance(ctx, "defaultRegistryEnterpriseInstance", &cr.RegistryEnterpriseInstanceArgs{
// 			PaymentType:  pulumi.String("Subscription"),
// 			Period:       pulumi.Int(1),
// 			InstanceType: pulumi.String("Advanced"),
// 			InstanceName: pulumi.String("name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultChartNamespace, err := cr.NewChartNamespace(ctx, "defaultChartNamespace", &cr.ChartNamespaceArgs{
// 			InstanceId:    defaultRegistryEnterpriseInstance.ID(),
// 			NamespaceName: pulumi.String("name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cr.NewChartRepository(ctx, "defaultChartRepository", &cr.ChartRepositoryArgs{
// 			RepoNamespaceName: defaultChartNamespace.NamespaceName,
// 			InstanceId:        pulumi.Any(local.Instance),
// 			RepoName:          pulumi.String("repo_name"),
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
// CR Chart Repository can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cr/chartRepository:ChartRepository example <instance_id>:<repo_namespace_name>:<repo_name>
// ```
type ChartRepository struct {
	pulumi.CustomResourceState

	// The ID of the Container Registry instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the repository that you want to create.
	RepoName pulumi.StringOutput `pulumi:"repoName"`
	// The namespace to which the repository belongs.
	RepoNamespaceName pulumi.StringOutput `pulumi:"repoNamespaceName"`
	// The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
	RepoType pulumi.StringOutput `pulumi:"repoType"`
	// The summary about the repository.
	Summary pulumi.StringPtrOutput `pulumi:"summary"`
}

// NewChartRepository registers a new resource with the given unique name, arguments, and options.
func NewChartRepository(ctx *pulumi.Context,
	name string, args *ChartRepositoryArgs, opts ...pulumi.ResourceOption) (*ChartRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.RepoName == nil {
		return nil, errors.New("invalid value for required argument 'RepoName'")
	}
	if args.RepoNamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'RepoNamespaceName'")
	}
	var resource ChartRepository
	err := ctx.RegisterResource("alicloud:cr/chartRepository:ChartRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChartRepository gets an existing ChartRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChartRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChartRepositoryState, opts ...pulumi.ResourceOption) (*ChartRepository, error) {
	var resource ChartRepository
	err := ctx.ReadResource("alicloud:cr/chartRepository:ChartRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChartRepository resources.
type chartRepositoryState struct {
	// The ID of the Container Registry instance.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the repository that you want to create.
	RepoName *string `pulumi:"repoName"`
	// The namespace to which the repository belongs.
	RepoNamespaceName *string `pulumi:"repoNamespaceName"`
	// The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
	RepoType *string `pulumi:"repoType"`
	// The summary about the repository.
	Summary *string `pulumi:"summary"`
}

type ChartRepositoryState struct {
	// The ID of the Container Registry instance.
	InstanceId pulumi.StringPtrInput
	// The name of the repository that you want to create.
	RepoName pulumi.StringPtrInput
	// The namespace to which the repository belongs.
	RepoNamespaceName pulumi.StringPtrInput
	// The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
	RepoType pulumi.StringPtrInput
	// The summary about the repository.
	Summary pulumi.StringPtrInput
}

func (ChartRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*chartRepositoryState)(nil)).Elem()
}

type chartRepositoryArgs struct {
	// The ID of the Container Registry instance.
	InstanceId string `pulumi:"instanceId"`
	// The name of the repository that you want to create.
	RepoName string `pulumi:"repoName"`
	// The namespace to which the repository belongs.
	RepoNamespaceName string `pulumi:"repoNamespaceName"`
	// The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
	RepoType *string `pulumi:"repoType"`
	// The summary about the repository.
	Summary *string `pulumi:"summary"`
}

// The set of arguments for constructing a ChartRepository resource.
type ChartRepositoryArgs struct {
	// The ID of the Container Registry instance.
	InstanceId pulumi.StringInput
	// The name of the repository that you want to create.
	RepoName pulumi.StringInput
	// The namespace to which the repository belongs.
	RepoNamespaceName pulumi.StringInput
	// The default repository type. Valid values: `PUBLIC`,`PRIVATE`.
	RepoType pulumi.StringPtrInput
	// The summary about the repository.
	Summary pulumi.StringPtrInput
}

func (ChartRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chartRepositoryArgs)(nil)).Elem()
}

type ChartRepositoryInput interface {
	pulumi.Input

	ToChartRepositoryOutput() ChartRepositoryOutput
	ToChartRepositoryOutputWithContext(ctx context.Context) ChartRepositoryOutput
}

func (*ChartRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*ChartRepository)(nil))
}

func (i *ChartRepository) ToChartRepositoryOutput() ChartRepositoryOutput {
	return i.ToChartRepositoryOutputWithContext(context.Background())
}

func (i *ChartRepository) ToChartRepositoryOutputWithContext(ctx context.Context) ChartRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChartRepositoryOutput)
}

func (i *ChartRepository) ToChartRepositoryPtrOutput() ChartRepositoryPtrOutput {
	return i.ToChartRepositoryPtrOutputWithContext(context.Background())
}

func (i *ChartRepository) ToChartRepositoryPtrOutputWithContext(ctx context.Context) ChartRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChartRepositoryPtrOutput)
}

type ChartRepositoryPtrInput interface {
	pulumi.Input

	ToChartRepositoryPtrOutput() ChartRepositoryPtrOutput
	ToChartRepositoryPtrOutputWithContext(ctx context.Context) ChartRepositoryPtrOutput
}

type chartRepositoryPtrType ChartRepositoryArgs

func (*chartRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ChartRepository)(nil))
}

func (i *chartRepositoryPtrType) ToChartRepositoryPtrOutput() ChartRepositoryPtrOutput {
	return i.ToChartRepositoryPtrOutputWithContext(context.Background())
}

func (i *chartRepositoryPtrType) ToChartRepositoryPtrOutputWithContext(ctx context.Context) ChartRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChartRepositoryPtrOutput)
}

// ChartRepositoryArrayInput is an input type that accepts ChartRepositoryArray and ChartRepositoryArrayOutput values.
// You can construct a concrete instance of `ChartRepositoryArrayInput` via:
//
//          ChartRepositoryArray{ ChartRepositoryArgs{...} }
type ChartRepositoryArrayInput interface {
	pulumi.Input

	ToChartRepositoryArrayOutput() ChartRepositoryArrayOutput
	ToChartRepositoryArrayOutputWithContext(context.Context) ChartRepositoryArrayOutput
}

type ChartRepositoryArray []ChartRepositoryInput

func (ChartRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ChartRepository)(nil)).Elem()
}

func (i ChartRepositoryArray) ToChartRepositoryArrayOutput() ChartRepositoryArrayOutput {
	return i.ToChartRepositoryArrayOutputWithContext(context.Background())
}

func (i ChartRepositoryArray) ToChartRepositoryArrayOutputWithContext(ctx context.Context) ChartRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChartRepositoryArrayOutput)
}

// ChartRepositoryMapInput is an input type that accepts ChartRepositoryMap and ChartRepositoryMapOutput values.
// You can construct a concrete instance of `ChartRepositoryMapInput` via:
//
//          ChartRepositoryMap{ "key": ChartRepositoryArgs{...} }
type ChartRepositoryMapInput interface {
	pulumi.Input

	ToChartRepositoryMapOutput() ChartRepositoryMapOutput
	ToChartRepositoryMapOutputWithContext(context.Context) ChartRepositoryMapOutput
}

type ChartRepositoryMap map[string]ChartRepositoryInput

func (ChartRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ChartRepository)(nil)).Elem()
}

func (i ChartRepositoryMap) ToChartRepositoryMapOutput() ChartRepositoryMapOutput {
	return i.ToChartRepositoryMapOutputWithContext(context.Background())
}

func (i ChartRepositoryMap) ToChartRepositoryMapOutputWithContext(ctx context.Context) ChartRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChartRepositoryMapOutput)
}

type ChartRepositoryOutput struct{ *pulumi.OutputState }

func (ChartRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChartRepository)(nil))
}

func (o ChartRepositoryOutput) ToChartRepositoryOutput() ChartRepositoryOutput {
	return o
}

func (o ChartRepositoryOutput) ToChartRepositoryOutputWithContext(ctx context.Context) ChartRepositoryOutput {
	return o
}

func (o ChartRepositoryOutput) ToChartRepositoryPtrOutput() ChartRepositoryPtrOutput {
	return o.ToChartRepositoryPtrOutputWithContext(context.Background())
}

func (o ChartRepositoryOutput) ToChartRepositoryPtrOutputWithContext(ctx context.Context) ChartRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChartRepository) *ChartRepository {
		return &v
	}).(ChartRepositoryPtrOutput)
}

type ChartRepositoryPtrOutput struct{ *pulumi.OutputState }

func (ChartRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChartRepository)(nil))
}

func (o ChartRepositoryPtrOutput) ToChartRepositoryPtrOutput() ChartRepositoryPtrOutput {
	return o
}

func (o ChartRepositoryPtrOutput) ToChartRepositoryPtrOutputWithContext(ctx context.Context) ChartRepositoryPtrOutput {
	return o
}

func (o ChartRepositoryPtrOutput) Elem() ChartRepositoryOutput {
	return o.ApplyT(func(v *ChartRepository) ChartRepository {
		if v != nil {
			return *v
		}
		var ret ChartRepository
		return ret
	}).(ChartRepositoryOutput)
}

type ChartRepositoryArrayOutput struct{ *pulumi.OutputState }

func (ChartRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChartRepository)(nil))
}

func (o ChartRepositoryArrayOutput) ToChartRepositoryArrayOutput() ChartRepositoryArrayOutput {
	return o
}

func (o ChartRepositoryArrayOutput) ToChartRepositoryArrayOutputWithContext(ctx context.Context) ChartRepositoryArrayOutput {
	return o
}

func (o ChartRepositoryArrayOutput) Index(i pulumi.IntInput) ChartRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChartRepository {
		return vs[0].([]ChartRepository)[vs[1].(int)]
	}).(ChartRepositoryOutput)
}

type ChartRepositoryMapOutput struct{ *pulumi.OutputState }

func (ChartRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ChartRepository)(nil))
}

func (o ChartRepositoryMapOutput) ToChartRepositoryMapOutput() ChartRepositoryMapOutput {
	return o
}

func (o ChartRepositoryMapOutput) ToChartRepositoryMapOutputWithContext(ctx context.Context) ChartRepositoryMapOutput {
	return o
}

func (o ChartRepositoryMapOutput) MapIndex(k pulumi.StringInput) ChartRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ChartRepository {
		return vs[0].(map[string]ChartRepository)[vs[1].(string)]
	}).(ChartRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChartRepositoryInput)(nil)).Elem(), &ChartRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChartRepositoryPtrInput)(nil)).Elem(), &ChartRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChartRepositoryArrayInput)(nil)).Elem(), ChartRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChartRepositoryMapInput)(nil)).Elem(), ChartRepositoryMap{})
	pulumi.RegisterOutputType(ChartRepositoryOutput{})
	pulumi.RegisterOutputType(ChartRepositoryPtrOutput{})
	pulumi.RegisterOutputType(ChartRepositoryArrayOutput{})
	pulumi.RegisterOutputType(ChartRepositoryMapOutput{})
}
