// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Alidns Custom Line resource.
 *
 * For information about Alidns Custom Line and how to use it, see [What is Custom Line](https://www.alibabacloud.com/help/en/doc-detail/145059.html).
 *
 * > **NOTE:** Available since v1.151.0.
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
 * const _default = new alicloud.dns.CustomLine("default", {
 *     customLineName: "tf-example",
 *     domainName: "alicloud-provider.com",
 *     ipSegmentLists: [{
 *         endIp: "192.0.2.125",
 *         startIp: "192.0.2.123",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Alidns Custom Line can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:dns/customLine:CustomLine example <id>
 * ```
 */
export class CustomLine extends pulumi.CustomResource {
    /**
     * Get an existing CustomLine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomLineState, opts?: pulumi.CustomResourceOptions): CustomLine {
        return new CustomLine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dns/customLine:CustomLine';

    /**
     * Returns true if the given object is an instance of CustomLine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomLine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomLine.__pulumiType;
    }

    /**
     * The name of the Custom Line.
     */
    public readonly customLineName!: pulumi.Output<string>;
    /**
     * The Domain name.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The IP segment list. See `ipSegmentList` below for details.
     */
    public readonly ipSegmentLists!: pulumi.Output<outputs.dns.CustomLineIpSegmentList[]>;
    /**
     * The lang.
     */
    public readonly lang!: pulumi.Output<string | undefined>;

    /**
     * Create a CustomLine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomLineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomLineArgs | CustomLineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomLineState | undefined;
            resourceInputs["customLineName"] = state ? state.customLineName : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["ipSegmentLists"] = state ? state.ipSegmentLists : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
        } else {
            const args = argsOrState as CustomLineArgs | undefined;
            if ((!args || args.customLineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customLineName'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.ipSegmentLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipSegmentLists'");
            }
            resourceInputs["customLineName"] = args ? args.customLineName : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["ipSegmentLists"] = args ? args.ipSegmentLists : undefined;
            resourceInputs["lang"] = args ? args.lang : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomLine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomLine resources.
 */
export interface CustomLineState {
    /**
     * The name of the Custom Line.
     */
    customLineName?: pulumi.Input<string>;
    /**
     * The Domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The IP segment list. See `ipSegmentList` below for details.
     */
    ipSegmentLists?: pulumi.Input<pulumi.Input<inputs.dns.CustomLineIpSegmentList>[]>;
    /**
     * The lang.
     */
    lang?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomLine resource.
 */
export interface CustomLineArgs {
    /**
     * The name of the Custom Line.
     */
    customLineName: pulumi.Input<string>;
    /**
     * The Domain name.
     */
    domainName: pulumi.Input<string>;
    /**
     * The IP segment list. See `ipSegmentList` below for details.
     */
    ipSegmentLists: pulumi.Input<pulumi.Input<inputs.dns.CustomLineIpSegmentList>[]>;
    /**
     * The lang.
     */
    lang?: pulumi.Input<string>;
}
