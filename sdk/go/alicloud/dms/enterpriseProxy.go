// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS Enterprise Proxy resource.
//
// For information about DMS Enterprise Proxy and how to use it, see [What is Proxy](https://next.api.alibabacloud.com/document/dms-enterprise/2018-11-01/CreateProxy).
//
// > **NOTE:** Available since v1.188.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/dms"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
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
//			current, err := alicloud.GetAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_default, err := alicloud.GetRegions(ctx, &alicloud.GetRegionsArgs{
//				Current: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetUserTenants, err := dms.GetUserTenants(ctx, &dms.GetUserTenantsArgs{
//				Status: pulumi.StringRef("ACTIVE"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetZones, err := rds.GetZones(ctx, &rds.GetZonesArgs{
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				InstanceChargeType:    pulumi.StringRef("PostPaid"),
//				Category:              pulumi.StringRef("HighAvailability"),
//				DbInstanceStorageType: pulumi.StringRef("cloud_essd"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultGetInstanceClasses, err := rds.GetInstanceClasses(ctx, &rds.GetInstanceClassesArgs{
//				ZoneId:                pulumi.StringRef(defaultGetZones.Zones[1].Id),
//				Engine:                pulumi.StringRef("MySQL"),
//				EngineVersion:         pulumi.StringRef("8.0"),
//				Category:              pulumi.StringRef("HighAvailability"),
//				DbInstanceStorageType: pulumi.StringRef("cloud_essd"),
//				InstanceChargeType:    pulumi.StringRef("PostPaid"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String(name),
//				CidrBlock: pulumi.String("10.4.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VswitchName: pulumi.String(name),
//				CidrBlock:   pulumi.String("10.4.0.0/24"),
//				VpcId:       defaultNetwork.ID(),
//				ZoneId:      pulumi.String(defaultGetZones.Zones[1].Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewSecurityGroup(ctx, "default", &ecs.SecurityGroupArgs{
//				Name:  pulumi.String(name),
//				VpcId: defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultInstance, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
//				Engine:                pulumi.String("MySQL"),
//				EngineVersion:         pulumi.String("8.0"),
//				DbInstanceStorageType: pulumi.String("cloud_essd"),
//				InstanceType:          pulumi.String(defaultGetInstanceClasses.InstanceClasses[0].InstanceClass),
//				InstanceStorage:       pulumi.String(defaultGetInstanceClasses.InstanceClasses[0].StorageRange.Min),
//				VswitchId:             defaultSwitch.ID(),
//				InstanceName:          pulumi.String(name),
//				SecurityIps: pulumi.StringArray{
//					pulumi.String("100.104.5.0/24"),
//					pulumi.String("192.168.0.6"),
//				},
//				Tags: pulumi.StringMap{
//					"Created": pulumi.String("TF"),
//					"For":     pulumi.String("example"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultAccount, err := rds.NewAccount(ctx, "default", &rds.AccountArgs{
//				DbInstanceId:    defaultInstance.ID(),
//				AccountName:     pulumi.String("tfexamplename"),
//				AccountPassword: pulumi.String("Example12345"),
//				AccountType:     pulumi.String("Normal"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEnterpriseInstance, err := dms.NewEnterpriseInstance(ctx, "default", &dms.EnterpriseInstanceArgs{
//				Tid:              pulumi.String(defaultGetUserTenants.Ids[0]),
//				InstanceType:     pulumi.String("mysql"),
//				InstanceSource:   pulumi.String("RDS"),
//				NetworkType:      pulumi.String("VPC"),
//				EnvType:          pulumi.String("dev"),
//				Host:             defaultInstance.ConnectionString,
//				Port:             pulumi.Int(3306),
//				DatabaseUser:     defaultAccount.AccountName,
//				DatabasePassword: defaultAccount.AccountPassword,
//				InstanceName:     pulumi.String(name),
//				DbaUid:           pulumi.String(current.Id),
//				SafeRule:         pulumi.String("自由操作"),
//				QueryTimeout:     pulumi.Int(60),
//				ExportTimeout:    pulumi.Int(600),
//				EcsRegion:        pulumi.String(_default.Regions[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dms.NewEnterpriseProxy(ctx, "default", &dms.EnterpriseProxyArgs{
//				InstanceId: defaultEnterpriseInstance.InstanceId,
//				Password:   pulumi.String("Example12345"),
//				Username:   pulumi.String("tfexamplename"),
//				Tid:        pulumi.String(defaultGetUserTenants.Ids[0]),
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
// DMS Enterprise Proxy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:dms/enterpriseProxy:EnterpriseProxy example <id>
// ```
type EnterpriseProxy struct {
	pulumi.CustomResourceState

	// The ID of the database instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The password of the database account.
	Password pulumi.StringOutput `pulumi:"password"`
	// The ID of the tenant.
	Tid pulumi.StringPtrOutput `pulumi:"tid"`
	// The username of the database account.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewEnterpriseProxy registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseProxy(ctx *pulumi.Context,
	name string, args *EnterpriseProxyArgs, opts ...pulumi.ResourceOption) (*EnterpriseProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	if args.Username != nil {
		args.Username = pulumi.ToSecret(args.Username).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"username",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseProxy
	err := ctx.RegisterResource("alicloud:dms/enterpriseProxy:EnterpriseProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseProxy gets an existing EnterpriseProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseProxyState, opts ...pulumi.ResourceOption) (*EnterpriseProxy, error) {
	var resource EnterpriseProxy
	err := ctx.ReadResource("alicloud:dms/enterpriseProxy:EnterpriseProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseProxy resources.
type enterpriseProxyState struct {
	// The ID of the database instance.
	InstanceId *string `pulumi:"instanceId"`
	// The password of the database account.
	Password *string `pulumi:"password"`
	// The ID of the tenant.
	Tid *string `pulumi:"tid"`
	// The username of the database account.
	Username *string `pulumi:"username"`
}

type EnterpriseProxyState struct {
	// The ID of the database instance.
	InstanceId pulumi.StringPtrInput
	// The password of the database account.
	Password pulumi.StringPtrInput
	// The ID of the tenant.
	Tid pulumi.StringPtrInput
	// The username of the database account.
	Username pulumi.StringPtrInput
}

func (EnterpriseProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseProxyState)(nil)).Elem()
}

type enterpriseProxyArgs struct {
	// The ID of the database instance.
	InstanceId string `pulumi:"instanceId"`
	// The password of the database account.
	Password string `pulumi:"password"`
	// The ID of the tenant.
	Tid *string `pulumi:"tid"`
	// The username of the database account.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a EnterpriseProxy resource.
type EnterpriseProxyArgs struct {
	// The ID of the database instance.
	InstanceId pulumi.StringInput
	// The password of the database account.
	Password pulumi.StringInput
	// The ID of the tenant.
	Tid pulumi.StringPtrInput
	// The username of the database account.
	Username pulumi.StringInput
}

func (EnterpriseProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseProxyArgs)(nil)).Elem()
}

type EnterpriseProxyInput interface {
	pulumi.Input

	ToEnterpriseProxyOutput() EnterpriseProxyOutput
	ToEnterpriseProxyOutputWithContext(ctx context.Context) EnterpriseProxyOutput
}

func (*EnterpriseProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseProxy)(nil)).Elem()
}

func (i *EnterpriseProxy) ToEnterpriseProxyOutput() EnterpriseProxyOutput {
	return i.ToEnterpriseProxyOutputWithContext(context.Background())
}

func (i *EnterpriseProxy) ToEnterpriseProxyOutputWithContext(ctx context.Context) EnterpriseProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseProxyOutput)
}

// EnterpriseProxyArrayInput is an input type that accepts EnterpriseProxyArray and EnterpriseProxyArrayOutput values.
// You can construct a concrete instance of `EnterpriseProxyArrayInput` via:
//
//	EnterpriseProxyArray{ EnterpriseProxyArgs{...} }
type EnterpriseProxyArrayInput interface {
	pulumi.Input

	ToEnterpriseProxyArrayOutput() EnterpriseProxyArrayOutput
	ToEnterpriseProxyArrayOutputWithContext(context.Context) EnterpriseProxyArrayOutput
}

type EnterpriseProxyArray []EnterpriseProxyInput

func (EnterpriseProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseProxy)(nil)).Elem()
}

func (i EnterpriseProxyArray) ToEnterpriseProxyArrayOutput() EnterpriseProxyArrayOutput {
	return i.ToEnterpriseProxyArrayOutputWithContext(context.Background())
}

func (i EnterpriseProxyArray) ToEnterpriseProxyArrayOutputWithContext(ctx context.Context) EnterpriseProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseProxyArrayOutput)
}

// EnterpriseProxyMapInput is an input type that accepts EnterpriseProxyMap and EnterpriseProxyMapOutput values.
// You can construct a concrete instance of `EnterpriseProxyMapInput` via:
//
//	EnterpriseProxyMap{ "key": EnterpriseProxyArgs{...} }
type EnterpriseProxyMapInput interface {
	pulumi.Input

	ToEnterpriseProxyMapOutput() EnterpriseProxyMapOutput
	ToEnterpriseProxyMapOutputWithContext(context.Context) EnterpriseProxyMapOutput
}

type EnterpriseProxyMap map[string]EnterpriseProxyInput

func (EnterpriseProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseProxy)(nil)).Elem()
}

func (i EnterpriseProxyMap) ToEnterpriseProxyMapOutput() EnterpriseProxyMapOutput {
	return i.ToEnterpriseProxyMapOutputWithContext(context.Background())
}

func (i EnterpriseProxyMap) ToEnterpriseProxyMapOutputWithContext(ctx context.Context) EnterpriseProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseProxyMapOutput)
}

type EnterpriseProxyOutput struct{ *pulumi.OutputState }

func (EnterpriseProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseProxy)(nil)).Elem()
}

func (o EnterpriseProxyOutput) ToEnterpriseProxyOutput() EnterpriseProxyOutput {
	return o
}

func (o EnterpriseProxyOutput) ToEnterpriseProxyOutputWithContext(ctx context.Context) EnterpriseProxyOutput {
	return o
}

// The ID of the database instance.
func (o EnterpriseProxyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseProxy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The password of the database account.
func (o EnterpriseProxyOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseProxy) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The ID of the tenant.
func (o EnterpriseProxyOutput) Tid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseProxy) pulumi.StringPtrOutput { return v.Tid }).(pulumi.StringPtrOutput)
}

// The username of the database account.
func (o EnterpriseProxyOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseProxy) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type EnterpriseProxyArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseProxy)(nil)).Elem()
}

func (o EnterpriseProxyArrayOutput) ToEnterpriseProxyArrayOutput() EnterpriseProxyArrayOutput {
	return o
}

func (o EnterpriseProxyArrayOutput) ToEnterpriseProxyArrayOutputWithContext(ctx context.Context) EnterpriseProxyArrayOutput {
	return o
}

func (o EnterpriseProxyArrayOutput) Index(i pulumi.IntInput) EnterpriseProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseProxy {
		return vs[0].([]*EnterpriseProxy)[vs[1].(int)]
	}).(EnterpriseProxyOutput)
}

type EnterpriseProxyMapOutput struct{ *pulumi.OutputState }

func (EnterpriseProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseProxy)(nil)).Elem()
}

func (o EnterpriseProxyMapOutput) ToEnterpriseProxyMapOutput() EnterpriseProxyMapOutput {
	return o
}

func (o EnterpriseProxyMapOutput) ToEnterpriseProxyMapOutputWithContext(ctx context.Context) EnterpriseProxyMapOutput {
	return o
}

func (o EnterpriseProxyMapOutput) MapIndex(k pulumi.StringInput) EnterpriseProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseProxy {
		return vs[0].(map[string]*EnterpriseProxy)[vs[1].(string)]
	}).(EnterpriseProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseProxyInput)(nil)).Elem(), &EnterpriseProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseProxyArrayInput)(nil)).Elem(), EnterpriseProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseProxyMapInput)(nil)).Elem(), EnterpriseProxyMap{})
	pulumi.RegisterOutputType(EnterpriseProxyOutput{})
	pulumi.RegisterOutputType(EnterpriseProxyArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseProxyMapOutput{})
}
