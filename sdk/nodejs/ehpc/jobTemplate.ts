// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Ehpc Job Template resource.
 *
 * For information about Ehpc Job Template and how to use it, see [What is Job Template](https://www.alibabacloud.com/help/product/57664.html).
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
 * const _default = new alicloud.ehpc.JobTemplate("default", {
 *     commandLine: "./LammpsTest/lammps.pbs",
 *     jobTemplateName: "example_value",
 * });
 * ```
 *
 * ## Import
 *
 * Ehpc Job Template can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import alicloud:ehpc/jobTemplate:JobTemplate example <id>
 * ```
 */
export class JobTemplate extends pulumi.CustomResource {
    /**
     * Get an existing JobTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobTemplateState, opts?: pulumi.CustomResourceOptions): JobTemplate {
        return new JobTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ehpc/jobTemplate:JobTemplate';

    /**
     * Returns true if the given object is an instance of JobTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobTemplate.__pulumiType;
    }

    /**
     * Queue Jobs, Is of the Form: 1-10:2.
     */
    public readonly arrayRequest!: pulumi.Output<string | undefined>;
    /**
     * Job Maximum Run Time.
     */
    public readonly clockTime!: pulumi.Output<string | undefined>;
    /**
     * Job Commands.
     */
    public readonly commandLine!: pulumi.Output<string>;
    /**
     * A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
     */
    public readonly gpu!: pulumi.Output<number | undefined>;
    /**
     * A Job Template Name.
     */
    public readonly jobTemplateName!: pulumi.Output<string>;
    /**
     * A Single Compute Node Maximum Memory.
     */
    public readonly mem!: pulumi.Output<string | undefined>;
    /**
     * Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
     */
    public readonly node!: pulumi.Output<number | undefined>;
    /**
     * Job Commands the Directory.
     */
    public readonly packagePath!: pulumi.Output<string | undefined>;
    /**
     * The Job Priority.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The Job Queue.
     */
    public readonly queue!: pulumi.Output<string | undefined>;
    /**
     * If the Job Is Support for the Re-Run.
     */
    public readonly reRunable!: pulumi.Output<boolean>;
    /**
     * The name of the user who performed the job.
     */
    public readonly runasUser!: pulumi.Output<string | undefined>;
    /**
     * Error Output Path.
     */
    public readonly stderrRedirectPath!: pulumi.Output<string | undefined>;
    /**
     * Standard Output Path and.
     */
    public readonly stdoutRedirectPath!: pulumi.Output<string | undefined>;
    /**
     * A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
     */
    public readonly task!: pulumi.Output<number | undefined>;
    /**
     * A Single Task and the Number of Required Threads.
     */
    public readonly thread!: pulumi.Output<number | undefined>;
    /**
     * The Job of the Environment Variable.
     */
    public readonly variables!: pulumi.Output<string | undefined>;

    /**
     * Create a JobTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobTemplateArgs | JobTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobTemplateState | undefined;
            resourceInputs["arrayRequest"] = state ? state.arrayRequest : undefined;
            resourceInputs["clockTime"] = state ? state.clockTime : undefined;
            resourceInputs["commandLine"] = state ? state.commandLine : undefined;
            resourceInputs["gpu"] = state ? state.gpu : undefined;
            resourceInputs["jobTemplateName"] = state ? state.jobTemplateName : undefined;
            resourceInputs["mem"] = state ? state.mem : undefined;
            resourceInputs["node"] = state ? state.node : undefined;
            resourceInputs["packagePath"] = state ? state.packagePath : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["queue"] = state ? state.queue : undefined;
            resourceInputs["reRunable"] = state ? state.reRunable : undefined;
            resourceInputs["runasUser"] = state ? state.runasUser : undefined;
            resourceInputs["stderrRedirectPath"] = state ? state.stderrRedirectPath : undefined;
            resourceInputs["stdoutRedirectPath"] = state ? state.stdoutRedirectPath : undefined;
            resourceInputs["task"] = state ? state.task : undefined;
            resourceInputs["thread"] = state ? state.thread : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as JobTemplateArgs | undefined;
            if ((!args || args.commandLine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commandLine'");
            }
            if ((!args || args.jobTemplateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobTemplateName'");
            }
            resourceInputs["arrayRequest"] = args ? args.arrayRequest : undefined;
            resourceInputs["clockTime"] = args ? args.clockTime : undefined;
            resourceInputs["commandLine"] = args ? args.commandLine : undefined;
            resourceInputs["gpu"] = args ? args.gpu : undefined;
            resourceInputs["jobTemplateName"] = args ? args.jobTemplateName : undefined;
            resourceInputs["mem"] = args ? args.mem : undefined;
            resourceInputs["node"] = args ? args.node : undefined;
            resourceInputs["packagePath"] = args ? args.packagePath : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["queue"] = args ? args.queue : undefined;
            resourceInputs["reRunable"] = args ? args.reRunable : undefined;
            resourceInputs["runasUser"] = args ? args.runasUser : undefined;
            resourceInputs["stderrRedirectPath"] = args ? args.stderrRedirectPath : undefined;
            resourceInputs["stdoutRedirectPath"] = args ? args.stdoutRedirectPath : undefined;
            resourceInputs["task"] = args ? args.task : undefined;
            resourceInputs["thread"] = args ? args.thread : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(JobTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering JobTemplate resources.
 */
export interface JobTemplateState {
    /**
     * Queue Jobs, Is of the Form: 1-10:2.
     */
    arrayRequest?: pulumi.Input<string>;
    /**
     * Job Maximum Run Time.
     */
    clockTime?: pulumi.Input<string>;
    /**
     * Job Commands.
     */
    commandLine?: pulumi.Input<string>;
    /**
     * A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
     */
    gpu?: pulumi.Input<number>;
    /**
     * A Job Template Name.
     */
    jobTemplateName?: pulumi.Input<string>;
    /**
     * A Single Compute Node Maximum Memory.
     */
    mem?: pulumi.Input<string>;
    /**
     * Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
     */
    node?: pulumi.Input<number>;
    /**
     * Job Commands the Directory.
     */
    packagePath?: pulumi.Input<string>;
    /**
     * The Job Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The Job Queue.
     */
    queue?: pulumi.Input<string>;
    /**
     * If the Job Is Support for the Re-Run.
     */
    reRunable?: pulumi.Input<boolean>;
    /**
     * The name of the user who performed the job.
     */
    runasUser?: pulumi.Input<string>;
    /**
     * Error Output Path.
     */
    stderrRedirectPath?: pulumi.Input<string>;
    /**
     * Standard Output Path and.
     */
    stdoutRedirectPath?: pulumi.Input<string>;
    /**
     * A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
     */
    task?: pulumi.Input<number>;
    /**
     * A Single Task and the Number of Required Threads.
     */
    thread?: pulumi.Input<number>;
    /**
     * The Job of the Environment Variable.
     */
    variables?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a JobTemplate resource.
 */
export interface JobTemplateArgs {
    /**
     * Queue Jobs, Is of the Form: 1-10:2.
     */
    arrayRequest?: pulumi.Input<string>;
    /**
     * Job Maximum Run Time.
     */
    clockTime?: pulumi.Input<string>;
    /**
     * Job Commands.
     */
    commandLine: pulumi.Input<string>;
    /**
     * A Single Compute Node Using the GPU Number.Possible Values: 1~20000.
     */
    gpu?: pulumi.Input<number>;
    /**
     * A Job Template Name.
     */
    jobTemplateName: pulumi.Input<string>;
    /**
     * A Single Compute Node Maximum Memory.
     */
    mem?: pulumi.Input<string>;
    /**
     * Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .
     */
    node?: pulumi.Input<number>;
    /**
     * Job Commands the Directory.
     */
    packagePath?: pulumi.Input<string>;
    /**
     * The Job Priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The Job Queue.
     */
    queue?: pulumi.Input<string>;
    /**
     * If the Job Is Support for the Re-Run.
     */
    reRunable?: pulumi.Input<boolean>;
    /**
     * The name of the user who performed the job.
     */
    runasUser?: pulumi.Input<string>;
    /**
     * Error Output Path.
     */
    stderrRedirectPath?: pulumi.Input<string>;
    /**
     * Standard Output Path and.
     */
    stdoutRedirectPath?: pulumi.Input<string>;
    /**
     * A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .
     */
    task?: pulumi.Input<number>;
    /**
     * A Single Task and the Number of Required Threads.
     */
    thread?: pulumi.Input<number>;
    /**
     * The Job of the Environment Variable.
     */
    variables?: pulumi.Input<string>;
}
