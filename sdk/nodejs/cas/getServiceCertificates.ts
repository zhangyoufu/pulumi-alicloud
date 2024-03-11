// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the Ssl Certificates Service Certificates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.129.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const certs = alicloud.cas.getCertificates({
 *     nameRegex: "^cas",
 *     ids: ["Certificate Id"],
 * });
 * export const cert = certs.then(certs => certs.certificates?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceCertificates(args?: GetServiceCertificatesArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceCertificatesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("alicloud:cas/getServiceCertificates:getServiceCertificates", {
        "enableDetails": args.enableDetails,
        "ids": args.ids,
        "lang": args.lang,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceCertificates.
 */
export interface GetServiceCertificatesArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: boolean;
    /**
     * A list of Certificate IDs.
     */
    ids?: string[];
    /**
     * The lang.
     */
    lang?: string;
    /**
     * A regex string to filter results by Certificate name.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getServiceCertificates.
 */
export interface GetServiceCertificatesResult {
    readonly certificates: outputs.cas.GetServiceCertificatesCertificate[];
    readonly enableDetails?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly lang?: string;
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}
/**
 * This data source provides the Ssl Certificates Service Certificates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.129.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const certs = alicloud.cas.getCertificates({
 *     nameRegex: "^cas",
 *     ids: ["Certificate Id"],
 * });
 * export const cert = certs.then(certs => certs.certificates?.[0]?.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServiceCertificatesOutput(args?: GetServiceCertificatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceCertificatesResult> {
    return pulumi.output(args).apply((a: any) => getServiceCertificates(a, opts))
}

/**
 * A collection of arguments for invoking getServiceCertificates.
 */
export interface GetServiceCertificatesOutputArgs {
    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     */
    enableDetails?: pulumi.Input<boolean>;
    /**
     * A list of Certificate IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The lang.
     */
    lang?: pulumi.Input<string>;
    /**
     * A regex string to filter results by Certificate name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results (after running `pulumi preview`).
     */
    outputFile?: pulumi.Input<string>;
}
