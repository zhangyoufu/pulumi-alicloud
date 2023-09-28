// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ebs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ebs.DiskReplicaGroupArgs;
import com.pulumi.alicloud.ebs.inputs.DiskReplicaGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a EBS Disk Replica Group resource.
 * 
 * For information about EBS Disk Replica Group and how to use it, see [What is Disk Replica Group](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/creatediskreplicagroup).
 * 
 * &gt; **NOTE:** Available since v1.187.0.
 * 
 * ## Import
 * 
 * EBS Disk Replica Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ebs/diskReplicaGroup:DiskReplicaGroup example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ebs/diskReplicaGroup:DiskReplicaGroup")
public class DiskReplicaGroup extends com.pulumi.resources.CustomResource {
    /**
     * The description of the consistent replication group.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the consistent replication group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of the region to which the disaster recovery site belongs.
     * 
     */
    @Export(name="destinationRegionId", type=String.class, parameters={})
    private Output<String> destinationRegionId;

    /**
     * @return The ID of the region to which the disaster recovery site belongs.
     * 
     */
    public Output<String> destinationRegionId() {
        return this.destinationRegionId;
    }
    /**
     * The ID of the zone to which the disaster recovery site belongs.
     * 
     */
    @Export(name="destinationZoneId", type=String.class, parameters={})
    private Output<String> destinationZoneId;

    /**
     * @return The ID of the zone to which the disaster recovery site belongs.
     * 
     */
    public Output<String> destinationZoneId() {
        return this.destinationZoneId;
    }
    /**
     * Consistent replication group name.
     * 
     */
    @Export(name="groupName", type=String.class, parameters={})
    private Output</* @Nullable */ String> groupName;

    /**
     * @return Consistent replication group name.
     * 
     */
    public Output<Optional<String>> groupName() {
        return Codegen.optional(this.groupName);
    }
    /**
     * The recovery point objective (RPO) of the replication pair-consistent group. Unit: seconds.
     * 
     */
    @Export(name="rpo", type=Integer.class, parameters={})
    private Output<Integer> rpo;

    /**
     * @return The recovery point objective (RPO) of the replication pair-consistent group. Unit: seconds.
     * 
     */
    public Output<Integer> rpo() {
        return this.rpo;
    }
    /**
     * The ID of the region to which the production site belongs.
     * 
     */
    @Export(name="sourceRegionId", type=String.class, parameters={})
    private Output<String> sourceRegionId;

    /**
     * @return The ID of the region to which the production site belongs.
     * 
     */
    public Output<String> sourceRegionId() {
        return this.sourceRegionId;
    }
    /**
     * The ID of the zone to which the production site belongs.
     * 
     */
    @Export(name="sourceZoneId", type=String.class, parameters={})
    private Output<String> sourceZoneId;

    /**
     * @return The ID of the zone to which the production site belongs.
     * 
     */
    public Output<String> sourceZoneId() {
        return this.sourceZoneId;
    }
    /**
     * The status of the consistent replication group.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the consistent replication group.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DiskReplicaGroup(String name) {
        this(name, DiskReplicaGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DiskReplicaGroup(String name, DiskReplicaGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DiskReplicaGroup(String name, DiskReplicaGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/diskReplicaGroup:DiskReplicaGroup", name, args == null ? DiskReplicaGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DiskReplicaGroup(String name, Output<String> id, @Nullable DiskReplicaGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ebs/diskReplicaGroup:DiskReplicaGroup", name, state, makeResourceOptions(options, id));
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
    public static DiskReplicaGroup get(String name, Output<String> id, @Nullable DiskReplicaGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DiskReplicaGroup(name, id, state, options);
    }
}
