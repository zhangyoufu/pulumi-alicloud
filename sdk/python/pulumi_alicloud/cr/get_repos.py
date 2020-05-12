# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetReposResult:
    """
    A collection of values returned by getRepos.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, name_regex=None, names=None, namespace=None, output_file=None, repos=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        __self__.enable_details = enable_details
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        """
        A list of matched Container Registry Repositories. Its element is set to `names`.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of repository names.
        """
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        __self__.namespace = namespace
        """
        Name of container registry namespace where repo is located.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
        if repos and not isinstance(repos, list):
            raise TypeError("Expected argument 'repos' to be a list")
        __self__.repos = repos
        """
        A list of matched Container Registry Namespaces. Each element contains the following attributes:
        """
class AwaitableGetReposResult(GetReposResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReposResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            namespace=self.namespace,
            output_file=self.output_file,
            repos=self.repos)

def get_repos(enable_details=None,name_regex=None,namespace=None,output_file=None,opts=None):
    """
    This data source provides a list Container Registry repositories on Alibaba Cloud.

    > **NOTE:** Available in v1.35.0+

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    my_repos = alicloud.cr.get_repos(name_regex="my-repos",
        output_file="my-repo-json")
    pulumi.export("output", my_repos.repos)
    ```



    :param bool enable_details: Boolean, false by default, only repository attributes are exported. Set to true if domain list and tags belong to this repository are needed. See `tags` in attributes.
    :param str name_regex: A regex string to filter results by repository name.
    :param str namespace: Name of container registry namespace where the repositories are located in.
    """
    __args__ = dict()


    __args__['enableDetails'] = enable_details
    __args__['nameRegex'] = name_regex
    __args__['namespace'] = namespace
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:cr/getRepos:getRepos', __args__, opts=opts).value

    return AwaitableGetReposResult(
        enable_details=__ret__.get('enableDetails'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        names=__ret__.get('names'),
        namespace=__ret__.get('namespace'),
        output_file=__ret__.get('outputFile'),
        repos=__ret__.get('repos'))
