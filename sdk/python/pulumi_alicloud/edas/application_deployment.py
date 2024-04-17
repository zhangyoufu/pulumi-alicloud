# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationDeploymentArgs', 'ApplicationDeployment']

@pulumi.input_type
class ApplicationDeploymentArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 war_url: pulumi.Input[str],
                 package_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationDeployment resource.
        :param pulumi.Input[str] app_id: The ID of the application that you want to deploy.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "war_url", war_url)
        if package_version is not None:
            pulumi.set(__self__, "package_version", package_version)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        The ID of the application that you want to deploy.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> pulumi.Input[str]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

    @war_url.setter
    def war_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "war_url", value)

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @package_version.setter
    def package_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_version", value)


@pulumi.input_type
class _ApplicationDeploymentState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 last_package_version: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationDeployment resources.
        :param pulumi.Input[str] app_id: The ID of the application that you want to deploy.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] last_package_version: Last package version deployed.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if last_package_version is not None:
            pulumi.set(__self__, "last_package_version", last_package_version)
        if package_version is not None:
            pulumi.set(__self__, "package_version", package_version)
        if war_url is not None:
            pulumi.set(__self__, "war_url", war_url)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the application that you want to deploy.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="lastPackageVersion")
    def last_package_version(self) -> Optional[pulumi.Input[str]]:
        """
        Last package version deployed.
        """
        return pulumi.get(self, "last_package_version")

    @last_package_version.setter
    def last_package_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_package_version", value)

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @package_version.setter
    def package_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_version", value)

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> Optional[pulumi.Input[str]]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

    @war_url.setter
    def war_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "war_url", value)


class ApplicationDeployment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Deploys applications on EDAS, see [What is EDAS Application Deployment](https://www.alibabacloud.com/help/en/edas/developer-reference/api-edas-2017-08-01-deployapplication).

        > **NOTE:** Available since v1.82.0.

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
        default = alicloud.get_regions(current=True)
        default_get_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_get_zones.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default_get_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default", vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("default",
            availability_zone=default_get_zones.zones[0].id,
            instance_name=name,
            image_id=default_get_images.images[0].id,
            instance_type=default_get_instance_types.instance_types[0].id,
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id)
        default_cluster = alicloud.edas.Cluster("default",
            cluster_name=name,
            cluster_type=2,
            network_mode=2,
            logical_region_id=default.regions[0].id,
            vpc_id=default_network.id)
        default_instance_cluster_attachment = alicloud.edas.InstanceClusterAttachment("default",
            cluster_id=default_cluster.id,
            instance_ids=[default_instance.id])
        default_application = alicloud.edas.Application("default",
            application_name=name,
            cluster_id=default_cluster.id,
            package_type="JAR")
        default_deploy_group = alicloud.edas.DeployGroup("default",
            app_id=default_application.id,
            group_name=name)
        default_application_deployment = alicloud.edas.ApplicationDeployment("default",
            app_id=default_application.id,
            group_id="all",
            war_url="http://edas-sz.oss-cn-shenzhen.aliyuncs.com/prod/demo/SPRING_CLOUD_CONSUMER.jar")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the application that you want to deploy.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationDeploymentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Deploys applications on EDAS, see [What is EDAS Application Deployment](https://www.alibabacloud.com/help/en/edas/developer-reference/api-edas-2017-08-01-deployapplication).

        > **NOTE:** Available since v1.82.0.

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
        default = alicloud.get_regions(current=True)
        default_get_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        default_get_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            owners="system")
        default_get_instance_types = alicloud.ecs.get_instance_types(availability_zone=default_get_zones.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        default_network = alicloud.vpc.Network("default",
            vpc_name=name,
            cidr_block="10.4.0.0/16")
        default_switch = alicloud.vpc.Switch("default",
            vswitch_name=name,
            cidr_block="10.4.0.0/24",
            vpc_id=default_network.id,
            zone_id=default_get_zones.zones[0].id)
        default_security_group = alicloud.ecs.SecurityGroup("default", vpc_id=default_network.id)
        default_instance = alicloud.ecs.Instance("default",
            availability_zone=default_get_zones.zones[0].id,
            instance_name=name,
            image_id=default_get_images.images[0].id,
            instance_type=default_get_instance_types.instance_types[0].id,
            security_groups=[default_security_group.id],
            vswitch_id=default_switch.id)
        default_cluster = alicloud.edas.Cluster("default",
            cluster_name=name,
            cluster_type=2,
            network_mode=2,
            logical_region_id=default.regions[0].id,
            vpc_id=default_network.id)
        default_instance_cluster_attachment = alicloud.edas.InstanceClusterAttachment("default",
            cluster_id=default_cluster.id,
            instance_ids=[default_instance.id])
        default_application = alicloud.edas.Application("default",
            application_name=name,
            cluster_id=default_cluster.id,
            package_type="JAR")
        default_deploy_group = alicloud.edas.DeployGroup("default",
            app_id=default_application.id,
            group_name=name)
        default_application_deployment = alicloud.edas.ApplicationDeployment("default",
            app_id=default_application.id,
            group_id="all",
            war_url="http://edas-sz.oss-cn-shenzhen.aliyuncs.com/prod/demo/SPRING_CLOUD_CONSUMER.jar")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ApplicationDeploymentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationDeploymentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 package_version: Optional[pulumi.Input[str]] = None,
                 war_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationDeploymentArgs.__new__(ApplicationDeploymentArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["package_version"] = package_version
            if war_url is None and not opts.urn:
                raise TypeError("Missing required property 'war_url'")
            __props__.__dict__["war_url"] = war_url
            __props__.__dict__["last_package_version"] = None
        super(ApplicationDeployment, __self__).__init__(
            'alicloud:edas/applicationDeployment:ApplicationDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            last_package_version: Optional[pulumi.Input[str]] = None,
            package_version: Optional[pulumi.Input[str]] = None,
            war_url: Optional[pulumi.Input[str]] = None) -> 'ApplicationDeployment':
        """
        Get an existing ApplicationDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the application that you want to deploy.
        :param pulumi.Input[str] group_id: The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        :param pulumi.Input[str] last_package_version: Last package version deployed.
        :param pulumi.Input[str] package_version: The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        :param pulumi.Input[str] war_url: The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationDeploymentState.__new__(_ApplicationDeploymentState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["last_package_version"] = last_package_version
        __props__.__dict__["package_version"] = package_version
        __props__.__dict__["war_url"] = war_url
        return ApplicationDeployment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        The ID of the application that you want to deploy.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="lastPackageVersion")
    def last_package_version(self) -> pulumi.Output[str]:
        """
        Last package version deployed.
        """
        return pulumi.get(self, "last_package_version")

    @property
    @pulumi.getter(name="packageVersion")
    def package_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.
        """
        return pulumi.get(self, "package_version")

    @property
    @pulumi.getter(name="warUrl")
    def war_url(self) -> pulumi.Output[str]:
        """
        The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url.
        """
        return pulumi.get(self, "war_url")

