// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a ECD Simple Office Site resource.
 *
 * For information about ECD Simple Office Site and how to use it, see [What is Simple Office Site](https://www.alibabacloud.com/help/en/wuying-workspace/developer-reference/api-ecd-2020-09-30-createsimpleofficesite).
 *
 * > **NOTE:** Available since v1.140.0.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * import * as random from "@pulumi/random";
 *
 * const defaultRandomInteger = new random.RandomInteger("defaultRandomInteger", {
 *     max: 99999,
 *     min: 10000,
 * });
 * const defaultSimpleOfficeSite = new alicloud.eds.SimpleOfficeSite("defaultSimpleOfficeSite", {
 *     cidrBlock: "172.16.0.0/12",
 *     desktopAccessType: "Internet",
 *     enableAdminAccess: true,
 *     officeSiteName: pulumi.interpolate`terraform-example-${defaultRandomInteger.result}`,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ECD Simple Office Site can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:eds/simpleOfficeSite:SimpleOfficeSite example <id>
 * ```
 */
export class SimpleOfficeSite extends pulumi.CustomResource {
    /**
     * Get an existing SimpleOfficeSite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SimpleOfficeSiteState, opts?: pulumi.CustomResourceOptions): SimpleOfficeSite {
        return new SimpleOfficeSite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/simpleOfficeSite:SimpleOfficeSite';

    /**
     * Returns true if the given object is an instance of SimpleOfficeSite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SimpleOfficeSite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SimpleOfficeSite.__pulumiType;
    }

    /**
     * The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
     *
     * @deprecated Field 'bandwidth' has been deprecated from provider version 1.142.0.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * Cloud Enterprise Network Instance ID.
     */
    public readonly cenId!: pulumi.Output<string | undefined>;
    /**
     * The cen owner id.
     */
    public readonly cenOwnerId!: pulumi.Output<string | undefined>;
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
     */
    public readonly desktopAccessType!: pulumi.Output<string>;
    /**
     * Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
     */
    public readonly enableAdminAccess!: pulumi.Output<boolean>;
    /**
     * Enable Cross-Desktop Access.
     */
    public readonly enableCrossDesktopAccess!: pulumi.Output<boolean>;
    /**
     * Whether the Open Internet Access Function.
     *
     * @deprecated Field 'enable_internet_access' has been deprecated from provider version 1.142.0.
     */
    public readonly enableInternetAccess!: pulumi.Output<boolean>;
    /**
     * Whether to Enable Multi-Factor Authentication MFA.
     */
    public readonly mfaEnabled!: pulumi.Output<boolean>;
    /**
     * The office site name.
     */
    public readonly officeSiteName!: pulumi.Output<string | undefined>;
    /**
     * Whether to Enable Single Sign-on (SSO) for User-Based SSO.
     */
    public readonly ssoEnabled!: pulumi.Output<boolean>;
    /**
     * Workspace State. Valid Values: `REGISTERED`,`REGISTERING`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a SimpleOfficeSite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SimpleOfficeSiteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SimpleOfficeSiteArgs | SimpleOfficeSiteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SimpleOfficeSiteState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["cenOwnerId"] = state ? state.cenOwnerId : undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["desktopAccessType"] = state ? state.desktopAccessType : undefined;
            resourceInputs["enableAdminAccess"] = state ? state.enableAdminAccess : undefined;
            resourceInputs["enableCrossDesktopAccess"] = state ? state.enableCrossDesktopAccess : undefined;
            resourceInputs["enableInternetAccess"] = state ? state.enableInternetAccess : undefined;
            resourceInputs["mfaEnabled"] = state ? state.mfaEnabled : undefined;
            resourceInputs["officeSiteName"] = state ? state.officeSiteName : undefined;
            resourceInputs["ssoEnabled"] = state ? state.ssoEnabled : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SimpleOfficeSiteArgs | undefined;
            if ((!args || args.cidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["cenOwnerId"] = args ? args.cenOwnerId : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["desktopAccessType"] = args ? args.desktopAccessType : undefined;
            resourceInputs["enableAdminAccess"] = args ? args.enableAdminAccess : undefined;
            resourceInputs["enableCrossDesktopAccess"] = args ? args.enableCrossDesktopAccess : undefined;
            resourceInputs["enableInternetAccess"] = args ? args.enableInternetAccess : undefined;
            resourceInputs["mfaEnabled"] = args ? args.mfaEnabled : undefined;
            resourceInputs["officeSiteName"] = args ? args.officeSiteName : undefined;
            resourceInputs["ssoEnabled"] = args ? args.ssoEnabled : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SimpleOfficeSite.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SimpleOfficeSite resources.
 */
export interface SimpleOfficeSiteState {
    /**
     * The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
     *
     * @deprecated Field 'bandwidth' has been deprecated from provider version 1.142.0.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Cloud Enterprise Network Instance ID.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The cen owner id.
     */
    cenOwnerId?: pulumi.Input<string>;
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
     */
    desktopAccessType?: pulumi.Input<string>;
    /**
     * Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
     */
    enableAdminAccess?: pulumi.Input<boolean>;
    /**
     * Enable Cross-Desktop Access.
     */
    enableCrossDesktopAccess?: pulumi.Input<boolean>;
    /**
     * Whether the Open Internet Access Function.
     *
     * @deprecated Field 'enable_internet_access' has been deprecated from provider version 1.142.0.
     */
    enableInternetAccess?: pulumi.Input<boolean>;
    /**
     * Whether to Enable Multi-Factor Authentication MFA.
     */
    mfaEnabled?: pulumi.Input<boolean>;
    /**
     * The office site name.
     */
    officeSiteName?: pulumi.Input<string>;
    /**
     * Whether to Enable Single Sign-on (SSO) for User-Based SSO.
     */
    ssoEnabled?: pulumi.Input<boolean>;
    /**
     * Workspace State. Valid Values: `REGISTERED`,`REGISTERING`.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SimpleOfficeSite resource.
 */
export interface SimpleOfficeSiteArgs {
    /**
     * The Internet Bandwidth Peak. It has been deprecated from version 1.142.0 and can be found in the new resource alicloud_ecd_network_package.
     *
     * @deprecated Field 'bandwidth' has been deprecated from provider version 1.142.0.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Cloud Enterprise Network Instance ID.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The cen owner id.
     */
    cenOwnerId?: pulumi.Input<string>;
    /**
     * Workspace Corresponds to the Security Office Network of IPv4 Segment.
     */
    cidrBlock: pulumi.Input<string>;
    /**
     * Connect to the Cloud Desktop Allows the Use of the Access Mode of. Valid values: `Any`, `Internet`, `VPC`.
     */
    desktopAccessType?: pulumi.Input<string>;
    /**
     * Whether to Use Cloud Desktop User Empowerment of Local Administrator Permissions.
     */
    enableAdminAccess?: pulumi.Input<boolean>;
    /**
     * Enable Cross-Desktop Access.
     */
    enableCrossDesktopAccess?: pulumi.Input<boolean>;
    /**
     * Whether the Open Internet Access Function.
     *
     * @deprecated Field 'enable_internet_access' has been deprecated from provider version 1.142.0.
     */
    enableInternetAccess?: pulumi.Input<boolean>;
    /**
     * Whether to Enable Multi-Factor Authentication MFA.
     */
    mfaEnabled?: pulumi.Input<boolean>;
    /**
     * The office site name.
     */
    officeSiteName?: pulumi.Input<string>;
    /**
     * Whether to Enable Single Sign-on (SSO) for User-Based SSO.
     */
    ssoEnabled?: pulumi.Input<boolean>;
}
