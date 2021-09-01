// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./config";
export * from "./getConfigs";
export * from "./getRules";
export * from "./rule";

// Import resources to register:
import { Config } from "./config";
import { Rule } from "./rule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:sddp/config:Config":
                return new Config(name, <any>undefined, { urn })
            case "alicloud:sddp/rule:Rule":
                return new Rule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "sddp/config", _module)
pulumi.runtime.registerResourceModule("alicloud", "sddp/rule", _module)
