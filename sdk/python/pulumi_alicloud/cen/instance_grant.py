# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceGrantArgs', 'InstanceGrant']

@pulumi.input_type
class InstanceGrantArgs:
    def __init__(__self__, *,
                 cen_id: pulumi.Input[str],
                 cen_owner_id: pulumi.Input[str],
                 child_instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a InstanceGrant resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN.
        :param pulumi.Input[str] cen_owner_id: The owner UID of the  CEN which the child instance granted to.
        :param pulumi.Input[str] child_instance_id: The ID of the child instance to grant.
        """
        pulumi.set(__self__, "cen_id", cen_id)
        pulumi.set(__self__, "cen_owner_id", cen_owner_id)
        pulumi.set(__self__, "child_instance_id", child_instance_id)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Input[str]:
        """
        The ID of the CEN.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="cenOwnerId")
    def cen_owner_id(self) -> pulumi.Input[str]:
        """
        The owner UID of the  CEN which the child instance granted to.
        """
        return pulumi.get(self, "cen_owner_id")

    @cen_owner_id.setter
    def cen_owner_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_owner_id", value)

    @property
    @pulumi.getter(name="childInstanceId")
    def child_instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the child instance to grant.
        """
        return pulumi.get(self, "child_instance_id")

    @child_instance_id.setter
    def child_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "child_instance_id", value)


@pulumi.input_type
class _InstanceGrantState:
    def __init__(__self__, *,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_owner_id: Optional[pulumi.Input[str]] = None,
                 child_instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceGrant resources.
        :param pulumi.Input[str] cen_id: The ID of the CEN.
        :param pulumi.Input[str] cen_owner_id: The owner UID of the  CEN which the child instance granted to.
        :param pulumi.Input[str] child_instance_id: The ID of the child instance to grant.
        """
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)
        if cen_owner_id is not None:
            pulumi.set(__self__, "cen_owner_id", cen_owner_id)
        if child_instance_id is not None:
            pulumi.set(__self__, "child_instance_id", child_instance_id)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the CEN.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="cenOwnerId")
    def cen_owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner UID of the  CEN which the child instance granted to.
        """
        return pulumi.get(self, "cen_owner_id")

    @cen_owner_id.setter
    def cen_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_owner_id", value)

    @property
    @pulumi.getter(name="childInstanceId")
    def child_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the child instance to grant.
        """
        return pulumi.get(self, "child_instance_id")

    @child_instance_id.setter
    def child_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "child_instance_id", value)


class InstanceGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_owner_id: Optional[pulumi.Input[str]] = None,
                 child_instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a CEN child instance grant resource, which allow you to authorize a VPC or VBR to a CEN of a different account.

        For more information about how to use it, see [Attach a network in a different account](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-attachcenchildinstance).

        > **NOTE:** Available since v1.37.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        another_uid = config.get("anotherUid")
        if another_uid is None:
            another_uid = "xxxx"
        your_account = alicloud.get_account()
        child_account = alicloud.get_account()
        default = alicloud.get_regions(current=True)
        example = alicloud.cen.Instance("example",
            cen_instance_name="tf_example",
            description="an example for cen")
        child_account_network = alicloud.vpc.Network("child_account",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        child_account_instance_grant = alicloud.cen.InstanceGrant("child_account",
            cen_id=example.id,
            child_instance_id=child_account_network.id,
            cen_owner_id=your_account.id)
        example_instance_attachment = alicloud.cen.InstanceAttachment("example",
            instance_id=example.id,
            child_instance_id=child_account_instance_grant.child_instance_id,
            child_instance_type="VPC",
            child_instance_region_id=default.regions[0].id,
            child_instance_owner_id=child_account.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        CEN instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/instanceGrant:InstanceGrant example cen-abc123456:vpc-abc123456:uid123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN.
        :param pulumi.Input[str] cen_owner_id: The owner UID of the  CEN which the child instance granted to.
        :param pulumi.Input[str] child_instance_id: The ID of the child instance to grant.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CEN child instance grant resource, which allow you to authorize a VPC or VBR to a CEN of a different account.

        For more information about how to use it, see [Attach a network in a different account](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-attachcenchildinstance).

        > **NOTE:** Available since v1.37.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        another_uid = config.get("anotherUid")
        if another_uid is None:
            another_uid = "xxxx"
        your_account = alicloud.get_account()
        child_account = alicloud.get_account()
        default = alicloud.get_regions(current=True)
        example = alicloud.cen.Instance("example",
            cen_instance_name="tf_example",
            description="an example for cen")
        child_account_network = alicloud.vpc.Network("child_account",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        child_account_instance_grant = alicloud.cen.InstanceGrant("child_account",
            cen_id=example.id,
            child_instance_id=child_account_network.id,
            cen_owner_id=your_account.id)
        example_instance_attachment = alicloud.cen.InstanceAttachment("example",
            instance_id=example.id,
            child_instance_id=child_account_instance_grant.child_instance_id,
            child_instance_type="VPC",
            child_instance_region_id=default.regions[0].id,
            child_instance_owner_id=child_account.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        CEN instance can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:cen/instanceGrant:InstanceGrant example cen-abc123456:vpc-abc123456:uid123456
        ```

        :param str resource_name: The name of the resource.
        :param InstanceGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_owner_id: Optional[pulumi.Input[str]] = None,
                 child_instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceGrantArgs.__new__(InstanceGrantArgs)

            if cen_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_id'")
            __props__.__dict__["cen_id"] = cen_id
            if cen_owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_owner_id'")
            __props__.__dict__["cen_owner_id"] = cen_owner_id
            if child_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'child_instance_id'")
            __props__.__dict__["child_instance_id"] = child_instance_id
        super(InstanceGrant, __self__).__init__(
            'alicloud:cen/instanceGrant:InstanceGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            cen_owner_id: Optional[pulumi.Input[str]] = None,
            child_instance_id: Optional[pulumi.Input[str]] = None) -> 'InstanceGrant':
        """
        Get an existing InstanceGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the CEN.
        :param pulumi.Input[str] cen_owner_id: The owner UID of the  CEN which the child instance granted to.
        :param pulumi.Input[str] child_instance_id: The ID of the child instance to grant.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceGrantState.__new__(_InstanceGrantState)

        __props__.__dict__["cen_id"] = cen_id
        __props__.__dict__["cen_owner_id"] = cen_owner_id
        __props__.__dict__["child_instance_id"] = child_instance_id
        return InstanceGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The ID of the CEN.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="cenOwnerId")
    def cen_owner_id(self) -> pulumi.Output[str]:
        """
        The owner UID of the  CEN which the child instance granted to.
        """
        return pulumi.get(self, "cen_owner_id")

    @property
    @pulumi.getter(name="childInstanceId")
    def child_instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the child instance to grant.
        """
        return pulumi.get(self, "child_instance_id")

