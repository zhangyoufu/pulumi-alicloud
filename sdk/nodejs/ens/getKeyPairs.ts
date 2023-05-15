// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ens Key Pairs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const nameRegex = alicloud.ens.getKeyPairs({
 *     version: "example_value",
 *     nameRegex: "^my-KeyPair",
 * });
 * export const ensKeyPairId1 = nameRegex.then(nameRegex => nameRegex.pairs?.[0]?.id);
 * ```
 */
export function getKeyPairs(args: GetKeyPairsArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyPairsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ens/getKeyPairs:getKeyPairs", {
        "keyPairName": args.keyPairName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeyPairs.
 */
export interface GetKeyPairsArgs {
    /**
     * The name of the key pair.
     */
    keyPairName?: string;
    /**
     * A regex string to filter results by Key Pair name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
    /**
     * The version number.
     */
    version: string;
}

/**
 * A collection of values returned by getKeyPairs.
 */
export interface GetKeyPairsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly keyPairName?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly pairs: outputs.ens.GetKeyPairsPair[];
    readonly version: string;
}
/**
 * This data source provides the Ens Key Pairs of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const nameRegex = alicloud.ens.getKeyPairs({
 *     version: "example_value",
 *     nameRegex: "^my-KeyPair",
 * });
 * export const ensKeyPairId1 = nameRegex.then(nameRegex => nameRegex.pairs?.[0]?.id);
 * ```
 */
export function getKeyPairsOutput(args: GetKeyPairsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeyPairsResult> {
    return pulumi.output(args).apply((a: any) => getKeyPairs(a, opts))
}

/**
 * A collection of arguments for invoking getKeyPairs.
 */
export interface GetKeyPairsOutputArgs {
    /**
     * The name of the key pair.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Key Pair name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The version number.
     */
    version: pulumi.Input<string>;
}
