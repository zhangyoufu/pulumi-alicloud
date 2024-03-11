# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcEndpointServiceConnectionArgs', 'VpcEndpointServiceConnection']

@pulumi.input_type
class VpcEndpointServiceConnectionArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a VpcEndpointServiceConnection resource.
        :param pulumi.Input[str] endpoint_id: The endpoint ID.
        :param pulumi.Input[str] service_id: The endpoint service ID.
        :param pulumi.Input[int] bandwidth: The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        :param pulumi.Input[bool] dry_run: Specifies whether to perform only a dry run, without performing the actual request. Valid values:
               - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
               - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "service_id", service_id)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        """
        The endpoint ID.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        The endpoint service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)


@pulumi.input_type
class _VpcEndpointServiceConnectionState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointServiceConnection resources.
        :param pulumi.Input[int] bandwidth: The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        :param pulumi.Input[bool] dry_run: Specifies whether to perform only a dry run, without performing the actual request. Valid values:
               - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
               - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        :param pulumi.Input[str] endpoint_id: The endpoint ID.
        :param pulumi.Input[str] service_id: The endpoint service ID.
        :param pulumi.Input[str] status: The state of the endpoint connection.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint ID.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the endpoint connection.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class VpcEndpointServiceConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Private Link Vpc Endpoint Connection resource. vpc endpoint connection.

        For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-enablevpcendpointzoneconnection).

        > **NOTE:** Available since v1.110.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        example_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        example_vpc_endpoint_service = alicloud.privatelink.VpcEndpointService("exampleVpcEndpointService",
            service_description=name,
            connect_bandwidth=103,
            auto_accept_connection=False)
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name=name,
            cidr_block="10.0.0.0/8")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vswitch_name=name,
            cidr_block="10.1.0.0/16",
            vpc_id=example_network.id,
            zone_id=example_zones.zones[0].id)
        example_security_group = alicloud.ecs.SecurityGroup("exampleSecurityGroup", vpc_id=example_network.id)
        example_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("exampleApplicationLoadBalancer",
            load_balancer_name=name,
            vswitch_id=example_switch.id,
            load_balancer_spec="slb.s2.small",
            address_type="intranet")
        example_vpc_endpoint_service_resource = alicloud.privatelink.VpcEndpointServiceResource("exampleVpcEndpointServiceResource",
            service_id=example_vpc_endpoint_service.id,
            resource_id=example_application_load_balancer.id,
            resource_type="slb")
        example_vpc_endpoint = alicloud.privatelink.VpcEndpoint("exampleVpcEndpoint",
            service_id=example_vpc_endpoint_service_resource.service_id,
            security_group_ids=[example_security_group.id],
            vpc_id=example_network.id,
            vpc_endpoint_name=name)
        example_vpc_endpoint_service_connection = alicloud.privatelink.VpcEndpointServiceConnection("exampleVpcEndpointServiceConnection",
            endpoint_id=example_vpc_endpoint.id,
            service_id=example_vpc_endpoint.service_id,
            bandwidth=1024)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Private Link Vpc Endpoint Connection can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example <service_id>:<endpoint_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        :param pulumi.Input[bool] dry_run: Specifies whether to perform only a dry run, without performing the actual request. Valid values:
               - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
               - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        :param pulumi.Input[str] endpoint_id: The endpoint ID.
        :param pulumi.Input[str] service_id: The endpoint service ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointServiceConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Private Link Vpc Endpoint Connection resource. vpc endpoint connection.

        For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-enablevpcendpointzoneconnection).

        > **NOTE:** Available since v1.110.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf_example"
        example_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        example_vpc_endpoint_service = alicloud.privatelink.VpcEndpointService("exampleVpcEndpointService",
            service_description=name,
            connect_bandwidth=103,
            auto_accept_connection=False)
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name=name,
            cidr_block="10.0.0.0/8")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vswitch_name=name,
            cidr_block="10.1.0.0/16",
            vpc_id=example_network.id,
            zone_id=example_zones.zones[0].id)
        example_security_group = alicloud.ecs.SecurityGroup("exampleSecurityGroup", vpc_id=example_network.id)
        example_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("exampleApplicationLoadBalancer",
            load_balancer_name=name,
            vswitch_id=example_switch.id,
            load_balancer_spec="slb.s2.small",
            address_type="intranet")
        example_vpc_endpoint_service_resource = alicloud.privatelink.VpcEndpointServiceResource("exampleVpcEndpointServiceResource",
            service_id=example_vpc_endpoint_service.id,
            resource_id=example_application_load_balancer.id,
            resource_type="slb")
        example_vpc_endpoint = alicloud.privatelink.VpcEndpoint("exampleVpcEndpoint",
            service_id=example_vpc_endpoint_service_resource.service_id,
            security_group_ids=[example_security_group.id],
            vpc_id=example_network.id,
            vpc_endpoint_name=name)
        example_vpc_endpoint_service_connection = alicloud.privatelink.VpcEndpointServiceConnection("exampleVpcEndpointServiceConnection",
            endpoint_id=example_vpc_endpoint.id,
            service_id=example_vpc_endpoint.service_id,
            bandwidth=1024)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Private Link Vpc Endpoint Connection can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example <service_id>:<endpoint_id>
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointServiceConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointServiceConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcEndpointServiceConnectionArgs.__new__(VpcEndpointServiceConnectionArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["dry_run"] = dry_run
            if endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__.__dict__["endpoint_id"] = endpoint_id
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["status"] = None
        super(VpcEndpointServiceConnection, __self__).__init__(
            'alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            endpoint_id: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointServiceConnection':
        """
        Get an existing VpcEndpointServiceConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        :param pulumi.Input[bool] dry_run: Specifies whether to perform only a dry run, without performing the actual request. Valid values:
               - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
               - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        :param pulumi.Input[str] endpoint_id: The endpoint ID.
        :param pulumi.Input[str] service_id: The endpoint service ID.
        :param pulumi.Input[str] status: The state of the endpoint connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointServiceConnectionState.__new__(_VpcEndpointServiceConnectionState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["endpoint_id"] = endpoint_id
        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["status"] = status
        return VpcEndpointServiceConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        The bandwidth of the endpoint connection. Valid values: 100 to 10240. Unit: Mbit/s.Note: The bandwidth of an endpoint connection is in the range of 100 to 10,240 Mbit/s. The default bandwidth is 1,024 Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is 1,024 to 10,240 Mbit/s.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to perform only a dry run, without performing the actual request. Valid values:
        - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
        - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        """
        The endpoint ID.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        The endpoint service ID.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The state of the endpoint connection.
        """
        return pulumi.get(self, "status")

