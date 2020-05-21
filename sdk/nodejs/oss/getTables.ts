// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the ots tables of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.40.0+.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const tablesDs = pulumi.output(alicloud.oss.getTables({
 *     instanceName: "sample-instance",
 *     nameRegex: "sample-table",
 *     outputFile: "tables.txt",
 * }, { async: true }));
 *
 * export const firstTableId = tablesDs.tables[0].id;
 * ```
 */
export function getTables(args: GetTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetTablesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:oss/getTables:getTables", {
        "ids": args.ids,
        "instanceName": args.instanceName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getTables.
 */
export interface GetTablesArgs {
    /**
     * A list of table IDs.
     */
    readonly ids?: string[];
    /**
     * The name of OTS instance.
     */
    readonly instanceName: string;
    /**
     * A regex string to filter results by table name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getTables.
 */
export interface GetTablesResult {
    /**
     * A list of table IDs.
     */
    readonly ids: string[];
    /**
     * The OTS instance name.
     */
    readonly instanceName: string;
    readonly nameRegex?: string;
    /**
     * A list of table names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of tables. Each element contains the following attributes:
     */
    readonly tables: outputs.oss.GetTablesTable[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
