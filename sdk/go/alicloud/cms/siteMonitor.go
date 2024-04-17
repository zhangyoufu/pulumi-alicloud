// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides a site monitor resource and it can be used to monitor public endpoints and websites.
// Details at https://www.alibabacloud.com/help/doc-detail/67907.htm
//
// > **NOTE:** Available since v1.72.0.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cms.NewSiteMonitor(ctx, "basic", &cms.SiteMonitorArgs{
//				Address:  pulumi.String("http://www.alibabacloud.com"),
//				TaskName: pulumi.String("tf-example"),
//				TaskType: pulumi.String("HTTP"),
//				Interval: pulumi.Int(5),
//				IspCities: cms.SiteMonitorIspCityArray{
//					&cms.SiteMonitorIspCityArgs{
//						City: pulumi.String("546"),
//						Isp:  pulumi.String("465"),
//					},
//				},
//				OptionsJson: pulumi.String(`{
//	    "http_method": "get",
//	    "waitTime_after_completion": null,
//	    "ipv6_task": false,
//	    "diagnosis_ping": false,
//	    "diagnosis_mtr": false,
//	    "assertions": [
//	        {
//	            "operator": "lessThan",
//	            "type": "response_time",
//	            "target": 1000
//	        }
//	    ],
//	    "time_out": 30000
//	}
//
// `),
//
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
// Cloud Monitor Service Site Monitor can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:cms/siteMonitor:SiteMonitor example <id>
// ```
type SiteMonitor struct {
	pulumi.CustomResourceState

	// The URL or IP address monitored by the site monitoring task.
	Address pulumi.StringOutput `pulumi:"address"`
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds pulumi.StringArrayOutput `pulumi:"alertIds"`
	// The time when the site monitoring task was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `ispCities` below.
	IspCities SiteMonitorIspCityArrayOutput `pulumi:"ispCities"`
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol. See [extended options](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-createsitemonitor#api-detail-35).
	OptionsJson pulumi.StringPtrOutput `pulumi:"optionsJson"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName pulumi.StringOutput `pulumi:"taskName"`
	// The status of the site monitoring task.
	TaskState pulumi.StringOutput `pulumi:"taskState"`
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// The time when the site monitoring task was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSiteMonitor registers a new resource with the given unique name, arguments, and options.
func NewSiteMonitor(ctx *pulumi.Context,
	name string, args *SiteMonitorArgs, opts ...pulumi.ResourceOption) (*SiteMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SiteMonitor
	err := ctx.RegisterResource("alicloud:cms/siteMonitor:SiteMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteMonitor gets an existing SiteMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteMonitorState, opts ...pulumi.ResourceOption) (*SiteMonitor, error) {
	var resource SiteMonitor
	err := ctx.ReadResource("alicloud:cms/siteMonitor:SiteMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteMonitor resources.
type siteMonitorState struct {
	// The URL or IP address monitored by the site monitoring task.
	Address *string `pulumi:"address"`
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds []string `pulumi:"alertIds"`
	// The time when the site monitoring task was created.
	CreateTime *string `pulumi:"createTime"`
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
	Interval *int `pulumi:"interval"`
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `ispCities` below.
	IspCities []SiteMonitorIspCity `pulumi:"ispCities"`
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol. See [extended options](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-createsitemonitor#api-detail-35).
	OptionsJson *string `pulumi:"optionsJson"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName *string `pulumi:"taskName"`
	// The status of the site monitoring task.
	TaskState *string `pulumi:"taskState"`
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType *string `pulumi:"taskType"`
	// The time when the site monitoring task was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type SiteMonitorState struct {
	// The URL or IP address monitored by the site monitoring task.
	Address pulumi.StringPtrInput
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds pulumi.StringArrayInput
	// The time when the site monitoring task was created.
	CreateTime pulumi.StringPtrInput
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
	Interval pulumi.IntPtrInput
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `ispCities` below.
	IspCities SiteMonitorIspCityArrayInput
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol. See [extended options](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-createsitemonitor#api-detail-35).
	OptionsJson pulumi.StringPtrInput
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName pulumi.StringPtrInput
	// The status of the site monitoring task.
	TaskState pulumi.StringPtrInput
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType pulumi.StringPtrInput
	// The time when the site monitoring task was updated.
	UpdateTime pulumi.StringPtrInput
}

func (SiteMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMonitorState)(nil)).Elem()
}

type siteMonitorArgs struct {
	// The URL or IP address monitored by the site monitoring task.
	Address string `pulumi:"address"`
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds []string `pulumi:"alertIds"`
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
	Interval *int `pulumi:"interval"`
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `ispCities` below.
	IspCities []SiteMonitorIspCity `pulumi:"ispCities"`
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol. See [extended options](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-createsitemonitor#api-detail-35).
	OptionsJson *string `pulumi:"optionsJson"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName string `pulumi:"taskName"`
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType string `pulumi:"taskType"`
}

// The set of arguments for constructing a SiteMonitor resource.
type SiteMonitorArgs struct {
	// The URL or IP address monitored by the site monitoring task.
	Address pulumi.StringInput
	// The IDs of existing alert rules to be associated with the site monitoring task.
	AlertIds pulumi.StringArrayInput
	// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
	Interval pulumi.IntPtrInput
	// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `ispCities` below.
	IspCities SiteMonitorIspCityArrayInput
	// The extended options of the protocol of the site monitoring task. The options vary according to the protocol. See [extended options](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-createsitemonitor#api-detail-35).
	OptionsJson pulumi.StringPtrInput
	// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
	TaskName pulumi.StringInput
	// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	TaskType pulumi.StringInput
}

func (SiteMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteMonitorArgs)(nil)).Elem()
}

type SiteMonitorInput interface {
	pulumi.Input

	ToSiteMonitorOutput() SiteMonitorOutput
	ToSiteMonitorOutputWithContext(ctx context.Context) SiteMonitorOutput
}

func (*SiteMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMonitor)(nil)).Elem()
}

func (i *SiteMonitor) ToSiteMonitorOutput() SiteMonitorOutput {
	return i.ToSiteMonitorOutputWithContext(context.Background())
}

func (i *SiteMonitor) ToSiteMonitorOutputWithContext(ctx context.Context) SiteMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorOutput)
}

// SiteMonitorArrayInput is an input type that accepts SiteMonitorArray and SiteMonitorArrayOutput values.
// You can construct a concrete instance of `SiteMonitorArrayInput` via:
//
//	SiteMonitorArray{ SiteMonitorArgs{...} }
type SiteMonitorArrayInput interface {
	pulumi.Input

	ToSiteMonitorArrayOutput() SiteMonitorArrayOutput
	ToSiteMonitorArrayOutputWithContext(context.Context) SiteMonitorArrayOutput
}

type SiteMonitorArray []SiteMonitorInput

func (SiteMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SiteMonitor)(nil)).Elem()
}

func (i SiteMonitorArray) ToSiteMonitorArrayOutput() SiteMonitorArrayOutput {
	return i.ToSiteMonitorArrayOutputWithContext(context.Background())
}

func (i SiteMonitorArray) ToSiteMonitorArrayOutputWithContext(ctx context.Context) SiteMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorArrayOutput)
}

// SiteMonitorMapInput is an input type that accepts SiteMonitorMap and SiteMonitorMapOutput values.
// You can construct a concrete instance of `SiteMonitorMapInput` via:
//
//	SiteMonitorMap{ "key": SiteMonitorArgs{...} }
type SiteMonitorMapInput interface {
	pulumi.Input

	ToSiteMonitorMapOutput() SiteMonitorMapOutput
	ToSiteMonitorMapOutputWithContext(context.Context) SiteMonitorMapOutput
}

type SiteMonitorMap map[string]SiteMonitorInput

func (SiteMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SiteMonitor)(nil)).Elem()
}

func (i SiteMonitorMap) ToSiteMonitorMapOutput() SiteMonitorMapOutput {
	return i.ToSiteMonitorMapOutputWithContext(context.Background())
}

func (i SiteMonitorMap) ToSiteMonitorMapOutputWithContext(ctx context.Context) SiteMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMonitorMapOutput)
}

type SiteMonitorOutput struct{ *pulumi.OutputState }

func (SiteMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMonitor)(nil)).Elem()
}

func (o SiteMonitorOutput) ToSiteMonitorOutput() SiteMonitorOutput {
	return o
}

func (o SiteMonitorOutput) ToSiteMonitorOutputWithContext(ctx context.Context) SiteMonitorOutput {
	return o
}

// The URL or IP address monitored by the site monitoring task.
func (o SiteMonitorOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The IDs of existing alert rules to be associated with the site monitoring task.
func (o SiteMonitorOutput) AlertIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringArrayOutput { return v.AlertIds }).(pulumi.StringArrayOutput)
}

// The time when the site monitoring task was created.
func (o SiteMonitorOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The monitoring interval of the site monitoring task. Unit: minutes. Valid values: `1`, `5`, `15`, `30` and `60`. Default value: `1`. **NOTE:** From version 1.207.0, `interval` can be set to `30`, `60`.
func (o SiteMonitorOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// The detection points in a JSON array. For example, `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring. See `ispCities` below.
func (o SiteMonitorOutput) IspCities() SiteMonitorIspCityArrayOutput {
	return o.ApplyT(func(v *SiteMonitor) SiteMonitorIspCityArrayOutput { return v.IspCities }).(SiteMonitorIspCityArrayOutput)
}

// The extended options of the protocol of the site monitoring task. The options vary according to the protocol. See [extended options](https://www.alibabacloud.com/help/en/cms/developer-reference/api-cms-2019-01-01-createsitemonitor#api-detail-35).
func (o SiteMonitorOutput) OptionsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringPtrOutput { return v.OptionsJson }).(pulumi.StringPtrOutput)
}

// The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.
func (o SiteMonitorOutput) TaskName() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringOutput { return v.TaskName }).(pulumi.StringOutput)
}

// The status of the site monitoring task.
func (o SiteMonitorOutput) TaskState() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringOutput { return v.TaskState }).(pulumi.StringOutput)
}

// The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
func (o SiteMonitorOutput) TaskType() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringOutput { return v.TaskType }).(pulumi.StringOutput)
}

// The time when the site monitoring task was updated.
func (o SiteMonitorOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteMonitor) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type SiteMonitorArrayOutput struct{ *pulumi.OutputState }

func (SiteMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SiteMonitor)(nil)).Elem()
}

func (o SiteMonitorArrayOutput) ToSiteMonitorArrayOutput() SiteMonitorArrayOutput {
	return o
}

func (o SiteMonitorArrayOutput) ToSiteMonitorArrayOutputWithContext(ctx context.Context) SiteMonitorArrayOutput {
	return o
}

func (o SiteMonitorArrayOutput) Index(i pulumi.IntInput) SiteMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SiteMonitor {
		return vs[0].([]*SiteMonitor)[vs[1].(int)]
	}).(SiteMonitorOutput)
}

type SiteMonitorMapOutput struct{ *pulumi.OutputState }

func (SiteMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SiteMonitor)(nil)).Elem()
}

func (o SiteMonitorMapOutput) ToSiteMonitorMapOutput() SiteMonitorMapOutput {
	return o
}

func (o SiteMonitorMapOutput) ToSiteMonitorMapOutputWithContext(ctx context.Context) SiteMonitorMapOutput {
	return o
}

func (o SiteMonitorMapOutput) MapIndex(k pulumi.StringInput) SiteMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SiteMonitor {
		return vs[0].(map[string]*SiteMonitor)[vs[1].(string)]
	}).(SiteMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorInput)(nil)).Elem(), &SiteMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorArrayInput)(nil)).Elem(), SiteMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteMonitorMapInput)(nil)).Elem(), SiteMonitorMap{})
	pulumi.RegisterOutputType(SiteMonitorOutput{})
	pulumi.RegisterOutputType(SiteMonitorArrayOutput{})
	pulumi.RegisterOutputType(SiteMonitorMapOutput{})
}
