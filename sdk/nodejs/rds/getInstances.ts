// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The `alicloud.rds.getInstances` data source provides a collection of RDS instances available in Alibaba Cloud account.
 * Filters support regular expression for the instance name, searches by tags, and other filters which are listed below.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const dbInstancesDs = pulumi.output(alicloud.rds.getInstances({
 *     nameRegex: "data-\\d+",
 *     status: "Running",
 *     tags: {
 *         size: "tiny",
 *         type: "database",
 *     },
 * }, { async: true }));
 * 
 * export const firstDbInstanceId = dbInstancesDs.instances[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/db_instances.html.markdown.
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:rds/getInstances:getInstances", {
        "connectionMode": args.connectionMode,
        "dbType": args.dbType,
        "engine": args.engine,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "vswitchId": args.vswitchId,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * `Standard` for standard access mode and `Safe` for high security access mode.
     */
    readonly connectionMode?: string;
    /**
     * `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
     */
    readonly dbType?: string;
    /**
     * Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
     */
    readonly engine?: string;
    /**
     * A list of RDS instance IDs. 
     */
    readonly ids?: string[];
    /**
     * A regex string to filter results by instance name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Status of the instance.
     */
    readonly status?: string;
    /**
     * A map of tags assigned to the DB instances. 
     * Note: Before 1.60.0, the value's format is a `json` string which including `TagKey` and `TagValue`. `TagKey` cannot be null, and `TagValue` can be empty. Format example `"{\"key1\":\"value1\"}"`
     */
    readonly tags?: {[key: string]: any};
    /**
     * Used to retrieve instances belong to specified VPC.
     */
    readonly vpcId?: string;
    /**
     * Used to retrieve instances belong to specified `vswitch` resources.
     */
    readonly vswitchId?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    /**
     * `Standard` for standard access mode and `Safe` for high security access mode.
     */
    readonly connectionMode?: string;
    /**
     * `Primary` for primary instance, `Readonly` for read-only instance, `Guard` for disaster recovery instance, and `Temp` for temporary instance.
     */
    readonly dbType?: string;
    /**
     * Database type. Options are `MySQL`, `SQLServer`, `PostgreSQL` and `PPAS`. If no value is specified, all types are returned.
     */
    readonly engine?: string;
    /**
     * A list of RDS instance IDs. 
     */
    readonly ids: string[];
    /**
     * A list of RDS instances. Each element contains the following attributes:
     */
    readonly instances: outputs.rds.GetInstancesInstance[];
    readonly nameRegex?: string;
    /**
     * A list of RDS instance names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * Status of the instance.
     */
    readonly status?: string;
    readonly tags?: {[key: string]: any};
    /**
     * ID of the VPC the instance belongs to.
     */
    readonly vpcId?: string;
    /**
     * ID of the VSwitch the instance belongs to.
     */
    readonly vswitchId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
