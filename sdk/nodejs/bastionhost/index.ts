// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetHostAccountsArgs, GetHostAccountsResult, GetHostAccountsOutputArgs } from "./getHostAccounts";
export const getHostAccounts: typeof import("./getHostAccounts").getHostAccounts = null as any;
export const getHostAccountsOutput: typeof import("./getHostAccounts").getHostAccountsOutput = null as any;
utilities.lazyLoad(exports, ["getHostAccounts","getHostAccountsOutput"], () => require("./getHostAccounts"));

export { GetHostGroupsArgs, GetHostGroupsResult, GetHostGroupsOutputArgs } from "./getHostGroups";
export const getHostGroups: typeof import("./getHostGroups").getHostGroups = null as any;
export const getHostGroupsOutput: typeof import("./getHostGroups").getHostGroupsOutput = null as any;
utilities.lazyLoad(exports, ["getHostGroups","getHostGroupsOutput"], () => require("./getHostGroups"));

export { GetHostShareKeysArgs, GetHostShareKeysResult, GetHostShareKeysOutputArgs } from "./getHostShareKeys";
export const getHostShareKeys: typeof import("./getHostShareKeys").getHostShareKeys = null as any;
export const getHostShareKeysOutput: typeof import("./getHostShareKeys").getHostShareKeysOutput = null as any;
utilities.lazyLoad(exports, ["getHostShareKeys","getHostShareKeysOutput"], () => require("./getHostShareKeys"));

export { GetHostsArgs, GetHostsResult, GetHostsOutputArgs } from "./getHosts";
export const getHosts: typeof import("./getHosts").getHosts = null as any;
export const getHostsOutput: typeof import("./getHosts").getHostsOutput = null as any;
utilities.lazyLoad(exports, ["getHosts","getHostsOutput"], () => require("./getHosts"));

export { GetInstancesArgs, GetInstancesResult, GetInstancesOutputArgs } from "./getInstances";
export const getInstances: typeof import("./getInstances").getInstances = null as any;
export const getInstancesOutput: typeof import("./getInstances").getInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getInstances","getInstancesOutput"], () => require("./getInstances"));

export { GetUserGroupsArgs, GetUserGroupsResult, GetUserGroupsOutputArgs } from "./getUserGroups";
export const getUserGroups: typeof import("./getUserGroups").getUserGroups = null as any;
export const getUserGroupsOutput: typeof import("./getUserGroups").getUserGroupsOutput = null as any;
utilities.lazyLoad(exports, ["getUserGroups","getUserGroupsOutput"], () => require("./getUserGroups"));

export { GetUsersArgs, GetUsersResult, GetUsersOutputArgs } from "./getUsers";
export const getUsers: typeof import("./getUsers").getUsers = null as any;
export const getUsersOutput: typeof import("./getUsers").getUsersOutput = null as any;
utilities.lazyLoad(exports, ["getUsers","getUsersOutput"], () => require("./getUsers"));

export { HostArgs, HostState } from "./host";
export type Host = import("./host").Host;
export const Host: typeof import("./host").Host = null as any;
utilities.lazyLoad(exports, ["Host"], () => require("./host"));

export { HostAccountArgs, HostAccountState } from "./hostAccount";
export type HostAccount = import("./hostAccount").HostAccount;
export const HostAccount: typeof import("./hostAccount").HostAccount = null as any;
utilities.lazyLoad(exports, ["HostAccount"], () => require("./hostAccount"));

export { HostAccountShareKeyAttachmentArgs, HostAccountShareKeyAttachmentState } from "./hostAccountShareKeyAttachment";
export type HostAccountShareKeyAttachment = import("./hostAccountShareKeyAttachment").HostAccountShareKeyAttachment;
export const HostAccountShareKeyAttachment: typeof import("./hostAccountShareKeyAttachment").HostAccountShareKeyAttachment = null as any;
utilities.lazyLoad(exports, ["HostAccountShareKeyAttachment"], () => require("./hostAccountShareKeyAttachment"));

export { HostAccountUserAttachmentArgs, HostAccountUserAttachmentState } from "./hostAccountUserAttachment";
export type HostAccountUserAttachment = import("./hostAccountUserAttachment").HostAccountUserAttachment;
export const HostAccountUserAttachment: typeof import("./hostAccountUserAttachment").HostAccountUserAttachment = null as any;
utilities.lazyLoad(exports, ["HostAccountUserAttachment"], () => require("./hostAccountUserAttachment"));

export { HostAccountUserGroupAttachmentArgs, HostAccountUserGroupAttachmentState } from "./hostAccountUserGroupAttachment";
export type HostAccountUserGroupAttachment = import("./hostAccountUserGroupAttachment").HostAccountUserGroupAttachment;
export const HostAccountUserGroupAttachment: typeof import("./hostAccountUserGroupAttachment").HostAccountUserGroupAttachment = null as any;
utilities.lazyLoad(exports, ["HostAccountUserGroupAttachment"], () => require("./hostAccountUserGroupAttachment"));

export { HostAttachmentArgs, HostAttachmentState } from "./hostAttachment";
export type HostAttachment = import("./hostAttachment").HostAttachment;
export const HostAttachment: typeof import("./hostAttachment").HostAttachment = null as any;
utilities.lazyLoad(exports, ["HostAttachment"], () => require("./hostAttachment"));

export { HostGroupArgs, HostGroupState } from "./hostGroup";
export type HostGroup = import("./hostGroup").HostGroup;
export const HostGroup: typeof import("./hostGroup").HostGroup = null as any;
utilities.lazyLoad(exports, ["HostGroup"], () => require("./hostGroup"));

export { HostGroupAccountUserAttachmentArgs, HostGroupAccountUserAttachmentState } from "./hostGroupAccountUserAttachment";
export type HostGroupAccountUserAttachment = import("./hostGroupAccountUserAttachment").HostGroupAccountUserAttachment;
export const HostGroupAccountUserAttachment: typeof import("./hostGroupAccountUserAttachment").HostGroupAccountUserAttachment = null as any;
utilities.lazyLoad(exports, ["HostGroupAccountUserAttachment"], () => require("./hostGroupAccountUserAttachment"));

export { HostGroupAccountUserGroupAttachmentArgs, HostGroupAccountUserGroupAttachmentState } from "./hostGroupAccountUserGroupAttachment";
export type HostGroupAccountUserGroupAttachment = import("./hostGroupAccountUserGroupAttachment").HostGroupAccountUserGroupAttachment;
export const HostGroupAccountUserGroupAttachment: typeof import("./hostGroupAccountUserGroupAttachment").HostGroupAccountUserGroupAttachment = null as any;
utilities.lazyLoad(exports, ["HostGroupAccountUserGroupAttachment"], () => require("./hostGroupAccountUserGroupAttachment"));

export { HostShareKeyArgs, HostShareKeyState } from "./hostShareKey";
export type HostShareKey = import("./hostShareKey").HostShareKey;
export const HostShareKey: typeof import("./hostShareKey").HostShareKey = null as any;
utilities.lazyLoad(exports, ["HostShareKey"], () => require("./hostShareKey"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));

export { UserAttachmentArgs, UserAttachmentState } from "./userAttachment";
export type UserAttachment = import("./userAttachment").UserAttachment;
export const UserAttachment: typeof import("./userAttachment").UserAttachment = null as any;
utilities.lazyLoad(exports, ["UserAttachment"], () => require("./userAttachment"));

export { UserGroupArgs, UserGroupState } from "./userGroup";
export type UserGroup = import("./userGroup").UserGroup;
export const UserGroup: typeof import("./userGroup").UserGroup = null as any;
utilities.lazyLoad(exports, ["UserGroup"], () => require("./userGroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:bastionhost/host:Host":
                return new Host(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostAccount:HostAccount":
                return new HostAccount(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostAccountShareKeyAttachment:HostAccountShareKeyAttachment":
                return new HostAccountShareKeyAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostAccountUserAttachment:HostAccountUserAttachment":
                return new HostAccountUserAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment":
                return new HostAccountUserGroupAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostAttachment:HostAttachment":
                return new HostAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostGroup:HostGroup":
                return new HostGroup(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostGroupAccountUserAttachment:HostGroupAccountUserAttachment":
                return new HostGroupAccountUserAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostGroupAccountUserGroupAttachment:HostGroupAccountUserGroupAttachment":
                return new HostGroupAccountUserGroupAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/hostShareKey:HostShareKey":
                return new HostShareKey(name, <any>undefined, { urn })
            case "alicloud:bastionhost/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "alicloud:bastionhost/user:User":
                return new User(name, <any>undefined, { urn })
            case "alicloud:bastionhost/userAttachment:UserAttachment":
                return new UserAttachment(name, <any>undefined, { urn })
            case "alicloud:bastionhost/userGroup:UserGroup":
                return new UserGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/host", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostAccount", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostAccountShareKeyAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostAccountUserAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostAccountUserGroupAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostGroupAccountUserAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostGroupAccountUserGroupAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/hostShareKey", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/instance", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/user", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/userAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "bastionhost/userGroup", _module)
