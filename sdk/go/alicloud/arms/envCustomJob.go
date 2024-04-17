// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ARMS Env Custom Job resource. Custom jobs in the arms environment.
//
// For information about ARMS Env Custom Job and how to use it, see [What is Env Custom Job](https://www.alibabacloud.com/help/en/arms/developer-reference/api-arms-2019-08-08-createenvcustomjob).
//
// > **NOTE:** Available since v1.212.0.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/arms"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewInteger(ctx, "default", &random.IntegerArgs{
//				Max: 99999,
//				Min: 10000,
//			})
//			if err != nil {
//				return err
//			}
//			cfg := config.New(ctx, "")
//			name := "terraform-example"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			vpc, err := vpc.NewNetwork(ctx, "vpc", &vpc.NetworkArgs{
//				Description: pulumi.String(name),
//				CidrBlock:   pulumi.String("172.16.0.0/12"),
//				VpcName:     pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = arms.NewEnvironment(ctx, "env-ecs", &arms.EnvironmentArgs{
//				EnvironmentType:    pulumi.String("ECS"),
//				EnvironmentName:    pulumi.String(fmt.Sprintf("terraform-example-%v", _default.Result)),
//				BindResourceId:     vpc.ID(),
//				EnvironmentSubType: pulumi.String("ECS"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = arms.NewEnvCustomJob(ctx, "default", &arms.EnvCustomJobArgs{
//				Status:           pulumi.String("run"),
//				EnvironmentId:    env_ecs.ID(),
//				EnvCustomJobName: pulumi.String(name),
//				ConfigYaml: pulumi.String(`scrape_configs:
//   - job_name: job-demo1
//     honor_timestamps: false
//     honor_labels: false
//     scrape_interval: 30s
//     scheme: http
//     metrics_path: /metric
//     static_configs:
//   - targets:
//   - 127.0.0.1:9090
//
// `),
//
//				AliyunLang: pulumi.String("en"),
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
// ARMS Env Custom Job can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:arms/envCustomJob:EnvCustomJob example <environment_id>:<env_custom_job_name>
// ```
type EnvCustomJob struct {
	pulumi.CustomResourceState

	// The locale. The default is Chinese zh | en.
	AliyunLang pulumi.StringPtrOutput `pulumi:"aliyunLang"`
	// Yaml configuration string.
	ConfigYaml pulumi.StringOutput `pulumi:"configYaml"`
	// Custom job name.
	EnvCustomJobName pulumi.StringOutput `pulumi:"envCustomJobName"`
	// Environment id.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// Status: run, stop.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewEnvCustomJob registers a new resource with the given unique name, arguments, and options.
func NewEnvCustomJob(ctx *pulumi.Context,
	name string, args *EnvCustomJobArgs, opts ...pulumi.ResourceOption) (*EnvCustomJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigYaml == nil {
		return nil, errors.New("invalid value for required argument 'ConfigYaml'")
	}
	if args.EnvCustomJobName == nil {
		return nil, errors.New("invalid value for required argument 'EnvCustomJobName'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvCustomJob
	err := ctx.RegisterResource("alicloud:arms/envCustomJob:EnvCustomJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvCustomJob gets an existing EnvCustomJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvCustomJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvCustomJobState, opts ...pulumi.ResourceOption) (*EnvCustomJob, error) {
	var resource EnvCustomJob
	err := ctx.ReadResource("alicloud:arms/envCustomJob:EnvCustomJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvCustomJob resources.
type envCustomJobState struct {
	// The locale. The default is Chinese zh | en.
	AliyunLang *string `pulumi:"aliyunLang"`
	// Yaml configuration string.
	ConfigYaml *string `pulumi:"configYaml"`
	// Custom job name.
	EnvCustomJobName *string `pulumi:"envCustomJobName"`
	// Environment id.
	EnvironmentId *string `pulumi:"environmentId"`
	// Status: run, stop.
	Status *string `pulumi:"status"`
}

type EnvCustomJobState struct {
	// The locale. The default is Chinese zh | en.
	AliyunLang pulumi.StringPtrInput
	// Yaml configuration string.
	ConfigYaml pulumi.StringPtrInput
	// Custom job name.
	EnvCustomJobName pulumi.StringPtrInput
	// Environment id.
	EnvironmentId pulumi.StringPtrInput
	// Status: run, stop.
	Status pulumi.StringPtrInput
}

func (EnvCustomJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*envCustomJobState)(nil)).Elem()
}

type envCustomJobArgs struct {
	// The locale. The default is Chinese zh | en.
	AliyunLang *string `pulumi:"aliyunLang"`
	// Yaml configuration string.
	ConfigYaml string `pulumi:"configYaml"`
	// Custom job name.
	EnvCustomJobName string `pulumi:"envCustomJobName"`
	// Environment id.
	EnvironmentId string `pulumi:"environmentId"`
	// Status: run, stop.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a EnvCustomJob resource.
type EnvCustomJobArgs struct {
	// The locale. The default is Chinese zh | en.
	AliyunLang pulumi.StringPtrInput
	// Yaml configuration string.
	ConfigYaml pulumi.StringInput
	// Custom job name.
	EnvCustomJobName pulumi.StringInput
	// Environment id.
	EnvironmentId pulumi.StringInput
	// Status: run, stop.
	Status pulumi.StringPtrInput
}

func (EnvCustomJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*envCustomJobArgs)(nil)).Elem()
}

type EnvCustomJobInput interface {
	pulumi.Input

	ToEnvCustomJobOutput() EnvCustomJobOutput
	ToEnvCustomJobOutputWithContext(ctx context.Context) EnvCustomJobOutput
}

func (*EnvCustomJob) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvCustomJob)(nil)).Elem()
}

func (i *EnvCustomJob) ToEnvCustomJobOutput() EnvCustomJobOutput {
	return i.ToEnvCustomJobOutputWithContext(context.Background())
}

func (i *EnvCustomJob) ToEnvCustomJobOutputWithContext(ctx context.Context) EnvCustomJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvCustomJobOutput)
}

// EnvCustomJobArrayInput is an input type that accepts EnvCustomJobArray and EnvCustomJobArrayOutput values.
// You can construct a concrete instance of `EnvCustomJobArrayInput` via:
//
//	EnvCustomJobArray{ EnvCustomJobArgs{...} }
type EnvCustomJobArrayInput interface {
	pulumi.Input

	ToEnvCustomJobArrayOutput() EnvCustomJobArrayOutput
	ToEnvCustomJobArrayOutputWithContext(context.Context) EnvCustomJobArrayOutput
}

type EnvCustomJobArray []EnvCustomJobInput

func (EnvCustomJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvCustomJob)(nil)).Elem()
}

func (i EnvCustomJobArray) ToEnvCustomJobArrayOutput() EnvCustomJobArrayOutput {
	return i.ToEnvCustomJobArrayOutputWithContext(context.Background())
}

func (i EnvCustomJobArray) ToEnvCustomJobArrayOutputWithContext(ctx context.Context) EnvCustomJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvCustomJobArrayOutput)
}

// EnvCustomJobMapInput is an input type that accepts EnvCustomJobMap and EnvCustomJobMapOutput values.
// You can construct a concrete instance of `EnvCustomJobMapInput` via:
//
//	EnvCustomJobMap{ "key": EnvCustomJobArgs{...} }
type EnvCustomJobMapInput interface {
	pulumi.Input

	ToEnvCustomJobMapOutput() EnvCustomJobMapOutput
	ToEnvCustomJobMapOutputWithContext(context.Context) EnvCustomJobMapOutput
}

type EnvCustomJobMap map[string]EnvCustomJobInput

func (EnvCustomJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvCustomJob)(nil)).Elem()
}

func (i EnvCustomJobMap) ToEnvCustomJobMapOutput() EnvCustomJobMapOutput {
	return i.ToEnvCustomJobMapOutputWithContext(context.Background())
}

func (i EnvCustomJobMap) ToEnvCustomJobMapOutputWithContext(ctx context.Context) EnvCustomJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvCustomJobMapOutput)
}

type EnvCustomJobOutput struct{ *pulumi.OutputState }

func (EnvCustomJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvCustomJob)(nil)).Elem()
}

func (o EnvCustomJobOutput) ToEnvCustomJobOutput() EnvCustomJobOutput {
	return o
}

func (o EnvCustomJobOutput) ToEnvCustomJobOutputWithContext(ctx context.Context) EnvCustomJobOutput {
	return o
}

// The locale. The default is Chinese zh | en.
func (o EnvCustomJobOutput) AliyunLang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvCustomJob) pulumi.StringPtrOutput { return v.AliyunLang }).(pulumi.StringPtrOutput)
}

// Yaml configuration string.
func (o EnvCustomJobOutput) ConfigYaml() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvCustomJob) pulumi.StringOutput { return v.ConfigYaml }).(pulumi.StringOutput)
}

// Custom job name.
func (o EnvCustomJobOutput) EnvCustomJobName() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvCustomJob) pulumi.StringOutput { return v.EnvCustomJobName }).(pulumi.StringOutput)
}

// Environment id.
func (o EnvCustomJobOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvCustomJob) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// Status: run, stop.
func (o EnvCustomJobOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvCustomJob) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type EnvCustomJobArrayOutput struct{ *pulumi.OutputState }

func (EnvCustomJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvCustomJob)(nil)).Elem()
}

func (o EnvCustomJobArrayOutput) ToEnvCustomJobArrayOutput() EnvCustomJobArrayOutput {
	return o
}

func (o EnvCustomJobArrayOutput) ToEnvCustomJobArrayOutputWithContext(ctx context.Context) EnvCustomJobArrayOutput {
	return o
}

func (o EnvCustomJobArrayOutput) Index(i pulumi.IntInput) EnvCustomJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnvCustomJob {
		return vs[0].([]*EnvCustomJob)[vs[1].(int)]
	}).(EnvCustomJobOutput)
}

type EnvCustomJobMapOutput struct{ *pulumi.OutputState }

func (EnvCustomJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvCustomJob)(nil)).Elem()
}

func (o EnvCustomJobMapOutput) ToEnvCustomJobMapOutput() EnvCustomJobMapOutput {
	return o
}

func (o EnvCustomJobMapOutput) ToEnvCustomJobMapOutputWithContext(ctx context.Context) EnvCustomJobMapOutput {
	return o
}

func (o EnvCustomJobMapOutput) MapIndex(k pulumi.StringInput) EnvCustomJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnvCustomJob {
		return vs[0].(map[string]*EnvCustomJob)[vs[1].(string)]
	}).(EnvCustomJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvCustomJobInput)(nil)).Elem(), &EnvCustomJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvCustomJobArrayInput)(nil)).Elem(), EnvCustomJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvCustomJobMapInput)(nil)).Elem(), EnvCustomJobMap{})
	pulumi.RegisterOutputType(EnvCustomJobOutput{})
	pulumi.RegisterOutputType(EnvCustomJobArrayOutput{})
	pulumi.RegisterOutputType(EnvCustomJobMapOutput{})
}
