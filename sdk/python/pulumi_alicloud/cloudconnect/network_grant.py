# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkGrantArgs', 'NetworkGrant']

@pulumi.input_type
class NetworkGrantArgs:
    def __init__(__self__, *,
                 ccn_id: pulumi.Input[str],
                 cen_id: pulumi.Input[str],
                 cen_uid: pulumi.Input[str]):
        """
        The set of arguments for constructing a NetworkGrant resource.
        :param pulumi.Input[str] ccn_id: The ID of the CCN instance.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] cen_uid: The ID of the account to which the CEN instance belongs.
        """
        pulumi.set(__self__, "ccn_id", ccn_id)
        pulumi.set(__self__, "cen_id", cen_id)
        pulumi.set(__self__, "cen_uid", cen_uid)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> pulumi.Input[str]:
        """
        The ID of the CCN instance.
        """
        return pulumi.get(self, "ccn_id")

    @ccn_id.setter
    def ccn_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ccn_id", value)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Input[str]:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="cenUid")
    def cen_uid(self) -> pulumi.Input[str]:
        """
        The ID of the account to which the CEN instance belongs.
        """
        return pulumi.get(self, "cen_uid")

    @cen_uid.setter
    def cen_uid(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_uid", value)


@pulumi.input_type
class _NetworkGrantState:
    def __init__(__self__, *,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_uid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkGrant resources.
        :param pulumi.Input[str] ccn_id: The ID of the CCN instance.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] cen_uid: The ID of the account to which the CEN instance belongs.
        """
        if ccn_id is not None:
            pulumi.set(__self__, "ccn_id", ccn_id)
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)
        if cen_uid is not None:
            pulumi.set(__self__, "cen_uid", cen_uid)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the CCN instance.
        """
        return pulumi.get(self, "ccn_id")

    @ccn_id.setter
    def ccn_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ccn_id", value)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="cenUid")
    def cen_uid(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the account to which the CEN instance belongs.
        """
        return pulumi.get(self, "cen_uid")

    @cen_uid.setter
    def cen_uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_uid", value)


class NetworkGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_uid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Cloud Connect Network Grant resource. If the CEN instance to be attached belongs to another account, authorization by the CEN instance is required.

        For information about Cloud Connect Network Grant and how to use it, see [What is Cloud Connect Network Grant](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/grantinstancetocbn).

        > **NOTE:** Available since v1.63.0.

        > **NOTE:** Only the following regions support create Cloud Connect Network Grant. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        cen_uid = config.get_float("cenUid")
        if cen_uid is None:
            cen_uid = 123456789
        default = alicloud.Provider("default", region="cn-shanghai")
        # Method 1: Use assume_role to operate resources in the target cen account, detail see https://registry.terraform.io/providers/aliyun/alicloud/latest/docs#assume-role
        cen_account = alicloud.Provider("cenAccount",
            region="cn-hangzhou",
            assume_role=alicloud.ProviderAssumeRoleArgs(
                role_arn=f"acs:ram::{cen_uid}:role/terraform-example-assume-role",
            ))
        # Method 2: Use the target cen account's access_key, secret_key
        # provider "alicloud" {
        #   region     = "cn-hangzhou"
        #   access_key = "access_key"
        #   secret_key = "secret_key"
        #   alias      = "cen_account"
        # }
        default_network = alicloud.cloudconnect.Network("defaultNetwork",
            description=name,
            cidr_block="192.168.0.0/24",
            is_default=True,
            opts=pulumi.ResourceOptions(provider=alicloud["default"]))
        cen = alicloud.cen.Instance("cen", cen_instance_name=name,
        opts=pulumi.ResourceOptions(provider=alicloud["cen_account"]))
        default_network_grant = alicloud.cloudconnect.NetworkGrant("defaultNetworkGrant",
            ccn_id=default_network.id,
            cen_id=cen.id,
            cen_uid=cen_uid,
            opts=pulumi.ResourceOptions(provider=alicloud["default"]))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The Cloud Connect Network Grant can be imported using the instance_id, e.g.

        ```sh
        $ pulumi import alicloud:cloudconnect/networkGrant:NetworkGrant example ccn-abc123456:cen-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ccn_id: The ID of the CCN instance.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] cen_uid: The ID of the account to which the CEN instance belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Cloud Connect Network Grant resource. If the CEN instance to be attached belongs to another account, authorization by the CEN instance is required.

        For information about Cloud Connect Network Grant and how to use it, see [What is Cloud Connect Network Grant](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/grantinstancetocbn).

        > **NOTE:** Available since v1.63.0.

        > **NOTE:** Only the following regions support create Cloud Connect Network Grant. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "tf-example"
        cen_uid = config.get_float("cenUid")
        if cen_uid is None:
            cen_uid = 123456789
        default = alicloud.Provider("default", region="cn-shanghai")
        # Method 1: Use assume_role to operate resources in the target cen account, detail see https://registry.terraform.io/providers/aliyun/alicloud/latest/docs#assume-role
        cen_account = alicloud.Provider("cenAccount",
            region="cn-hangzhou",
            assume_role=alicloud.ProviderAssumeRoleArgs(
                role_arn=f"acs:ram::{cen_uid}:role/terraform-example-assume-role",
            ))
        # Method 2: Use the target cen account's access_key, secret_key
        # provider "alicloud" {
        #   region     = "cn-hangzhou"
        #   access_key = "access_key"
        #   secret_key = "secret_key"
        #   alias      = "cen_account"
        # }
        default_network = alicloud.cloudconnect.Network("defaultNetwork",
            description=name,
            cidr_block="192.168.0.0/24",
            is_default=True,
            opts=pulumi.ResourceOptions(provider=alicloud["default"]))
        cen = alicloud.cen.Instance("cen", cen_instance_name=name,
        opts=pulumi.ResourceOptions(provider=alicloud["cen_account"]))
        default_network_grant = alicloud.cloudconnect.NetworkGrant("defaultNetworkGrant",
            ccn_id=default_network.id,
            cen_id=cen.id,
            cen_uid=cen_uid,
            opts=pulumi.ResourceOptions(provider=alicloud["default"]))
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The Cloud Connect Network Grant can be imported using the instance_id, e.g.

        ```sh
        $ pulumi import alicloud:cloudconnect/networkGrant:NetworkGrant example ccn-abc123456:cen-abc123456
        ```

        :param str resource_name: The name of the resource.
        :param NetworkGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_uid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkGrantArgs.__new__(NetworkGrantArgs)

            if ccn_id is None and not opts.urn:
                raise TypeError("Missing required property 'ccn_id'")
            __props__.__dict__["ccn_id"] = ccn_id
            if cen_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_id'")
            __props__.__dict__["cen_id"] = cen_id
            if cen_uid is None and not opts.urn:
                raise TypeError("Missing required property 'cen_uid'")
            __props__.__dict__["cen_uid"] = cen_uid
        super(NetworkGrant, __self__).__init__(
            'alicloud:cloudconnect/networkGrant:NetworkGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ccn_id: Optional[pulumi.Input[str]] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            cen_uid: Optional[pulumi.Input[str]] = None) -> 'NetworkGrant':
        """
        Get an existing NetworkGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ccn_id: The ID of the CCN instance.
        :param pulumi.Input[str] cen_id: The ID of the CEN instance.
        :param pulumi.Input[str] cen_uid: The ID of the account to which the CEN instance belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkGrantState.__new__(_NetworkGrantState)

        __props__.__dict__["ccn_id"] = ccn_id
        __props__.__dict__["cen_id"] = cen_id
        __props__.__dict__["cen_uid"] = cen_uid
        return NetworkGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> pulumi.Output[str]:
        """
        The ID of the CCN instance.
        """
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The ID of the CEN instance.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="cenUid")
    def cen_uid(self) -> pulumi.Output[str]:
        """
        The ID of the account to which the CEN instance belongs.
        """
        return pulumi.get(self, "cen_uid")

