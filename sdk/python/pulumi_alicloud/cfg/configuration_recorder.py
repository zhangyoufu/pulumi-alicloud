# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConfigurationRecorderArgs', 'ConfigurationRecorder']

@pulumi.input_type
class ConfigurationRecorderArgs:
    def __init__(__self__, *,
                 enterprise_edition: Optional[pulumi.Input[bool]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ConfigurationRecorder resource.
        :param pulumi.Input[bool] enterprise_edition: Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
               * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
               * If you use an enterprise account, the `resource_types` does not support updating.
        """
        if enterprise_edition is not None:
            pulumi.set(__self__, "enterprise_edition", enterprise_edition)
        if resource_types is not None:
            pulumi.set(__self__, "resource_types", resource_types)

    @property
    @pulumi.getter(name="enterpriseEdition")
    def enterprise_edition(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        """
        return pulumi.get(self, "enterprise_edition")

    @enterprise_edition.setter
    def enterprise_edition(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enterprise_edition", value)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
        * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
        * If you use an enterprise account, the `resource_types` does not support updating.
        """
        return pulumi.get(self, "resource_types")

    @resource_types.setter
    def resource_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_types", value)


@pulumi.input_type
class _ConfigurationRecorderState:
    def __init__(__self__, *,
                 enterprise_edition: Optional[pulumi.Input[bool]] = None,
                 organization_enable_status: Optional[pulumi.Input[str]] = None,
                 organization_master_id: Optional[pulumi.Input[int]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConfigurationRecorder resources.
        :param pulumi.Input[bool] enterprise_edition: Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        :param pulumi.Input[str] organization_enable_status: Enterprise version configuration audit enabled status. Values: `REGISTRABLE`: Not enabled, `BUILDING`: Building and `REGISTERED`: Enabled.
        :param pulumi.Input[int] organization_master_id: The ID of the Enterprise management account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
               * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
               * If you use an enterprise account, the `resource_types` does not support updating.
        :param pulumi.Input[str] status: Status of resource monitoring. Values: `REGISTRABLE`: Not registered, `BUILDING`: Under construction, `REGISTERED`: Registered and `REBUILDING`: Rebuilding.
        """
        if enterprise_edition is not None:
            pulumi.set(__self__, "enterprise_edition", enterprise_edition)
        if organization_enable_status is not None:
            pulumi.set(__self__, "organization_enable_status", organization_enable_status)
        if organization_master_id is not None:
            pulumi.set(__self__, "organization_master_id", organization_master_id)
        if resource_types is not None:
            pulumi.set(__self__, "resource_types", resource_types)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="enterpriseEdition")
    def enterprise_edition(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        """
        return pulumi.get(self, "enterprise_edition")

    @enterprise_edition.setter
    def enterprise_edition(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enterprise_edition", value)

    @property
    @pulumi.getter(name="organizationEnableStatus")
    def organization_enable_status(self) -> Optional[pulumi.Input[str]]:
        """
        Enterprise version configuration audit enabled status. Values: `REGISTRABLE`: Not enabled, `BUILDING`: Building and `REGISTERED`: Enabled.
        """
        return pulumi.get(self, "organization_enable_status")

    @organization_enable_status.setter
    def organization_enable_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_enable_status", value)

    @property
    @pulumi.getter(name="organizationMasterId")
    def organization_master_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the Enterprise management account.
        """
        return pulumi.get(self, "organization_master_id")

    @organization_master_id.setter
    def organization_master_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "organization_master_id", value)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
        * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
        * If you use an enterprise account, the `resource_types` does not support updating.
        """
        return pulumi.get(self, "resource_types")

    @resource_types.setter
    def resource_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_types", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of resource monitoring. Values: `REGISTRABLE`: Not registered, `BUILDING`: Under construction, `REGISTERED`: Registered and `REBUILDING`: Rebuilding.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ConfigurationRecorder(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enterprise_edition: Optional[pulumi.Input[bool]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Alicloud Config Configuration Recorder resource. Cloud Config is a specialized service for evaluating resources. Cloud Config tracks configuration changes of your resources and evaluates configuration compliance. Cloud Config can help you evaluate numerous resources and maintain the continuous compliance of your cloud infrastructure.
        For information about Alicloud Config Configuration Recorder and how to use it, see [What is Configuration Recorder.](https://www.alibabacloud.com/help/en/cloud-config/latest/startconfigurationrecorder)

        > **NOTE:** Available since v1.99.0.

        > **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-southeast-1`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.cfg.ConfigurationRecorder("example", resource_types=[
            "ACS::ECS::Instance",
            "ACS::ECS::Disk",
        ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Alicloud Config Configuration Recorder can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cfg/configurationRecorder:ConfigurationRecorder example 122378463********
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enterprise_edition: Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
               * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
               * If you use an enterprise account, the `resource_types` does not support updating.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfigurationRecorderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Alicloud Config Configuration Recorder resource. Cloud Config is a specialized service for evaluating resources. Cloud Config tracks configuration changes of your resources and evaluates configuration compliance. Cloud Config can help you evaluate numerous resources and maintain the continuous compliance of your cloud infrastructure.
        For information about Alicloud Config Configuration Recorder and how to use it, see [What is Configuration Recorder.](https://www.alibabacloud.com/help/en/cloud-config/latest/startconfigurationrecorder)

        > **NOTE:** Available since v1.99.0.

        > **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-southeast-1`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.cfg.ConfigurationRecorder("example", resource_types=[
            "ACS::ECS::Instance",
            "ACS::ECS::Disk",
        ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Alicloud Config Configuration Recorder can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cfg/configurationRecorder:ConfigurationRecorder example 122378463********
        ```

        :param str resource_name: The name of the resource.
        :param ConfigurationRecorderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigurationRecorderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enterprise_edition: Optional[pulumi.Input[bool]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigurationRecorderArgs.__new__(ConfigurationRecorderArgs)

            __props__.__dict__["enterprise_edition"] = enterprise_edition
            __props__.__dict__["resource_types"] = resource_types
            __props__.__dict__["organization_enable_status"] = None
            __props__.__dict__["organization_master_id"] = None
            __props__.__dict__["status"] = None
        super(ConfigurationRecorder, __self__).__init__(
            'alicloud:cfg/configurationRecorder:ConfigurationRecorder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enterprise_edition: Optional[pulumi.Input[bool]] = None,
            organization_enable_status: Optional[pulumi.Input[str]] = None,
            organization_master_id: Optional[pulumi.Input[int]] = None,
            resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ConfigurationRecorder':
        """
        Get an existing ConfigurationRecorder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enterprise_edition: Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        :param pulumi.Input[str] organization_enable_status: Enterprise version configuration audit enabled status. Values: `REGISTRABLE`: Not enabled, `BUILDING`: Building and `REGISTERED`: Enabled.
        :param pulumi.Input[int] organization_master_id: The ID of the Enterprise management account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
               * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
               * If you use an enterprise account, the `resource_types` does not support updating.
        :param pulumi.Input[str] status: Status of resource monitoring. Values: `REGISTRABLE`: Not registered, `BUILDING`: Under construction, `REGISTERED`: Registered and `REBUILDING`: Rebuilding.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigurationRecorderState.__new__(_ConfigurationRecorderState)

        __props__.__dict__["enterprise_edition"] = enterprise_edition
        __props__.__dict__["organization_enable_status"] = organization_enable_status
        __props__.__dict__["organization_master_id"] = organization_master_id
        __props__.__dict__["resource_types"] = resource_types
        __props__.__dict__["status"] = status
        return ConfigurationRecorder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enterpriseEdition")
    def enterprise_edition(self) -> pulumi.Output[bool]:
        """
        Whether to use the enterprise version configuration audit. Valid values: `true` and `false`. Default value `false`. For enterprise accounts, We recommend you to use the resource alicloud_config_aggregator.
        """
        return pulumi.get(self, "enterprise_edition")

    @property
    @pulumi.getter(name="organizationEnableStatus")
    def organization_enable_status(self) -> pulumi.Output[str]:
        """
        Enterprise version configuration audit enabled status. Values: `REGISTRABLE`: Not enabled, `BUILDING`: Building and `REGISTERED`: Enabled.
        """
        return pulumi.get(self, "organization_enable_status")

    @property
    @pulumi.getter(name="organizationMasterId")
    def organization_master_id(self) -> pulumi.Output[int]:
        """
        The ID of the Enterprise management account.
        """
        return pulumi.get(self, "organization_master_id")

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of resource types to be monitored. [Resource types that support Cloud Config.](https://www.alibabacloud.com/help/en/doc-detail/127411.htm)
        * If you use an ordinary account, the `resource_types` supports the update operation after the process of creation is completed.
        * If you use an enterprise account, the `resource_types` does not support updating.
        """
        return pulumi.get(self, "resource_types")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of resource monitoring. Values: `REGISTRABLE`: Not registered, `BUILDING`: Under construction, `REGISTERED`: Registered and `REBUILDING`: Rebuilding.
        """
        return pulumi.get(self, "status")

