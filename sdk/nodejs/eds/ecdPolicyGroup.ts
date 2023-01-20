// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Elastic Desktop Service (ECD) Policy Group resource.
 *
 * For information about Elastic Desktop Service (ECD) Policy Group and how to use it, see [What is Policy Group](https://help.aliyun.com/document_detail/188382.html).
 *
 * > **NOTE:** Available in v1.130.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.eds.EcdPolicyGroup("default", {
 *     authorizeAccessPolicyRules: [{
 *         cidrIp: "1.2.3.45/24",
 *         description: "my-description1",
 *     }],
 *     authorizeSecurityPolicyRules: [{
 *         cidrIp: "1.2.3.4/24",
 *         description: "my-description",
 *         ipProtocol: "TCP",
 *         policy: "accept",
 *         portRange: "80/80",
 *         priority: "1",
 *         type: "inflow",
 *     }],
 *     clipboard: "read",
 *     localDrive: "read",
 *     policyGroupName: "my-policy-group",
 *     usbRedirect: "off",
 *     watermark: "off",
 * });
 * ```
 *
 * ## Import
 *
 * Elastic Desktop Service (ECD) Policy Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eds/ecdPolicyGroup:EcdPolicyGroup example <id>
 * ```
 */
export class EcdPolicyGroup extends pulumi.CustomResource {
    /**
     * Get an existing EcdPolicyGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcdPolicyGroupState, opts?: pulumi.CustomResourceOptions): EcdPolicyGroup {
        return new EcdPolicyGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eds/ecdPolicyGroup:EcdPolicyGroup';

    /**
     * Returns true if the given object is an instance of EcdPolicyGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcdPolicyGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcdPolicyGroup.__pulumiType;
    }

    /**
     * The rule of authorize access rule.
     */
    public readonly authorizeAccessPolicyRules!: pulumi.Output<outputs.eds.EcdPolicyGroupAuthorizeAccessPolicyRule[] | undefined>;
    /**
     * The policy rule.
     */
    public readonly authorizeSecurityPolicyRules!: pulumi.Output<outputs.eds.EcdPolicyGroupAuthorizeSecurityPolicyRule[] | undefined>;
    /**
     * Whether to enable local camera redirection. Valid values: `on`, `off`.
     */
    public readonly cameraRedirect!: pulumi.Output<string>;
    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     */
    public readonly clipboard!: pulumi.Output<string>;
    /**
     * The list of domain.
     */
    public readonly domainList!: pulumi.Output<string | undefined>;
    /**
     * The access of html5. Valid values: `off`, `on`.
     */
    public readonly htmlAccess!: pulumi.Output<string>;
    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     */
    public readonly htmlFileTransfer!: pulumi.Output<string>;
    /**
     * Local drive redirect policy. Valid values: ` readwrite`, `off`, `read`.
     */
    public readonly localDrive!: pulumi.Output<string>;
    /**
     * The name of policy group.
     */
    public readonly policyGroupName!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable screen recording. Valid values: `off`, `alltime`, `period`.
     */
    public readonly recording!: pulumi.Output<string>;
    /**
     * The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     */
    public readonly recordingEndTime!: pulumi.Output<string | undefined>;
    /**
     * The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `alltime`.
     */
    public readonly recordingExpires!: pulumi.Output<number>;
    /**
     * The fps of recording. Valid values: `2`, `5`, `10`, `15`.
     */
    public readonly recordingFps!: pulumi.Output<number>;
    /**
     * The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     */
    public readonly recordingStartTime!: pulumi.Output<string | undefined>;
    /**
     * The status of policy.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     */
    public readonly usbRedirect!: pulumi.Output<string>;
    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     */
    public readonly visualQuality!: pulumi.Output<string>;
    /**
     * The watermark policy. Valid values: `off`, `on`.
     */
    public readonly watermark!: pulumi.Output<string>;
    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     */
    public readonly watermarkTransparency!: pulumi.Output<string>;
    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     */
    public readonly watermarkType!: pulumi.Output<string>;

    /**
     * Create a EcdPolicyGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EcdPolicyGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcdPolicyGroupArgs | EcdPolicyGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcdPolicyGroupState | undefined;
            resourceInputs["authorizeAccessPolicyRules"] = state ? state.authorizeAccessPolicyRules : undefined;
            resourceInputs["authorizeSecurityPolicyRules"] = state ? state.authorizeSecurityPolicyRules : undefined;
            resourceInputs["cameraRedirect"] = state ? state.cameraRedirect : undefined;
            resourceInputs["clipboard"] = state ? state.clipboard : undefined;
            resourceInputs["domainList"] = state ? state.domainList : undefined;
            resourceInputs["htmlAccess"] = state ? state.htmlAccess : undefined;
            resourceInputs["htmlFileTransfer"] = state ? state.htmlFileTransfer : undefined;
            resourceInputs["localDrive"] = state ? state.localDrive : undefined;
            resourceInputs["policyGroupName"] = state ? state.policyGroupName : undefined;
            resourceInputs["recording"] = state ? state.recording : undefined;
            resourceInputs["recordingEndTime"] = state ? state.recordingEndTime : undefined;
            resourceInputs["recordingExpires"] = state ? state.recordingExpires : undefined;
            resourceInputs["recordingFps"] = state ? state.recordingFps : undefined;
            resourceInputs["recordingStartTime"] = state ? state.recordingStartTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["usbRedirect"] = state ? state.usbRedirect : undefined;
            resourceInputs["visualQuality"] = state ? state.visualQuality : undefined;
            resourceInputs["watermark"] = state ? state.watermark : undefined;
            resourceInputs["watermarkTransparency"] = state ? state.watermarkTransparency : undefined;
            resourceInputs["watermarkType"] = state ? state.watermarkType : undefined;
        } else {
            const args = argsOrState as EcdPolicyGroupArgs | undefined;
            resourceInputs["authorizeAccessPolicyRules"] = args ? args.authorizeAccessPolicyRules : undefined;
            resourceInputs["authorizeSecurityPolicyRules"] = args ? args.authorizeSecurityPolicyRules : undefined;
            resourceInputs["cameraRedirect"] = args ? args.cameraRedirect : undefined;
            resourceInputs["clipboard"] = args ? args.clipboard : undefined;
            resourceInputs["domainList"] = args ? args.domainList : undefined;
            resourceInputs["htmlAccess"] = args ? args.htmlAccess : undefined;
            resourceInputs["htmlFileTransfer"] = args ? args.htmlFileTransfer : undefined;
            resourceInputs["localDrive"] = args ? args.localDrive : undefined;
            resourceInputs["policyGroupName"] = args ? args.policyGroupName : undefined;
            resourceInputs["recording"] = args ? args.recording : undefined;
            resourceInputs["recordingEndTime"] = args ? args.recordingEndTime : undefined;
            resourceInputs["recordingExpires"] = args ? args.recordingExpires : undefined;
            resourceInputs["recordingFps"] = args ? args.recordingFps : undefined;
            resourceInputs["recordingStartTime"] = args ? args.recordingStartTime : undefined;
            resourceInputs["usbRedirect"] = args ? args.usbRedirect : undefined;
            resourceInputs["visualQuality"] = args ? args.visualQuality : undefined;
            resourceInputs["watermark"] = args ? args.watermark : undefined;
            resourceInputs["watermarkTransparency"] = args ? args.watermarkTransparency : undefined;
            resourceInputs["watermarkType"] = args ? args.watermarkType : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcdPolicyGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcdPolicyGroup resources.
 */
export interface EcdPolicyGroupState {
    /**
     * The rule of authorize access rule.
     */
    authorizeAccessPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeAccessPolicyRule>[]>;
    /**
     * The policy rule.
     */
    authorizeSecurityPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeSecurityPolicyRule>[]>;
    /**
     * Whether to enable local camera redirection. Valid values: `on`, `off`.
     */
    cameraRedirect?: pulumi.Input<string>;
    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     */
    clipboard?: pulumi.Input<string>;
    /**
     * The list of domain.
     */
    domainList?: pulumi.Input<string>;
    /**
     * The access of html5. Valid values: `off`, `on`.
     */
    htmlAccess?: pulumi.Input<string>;
    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     */
    htmlFileTransfer?: pulumi.Input<string>;
    /**
     * Local drive redirect policy. Valid values: ` readwrite`, `off`, `read`.
     */
    localDrive?: pulumi.Input<string>;
    /**
     * The name of policy group.
     */
    policyGroupName?: pulumi.Input<string>;
    /**
     * Whether to enable screen recording. Valid values: `off`, `alltime`, `period`.
     */
    recording?: pulumi.Input<string>;
    /**
     * The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     */
    recordingEndTime?: pulumi.Input<string>;
    /**
     * The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `alltime`.
     */
    recordingExpires?: pulumi.Input<number>;
    /**
     * The fps of recording. Valid values: `2`, `5`, `10`, `15`.
     */
    recordingFps?: pulumi.Input<number>;
    /**
     * The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     */
    recordingStartTime?: pulumi.Input<string>;
    /**
     * The status of policy.
     */
    status?: pulumi.Input<string>;
    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     */
    usbRedirect?: pulumi.Input<string>;
    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     */
    visualQuality?: pulumi.Input<string>;
    /**
     * The watermark policy. Valid values: `off`, `on`.
     */
    watermark?: pulumi.Input<string>;
    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     */
    watermarkTransparency?: pulumi.Input<string>;
    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     */
    watermarkType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcdPolicyGroup resource.
 */
export interface EcdPolicyGroupArgs {
    /**
     * The rule of authorize access rule.
     */
    authorizeAccessPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeAccessPolicyRule>[]>;
    /**
     * The policy rule.
     */
    authorizeSecurityPolicyRules?: pulumi.Input<pulumi.Input<inputs.eds.EcdPolicyGroupAuthorizeSecurityPolicyRule>[]>;
    /**
     * Whether to enable local camera redirection. Valid values: `on`, `off`.
     */
    cameraRedirect?: pulumi.Input<string>;
    /**
     * The clipboard policy. Valid values: `off`, `read`, `readwrite`.
     */
    clipboard?: pulumi.Input<string>;
    /**
     * The list of domain.
     */
    domainList?: pulumi.Input<string>;
    /**
     * The access of html5. Valid values: `off`, `on`.
     */
    htmlAccess?: pulumi.Input<string>;
    /**
     * The html5 file transfer. Valid values: `all`, `download`, `off`, `upload`.
     */
    htmlFileTransfer?: pulumi.Input<string>;
    /**
     * Local drive redirect policy. Valid values: ` readwrite`, `off`, `read`.
     */
    localDrive?: pulumi.Input<string>;
    /**
     * The name of policy group.
     */
    policyGroupName?: pulumi.Input<string>;
    /**
     * Whether to enable screen recording. Valid values: `off`, `alltime`, `period`.
     */
    recording?: pulumi.Input<string>;
    /**
     * The end time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     */
    recordingEndTime?: pulumi.Input<string>;
    /**
     * The screen recording video retention. Valid values between 30 and 180. This return value is meaningful only when the value of `recording` is `period` or `alltime`.
     */
    recordingExpires?: pulumi.Input<number>;
    /**
     * The fps of recording. Valid values: `2`, `5`, `10`, `15`.
     */
    recordingFps?: pulumi.Input<number>;
    /**
     * The start time of recording, value: `HH:MM:SS`. This return value is meaningful only when the value of `recording` is `period`.
     */
    recordingStartTime?: pulumi.Input<string>;
    /**
     * The usb redirect policy. Valid values: `off`, `on`.
     */
    usbRedirect?: pulumi.Input<string>;
    /**
     * The quality of visual. Valid values: `high`, `lossless`, `low`, `medium`.
     */
    visualQuality?: pulumi.Input<string>;
    /**
     * The watermark policy. Valid values: `off`, `on`.
     */
    watermark?: pulumi.Input<string>;
    /**
     * The watermark transparency. Valid values: `DARK`, `LIGHT`, `MIDDLE`.
     */
    watermarkTransparency?: pulumi.Input<string>;
    /**
     * The type of watemark. Valid values: `EndUserId`, `HostName`.
     */
    watermarkType?: pulumi.Input<string>;
}
