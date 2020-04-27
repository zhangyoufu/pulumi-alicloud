// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides available scheduled task resources. 
 * 
 * > **NOTE:** Available in 1.72.0+
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const ds = pulumi.output(alicloud.ess.getScheduledTasks({
 *     nameRegex: "scheduledTaskName",
 *     scheduledTaskId: "scheduledTaskId",
 * }, { async: true }));
 * 
 * export const firstScheduledTask = ds.tasks[0].id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ess_scheduled_tasks.html.markdown.
 */
export function getScheduledTasks(args?: GetScheduledTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduledTasksResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:ess/getScheduledTasks:getScheduledTasks", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "scheduledAction": args.scheduledAction,
        "scheduledTaskId": args.scheduledTaskId,
    }, opts);
}

/**
 * A collection of arguments for invoking getScheduledTasks.
 */
export interface GetScheduledTasksArgs {
    /**
     * A list of scheduled task IDs.
     */
    readonly ids?: string[];
    /**
     * A regex string to filter resulting scheduled tasks by name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The operation to be performed when a scheduled task is triggered.
     */
    readonly scheduledAction?: string;
    /**
     * The id of the scheduled task.
     */
    readonly scheduledTaskId?: string;
}

/**
 * A collection of values returned by getScheduledTasks.
 */
export interface GetScheduledTasksResult {
    /**
     * A list of scheduled task ids.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of scheduled task names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * The operation to be performed when a scheduled task is triggered.
     */
    readonly scheduledAction?: string;
    readonly scheduledTaskId?: string;
    /**
     * A list of scheduled tasks. Each element contains the following attributes:
     */
    readonly tasks: outputs.ess.GetScheduledTasksTask[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
