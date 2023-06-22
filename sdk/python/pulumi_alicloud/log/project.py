# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[str] description: Description of the log project.
        :param pulumi.Input[str] name: The name of the log project. It is the only in one Alicloud account.
        :param pulumi.Input[str] policy: Log project policy, used to set a policy for a project.
        :param pulumi.Input[Mapping[str, Any]] tags: Log project tags.
               - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
               - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the log project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the log project. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Log project policy, used to set a policy for a project.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Log project tags.
        - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ProjectState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Project resources.
        :param pulumi.Input[str] description: Description of the log project.
        :param pulumi.Input[str] name: The name of the log project. It is the only in one Alicloud account.
        :param pulumi.Input[str] policy: Log project policy, used to set a policy for a project.
        :param pulumi.Input[Mapping[str, Any]] tags: Log project tags.
               - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
               - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the log project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the log project. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Log project policy, used to set a policy for a project.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Log project tags.
        - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        The project is the resource management unit in Log Service and is used to isolate and control resources.
        You can manage all the logs and the related log sources of an application by using projects. [Refer to details](https://www.alibabacloud.com/help/doc-detail/48873.htm).

        > **NOTE:** Available since v1.9.5.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default = random.RandomInteger("default",
            max=99999,
            min=10000)
        example = alicloud.log.Project("example",
            description="terraform-example",
            tags={
                "Created": "TF",
                "For": "example",
            })
        ```

        Project With Policy Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default = random.RandomInteger("default",
            max=99999,
            min=10000)
        example_policy = alicloud.log.Project("examplePolicy",
            description="terraform-example",
            policy=\"\"\"{
          "Statement": [
            {
              "Action": [
                "log:PostLogStoreLogs"
              ],
              "Condition": {
                "StringNotLike": {
                  "acs:SourceVpc": [
                    "vpc-*"
                  ]
                }
              },
              "Effect": "Deny",
              "Resource": "acs:log:*:*:project/tf-log/*"
            }
          ],
          "Version": "1"
        }

        \"\"\")
        ```
        ## Module Support

        You can use the existing sls module
        to create SLS project, store and store index one-click, like ECS instances.

        ## Import

        Log project can be imported using the id or name, e.g.

        ```sh
         $ pulumi import alicloud:log/project:Project example tf-log
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the log project.
        :param pulumi.Input[str] name: The name of the log project. It is the only in one Alicloud account.
        :param pulumi.Input[str] policy: Log project policy, used to set a policy for a project.
        :param pulumi.Input[Mapping[str, Any]] tags: Log project tags.
               - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
               - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProjectArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The project is the resource management unit in Log Service and is used to isolate and control resources.
        You can manage all the logs and the related log sources of an application by using projects. [Refer to details](https://www.alibabacloud.com/help/doc-detail/48873.htm).

        > **NOTE:** Available since v1.9.5.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default = random.RandomInteger("default",
            max=99999,
            min=10000)
        example = alicloud.log.Project("example",
            description="terraform-example",
            tags={
                "Created": "TF",
                "For": "example",
            })
        ```

        Project With Policy Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud
        import pulumi_random as random

        default = random.RandomInteger("default",
            max=99999,
            min=10000)
        example_policy = alicloud.log.Project("examplePolicy",
            description="terraform-example",
            policy=\"\"\"{
          "Statement": [
            {
              "Action": [
                "log:PostLogStoreLogs"
              ],
              "Condition": {
                "StringNotLike": {
                  "acs:SourceVpc": [
                    "vpc-*"
                  ]
                }
              },
              "Effect": "Deny",
              "Resource": "acs:log:*:*:project/tf-log/*"
            }
          ],
          "Version": "1"
        }

        \"\"\")
        ```
        ## Module Support

        You can use the existing sls module
        to create SLS project, store and store index one-click, like ECS instances.

        ## Import

        Log project can be imported using the id or name, e.g.

        ```sh
         $ pulumi import alicloud:log/project:Project example tf-log
        ```

        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["policy"] = policy
            __props__.__dict__["tags"] = tags
        super(Project, __self__).__init__(
            'alicloud:log/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the log project.
        :param pulumi.Input[str] name: The name of the log project. It is the only in one Alicloud account.
        :param pulumi.Input[str] policy: Log project policy, used to set a policy for a project.
        :param pulumi.Input[Mapping[str, Any]] tags: Log project tags.
               - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
               - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectState.__new__(_ProjectState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["policy"] = policy
        __props__.__dict__["tags"] = tags
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the log project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the log project. It is the only in one Alicloud account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[Optional[str]]:
        """
        Log project policy, used to set a policy for a project.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Log project tags.
        - Key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://".
        """
        return pulumi.get(self, "tags")

