// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.LogTailAttachmentArgs;
import com.pulumi.alicloud.log.inputs.LogTailAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The Logtail access service is a log collection agent provided by Log Service.
 * You can use Logtail to collect logs from servers such as Alibaba Cloud Elastic
 * Compute Service (ECS) instances in real time in the Log Service console. [Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm)
 * 
 * This resource amis to attach one logtail configure to a machine group.
 * 
 * &gt; **NOTE:** One logtail configure can be attached to multiple machine groups and one machine group can attach several logtail configures.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomInteger;
 * import com.pulumi.random.RandomIntegerArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.log.Store;
 * import com.pulumi.alicloud.log.StoreArgs;
 * import com.pulumi.alicloud.log.LogTailConfig;
 * import com.pulumi.alicloud.log.LogTailConfigArgs;
 * import com.pulumi.alicloud.log.MachineGroup;
 * import com.pulumi.alicloud.log.MachineGroupArgs;
 * import com.pulumi.alicloud.log.LogTailAttachment;
 * import com.pulumi.alicloud.log.LogTailAttachmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new RandomInteger(&#34;default&#34;, RandomIntegerArgs.builder()        
 *             .max(99999)
 *             .min(10000)
 *             .build());
 * 
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *         var exampleStore = new Store(&#34;exampleStore&#34;, StoreArgs.builder()        
 *             .project(exampleProject.name())
 *             .retentionPeriod(3650)
 *             .shardCount(3)
 *             .autoSplit(true)
 *             .maxSplitShardCount(60)
 *             .appendMeta(true)
 *             .build());
 * 
 *         var exampleLogTailConfig = new LogTailConfig(&#34;exampleLogTailConfig&#34;, LogTailConfigArgs.builder()        
 *             .project(exampleProject.name())
 *             .logstore(exampleStore.name())
 *             .inputType(&#34;file&#34;)
 *             .outputType(&#34;LogService&#34;)
 *             .inputDetail(&#34;&#34;&#34;
 *   	{
 * 		&#34;logPath&#34;: &#34;/logPath&#34;,
 * 		&#34;filePattern&#34;: &#34;access.log&#34;,
 * 		&#34;logType&#34;: &#34;json_log&#34;,
 * 		&#34;topicFormat&#34;: &#34;default&#34;,
 * 		&#34;discardUnmatch&#34;: false,
 * 		&#34;enableRawLog&#34;: true,
 * 		&#34;fileEncoding&#34;: &#34;gbk&#34;,
 * 		&#34;maxDepth&#34;: 10
 * 	}
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var exampleMachineGroup = new MachineGroup(&#34;exampleMachineGroup&#34;, MachineGroupArgs.builder()        
 *             .project(exampleProject.name())
 *             .identifyType(&#34;ip&#34;)
 *             .topic(&#34;terraform&#34;)
 *             .identifyLists(            
 *                 &#34;10.0.0.1&#34;,
 *                 &#34;10.0.0.2&#34;)
 *             .build());
 * 
 *         var exampleLogTailAttachment = new LogTailAttachment(&#34;exampleLogTailAttachment&#34;, LogTailAttachmentArgs.builder()        
 *             .project(exampleProject.name())
 *             .logtailConfigName(exampleLogTailConfig.name())
 *             .machineGroupName(exampleMachineGroup.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Logtial to machine group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:log/logTailAttachment:LogTailAttachment example tf-log:tf-log-config:tf-log-machine-group
 * ```
 * 
 */
@ResourceType(type="alicloud:log/logTailAttachment:LogTailAttachment")
public class LogTailAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The Logtail configuration name, which is unique in the same project.
     * 
     */
    @Export(name="logtailConfigName", type=String.class, parameters={})
    private Output<String> logtailConfigName;

    /**
     * @return The Logtail configuration name, which is unique in the same project.
     * 
     */
    public Output<String> logtailConfigName() {
        return this.logtailConfigName;
    }
    /**
     * The machine group name, which is unique in the same project.
     * 
     */
    @Export(name="machineGroupName", type=String.class, parameters={})
    private Output<String> machineGroupName;

    /**
     * @return The machine group name, which is unique in the same project.
     * 
     */
    public Output<String> machineGroupName() {
        return this.machineGroupName;
    }
    /**
     * The project name to the log store belongs.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The project name to the log store belongs.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LogTailAttachment(String name) {
        this(name, LogTailAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LogTailAttachment(String name, LogTailAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LogTailAttachment(String name, LogTailAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/logTailAttachment:LogTailAttachment", name, args == null ? LogTailAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LogTailAttachment(String name, Output<String> id, @Nullable LogTailAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/logTailAttachment:LogTailAttachment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static LogTailAttachment get(String name, Output<String> id, @Nullable LogTailAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LogTailAttachment(name, id, state, options);
    }
}
