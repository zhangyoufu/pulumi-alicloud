# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetHandshakesResult:
    """
    A collection of values returned by getHandshakes.
    """
    def __init__(__self__, handshakes=None, id=None, ids=None, output_file=None):
        if handshakes and not isinstance(handshakes, list):
            raise TypeError("Expected argument 'handshakes' to be a list")
        __self__.handshakes = handshakes
        """
        A list of Resource Manager Handshakes. Each element contains the following attributes:
        """
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
        A list of Resource Manager Handshake IDs.
        """
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        __self__.output_file = output_file
class AwaitableGetHandshakesResult(GetHandshakesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHandshakesResult(
            handshakes=self.handshakes,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file)

def get_handshakes(ids=None,output_file=None,opts=None):
    """
    This data source provides the Resource Manager Handshakes of the current Alibaba Cloud user.

    > **NOTE:**  Available in 1.86.0+.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.resourcemanager.get_handshakes()
    pulumi.export("firstHandshakeId", example.handshakes[0]["id"])
    ```



    :param list ids: A list of Resource Manager Handshake IDs.
    """
    __args__ = dict()


    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('alicloud:resourcemanager/getHandshakes:getHandshakes', __args__, opts=opts).value

    return AwaitableGetHandshakesResult(
        handshakes=__ret__.get('handshakes'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        output_file=__ret__.get('outputFile'))
