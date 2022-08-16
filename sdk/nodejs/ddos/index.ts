// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./basicDefenseThreshold";
export * from "./bgpIp";
export * from "./ddosBgpInstance";
export * from "./ddosCooInstance";
export * from "./domainResource";
export * from "./getDdosBgpInstances";
export * from "./getDdosBgpIps";
export * from "./getDdosCooDomainResources";
export * from "./getDdosCooInstances";
export * from "./getDdosCooPorts";
export * from "./port";
export * from "./schedulerRule";

// Import resources to register:
import { BasicDefenseThreshold } from "./basicDefenseThreshold";
import { BgpIp } from "./bgpIp";
import { DdosBgpInstance } from "./ddosBgpInstance";
import { DdosCooInstance } from "./ddosCooInstance";
import { DomainResource } from "./domainResource";
import { Port } from "./port";
import { SchedulerRule } from "./schedulerRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:ddos/basicDefenseThreshold:BasicDefenseThreshold":
                return new BasicDefenseThreshold(name, <any>undefined, { urn })
            case "alicloud:ddos/bgpIp:BgpIp":
                return new BgpIp(name, <any>undefined, { urn })
            case "alicloud:ddos/ddosBgpInstance:DdosBgpInstance":
                return new DdosBgpInstance(name, <any>undefined, { urn })
            case "alicloud:ddos/ddosCooInstance:DdosCooInstance":
                return new DdosCooInstance(name, <any>undefined, { urn })
            case "alicloud:ddos/domainResource:DomainResource":
                return new DomainResource(name, <any>undefined, { urn })
            case "alicloud:ddos/port:Port":
                return new Port(name, <any>undefined, { urn })
            case "alicloud:ddos/schedulerRule:SchedulerRule":
                return new SchedulerRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "ddos/basicDefenseThreshold", _module)
pulumi.runtime.registerResourceModule("alicloud", "ddos/bgpIp", _module)
pulumi.runtime.registerResourceModule("alicloud", "ddos/ddosBgpInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "ddos/ddosCooInstance", _module)
pulumi.runtime.registerResourceModule("alicloud", "ddos/domainResource", _module)
pulumi.runtime.registerResourceModule("alicloud", "ddos/port", _module)
pulumi.runtime.registerResourceModule("alicloud", "ddos/schedulerRule", _module)
