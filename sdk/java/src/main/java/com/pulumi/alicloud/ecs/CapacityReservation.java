// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.CapacityReservationArgs;
import com.pulumi.alicloud.ecs.inputs.CapacityReservationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Ecs Capacity Reservation resource.
 * 
 * For information about Ecs Capacity Reservation and how to use it, see [What is Capacity Reservation](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/createcapacityreservation).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
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
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.ecs.CapacityReservation;
 * import com.pulumi.alicloud.ecs.CapacityReservationArgs;
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
 *         final var defaultInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .instanceTypeFamily(&#34;ecs.g5&#34;)
 *             .build());
 * 
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;Instance&#34;)
 *             .availableInstanceType(defaultInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.ids()[0]))
 *             .build());
 * 
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status(&#34;OK&#34;)
 *             .build());
 * 
 *         var defaultCapacityReservation = new CapacityReservation(&#34;defaultCapacityReservation&#34;, CapacityReservationArgs.builder()        
 *             .description(&#34;terraform-example&#34;)
 *             .platform(&#34;linux&#34;)
 *             .capacityReservationName(&#34;terraform-example&#34;)
 *             .endTimeType(&#34;Unlimited&#34;)
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .instanceAmount(1)
 *             .instanceType(defaultInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.ids()[0]))
 *             .matchCriteria(&#34;Open&#34;)
 *             .tags(Map.of(&#34;Created&#34;, &#34;terraform-example&#34;))
 *             .zoneIds(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Ecs Capacity Reservation can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/capacityReservation:CapacityReservation example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/capacityReservation:CapacityReservation")
public class CapacityReservation extends com.pulumi.resources.CustomResource {
    /**
     * Capacity reservation service name.
     * 
     */
    @Export(name="capacityReservationName", refs={String.class}, tree="[0]")
    private Output<String> capacityReservationName;

    /**
     * @return Capacity reservation service name.
     * 
     */
    public Output<String> capacityReservationName() {
        return this.capacityReservationName;
    }
    /**
     * description of the capacity reservation instance.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return description of the capacity reservation instance.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies whether to pre-check the API request. Valid values: `true` and `false`.
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to pre-check the API request. Valid values: `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * end time of the capacity reservation. the capacity reservation will be  released at the end time automatically if set. otherwise it will last until manually released
     * 
     */
    @Export(name="endTime", refs={String.class}, tree="[0]")
    private Output<String> endTime;

    /**
     * @return end time of the capacity reservation. the capacity reservation will be  released at the end time automatically if set. otherwise it will last until manually released
     * 
     */
    public Output<String> endTime() {
        return this.endTime;
    }
    /**
     * Release mode of capacity reservation service. Value range:Limited: release at specified time. The EndTime parameter must be specified at the same time.Unlimited: manual release. No time limit.
     * 
     */
    @Export(name="endTimeType", refs={String.class}, tree="[0]")
    private Output<String> endTimeType;

    /**
     * @return Release mode of capacity reservation service. Value range:Limited: release at specified time. The EndTime parameter must be specified at the same time.Unlimited: manual release. No time limit.
     * 
     */
    public Output<String> endTimeType() {
        return this.endTimeType;
    }
    /**
     * The total number of instances that need to be reserved within the capacity reservation.
     * 
     */
    @Export(name="instanceAmount", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceAmount;

    /**
     * @return The total number of instances that need to be reserved within the capacity reservation.
     * 
     */
    public Output<Integer> instanceAmount() {
        return this.instanceAmount;
    }
    /**
     * Instance type. Currently, you can only set the capacity reservation service for one instance type.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    /**
     * @return Instance type. Currently, you can only set the capacity reservation service for one instance type.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The type of private resource pool generated after the capacity reservation service takes effect. Value range:Open: Open mode.Target: dedicated mode.Default value: Open
     * 
     */
    @Export(name="matchCriteria", refs={String.class}, tree="[0]")
    private Output<String> matchCriteria;

    /**
     * @return The type of private resource pool generated after the capacity reservation service takes effect. Value range:Open: Open mode.Target: dedicated mode.Default value: Open
     * 
     */
    public Output<String> matchCriteria() {
        return this.matchCriteria;
    }
    /**
     * The payment type of the resource
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * platform of the capacity reservation, value range `windows`, `linux`.
     * 
     */
    @Export(name="platform", refs={String.class}, tree="[0]")
    private Output<String> platform;

    /**
     * @return platform of the capacity reservation, value range `windows`, `linux`.
     * 
     */
    public Output<String> platform() {
        return this.platform;
    }
    /**
     * The resource group id.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The resource group id.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * time of the capacity reservation which become active.
     * 
     */
    @Export(name="startTime", refs={String.class}, tree="[0]")
    private Output<String> startTime;

    /**
     * @return time of the capacity reservation which become active.
     * 
     */
    public Output<String> startTime() {
        return this.startTime;
    }
    /**
     * The capacity is scheduled to take effect. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
     * 
     */
    @Export(name="startTimeType", refs={String.class}, tree="[0]")
    private Output<String> startTimeType;

    /**
     * @return The capacity is scheduled to take effect. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
     * 
     */
    public Output<String> startTimeType() {
        return this.startTimeType;
    }
    /**
     * The status of the capacity reservation.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the capacity reservation.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tag of the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * This parameter is under test and is not yet open for use.
     * 
     */
    @Export(name="timeSlot", refs={String.class}, tree="[0]")
    private Output<String> timeSlot;

    /**
     * @return This parameter is under test and is not yet open for use.
     * 
     */
    public Output<String> timeSlot() {
        return this.timeSlot;
    }
    /**
     * The ID of the zone in the region to which the capacity reservation service belongs. Currently, it is only supported to create a capacity reservation service in one zone.
     * 
     */
    @Export(name="zoneIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> zoneIds;

    /**
     * @return The ID of the zone in the region to which the capacity reservation service belongs. Currently, it is only supported to create a capacity reservation service in one zone.
     * 
     */
    public Output<List<String>> zoneIds() {
        return this.zoneIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CapacityReservation(String name) {
        this(name, CapacityReservationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CapacityReservation(String name, CapacityReservationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CapacityReservation(String name, CapacityReservationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/capacityReservation:CapacityReservation", name, args == null ? CapacityReservationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CapacityReservation(String name, Output<String> id, @Nullable CapacityReservationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/capacityReservation:CapacityReservation", name, state, makeResourceOptions(options, id));
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
    public static CapacityReservation get(String name, Output<String> id, @Nullable CapacityReservationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CapacityReservation(name, id, state, options);
    }
}
