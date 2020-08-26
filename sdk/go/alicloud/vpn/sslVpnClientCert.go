// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SslVpnClientCert struct {
	pulumi.CustomResourceState

	// The client ca cert.
	CaCert pulumi.StringOutput `pulumi:"caCert"`
	// The client cert.
	ClientCert pulumi.StringOutput `pulumi:"clientCert"`
	// The vpn client config.
	ClientConfig pulumi.StringOutput `pulumi:"clientConfig"`
	// The client key.
	ClientKey pulumi.StringOutput `pulumi:"clientKey"`
	// The name of the client certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the SSL-VPN server.
	SslVpnServerId pulumi.StringOutput `pulumi:"sslVpnServerId"`
	// The status of the client certificate.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSslVpnClientCert registers a new resource with the given unique name, arguments, and options.
func NewSslVpnClientCert(ctx *pulumi.Context,
	name string, args *SslVpnClientCertArgs, opts ...pulumi.ResourceOption) (*SslVpnClientCert, error) {
	if args == nil || args.SslVpnServerId == nil {
		return nil, errors.New("missing required argument 'SslVpnServerId'")
	}
	if args == nil {
		args = &SslVpnClientCertArgs{}
	}
	var resource SslVpnClientCert
	err := ctx.RegisterResource("alicloud:vpn/sslVpnClientCert:SslVpnClientCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslVpnClientCert gets an existing SslVpnClientCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslVpnClientCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslVpnClientCertState, opts ...pulumi.ResourceOption) (*SslVpnClientCert, error) {
	var resource SslVpnClientCert
	err := ctx.ReadResource("alicloud:vpn/sslVpnClientCert:SslVpnClientCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslVpnClientCert resources.
type sslVpnClientCertState struct {
	// The client ca cert.
	CaCert *string `pulumi:"caCert"`
	// The client cert.
	ClientCert *string `pulumi:"clientCert"`
	// The vpn client config.
	ClientConfig *string `pulumi:"clientConfig"`
	// The client key.
	ClientKey *string `pulumi:"clientKey"`
	// The name of the client certificate.
	Name *string `pulumi:"name"`
	// The ID of the SSL-VPN server.
	SslVpnServerId *string `pulumi:"sslVpnServerId"`
	// The status of the client certificate.
	Status *string `pulumi:"status"`
}

type SslVpnClientCertState struct {
	// The client ca cert.
	CaCert pulumi.StringPtrInput
	// The client cert.
	ClientCert pulumi.StringPtrInput
	// The vpn client config.
	ClientConfig pulumi.StringPtrInput
	// The client key.
	ClientKey pulumi.StringPtrInput
	// The name of the client certificate.
	Name pulumi.StringPtrInput
	// The ID of the SSL-VPN server.
	SslVpnServerId pulumi.StringPtrInput
	// The status of the client certificate.
	Status pulumi.StringPtrInput
}

func (SslVpnClientCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslVpnClientCertState)(nil)).Elem()
}

type sslVpnClientCertArgs struct {
	// The name of the client certificate.
	Name *string `pulumi:"name"`
	// The ID of the SSL-VPN server.
	SslVpnServerId string `pulumi:"sslVpnServerId"`
}

// The set of arguments for constructing a SslVpnClientCert resource.
type SslVpnClientCertArgs struct {
	// The name of the client certificate.
	Name pulumi.StringPtrInput
	// The ID of the SSL-VPN server.
	SslVpnServerId pulumi.StringInput
}

func (SslVpnClientCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslVpnClientCertArgs)(nil)).Elem()
}
