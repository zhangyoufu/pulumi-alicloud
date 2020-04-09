// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pvtz

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ZoneRecord struct {
	pulumi.CustomResourceState

	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-50]. Default to 1.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The Private Zone Record ID.
	RecordId pulumi.IntOutput `pulumi:"recordId"`
	// The resource record of the Private Zone Record.
	ResourceRecord pulumi.StringOutput `pulumi:"resourceRecord"`
	// The ttl of the Private Zone Record.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the Private Zone Record.
	Value pulumi.StringOutput `pulumi:"value"`
	// The name of the Private Zone Record.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneRecord registers a new resource with the given unique name, arguments, and options.
func NewZoneRecord(ctx *pulumi.Context,
	name string, args *ZoneRecordArgs, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	if args == nil || args.ResourceRecord == nil {
		return nil, errors.New("missing required argument 'ResourceRecord'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &ZoneRecordArgs{}
	}
	var resource ZoneRecord
	err := ctx.RegisterResource("alicloud:pvtz/zoneRecord:ZoneRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneRecord gets an existing ZoneRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneRecordState, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	var resource ZoneRecord
	err := ctx.ReadResource("alicloud:pvtz/zoneRecord:ZoneRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneRecord resources.
type zoneRecordState struct {
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-50]. Default to 1.
	Priority *int `pulumi:"priority"`
	// The Private Zone Record ID.
	RecordId *int `pulumi:"recordId"`
	// The resource record of the Private Zone Record.
	ResourceRecord *string `pulumi:"resourceRecord"`
	// The ttl of the Private Zone Record.
	Ttl *int `pulumi:"ttl"`
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR.
	Type *string `pulumi:"type"`
	// The value of the Private Zone Record.
	Value *string `pulumi:"value"`
	// The name of the Private Zone Record.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneRecordState struct {
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-50]. Default to 1.
	Priority pulumi.IntPtrInput
	// The Private Zone Record ID.
	RecordId pulumi.IntPtrInput
	// The resource record of the Private Zone Record.
	ResourceRecord pulumi.StringPtrInput
	// The ttl of the Private Zone Record.
	Ttl pulumi.IntPtrInput
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR.
	Type pulumi.StringPtrInput
	// The value of the Private Zone Record.
	Value pulumi.StringPtrInput
	// The name of the Private Zone Record.
	ZoneId pulumi.StringPtrInput
}

func (ZoneRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordState)(nil)).Elem()
}

type zoneRecordArgs struct {
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-50]. Default to 1.
	Priority *int `pulumi:"priority"`
	// The resource record of the Private Zone Record.
	ResourceRecord string `pulumi:"resourceRecord"`
	// The ttl of the Private Zone Record.
	Ttl *int `pulumi:"ttl"`
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR.
	Type string `pulumi:"type"`
	// The value of the Private Zone Record.
	Value string `pulumi:"value"`
	// The name of the Private Zone Record.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneRecord resource.
type ZoneRecordArgs struct {
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-50]. Default to 1.
	Priority pulumi.IntPtrInput
	// The resource record of the Private Zone Record.
	ResourceRecord pulumi.StringInput
	// The ttl of the Private Zone Record.
	Ttl pulumi.IntPtrInput
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR.
	Type pulumi.StringInput
	// The value of the Private Zone Record.
	Value pulumi.StringInput
	// The name of the Private Zone Record.
	ZoneId pulumi.StringInput
}

func (ZoneRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordArgs)(nil)).Elem()
}
