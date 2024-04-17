// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataworks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Data Works Folder resource.
//
// For information about Data Works Folder and how to use it, see [What is Folder](https://help.aliyun.com/document_detail/173940.html).
//
// > **NOTE:** Available in v1.131.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dataworks"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataworks.NewFolder(ctx, "example", &dataworks.FolderArgs{
//				ProjectId:  pulumi.String("320687"),
//				FolderPath: pulumi.String("Business Flow/tfTestAcc/folderDi/tftest1"),
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
// Data Works Folder can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dataworks/folder:Folder example <folder_id>:<$.ProjectId>
// ```
type Folder struct {
	pulumi.CustomResourceState

	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
	FolderPath pulumi.StringOutput `pulumi:"folderPath"`
	// The ID of the project.
	ProjectId         pulumi.StringPtrOutput `pulumi:"projectId"`
	ProjectIdentifier pulumi.StringPtrOutput `pulumi:"projectIdentifier"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FolderPath == nil {
		return nil, errors.New("invalid value for required argument 'FolderPath'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Folder
	err := ctx.RegisterResource("alicloud:dataworks/folder:Folder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	err := ctx.ReadResource("alicloud:dataworks/folder:Folder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	FolderId *string `pulumi:"folderId"`
	// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
	FolderPath *string `pulumi:"folderPath"`
	// The ID of the project.
	ProjectId         *string `pulumi:"projectId"`
	ProjectIdentifier *string `pulumi:"projectIdentifier"`
}

type FolderState struct {
	FolderId pulumi.StringPtrInput
	// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
	FolderPath pulumi.StringPtrInput
	// The ID of the project.
	ProjectId         pulumi.StringPtrInput
	ProjectIdentifier pulumi.StringPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
	FolderPath string `pulumi:"folderPath"`
	// The ID of the project.
	ProjectId         *string `pulumi:"projectId"`
	ProjectIdentifier *string `pulumi:"projectIdentifier"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
	FolderPath pulumi.StringInput
	// The ID of the project.
	ProjectId         pulumi.StringPtrInput
	ProjectIdentifier pulumi.StringPtrInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

// FolderArrayInput is an input type that accepts FolderArray and FolderArrayOutput values.
// You can construct a concrete instance of `FolderArrayInput` via:
//
//	FolderArray{ FolderArgs{...} }
type FolderArrayInput interface {
	pulumi.Input

	ToFolderArrayOutput() FolderArrayOutput
	ToFolderArrayOutputWithContext(context.Context) FolderArrayOutput
}

type FolderArray []FolderInput

func (FolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (i FolderArray) ToFolderArrayOutput() FolderArrayOutput {
	return i.ToFolderArrayOutputWithContext(context.Background())
}

func (i FolderArray) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderArrayOutput)
}

// FolderMapInput is an input type that accepts FolderMap and FolderMapOutput values.
// You can construct a concrete instance of `FolderMapInput` via:
//
//	FolderMap{ "key": FolderArgs{...} }
type FolderMapInput interface {
	pulumi.Input

	ToFolderMapOutput() FolderMapOutput
	ToFolderMapOutputWithContext(context.Context) FolderMapOutput
}

type FolderMap map[string]FolderInput

func (FolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (i FolderMap) ToFolderMapOutput() FolderMapOutput {
	return i.ToFolderMapOutputWithContext(context.Background())
}

func (i FolderMap) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderMapOutput)
}

type FolderOutput struct{ *pulumi.OutputState }

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

func (o FolderOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Folder Path. The folder path composed with for part: `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}`. The first segment of path must be `Business Flow`, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:`folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined`. Then the finial part of folder path can be specified in yourself.
func (o FolderOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

// The ID of the project.
func (o FolderOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o FolderOutput) ProjectIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringPtrOutput { return v.ProjectIdentifier }).(pulumi.StringPtrOutput)
}

type FolderArrayOutput struct{ *pulumi.OutputState }

func (FolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (o FolderArrayOutput) ToFolderArrayOutput() FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) Index(i pulumi.IntInput) FolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].([]*Folder)[vs[1].(int)]
	}).(FolderOutput)
}

type FolderMapOutput struct{ *pulumi.OutputState }

func (FolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (o FolderMapOutput) ToFolderMapOutput() FolderMapOutput {
	return o
}

func (o FolderMapOutput) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return o
}

func (o FolderMapOutput) MapIndex(k pulumi.StringInput) FolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].(map[string]*Folder)[vs[1].(string)]
	}).(FolderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderInput)(nil)).Elem(), &Folder{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderArrayInput)(nil)).Elem(), FolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderMapInput)(nil)).Elem(), FolderMap{})
	pulumi.RegisterOutputType(FolderOutput{})
	pulumi.RegisterOutputType(FolderArrayOutput{})
	pulumi.RegisterOutputType(FolderMapOutput{})
}
