// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **DEPRECATED:** This resource manages applications in swarm cluster only, which is being deprecated and will be replaced by Kubernetes cluster.
//
// This resource use an orchestration template to define and deploy a multi-container application. An application is created by using an orchestration template.
// Each application can contain one or more services.
//
// > **NOTE:** Application orchestration template must be a valid Docker Compose YAML template.
//
// > **NOTE:** At present, this resource only support swarm cluster.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cs_application.html.markdown.
type Application struct {
	pulumi.CustomResourceState

	// Wherther to use "Blue Green" method when release a new version. Default to false.
	BlueGreen pulumi.BoolPtrOutput `pulumi:"blueGreen"`
	// Whether to confirm a "Blue Green" application. Default to false. It will be ignored when `blueGreen` is false.
	BlueGreenConfirm pulumi.BoolPtrOutput `pulumi:"blueGreenConfirm"`
	// The swarm cluster's name.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The application default domain and it can be used to configure routing service.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// The description of application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A key/value map used to replace the variable parameter in the Compose template.
	Environment pulumi.MapOutput `pulumi:"environment"`
	// Whether to use latest docker image while each updating application. Default to false.
	LatestImage pulumi.BoolPtrOutput `pulumi:"latestImage"`
	// The application name. It should be 1-64 characters long, and can contain numbers, English letters and hyphens, but cannot start with hyphens.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of services in the application. It contains several attributes to `Block Nodes`.
	Services ApplicationServiceArrayOutput `pulumi:"services"`
	// The application deployment template and it must be [Docker Compose format](https://docs.docker.com/compose/).
	Template pulumi.StringOutput `pulumi:"template"`
	// The application deploying version. Each updating, it must be different with current. Default to "1.0"
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.Template == nil {
		return nil, errors.New("missing required argument 'Template'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("alicloud:cs/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("alicloud:cs/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Wherther to use "Blue Green" method when release a new version. Default to false.
	BlueGreen *bool `pulumi:"blueGreen"`
	// Whether to confirm a "Blue Green" application. Default to false. It will be ignored when `blueGreen` is false.
	BlueGreenConfirm *bool `pulumi:"blueGreenConfirm"`
	// The swarm cluster's name.
	ClusterName *string `pulumi:"clusterName"`
	// The application default domain and it can be used to configure routing service.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// The description of application.
	Description *string `pulumi:"description"`
	// A key/value map used to replace the variable parameter in the Compose template.
	Environment map[string]interface{} `pulumi:"environment"`
	// Whether to use latest docker image while each updating application. Default to false.
	LatestImage *bool `pulumi:"latestImage"`
	// The application name. It should be 1-64 characters long, and can contain numbers, English letters and hyphens, but cannot start with hyphens.
	Name *string `pulumi:"name"`
	// List of services in the application. It contains several attributes to `Block Nodes`.
	Services []ApplicationService `pulumi:"services"`
	// The application deployment template and it must be [Docker Compose format](https://docs.docker.com/compose/).
	Template *string `pulumi:"template"`
	// The application deploying version. Each updating, it must be different with current. Default to "1.0"
	Version *string `pulumi:"version"`
}

type ApplicationState struct {
	// Wherther to use "Blue Green" method when release a new version. Default to false.
	BlueGreen pulumi.BoolPtrInput
	// Whether to confirm a "Blue Green" application. Default to false. It will be ignored when `blueGreen` is false.
	BlueGreenConfirm pulumi.BoolPtrInput
	// The swarm cluster's name.
	ClusterName pulumi.StringPtrInput
	// The application default domain and it can be used to configure routing service.
	DefaultDomain pulumi.StringPtrInput
	// The description of application.
	Description pulumi.StringPtrInput
	// A key/value map used to replace the variable parameter in the Compose template.
	Environment pulumi.MapInput
	// Whether to use latest docker image while each updating application. Default to false.
	LatestImage pulumi.BoolPtrInput
	// The application name. It should be 1-64 characters long, and can contain numbers, English letters and hyphens, but cannot start with hyphens.
	Name pulumi.StringPtrInput
	// List of services in the application. It contains several attributes to `Block Nodes`.
	Services ApplicationServiceArrayInput
	// The application deployment template and it must be [Docker Compose format](https://docs.docker.com/compose/).
	Template pulumi.StringPtrInput
	// The application deploying version. Each updating, it must be different with current. Default to "1.0"
	Version pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Wherther to use "Blue Green" method when release a new version. Default to false.
	BlueGreen *bool `pulumi:"blueGreen"`
	// Whether to confirm a "Blue Green" application. Default to false. It will be ignored when `blueGreen` is false.
	BlueGreenConfirm *bool `pulumi:"blueGreenConfirm"`
	// The swarm cluster's name.
	ClusterName string `pulumi:"clusterName"`
	// The description of application.
	Description *string `pulumi:"description"`
	// A key/value map used to replace the variable parameter in the Compose template.
	Environment map[string]interface{} `pulumi:"environment"`
	// Whether to use latest docker image while each updating application. Default to false.
	LatestImage *bool `pulumi:"latestImage"`
	// The application name. It should be 1-64 characters long, and can contain numbers, English letters and hyphens, but cannot start with hyphens.
	Name *string `pulumi:"name"`
	// The application deployment template and it must be [Docker Compose format](https://docs.docker.com/compose/).
	Template string `pulumi:"template"`
	// The application deploying version. Each updating, it must be different with current. Default to "1.0"
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Wherther to use "Blue Green" method when release a new version. Default to false.
	BlueGreen pulumi.BoolPtrInput
	// Whether to confirm a "Blue Green" application. Default to false. It will be ignored when `blueGreen` is false.
	BlueGreenConfirm pulumi.BoolPtrInput
	// The swarm cluster's name.
	ClusterName pulumi.StringInput
	// The description of application.
	Description pulumi.StringPtrInput
	// A key/value map used to replace the variable parameter in the Compose template.
	Environment pulumi.MapInput
	// Whether to use latest docker image while each updating application. Default to false.
	LatestImage pulumi.BoolPtrInput
	// The application name. It should be 1-64 characters long, and can contain numbers, English letters and hyphens, but cannot start with hyphens.
	Name pulumi.StringPtrInput
	// The application deployment template and it must be [Docker Compose format](https://docs.docker.com/compose/).
	Template pulumi.StringInput
	// The application deploying version. Each updating, it must be different with current. Default to "1.0"
	Version pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
