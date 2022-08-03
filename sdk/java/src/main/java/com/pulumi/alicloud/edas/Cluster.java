// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.edas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.edas.ClusterArgs;
import com.pulumi.alicloud.edas.inputs.ClusterState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an EDAS cluster resource.
 * 
 * &gt; **NOTE:** Available in 1.82.0+
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
 * import com.pulumi.alicloud.edas.Cluster;
 * import com.pulumi.alicloud.edas.ClusterArgs;
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
 *         var default_ = new Cluster(&#34;default&#34;, ClusterArgs.builder()        
 *             .clusterName(var_.cluster_name())
 *             .clusterType(var_.cluster_type())
 *             .networkMode(var_.network_mode())
 *             .logicalRegionId(var_.logical_region_id())
 *             .vpcId(var_.vpc_id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EDAS cluster can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:edas/cluster:Cluster cluster cluster_id
 * ```
 * 
 */
@ResourceType(type="alicloud:edas/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * The name of the cluster that you want to create.
     * 
     */
    @Export(name="clusterName", type=String.class, parameters={})
    private Output<String> clusterName;

    /**
     * @return The name of the cluster that you want to create.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * The type of the cluster that you want to create. Valid values only: 2: ECS cluster.
     * 
     */
    @Export(name="clusterType", type=Integer.class, parameters={})
    private Output<Integer> clusterType;

    /**
     * @return The type of the cluster that you want to create. Valid values only: 2: ECS cluster.
     * 
     */
    public Output<Integer> clusterType() {
        return this.clusterType;
    }
    /**
     * The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
     * 
     */
    @Export(name="logicalRegionId", type=String.class, parameters={})
    private Output</* @Nullable */ String> logicalRegionId;

    /**
     * @return The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
     * 
     */
    public Output<Optional<String>> logicalRegionId() {
        return Codegen.optional(this.logicalRegionId);
    }
    /**
     * The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
     * 
     */
    @Export(name="networkMode", type=Integer.class, parameters={})
    private Output<Integer> networkMode;

    /**
     * @return The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
     * 
     */
    public Output<Integer> networkMode() {
        return this.networkMode;
    }
    /**
     * The ID of the Virtual Private Cloud (VPC) for the cluster.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The ID of the Virtual Private Cloud (VPC) for the cluster.
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:edas/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:edas/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
