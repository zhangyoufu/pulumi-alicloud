// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DiskArgs, DiskState } from "./disk";
export type Disk = import("./disk").Disk;
export const Disk: typeof import("./disk").Disk = null as any;
utilities.lazyLoad(exports, ["Disk"], () => require("./disk"));

export { DiskInstanceAttachmentArgs, DiskInstanceAttachmentState } from "./diskInstanceAttachment";
export type DiskInstanceAttachment = import("./diskInstanceAttachment").DiskInstanceAttachment;
export const DiskInstanceAttachment: typeof import("./diskInstanceAttachment").DiskInstanceAttachment = null as any;
utilities.lazyLoad(exports, ["DiskInstanceAttachment"], () => require("./diskInstanceAttachment"));

export { EipArgs, EipState } from "./eip";
export type Eip = import("./eip").Eip;
export const Eip: typeof import("./eip").Eip = null as any;
utilities.lazyLoad(exports, ["Eip"], () => require("./eip"));

export { EipInstanceAttachmentArgs, EipInstanceAttachmentState } from "./eipInstanceAttachment";
export type EipInstanceAttachment = import("./eipInstanceAttachment").EipInstanceAttachment;
export const EipInstanceAttachment: typeof import("./eipInstanceAttachment").EipInstanceAttachment = null as any;
utilities.lazyLoad(exports, ["EipInstanceAttachment"], () => require("./eipInstanceAttachment"));

export { GetKeyPairsArgs, GetKeyPairsResult, GetKeyPairsOutputArgs } from "./getKeyPairs";
export const getKeyPairs: typeof import("./getKeyPairs").getKeyPairs = null as any;
export const getKeyPairsOutput: typeof import("./getKeyPairs").getKeyPairsOutput = null as any;
utilities.lazyLoad(exports, ["getKeyPairs","getKeyPairsOutput"], () => require("./getKeyPairs"));

export { ImageArgs, ImageState } from "./image";
export type Image = import("./image").Image;
export const Image: typeof import("./image").Image = null as any;
utilities.lazyLoad(exports, ["Image"], () => require("./image"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstanceSecurityGroupAttachmentArgs, InstanceSecurityGroupAttachmentState } from "./instanceSecurityGroupAttachment";
export type InstanceSecurityGroupAttachment = import("./instanceSecurityGroupAttachment").InstanceSecurityGroupAttachment;
export const InstanceSecurityGroupAttachment: typeof import("./instanceSecurityGroupAttachment").InstanceSecurityGroupAttachment = null as any;
utilities.lazyLoad(exports, ["InstanceSecurityGroupAttachment"], () => require("./instanceSecurityGroupAttachment"));

export { KeyPairArgs, KeyPairState } from "./keyPair";
export type KeyPair = import("./keyPair").KeyPair;
export const KeyPair: typeof import("./keyPair").KeyPair = null as any;
utilities.lazyLoad(exports, ["KeyPair"], () => require("./keyPair"));

export { LoadBalancerArgs, LoadBalancerState } from "./loadBalancer";
export type LoadBalancer = import("./loadBalancer").LoadBalancer;
export const LoadBalancer: typeof import("./loadBalancer").LoadBalancer = null as any;
utilities.lazyLoad(exports, ["LoadBalancer"], () => require("./loadBalancer"));

export { NatGatewayArgs, NatGatewayState } from "./natGateway";
export type NatGateway = import("./natGateway").NatGateway;
export const NatGateway: typeof import("./natGateway").NatGateway = null as any;
utilities.lazyLoad(exports, ["NatGateway"], () => require("./natGateway"));

export { NetworkArgs, NetworkState } from "./network";
export type Network = import("./network").Network;
export const Network: typeof import("./network").Network = null as any;
utilities.lazyLoad(exports, ["Network"], () => require("./network"));

export { SecurityGroupArgs, SecurityGroupState } from "./securityGroup";
export type SecurityGroup = import("./securityGroup").SecurityGroup;
export const SecurityGroup: typeof import("./securityGroup").SecurityGroup = null as any;
utilities.lazyLoad(exports, ["SecurityGroup"], () => require("./securityGroup"));

export { SnapshotArgs, SnapshotState } from "./snapshot";
export type Snapshot = import("./snapshot").Snapshot;
export const Snapshot: typeof import("./snapshot").Snapshot = null as any;
utilities.lazyLoad(exports, ["Snapshot"], () => require("./snapshot"));

export { VswitchArgs, VswitchState } from "./vswitch";
export type Vswitch = import("./vswitch").Vswitch;
export const Vswitch: typeof import("./vswitch").Vswitch = null as any;
utilities.lazyLoad(exports, ["Vswitch"], () => require("./vswitch"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:ens/disk:Disk":
                return new Disk(name, <any>undefined, { urn })
            case "alicloud:ens/diskInstanceAttachment:DiskInstanceAttachment":
                return new DiskInstanceAttachment(name, <any>undefined, { urn })
            case "alicloud:ens/eip:Eip":
                return new Eip(name, <any>undefined, { urn })
            case "alicloud:ens/eipInstanceAttachment:EipInstanceAttachment":
                return new EipInstanceAttachment(name, <any>undefined, { urn })
            case "alicloud:ens/image:Image":
                return new Image(name, <any>undefined, { urn })
            case "alicloud:ens/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "alicloud:ens/instanceSecurityGroupAttachment:InstanceSecurityGroupAttachment":
                return new InstanceSecurityGroupAttachment(name, <any>undefined, { urn })
            case "alicloud:ens/keyPair:KeyPair":
                return new KeyPair(name, <any>undefined, { urn })
            case "alicloud:ens/loadBalancer:LoadBalancer":
                return new LoadBalancer(name, <any>undefined, { urn })
            case "alicloud:ens/natGateway:NatGateway":
                return new NatGateway(name, <any>undefined, { urn })
            case "alicloud:ens/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "alicloud:ens/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "alicloud:ens/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "alicloud:ens/vswitch:Vswitch":
                return new Vswitch(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "ens/disk", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/diskInstanceAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/eip", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/eipInstanceAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/image", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/instance", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/instanceSecurityGroupAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/keyPair", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/loadBalancer", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/natGateway", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/network", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/securityGroup", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/snapshot", _module)
pulumi.runtime.registerResourceModule("alicloud", "ens/vswitch", _module)
