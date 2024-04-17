// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.simpleapplicationserver;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.simpleapplicationserver.SnapshotArgs;
import com.pulumi.alicloud.simpleapplicationserver.inputs.SnapshotState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Simple Application Server Snapshot resource.
 * 
 * For information about Simple Application Server Snapshot and how to use it, see [What is Snapshot](https://www.alibabacloud.com/help/doc-detail/190452.htm).
 * 
 * &gt; **NOTE:** Available since v1.143.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.simpleapplicationserver.SimpleapplicationserverFunctions;
 * import com.pulumi.alicloud.simpleapplicationserver.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.simpleapplicationserver.inputs.GetServerPlansArgs;
 * import com.pulumi.alicloud.simpleapplicationserver.Instance;
 * import com.pulumi.alicloud.simpleapplicationserver.InstanceArgs;
 * import com.pulumi.alicloud.simpleapplicationserver.inputs.GetServerDisksArgs;
 * import com.pulumi.alicloud.simpleapplicationserver.Snapshot;
 * import com.pulumi.alicloud.simpleapplicationserver.SnapshotArgs;
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
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var default = SimpleapplicationserverFunctions.getImages();
 * 
 *         final var defaultGetServerPlans = SimpleapplicationserverFunctions.getServerPlans();
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .paymentType(&#34;Subscription&#34;)
 *             .planId(defaultGetServerPlans.applyValue(getServerPlansResult -&gt; getServerPlansResult.plans()[0].id()))
 *             .instanceName(name)
 *             .imageId(default_.images()[0].id())
 *             .period(1)
 *             .dataDiskSize(100)
 *             .build());
 * 
 *         final var defaultGetServerDisks = SimpleapplicationserverFunctions.getServerDisks(GetServerDisksArgs.builder()
 *             .instanceId(defaultInstance.id())
 *             .build());
 * 
 *         var defaultSnapshot = new Snapshot(&#34;defaultSnapshot&#34;, SnapshotArgs.builder()        
 *             .diskId(defaultGetServerDisks.applyValue(getServerDisksResult -&gt; getServerDisksResult).applyValue(defaultGetServerDisks -&gt; defaultGetServerDisks.applyValue(getServerDisksResult -&gt; getServerDisksResult.ids()[0])))
 *             .snapshotName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Simple Application Server Snapshot can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:simpleapplicationserver/snapshot:Snapshot example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:simpleapplicationserver/snapshot:Snapshot")
public class Snapshot extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the disk.
     * 
     */
    @Export(name="diskId", refs={String.class}, tree="[0]")
    private Output<String> diskId;

    /**
     * @return The ID of the disk.
     * 
     */
    public Output<String> diskId() {
        return this.diskId;
    }
    /**
     * The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     * 
     */
    @Export(name="snapshotName", refs={String.class}, tree="[0]")
    private Output<String> snapshotName;

    /**
     * @return The name of the snapshot. The name must be `2` to `50` characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.),and hyphens (-).
     * 
     */
    public Output<String> snapshotName() {
        return this.snapshotName;
    }
    /**
     * The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the snapshot. Valid values: `Progressing`, `Accomplished` and `Failed`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Snapshot(String name) {
        this(name, SnapshotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Snapshot(String name, SnapshotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Snapshot(String name, SnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:simpleapplicationserver/snapshot:Snapshot", name, args == null ? SnapshotArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Snapshot(String name, Output<String> id, @Nullable SnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:simpleapplicationserver/snapshot:Snapshot", name, state, makeResourceOptions(options, id));
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
    public static Snapshot get(String name, Output<String> id, @Nullable SnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Snapshot(name, id, state, options);
    }
}
