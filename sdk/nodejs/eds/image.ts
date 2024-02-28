// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Image resource.
 *
 * For information about ECD Image and how to use it, see [What is Image](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createimage).
 *
 * > **NOTE:** Available since v1.146.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "terraform-example";
 * const defaultRandomInteger = new random.RandomInteger("defaultRandomInteger", {
 *     min: 10000,
 *     max: 99999,
 * });
 * const defaultSimpleOfficeSite = new alicloud.eds.SimpleOfficeSite("defaultSimpleOfficeSite", {
 *     cidrBlock: "172.16.0.0/12",
 *     enableAdminAccess: true,
 *     desktopAccessType: "Internet",
 *     officeSiteName: pulumi.interpolate`${name}-${defaultRandomInteger.result}`,
 * });
 * const defaultEcdPolicyGroup = new alicloud.eds.EcdPolicyGroup("defaultEcdPolicyGroup", {
 *     policyGroupName: name,
 *     clipboard: "read",
 *     localDrive: "read",
 *     usbRedirect: "off",
 *     watermark: "off",
 *     authorizeAccessPolicyRules: [{
 *         description: name,
 *         cidrIp: "1.2.3.45/24",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         type: "inflow",
 *         policy: "accept",
 *         description: name,
 *         portRange: "80/80",
 *         ipProtocol: "TCP",
 *         priority: "1",
 *         cidrIp: "1.2.3.4/24",
 *     }],
 * });
 * const defaultBundles = alicloud.eds.getBundles({
 *     bundleType: "SYSTEM",
 * });
 * const defaultDesktop = new alicloud.eds.Desktop("defaultDesktop", {
 *     officeSiteId: defaultSimpleOfficeSite.id,
 *     policyGroupId: defaultEcdPolicyGroup.id,
 *     bundleId: defaultBundles.then(defaultBundles => defaultBundles.bundles?.[1]?.id),
 *     desktopName: name,
 * });
 * const defaultImage = new alicloud.eds.Image("defaultImage", {
 *     imageName: name,
 *     desktopId: defaultDesktop.id,
 *     description: name,
 * });
 * ```
 *
 * ## Import
 *
 * ECD Image can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/image:Image example <id>
 * ```
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * The description of the image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The desktop id of the desktop.
     */
    public readonly desktopId!: pulumi.Output<string>;
    /**
     * The name of the image.
     */
    public readonly imageName!: pulumi.Output<string | undefined>;
    /**
     * The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["desktopId"] = state ? state.desktopId : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            if ((!args || args.desktopId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desktopId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["desktopId"] = args ? args.desktopId : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    /**
     * The description of the image.
     */
    description?: pulumi.Input<string>;
    /**
     * The desktop id of the desktop.
     */
    desktopId?: pulumi.Input<string>;
    /**
     * The name of the image.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The description of the image.
     */
    description?: pulumi.Input<string>;
    /**
     * The desktop id of the desktop.
     */
    desktopId: pulumi.Input<string>;
    /**
     * The name of the image.
     */
    imageName?: pulumi.Input<string>;
}
