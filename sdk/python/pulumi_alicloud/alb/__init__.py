# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .acl import *
from .get_acls import *
from .get_listeners import *
from .get_load_balancers import *
from .get_rules import *
from .get_security_policies import *
from .get_server_groups import *
from .get_zones import *
from .listener import *
from .load_balancer import *
from .rule import *
from .security_policy import *
from .server_group import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "alicloud:alb/acl:Acl":
                return Acl(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:alb/listener:Listener":
                return Listener(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:alb/loadBalancer:LoadBalancer":
                return LoadBalancer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:alb/rule:Rule":
                return Rule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:alb/securityPolicy:SecurityPolicy":
                return SecurityPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "alicloud:alb/serverGroup:ServerGroup":
                return ServerGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("alicloud", "alb/acl", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "alb/listener", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "alb/loadBalancer", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "alb/rule", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "alb/securityPolicy", _module_instance)
    pulumi.runtime.register_resource_module("alicloud", "alb/serverGroup", _module_instance)

_register_module()
