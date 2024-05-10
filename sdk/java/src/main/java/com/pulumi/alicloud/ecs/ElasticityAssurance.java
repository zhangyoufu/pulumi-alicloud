// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.ElasticityAssuranceArgs;
import com.pulumi.alicloud.ecs.inputs.ElasticityAssuranceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Ecs Elasticity Assurance resource.
 * 
 * For information about Ecs Elasticity Assurance and how to use it, see [What is Elasticity Assurance](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/createelasticityassurance).
 * 
 * &gt; **NOTE:** Available in v1.196.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.ElasticityAssurance;
 * import com.pulumi.alicloud.ecs.ElasticityAssuranceArgs;
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
 *         final var default = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status("OK")
 *             .build());
 * 
 *         final var defaultGetZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("Instance")
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .instanceTypeFamily("ecs.c6")
 *             .build());
 * 
 *         var defaultElasticityAssurance = new ElasticityAssurance("defaultElasticityAssurance", ElasticityAssuranceArgs.builder()        
 *             .instanceAmount(1)
 *             .description("before")
 *             .zoneIds(defaultGetZones.applyValue(getZonesResult -> getZonesResult.zones()[0].id()))
 *             .privatePoolOptionsName("test_before")
 *             .period(1)
 *             .privatePoolOptionsMatchCriteria("Open")
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -> getInstanceTypesResult.instanceTypes()[0].id()))
 *             .periodUnit("Month")
 *             .assuranceTimes("Unlimited")
 *             .resourceGroupId(default_.ids()[0])
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Ecs Elasticity Assurance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/elasticityAssurance:ElasticityAssurance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/elasticityAssurance:ElasticityAssurance")
public class ElasticityAssurance extends com.pulumi.resources.CustomResource {
    /**
     * The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective duration. Default value: Unlimited.
     * 
     */
    @Export(name="assuranceTimes", refs={String.class}, tree="[0]")
    private Output<String> assuranceTimes;

    /**
     * @return The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective duration. Default value: Unlimited.
     * 
     */
    public Output<String> assuranceTimes() {
        return this.assuranceTimes;
    }
    /**
     * Description of flexible guarantee service.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of flexible guarantee service.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The first ID of the resource
     * 
     */
    @Export(name="elasticityAssuranceId", refs={String.class}, tree="[0]")
    private Output<String> elasticityAssuranceId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> elasticityAssuranceId() {
        return this.elasticityAssuranceId;
    }
    /**
     * Flexible guarantee service failure time.
     * 
     */
    @Export(name="endTime", refs={String.class}, tree="[0]")
    private Output<String> endTime;

    /**
     * @return Flexible guarantee service failure time.
     * 
     */
    public Output<String> endTime() {
        return this.endTime;
    }
    /**
     * The total number of instances for which to reserve the capacity of an instance type. Valid values: 1 to 1000.
     * 
     */
    @Export(name="instanceAmount", refs={Integer.class}, tree="[0]")
    private Output<Integer> instanceAmount;

    /**
     * @return The total number of instances for which to reserve the capacity of an instance type. Valid values: 1 to 1000.
     * 
     */
    public Output<Integer> instanceAmount() {
        return this.instanceAmount;
    }
    /**
     * The billing method of the instance. Possible value: PostPaid. Currently, only pay-as-you-go is supported.
     * 
     */
    @Export(name="instanceChargeType", refs={String.class}, tree="[0]")
    private Output<String> instanceChargeType;

    /**
     * @return The billing method of the instance. Possible value: PostPaid. Currently, only pay-as-you-go is supported.
     * 
     */
    public Output<String> instanceChargeType() {
        return this.instanceChargeType;
    }
    /**
     * Instance type. Currently, only one instance type is supported.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    /**
     * @return Instance type. Currently, only one instance type is supported.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * Length of purchase. The unit of duration is determined by the &#39;period_unit&#39; parameter. Default value: 1.
     * - When the `period_unit` parameter is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
     * - When the `period_unit` parameter is set to Year, the valid values are 1, 2, 3, 4, and 5.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return Length of purchase. The unit of duration is determined by the &#39;period_unit&#39; parameter. Default value: 1.
     * - When the `period_unit` parameter is set to Month, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
     * - When the `period_unit` parameter is set to Year, the valid values are 1, 2, 3, 4, and 5.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * Duration unit. Value range:-Month: Month-Year: YearDefault value: Year
     * 
     */
    @Export(name="periodUnit", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> periodUnit;

    /**
     * @return Duration unit. Value range:-Month: Month-Year: YearDefault value: Year
     * 
     */
    public Output<Optional<String>> periodUnit() {
        return Codegen.optional(this.periodUnit);
    }
    /**
     * The matching mode of flexible guarantee service. Possible values:-Open: flexible guarantee service for Open mode.-Target: specifies the flexible guarantee service of the mode.
     * 
     */
    @Export(name="privatePoolOptionsMatchCriteria", refs={String.class}, tree="[0]")
    private Output<String> privatePoolOptionsMatchCriteria;

    /**
     * @return The matching mode of flexible guarantee service. Possible values:-Open: flexible guarantee service for Open mode.-Target: specifies the flexible guarantee service of the mode.
     * 
     */
    public Output<String> privatePoolOptionsMatchCriteria() {
        return this.privatePoolOptionsMatchCriteria;
    }
    /**
     * The name of the flexible protection service.
     * 
     */
    @Export(name="privatePoolOptionsName", refs={String.class}, tree="[0]")
    private Output<String> privatePoolOptionsName;

    /**
     * @return The name of the flexible protection service.
     * 
     */
    public Output<String> privatePoolOptionsName() {
        return this.privatePoolOptionsName;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * Flexible guarantee service effective time.
     * 
     */
    @Export(name="startTime", refs={String.class}, tree="[0]")
    private Output<String> startTime;

    /**
     * @return Flexible guarantee service effective time.
     * 
     */
    public Output<String> startTime() {
        return this.startTime;
    }
    /**
     * Flexible guarantee effective way. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
     * 
     */
    @Export(name="startTimeType", refs={String.class}, tree="[0]")
    private Output<String> startTimeType;

    /**
     * @return Flexible guarantee effective way. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
     * 
     */
    public Output<String> startTimeType() {
        return this.startTimeType;
    }
    /**
     * The status of flexible guarantee services. Possible values:-Preparing: in preparation.-Prepared: to take effect.-Active: in effect.-Released: Released.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of flexible guarantee services. Possible values:-Preparing: in preparation.-Prepared: to take effect.-Active: in effect.-Released: Released.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tag key-value pair information bound by the elastic guarantee service.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tag key-value pair information bound by the elastic guarantee service.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * This parameter is not yet available.
     * 
     */
    @Export(name="usedAssuranceTimes", refs={Integer.class}, tree="[0]")
    private Output<Integer> usedAssuranceTimes;

    /**
     * @return This parameter is not yet available.
     * 
     */
    public Output<Integer> usedAssuranceTimes() {
        return this.usedAssuranceTimes;
    }
    /**
     * The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
     * 
     */
    @Export(name="zoneIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> zoneIds;

    /**
     * @return The zone ID of the region to which the elastic Protection Service belongs. Currently, only the creation of flexible protection services in one available area is supported.
     * 
     */
    public Output<List<String>> zoneIds() {
        return this.zoneIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ElasticityAssurance(String name) {
        this(name, ElasticityAssuranceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ElasticityAssurance(String name, ElasticityAssuranceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ElasticityAssurance(String name, ElasticityAssuranceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/elasticityAssurance:ElasticityAssurance", name, args == null ? ElasticityAssuranceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ElasticityAssurance(String name, Output<String> id, @Nullable ElasticityAssuranceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/elasticityAssurance:ElasticityAssurance", name, state, makeResourceOptions(options, id));
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
    public static ElasticityAssurance get(String name, Output<String> id, @Nullable ElasticityAssuranceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ElasticityAssurance(name, id, state, options);
    }
}
