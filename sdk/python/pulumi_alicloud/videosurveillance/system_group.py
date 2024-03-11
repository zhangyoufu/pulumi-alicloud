# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SystemGroupArgs', 'SystemGroup']

@pulumi.input_type
class SystemGroupArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[str],
                 in_protocol: pulumi.Input[str],
                 out_protocol: pulumi.Input[str],
                 play_domain: pulumi.Input[str],
                 push_domain: pulumi.Input[str],
                 callback: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SystemGroup resource.
        :param pulumi.Input[str] group_name: The Group Name.
        :param pulumi.Input[str] in_protocol: The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        :param pulumi.Input[str] out_protocol: The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        :param pulumi.Input[str] play_domain: The domain name of plan streaming used by the group.
        :param pulumi.Input[str] push_domain: The domain name of push streaming used by the group.
        :param pulumi.Input[str] callback: The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        :param pulumi.Input[str] description: The description of Group.
        :param pulumi.Input[bool] enabled: Whether to open Group.
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "in_protocol", in_protocol)
        pulumi.set(__self__, "out_protocol", out_protocol)
        pulumi.set(__self__, "play_domain", play_domain)
        pulumi.set(__self__, "push_domain", push_domain)
        if callback is not None:
            pulumi.set(__self__, "callback", callback)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        The Group Name.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="inProtocol")
    def in_protocol(self) -> pulumi.Input[str]:
        """
        The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        """
        return pulumi.get(self, "in_protocol")

    @in_protocol.setter
    def in_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "in_protocol", value)

    @property
    @pulumi.getter(name="outProtocol")
    def out_protocol(self) -> pulumi.Input[str]:
        """
        The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        """
        return pulumi.get(self, "out_protocol")

    @out_protocol.setter
    def out_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "out_protocol", value)

    @property
    @pulumi.getter(name="playDomain")
    def play_domain(self) -> pulumi.Input[str]:
        """
        The domain name of plan streaming used by the group.
        """
        return pulumi.get(self, "play_domain")

    @play_domain.setter
    def play_domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "play_domain", value)

    @property
    @pulumi.getter(name="pushDomain")
    def push_domain(self) -> pulumi.Input[str]:
        """
        The domain name of push streaming used by the group.
        """
        return pulumi.get(self, "push_domain")

    @push_domain.setter
    def push_domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "push_domain", value)

    @property
    @pulumi.getter
    def callback(self) -> Optional[pulumi.Input[str]]:
        """
        The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        """
        return pulumi.get(self, "callback")

    @callback.setter
    def callback(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "callback", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of Group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to open Group.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _SystemGroupState:
    def __init__(__self__, *,
                 callback: Optional[pulumi.Input[str]] = None,
                 capture_image: Optional[pulumi.Input[int]] = None,
                 capture_interval: Optional[pulumi.Input[int]] = None,
                 capture_oss_bucket: Optional[pulumi.Input[str]] = None,
                 capture_oss_path: Optional[pulumi.Input[str]] = None,
                 capture_video: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 in_protocol: Optional[pulumi.Input[str]] = None,
                 lazy_pull: Optional[pulumi.Input[bool]] = None,
                 out_protocol: Optional[pulumi.Input[str]] = None,
                 play_domain: Optional[pulumi.Input[str]] = None,
                 push_domain: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering SystemGroup resources.
        :param pulumi.Input[str] callback: The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        :param pulumi.Input[int] capture_image: The capture image.
        :param pulumi.Input[int] capture_interval: The capture interval.
        :param pulumi.Input[str] capture_oss_bucket: The capture oss bucket.
        :param pulumi.Input[str] capture_oss_path: The capture oss path.
        :param pulumi.Input[int] capture_video: The capture video.
        :param pulumi.Input[str] description: The description of Group.
        :param pulumi.Input[bool] enabled: Whether to open Group.
        :param pulumi.Input[str] group_name: The Group Name.
        :param pulumi.Input[str] in_protocol: The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        :param pulumi.Input[bool] lazy_pull: Whether to enable on-demand streaming. Default value:`false`.
        :param pulumi.Input[str] out_protocol: The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        :param pulumi.Input[str] play_domain: The domain name of plan streaming used by the group.
        :param pulumi.Input[str] push_domain: The domain name of push streaming used by the group.
        :param pulumi.Input[bool] status: Whether to open Group. Valid values: `on`,`off`.
        """
        if callback is not None:
            pulumi.set(__self__, "callback", callback)
        if capture_image is not None:
            pulumi.set(__self__, "capture_image", capture_image)
        if capture_interval is not None:
            pulumi.set(__self__, "capture_interval", capture_interval)
        if capture_oss_bucket is not None:
            pulumi.set(__self__, "capture_oss_bucket", capture_oss_bucket)
        if capture_oss_path is not None:
            pulumi.set(__self__, "capture_oss_path", capture_oss_path)
        if capture_video is not None:
            pulumi.set(__self__, "capture_video", capture_video)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if in_protocol is not None:
            pulumi.set(__self__, "in_protocol", in_protocol)
        if lazy_pull is not None:
            pulumi.set(__self__, "lazy_pull", lazy_pull)
        if out_protocol is not None:
            pulumi.set(__self__, "out_protocol", out_protocol)
        if play_domain is not None:
            pulumi.set(__self__, "play_domain", play_domain)
        if push_domain is not None:
            pulumi.set(__self__, "push_domain", push_domain)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def callback(self) -> Optional[pulumi.Input[str]]:
        """
        The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        """
        return pulumi.get(self, "callback")

    @callback.setter
    def callback(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "callback", value)

    @property
    @pulumi.getter(name="captureImage")
    def capture_image(self) -> Optional[pulumi.Input[int]]:
        """
        The capture image.
        """
        return pulumi.get(self, "capture_image")

    @capture_image.setter
    def capture_image(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capture_image", value)

    @property
    @pulumi.getter(name="captureInterval")
    def capture_interval(self) -> Optional[pulumi.Input[int]]:
        """
        The capture interval.
        """
        return pulumi.get(self, "capture_interval")

    @capture_interval.setter
    def capture_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capture_interval", value)

    @property
    @pulumi.getter(name="captureOssBucket")
    def capture_oss_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The capture oss bucket.
        """
        return pulumi.get(self, "capture_oss_bucket")

    @capture_oss_bucket.setter
    def capture_oss_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "capture_oss_bucket", value)

    @property
    @pulumi.getter(name="captureOssPath")
    def capture_oss_path(self) -> Optional[pulumi.Input[str]]:
        """
        The capture oss path.
        """
        return pulumi.get(self, "capture_oss_path")

    @capture_oss_path.setter
    def capture_oss_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "capture_oss_path", value)

    @property
    @pulumi.getter(name="captureVideo")
    def capture_video(self) -> Optional[pulumi.Input[int]]:
        """
        The capture video.
        """
        return pulumi.get(self, "capture_video")

    @capture_video.setter
    def capture_video(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capture_video", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of Group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to open Group.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Group Name.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="inProtocol")
    def in_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        """
        return pulumi.get(self, "in_protocol")

    @in_protocol.setter
    def in_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "in_protocol", value)

    @property
    @pulumi.getter(name="lazyPull")
    def lazy_pull(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable on-demand streaming. Default value:`false`.
        """
        return pulumi.get(self, "lazy_pull")

    @lazy_pull.setter
    def lazy_pull(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lazy_pull", value)

    @property
    @pulumi.getter(name="outProtocol")
    def out_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        """
        return pulumi.get(self, "out_protocol")

    @out_protocol.setter
    def out_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "out_protocol", value)

    @property
    @pulumi.getter(name="playDomain")
    def play_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name of plan streaming used by the group.
        """
        return pulumi.get(self, "play_domain")

    @play_domain.setter
    def play_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "play_domain", value)

    @property
    @pulumi.getter(name="pushDomain")
    def push_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name of push streaming used by the group.
        """
        return pulumi.get(self, "push_domain")

    @push_domain.setter
    def push_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_domain", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to open Group. Valid values: `on`,`off`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status", value)


class SystemGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 callback: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 in_protocol: Optional[pulumi.Input[str]] = None,
                 out_protocol: Optional[pulumi.Input[str]] = None,
                 play_domain: Optional[pulumi.Input[str]] = None,
                 push_domain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Video Surveillance System Group resource.

        For information about Video Surveillance System Group and how to use it, see [What is Group](https://help.aliyun.com/product/108765.html).

        > **NOTE:** Available in v1.135.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.videosurveillance.SystemGroup("default",
            group_name="your_group_name",
            in_protocol="rtmp",
            out_protocol="flv",
            play_domain="your_plan_domain",
            push_domain="your_push_domain")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Video Surveillance System Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:videosurveillance/systemGroup:SystemGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] callback: The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        :param pulumi.Input[str] description: The description of Group.
        :param pulumi.Input[bool] enabled: Whether to open Group.
        :param pulumi.Input[str] group_name: The Group Name.
        :param pulumi.Input[str] in_protocol: The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        :param pulumi.Input[str] out_protocol: The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        :param pulumi.Input[str] play_domain: The domain name of plan streaming used by the group.
        :param pulumi.Input[str] push_domain: The domain name of push streaming used by the group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Video Surveillance System Group resource.

        For information about Video Surveillance System Group and how to use it, see [What is Group](https://help.aliyun.com/product/108765.html).

        > **NOTE:** Available in v1.135.0+.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.videosurveillance.SystemGroup("default",
            group_name="your_group_name",
            in_protocol="rtmp",
            out_protocol="flv",
            play_domain="your_plan_domain",
            push_domain="your_push_domain")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Video Surveillance System Group can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:videosurveillance/systemGroup:SystemGroup example <id>
        ```

        :param str resource_name: The name of the resource.
        :param SystemGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 callback: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 in_protocol: Optional[pulumi.Input[str]] = None,
                 out_protocol: Optional[pulumi.Input[str]] = None,
                 play_domain: Optional[pulumi.Input[str]] = None,
                 push_domain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemGroupArgs.__new__(SystemGroupArgs)

            __props__.__dict__["callback"] = callback
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            if in_protocol is None and not opts.urn:
                raise TypeError("Missing required property 'in_protocol'")
            __props__.__dict__["in_protocol"] = in_protocol
            if out_protocol is None and not opts.urn:
                raise TypeError("Missing required property 'out_protocol'")
            __props__.__dict__["out_protocol"] = out_protocol
            if play_domain is None and not opts.urn:
                raise TypeError("Missing required property 'play_domain'")
            __props__.__dict__["play_domain"] = play_domain
            if push_domain is None and not opts.urn:
                raise TypeError("Missing required property 'push_domain'")
            __props__.__dict__["push_domain"] = push_domain
            __props__.__dict__["capture_image"] = None
            __props__.__dict__["capture_interval"] = None
            __props__.__dict__["capture_oss_bucket"] = None
            __props__.__dict__["capture_oss_path"] = None
            __props__.__dict__["capture_video"] = None
            __props__.__dict__["lazy_pull"] = None
            __props__.__dict__["status"] = None
        super(SystemGroup, __self__).__init__(
            'alicloud:videosurveillance/systemGroup:SystemGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            callback: Optional[pulumi.Input[str]] = None,
            capture_image: Optional[pulumi.Input[int]] = None,
            capture_interval: Optional[pulumi.Input[int]] = None,
            capture_oss_bucket: Optional[pulumi.Input[str]] = None,
            capture_oss_path: Optional[pulumi.Input[str]] = None,
            capture_video: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            in_protocol: Optional[pulumi.Input[str]] = None,
            lazy_pull: Optional[pulumi.Input[bool]] = None,
            out_protocol: Optional[pulumi.Input[str]] = None,
            play_domain: Optional[pulumi.Input[str]] = None,
            push_domain: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[bool]] = None) -> 'SystemGroup':
        """
        Get an existing SystemGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] callback: The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        :param pulumi.Input[int] capture_image: The capture image.
        :param pulumi.Input[int] capture_interval: The capture interval.
        :param pulumi.Input[str] capture_oss_bucket: The capture oss bucket.
        :param pulumi.Input[str] capture_oss_path: The capture oss path.
        :param pulumi.Input[int] capture_video: The capture video.
        :param pulumi.Input[str] description: The description of Group.
        :param pulumi.Input[bool] enabled: Whether to open Group.
        :param pulumi.Input[str] group_name: The Group Name.
        :param pulumi.Input[str] in_protocol: The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        :param pulumi.Input[bool] lazy_pull: Whether to enable on-demand streaming. Default value:`false`.
        :param pulumi.Input[str] out_protocol: The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        :param pulumi.Input[str] play_domain: The domain name of plan streaming used by the group.
        :param pulumi.Input[str] push_domain: The domain name of push streaming used by the group.
        :param pulumi.Input[bool] status: Whether to open Group. Valid values: `on`,`off`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemGroupState.__new__(_SystemGroupState)

        __props__.__dict__["callback"] = callback
        __props__.__dict__["capture_image"] = capture_image
        __props__.__dict__["capture_interval"] = capture_interval
        __props__.__dict__["capture_oss_bucket"] = capture_oss_bucket
        __props__.__dict__["capture_oss_path"] = capture_oss_path
        __props__.__dict__["capture_video"] = capture_video
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["in_protocol"] = in_protocol
        __props__.__dict__["lazy_pull"] = lazy_pull
        __props__.__dict__["out_protocol"] = out_protocol
        __props__.__dict__["play_domain"] = play_domain
        __props__.__dict__["push_domain"] = push_domain
        __props__.__dict__["status"] = status
        return SystemGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def callback(self) -> pulumi.Output[Optional[str]]:
        """
        The space within the device status update of the callback, need to start with http:// or https:// at the beginning.
        """
        return pulumi.get(self, "callback")

    @property
    @pulumi.getter(name="captureImage")
    def capture_image(self) -> pulumi.Output[int]:
        """
        The capture image.
        """
        return pulumi.get(self, "capture_image")

    @property
    @pulumi.getter(name="captureInterval")
    def capture_interval(self) -> pulumi.Output[int]:
        """
        The capture interval.
        """
        return pulumi.get(self, "capture_interval")

    @property
    @pulumi.getter(name="captureOssBucket")
    def capture_oss_bucket(self) -> pulumi.Output[str]:
        """
        The capture oss bucket.
        """
        return pulumi.get(self, "capture_oss_bucket")

    @property
    @pulumi.getter(name="captureOssPath")
    def capture_oss_path(self) -> pulumi.Output[str]:
        """
        The capture oss path.
        """
        return pulumi.get(self, "capture_oss_path")

    @property
    @pulumi.getter(name="captureVideo")
    def capture_video(self) -> pulumi.Output[int]:
        """
        The capture video.
        """
        return pulumi.get(self, "capture_video")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of Group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether to open Group.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        The Group Name.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="inProtocol")
    def in_protocol(self) -> pulumi.Output[str]:
        """
        The use of the access protocol support gb28181, Real Time Messaging Protocol (rtmp). Valid values: `gb28181`, `rtmp`.
        """
        return pulumi.get(self, "in_protocol")

    @property
    @pulumi.getter(name="lazyPull")
    def lazy_pull(self) -> pulumi.Output[bool]:
        """
        Whether to enable on-demand streaming. Default value:`false`.
        """
        return pulumi.get(self, "lazy_pull")

    @property
    @pulumi.getter(name="outProtocol")
    def out_protocol(self) -> pulumi.Output[str]:
        """
        The playback protocol used by the space, multiple values are separated by commas (,). Valid values: `flv`,`hls`, `rtmp`.
        """
        return pulumi.get(self, "out_protocol")

    @property
    @pulumi.getter(name="playDomain")
    def play_domain(self) -> pulumi.Output[str]:
        """
        The domain name of plan streaming used by the group.
        """
        return pulumi.get(self, "play_domain")

    @property
    @pulumi.getter(name="pushDomain")
    def push_domain(self) -> pulumi.Output[str]:
        """
        The domain name of push streaming used by the group.
        """
        return pulumi.get(self, "push_domain")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[bool]:
        """
        Whether to open Group. Valid values: `on`,`off`.
        """
        return pulumi.get(self, "status")

