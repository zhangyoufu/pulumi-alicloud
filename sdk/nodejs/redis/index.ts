// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { TairInstanceArgs, TairInstanceState } from "./tairInstance";
export type TairInstance = import("./tairInstance").TairInstance;
export const TairInstance: typeof import("./tairInstance").TairInstance = null as any;
utilities.lazyLoad(exports, ["TairInstance"], () => require("./tairInstance"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:redis/tairInstance:TairInstance":
                return new TairInstance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "redis/tairInstance", _module)
