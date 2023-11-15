// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource will help you to manager Container Registry Enterprise Edition repositories.
//
// For information about Container Registry Enterprise Edition repository and how to use it, see [Create a Repository](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createrepository)
//
// > **NOTE:** Available since v1.86.0.
//
// > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cr"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
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
//			exampleRegistryEnterpriseInstance, err := cr.NewRegistryEnterpriseInstance(ctx, "exampleRegistryEnterpriseInstance", &cr.RegistryEnterpriseInstanceArgs{
//				PaymentType:   pulumi.String("Subscription"),
//				Period:        pulumi.Int(1),
//				RenewPeriod:   pulumi.Int(0),
//				RenewalStatus: pulumi.String("ManualRenewal"),
//				InstanceType:  pulumi.String("Advanced"),
//				InstanceName:  pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			exampleRegistryEnterpriseNamespace, err := cs.NewRegistryEnterpriseNamespace(ctx, "exampleRegistryEnterpriseNamespace", &cs.RegistryEnterpriseNamespaceArgs{
//				InstanceId:        exampleRegistryEnterpriseInstance.ID(),
//				AutoCreate:        pulumi.Bool(false),
//				DefaultVisibility: pulumi.String("PUBLIC"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cs.NewRegistryEnterpriseRepo(ctx, "exampleRegistryEnterpriseRepo", &cs.RegistryEnterpriseRepoArgs{
//				InstanceId: exampleRegistryEnterpriseInstance.ID(),
//				Namespace:  exampleRegistryEnterpriseNamespace.Name,
//				Summary:    pulumi.String("this is summary of my new repo"),
//				RepoType:   pulumi.String("PUBLIC"),
//				Detail:     pulumi.String("this is a public repo"),
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
// Container Registry Enterprise Edition repository can be imported using the `{instance_id}:{namespace}:{repository}`, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo default `cri-xxx:my-namespace:my-repo`
//
// ```
type RegistryEnterpriseRepo struct {
	pulumi.CustomResourceState

	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail pulumi.StringPtrOutput `pulumi:"detail"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The uuid of Container Registry Enterprise Edition repository.
	RepoId pulumi.StringOutput `pulumi:"repoId"`
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType pulumi.StringOutput `pulumi:"repoType"`
	// The repository general information. It can contain 1 to 100 characters.
	Summary pulumi.StringOutput `pulumi:"summary"`
}

// NewRegistryEnterpriseRepo registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnterpriseRepo(ctx *pulumi.Context,
	name string, args *RegistryEnterpriseRepoArgs, opts ...pulumi.ResourceOption) (*RegistryEnterpriseRepo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.RepoType == nil {
		return nil, errors.New("invalid value for required argument 'RepoType'")
	}
	if args.Summary == nil {
		return nil, errors.New("invalid value for required argument 'Summary'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegistryEnterpriseRepo
	err := ctx.RegisterResource("alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnterpriseRepo gets an existing RegistryEnterpriseRepo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnterpriseRepo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnterpriseRepoState, opts ...pulumi.ResourceOption) (*RegistryEnterpriseRepo, error) {
	var resource RegistryEnterpriseRepo
	err := ctx.ReadResource("alicloud:cs/registryEnterpriseRepo:RegistryEnterpriseRepo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnterpriseRepo resources.
type registryEnterpriseRepoState struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail *string `pulumi:"detail"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId *string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
	Name *string `pulumi:"name"`
	// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
	Namespace *string `pulumi:"namespace"`
	// The uuid of Container Registry Enterprise Edition repository.
	RepoId *string `pulumi:"repoId"`
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType *string `pulumi:"repoType"`
	// The repository general information. It can contain 1 to 100 characters.
	Summary *string `pulumi:"summary"`
}

type RegistryEnterpriseRepoState struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail pulumi.StringPtrInput
	// ID of Container Registry Enterprise Edition instance.
	InstanceId pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
	Name pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
	Namespace pulumi.StringPtrInput
	// The uuid of Container Registry Enterprise Edition repository.
	RepoId pulumi.StringPtrInput
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType pulumi.StringPtrInput
	// The repository general information. It can contain 1 to 100 characters.
	Summary pulumi.StringPtrInput
}

func (RegistryEnterpriseRepoState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseRepoState)(nil)).Elem()
}

type registryEnterpriseRepoArgs struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail *string `pulumi:"detail"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
	Name *string `pulumi:"name"`
	// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
	Namespace string `pulumi:"namespace"`
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType string `pulumi:"repoType"`
	// The repository general information. It can contain 1 to 100 characters.
	Summary string `pulumi:"summary"`
}

// The set of arguments for constructing a RegistryEnterpriseRepo resource.
type RegistryEnterpriseRepoArgs struct {
	// The repository specific information. MarkDown format is supported, and the length limit is 2000.
	Detail pulumi.StringPtrInput
	// ID of Container Registry Enterprise Edition instance.
	InstanceId pulumi.StringInput
	// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
	Name pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
	Namespace pulumi.StringInput
	// `PUBLIC` or `PRIVATE`, repo's visibility.
	RepoType pulumi.StringInput
	// The repository general information. It can contain 1 to 100 characters.
	Summary pulumi.StringInput
}

func (RegistryEnterpriseRepoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseRepoArgs)(nil)).Elem()
}

type RegistryEnterpriseRepoInput interface {
	pulumi.Input

	ToRegistryEnterpriseRepoOutput() RegistryEnterpriseRepoOutput
	ToRegistryEnterpriseRepoOutputWithContext(ctx context.Context) RegistryEnterpriseRepoOutput
}

func (*RegistryEnterpriseRepo) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseRepo)(nil)).Elem()
}

func (i *RegistryEnterpriseRepo) ToRegistryEnterpriseRepoOutput() RegistryEnterpriseRepoOutput {
	return i.ToRegistryEnterpriseRepoOutputWithContext(context.Background())
}

func (i *RegistryEnterpriseRepo) ToRegistryEnterpriseRepoOutputWithContext(ctx context.Context) RegistryEnterpriseRepoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseRepoOutput)
}

// RegistryEnterpriseRepoArrayInput is an input type that accepts RegistryEnterpriseRepoArray and RegistryEnterpriseRepoArrayOutput values.
// You can construct a concrete instance of `RegistryEnterpriseRepoArrayInput` via:
//
//	RegistryEnterpriseRepoArray{ RegistryEnterpriseRepoArgs{...} }
type RegistryEnterpriseRepoArrayInput interface {
	pulumi.Input

	ToRegistryEnterpriseRepoArrayOutput() RegistryEnterpriseRepoArrayOutput
	ToRegistryEnterpriseRepoArrayOutputWithContext(context.Context) RegistryEnterpriseRepoArrayOutput
}

type RegistryEnterpriseRepoArray []RegistryEnterpriseRepoInput

func (RegistryEnterpriseRepoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseRepo)(nil)).Elem()
}

func (i RegistryEnterpriseRepoArray) ToRegistryEnterpriseRepoArrayOutput() RegistryEnterpriseRepoArrayOutput {
	return i.ToRegistryEnterpriseRepoArrayOutputWithContext(context.Background())
}

func (i RegistryEnterpriseRepoArray) ToRegistryEnterpriseRepoArrayOutputWithContext(ctx context.Context) RegistryEnterpriseRepoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseRepoArrayOutput)
}

// RegistryEnterpriseRepoMapInput is an input type that accepts RegistryEnterpriseRepoMap and RegistryEnterpriseRepoMapOutput values.
// You can construct a concrete instance of `RegistryEnterpriseRepoMapInput` via:
//
//	RegistryEnterpriseRepoMap{ "key": RegistryEnterpriseRepoArgs{...} }
type RegistryEnterpriseRepoMapInput interface {
	pulumi.Input

	ToRegistryEnterpriseRepoMapOutput() RegistryEnterpriseRepoMapOutput
	ToRegistryEnterpriseRepoMapOutputWithContext(context.Context) RegistryEnterpriseRepoMapOutput
}

type RegistryEnterpriseRepoMap map[string]RegistryEnterpriseRepoInput

func (RegistryEnterpriseRepoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseRepo)(nil)).Elem()
}

func (i RegistryEnterpriseRepoMap) ToRegistryEnterpriseRepoMapOutput() RegistryEnterpriseRepoMapOutput {
	return i.ToRegistryEnterpriseRepoMapOutputWithContext(context.Background())
}

func (i RegistryEnterpriseRepoMap) ToRegistryEnterpriseRepoMapOutputWithContext(ctx context.Context) RegistryEnterpriseRepoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryEnterpriseRepoMapOutput)
}

type RegistryEnterpriseRepoOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseRepoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryEnterpriseRepo)(nil)).Elem()
}

func (o RegistryEnterpriseRepoOutput) ToRegistryEnterpriseRepoOutput() RegistryEnterpriseRepoOutput {
	return o
}

func (o RegistryEnterpriseRepoOutput) ToRegistryEnterpriseRepoOutputWithContext(ctx context.Context) RegistryEnterpriseRepoOutput {
	return o
}

// The repository specific information. MarkDown format is supported, and the length limit is 2000.
func (o RegistryEnterpriseRepoOutput) Detail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringPtrOutput { return v.Detail }).(pulumi.StringPtrOutput)
}

// ID of Container Registry Enterprise Edition instance.
func (o RegistryEnterpriseRepoOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.
func (o RegistryEnterpriseRepoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.
func (o RegistryEnterpriseRepoOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// The uuid of Container Registry Enterprise Edition repository.
func (o RegistryEnterpriseRepoOutput) RepoId() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringOutput { return v.RepoId }).(pulumi.StringOutput)
}

// `PUBLIC` or `PRIVATE`, repo's visibility.
func (o RegistryEnterpriseRepoOutput) RepoType() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringOutput { return v.RepoType }).(pulumi.StringOutput)
}

// The repository general information. It can contain 1 to 100 characters.
func (o RegistryEnterpriseRepoOutput) Summary() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryEnterpriseRepo) pulumi.StringOutput { return v.Summary }).(pulumi.StringOutput)
}

type RegistryEnterpriseRepoArrayOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseRepoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryEnterpriseRepo)(nil)).Elem()
}

func (o RegistryEnterpriseRepoArrayOutput) ToRegistryEnterpriseRepoArrayOutput() RegistryEnterpriseRepoArrayOutput {
	return o
}

func (o RegistryEnterpriseRepoArrayOutput) ToRegistryEnterpriseRepoArrayOutputWithContext(ctx context.Context) RegistryEnterpriseRepoArrayOutput {
	return o
}

func (o RegistryEnterpriseRepoArrayOutput) Index(i pulumi.IntInput) RegistryEnterpriseRepoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegistryEnterpriseRepo {
		return vs[0].([]*RegistryEnterpriseRepo)[vs[1].(int)]
	}).(RegistryEnterpriseRepoOutput)
}

type RegistryEnterpriseRepoMapOutput struct{ *pulumi.OutputState }

func (RegistryEnterpriseRepoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryEnterpriseRepo)(nil)).Elem()
}

func (o RegistryEnterpriseRepoMapOutput) ToRegistryEnterpriseRepoMapOutput() RegistryEnterpriseRepoMapOutput {
	return o
}

func (o RegistryEnterpriseRepoMapOutput) ToRegistryEnterpriseRepoMapOutputWithContext(ctx context.Context) RegistryEnterpriseRepoMapOutput {
	return o
}

func (o RegistryEnterpriseRepoMapOutput) MapIndex(k pulumi.StringInput) RegistryEnterpriseRepoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegistryEnterpriseRepo {
		return vs[0].(map[string]*RegistryEnterpriseRepo)[vs[1].(string)]
	}).(RegistryEnterpriseRepoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseRepoInput)(nil)).Elem(), &RegistryEnterpriseRepo{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseRepoArrayInput)(nil)).Elem(), RegistryEnterpriseRepoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryEnterpriseRepoMapInput)(nil)).Elem(), RegistryEnterpriseRepoMap{})
	pulumi.RegisterOutputType(RegistryEnterpriseRepoOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseRepoArrayOutput{})
	pulumi.RegisterOutputType(RegistryEnterpriseRepoMapOutput{})
}
