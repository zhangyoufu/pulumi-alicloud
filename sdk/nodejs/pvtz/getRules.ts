// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the PrivateZone Rules of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.143.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.pvtz.getRules({});
 * export const pvtzRuleId1 = ids.then(ids => ids.rules?.[0]?.id);
 * const nameRegex = alicloud.pvtz.getRules({
 *     nameRegex: "^my-Rule",
 * });
 * export const pvtzRuleId2 = nameRegex.then(nameRegex => nameRegex.rules?.[0]?.id);
 * ```
 */
export function getRules(args?: GetRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetRulesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:pvtz/getRules:getRules", {
        "endpointId": args.endpointId,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getRules.
 */
export interface GetRulesArgs {
    /**
     * The ID of the Endpoint.
     */
    endpointId?: string;
    /**
     * A list of Rule IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Rule name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getRules.
 */
export interface GetRulesResult {
    readonly endpointId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly rules: outputs.pvtz.GetRulesRule[];
}
/**
 * This data source provides the PrivateZone Rules of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.143.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.pvtz.getRules({});
 * export const pvtzRuleId1 = ids.then(ids => ids.rules?.[0]?.id);
 * const nameRegex = alicloud.pvtz.getRules({
 *     nameRegex: "^my-Rule",
 * });
 * export const pvtzRuleId2 = nameRegex.then(nameRegex => nameRegex.rules?.[0]?.id);
 * ```
 */
export function getRulesOutput(args?: GetRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRulesResult> {
    return pulumi.output(args).apply((a: any) => getRules(a, opts))
}

/**
 * A collection of arguments for invoking getRules.
 */
export interface GetRulesOutputArgs {
    /**
     * The ID of the Endpoint.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * A list of Rule IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Rule name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
