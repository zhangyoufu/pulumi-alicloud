// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the api groups of the current Alibaba Cloud user.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const dataApigatway = pulumi.output(alicloud.apigateway.getGroups({
 *     outputFile: "outgroups",
 * }, { async: true }));
 * 
 * export const firstGroupId = dataApigatway.groups[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/api_gateway_groups.html.markdown.
 */
export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:apigateway/getGroups:getGroups", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    /**
     * A list of api group IDs. 
     */
    readonly ids?: string[];
    /**
     * A regex string to filter api gateway groups by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    /**
     * A list of api groups. Each element contains the following attributes:
     */
    readonly groups: outputs.apigateway.GetGroupsGroup[];
    /**
     * A list of api group IDs. 
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of api group names. 
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
