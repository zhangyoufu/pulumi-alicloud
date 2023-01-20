// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BackupPlanArgs, BackupPlanState } from "./backupPlan";
export type BackupPlan = import("./backupPlan").BackupPlan;
export const BackupPlan: typeof import("./backupPlan").BackupPlan = null as any;
utilities.lazyLoad(exports, ["BackupPlan"], () => require("./backupPlan"));

export { GetBackupPlansArgs, GetBackupPlansResult, GetBackupPlansOutputArgs } from "./getBackupPlans";
export const getBackupPlans: typeof import("./getBackupPlans").getBackupPlans = null as any;
export const getBackupPlansOutput: typeof import("./getBackupPlans").getBackupPlansOutput = null as any;
utilities.lazyLoad(exports, ["getBackupPlans","getBackupPlansOutput"], () => require("./getBackupPlans"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:dbs/backupPlan:BackupPlan":
                return new BackupPlan(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "dbs/backupPlan", _module)
