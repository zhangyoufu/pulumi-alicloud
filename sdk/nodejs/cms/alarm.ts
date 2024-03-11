// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Monitor Service Alarm resource.
 *
 * For information about Cloud Monitor Service Alarm and how to use it, see [What is Alarm](https://www.alibabacloud.com/help/en/cloudmonitor/latest/putresourcemetricrule).
 *
 * > **NOTE:** Available since v1.9.1.
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
 * const name = config.get("name") || "tf-example";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: "Instance",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?.[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     owners: "system",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {
 *     vpcName: name,
 *     cidrBlock: "10.4.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vswitchName: name,
 *     cidrBlock: "10.4.0.0/24",
 *     vpcId: defaultNetwork.id,
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 * });
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetwork.id});
 * const defaultInstance = new alicloud.ecs.Instance("defaultInstance", {
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones?.[0]?.id),
 *     instanceName: name,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?.[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?.[0]?.id),
 *     securityGroups: [defaultSecurityGroup.id],
 *     vswitchId: defaultSwitch.id,
 * });
 * const defaultAlarmContactGroup = new alicloud.cms.AlarmContactGroup("defaultAlarmContactGroup", {alarmContactGroupName: name});
 * const defaultAlarm = new alicloud.cms.Alarm("defaultAlarm", {
 *     project: "acs_ecs_dashboard",
 *     metric: "disk_writebytes",
 *     period: 900,
 *     contactGroups: [defaultAlarmContactGroup.alarmContactGroupName],
 *     effectiveInterval: "06:00-20:00",
 *     metricDimensions: pulumi.interpolate`  [
 *     {
 *       "instanceId": "${defaultInstance.id}",
 *       "device": "/dev/vda1"
 *     }
 *   ]
 * `,
 *     escalationsCritical: {
 *         statistics: "Average",
 *         comparisonOperator: "<=",
 *         threshold: "35",
 *         times: 2,
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Cloud Monitor Service Alarm can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:cms/alarm:Alarm example <id>
 * ```
 */
export class Alarm extends pulumi.CustomResource {
    /**
     * Get an existing Alarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmState, opts?: pulumi.CustomResourceOptions): Alarm {
        return new Alarm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cms/alarm:Alarm';

    /**
     * Returns true if the given object is an instance of Alarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alarm.__pulumiType;
    }

    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     */
    public readonly contactGroups!: pulumi.Output<string[]>;
    /**
     * Field `dimensions` has been deprecated from provider version 1.173.0. New field `metricDimensions` instead.
     *
     * @deprecated Field `dimensions` has been deprecated from provider version 1.173.0. New field `metric_dimensions` instead.
     */
    public readonly dimensions!: pulumi.Output<{[key: string]: any}>;
    /**
     * The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
     */
    public readonly effectiveInterval!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable alarm rule. Default value: `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Field `endTime` has been deprecated from provider version 1.50.0. New field `effectiveInterval` instead.
     *
     * @deprecated Field `end_time` has been deprecated from provider version 1.50.0. New field `effective_interval` instead.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * A configuration of critical alarm. See `escalationsCritical` below.
     */
    public readonly escalationsCritical!: pulumi.Output<outputs.cms.AlarmEscalationsCritical>;
    /**
     * A configuration of critical info. See `escalationsInfo` below.
     */
    public readonly escalationsInfo!: pulumi.Output<outputs.cms.AlarmEscalationsInfo>;
    /**
     * A configuration of critical warn. See `escalationsWarn` below.
     */
    public readonly escalationsWarn!: pulumi.Output<outputs.cms.AlarmEscalationsWarn>;
    /**
     * The name of the metric, such as `CPUUtilization` and `networkinRate`. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    public readonly metric!: pulumi.Output<string>;
    /**
     * Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    public readonly metricDimensions!: pulumi.Output<string>;
    /**
     * The name of the alert rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The namespace of the cloud service, such as `acsEcsDashboard` and `acsRdsDashboard`. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * **NOTE:** The `dimensions` and `metricDimensions` must be empty when `project` is `acsPrometheus`, otherwise, one of them must be set.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The Prometheus alert rule. See `prometheus` below. **Note:** This parameter is required only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
     */
    public readonly prometheuses!: pulumi.Output<outputs.cms.AlarmPrometheus[]>;
    /**
     * Notification silence period in the alarm state, in seconds. Default value: `86400`. Valid value range: [300, 86400].
     */
    public readonly silenceTime!: pulumi.Output<number | undefined>;
    /**
     * Field `startTime` has been deprecated from provider version 1.50.0. New field `effectiveInterval` instead.
     *
     * @deprecated Field `start_time` has been deprecated from provider version 1.50.0. New field `effective_interval` instead.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * The status of the Alarm.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The information about the resource for which alerts are triggered. See `targets` below.
     */
    public readonly targets!: pulumi.Output<outputs.cms.AlarmTarget[]>;
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     */
    public readonly webhook!: pulumi.Output<string | undefined>;

    /**
     * Create a Alarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmArgs | AlarmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmState | undefined;
            resourceInputs["contactGroups"] = state ? state.contactGroups : undefined;
            resourceInputs["dimensions"] = state ? state.dimensions : undefined;
            resourceInputs["effectiveInterval"] = state ? state.effectiveInterval : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["escalationsCritical"] = state ? state.escalationsCritical : undefined;
            resourceInputs["escalationsInfo"] = state ? state.escalationsInfo : undefined;
            resourceInputs["escalationsWarn"] = state ? state.escalationsWarn : undefined;
            resourceInputs["metric"] = state ? state.metric : undefined;
            resourceInputs["metricDimensions"] = state ? state.metricDimensions : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["prometheuses"] = state ? state.prometheuses : undefined;
            resourceInputs["silenceTime"] = state ? state.silenceTime : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["targets"] = state ? state.targets : undefined;
            resourceInputs["webhook"] = state ? state.webhook : undefined;
        } else {
            const args = argsOrState as AlarmArgs | undefined;
            if ((!args || args.contactGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactGroups'");
            }
            if ((!args || args.metric === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metric'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["contactGroups"] = args ? args.contactGroups : undefined;
            resourceInputs["dimensions"] = args ? args.dimensions : undefined;
            resourceInputs["effectiveInterval"] = args ? args.effectiveInterval : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["escalationsCritical"] = args ? args.escalationsCritical : undefined;
            resourceInputs["escalationsInfo"] = args ? args.escalationsInfo : undefined;
            resourceInputs["escalationsWarn"] = args ? args.escalationsWarn : undefined;
            resourceInputs["metric"] = args ? args.metric : undefined;
            resourceInputs["metricDimensions"] = args ? args.metricDimensions : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["prometheuses"] = args ? args.prometheuses : undefined;
            resourceInputs["silenceTime"] = args ? args.silenceTime : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["webhook"] = args ? args.webhook : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alarm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alarm resources.
 */
export interface AlarmState {
    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     */
    contactGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Field `dimensions` has been deprecated from provider version 1.173.0. New field `metricDimensions` instead.
     *
     * @deprecated Field `dimensions` has been deprecated from provider version 1.173.0. New field `metric_dimensions` instead.
     */
    dimensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
     */
    effectiveInterval?: pulumi.Input<string>;
    /**
     * Whether to enable alarm rule. Default value: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Field `endTime` has been deprecated from provider version 1.50.0. New field `effectiveInterval` instead.
     *
     * @deprecated Field `end_time` has been deprecated from provider version 1.50.0. New field `effective_interval` instead.
     */
    endTime?: pulumi.Input<number>;
    /**
     * A configuration of critical alarm. See `escalationsCritical` below.
     */
    escalationsCritical?: pulumi.Input<inputs.cms.AlarmEscalationsCritical>;
    /**
     * A configuration of critical info. See `escalationsInfo` below.
     */
    escalationsInfo?: pulumi.Input<inputs.cms.AlarmEscalationsInfo>;
    /**
     * A configuration of critical warn. See `escalationsWarn` below.
     */
    escalationsWarn?: pulumi.Input<inputs.cms.AlarmEscalationsWarn>;
    /**
     * The name of the metric, such as `CPUUtilization` and `networkinRate`. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    metric?: pulumi.Input<string>;
    /**
     * Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    metricDimensions?: pulumi.Input<string>;
    /**
     * The name of the alert rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     */
    period?: pulumi.Input<number>;
    /**
     * The namespace of the cloud service, such as `acsEcsDashboard` and `acsRdsDashboard`. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * **NOTE:** The `dimensions` and `metricDimensions` must be empty when `project` is `acsPrometheus`, otherwise, one of them must be set.
     */
    project?: pulumi.Input<string>;
    /**
     * The Prometheus alert rule. See `prometheus` below. **Note:** This parameter is required only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
     */
    prometheuses?: pulumi.Input<pulumi.Input<inputs.cms.AlarmPrometheus>[]>;
    /**
     * Notification silence period in the alarm state, in seconds. Default value: `86400`. Valid value range: [300, 86400].
     */
    silenceTime?: pulumi.Input<number>;
    /**
     * Field `startTime` has been deprecated from provider version 1.50.0. New field `effectiveInterval` instead.
     *
     * @deprecated Field `start_time` has been deprecated from provider version 1.50.0. New field `effective_interval` instead.
     */
    startTime?: pulumi.Input<number>;
    /**
     * The status of the Alarm.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The information about the resource for which alerts are triggered. See `targets` below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.cms.AlarmTarget>[]>;
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     */
    webhook?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * List contact groups of the alarm rule, which must have been created on the console.
     */
    contactGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Field `dimensions` has been deprecated from provider version 1.173.0. New field `metricDimensions` instead.
     *
     * @deprecated Field `dimensions` has been deprecated from provider version 1.173.0. New field `metric_dimensions` instead.
     */
    dimensions?: pulumi.Input<{[key: string]: any}>;
    /**
     * The interval of effecting alarm rule. It format as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".
     */
    effectiveInterval?: pulumi.Input<string>;
    /**
     * Whether to enable alarm rule. Default value: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Field `endTime` has been deprecated from provider version 1.50.0. New field `effectiveInterval` instead.
     *
     * @deprecated Field `end_time` has been deprecated from provider version 1.50.0. New field `effective_interval` instead.
     */
    endTime?: pulumi.Input<number>;
    /**
     * A configuration of critical alarm. See `escalationsCritical` below.
     */
    escalationsCritical?: pulumi.Input<inputs.cms.AlarmEscalationsCritical>;
    /**
     * A configuration of critical info. See `escalationsInfo` below.
     */
    escalationsInfo?: pulumi.Input<inputs.cms.AlarmEscalationsInfo>;
    /**
     * A configuration of critical warn. See `escalationsWarn` below.
     */
    escalationsWarn?: pulumi.Input<inputs.cms.AlarmEscalationsWarn>;
    /**
     * The name of the metric, such as `CPUUtilization` and `networkinRate`. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    metric: pulumi.Input<string>;
    /**
     * Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string, and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     */
    metricDimensions?: pulumi.Input<string>;
    /**
     * The name of the alert rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.
     */
    period?: pulumi.Input<number>;
    /**
     * The namespace of the cloud service, such as `acsEcsDashboard` and `acsRdsDashboard`. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).
     * **NOTE:** The `dimensions` and `metricDimensions` must be empty when `project` is `acsPrometheus`, otherwise, one of them must be set.
     */
    project: pulumi.Input<string>;
    /**
     * The Prometheus alert rule. See `prometheus` below. **Note:** This parameter is required only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
     */
    prometheuses?: pulumi.Input<pulumi.Input<inputs.cms.AlarmPrometheus>[]>;
    /**
     * Notification silence period in the alarm state, in seconds. Default value: `86400`. Valid value range: [300, 86400].
     */
    silenceTime?: pulumi.Input<number>;
    /**
     * Field `startTime` has been deprecated from provider version 1.50.0. New field `effectiveInterval` instead.
     *
     * @deprecated Field `start_time` has been deprecated from provider version 1.50.0. New field `effective_interval` instead.
     */
    startTime?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The information about the resource for which alerts are triggered. See `targets` below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.cms.AlarmTarget>[]>;
    /**
     * The webhook that should be called when the alarm is triggered. Currently, only http protocol is supported. Default is empty string.
     */
    webhook?: pulumi.Input<string>;
}
