// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AlertNotificationList struct {
	// Notice content of alarm.
	Content string `pulumi:"content"`
	// Email address list.
	EmailLists []string `pulumi:"emailLists"`
	// SMS sending mobile number.
	MobileLists []string `pulumi:"mobileLists"`
	// Request address.
	ServiceUri *string `pulumi:"serviceUri"`
	// Notification type. support Email, SMS, DingTalk.
	Type string `pulumi:"type"`
}

// AlertNotificationListInput is an input type that accepts AlertNotificationListArgs and AlertNotificationListOutput values.
// You can construct a concrete instance of `AlertNotificationListInput` via:
//
//          AlertNotificationListArgs{...}
type AlertNotificationListInput interface {
	pulumi.Input

	ToAlertNotificationListOutput() AlertNotificationListOutput
	ToAlertNotificationListOutputWithContext(context.Context) AlertNotificationListOutput
}

type AlertNotificationListArgs struct {
	// Notice content of alarm.
	Content pulumi.StringInput `pulumi:"content"`
	// Email address list.
	EmailLists pulumi.StringArrayInput `pulumi:"emailLists"`
	// SMS sending mobile number.
	MobileLists pulumi.StringArrayInput `pulumi:"mobileLists"`
	// Request address.
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
	// Notification type. support Email, SMS, DingTalk.
	Type pulumi.StringInput `pulumi:"type"`
}

func (AlertNotificationListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertNotificationList)(nil)).Elem()
}

func (i AlertNotificationListArgs) ToAlertNotificationListOutput() AlertNotificationListOutput {
	return i.ToAlertNotificationListOutputWithContext(context.Background())
}

func (i AlertNotificationListArgs) ToAlertNotificationListOutputWithContext(ctx context.Context) AlertNotificationListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertNotificationListOutput)
}

// AlertNotificationListArrayInput is an input type that accepts AlertNotificationListArray and AlertNotificationListArrayOutput values.
// You can construct a concrete instance of `AlertNotificationListArrayInput` via:
//
//          AlertNotificationListArray{ AlertNotificationListArgs{...} }
type AlertNotificationListArrayInput interface {
	pulumi.Input

	ToAlertNotificationListArrayOutput() AlertNotificationListArrayOutput
	ToAlertNotificationListArrayOutputWithContext(context.Context) AlertNotificationListArrayOutput
}

type AlertNotificationListArray []AlertNotificationListInput

func (AlertNotificationListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertNotificationList)(nil)).Elem()
}

func (i AlertNotificationListArray) ToAlertNotificationListArrayOutput() AlertNotificationListArrayOutput {
	return i.ToAlertNotificationListArrayOutputWithContext(context.Background())
}

func (i AlertNotificationListArray) ToAlertNotificationListArrayOutputWithContext(ctx context.Context) AlertNotificationListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertNotificationListArrayOutput)
}

type AlertNotificationListOutput struct{ *pulumi.OutputState }

func (AlertNotificationListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertNotificationList)(nil)).Elem()
}

func (o AlertNotificationListOutput) ToAlertNotificationListOutput() AlertNotificationListOutput {
	return o
}

func (o AlertNotificationListOutput) ToAlertNotificationListOutputWithContext(ctx context.Context) AlertNotificationListOutput {
	return o
}

// Notice content of alarm.
func (o AlertNotificationListOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v AlertNotificationList) string { return v.Content }).(pulumi.StringOutput)
}

// Email address list.
func (o AlertNotificationListOutput) EmailLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertNotificationList) []string { return v.EmailLists }).(pulumi.StringArrayOutput)
}

// SMS sending mobile number.
func (o AlertNotificationListOutput) MobileLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertNotificationList) []string { return v.MobileLists }).(pulumi.StringArrayOutput)
}

// Request address.
func (o AlertNotificationListOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertNotificationList) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

// Notification type. support Email, SMS, DingTalk.
func (o AlertNotificationListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AlertNotificationList) string { return v.Type }).(pulumi.StringOutput)
}

type AlertNotificationListArrayOutput struct{ *pulumi.OutputState }

func (AlertNotificationListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertNotificationList)(nil)).Elem()
}

func (o AlertNotificationListArrayOutput) ToAlertNotificationListArrayOutput() AlertNotificationListArrayOutput {
	return o
}

func (o AlertNotificationListArrayOutput) ToAlertNotificationListArrayOutputWithContext(ctx context.Context) AlertNotificationListArrayOutput {
	return o
}

func (o AlertNotificationListArrayOutput) Index(i pulumi.IntInput) AlertNotificationListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertNotificationList {
		return vs[0].([]AlertNotificationList)[vs[1].(int)]
	}).(AlertNotificationListOutput)
}

type AlertQueryList struct {
	// chart title
	ChartTitle string `pulumi:"chartTitle"`
	// end time. example: 20s.
	End string `pulumi:"end"`
	// Query logstore
	Logstore string `pulumi:"logstore"`
	// query corresponding to chart. example: * AND aliyun.
	Query string `pulumi:"query"`
	// begin time. example: -60s.
	Start string `pulumi:"start"`
	// default Custom. No need to configure this parameter.
	TimeSpanType *string `pulumi:"timeSpanType"`
}

// AlertQueryListInput is an input type that accepts AlertQueryListArgs and AlertQueryListOutput values.
// You can construct a concrete instance of `AlertQueryListInput` via:
//
//          AlertQueryListArgs{...}
type AlertQueryListInput interface {
	pulumi.Input

	ToAlertQueryListOutput() AlertQueryListOutput
	ToAlertQueryListOutputWithContext(context.Context) AlertQueryListOutput
}

type AlertQueryListArgs struct {
	// chart title
	ChartTitle pulumi.StringInput `pulumi:"chartTitle"`
	// end time. example: 20s.
	End pulumi.StringInput `pulumi:"end"`
	// Query logstore
	Logstore pulumi.StringInput `pulumi:"logstore"`
	// query corresponding to chart. example: * AND aliyun.
	Query pulumi.StringInput `pulumi:"query"`
	// begin time. example: -60s.
	Start pulumi.StringInput `pulumi:"start"`
	// default Custom. No need to configure this parameter.
	TimeSpanType pulumi.StringPtrInput `pulumi:"timeSpanType"`
}

func (AlertQueryListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertQueryList)(nil)).Elem()
}

func (i AlertQueryListArgs) ToAlertQueryListOutput() AlertQueryListOutput {
	return i.ToAlertQueryListOutputWithContext(context.Background())
}

func (i AlertQueryListArgs) ToAlertQueryListOutputWithContext(ctx context.Context) AlertQueryListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertQueryListOutput)
}

// AlertQueryListArrayInput is an input type that accepts AlertQueryListArray and AlertQueryListArrayOutput values.
// You can construct a concrete instance of `AlertQueryListArrayInput` via:
//
//          AlertQueryListArray{ AlertQueryListArgs{...} }
type AlertQueryListArrayInput interface {
	pulumi.Input

	ToAlertQueryListArrayOutput() AlertQueryListArrayOutput
	ToAlertQueryListArrayOutputWithContext(context.Context) AlertQueryListArrayOutput
}

type AlertQueryListArray []AlertQueryListInput

func (AlertQueryListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertQueryList)(nil)).Elem()
}

func (i AlertQueryListArray) ToAlertQueryListArrayOutput() AlertQueryListArrayOutput {
	return i.ToAlertQueryListArrayOutputWithContext(context.Background())
}

func (i AlertQueryListArray) ToAlertQueryListArrayOutputWithContext(ctx context.Context) AlertQueryListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertQueryListArrayOutput)
}

type AlertQueryListOutput struct{ *pulumi.OutputState }

func (AlertQueryListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertQueryList)(nil)).Elem()
}

func (o AlertQueryListOutput) ToAlertQueryListOutput() AlertQueryListOutput {
	return o
}

func (o AlertQueryListOutput) ToAlertQueryListOutputWithContext(ctx context.Context) AlertQueryListOutput {
	return o
}

// chart title
func (o AlertQueryListOutput) ChartTitle() pulumi.StringOutput {
	return o.ApplyT(func(v AlertQueryList) string { return v.ChartTitle }).(pulumi.StringOutput)
}

// end time. example: 20s.
func (o AlertQueryListOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v AlertQueryList) string { return v.End }).(pulumi.StringOutput)
}

// Query logstore
func (o AlertQueryListOutput) Logstore() pulumi.StringOutput {
	return o.ApplyT(func(v AlertQueryList) string { return v.Logstore }).(pulumi.StringOutput)
}

// query corresponding to chart. example: * AND aliyun.
func (o AlertQueryListOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v AlertQueryList) string { return v.Query }).(pulumi.StringOutput)
}

// begin time. example: -60s.
func (o AlertQueryListOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v AlertQueryList) string { return v.Start }).(pulumi.StringOutput)
}

// default Custom. No need to configure this parameter.
func (o AlertQueryListOutput) TimeSpanType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertQueryList) *string { return v.TimeSpanType }).(pulumi.StringPtrOutput)
}

type AlertQueryListArrayOutput struct{ *pulumi.OutputState }

func (AlertQueryListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertQueryList)(nil)).Elem()
}

func (o AlertQueryListArrayOutput) ToAlertQueryListArrayOutput() AlertQueryListArrayOutput {
	return o
}

func (o AlertQueryListArrayOutput) ToAlertQueryListArrayOutputWithContext(ctx context.Context) AlertQueryListArrayOutput {
	return o
}

func (o AlertQueryListArrayOutput) Index(i pulumi.IntInput) AlertQueryListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertQueryList {
		return vs[0].([]AlertQueryList)[vs[1].(int)]
	}).(AlertQueryListOutput)
}

type StoreIndexFieldSearch struct {
	// The alias of one field.
	Alias *string `pulumi:"alias"`
	// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
	CaseSensitive *bool `pulumi:"caseSensitive"`
	// Whether to enable field analytics. Default to true.
	EnableAnalytics *bool `pulumi:"enableAnalytics"`
	// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
	IncludeChinese *bool `pulumi:"includeChinese"`
	// Use nested index when type is json
	JsonKeys []StoreIndexFieldSearchJsonKey `pulumi:"jsonKeys"`
	// When using the jsonKeys field, this field is required.
	Name string `pulumi:"name"`
	// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
	Token *string `pulumi:"token"`
	// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
	Type *string `pulumi:"type"`
}

// StoreIndexFieldSearchInput is an input type that accepts StoreIndexFieldSearchArgs and StoreIndexFieldSearchOutput values.
// You can construct a concrete instance of `StoreIndexFieldSearchInput` via:
//
//          StoreIndexFieldSearchArgs{...}
type StoreIndexFieldSearchInput interface {
	pulumi.Input

	ToStoreIndexFieldSearchOutput() StoreIndexFieldSearchOutput
	ToStoreIndexFieldSearchOutputWithContext(context.Context) StoreIndexFieldSearchOutput
}

type StoreIndexFieldSearchArgs struct {
	// The alias of one field.
	Alias pulumi.StringPtrInput `pulumi:"alias"`
	// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
	CaseSensitive pulumi.BoolPtrInput `pulumi:"caseSensitive"`
	// Whether to enable field analytics. Default to true.
	EnableAnalytics pulumi.BoolPtrInput `pulumi:"enableAnalytics"`
	// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
	IncludeChinese pulumi.BoolPtrInput `pulumi:"includeChinese"`
	// Use nested index when type is json
	JsonKeys StoreIndexFieldSearchJsonKeyArrayInput `pulumi:"jsonKeys"`
	// When using the jsonKeys field, this field is required.
	Name pulumi.StringInput `pulumi:"name"`
	// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
	Token pulumi.StringPtrInput `pulumi:"token"`
	// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (StoreIndexFieldSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreIndexFieldSearch)(nil)).Elem()
}

func (i StoreIndexFieldSearchArgs) ToStoreIndexFieldSearchOutput() StoreIndexFieldSearchOutput {
	return i.ToStoreIndexFieldSearchOutputWithContext(context.Background())
}

func (i StoreIndexFieldSearchArgs) ToStoreIndexFieldSearchOutputWithContext(ctx context.Context) StoreIndexFieldSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFieldSearchOutput)
}

// StoreIndexFieldSearchArrayInput is an input type that accepts StoreIndexFieldSearchArray and StoreIndexFieldSearchArrayOutput values.
// You can construct a concrete instance of `StoreIndexFieldSearchArrayInput` via:
//
//          StoreIndexFieldSearchArray{ StoreIndexFieldSearchArgs{...} }
type StoreIndexFieldSearchArrayInput interface {
	pulumi.Input

	ToStoreIndexFieldSearchArrayOutput() StoreIndexFieldSearchArrayOutput
	ToStoreIndexFieldSearchArrayOutputWithContext(context.Context) StoreIndexFieldSearchArrayOutput
}

type StoreIndexFieldSearchArray []StoreIndexFieldSearchInput

func (StoreIndexFieldSearchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreIndexFieldSearch)(nil)).Elem()
}

func (i StoreIndexFieldSearchArray) ToStoreIndexFieldSearchArrayOutput() StoreIndexFieldSearchArrayOutput {
	return i.ToStoreIndexFieldSearchArrayOutputWithContext(context.Background())
}

func (i StoreIndexFieldSearchArray) ToStoreIndexFieldSearchArrayOutputWithContext(ctx context.Context) StoreIndexFieldSearchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFieldSearchArrayOutput)
}

type StoreIndexFieldSearchOutput struct{ *pulumi.OutputState }

func (StoreIndexFieldSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreIndexFieldSearch)(nil)).Elem()
}

func (o StoreIndexFieldSearchOutput) ToStoreIndexFieldSearchOutput() StoreIndexFieldSearchOutput {
	return o
}

func (o StoreIndexFieldSearchOutput) ToStoreIndexFieldSearchOutputWithContext(ctx context.Context) StoreIndexFieldSearchOutput {
	return o
}

// The alias of one field.
func (o StoreIndexFieldSearchOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
func (o StoreIndexFieldSearchOutput) CaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) *bool { return v.CaseSensitive }).(pulumi.BoolPtrOutput)
}

// Whether to enable field analytics. Default to true.
func (o StoreIndexFieldSearchOutput) EnableAnalytics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) *bool { return v.EnableAnalytics }).(pulumi.BoolPtrOutput)
}

// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
func (o StoreIndexFieldSearchOutput) IncludeChinese() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) *bool { return v.IncludeChinese }).(pulumi.BoolPtrOutput)
}

// Use nested index when type is json
func (o StoreIndexFieldSearchOutput) JsonKeys() StoreIndexFieldSearchJsonKeyArrayOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) []StoreIndexFieldSearchJsonKey { return v.JsonKeys }).(StoreIndexFieldSearchJsonKeyArrayOutput)
}

// When using the jsonKeys field, this field is required.
func (o StoreIndexFieldSearchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) string { return v.Name }).(pulumi.StringOutput)
}

// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
func (o StoreIndexFieldSearchOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) *string { return v.Token }).(pulumi.StringPtrOutput)
}

// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
func (o StoreIndexFieldSearchOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearch) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StoreIndexFieldSearchArrayOutput struct{ *pulumi.OutputState }

func (StoreIndexFieldSearchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreIndexFieldSearch)(nil)).Elem()
}

func (o StoreIndexFieldSearchArrayOutput) ToStoreIndexFieldSearchArrayOutput() StoreIndexFieldSearchArrayOutput {
	return o
}

func (o StoreIndexFieldSearchArrayOutput) ToStoreIndexFieldSearchArrayOutputWithContext(ctx context.Context) StoreIndexFieldSearchArrayOutput {
	return o
}

func (o StoreIndexFieldSearchArrayOutput) Index(i pulumi.IntInput) StoreIndexFieldSearchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StoreIndexFieldSearch {
		return vs[0].([]StoreIndexFieldSearch)[vs[1].(int)]
	}).(StoreIndexFieldSearchOutput)
}

type StoreIndexFieldSearchJsonKey struct {
	// The alias of one field.
	Alias *string `pulumi:"alias"`
	// Whether to enable statistics. default to true.
	DocValue *bool `pulumi:"docValue"`
	// When using the jsonKeys field, this field is required.
	Name string `pulumi:"name"`
	// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
	Type *string `pulumi:"type"`
}

// StoreIndexFieldSearchJsonKeyInput is an input type that accepts StoreIndexFieldSearchJsonKeyArgs and StoreIndexFieldSearchJsonKeyOutput values.
// You can construct a concrete instance of `StoreIndexFieldSearchJsonKeyInput` via:
//
//          StoreIndexFieldSearchJsonKeyArgs{...}
type StoreIndexFieldSearchJsonKeyInput interface {
	pulumi.Input

	ToStoreIndexFieldSearchJsonKeyOutput() StoreIndexFieldSearchJsonKeyOutput
	ToStoreIndexFieldSearchJsonKeyOutputWithContext(context.Context) StoreIndexFieldSearchJsonKeyOutput
}

type StoreIndexFieldSearchJsonKeyArgs struct {
	// The alias of one field.
	Alias pulumi.StringPtrInput `pulumi:"alias"`
	// Whether to enable statistics. default to true.
	DocValue pulumi.BoolPtrInput `pulumi:"docValue"`
	// When using the jsonKeys field, this field is required.
	Name pulumi.StringInput `pulumi:"name"`
	// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (StoreIndexFieldSearchJsonKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreIndexFieldSearchJsonKey)(nil)).Elem()
}

func (i StoreIndexFieldSearchJsonKeyArgs) ToStoreIndexFieldSearchJsonKeyOutput() StoreIndexFieldSearchJsonKeyOutput {
	return i.ToStoreIndexFieldSearchJsonKeyOutputWithContext(context.Background())
}

func (i StoreIndexFieldSearchJsonKeyArgs) ToStoreIndexFieldSearchJsonKeyOutputWithContext(ctx context.Context) StoreIndexFieldSearchJsonKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFieldSearchJsonKeyOutput)
}

// StoreIndexFieldSearchJsonKeyArrayInput is an input type that accepts StoreIndexFieldSearchJsonKeyArray and StoreIndexFieldSearchJsonKeyArrayOutput values.
// You can construct a concrete instance of `StoreIndexFieldSearchJsonKeyArrayInput` via:
//
//          StoreIndexFieldSearchJsonKeyArray{ StoreIndexFieldSearchJsonKeyArgs{...} }
type StoreIndexFieldSearchJsonKeyArrayInput interface {
	pulumi.Input

	ToStoreIndexFieldSearchJsonKeyArrayOutput() StoreIndexFieldSearchJsonKeyArrayOutput
	ToStoreIndexFieldSearchJsonKeyArrayOutputWithContext(context.Context) StoreIndexFieldSearchJsonKeyArrayOutput
}

type StoreIndexFieldSearchJsonKeyArray []StoreIndexFieldSearchJsonKeyInput

func (StoreIndexFieldSearchJsonKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreIndexFieldSearchJsonKey)(nil)).Elem()
}

func (i StoreIndexFieldSearchJsonKeyArray) ToStoreIndexFieldSearchJsonKeyArrayOutput() StoreIndexFieldSearchJsonKeyArrayOutput {
	return i.ToStoreIndexFieldSearchJsonKeyArrayOutputWithContext(context.Background())
}

func (i StoreIndexFieldSearchJsonKeyArray) ToStoreIndexFieldSearchJsonKeyArrayOutputWithContext(ctx context.Context) StoreIndexFieldSearchJsonKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFieldSearchJsonKeyArrayOutput)
}

type StoreIndexFieldSearchJsonKeyOutput struct{ *pulumi.OutputState }

func (StoreIndexFieldSearchJsonKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreIndexFieldSearchJsonKey)(nil)).Elem()
}

func (o StoreIndexFieldSearchJsonKeyOutput) ToStoreIndexFieldSearchJsonKeyOutput() StoreIndexFieldSearchJsonKeyOutput {
	return o
}

func (o StoreIndexFieldSearchJsonKeyOutput) ToStoreIndexFieldSearchJsonKeyOutputWithContext(ctx context.Context) StoreIndexFieldSearchJsonKeyOutput {
	return o
}

// The alias of one field.
func (o StoreIndexFieldSearchJsonKeyOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearchJsonKey) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

// Whether to enable statistics. default to true.
func (o StoreIndexFieldSearchJsonKeyOutput) DocValue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearchJsonKey) *bool { return v.DocValue }).(pulumi.BoolPtrOutput)
}

// When using the jsonKeys field, this field is required.
func (o StoreIndexFieldSearchJsonKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StoreIndexFieldSearchJsonKey) string { return v.Name }).(pulumi.StringOutput)
}

// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
func (o StoreIndexFieldSearchJsonKeyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreIndexFieldSearchJsonKey) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type StoreIndexFieldSearchJsonKeyArrayOutput struct{ *pulumi.OutputState }

func (StoreIndexFieldSearchJsonKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreIndexFieldSearchJsonKey)(nil)).Elem()
}

func (o StoreIndexFieldSearchJsonKeyArrayOutput) ToStoreIndexFieldSearchJsonKeyArrayOutput() StoreIndexFieldSearchJsonKeyArrayOutput {
	return o
}

func (o StoreIndexFieldSearchJsonKeyArrayOutput) ToStoreIndexFieldSearchJsonKeyArrayOutputWithContext(ctx context.Context) StoreIndexFieldSearchJsonKeyArrayOutput {
	return o
}

func (o StoreIndexFieldSearchJsonKeyArrayOutput) Index(i pulumi.IntInput) StoreIndexFieldSearchJsonKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StoreIndexFieldSearchJsonKey {
		return vs[0].([]StoreIndexFieldSearchJsonKey)[vs[1].(int)]
	}).(StoreIndexFieldSearchJsonKeyOutput)
}

type StoreIndexFullText struct {
	// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
	CaseSensitive *bool `pulumi:"caseSensitive"`
	// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
	IncludeChinese *bool `pulumi:"includeChinese"`
	// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
	Token *string `pulumi:"token"`
}

// StoreIndexFullTextInput is an input type that accepts StoreIndexFullTextArgs and StoreIndexFullTextOutput values.
// You can construct a concrete instance of `StoreIndexFullTextInput` via:
//
//          StoreIndexFullTextArgs{...}
type StoreIndexFullTextInput interface {
	pulumi.Input

	ToStoreIndexFullTextOutput() StoreIndexFullTextOutput
	ToStoreIndexFullTextOutputWithContext(context.Context) StoreIndexFullTextOutput
}

type StoreIndexFullTextArgs struct {
	// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
	CaseSensitive pulumi.BoolPtrInput `pulumi:"caseSensitive"`
	// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
	IncludeChinese pulumi.BoolPtrInput `pulumi:"includeChinese"`
	// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
	Token pulumi.StringPtrInput `pulumi:"token"`
}

func (StoreIndexFullTextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreIndexFullText)(nil)).Elem()
}

func (i StoreIndexFullTextArgs) ToStoreIndexFullTextOutput() StoreIndexFullTextOutput {
	return i.ToStoreIndexFullTextOutputWithContext(context.Background())
}

func (i StoreIndexFullTextArgs) ToStoreIndexFullTextOutputWithContext(ctx context.Context) StoreIndexFullTextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFullTextOutput)
}

func (i StoreIndexFullTextArgs) ToStoreIndexFullTextPtrOutput() StoreIndexFullTextPtrOutput {
	return i.ToStoreIndexFullTextPtrOutputWithContext(context.Background())
}

func (i StoreIndexFullTextArgs) ToStoreIndexFullTextPtrOutputWithContext(ctx context.Context) StoreIndexFullTextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFullTextOutput).ToStoreIndexFullTextPtrOutputWithContext(ctx)
}

// StoreIndexFullTextPtrInput is an input type that accepts StoreIndexFullTextArgs, StoreIndexFullTextPtr and StoreIndexFullTextPtrOutput values.
// You can construct a concrete instance of `StoreIndexFullTextPtrInput` via:
//
//          StoreIndexFullTextArgs{...}
//
//  or:
//
//          nil
type StoreIndexFullTextPtrInput interface {
	pulumi.Input

	ToStoreIndexFullTextPtrOutput() StoreIndexFullTextPtrOutput
	ToStoreIndexFullTextPtrOutputWithContext(context.Context) StoreIndexFullTextPtrOutput
}

type storeIndexFullTextPtrType StoreIndexFullTextArgs

func StoreIndexFullTextPtr(v *StoreIndexFullTextArgs) StoreIndexFullTextPtrInput {
	return (*storeIndexFullTextPtrType)(v)
}

func (*storeIndexFullTextPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreIndexFullText)(nil)).Elem()
}

func (i *storeIndexFullTextPtrType) ToStoreIndexFullTextPtrOutput() StoreIndexFullTextPtrOutput {
	return i.ToStoreIndexFullTextPtrOutputWithContext(context.Background())
}

func (i *storeIndexFullTextPtrType) ToStoreIndexFullTextPtrOutputWithContext(ctx context.Context) StoreIndexFullTextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreIndexFullTextPtrOutput)
}

type StoreIndexFullTextOutput struct{ *pulumi.OutputState }

func (StoreIndexFullTextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreIndexFullText)(nil)).Elem()
}

func (o StoreIndexFullTextOutput) ToStoreIndexFullTextOutput() StoreIndexFullTextOutput {
	return o
}

func (o StoreIndexFullTextOutput) ToStoreIndexFullTextOutputWithContext(ctx context.Context) StoreIndexFullTextOutput {
	return o
}

func (o StoreIndexFullTextOutput) ToStoreIndexFullTextPtrOutput() StoreIndexFullTextPtrOutput {
	return o.ToStoreIndexFullTextPtrOutputWithContext(context.Background())
}

func (o StoreIndexFullTextOutput) ToStoreIndexFullTextPtrOutputWithContext(ctx context.Context) StoreIndexFullTextPtrOutput {
	return o.ApplyT(func(v StoreIndexFullText) *StoreIndexFullText {
		return &v
	}).(StoreIndexFullTextPtrOutput)
}

// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
func (o StoreIndexFullTextOutput) CaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StoreIndexFullText) *bool { return v.CaseSensitive }).(pulumi.BoolPtrOutput)
}

// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
func (o StoreIndexFullTextOutput) IncludeChinese() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StoreIndexFullText) *bool { return v.IncludeChinese }).(pulumi.BoolPtrOutput)
}

// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
func (o StoreIndexFullTextOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreIndexFullText) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type StoreIndexFullTextPtrOutput struct{ *pulumi.OutputState }

func (StoreIndexFullTextPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreIndexFullText)(nil)).Elem()
}

func (o StoreIndexFullTextPtrOutput) ToStoreIndexFullTextPtrOutput() StoreIndexFullTextPtrOutput {
	return o
}

func (o StoreIndexFullTextPtrOutput) ToStoreIndexFullTextPtrOutputWithContext(ctx context.Context) StoreIndexFullTextPtrOutput {
	return o
}

func (o StoreIndexFullTextPtrOutput) Elem() StoreIndexFullTextOutput {
	return o.ApplyT(func(v *StoreIndexFullText) StoreIndexFullText { return *v }).(StoreIndexFullTextOutput)
}

// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
func (o StoreIndexFullTextPtrOutput) CaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StoreIndexFullText) *bool {
		if v == nil {
			return nil
		}
		return v.CaseSensitive
	}).(pulumi.BoolPtrOutput)
}

// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
func (o StoreIndexFullTextPtrOutput) IncludeChinese() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StoreIndexFullText) *bool {
		if v == nil {
			return nil
		}
		return v.IncludeChinese
	}).(pulumi.BoolPtrOutput)
}

// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
func (o StoreIndexFullTextPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoreIndexFullText) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type StoreShard struct {
	BeginKey *string `pulumi:"beginKey"`
	EndKey   *string `pulumi:"endKey"`
	// The ID of the log project. It formats of `<project>:<name>`.
	Id     *int    `pulumi:"id"`
	Status *string `pulumi:"status"`
}

// StoreShardInput is an input type that accepts StoreShardArgs and StoreShardOutput values.
// You can construct a concrete instance of `StoreShardInput` via:
//
//          StoreShardArgs{...}
type StoreShardInput interface {
	pulumi.Input

	ToStoreShardOutput() StoreShardOutput
	ToStoreShardOutputWithContext(context.Context) StoreShardOutput
}

type StoreShardArgs struct {
	BeginKey pulumi.StringPtrInput `pulumi:"beginKey"`
	EndKey   pulumi.StringPtrInput `pulumi:"endKey"`
	// The ID of the log project. It formats of `<project>:<name>`.
	Id     pulumi.IntPtrInput    `pulumi:"id"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (StoreShardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreShard)(nil)).Elem()
}

func (i StoreShardArgs) ToStoreShardOutput() StoreShardOutput {
	return i.ToStoreShardOutputWithContext(context.Background())
}

func (i StoreShardArgs) ToStoreShardOutputWithContext(ctx context.Context) StoreShardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreShardOutput)
}

// StoreShardArrayInput is an input type that accepts StoreShardArray and StoreShardArrayOutput values.
// You can construct a concrete instance of `StoreShardArrayInput` via:
//
//          StoreShardArray{ StoreShardArgs{...} }
type StoreShardArrayInput interface {
	pulumi.Input

	ToStoreShardArrayOutput() StoreShardArrayOutput
	ToStoreShardArrayOutputWithContext(context.Context) StoreShardArrayOutput
}

type StoreShardArray []StoreShardInput

func (StoreShardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreShard)(nil)).Elem()
}

func (i StoreShardArray) ToStoreShardArrayOutput() StoreShardArrayOutput {
	return i.ToStoreShardArrayOutputWithContext(context.Background())
}

func (i StoreShardArray) ToStoreShardArrayOutputWithContext(ctx context.Context) StoreShardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreShardArrayOutput)
}

type StoreShardOutput struct{ *pulumi.OutputState }

func (StoreShardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreShard)(nil)).Elem()
}

func (o StoreShardOutput) ToStoreShardOutput() StoreShardOutput {
	return o
}

func (o StoreShardOutput) ToStoreShardOutputWithContext(ctx context.Context) StoreShardOutput {
	return o
}

func (o StoreShardOutput) BeginKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreShard) *string { return v.BeginKey }).(pulumi.StringPtrOutput)
}

func (o StoreShardOutput) EndKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreShard) *string { return v.EndKey }).(pulumi.StringPtrOutput)
}

// The ID of the log project. It formats of `<project>:<name>`.
func (o StoreShardOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StoreShard) *int { return v.Id }).(pulumi.IntPtrOutput)
}

func (o StoreShardOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StoreShard) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type StoreShardArrayOutput struct{ *pulumi.OutputState }

func (StoreShardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreShard)(nil)).Elem()
}

func (o StoreShardArrayOutput) ToStoreShardArrayOutput() StoreShardArrayOutput {
	return o
}

func (o StoreShardArrayOutput) ToStoreShardArrayOutputWithContext(ctx context.Context) StoreShardArrayOutput {
	return o
}

func (o StoreShardArrayOutput) Index(i pulumi.IntInput) StoreShardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StoreShard {
		return vs[0].([]StoreShard)[vs[1].(int)]
	}).(StoreShardOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertNotificationListOutput{})
	pulumi.RegisterOutputType(AlertNotificationListArrayOutput{})
	pulumi.RegisterOutputType(AlertQueryListOutput{})
	pulumi.RegisterOutputType(AlertQueryListArrayOutput{})
	pulumi.RegisterOutputType(StoreIndexFieldSearchOutput{})
	pulumi.RegisterOutputType(StoreIndexFieldSearchArrayOutput{})
	pulumi.RegisterOutputType(StoreIndexFieldSearchJsonKeyOutput{})
	pulumi.RegisterOutputType(StoreIndexFieldSearchJsonKeyArrayOutput{})
	pulumi.RegisterOutputType(StoreIndexFullTextOutput{})
	pulumi.RegisterOutputType(StoreIndexFullTextPtrOutput{})
	pulumi.RegisterOutputType(StoreShardOutput{})
	pulumi.RegisterOutputType(StoreShardArrayOutput{})
}
