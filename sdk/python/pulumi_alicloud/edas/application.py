# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Application']


class Application(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 build_pack_id: Optional[pulumi.Input[float]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 descriotion: Optional[pulumi.Input[str]] = None,
                 ecu_infos: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 health_check_url: Optional[pulumi.Input[str]] = None,
                 logical_region_id: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an EDAS ecs application on EDAS. The application will be deployed when `group_id` and `war_url` are given.

        > **NOTE:** Available in 1.82.0+

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param pulumi.Input[float] build_pack_id: The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        :param pulumi.Input[str] cluster_id: The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        :param pulumi.Input[str] descriotion: The description of the application that you want to create.
        :param pulumi.Input[List[pulumi.Input[str]]] ecu_infos: The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] health_check_url: The URL for health checking of the application.
        :param pulumi.Input[str] logical_region_id: The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        :param pulumi.Input[str] package_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if application_name is None:
                raise TypeError("Missing required property 'application_name'")
            __props__['application_name'] = application_name
            __props__['build_pack_id'] = build_pack_id
            if cluster_id is None:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            __props__['descriotion'] = descriotion
            __props__['ecu_infos'] = ecu_infos
            __props__['group_id'] = group_id
            __props__['health_check_url'] = health_check_url
            __props__['logical_region_id'] = logical_region_id
            if package_type is None:
                raise TypeError("Missing required property 'package_type'")
            __props__['package_type'] = package_type
            __props__['package_version'] = package_version
            __props__['war_url'] = war_url
        super(Application, __self__).__init__(
            'alicloud:edas/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_name: Optional[pulumi.Input[str]] = None,
            build_pack_id: Optional[pulumi.Input[float]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            descriotion: Optional[pulumi.Input[str]] = None,
            ecu_infos: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            health_check_url: Optional[pulumi.Input[str]] = None,
            logical_region_id: Optional[pulumi.Input[str]] = None,
            package_type: Optional[pulumi.Input[str]] = None,
            package_version: Optional[pulumi.Input[str]] = None,
            war_url: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        :param pulumi.Input[float] build_pack_id: The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        :param pulumi.Input[str] cluster_id: The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        :param pulumi.Input[str] descriotion: The description of the application that you want to create.
        :param pulumi.Input[List[pulumi.Input[str]]] ecu_infos: The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] health_check_url: The URL for health checking of the application.
        :param pulumi.Input[str] logical_region_id: The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        :param pulumi.Input[str] package_type: The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_name"] = application_name
        __props__["build_pack_id"] = build_pack_id
        __props__["cluster_id"] = cluster_id
        __props__["descriotion"] = descriotion
        __props__["ecu_infos"] = ecu_infos
        __props__["group_id"] = group_id
        __props__["health_check_url"] = health_check_url
        __props__["logical_region_id"] = logical_region_id
        __props__["package_type"] = package_type
        __props__["package_version"] = package_version
        __props__["war_url"] = war_url
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> str:
        """
        Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
        """
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="buildPackId")
    def build_pack_id(self) -> Optional[float]:
        """
        The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.
        """
        return pulumi.get(self, "build_pack_id")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def descriotion(self) -> Optional[str]:
        """
        The description of the application that you want to create.
        """
        return pulumi.get(self, "descriotion")

    @property
    @pulumi.getter(name="ecuInfos")
    def ecu_infos(self) -> Optional[List[str]]:
        """
        The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.
        """
        return pulumi.get(self, "ecu_infos")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="healthCheckUrl")
    def health_check_url(self) -> Optional[str]:
        """
        The URL for health checking of the application.
        """
        return pulumi.get(self, "health_check_url")

    @property
    @pulumi.getter(name="logicalRegionId")
    def logical_region_id(self) -> Optional[str]:
        """
        The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
        """
        return pulumi.get(self, "logical_region_id")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        """
        The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
        """
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> Optional[str]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> Optional[str]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

