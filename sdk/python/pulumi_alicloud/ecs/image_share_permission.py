# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ImageSharePermissionArgs', 'ImageSharePermission']

@pulumi.input_type
class ImageSharePermissionArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 image_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ImageSharePermission resource.
        :param pulumi.Input[str] account_id: Alibaba Cloud Account ID. It is used to share images.
        :param pulumi.Input[str] image_id: The source image ID.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "image_id", image_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        """
        Alibaba Cloud Account ID. It is used to share images.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[str]:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_id", value)


@pulumi.input_type
class _ImageSharePermissionState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ImageSharePermission resources.
        :param pulumi.Input[str] account_id: Alibaba Cloud Account ID. It is used to share images.
        :param pulumi.Input[str] image_id: The source image ID.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Alibaba Cloud Account ID. It is used to share images.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)


class ImageSharePermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage image sharing permissions. You can share your custom image to other Alibaba Cloud users. The user can use the shared custom image to create ECS instances or replace the system disk of the instance.

        > **NOTE:** You can only share your own custom images to other Alibaba Cloud users.

        > **NOTE:** Each custom image can be shared with up to 50 Alibaba Cloud accounts. You can submit a ticket to share with more users.

        > **NOTE:** After creating an ECS instance using a shared image, once the custom image owner releases the image sharing relationship or deletes the custom image, the instance cannot initialize the system disk.

        > **NOTE:** Available in 1.68.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_resource_creation="Instance")
        default_instance_types = alicloud.ecs.get_instance_types(instance_type_family="ecs.sn1ne")
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("defaultInstance",
            availability_zone=default_zones.zones[0].id,
            instance_name="terraform-example",
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id,
            instance_type=default_instance_types.ids[0],
            image_id=default_images.ids[0],
            internet_max_bandwidth_out=10)
        default_image = alicloud.ecs.Image("defaultImage",
            instance_id=default_instance.id,
            image_name="terraform-example",
            description="terraform-example")
        config = pulumi.Config()
        account_id = config.get("accountId")
        if account_id is None:
            account_id = "123456789"
        default_image_share_permission = alicloud.ecs.ImageSharePermission("defaultImageSharePermission",
            image_id=default_image.id,
            account_id=account_id)
        ```
        ## Attributes Reference0

         The following attributes are exported:

        * `id` - ID of the image. It formats as `<image_id>:<account_id>`

        ## Import

        image can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/imageSharePermission:ImageSharePermission default m-uf66yg1q:123456789
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: Alibaba Cloud Account ID. It is used to share images.
        :param pulumi.Input[str] image_id: The source image ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageSharePermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage image sharing permissions. You can share your custom image to other Alibaba Cloud users. The user can use the shared custom image to create ECS instances or replace the system disk of the instance.

        > **NOTE:** You can only share your own custom images to other Alibaba Cloud users.

        > **NOTE:** Each custom image can be shared with up to 50 Alibaba Cloud accounts. You can submit a ticket to share with more users.

        > **NOTE:** After creating an ECS instance using a shared image, once the custom image owner releases the image sharing relationship or deletes the custom image, the instance cannot initialize the system disk.

        > **NOTE:** Available in 1.68.0+.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default_zones = alicloud.get_zones(available_resource_creation="Instance")
        default_instance_types = alicloud.ecs.get_instance_types(instance_type_family="ecs.sn1ne")
        default_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        default_network = alicloud.vpc.Network("defaultNetwork",
            vpc_name="terraform-example",
            cidr_block="172.17.3.0/24")
        default_switch = alicloud.vpc.Switch("defaultSwitch",
            vswitch_name="terraform-example",
            cidr_block="172.17.3.0/24",
            vpc_id=default_network.id,
            zone_id=default_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("defaultInstance",
            availability_zone=default_zones.zones[0].id,
            instance_name="terraform-example",
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id,
            instance_type=default_instance_types.ids[0],
            image_id=default_images.ids[0],
            internet_max_bandwidth_out=10)
        default_image = alicloud.ecs.Image("defaultImage",
            instance_id=default_instance.id,
            image_name="terraform-example",
            description="terraform-example")
        config = pulumi.Config()
        account_id = config.get("accountId")
        if account_id is None:
            account_id = "123456789"
        default_image_share_permission = alicloud.ecs.ImageSharePermission("defaultImageSharePermission",
            image_id=default_image.id,
            account_id=account_id)
        ```
        ## Attributes Reference0

         The following attributes are exported:

        * `id` - ID of the image. It formats as `<image_id>:<account_id>`

        ## Import

        image can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ecs/imageSharePermission:ImageSharePermission default m-uf66yg1q:123456789
        ```

        :param str resource_name: The name of the resource.
        :param ImageSharePermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageSharePermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageSharePermissionArgs.__new__(ImageSharePermissionArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = account_id
            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
        super(ImageSharePermission, __self__).__init__(
            'alicloud:ecs/imageSharePermission:ImageSharePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None) -> 'ImageSharePermission':
        """
        Get an existing ImageSharePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: Alibaba Cloud Account ID. It is used to share images.
        :param pulumi.Input[str] image_id: The source image ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageSharePermissionState.__new__(_ImageSharePermissionState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["image_id"] = image_id
        return ImageSharePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        Alibaba Cloud Account ID. It is used to share images.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        The source image ID.
        """
        return pulumi.get(self, "image_id")

