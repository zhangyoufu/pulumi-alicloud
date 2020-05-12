# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SlbAttachment(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    The ID of the applicaton to which you want to bind an SLB instance.
    """
    listener_port: pulumi.Output[float]
    """
    The listening port for the bound SLB instance.
    """
    slb_id: pulumi.Output[str]
    """
    The ID of the SLB instance that is going to be bound.
    """
    slb_ip: pulumi.Output[str]
    """
    The IP address that is allocated to the bound SLB instance.
    """
    slb_status: pulumi.Output[str]
    """
    Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and foward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.
    """
    type: pulumi.Output[str]
    """
    The type of the bound SLB instance.
    """
    vserver_group_id: pulumi.Output[str]
    """
    The ID of the virtual server (VServer) group associated with the intranet SLB instance.
    """
    vswitch_id: pulumi.Output[str]
    """
    VPC related vswitch ID.
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, listener_port=None, slb_id=None, slb_ip=None, type=None, vserver_group_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an EDAS slb application attachment resource.

        > **NOTE:** Available in 1.82.0+

        ## Example Usage



        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.edas.SlbAttachment("default",
            app_id=var["app_id"],
            slb_id=var["slb_id"],
            slb_ip=var["slb_ip"],
            type=var["type"],
            listener_port=var["listener_port"],
            vserver_group_id=var["vserver_group_id"])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the applicaton to which you want to bind an SLB instance.
        :param pulumi.Input[float] listener_port: The listening port for the bound SLB instance.
        :param pulumi.Input[str] slb_id: The ID of the SLB instance that is going to be bound.
        :param pulumi.Input[str] slb_ip: The IP address that is allocated to the bound SLB instance.
        :param pulumi.Input[str] type: The type of the bound SLB instance.
        :param pulumi.Input[str] vserver_group_id: The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if app_id is None:
                raise TypeError("Missing required property 'app_id'")
            __props__['app_id'] = app_id
            __props__['listener_port'] = listener_port
            if slb_id is None:
                raise TypeError("Missing required property 'slb_id'")
            __props__['slb_id'] = slb_id
            if slb_ip is None:
                raise TypeError("Missing required property 'slb_ip'")
            __props__['slb_ip'] = slb_ip
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['vserver_group_id'] = vserver_group_id
            __props__['slb_status'] = None
            __props__['vswitch_id'] = None
        super(SlbAttachment, __self__).__init__(
            'alicloud:edas/slbAttachment:SlbAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, listener_port=None, slb_id=None, slb_ip=None, slb_status=None, type=None, vserver_group_id=None, vswitch_id=None):
        """
        Get an existing SlbAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the applicaton to which you want to bind an SLB instance.
        :param pulumi.Input[float] listener_port: The listening port for the bound SLB instance.
        :param pulumi.Input[str] slb_id: The ID of the SLB instance that is going to be bound.
        :param pulumi.Input[str] slb_ip: The IP address that is allocated to the bound SLB instance.
        :param pulumi.Input[str] slb_status: Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and foward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.
        :param pulumi.Input[str] type: The type of the bound SLB instance.
        :param pulumi.Input[str] vserver_group_id: The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        :param pulumi.Input[str] vswitch_id: VPC related vswitch ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["listener_port"] = listener_port
        __props__["slb_id"] = slb_id
        __props__["slb_ip"] = slb_ip
        __props__["slb_status"] = slb_status
        __props__["type"] = type
        __props__["vserver_group_id"] = vserver_group_id
        __props__["vswitch_id"] = vswitch_id
        return SlbAttachment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

