# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 instance_class: pulumi.Input[str],
                 instance_group_count: pulumi.Input[str],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 security_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] instance_class: Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        :param pulumi.Input[str] instance_group_count: The number of groups. Valid values: [2,4,8,16,32]
        :param pulumi.Input[str] description: The name of DB instance. It a string of 2 to 256 characters.
        :param pulumi.Input[str] engine: Database engine: gpdb. System Default value: gpdb.
        :param pulumi.Input[str] engine_version: Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        :param pulumi.Input[str] instance_charge_type: Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_ip_lists: List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: The virtual switch ID to launch DB instances in one VPC.
        """
        pulumi.set(__self__, "instance_class", instance_class)
        pulumi.set(__self__, "instance_group_count", instance_group_count)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if instance_charge_type is not None:
            pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if security_ip_lists is not None:
            pulumi.set(__self__, "security_ip_lists", security_ip_lists)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Input[str]:
        """
        Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="instanceGroupCount")
    def instance_group_count(self) -> pulumi.Input[str]:
        """
        The number of groups. Valid values: [2,4,8,16,32]
        """
        return pulumi.get(self, "instance_group_count")

    @instance_group_count.setter
    def instance_group_count(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_group_count", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The name of DB instance. It a string of 2 to 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        Database engine: gpdb. System Default value: gpdb.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        """
        return pulumi.get(self, "instance_charge_type")

    @instance_charge_type.setter
    def instance_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_charge_type", value)

    @property
    @pulumi.getter(name="securityIpLists")
    def security_ip_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        """
        return pulumi.get(self, "security_ip_lists")

    @security_ip_lists.setter
    def security_ip_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_ip_lists", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The virtual switch ID to launch DB instances in one VPC.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 instance_group_count: Optional[pulumi.Input[str]] = None,
                 security_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] description: The name of DB instance. It a string of 2 to 256 characters.
        :param pulumi.Input[str] engine: Database engine: gpdb. System Default value: gpdb.
        :param pulumi.Input[str] engine_version: Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        :param pulumi.Input[str] instance_charge_type: Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        :param pulumi.Input[str] instance_class: Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        :param pulumi.Input[str] instance_group_count: The number of groups. Valid values: [2,4,8,16,32]
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_ip_lists: List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: The virtual switch ID to launch DB instances in one VPC.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if instance_charge_type is not None:
            pulumi.set(__self__, "instance_charge_type", instance_charge_type)
        if instance_class is not None:
            pulumi.set(__self__, "instance_class", instance_class)
        if instance_group_count is not None:
            pulumi.set(__self__, "instance_group_count", instance_group_count)
        if security_ip_lists is not None:
            pulumi.set(__self__, "security_ip_lists", security_ip_lists)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The name of DB instance. It a string of 2 to 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        Database engine: gpdb. System Default value: gpdb.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        """
        return pulumi.get(self, "instance_charge_type")

    @instance_charge_type.setter
    def instance_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_charge_type", value)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> Optional[pulumi.Input[str]]:
        """
        Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="instanceGroupCount")
    def instance_group_count(self) -> Optional[pulumi.Input[str]]:
        """
        The number of groups. Valid values: [2,4,8,16,32]
        """
        return pulumi.get(self, "instance_group_count")

    @instance_group_count.setter
    def instance_group_count(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_group_count", value)

    @property
    @pulumi.getter(name="securityIpLists")
    def security_ip_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        """
        return pulumi.get(self, "security_ip_lists")

    @security_ip_lists.setter
    def security_ip_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_ip_lists", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The virtual switch ID to launch DB instances in one VPC.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 instance_group_count: Optional[pulumi.Input[str]] = None,
                 security_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a AnalyticDB for PostgreSQL instance resource supports replica set instances only. the AnalyticDB for PostgreSQL provides stable, reliable, and automatic scalable database services.
        You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/35387.htm)

        > **NOTE:**  Available in 1.47.0+

        > **NOTE:**  The following regions don't support create Classic network Gpdb instance.
        [`ap-southeast-2`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`me-east-1`,`ap-northeast-1`,`eu-west-1`,`us-east-1`,`eu-central-1`,`cn-shanghai-finance-1`,`cn-shenzhen-finance-1`,`cn-hangzhou-finance`]

        > **NOTE:**  Create instance or change instance would cost 10~15 minutes. Please make full preparation.

        > **NOTE:**  This resource is used to manage a Reserved Storage Mode instance, and creating a new reserved storage mode instance is no longer supported since v1.127.0.
        You can still use this resource to manage the instance which has been already created, but can not create a new one.

        ## Example Usage
        ### Create a Gpdb instance

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_resource_creation="Gpdb")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            zone_id=default_zones.zones[0].id,
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            vswitch_name="vpc-123456")
        example = alicloud.gpdb.Instance("example",
            description="tf-gpdb-test",
            engine="gpdb",
            engine_version="4.3",
            instance_class="gpdb.group.segsdx2",
            instance_group_count="2",
            vswitch_id=default_switch.id,
            security_ip_lists=[
                "10.168.1.12",
                "100.69.7.112",
            ])
        ```

        ## Import

        AnalyticDB for PostgreSQL can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:gpdb/instance:Instance example gp-bp1291daeda44194
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The name of DB instance. It a string of 2 to 256 characters.
        :param pulumi.Input[str] engine: Database engine: gpdb. System Default value: gpdb.
        :param pulumi.Input[str] engine_version: Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        :param pulumi.Input[str] instance_charge_type: Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        :param pulumi.Input[str] instance_class: Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        :param pulumi.Input[str] instance_group_count: The number of groups. Valid values: [2,4,8,16,32]
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_ip_lists: List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: The virtual switch ID to launch DB instances in one VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a AnalyticDB for PostgreSQL instance resource supports replica set instances only. the AnalyticDB for PostgreSQL provides stable, reliable, and automatic scalable database services.
        You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/35387.htm)

        > **NOTE:**  Available in 1.47.0+

        > **NOTE:**  The following regions don't support create Classic network Gpdb instance.
        [`ap-southeast-2`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`me-east-1`,`ap-northeast-1`,`eu-west-1`,`us-east-1`,`eu-central-1`,`cn-shanghai-finance-1`,`cn-shenzhen-finance-1`,`cn-hangzhou-finance`]

        > **NOTE:**  Create instance or change instance would cost 10~15 minutes. Please make full preparation.

        > **NOTE:**  This resource is used to manage a Reserved Storage Mode instance, and creating a new reserved storage mode instance is no longer supported since v1.127.0.
        You can still use this resource to manage the instance which has been already created, but can not create a new one.

        ## Example Usage
        ### Create a Gpdb instance

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_resource_creation="Gpdb")
        default_network = alicloud.vpc.Network("defaultNetwork", cidr_block="172.16.0.0/16")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            zone_id=default_zones.zones[0].id,
            vpc_id=default_network.id,
            cidr_block="172.16.0.0/24",
            vswitch_name="vpc-123456")
        example = alicloud.gpdb.Instance("example",
            description="tf-gpdb-test",
            engine="gpdb",
            engine_version="4.3",
            instance_class="gpdb.group.segsdx2",
            instance_group_count="2",
            vswitch_id=default_switch.id,
            security_ip_lists=[
                "10.168.1.12",
                "100.69.7.112",
            ])
        ```

        ## Import

        AnalyticDB for PostgreSQL can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:gpdb/instance:Instance example gp-bp1291daeda44194
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 instance_charge_type: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 instance_group_count: Optional[pulumi.Input[str]] = None,
                 security_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["description"] = description
            __props__.__dict__["engine"] = engine
            __props__.__dict__["engine_version"] = engine_version
            __props__.__dict__["instance_charge_type"] = instance_charge_type
            if instance_class is None and not opts.urn:
                raise TypeError("Missing required property 'instance_class'")
            __props__.__dict__["instance_class"] = instance_class
            if instance_group_count is None and not opts.urn:
                raise TypeError("Missing required property 'instance_group_count'")
            __props__.__dict__["instance_group_count"] = instance_group_count
            __props__.__dict__["security_ip_lists"] = security_ip_lists
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vswitch_id"] = vswitch_id
        super(Instance, __self__).__init__(
            'alicloud:gpdb/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            instance_charge_type: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            instance_group_count: Optional[pulumi.Input[str]] = None,
            security_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The name of DB instance. It a string of 2 to 256 characters.
        :param pulumi.Input[str] engine: Database engine: gpdb. System Default value: gpdb.
        :param pulumi.Input[str] engine_version: Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        :param pulumi.Input[str] instance_charge_type: Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        :param pulumi.Input[str] instance_class: Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        :param pulumi.Input[str] instance_group_count: The number of groups. Valid values: [2,4,8,16,32]
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_ip_lists: List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: The virtual switch ID to launch DB instances in one VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["description"] = description
        __props__.__dict__["engine"] = engine
        __props__.__dict__["engine_version"] = engine_version
        __props__.__dict__["instance_charge_type"] = instance_charge_type
        __props__.__dict__["instance_class"] = instance_class
        __props__.__dict__["instance_group_count"] = instance_group_count
        __props__.__dict__["security_ip_lists"] = security_ip_lists
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vswitch_id"] = vswitch_id
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The name of DB instance. It a string of 2 to 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        Database engine: gpdb. System Default value: gpdb.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) `EngineVersion`.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="instanceChargeType")
    def instance_charge_type(self) -> pulumi.Output[str]:
        """
        Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
        """
        return pulumi.get(self, "instance_charge_type")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Output[str]:
        """
        Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="instanceGroupCount")
    def instance_group_count(self) -> pulumi.Output[str]:
        """
        The number of groups. Valid values: [2,4,8,16,32]
        """
        return pulumi.get(self, "instance_group_count")

    @property
    @pulumi.getter(name="securityIpLists")
    def security_ip_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
        """
        return pulumi.get(self, "security_ip_lists")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        The virtual switch ID to launch DB instances in one VPC.
        """
        return pulumi.get(self, "vswitch_id")

