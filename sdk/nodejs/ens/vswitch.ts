// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ENS Vswitch resource.
 *
 * For information about ENS Vswitch and how to use it, see [What is Vswitch](https://www.alibabacloud.com/help/en/ens/developer-reference/api-createvswitch).
 *
 * > **NOTE:** Available since v1.213.0.
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
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const _default = new alicloud.ens.Network("default", {
 *     networkName: name,
 *     description: name,
 *     cidrBlock: "192.168.2.0/24",
 *     ensRegionId: "cn-chenzhou-telecom_unicom_cmcc",
 * });
 * const defaultVswitch = new alicloud.ens.Vswitch("default", {
 *     description: name,
 *     cidrBlock: "192.168.2.0/24",
 *     vswitchName: name,
 *     ensRegionId: "cn-chenzhou-telecom_unicom_cmcc",
 *     networkId: _default.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ENS Vswitch can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ens/vswitch:Vswitch example <id>
 * ```
 */
export class Vswitch extends pulumi.CustomResource {
    /**
     * Get an existing Vswitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VswitchState, opts?: pulumi.CustomResourceOptions): Vswitch {
        return new Vswitch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ens/vswitch:Vswitch';

    /**
     * Returns true if the given object is an instance of Vswitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vswitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vswitch.__pulumiType;
    }

    /**
     * IPv4 CIDR block of the VSwitch instance.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The creation time of the VSwitch instance, in the UTC time format, yyyy-MM-ddTHH:mm:ssZ.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the VSwitch Instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ENS Region ID.
     */
    public readonly ensRegionId!: pulumi.Output<string>;
    /**
     * Network ID of the VSwitch instance.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Status of the switch instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Name of the switch instance.
     */
    public readonly vswitchName!: pulumi.Output<string | undefined>;

    /**
     * Create a Vswitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VswitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VswitchArgs | VswitchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VswitchState | undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ensRegionId"] = state ? state.ensRegionId : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vswitchName"] = state ? state.vswitchName : undefined;
        } else {
            const args = argsOrState as VswitchArgs | undefined;
            if ((!args || args.cidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if ((!args || args.ensRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ensRegionId'");
            }
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ensRegionId"] = args ? args.ensRegionId : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["vswitchName"] = args ? args.vswitchName : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vswitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vswitch resources.
 */
export interface VswitchState {
    /**
     * IPv4 CIDR block of the VSwitch instance.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The creation time of the VSwitch instance, in the UTC time format, yyyy-MM-ddTHH:mm:ssZ.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the VSwitch Instance.
     */
    description?: pulumi.Input<string>;
    /**
     * ENS Region ID.
     */
    ensRegionId?: pulumi.Input<string>;
    /**
     * Network ID of the VSwitch instance.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Status of the switch instance.
     */
    status?: pulumi.Input<string>;
    /**
     * Name of the switch instance.
     */
    vswitchName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vswitch resource.
 */
export interface VswitchArgs {
    /**
     * IPv4 CIDR block of the VSwitch instance.
     */
    cidrBlock: pulumi.Input<string>;
    /**
     * Description of the VSwitch Instance.
     */
    description?: pulumi.Input<string>;
    /**
     * ENS Region ID.
     */
    ensRegionId: pulumi.Input<string>;
    /**
     * Network ID of the VSwitch instance.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Name of the switch instance.
     */
    vswitchName?: pulumi.Input<string>;
}
