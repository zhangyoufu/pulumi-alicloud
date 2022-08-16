// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directmail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetDomainsDomain struct {
	// Track verification.
	CnameAuthStatus string `pulumi:"cnameAuthStatus"`
	// Indicates whether the CNAME record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
	CnameConfirmStatus string `pulumi:"cnameConfirmStatus"`
	// The value of the CNAME record.
	CnameRecord string `pulumi:"cnameRecord"`
	// The time when the DNS record was created.
	CreateTime string `pulumi:"createTime"`
	// The default domain name.
	DefaultDomain string `pulumi:"defaultDomain"`
	// The value of the MX record.
	DnsMx string `pulumi:"dnsMx"`
	// The value of the SPF record.
	DnsSpf string `pulumi:"dnsSpf"`
	// The value of the TXT ownership record.
	DnsTxt string `pulumi:"dnsTxt"`
	// The ID of the domain name.
	DomainId string `pulumi:"domainId"`
	// The domain name.
	DomainName string `pulumi:"domainName"`
	// The type of the domain.
	DomainType string `pulumi:"domainType"`
	// The status of ICP filing. Valid values: `0` and `1`. `0`: indicates that the domain name is not filed. `1`: indicates that the domain name is filed.
	IcpStatus string `pulumi:"icpStatus"`
	// The ID of the Domain.
	Id string `pulumi:"id"`
	// Indicates whether the MX record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
	MxAuthStatus string `pulumi:"mxAuthStatus"`
	// The MX verification record provided by Alibaba Cloud DNS.
	MxRecord string `pulumi:"mxRecord"`
	// Indicates whether the SPF record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
	SpfAuthStatus string `pulumi:"spfAuthStatus"`
	// The SPF verification record provided by Alibaba Cloud DNS.
	SpfRecord string `pulumi:"spfRecord"`
	// The status of the domain name. Valid values:`0` to `4`. `0`:Available, Passed. `1`: Unavailable, No passed. `2`: Available, cname no passed, icp no passed. `3`: Available, icp no passed. `4`: Available, cname no passed.
	Status string `pulumi:"status"`
	// The primary domain name.
	TlDomainName string `pulumi:"tlDomainName"`
	// The CNAME verification record provided by Alibaba Cloud DNS.
	TracefRecord string `pulumi:"tracefRecord"`
}

// GetDomainsDomainInput is an input type that accepts GetDomainsDomainArgs and GetDomainsDomainOutput values.
// You can construct a concrete instance of `GetDomainsDomainInput` via:
//
//	GetDomainsDomainArgs{...}
type GetDomainsDomainInput interface {
	pulumi.Input

	ToGetDomainsDomainOutput() GetDomainsDomainOutput
	ToGetDomainsDomainOutputWithContext(context.Context) GetDomainsDomainOutput
}

type GetDomainsDomainArgs struct {
	// Track verification.
	CnameAuthStatus pulumi.StringInput `pulumi:"cnameAuthStatus"`
	// Indicates whether the CNAME record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
	CnameConfirmStatus pulumi.StringInput `pulumi:"cnameConfirmStatus"`
	// The value of the CNAME record.
	CnameRecord pulumi.StringInput `pulumi:"cnameRecord"`
	// The time when the DNS record was created.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The default domain name.
	DefaultDomain pulumi.StringInput `pulumi:"defaultDomain"`
	// The value of the MX record.
	DnsMx pulumi.StringInput `pulumi:"dnsMx"`
	// The value of the SPF record.
	DnsSpf pulumi.StringInput `pulumi:"dnsSpf"`
	// The value of the TXT ownership record.
	DnsTxt pulumi.StringInput `pulumi:"dnsTxt"`
	// The ID of the domain name.
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The domain name.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The type of the domain.
	DomainType pulumi.StringInput `pulumi:"domainType"`
	// The status of ICP filing. Valid values: `0` and `1`. `0`: indicates that the domain name is not filed. `1`: indicates that the domain name is filed.
	IcpStatus pulumi.StringInput `pulumi:"icpStatus"`
	// The ID of the Domain.
	Id pulumi.StringInput `pulumi:"id"`
	// Indicates whether the MX record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
	MxAuthStatus pulumi.StringInput `pulumi:"mxAuthStatus"`
	// The MX verification record provided by Alibaba Cloud DNS.
	MxRecord pulumi.StringInput `pulumi:"mxRecord"`
	// Indicates whether the SPF record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
	SpfAuthStatus pulumi.StringInput `pulumi:"spfAuthStatus"`
	// The SPF verification record provided by Alibaba Cloud DNS.
	SpfRecord pulumi.StringInput `pulumi:"spfRecord"`
	// The status of the domain name. Valid values:`0` to `4`. `0`:Available, Passed. `1`: Unavailable, No passed. `2`: Available, cname no passed, icp no passed. `3`: Available, icp no passed. `4`: Available, cname no passed.
	Status pulumi.StringInput `pulumi:"status"`
	// The primary domain name.
	TlDomainName pulumi.StringInput `pulumi:"tlDomainName"`
	// The CNAME verification record provided by Alibaba Cloud DNS.
	TracefRecord pulumi.StringInput `pulumi:"tracefRecord"`
}

func (GetDomainsDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomain)(nil)).Elem()
}

func (i GetDomainsDomainArgs) ToGetDomainsDomainOutput() GetDomainsDomainOutput {
	return i.ToGetDomainsDomainOutputWithContext(context.Background())
}

func (i GetDomainsDomainArgs) ToGetDomainsDomainOutputWithContext(ctx context.Context) GetDomainsDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainOutput)
}

// GetDomainsDomainArrayInput is an input type that accepts GetDomainsDomainArray and GetDomainsDomainArrayOutput values.
// You can construct a concrete instance of `GetDomainsDomainArrayInput` via:
//
//	GetDomainsDomainArray{ GetDomainsDomainArgs{...} }
type GetDomainsDomainArrayInput interface {
	pulumi.Input

	ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput
	ToGetDomainsDomainArrayOutputWithContext(context.Context) GetDomainsDomainArrayOutput
}

type GetDomainsDomainArray []GetDomainsDomainInput

func (GetDomainsDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomain)(nil)).Elem()
}

func (i GetDomainsDomainArray) ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput {
	return i.ToGetDomainsDomainArrayOutputWithContext(context.Background())
}

func (i GetDomainsDomainArray) ToGetDomainsDomainArrayOutputWithContext(ctx context.Context) GetDomainsDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDomainsDomainArrayOutput)
}

type GetDomainsDomainOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsDomain)(nil)).Elem()
}

func (o GetDomainsDomainOutput) ToGetDomainsDomainOutput() GetDomainsDomainOutput {
	return o
}

func (o GetDomainsDomainOutput) ToGetDomainsDomainOutputWithContext(ctx context.Context) GetDomainsDomainOutput {
	return o
}

// Track verification.
func (o GetDomainsDomainOutput) CnameAuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.CnameAuthStatus }).(pulumi.StringOutput)
}

// Indicates whether the CNAME record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
func (o GetDomainsDomainOutput) CnameConfirmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.CnameConfirmStatus }).(pulumi.StringOutput)
}

// The value of the CNAME record.
func (o GetDomainsDomainOutput) CnameRecord() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.CnameRecord }).(pulumi.StringOutput)
}

// The time when the DNS record was created.
func (o GetDomainsDomainOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The default domain name.
func (o GetDomainsDomainOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DefaultDomain }).(pulumi.StringOutput)
}

// The value of the MX record.
func (o GetDomainsDomainOutput) DnsMx() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DnsMx }).(pulumi.StringOutput)
}

// The value of the SPF record.
func (o GetDomainsDomainOutput) DnsSpf() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DnsSpf }).(pulumi.StringOutput)
}

// The value of the TXT ownership record.
func (o GetDomainsDomainOutput) DnsTxt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DnsTxt }).(pulumi.StringOutput)
}

// The ID of the domain name.
func (o GetDomainsDomainOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DomainId }).(pulumi.StringOutput)
}

// The domain name.
func (o GetDomainsDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DomainName }).(pulumi.StringOutput)
}

// The type of the domain.
func (o GetDomainsDomainOutput) DomainType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.DomainType }).(pulumi.StringOutput)
}

// The status of ICP filing. Valid values: `0` and `1`. `0`: indicates that the domain name is not filed. `1`: indicates that the domain name is filed.
func (o GetDomainsDomainOutput) IcpStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.IcpStatus }).(pulumi.StringOutput)
}

// The ID of the Domain.
func (o GetDomainsDomainOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether the MX record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
func (o GetDomainsDomainOutput) MxAuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.MxAuthStatus }).(pulumi.StringOutput)
}

// The MX verification record provided by Alibaba Cloud DNS.
func (o GetDomainsDomainOutput) MxRecord() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.MxRecord }).(pulumi.StringOutput)
}

// Indicates whether the SPF record is successfully verified. Valid values: `0` and `1`. `0`: indicates the verification is successful. `1`: indicates that the verification fails.
func (o GetDomainsDomainOutput) SpfAuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.SpfAuthStatus }).(pulumi.StringOutput)
}

// The SPF verification record provided by Alibaba Cloud DNS.
func (o GetDomainsDomainOutput) SpfRecord() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.SpfRecord }).(pulumi.StringOutput)
}

// The status of the domain name. Valid values:`0` to `4`. `0`:Available, Passed. `1`: Unavailable, No passed. `2`: Available, cname no passed, icp no passed. `3`: Available, icp no passed. `4`: Available, cname no passed.
func (o GetDomainsDomainOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.Status }).(pulumi.StringOutput)
}

// The primary domain name.
func (o GetDomainsDomainOutput) TlDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.TlDomainName }).(pulumi.StringOutput)
}

// The CNAME verification record provided by Alibaba Cloud DNS.
func (o GetDomainsDomainOutput) TracefRecord() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsDomain) string { return v.TracefRecord }).(pulumi.StringOutput)
}

type GetDomainsDomainArrayOutput struct{ *pulumi.OutputState }

func (GetDomainsDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDomainsDomain)(nil)).Elem()
}

func (o GetDomainsDomainArrayOutput) ToGetDomainsDomainArrayOutput() GetDomainsDomainArrayOutput {
	return o
}

func (o GetDomainsDomainArrayOutput) ToGetDomainsDomainArrayOutputWithContext(ctx context.Context) GetDomainsDomainArrayOutput {
	return o
}

func (o GetDomainsDomainArrayOutput) Index(i pulumi.IntInput) GetDomainsDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDomainsDomain {
		return vs[0].([]GetDomainsDomain)[vs[1].(int)]
	}).(GetDomainsDomainOutput)
}

type GetMailAddressesAddress struct {
	// The sender address.
	AccountName string `pulumi:"accountName"`
	// The creation of the record time.
	CreateTime string `pulumi:"createTime"`
	// On the quota limit.
	DailyCount string `pulumi:"dailyCount"`
	// On the quota.
	DailyReqCount string `pulumi:"dailyReqCount"`
	// Domain name status. Valid values: `0`, `1`.
	DomainStatus string `pulumi:"domainStatus"`
	// The ID of the Mail Address.
	Id string `pulumi:"id"`
	// The sender address ID.
	MailAddressId string `pulumi:"mailAddressId"`
	// Monthly quota limit.
	MonthCount string `pulumi:"monthCount"`
	// Months amount.
	MonthReqCount string `pulumi:"monthReqCount"`
	// Return address.
	ReplyAddress string `pulumi:"replyAddress"`
	// If using STMP address status.
	ReplyStatus string `pulumi:"replyStatus"`
	// Account type.
	Sendtype string `pulumi:"sendtype"`
	// Account Status. Valid values: `0`, `1`. Freeze: 1, normal: 0.
	Status string `pulumi:"status"`
}

// GetMailAddressesAddressInput is an input type that accepts GetMailAddressesAddressArgs and GetMailAddressesAddressOutput values.
// You can construct a concrete instance of `GetMailAddressesAddressInput` via:
//
//	GetMailAddressesAddressArgs{...}
type GetMailAddressesAddressInput interface {
	pulumi.Input

	ToGetMailAddressesAddressOutput() GetMailAddressesAddressOutput
	ToGetMailAddressesAddressOutputWithContext(context.Context) GetMailAddressesAddressOutput
}

type GetMailAddressesAddressArgs struct {
	// The sender address.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The creation of the record time.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// On the quota limit.
	DailyCount pulumi.StringInput `pulumi:"dailyCount"`
	// On the quota.
	DailyReqCount pulumi.StringInput `pulumi:"dailyReqCount"`
	// Domain name status. Valid values: `0`, `1`.
	DomainStatus pulumi.StringInput `pulumi:"domainStatus"`
	// The ID of the Mail Address.
	Id pulumi.StringInput `pulumi:"id"`
	// The sender address ID.
	MailAddressId pulumi.StringInput `pulumi:"mailAddressId"`
	// Monthly quota limit.
	MonthCount pulumi.StringInput `pulumi:"monthCount"`
	// Months amount.
	MonthReqCount pulumi.StringInput `pulumi:"monthReqCount"`
	// Return address.
	ReplyAddress pulumi.StringInput `pulumi:"replyAddress"`
	// If using STMP address status.
	ReplyStatus pulumi.StringInput `pulumi:"replyStatus"`
	// Account type.
	Sendtype pulumi.StringInput `pulumi:"sendtype"`
	// Account Status. Valid values: `0`, `1`. Freeze: 1, normal: 0.
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetMailAddressesAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMailAddressesAddress)(nil)).Elem()
}

func (i GetMailAddressesAddressArgs) ToGetMailAddressesAddressOutput() GetMailAddressesAddressOutput {
	return i.ToGetMailAddressesAddressOutputWithContext(context.Background())
}

func (i GetMailAddressesAddressArgs) ToGetMailAddressesAddressOutputWithContext(ctx context.Context) GetMailAddressesAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMailAddressesAddressOutput)
}

// GetMailAddressesAddressArrayInput is an input type that accepts GetMailAddressesAddressArray and GetMailAddressesAddressArrayOutput values.
// You can construct a concrete instance of `GetMailAddressesAddressArrayInput` via:
//
//	GetMailAddressesAddressArray{ GetMailAddressesAddressArgs{...} }
type GetMailAddressesAddressArrayInput interface {
	pulumi.Input

	ToGetMailAddressesAddressArrayOutput() GetMailAddressesAddressArrayOutput
	ToGetMailAddressesAddressArrayOutputWithContext(context.Context) GetMailAddressesAddressArrayOutput
}

type GetMailAddressesAddressArray []GetMailAddressesAddressInput

func (GetMailAddressesAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMailAddressesAddress)(nil)).Elem()
}

func (i GetMailAddressesAddressArray) ToGetMailAddressesAddressArrayOutput() GetMailAddressesAddressArrayOutput {
	return i.ToGetMailAddressesAddressArrayOutputWithContext(context.Background())
}

func (i GetMailAddressesAddressArray) ToGetMailAddressesAddressArrayOutputWithContext(ctx context.Context) GetMailAddressesAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetMailAddressesAddressArrayOutput)
}

type GetMailAddressesAddressOutput struct{ *pulumi.OutputState }

func (GetMailAddressesAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMailAddressesAddress)(nil)).Elem()
}

func (o GetMailAddressesAddressOutput) ToGetMailAddressesAddressOutput() GetMailAddressesAddressOutput {
	return o
}

func (o GetMailAddressesAddressOutput) ToGetMailAddressesAddressOutputWithContext(ctx context.Context) GetMailAddressesAddressOutput {
	return o
}

// The sender address.
func (o GetMailAddressesAddressOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.AccountName }).(pulumi.StringOutput)
}

// The creation of the record time.
func (o GetMailAddressesAddressOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.CreateTime }).(pulumi.StringOutput)
}

// On the quota limit.
func (o GetMailAddressesAddressOutput) DailyCount() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.DailyCount }).(pulumi.StringOutput)
}

// On the quota.
func (o GetMailAddressesAddressOutput) DailyReqCount() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.DailyReqCount }).(pulumi.StringOutput)
}

// Domain name status. Valid values: `0`, `1`.
func (o GetMailAddressesAddressOutput) DomainStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.DomainStatus }).(pulumi.StringOutput)
}

// The ID of the Mail Address.
func (o GetMailAddressesAddressOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.Id }).(pulumi.StringOutput)
}

// The sender address ID.
func (o GetMailAddressesAddressOutput) MailAddressId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.MailAddressId }).(pulumi.StringOutput)
}

// Monthly quota limit.
func (o GetMailAddressesAddressOutput) MonthCount() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.MonthCount }).(pulumi.StringOutput)
}

// Months amount.
func (o GetMailAddressesAddressOutput) MonthReqCount() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.MonthReqCount }).(pulumi.StringOutput)
}

// Return address.
func (o GetMailAddressesAddressOutput) ReplyAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.ReplyAddress }).(pulumi.StringOutput)
}

// If using STMP address status.
func (o GetMailAddressesAddressOutput) ReplyStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.ReplyStatus }).(pulumi.StringOutput)
}

// Account type.
func (o GetMailAddressesAddressOutput) Sendtype() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.Sendtype }).(pulumi.StringOutput)
}

// Account Status. Valid values: `0`, `1`. Freeze: 1, normal: 0.
func (o GetMailAddressesAddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetMailAddressesAddress) string { return v.Status }).(pulumi.StringOutput)
}

type GetMailAddressesAddressArrayOutput struct{ *pulumi.OutputState }

func (GetMailAddressesAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetMailAddressesAddress)(nil)).Elem()
}

func (o GetMailAddressesAddressArrayOutput) ToGetMailAddressesAddressArrayOutput() GetMailAddressesAddressArrayOutput {
	return o
}

func (o GetMailAddressesAddressArrayOutput) ToGetMailAddressesAddressArrayOutputWithContext(ctx context.Context) GetMailAddressesAddressArrayOutput {
	return o
}

func (o GetMailAddressesAddressArrayOutput) Index(i pulumi.IntInput) GetMailAddressesAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetMailAddressesAddress {
		return vs[0].([]GetMailAddressesAddress)[vs[1].(int)]
	}).(GetMailAddressesAddressOutput)
}

type GetReceiversReceiverse struct {
	// The creation time of the resource.
	CreateTime string `pulumi:"createTime"`
	// The description.
	Description string `pulumi:"description"`
	// The ID of the Receivers.
	Id string `pulumi:"id"`
	// -The Receivers Alias.
	ReceiversAlias string `pulumi:"receiversAlias"`
	// The first ID of the resource.
	ReceiversId string `pulumi:"receiversId"`
	// The name of the resource.
	ReceiversName string `pulumi:"receiversName"`
	// The status of the resource.
	Status int `pulumi:"status"`
}

// GetReceiversReceiverseInput is an input type that accepts GetReceiversReceiverseArgs and GetReceiversReceiverseOutput values.
// You can construct a concrete instance of `GetReceiversReceiverseInput` via:
//
//	GetReceiversReceiverseArgs{...}
type GetReceiversReceiverseInput interface {
	pulumi.Input

	ToGetReceiversReceiverseOutput() GetReceiversReceiverseOutput
	ToGetReceiversReceiverseOutputWithContext(context.Context) GetReceiversReceiverseOutput
}

type GetReceiversReceiverseArgs struct {
	// The creation time of the resource.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The description.
	Description pulumi.StringInput `pulumi:"description"`
	// The ID of the Receivers.
	Id pulumi.StringInput `pulumi:"id"`
	// -The Receivers Alias.
	ReceiversAlias pulumi.StringInput `pulumi:"receiversAlias"`
	// The first ID of the resource.
	ReceiversId pulumi.StringInput `pulumi:"receiversId"`
	// The name of the resource.
	ReceiversName pulumi.StringInput `pulumi:"receiversName"`
	// The status of the resource.
	Status pulumi.IntInput `pulumi:"status"`
}

func (GetReceiversReceiverseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReceiversReceiverse)(nil)).Elem()
}

func (i GetReceiversReceiverseArgs) ToGetReceiversReceiverseOutput() GetReceiversReceiverseOutput {
	return i.ToGetReceiversReceiverseOutputWithContext(context.Background())
}

func (i GetReceiversReceiverseArgs) ToGetReceiversReceiverseOutputWithContext(ctx context.Context) GetReceiversReceiverseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetReceiversReceiverseOutput)
}

// GetReceiversReceiverseArrayInput is an input type that accepts GetReceiversReceiverseArray and GetReceiversReceiverseArrayOutput values.
// You can construct a concrete instance of `GetReceiversReceiverseArrayInput` via:
//
//	GetReceiversReceiverseArray{ GetReceiversReceiverseArgs{...} }
type GetReceiversReceiverseArrayInput interface {
	pulumi.Input

	ToGetReceiversReceiverseArrayOutput() GetReceiversReceiverseArrayOutput
	ToGetReceiversReceiverseArrayOutputWithContext(context.Context) GetReceiversReceiverseArrayOutput
}

type GetReceiversReceiverseArray []GetReceiversReceiverseInput

func (GetReceiversReceiverseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetReceiversReceiverse)(nil)).Elem()
}

func (i GetReceiversReceiverseArray) ToGetReceiversReceiverseArrayOutput() GetReceiversReceiverseArrayOutput {
	return i.ToGetReceiversReceiverseArrayOutputWithContext(context.Background())
}

func (i GetReceiversReceiverseArray) ToGetReceiversReceiverseArrayOutputWithContext(ctx context.Context) GetReceiversReceiverseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetReceiversReceiverseArrayOutput)
}

type GetReceiversReceiverseOutput struct{ *pulumi.OutputState }

func (GetReceiversReceiverseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReceiversReceiverse)(nil)).Elem()
}

func (o GetReceiversReceiverseOutput) ToGetReceiversReceiverseOutput() GetReceiversReceiverseOutput {
	return o
}

func (o GetReceiversReceiverseOutput) ToGetReceiversReceiverseOutputWithContext(ctx context.Context) GetReceiversReceiverseOutput {
	return o
}

// The creation time of the resource.
func (o GetReceiversReceiverseOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The description.
func (o GetReceiversReceiverseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of the Receivers.
func (o GetReceiversReceiverseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) string { return v.Id }).(pulumi.StringOutput)
}

// -The Receivers Alias.
func (o GetReceiversReceiverseOutput) ReceiversAlias() pulumi.StringOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) string { return v.ReceiversAlias }).(pulumi.StringOutput)
}

// The first ID of the resource.
func (o GetReceiversReceiverseOutput) ReceiversId() pulumi.StringOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) string { return v.ReceiversId }).(pulumi.StringOutput)
}

// The name of the resource.
func (o GetReceiversReceiverseOutput) ReceiversName() pulumi.StringOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) string { return v.ReceiversName }).(pulumi.StringOutput)
}

// The status of the resource.
func (o GetReceiversReceiverseOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetReceiversReceiverse) int { return v.Status }).(pulumi.IntOutput)
}

type GetReceiversReceiverseArrayOutput struct{ *pulumi.OutputState }

func (GetReceiversReceiverseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetReceiversReceiverse)(nil)).Elem()
}

func (o GetReceiversReceiverseArrayOutput) ToGetReceiversReceiverseArrayOutput() GetReceiversReceiverseArrayOutput {
	return o
}

func (o GetReceiversReceiverseArrayOutput) ToGetReceiversReceiverseArrayOutputWithContext(ctx context.Context) GetReceiversReceiverseArrayOutput {
	return o
}

func (o GetReceiversReceiverseArrayOutput) Index(i pulumi.IntInput) GetReceiversReceiverseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetReceiversReceiverse {
		return vs[0].([]GetReceiversReceiverse)[vs[1].(int)]
	}).(GetReceiversReceiverseOutput)
}

type GetTagsTag struct {
	// The ID of the tag.
	Id string `pulumi:"id"`
	// The ID of the tag.
	TagId string `pulumi:"tagId"`
	// The name of the tag.
	TagName string `pulumi:"tagName"`
}

// GetTagsTagInput is an input type that accepts GetTagsTagArgs and GetTagsTagOutput values.
// You can construct a concrete instance of `GetTagsTagInput` via:
//
//	GetTagsTagArgs{...}
type GetTagsTagInput interface {
	pulumi.Input

	ToGetTagsTagOutput() GetTagsTagOutput
	ToGetTagsTagOutputWithContext(context.Context) GetTagsTagOutput
}

type GetTagsTagArgs struct {
	// The ID of the tag.
	Id pulumi.StringInput `pulumi:"id"`
	// The ID of the tag.
	TagId pulumi.StringInput `pulumi:"tagId"`
	// The name of the tag.
	TagName pulumi.StringInput `pulumi:"tagName"`
}

func (GetTagsTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTag)(nil)).Elem()
}

func (i GetTagsTagArgs) ToGetTagsTagOutput() GetTagsTagOutput {
	return i.ToGetTagsTagOutputWithContext(context.Background())
}

func (i GetTagsTagArgs) ToGetTagsTagOutputWithContext(ctx context.Context) GetTagsTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTagsTagOutput)
}

// GetTagsTagArrayInput is an input type that accepts GetTagsTagArray and GetTagsTagArrayOutput values.
// You can construct a concrete instance of `GetTagsTagArrayInput` via:
//
//	GetTagsTagArray{ GetTagsTagArgs{...} }
type GetTagsTagArrayInput interface {
	pulumi.Input

	ToGetTagsTagArrayOutput() GetTagsTagArrayOutput
	ToGetTagsTagArrayOutputWithContext(context.Context) GetTagsTagArrayOutput
}

type GetTagsTagArray []GetTagsTagInput

func (GetTagsTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTagsTag)(nil)).Elem()
}

func (i GetTagsTagArray) ToGetTagsTagArrayOutput() GetTagsTagArrayOutput {
	return i.ToGetTagsTagArrayOutputWithContext(context.Background())
}

func (i GetTagsTagArray) ToGetTagsTagArrayOutputWithContext(ctx context.Context) GetTagsTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTagsTagArrayOutput)
}

type GetTagsTagOutput struct{ *pulumi.OutputState }

func (GetTagsTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTag)(nil)).Elem()
}

func (o GetTagsTagOutput) ToGetTagsTagOutput() GetTagsTagOutput {
	return o
}

func (o GetTagsTagOutput) ToGetTagsTagOutputWithContext(ctx context.Context) GetTagsTagOutput {
	return o
}

// The ID of the tag.
func (o GetTagsTagOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTag) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the tag.
func (o GetTagsTagOutput) TagId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTag) string { return v.TagId }).(pulumi.StringOutput)
}

// The name of the tag.
func (o GetTagsTagOutput) TagName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTag) string { return v.TagName }).(pulumi.StringOutput)
}

type GetTagsTagArrayOutput struct{ *pulumi.OutputState }

func (GetTagsTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTagsTag)(nil)).Elem()
}

func (o GetTagsTagArrayOutput) ToGetTagsTagArrayOutput() GetTagsTagArrayOutput {
	return o
}

func (o GetTagsTagArrayOutput) ToGetTagsTagArrayOutputWithContext(ctx context.Context) GetTagsTagArrayOutput {
	return o
}

func (o GetTagsTagArrayOutput) Index(i pulumi.IntInput) GetTagsTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTagsTag {
		return vs[0].([]GetTagsTag)[vs[1].(int)]
	}).(GetTagsTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainInput)(nil)).Elem(), GetDomainsDomainArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDomainsDomainArrayInput)(nil)).Elem(), GetDomainsDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMailAddressesAddressInput)(nil)).Elem(), GetMailAddressesAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetMailAddressesAddressArrayInput)(nil)).Elem(), GetMailAddressesAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetReceiversReceiverseInput)(nil)).Elem(), GetReceiversReceiverseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetReceiversReceiverseArrayInput)(nil)).Elem(), GetReceiversReceiverseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTagsTagInput)(nil)).Elem(), GetTagsTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTagsTagArrayInput)(nil)).Elem(), GetTagsTagArray{})
	pulumi.RegisterOutputType(GetDomainsDomainOutput{})
	pulumi.RegisterOutputType(GetDomainsDomainArrayOutput{})
	pulumi.RegisterOutputType(GetMailAddressesAddressOutput{})
	pulumi.RegisterOutputType(GetMailAddressesAddressArrayOutput{})
	pulumi.RegisterOutputType(GetReceiversReceiverseOutput{})
	pulumi.RegisterOutputType(GetReceiversReceiverseArrayOutput{})
	pulumi.RegisterOutputType(GetTagsTagOutput{})
	pulumi.RegisterOutputType(GetTagsTagArrayOutput{})
}
