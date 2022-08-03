# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DBAuditInstanceArgs', 'DBAuditInstance']

@pulumi.input_type
class DBAuditInstanceArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 period: pulumi.Input[int],
                 plan_code: pulumi.Input[str],
                 vswitch_id: pulumi.Input[str],
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a DBAuditInstance resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] plan_code: Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        :param pulumi.Input[str] vswitch_id: vSwtich ID configured to audit
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "period", period)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "vswitch_id", vswitch_id)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Description of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def period(self) -> pulumi.Input[int]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: pulumi.Input[int]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> pulumi.Input[str]:
        """
        Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        """
        return pulumi.get(self, "plan_code")

    @plan_code.setter
    def plan_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan_code", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Input[str]:
        """
        vSwtich ID configured to audit
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vswitch_id", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

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


@pulumi.input_type
class _DBAuditInstanceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 plan_code: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DBAuditInstance resources.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] plan_code: Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: vSwtich ID configured to audit
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if plan_code is not None:
            pulumi.set(__self__, "plan_code", plan_code)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> Optional[pulumi.Input[str]]:
        """
        Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        """
        return pulumi.get(self, "plan_code")

    @plan_code.setter
    def plan_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_code", value)

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

    @resource_group_id.setter
    def resource_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_id", value)

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
        vSwtich ID configured to audit
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class DBAuditInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 plan_code: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Yundun_dbaudit instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:yundun/dBAuditInstance:DBAuditInstance example dbaudit-exampe123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] plan_code: Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: vSwtich ID configured to audit
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DBAuditInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Yundun_dbaudit instance can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:yundun/dBAuditInstance:DBAuditInstance example dbaudit-exampe123456
        ```

        :param str resource_name: The name of the resource.
        :param DBAuditInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DBAuditInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 plan_code: Optional[pulumi.Input[str]] = None,
                 resource_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DBAuditInstanceArgs.__new__(DBAuditInstanceArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if period is None and not opts.urn:
                raise TypeError("Missing required property 'period'")
            __props__.__dict__["period"] = period
            if plan_code is None and not opts.urn:
                raise TypeError("Missing required property 'plan_code'")
            __props__.__dict__["plan_code"] = plan_code
            __props__.__dict__["resource_group_id"] = resource_group_id
            __props__.__dict__["tags"] = tags
            if vswitch_id is None and not opts.urn:
                raise TypeError("Missing required property 'vswitch_id'")
            __props__.__dict__["vswitch_id"] = vswitch_id
        super(DBAuditInstance, __self__).__init__(
            'alicloud:yundun/dBAuditInstance:DBAuditInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            plan_code: Optional[pulumi.Input[str]] = None,
            resource_group_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'DBAuditInstance':
        """
        Get an existing DBAuditInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance. This name can have a string of 1 to 63 characters.
        :param pulumi.Input[str] plan_code: Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        :param pulumi.Input[str] resource_group_id: The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        :param pulumi.Input[Mapping[str, Any]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vswitch_id: vSwtich ID configured to audit
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DBAuditInstanceState.__new__(_DBAuditInstanceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["period"] = period
        __props__.__dict__["plan_code"] = plan_code
        __props__.__dict__["resource_group_id"] = resource_group_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vswitch_id"] = vswitch_id
        return DBAuditInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the instance. This name can have a string of 1 to 63 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[int]:
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> pulumi.Output[str]:
        """
        Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group.
        """
        return pulumi.get(self, "resource_group_id")

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
        vSwtich ID configured to audit
        """
        return pulumi.get(self, "vswitch_id")

