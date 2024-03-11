# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetClustersResult',
    'AwaitableGetClustersResult',
    'get_clusters',
    'get_clusters_output',
]

@pulumi.output_type
class GetClustersResult:
    """
    A collection of values returned by getClusters.
    """
    def __init__(__self__, cluster_name=None, cluster_type_lists=None, clusters=None, create_type=None, default_status=None, deposit_type=None, enable_details=None, id=None, ids=None, is_desc=None, machine_type=None, name_regex=None, names=None, output_file=None, page_number=None, page_size=None, resource_group_id=None, status_lists=None, total_count=None, vpc_id=None):
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if cluster_type_lists and not isinstance(cluster_type_lists, list):
            raise TypeError("Expected argument 'cluster_type_lists' to be a list")
        pulumi.set(__self__, "cluster_type_lists", cluster_type_lists)
        if clusters and not isinstance(clusters, list):
            raise TypeError("Expected argument 'clusters' to be a list")
        pulumi.set(__self__, "clusters", clusters)
        if create_type and not isinstance(create_type, str):
            raise TypeError("Expected argument 'create_type' to be a str")
        pulumi.set(__self__, "create_type", create_type)
        if default_status and not isinstance(default_status, bool):
            raise TypeError("Expected argument 'default_status' to be a bool")
        pulumi.set(__self__, "default_status", default_status)
        if deposit_type and not isinstance(deposit_type, str):
            raise TypeError("Expected argument 'deposit_type' to be a str")
        pulumi.set(__self__, "deposit_type", deposit_type)
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if is_desc and not isinstance(is_desc, bool):
            raise TypeError("Expected argument 'is_desc' to be a bool")
        pulumi.set(__self__, "is_desc", is_desc)
        if machine_type and not isinstance(machine_type, str):
            raise TypeError("Expected argument 'machine_type' to be a str")
        pulumi.set(__self__, "machine_type", machine_type)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if page_number and not isinstance(page_number, int):
            raise TypeError("Expected argument 'page_number' to be a int")
        pulumi.set(__self__, "page_number", page_number)
        if page_size and not isinstance(page_size, int):
            raise TypeError("Expected argument 'page_size' to be a int")
        pulumi.set(__self__, "page_size", page_size)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status_lists and not isinstance(status_lists, list):
            raise TypeError("Expected argument 'status_lists' to be a list")
        pulumi.set(__self__, "status_lists", status_lists)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[str]:
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="clusterTypeLists")
    def cluster_type_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cluster_type_lists")

    @property
    @pulumi.getter
    def clusters(self) -> Sequence['outputs.GetClustersClusterResult']:
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="createType")
    def create_type(self) -> Optional[str]:
        return pulumi.get(self, "create_type")

    @property
    @pulumi.getter(name="defaultStatus")
    def default_status(self) -> Optional[bool]:
        return pulumi.get(self, "default_status")

    @property
    @pulumi.getter(name="depositType")
    def deposit_type(self) -> Optional[str]:
        return pulumi.get(self, "deposit_type")

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="isDesc")
    def is_desc(self) -> Optional[bool]:
        return pulumi.get(self, "is_desc")

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> Optional[str]:
        return pulumi.get(self, "machine_type")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="pageNumber")
    def page_number(self) -> Optional[int]:
        return pulumi.get(self, "page_number")

    @property
    @pulumi.getter(name="pageSize")
    def page_size(self) -> Optional[int]:
        return pulumi.get(self, "page_size")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter(name="statusLists")
    def status_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "status_lists")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableGetClustersResult(GetClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClustersResult(
            cluster_name=self.cluster_name,
            cluster_type_lists=self.cluster_type_lists,
            clusters=self.clusters,
            create_type=self.create_type,
            default_status=self.default_status,
            deposit_type=self.deposit_type,
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            is_desc=self.is_desc,
            machine_type=self.machine_type,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            page_number=self.page_number,
            page_size=self.page_size,
            resource_group_id=self.resource_group_id,
            status_lists=self.status_lists,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def get_clusters(cluster_name: Optional[str] = None,
                 cluster_type_lists: Optional[Sequence[str]] = None,
                 create_type: Optional[str] = None,
                 default_status: Optional[bool] = None,
                 deposit_type: Optional[str] = None,
                 enable_details: Optional[bool] = None,
                 ids: Optional[Sequence[str]] = None,
                 is_desc: Optional[bool] = None,
                 machine_type: Optional[str] = None,
                 name_regex: Optional[str] = None,
                 output_file: Optional[str] = None,
                 page_number: Optional[int] = None,
                 page_size: Optional[int] = None,
                 resource_group_id: Optional[str] = None,
                 status_lists: Optional[Sequence[str]] = None,
                 vpc_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClustersResult:
    """
    This data source provides the Emr Clusters of the current Alibaba Cloud user.

    > **DEPRECATED:**  This datasource has been deprecated from version `1.204.0`. Please use new datasource emrv2_clusters.

    > **NOTE:** Available in v1.146.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testAccClusters"
    default_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
    default_main_versions = alicloud.emr.get_main_versions()
    default_instance_types = alicloud.emr.get_instance_types(destination_resource="InstanceType",
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        support_local_storage=False,
        instance_charge_type="PostPaid",
        support_node_types=[
            "MASTER",
            "CORE",
            "TASK",
        ])
    data_disk = alicloud.emr.get_disk_types(destination_resource="DataDisk",
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        instance_charge_type="PostPaid",
        instance_type=default_instance_types.types[0].id,
        zone_id=default_instance_types.types[0].zone_id)
    system_disk = alicloud.emr.get_disk_types(destination_resource="SystemDisk",
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        instance_charge_type="PostPaid",
        instance_type=default_instance_types.types[0].id,
        zone_id=default_instance_types.types[0].zone_id)
    default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
    default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_networks.ids[0])
    default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
        zone_id=default_instance_types.types[0].zone_id)
    default_role = alicloud.ram.Role("defaultRole",
        document=\"\"\"    {
            "Statement": [
            {
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Principal": {
                "Service": [
                    "emr.aliyuncs.com",
                    "ecs.aliyuncs.com"
                ]
                }
            }
            ],
            "Version": "1"
        }
    \"\"\",
        description="this is a role test.",
        force=True)
    default_cluster = alicloud.emr.Cluster("defaultCluster",
        emr_ver=default_main_versions.main_versions[0].emr_version,
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        host_groups=[
            alicloud.emr.ClusterHostGroupArgs(
                host_group_name="master_group",
                host_group_type="MASTER",
                node_count="2",
                instance_type=default_instance_types.types[0].id,
                disk_type=data_disk.types[0].value,
                disk_capacity=data_disk.types[0].min if data_disk.types[0].min > 160 else "160",
                disk_count="1",
                sys_disk_type=system_disk.types[0].value,
                sys_disk_capacity=system_disk.types[0].min if system_disk.types[0].min > 160 else "160",
            ),
            alicloud.emr.ClusterHostGroupArgs(
                host_group_name="core_group",
                host_group_type="CORE",
                node_count="3",
                instance_type=default_instance_types.types[0].id,
                disk_type=data_disk.types[0].value,
                disk_capacity=data_disk.types[0].min if data_disk.types[0].min > 160 else "160",
                disk_count="4",
                sys_disk_type=system_disk.types[0].value,
                sys_disk_capacity=system_disk.types[0].min if system_disk.types[0].min > 160 else "160",
            ),
            alicloud.emr.ClusterHostGroupArgs(
                host_group_name="task_group",
                host_group_type="TASK",
                node_count="2",
                instance_type=default_instance_types.types[0].id,
                disk_type=data_disk.types[0].value,
                disk_capacity=data_disk.types[0].min if data_disk.types[0].min > 160 else "160",
                disk_count="4",
                sys_disk_type=system_disk.types[0].value,
                sys_disk_capacity=system_disk.types[0].min if system_disk.types[0].min > 160 else "160",
            ),
        ],
        high_availability_enable=True,
        zone_id=default_instance_types.types[0].zone_id,
        security_group_id=default_security_group.id,
        is_open_public_ip=True,
        charge_type="PostPaid",
        vswitch_id=default_switches.ids[0],
        user_defined_emr_ecs_role=default_role.name,
        ssh_enable=True,
        master_pwd="ABCtest1234!",
        tags={
            "Created": "TF",
            "For": "acceptance test",
        })
    ids = alicloud.emr.get_clusters()
    pulumi.export("emrClusterId1", ids.clusters[0].id)
    name_regex = alicloud.emr.get_clusters_output(name_regex=default_cluster.name)
    pulumi.export("emrClusterId2", name_regex.clusters[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_name: The name of the associated cluster.
    :param Sequence[str] cluster_type_lists: The cluster type list.
    :param str create_type: How to create a cluster.
    :param bool default_status: The default status.
    :param str deposit_type: The hosting type of the cluster.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Cluster IDs.
    :param bool is_desc: The is desc.
    :param str machine_type: The host type of the cluster. The default is ECS.
    :param str name_regex: A regex string to filter results by Cluster name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The Resource Group ID.
    :param Sequence[str] status_lists: The status list. Valid values: `ABNORMAL`, `CREATE_FAILED`, `CREATING`, `IDLE`, `RELEASED`, `RELEASE_FAILED`, `RELEASING`, `RUNNING`, `WAIT_FOR_PAY`.
    :param str vpc_id: The VPC ID.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['clusterTypeLists'] = cluster_type_lists
    __args__['createType'] = create_type
    __args__['defaultStatus'] = default_status
    __args__['depositType'] = deposit_type
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['isDesc'] = is_desc
    __args__['machineType'] = machine_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['pageNumber'] = page_number
    __args__['pageSize'] = page_size
    __args__['resourceGroupId'] = resource_group_id
    __args__['statusLists'] = status_lists
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:emr/getClusters:getClusters', __args__, opts=opts, typ=GetClustersResult).value

    return AwaitableGetClustersResult(
        cluster_name=pulumi.get(__ret__, 'cluster_name'),
        cluster_type_lists=pulumi.get(__ret__, 'cluster_type_lists'),
        clusters=pulumi.get(__ret__, 'clusters'),
        create_type=pulumi.get(__ret__, 'create_type'),
        default_status=pulumi.get(__ret__, 'default_status'),
        deposit_type=pulumi.get(__ret__, 'deposit_type'),
        enable_details=pulumi.get(__ret__, 'enable_details'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        is_desc=pulumi.get(__ret__, 'is_desc'),
        machine_type=pulumi.get(__ret__, 'machine_type'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        page_number=pulumi.get(__ret__, 'page_number'),
        page_size=pulumi.get(__ret__, 'page_size'),
        resource_group_id=pulumi.get(__ret__, 'resource_group_id'),
        status_lists=pulumi.get(__ret__, 'status_lists'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_clusters)
def get_clusters_output(cluster_name: Optional[pulumi.Input[Optional[str]]] = None,
                        cluster_type_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        create_type: Optional[pulumi.Input[Optional[str]]] = None,
                        default_status: Optional[pulumi.Input[Optional[bool]]] = None,
                        deposit_type: Optional[pulumi.Input[Optional[str]]] = None,
                        enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        is_desc: Optional[pulumi.Input[Optional[bool]]] = None,
                        machine_type: Optional[pulumi.Input[Optional[str]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        page_number: Optional[pulumi.Input[Optional[int]]] = None,
                        page_size: Optional[pulumi.Input[Optional[int]]] = None,
                        resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                        status_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClustersResult]:
    """
    This data source provides the Emr Clusters of the current Alibaba Cloud user.

    > **DEPRECATED:**  This datasource has been deprecated from version `1.204.0`. Please use new datasource emrv2_clusters.

    > **NOTE:** Available in v1.146.0+.

    ## Example Usage

    Basic Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "tf-testAccClusters"
    default_resource_groups = alicloud.resourcemanager.get_resource_groups(status="OK")
    default_main_versions = alicloud.emr.get_main_versions()
    default_instance_types = alicloud.emr.get_instance_types(destination_resource="InstanceType",
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        support_local_storage=False,
        instance_charge_type="PostPaid",
        support_node_types=[
            "MASTER",
            "CORE",
            "TASK",
        ])
    data_disk = alicloud.emr.get_disk_types(destination_resource="DataDisk",
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        instance_charge_type="PostPaid",
        instance_type=default_instance_types.types[0].id,
        zone_id=default_instance_types.types[0].zone_id)
    system_disk = alicloud.emr.get_disk_types(destination_resource="SystemDisk",
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        instance_charge_type="PostPaid",
        instance_type=default_instance_types.types[0].id,
        zone_id=default_instance_types.types[0].zone_id)
    default_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
    default_security_group = alicloud.ecs.SecurityGroup("defaultSecurityGroup", vpc_id=default_networks.ids[0])
    default_switches = alicloud.vpc.get_switches(vpc_id=default_networks.ids[0],
        zone_id=default_instance_types.types[0].zone_id)
    default_role = alicloud.ram.Role("defaultRole",
        document=\"\"\"    {
            "Statement": [
            {
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Principal": {
                "Service": [
                    "emr.aliyuncs.com",
                    "ecs.aliyuncs.com"
                ]
                }
            }
            ],
            "Version": "1"
        }
    \"\"\",
        description="this is a role test.",
        force=True)
    default_cluster = alicloud.emr.Cluster("defaultCluster",
        emr_ver=default_main_versions.main_versions[0].emr_version,
        cluster_type=default_main_versions.main_versions[0].cluster_types[0],
        host_groups=[
            alicloud.emr.ClusterHostGroupArgs(
                host_group_name="master_group",
                host_group_type="MASTER",
                node_count="2",
                instance_type=default_instance_types.types[0].id,
                disk_type=data_disk.types[0].value,
                disk_capacity=data_disk.types[0].min if data_disk.types[0].min > 160 else "160",
                disk_count="1",
                sys_disk_type=system_disk.types[0].value,
                sys_disk_capacity=system_disk.types[0].min if system_disk.types[0].min > 160 else "160",
            ),
            alicloud.emr.ClusterHostGroupArgs(
                host_group_name="core_group",
                host_group_type="CORE",
                node_count="3",
                instance_type=default_instance_types.types[0].id,
                disk_type=data_disk.types[0].value,
                disk_capacity=data_disk.types[0].min if data_disk.types[0].min > 160 else "160",
                disk_count="4",
                sys_disk_type=system_disk.types[0].value,
                sys_disk_capacity=system_disk.types[0].min if system_disk.types[0].min > 160 else "160",
            ),
            alicloud.emr.ClusterHostGroupArgs(
                host_group_name="task_group",
                host_group_type="TASK",
                node_count="2",
                instance_type=default_instance_types.types[0].id,
                disk_type=data_disk.types[0].value,
                disk_capacity=data_disk.types[0].min if data_disk.types[0].min > 160 else "160",
                disk_count="4",
                sys_disk_type=system_disk.types[0].value,
                sys_disk_capacity=system_disk.types[0].min if system_disk.types[0].min > 160 else "160",
            ),
        ],
        high_availability_enable=True,
        zone_id=default_instance_types.types[0].zone_id,
        security_group_id=default_security_group.id,
        is_open_public_ip=True,
        charge_type="PostPaid",
        vswitch_id=default_switches.ids[0],
        user_defined_emr_ecs_role=default_role.name,
        ssh_enable=True,
        master_pwd="ABCtest1234!",
        tags={
            "Created": "TF",
            "For": "acceptance test",
        })
    ids = alicloud.emr.get_clusters()
    pulumi.export("emrClusterId1", ids.clusters[0].id)
    name_regex = alicloud.emr.get_clusters_output(name_regex=default_cluster.name)
    pulumi.export("emrClusterId2", name_regex.clusters[0].id)
    ```
    <!--End PulumiCodeChooser -->


    :param str cluster_name: The name of the associated cluster.
    :param Sequence[str] cluster_type_lists: The cluster type list.
    :param str create_type: How to create a cluster.
    :param bool default_status: The default status.
    :param str deposit_type: The hosting type of the cluster.
    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Cluster IDs.
    :param bool is_desc: The is desc.
    :param str machine_type: The host type of the cluster. The default is ECS.
    :param str name_regex: A regex string to filter results by Cluster name.
    :param str output_file: File name where to save data source results (after running `pulumi preview`).
    :param str resource_group_id: The Resource Group ID.
    :param Sequence[str] status_lists: The status list. Valid values: `ABNORMAL`, `CREATE_FAILED`, `CREATING`, `IDLE`, `RELEASED`, `RELEASE_FAILED`, `RELEASING`, `RUNNING`, `WAIT_FOR_PAY`.
    :param str vpc_id: The VPC ID.
    """
    ...
