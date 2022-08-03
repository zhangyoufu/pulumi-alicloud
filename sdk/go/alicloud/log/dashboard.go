// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The dashboard is a real-time data analysis platform provided by the log service. You can display frequently used query and analysis statements in the form of charts and save statistical charts to the dashboard.
// [Refer to details](https://www.alibabacloud.com/help/doc-detail/102530.htm).
//
// > **NOTE:** Available in 1.86.0, parameter "action" in charList is supported since 1.164.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/log"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := log.NewProject(ctx, "defaultProject", &log.ProjectArgs{
// 			Description: pulumi.String("tf unit test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = log.NewStore(ctx, "defaultStore", &log.StoreArgs{
// 			Project:         pulumi.String("tf-project"),
// 			RetentionPeriod: pulumi.Int(3000),
// 			ShardCount:      pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = log.NewDashboard(ctx, "example", &log.DashboardArgs{
// 			CharList: pulumi.String(fmt.Sprintf(`  [
//     {
//       "action": {},
//       "title":"new_title",
//       "type":"map",
//       "search":{
//         "logstore":"tf-logstore",
//         "topic":"new_topic",
//         "query":"* | SELECT COUNT(name) as ct_name, COUNT(product) as ct_product, name,product GROUP BY name,product",
//         "start":"-86400s",
//         "end":"now"
//       },
//       "display":{
//         "xAxis":[
//           "ct_name"
//         ],
//         "yAxis":[
//           "ct_product"
//         ],
//         "xPos":0,
//         "yPos":0,
//         "width":10,
//         "height":12,
//         "displayName":"xixihaha911"
//       }
//     }
//   ]
//
// `)),
// 			DashboardName: pulumi.String("tf-dashboard"),
// 			ProjectName:   pulumi.String("tf-project"),
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
// Log Dashboard can be imported using the id or name, e.g.
//
// ```sh
//  $ pulumi import alicloud:log/dashboard:Dashboard example tf-project:tf-logstore:tf-dashboard
// ```
type Dashboard struct {
	pulumi.CustomResourceState

	// Configuration of charts in the dashboard.
	CharList pulumi.StringOutput `pulumi:"charList"`
	// The name of the Log Dashboard.
	DashboardName pulumi.StringOutput `pulumi:"dashboardName"`
	// Dashboard alias.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CharList == nil {
		return nil, errors.New("invalid value for required argument 'CharList'")
	}
	if args.DashboardName == nil {
		return nil, errors.New("invalid value for required argument 'DashboardName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	var resource Dashboard
	err := ctx.RegisterResource("alicloud:log/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("alicloud:log/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	// Configuration of charts in the dashboard.
	CharList *string `pulumi:"charList"`
	// The name of the Log Dashboard.
	DashboardName *string `pulumi:"dashboardName"`
	// Dashboard alias.
	DisplayName *string `pulumi:"displayName"`
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName *string `pulumi:"projectName"`
}

type DashboardState struct {
	// Configuration of charts in the dashboard.
	CharList pulumi.StringPtrInput
	// The name of the Log Dashboard.
	DashboardName pulumi.StringPtrInput
	// Dashboard alias.
	DisplayName pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName pulumi.StringPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// Configuration of charts in the dashboard.
	CharList string `pulumi:"charList"`
	// The name of the Log Dashboard.
	DashboardName string `pulumi:"dashboardName"`
	// Dashboard alias.
	DisplayName *string `pulumi:"displayName"`
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName string `pulumi:"projectName"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// Configuration of charts in the dashboard.
	CharList pulumi.StringInput
	// The name of the Log Dashboard.
	DashboardName pulumi.StringInput
	// Dashboard alias.
	DisplayName pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	ProjectName pulumi.StringInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//          DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//          DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

// Configuration of charts in the dashboard.
func (o DashboardOutput) CharList() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.CharList }).(pulumi.StringOutput)
}

// The name of the Log Dashboard.
func (o DashboardOutput) DashboardName() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.DashboardName }).(pulumi.StringOutput)
}

// Dashboard alias.
func (o DashboardOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The name of the log project. It is the only in one Alicloud account.
func (o DashboardOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].([]*Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].(map[string]*Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardInput)(nil)).Elem(), &Dashboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardArrayInput)(nil)).Elem(), DashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardMapInput)(nil)).Elem(), DashboardMap{})
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}
