# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReplicaPairDrillArgs', 'ReplicaPairDrill']

@pulumi.input_type
class ReplicaPairDrillArgs:
    def __init__(__self__, *,
                 pair_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ReplicaPairDrill resource.
        :param pulumi.Input[str] pair_id: Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        """
        pulumi.set(__self__, "pair_id", pair_id)

    @property
    @pulumi.getter(name="pairId")
    def pair_id(self) -> pulumi.Input[str]:
        """
        Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        """
        return pulumi.get(self, "pair_id")

    @pair_id.setter
    def pair_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pair_id", value)


@pulumi.input_type
class _ReplicaPairDrillState:
    def __init__(__self__, *,
                 pair_id: Optional[pulumi.Input[str]] = None,
                 replica_pair_drill_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReplicaPairDrill resources.
        :param pulumi.Input[str] pair_id: Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        :param pulumi.Input[str] replica_pair_drill_id: The first ID of the resource.
        :param pulumi.Input[str] status: Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
        """
        if pair_id is not None:
            pulumi.set(__self__, "pair_id", pair_id)
        if replica_pair_drill_id is not None:
            pulumi.set(__self__, "replica_pair_drill_id", replica_pair_drill_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="pairId")
    def pair_id(self) -> Optional[pulumi.Input[str]]:
        """
        Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        """
        return pulumi.get(self, "pair_id")

    @pair_id.setter
    def pair_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pair_id", value)

    @property
    @pulumi.getter(name="replicaPairDrillId")
    def replica_pair_drill_id(self) -> Optional[pulumi.Input[str]]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "replica_pair_drill_id")

    @replica_pair_drill_id.setter
    def replica_pair_drill_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replica_pair_drill_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ReplicaPairDrill(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pair_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a EBS Replica Pair Drill resource.

        For information about EBS Replica Pair Drill and how to use it, see [What is Replica Pair Drill](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.215.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.ebs.ReplicaPairDrill("default", pair_id="pair-cn-wwo3kjfq5001")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        EBS Replica Pair Drill can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ebs/replicaPairDrill:ReplicaPairDrill example <pair_id>:<replica_pair_drill_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pair_id: Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicaPairDrillArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a EBS Replica Pair Drill resource.

        For information about EBS Replica Pair Drill and how to use it, see [What is Replica Pair Drill](https://www.alibabacloud.com/help/en/).

        > **NOTE:** Available since v1.215.0.

        ## Example Usage

        Basic Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        name = config.get("name")
        if name is None:
            name = "terraform-example"
        default = alicloud.ebs.ReplicaPairDrill("default", pair_id="pair-cn-wwo3kjfq5001")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        EBS Replica Pair Drill can be imported using the id, e.g.

        ```sh
        $ pulumi import alicloud:ebs/replicaPairDrill:ReplicaPairDrill example <pair_id>:<replica_pair_drill_id>
        ```

        :param str resource_name: The name of the resource.
        :param ReplicaPairDrillArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicaPairDrillArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pair_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicaPairDrillArgs.__new__(ReplicaPairDrillArgs)

            if pair_id is None and not opts.urn:
                raise TypeError("Missing required property 'pair_id'")
            __props__.__dict__["pair_id"] = pair_id
            __props__.__dict__["replica_pair_drill_id"] = None
            __props__.__dict__["status"] = None
        super(ReplicaPairDrill, __self__).__init__(
            'alicloud:ebs/replicaPairDrill:ReplicaPairDrill',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            pair_id: Optional[pulumi.Input[str]] = None,
            replica_pair_drill_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ReplicaPairDrill':
        """
        Get an existing ReplicaPairDrill resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pair_id: Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        :param pulumi.Input[str] replica_pair_drill_id: The first ID of the resource.
        :param pulumi.Input[str] status: Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicaPairDrillState.__new__(_ReplicaPairDrillState)

        __props__.__dict__["pair_id"] = pair_id
        __props__.__dict__["replica_pair_drill_id"] = replica_pair_drill_id
        __props__.__dict__["status"] = status
        return ReplicaPairDrill(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="pairId")
    def pair_id(self) -> pulumi.Output[str]:
        """
        Copy the ID of the pair. You can call DescribeDiskReplicaPairs to query the list of asynchronous replication pairs to obtain the replication pair ID.
        """
        return pulumi.get(self, "pair_id")

    @property
    @pulumi.getter(name="replicaPairDrillId")
    def replica_pair_drill_id(self) -> pulumi.Output[str]:
        """
        The first ID of the resource.
        """
        return pulumi.get(self, "replica_pair_drill_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Walkthrough status. _failed: Execution failed._failed: Cleanup failed.
        """
        return pulumi.get(self, "status")

