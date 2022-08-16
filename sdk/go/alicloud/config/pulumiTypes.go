// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssumeRole struct {
	Policy            *string `pulumi:"policy"`
	RoleArn           string  `pulumi:"roleArn"`
	SessionExpiration *int    `pulumi:"sessionExpiration"`
	SessionName       *string `pulumi:"sessionName"`
}

// AssumeRoleInput is an input type that accepts AssumeRoleArgs and AssumeRoleOutput values.
// You can construct a concrete instance of `AssumeRoleInput` via:
//
//	AssumeRoleArgs{...}
type AssumeRoleInput interface {
	pulumi.Input

	ToAssumeRoleOutput() AssumeRoleOutput
	ToAssumeRoleOutputWithContext(context.Context) AssumeRoleOutput
}

type AssumeRoleArgs struct {
	Policy            pulumi.StringPtrInput `pulumi:"policy"`
	RoleArn           pulumi.StringInput    `pulumi:"roleArn"`
	SessionExpiration pulumi.IntPtrInput    `pulumi:"sessionExpiration"`
	SessionName       pulumi.StringPtrInput `pulumi:"sessionName"`
}

func (AssumeRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssumeRole)(nil)).Elem()
}

func (i AssumeRoleArgs) ToAssumeRoleOutput() AssumeRoleOutput {
	return i.ToAssumeRoleOutputWithContext(context.Background())
}

func (i AssumeRoleArgs) ToAssumeRoleOutputWithContext(ctx context.Context) AssumeRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssumeRoleOutput)
}

type AssumeRoleOutput struct{ *pulumi.OutputState }

func (AssumeRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssumeRole)(nil)).Elem()
}

func (o AssumeRoleOutput) ToAssumeRoleOutput() AssumeRoleOutput {
	return o
}

func (o AssumeRoleOutput) ToAssumeRoleOutputWithContext(ctx context.Context) AssumeRoleOutput {
	return o
}

func (o AssumeRoleOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssumeRole) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func (o AssumeRoleOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v AssumeRole) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o AssumeRoleOutput) SessionExpiration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AssumeRole) *int { return v.SessionExpiration }).(pulumi.IntPtrOutput)
}

func (o AssumeRoleOutput) SessionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssumeRole) *string { return v.SessionName }).(pulumi.StringPtrOutput)
}

type Endpoints struct {
	Acr                 *string `pulumi:"acr"`
	Actiontrail         *string `pulumi:"actiontrail"`
	Adb                 *string `pulumi:"adb"`
	Alb                 *string `pulumi:"alb"`
	Alidfs              *string `pulumi:"alidfs"`
	Alidns              *string `pulumi:"alidns"`
	Alikafka            *string `pulumi:"alikafka"`
	Apigateway          *string `pulumi:"apigateway"`
	Arms                *string `pulumi:"arms"`
	Bastionhost         *string `pulumi:"bastionhost"`
	BrainIndustrial     *string `pulumi:"brainIndustrial"`
	Bssopenapi          *string `pulumi:"bssopenapi"`
	Cas                 *string `pulumi:"cas"`
	Cassandra           *string `pulumi:"cassandra"`
	Cbn                 *string `pulumi:"cbn"`
	Cddc                *string `pulumi:"cddc"`
	Cdn                 *string `pulumi:"cdn"`
	Cds                 *string `pulumi:"cds"`
	Clickhouse          *string `pulumi:"clickhouse"`
	Cloudauth           *string `pulumi:"cloudauth"`
	Cloudfw             *string `pulumi:"cloudfw"`
	Cloudphone          *string `pulumi:"cloudphone"`
	Cloudsso            *string `pulumi:"cloudsso"`
	Cms                 *string `pulumi:"cms"`
	Config              *string `pulumi:"config"`
	Cr                  *string `pulumi:"cr"`
	Cs                  *string `pulumi:"cs"`
	Datahub             *string `pulumi:"datahub"`
	Dataworkspublic     *string `pulumi:"dataworkspublic"`
	Dbfs                *string `pulumi:"dbfs"`
	Dcdn                *string `pulumi:"dcdn"`
	Ddosbasic           *string `pulumi:"ddosbasic"`
	Ddosbgp             *string `pulumi:"ddosbgp"`
	Ddoscoo             *string `pulumi:"ddoscoo"`
	Dds                 *string `pulumi:"dds"`
	Devopsrdc           *string `pulumi:"devopsrdc"`
	Dg                  *string `pulumi:"dg"`
	Dm                  *string `pulumi:"dm"`
	DmsEnterprise       *string `pulumi:"dmsEnterprise"`
	Dns                 *string `pulumi:"dns"`
	Drds                *string `pulumi:"drds"`
	Dts                 *string `pulumi:"dts"`
	Dysms               *string `pulumi:"dysms"`
	Eais                *string `pulumi:"eais"`
	Eci                 *string `pulumi:"eci"`
	Ecs                 *string `pulumi:"ecs"`
	Edas                *string `pulumi:"edas"`
	Edasschedulerx      *string `pulumi:"edasschedulerx"`
	Edsuser             *string `pulumi:"edsuser"`
	Ehpc                *string `pulumi:"ehpc"`
	Ehs                 *string `pulumi:"ehs"`
	Eipanycast          *string `pulumi:"eipanycast"`
	Elasticsearch       *string `pulumi:"elasticsearch"`
	Emr                 *string `pulumi:"emr"`
	Ens                 *string `pulumi:"ens"`
	Ess                 *string `pulumi:"ess"`
	Eventbridge         *string `pulumi:"eventbridge"`
	Fc                  *string `pulumi:"fc"`
	Fnf                 *string `pulumi:"fnf"`
	Ga                  *string `pulumi:"ga"`
	Gaplus              *string `pulumi:"gaplus"`
	Gds                 *string `pulumi:"gds"`
	Gpdb                *string `pulumi:"gpdb"`
	Gwsecd              *string `pulumi:"gwsecd"`
	Hbr                 *string `pulumi:"hbr"`
	HcsSgw              *string `pulumi:"hcsSgw"`
	Hitsdb              *string `pulumi:"hitsdb"`
	Imm                 *string `pulumi:"imm"`
	Imp                 *string `pulumi:"imp"`
	Ims                 *string `pulumi:"ims"`
	Iot                 *string `pulumi:"iot"`
	Kms                 *string `pulumi:"kms"`
	Kvstore             *string `pulumi:"kvstore"`
	Location            *string `pulumi:"location"`
	Log                 *string `pulumi:"log"`
	Market              *string `pulumi:"market"`
	Maxcompute          *string `pulumi:"maxcompute"`
	Mhub                *string `pulumi:"mhub"`
	Mns                 *string `pulumi:"mns"`
	Mscopensubscription *string `pulumi:"mscopensubscription"`
	Mse                 *string `pulumi:"mse"`
	Nas                 *string `pulumi:"nas"`
	Ons                 *string `pulumi:"ons"`
	Onsproxy            *string `pulumi:"onsproxy"`
	Oos                 *string `pulumi:"oos"`
	Opensearch          *string `pulumi:"opensearch"`
	Oss                 *string `pulumi:"oss"`
	Ots                 *string `pulumi:"ots"`
	Polardb             *string `pulumi:"polardb"`
	Privatelink         *string `pulumi:"privatelink"`
	Pvtz                *string `pulumi:"pvtz"`
	Quickbi             *string `pulumi:"quickbi"`
	Quotas              *string `pulumi:"quotas"`
	RKvstore            *string `pulumi:"rKvstore"`
	Ram                 *string `pulumi:"ram"`
	Rds                 *string `pulumi:"rds"`
	Redisa              *string `pulumi:"redisa"`
	Resourcemanager     *string `pulumi:"resourcemanager"`
	Resourcesharing     *string `pulumi:"resourcesharing"`
	Ros                 *string `pulumi:"ros"`
	Sas                 *string `pulumi:"sas"`
	Scdn                *string `pulumi:"scdn"`
	Sddp                *string `pulumi:"sddp"`
	Serverless          *string `pulumi:"serverless"`
	Servicemesh         *string `pulumi:"servicemesh"`
	Sgw                 *string `pulumi:"sgw"`
	Slb                 *string `pulumi:"slb"`
	Smartag             *string `pulumi:"smartag"`
	Sts                 *string `pulumi:"sts"`
	Swas                *string `pulumi:"swas"`
	Tag                 *string `pulumi:"tag"`
	Vod                 *string `pulumi:"vod"`
	Vpc                 *string `pulumi:"vpc"`
	Vs                  *string `pulumi:"vs"`
	Waf                 *string `pulumi:"waf"`
	WafOpenapi          *string `pulumi:"wafOpenapi"`
}

// EndpointsInput is an input type that accepts EndpointsArgs and EndpointsOutput values.
// You can construct a concrete instance of `EndpointsInput` via:
//
//	EndpointsArgs{...}
type EndpointsInput interface {
	pulumi.Input

	ToEndpointsOutput() EndpointsOutput
	ToEndpointsOutputWithContext(context.Context) EndpointsOutput
}

type EndpointsArgs struct {
	Acr                 pulumi.StringPtrInput `pulumi:"acr"`
	Actiontrail         pulumi.StringPtrInput `pulumi:"actiontrail"`
	Adb                 pulumi.StringPtrInput `pulumi:"adb"`
	Alb                 pulumi.StringPtrInput `pulumi:"alb"`
	Alidfs              pulumi.StringPtrInput `pulumi:"alidfs"`
	Alidns              pulumi.StringPtrInput `pulumi:"alidns"`
	Alikafka            pulumi.StringPtrInput `pulumi:"alikafka"`
	Apigateway          pulumi.StringPtrInput `pulumi:"apigateway"`
	Arms                pulumi.StringPtrInput `pulumi:"arms"`
	Bastionhost         pulumi.StringPtrInput `pulumi:"bastionhost"`
	BrainIndustrial     pulumi.StringPtrInput `pulumi:"brainIndustrial"`
	Bssopenapi          pulumi.StringPtrInput `pulumi:"bssopenapi"`
	Cas                 pulumi.StringPtrInput `pulumi:"cas"`
	Cassandra           pulumi.StringPtrInput `pulumi:"cassandra"`
	Cbn                 pulumi.StringPtrInput `pulumi:"cbn"`
	Cddc                pulumi.StringPtrInput `pulumi:"cddc"`
	Cdn                 pulumi.StringPtrInput `pulumi:"cdn"`
	Cds                 pulumi.StringPtrInput `pulumi:"cds"`
	Clickhouse          pulumi.StringPtrInput `pulumi:"clickhouse"`
	Cloudauth           pulumi.StringPtrInput `pulumi:"cloudauth"`
	Cloudfw             pulumi.StringPtrInput `pulumi:"cloudfw"`
	Cloudphone          pulumi.StringPtrInput `pulumi:"cloudphone"`
	Cloudsso            pulumi.StringPtrInput `pulumi:"cloudsso"`
	Cms                 pulumi.StringPtrInput `pulumi:"cms"`
	Config              pulumi.StringPtrInput `pulumi:"config"`
	Cr                  pulumi.StringPtrInput `pulumi:"cr"`
	Cs                  pulumi.StringPtrInput `pulumi:"cs"`
	Datahub             pulumi.StringPtrInput `pulumi:"datahub"`
	Dataworkspublic     pulumi.StringPtrInput `pulumi:"dataworkspublic"`
	Dbfs                pulumi.StringPtrInput `pulumi:"dbfs"`
	Dcdn                pulumi.StringPtrInput `pulumi:"dcdn"`
	Ddosbasic           pulumi.StringPtrInput `pulumi:"ddosbasic"`
	Ddosbgp             pulumi.StringPtrInput `pulumi:"ddosbgp"`
	Ddoscoo             pulumi.StringPtrInput `pulumi:"ddoscoo"`
	Dds                 pulumi.StringPtrInput `pulumi:"dds"`
	Devopsrdc           pulumi.StringPtrInput `pulumi:"devopsrdc"`
	Dg                  pulumi.StringPtrInput `pulumi:"dg"`
	Dm                  pulumi.StringPtrInput `pulumi:"dm"`
	DmsEnterprise       pulumi.StringPtrInput `pulumi:"dmsEnterprise"`
	Dns                 pulumi.StringPtrInput `pulumi:"dns"`
	Drds                pulumi.StringPtrInput `pulumi:"drds"`
	Dts                 pulumi.StringPtrInput `pulumi:"dts"`
	Dysms               pulumi.StringPtrInput `pulumi:"dysms"`
	Eais                pulumi.StringPtrInput `pulumi:"eais"`
	Eci                 pulumi.StringPtrInput `pulumi:"eci"`
	Ecs                 pulumi.StringPtrInput `pulumi:"ecs"`
	Edas                pulumi.StringPtrInput `pulumi:"edas"`
	Edasschedulerx      pulumi.StringPtrInput `pulumi:"edasschedulerx"`
	Edsuser             pulumi.StringPtrInput `pulumi:"edsuser"`
	Ehpc                pulumi.StringPtrInput `pulumi:"ehpc"`
	Ehs                 pulumi.StringPtrInput `pulumi:"ehs"`
	Eipanycast          pulumi.StringPtrInput `pulumi:"eipanycast"`
	Elasticsearch       pulumi.StringPtrInput `pulumi:"elasticsearch"`
	Emr                 pulumi.StringPtrInput `pulumi:"emr"`
	Ens                 pulumi.StringPtrInput `pulumi:"ens"`
	Ess                 pulumi.StringPtrInput `pulumi:"ess"`
	Eventbridge         pulumi.StringPtrInput `pulumi:"eventbridge"`
	Fc                  pulumi.StringPtrInput `pulumi:"fc"`
	Fnf                 pulumi.StringPtrInput `pulumi:"fnf"`
	Ga                  pulumi.StringPtrInput `pulumi:"ga"`
	Gaplus              pulumi.StringPtrInput `pulumi:"gaplus"`
	Gds                 pulumi.StringPtrInput `pulumi:"gds"`
	Gpdb                pulumi.StringPtrInput `pulumi:"gpdb"`
	Gwsecd              pulumi.StringPtrInput `pulumi:"gwsecd"`
	Hbr                 pulumi.StringPtrInput `pulumi:"hbr"`
	HcsSgw              pulumi.StringPtrInput `pulumi:"hcsSgw"`
	Hitsdb              pulumi.StringPtrInput `pulumi:"hitsdb"`
	Imm                 pulumi.StringPtrInput `pulumi:"imm"`
	Imp                 pulumi.StringPtrInput `pulumi:"imp"`
	Ims                 pulumi.StringPtrInput `pulumi:"ims"`
	Iot                 pulumi.StringPtrInput `pulumi:"iot"`
	Kms                 pulumi.StringPtrInput `pulumi:"kms"`
	Kvstore             pulumi.StringPtrInput `pulumi:"kvstore"`
	Location            pulumi.StringPtrInput `pulumi:"location"`
	Log                 pulumi.StringPtrInput `pulumi:"log"`
	Market              pulumi.StringPtrInput `pulumi:"market"`
	Maxcompute          pulumi.StringPtrInput `pulumi:"maxcompute"`
	Mhub                pulumi.StringPtrInput `pulumi:"mhub"`
	Mns                 pulumi.StringPtrInput `pulumi:"mns"`
	Mscopensubscription pulumi.StringPtrInput `pulumi:"mscopensubscription"`
	Mse                 pulumi.StringPtrInput `pulumi:"mse"`
	Nas                 pulumi.StringPtrInput `pulumi:"nas"`
	Ons                 pulumi.StringPtrInput `pulumi:"ons"`
	Onsproxy            pulumi.StringPtrInput `pulumi:"onsproxy"`
	Oos                 pulumi.StringPtrInput `pulumi:"oos"`
	Opensearch          pulumi.StringPtrInput `pulumi:"opensearch"`
	Oss                 pulumi.StringPtrInput `pulumi:"oss"`
	Ots                 pulumi.StringPtrInput `pulumi:"ots"`
	Polardb             pulumi.StringPtrInput `pulumi:"polardb"`
	Privatelink         pulumi.StringPtrInput `pulumi:"privatelink"`
	Pvtz                pulumi.StringPtrInput `pulumi:"pvtz"`
	Quickbi             pulumi.StringPtrInput `pulumi:"quickbi"`
	Quotas              pulumi.StringPtrInput `pulumi:"quotas"`
	RKvstore            pulumi.StringPtrInput `pulumi:"rKvstore"`
	Ram                 pulumi.StringPtrInput `pulumi:"ram"`
	Rds                 pulumi.StringPtrInput `pulumi:"rds"`
	Redisa              pulumi.StringPtrInput `pulumi:"redisa"`
	Resourcemanager     pulumi.StringPtrInput `pulumi:"resourcemanager"`
	Resourcesharing     pulumi.StringPtrInput `pulumi:"resourcesharing"`
	Ros                 pulumi.StringPtrInput `pulumi:"ros"`
	Sas                 pulumi.StringPtrInput `pulumi:"sas"`
	Scdn                pulumi.StringPtrInput `pulumi:"scdn"`
	Sddp                pulumi.StringPtrInput `pulumi:"sddp"`
	Serverless          pulumi.StringPtrInput `pulumi:"serverless"`
	Servicemesh         pulumi.StringPtrInput `pulumi:"servicemesh"`
	Sgw                 pulumi.StringPtrInput `pulumi:"sgw"`
	Slb                 pulumi.StringPtrInput `pulumi:"slb"`
	Smartag             pulumi.StringPtrInput `pulumi:"smartag"`
	Sts                 pulumi.StringPtrInput `pulumi:"sts"`
	Swas                pulumi.StringPtrInput `pulumi:"swas"`
	Tag                 pulumi.StringPtrInput `pulumi:"tag"`
	Vod                 pulumi.StringPtrInput `pulumi:"vod"`
	Vpc                 pulumi.StringPtrInput `pulumi:"vpc"`
	Vs                  pulumi.StringPtrInput `pulumi:"vs"`
	Waf                 pulumi.StringPtrInput `pulumi:"waf"`
	WafOpenapi          pulumi.StringPtrInput `pulumi:"wafOpenapi"`
}

func (EndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoints)(nil)).Elem()
}

func (i EndpointsArgs) ToEndpointsOutput() EndpointsOutput {
	return i.ToEndpointsOutputWithContext(context.Background())
}

func (i EndpointsArgs) ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsOutput)
}

// EndpointsArrayInput is an input type that accepts EndpointsArray and EndpointsArrayOutput values.
// You can construct a concrete instance of `EndpointsArrayInput` via:
//
//	EndpointsArray{ EndpointsArgs{...} }
type EndpointsArrayInput interface {
	pulumi.Input

	ToEndpointsArrayOutput() EndpointsArrayOutput
	ToEndpointsArrayOutputWithContext(context.Context) EndpointsArrayOutput
}

type EndpointsArray []EndpointsInput

func (EndpointsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoints)(nil)).Elem()
}

func (i EndpointsArray) ToEndpointsArrayOutput() EndpointsArrayOutput {
	return i.ToEndpointsArrayOutputWithContext(context.Background())
}

func (i EndpointsArray) ToEndpointsArrayOutputWithContext(ctx context.Context) EndpointsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsArrayOutput)
}

type EndpointsOutput struct{ *pulumi.OutputState }

func (EndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoints)(nil)).Elem()
}

func (o EndpointsOutput) ToEndpointsOutput() EndpointsOutput {
	return o
}

func (o EndpointsOutput) ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput {
	return o
}

func (o EndpointsOutput) Acr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Acr }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Actiontrail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Actiontrail }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Adb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Adb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Alb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Alb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Alidfs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Alidfs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Alidns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Alidns }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Alikafka() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Alikafka }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Apigateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Apigateway }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Arms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Arms }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Bastionhost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Bastionhost }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) BrainIndustrial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.BrainIndustrial }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Bssopenapi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Bssopenapi }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cassandra() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cassandra }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cbn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cbn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cddc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cddc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cdn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Clickhouse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Clickhouse }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cloudauth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cloudauth }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cloudfw() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cloudfw }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cloudphone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cloudphone }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cloudsso() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cloudsso }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cms }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Config }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cr }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Cs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Cs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Datahub() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Datahub }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dataworkspublic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dataworkspublic }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dbfs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dbfs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dcdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dcdn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ddosbasic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ddosbasic }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ddosbgp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ddosbgp }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ddoscoo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ddoscoo }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Devopsrdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Devopsrdc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dg() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dg }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dm }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) DmsEnterprise() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.DmsEnterprise }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dns }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Drds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Drds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dts }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Dysms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Dysms }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Eais() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Eais }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Eci() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Eci }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ecs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ecs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Edas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Edas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Edasschedulerx() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Edasschedulerx }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Edsuser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Edsuser }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ehpc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ehpc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ehs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ehs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Eipanycast() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Eipanycast }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Elasticsearch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Elasticsearch }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Emr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Emr }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ens() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ens }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ess }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Eventbridge() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Eventbridge }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Fc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Fc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Fnf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Fnf }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ga() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ga }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Gaplus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Gaplus }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Gds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Gds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Gpdb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Gpdb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Gwsecd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Gwsecd }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Hbr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Hbr }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) HcsSgw() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.HcsSgw }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Hitsdb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Hitsdb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Imm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Imm }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Imp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Imp }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ims() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ims }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Iot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Iot }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Kms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Kms }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Kvstore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Kvstore }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Log() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Log }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Market() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Market }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Maxcompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Maxcompute }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Mhub() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Mhub }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Mns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Mns }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Mscopensubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Mscopensubscription }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Mse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Mse }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Nas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Nas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ons() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ons }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Onsproxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Onsproxy }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Oos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Oos }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Opensearch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Opensearch }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Oss() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Oss }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ots() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ots }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Polardb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Polardb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Privatelink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Privatelink }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Pvtz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Pvtz }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Quickbi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Quickbi }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Quotas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Quotas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) RKvstore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.RKvstore }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ram() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ram }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Rds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Rds }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Redisa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Redisa }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Resourcemanager() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Resourcemanager }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Resourcesharing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Resourcesharing }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Ros() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Ros }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Sas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Sas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Scdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Scdn }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Sddp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Sddp }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Serverless() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Serverless }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Servicemesh() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Servicemesh }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Sgw() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Sgw }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Slb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Slb }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Smartag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Smartag }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Sts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Sts }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Swas() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Swas }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Vod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Vod }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Vpc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Vpc }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Vs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Vs }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) Waf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.Waf }).(pulumi.StringPtrOutput)
}

func (o EndpointsOutput) WafOpenapi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Endpoints) *string { return v.WafOpenapi }).(pulumi.StringPtrOutput)
}

type EndpointsArrayOutput struct{ *pulumi.OutputState }

func (EndpointsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoints)(nil)).Elem()
}

func (o EndpointsArrayOutput) ToEndpointsArrayOutput() EndpointsArrayOutput {
	return o
}

func (o EndpointsArrayOutput) ToEndpointsArrayOutputWithContext(ctx context.Context) EndpointsArrayOutput {
	return o
}

func (o EndpointsArrayOutput) Index(i pulumi.IntInput) EndpointsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Endpoints {
		return vs[0].([]Endpoints)[vs[1].(int)]
	}).(EndpointsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssumeRoleInput)(nil)).Elem(), AssumeRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsInput)(nil)).Elem(), EndpointsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsArrayInput)(nil)).Elem(), EndpointsArray{})
	pulumi.RegisterOutputType(AssumeRoleOutput{})
	pulumi.RegisterOutputType(EndpointsOutput{})
	pulumi.RegisterOutputType(EndpointsArrayOutput{})
}
