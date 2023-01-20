// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ram Saml Providers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.114.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ram.getSamlProviders({
 *     ids: ["samlProviderName"],
 *     nameRegex: "tf-testAcc",
 * });
 * export const firstRamSamlProviderId = example.then(example => example.providers?.[0]?.id);
 * ```
 */
export function getSamlProviders(args?: GetSamlProvidersArgs, opts?: pulumi.InvokeOptions): Promise<GetSamlProvidersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:ram/getSamlProviders:getSamlProviders", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSamlProviders.
 */
export interface GetSamlProvidersArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of SAML Provider IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter results by SAML Provider name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getSamlProviders.
 */
export interface GetSamlProvidersResult {
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
    readonly providers: outputs.ram.GetSamlProvidersProvider[];
}
/**
 * This data source provides the Ram Saml Providers of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.114.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = alicloud.ram.getSamlProviders({
 *     ids: ["samlProviderName"],
 *     nameRegex: "tf-testAcc",
 * });
 * export const firstRamSamlProviderId = example.then(example => example.providers?.[0]?.id);
 * ```
 */
export function getSamlProvidersOutput(args?: GetSamlProvidersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSamlProvidersResult> {
    return pulumi.output(args).apply((a: any) => getSamlProviders(a, opts))
}

/**
 * A collection of arguments for invoking getSamlProviders.
 */
export interface GetSamlProvidersOutputArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of SAML Provider IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by SAML Provider name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}
