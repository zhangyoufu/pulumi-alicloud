// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package polardb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a PolarDB endpoint resource to manage cluster endpoint of PolarDB cluster.
//
// > **NOTE:** Available since v1.217.0
//
// > **NOTE:** The default cluster endpoint can not be created or deleted manually.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/polardb"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := polardb.GetNodeClasses(ctx, &polardb.GetNodeClassesArgs{
//				DbType:    pulumi.StringRef("MySQL"),
//				DbVersion: pulumi.StringRef("8.0"),
//				PayType:   "PostPaid",
//				Category:  pulumi.StringRef("Normal"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "default", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("terraform-example"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "default", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      pulumi.String(_default.Classes[0].ZoneId),
//				VswitchName: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultCluster, err := polardb.NewCluster(ctx, "default", &polardb.ClusterArgs{
//				DbType:      pulumi.String("MySQL"),
//				DbVersion:   pulumi.String("8.0"),
//				DbNodeClass: pulumi.String(_default.Classes[0].SupportedEngines[0].AvailableResources[0].DbNodeClass),
//				PayType:     pulumi.String("PostPaid"),
//				VswitchId:   defaultSwitch.ID(),
//				Description: pulumi.String("terraform-example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = polardb.NewClusterEndpoint(ctx, "default", &polardb.ClusterEndpointArgs{
//				DbClusterId: defaultCluster.ID(),
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
// PolarDB endpoint can be imported using the id, e.g.
//
// ```sh
// $ pulumi import alicloud:polardb/clusterEndpoint:ClusterEndpoint example pc-abc123456:pe-abc123456
// ```
type ClusterEndpoint struct {
	pulumi.CustomResourceState

	// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
	AutoAddNewNodes pulumi.StringOutput `pulumi:"autoAddNewNodes"`
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix pulumi.StringOutput `pulumi:"connectionPrefix"`
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringOutput `pulumi:"dbClusterId"`
	// The name of the endpoint.
	DbEndpointDescription pulumi.StringPtrOutput `pulumi:"dbEndpointDescription"`
	// The ID of the cluster endpoint.
	DbEndpointId pulumi.StringOutput `pulumi:"dbEndpointId"`
	// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
	EndpointConfig pulumi.StringMapOutput `pulumi:"endpointConfig"`
	// Type of endpoint.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// The network type of the endpoint address.
	NetType pulumi.StringPtrOutput `pulumi:"netType"`
	// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
	Nodes pulumi.StringArrayOutput `pulumi:"nodes"`
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port pulumi.StringOutput `pulumi:"port"`
	// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
	ReadWriteMode pulumi.StringOutput `pulumi:"readWriteMode"`
	// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
	// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
	// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
	SslAutoRotate pulumi.StringPtrOutput `pulumi:"sslAutoRotate"`
	// The specifies SSL certificate download link.
	SslCertificateUrl pulumi.StringOutput `pulumi:"sslCertificateUrl"`
	// The SSL connection string.
	SslConnectionString pulumi.StringOutput `pulumi:"sslConnectionString"`
	// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
	SslEnabled pulumi.StringPtrOutput `pulumi:"sslEnabled"`
	// The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SslExpireTime pulumi.StringOutput `pulumi:"sslExpireTime"`
}

// NewClusterEndpoint registers a new resource with the given unique name, arguments, and options.
func NewClusterEndpoint(ctx *pulumi.Context,
	name string, args *ClusterEndpointArgs, opts ...pulumi.ResourceOption) (*ClusterEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterEndpoint
	err := ctx.RegisterResource("alicloud:polardb/clusterEndpoint:ClusterEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterEndpoint gets an existing ClusterEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterEndpointState, opts ...pulumi.ResourceOption) (*ClusterEndpoint, error) {
	var resource ClusterEndpoint
	err := ctx.ReadResource("alicloud:polardb/clusterEndpoint:ClusterEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterEndpoint resources.
type clusterEndpointState struct {
	// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
	AutoAddNewNodes *string `pulumi:"autoAddNewNodes"`
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// The Id of cluster that can run database.
	DbClusterId *string `pulumi:"dbClusterId"`
	// The name of the endpoint.
	DbEndpointDescription *string `pulumi:"dbEndpointDescription"`
	// The ID of the cluster endpoint.
	DbEndpointId *string `pulumi:"dbEndpointId"`
	// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
	EndpointConfig map[string]string `pulumi:"endpointConfig"`
	// Type of endpoint.
	EndpointType *string `pulumi:"endpointType"`
	// The network type of the endpoint address.
	NetType *string `pulumi:"netType"`
	// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
	Nodes []string `pulumi:"nodes"`
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port *string `pulumi:"port"`
	// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
	ReadWriteMode *string `pulumi:"readWriteMode"`
	// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
	// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
	// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
	SslAutoRotate *string `pulumi:"sslAutoRotate"`
	// The specifies SSL certificate download link.
	SslCertificateUrl *string `pulumi:"sslCertificateUrl"`
	// The SSL connection string.
	SslConnectionString *string `pulumi:"sslConnectionString"`
	// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
	SslEnabled *string `pulumi:"sslEnabled"`
	// The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SslExpireTime *string `pulumi:"sslExpireTime"`
}

type ClusterEndpointState struct {
	// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
	AutoAddNewNodes pulumi.StringPtrInput
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringPtrInput
	// The name of the endpoint.
	DbEndpointDescription pulumi.StringPtrInput
	// The ID of the cluster endpoint.
	DbEndpointId pulumi.StringPtrInput
	// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
	EndpointConfig pulumi.StringMapInput
	// Type of endpoint.
	EndpointType pulumi.StringPtrInput
	// The network type of the endpoint address.
	NetType pulumi.StringPtrInput
	// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
	Nodes pulumi.StringArrayInput
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port pulumi.StringPtrInput
	// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
	ReadWriteMode pulumi.StringPtrInput
	// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
	// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
	// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
	SslAutoRotate pulumi.StringPtrInput
	// The specifies SSL certificate download link.
	SslCertificateUrl pulumi.StringPtrInput
	// The SSL connection string.
	SslConnectionString pulumi.StringPtrInput
	// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
	SslEnabled pulumi.StringPtrInput
	// The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	SslExpireTime pulumi.StringPtrInput
}

func (ClusterEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterEndpointState)(nil)).Elem()
}

type clusterEndpointArgs struct {
	// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
	AutoAddNewNodes *string `pulumi:"autoAddNewNodes"`
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// The Id of cluster that can run database.
	DbClusterId string `pulumi:"dbClusterId"`
	// The name of the endpoint.
	DbEndpointDescription *string `pulumi:"dbEndpointDescription"`
	// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
	EndpointConfig map[string]string `pulumi:"endpointConfig"`
	// The network type of the endpoint address.
	NetType *string `pulumi:"netType"`
	// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
	Nodes []string `pulumi:"nodes"`
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port *string `pulumi:"port"`
	// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
	ReadWriteMode *string `pulumi:"readWriteMode"`
	// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
	// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
	// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
	SslAutoRotate *string `pulumi:"sslAutoRotate"`
	// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
	SslEnabled *string `pulumi:"sslEnabled"`
}

// The set of arguments for constructing a ClusterEndpoint resource.
type ClusterEndpointArgs struct {
	// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
	AutoAddNewNodes pulumi.StringPtrInput
	// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
	ConnectionPrefix pulumi.StringPtrInput
	// The Id of cluster that can run database.
	DbClusterId pulumi.StringInput
	// The name of the endpoint.
	DbEndpointDescription pulumi.StringPtrInput
	// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
	EndpointConfig pulumi.StringMapInput
	// The network type of the endpoint address.
	NetType pulumi.StringPtrInput
	// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
	Nodes pulumi.StringArrayInput
	// Port of the specified endpoint. Valid values: 3000 to 5999.
	Port pulumi.StringPtrInput
	// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
	ReadWriteMode pulumi.StringPtrInput
	// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
	// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
	// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
	SslAutoRotate pulumi.StringPtrInput
	// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
	SslEnabled pulumi.StringPtrInput
}

func (ClusterEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterEndpointArgs)(nil)).Elem()
}

type ClusterEndpointInput interface {
	pulumi.Input

	ToClusterEndpointOutput() ClusterEndpointOutput
	ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput
}

func (*ClusterEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (i *ClusterEndpoint) ToClusterEndpointOutput() ClusterEndpointOutput {
	return i.ToClusterEndpointOutputWithContext(context.Background())
}

func (i *ClusterEndpoint) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointOutput)
}

// ClusterEndpointArrayInput is an input type that accepts ClusterEndpointArray and ClusterEndpointArrayOutput values.
// You can construct a concrete instance of `ClusterEndpointArrayInput` via:
//
//	ClusterEndpointArray{ ClusterEndpointArgs{...} }
type ClusterEndpointArrayInput interface {
	pulumi.Input

	ToClusterEndpointArrayOutput() ClusterEndpointArrayOutput
	ToClusterEndpointArrayOutputWithContext(context.Context) ClusterEndpointArrayOutput
}

type ClusterEndpointArray []ClusterEndpointInput

func (ClusterEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterEndpoint)(nil)).Elem()
}

func (i ClusterEndpointArray) ToClusterEndpointArrayOutput() ClusterEndpointArrayOutput {
	return i.ToClusterEndpointArrayOutputWithContext(context.Background())
}

func (i ClusterEndpointArray) ToClusterEndpointArrayOutputWithContext(ctx context.Context) ClusterEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointArrayOutput)
}

// ClusterEndpointMapInput is an input type that accepts ClusterEndpointMap and ClusterEndpointMapOutput values.
// You can construct a concrete instance of `ClusterEndpointMapInput` via:
//
//	ClusterEndpointMap{ "key": ClusterEndpointArgs{...} }
type ClusterEndpointMapInput interface {
	pulumi.Input

	ToClusterEndpointMapOutput() ClusterEndpointMapOutput
	ToClusterEndpointMapOutputWithContext(context.Context) ClusterEndpointMapOutput
}

type ClusterEndpointMap map[string]ClusterEndpointInput

func (ClusterEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterEndpoint)(nil)).Elem()
}

func (i ClusterEndpointMap) ToClusterEndpointMapOutput() ClusterEndpointMapOutput {
	return i.ToClusterEndpointMapOutputWithContext(context.Background())
}

func (i ClusterEndpointMap) ToClusterEndpointMapOutputWithContext(ctx context.Context) ClusterEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointMapOutput)
}

type ClusterEndpointOutput struct{ *pulumi.OutputState }

func (ClusterEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointOutput) ToClusterEndpointOutput() ClusterEndpointOutput {
	return o
}

func (o ClusterEndpointOutput) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return o
}

// Whether the new node automatically joins the default cluster address. Valid values are `Enable`, `Disable`. When creating a new custom endpoint, default to `Disable`.
func (o ClusterEndpointOutput) AutoAddNewNodes() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.AutoAddNewNodes }).(pulumi.StringOutput)
}

// Prefix of the specified endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter.
func (o ClusterEndpointOutput) ConnectionPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.ConnectionPrefix }).(pulumi.StringOutput)
}

// The Id of cluster that can run database.
func (o ClusterEndpointOutput) DbClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.DbClusterId }).(pulumi.StringOutput)
}

// The name of the endpoint.
func (o ClusterEndpointOutput) DbEndpointDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.DbEndpointDescription }).(pulumi.StringPtrOutput)
}

// The ID of the cluster endpoint.
func (o ClusterEndpointOutput) DbEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.DbEndpointId }).(pulumi.StringOutput)
}

// The advanced settings of the endpoint of Apsara PolarDB clusters are in JSON format. Including the settings of consistency level, transaction splitting, connection pool, and offload reads from primary node. For more details, see the [description of EndpointConfig in the Request parameters table for details](https://www.alibabacloud.com/help/doc-detail/116593.htm).
func (o ClusterEndpointOutput) EndpointConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringMapOutput { return v.EndpointConfig }).(pulumi.StringMapOutput)
}

// Type of endpoint.
func (o ClusterEndpointOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// The network type of the endpoint address.
func (o ClusterEndpointOutput) NetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.NetType }).(pulumi.StringPtrOutput)
}

// Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.
func (o ClusterEndpointOutput) Nodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringArrayOutput { return v.Nodes }).(pulumi.StringArrayOutput)
}

// Port of the specified endpoint. Valid values: 3000 to 5999.
func (o ClusterEndpointOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Read or write mode. Valid values are `ReadWrite`, `ReadOnly`. When creating a new custom endpoint, default to `ReadOnly`.
func (o ClusterEndpointOutput) ReadWriteMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.ReadWriteMode }).(pulumi.StringOutput)
}

// Specifies whether automatic rotation of SSL certificates is enabled. Valid values: `Enable`,`Disable`.
// **NOTE:** For a PolarDB for MySQL cluster, this parameter is required, and only one connection string in each endpoint can enable the ssl, for other notes, see [Configure SSL encryption](https://www.alibabacloud.com/help/doc-detail/153182.htm).
// For a PolarDB for PostgreSQL cluster or a PolarDB-O cluster, this parameter is not required, by default, SSL encryption is enabled for all endpoints.
func (o ClusterEndpointOutput) SslAutoRotate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.SslAutoRotate }).(pulumi.StringPtrOutput)
}

// The specifies SSL certificate download link.
func (o ClusterEndpointOutput) SslCertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.SslCertificateUrl }).(pulumi.StringOutput)
}

// The SSL connection string.
func (o ClusterEndpointOutput) SslConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.SslConnectionString }).(pulumi.StringOutput)
}

// Specifies how to modify the SSL encryption status. Valid values: `Disable`, `Enable`, `Update`.
func (o ClusterEndpointOutput) SslEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.SslEnabled }).(pulumi.StringPtrOutput)
}

// The time when the SSL certificate expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
func (o ClusterEndpointOutput) SslExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.SslExpireTime }).(pulumi.StringOutput)
}

type ClusterEndpointArrayOutput struct{ *pulumi.OutputState }

func (ClusterEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointArrayOutput) ToClusterEndpointArrayOutput() ClusterEndpointArrayOutput {
	return o
}

func (o ClusterEndpointArrayOutput) ToClusterEndpointArrayOutputWithContext(ctx context.Context) ClusterEndpointArrayOutput {
	return o
}

func (o ClusterEndpointArrayOutput) Index(i pulumi.IntInput) ClusterEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterEndpoint {
		return vs[0].([]*ClusterEndpoint)[vs[1].(int)]
	}).(ClusterEndpointOutput)
}

type ClusterEndpointMapOutput struct{ *pulumi.OutputState }

func (ClusterEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointMapOutput) ToClusterEndpointMapOutput() ClusterEndpointMapOutput {
	return o
}

func (o ClusterEndpointMapOutput) ToClusterEndpointMapOutputWithContext(ctx context.Context) ClusterEndpointMapOutput {
	return o
}

func (o ClusterEndpointMapOutput) MapIndex(k pulumi.StringInput) ClusterEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterEndpoint {
		return vs[0].(map[string]*ClusterEndpoint)[vs[1].(string)]
	}).(ClusterEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointInput)(nil)).Elem(), &ClusterEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointArrayInput)(nil)).Elem(), ClusterEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointMapInput)(nil)).Elem(), ClusterEndpointMap{})
	pulumi.RegisterOutputType(ClusterEndpointOutput{})
	pulumi.RegisterOutputType(ClusterEndpointArrayOutput{})
	pulumi.RegisterOutputType(ClusterEndpointMapOutput{})
}
