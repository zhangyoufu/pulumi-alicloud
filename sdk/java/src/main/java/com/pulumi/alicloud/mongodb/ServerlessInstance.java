// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mongodb.ServerlessInstanceArgs;
import com.pulumi.alicloud.mongodb.inputs.ServerlessInstanceState;
import com.pulumi.alicloud.mongodb.outputs.ServerlessInstanceSecurityIpGroup;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a MongoDB Serverless Instance resource.
 * 
 * For information about MongoDB Serverless Instance and how to use it, see [What is Serverless Instance](https://www.alibabacloud.com/help/doc-detail/26558.html).
 * 
 * &gt; **NOTE:** Deprecated since v1.214.0.
 * 
 * &gt; **DEPRECATED:**  This resource has been deprecated from version `1.214.0`.
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
 * import com.pulumi.alicloud.mongodb.MongodbFunctions;
 * import com.pulumi.alicloud.mongodb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.mongodb.ServerlessInstance;
 * import com.pulumi.alicloud.mongodb.ServerlessInstanceArgs;
 * import com.pulumi.alicloud.mongodb.inputs.ServerlessInstanceSecurityIpGroupArgs;
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
 *         final var default = MongodbFunctions.getZones();
 * 
 *         final var defaultGetNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex("default-NODELETING")
 *             .build());
 * 
 *         final var defaultGetSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultGetNetworks.applyValue(getNetworksResult -> getNetworksResult.ids()[0]))
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         final var defaultGetResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var example = new ServerlessInstance("example", ServerlessInstanceArgs.builder()
 *             .accountPassword("Abc12345")
 *             .dbInstanceDescription("example_value")
 *             .dbInstanceStorage(5)
 *             .storageEngine("WiredTiger")
 *             .capacityUnit(100)
 *             .engine("MongoDB")
 *             .resourceGroupId(defaultGetResourceGroups.applyValue(getResourceGroupsResult -> getResourceGroupsResult.groups()[0].id()))
 *             .engineVersion("4.2")
 *             .period(1)
 *             .periodPriceType("Month")
 *             .vpcId(defaultGetNetworks.applyValue(getNetworksResult -> getNetworksResult.ids()[0]))
 *             .zoneId(default_.zones()[0].id())
 *             .vswitchId(defaultGetSwitches.applyValue(getSwitchesResult -> getSwitchesResult.ids()[0]))
 *             .tags(Map.ofEntries(
 *                 Map.entry("Created", "MongodbServerlessInstance"),
 *                 Map.entry("For", "TF")
 *             ))
 *             .securityIpGroups(ServerlessInstanceSecurityIpGroupArgs.builder()
 *                 .securityIpGroupAttribute("example_value")
 *                 .securityIpGroupName("example_value")
 *                 .securityIpList("192.168.0.1")
 *                 .build())
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
 * MongoDB Serverless Instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mongodb/serverlessInstance:ServerlessInstance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:mongodb/serverlessInstance:ServerlessInstance")
public class ServerlessInstance extends com.pulumi.resources.CustomResource {
    /**
     * The password of the database logon account.
     * * The password length is `8` to `32` bits.
     * * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%^&amp;*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
     * 
     */
    @Export(name="accountPassword", refs={String.class}, tree="[0]")
    private Output<String> accountPassword;

    /**
     * @return The password of the database logon account.
     * * The password length is `8` to `32` bits.
     * * The password consists of at least any three of uppercase letters, lowercase letters, numbers, and special characters. The special character is `!#$%^&amp;*()_+-=`. The MongoDB Serverless instance provides a default database login account. This account cannot be modified. You can only set or modify the password for this account.
     * 
     */
    public Output<String> accountPassword() {
        return this.accountPassword;
    }
    /**
     * Set whether the instance is automatically renewed.
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Set whether the instance is automatically renewed.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
     * 
     */
    @Export(name="capacityUnit", refs={Integer.class}, tree="[0]")
    private Output<Integer> capacityUnit;

    /**
     * @return The I/O throughput consumed by the instance. Valid values: `100` to `8000`.
     * 
     */
    public Output<Integer> capacityUnit() {
        return this.capacityUnit;
    }
    /**
     * The db instance description.
     * 
     */
    @Export(name="dbInstanceDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dbInstanceDescription;

    /**
     * @return The db instance description.
     * 
     */
    public Output<Optional<String>> dbInstanceDescription() {
        return Codegen.optional(this.dbInstanceDescription);
    }
    /**
     * The db instance storage. Valid values: `1` to `100`.
     * 
     */
    @Export(name="dbInstanceStorage", refs={Integer.class}, tree="[0]")
    private Output<Integer> dbInstanceStorage;

    /**
     * @return The db instance storage. Valid values: `1` to `100`.
     * 
     */
    public Output<Integer> dbInstanceStorage() {
        return this.dbInstanceStorage;
    }
    /**
     * The database engine of the instance. Valid values: `MongoDB`.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output<String> engine;

    /**
     * @return The database engine of the instance. Valid values: `MongoDB`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * The database version number. Valid values: `4.2`.
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return The database version number. Valid values: `4.2`.
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintain_start_time` is `01:00Z`, `maintain_end_time` must be `02:00Z`.
     * 
     */
    @Export(name="maintainEndTime", refs={String.class}, tree="[0]")
    private Output<String> maintainEndTime;

    /**
     * @return The end time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC. **NOTE:** The difference between the start time and end time must be one hour. For example, if `maintain_start_time` is `01:00Z`, `maintain_end_time` must be `02:00Z`.
     * 
     */
    public Output<String> maintainEndTime() {
        return this.maintainEndTime;
    }
    /**
     * The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
     * 
     */
    @Export(name="maintainStartTime", refs={String.class}, tree="[0]")
    private Output<String> maintainStartTime;

    /**
     * @return The start time of the maintenance window. Specify the time in the `HH:mmZ` format. The time must be in UTC.
     * 
     */
    public Output<String> maintainStartTime() {
        return this.maintainStartTime;
    }
    /**
     * The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The purchase duration of the instance, in months. Valid values: `1` to `9`, `12`, `24`, `36`, `60`.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The period price type. Valid values: `Day`, `Month`.
     * 
     */
    @Export(name="periodPriceType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> periodPriceType;

    /**
     * @return The period price type. Valid values: `Day`, `Month`.
     * 
     */
    public Output<Optional<String>> periodPriceType() {
        return Codegen.optional(this.periodPriceType);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * An array that consists of the information of IP whitelists.
     * 
     */
    @Export(name="securityIpGroups", refs={List.class,ServerlessInstanceSecurityIpGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ServerlessInstanceSecurityIpGroup>> securityIpGroups;

    /**
     * @return An array that consists of the information of IP whitelists.
     * 
     */
    public Output<Optional<List<ServerlessInstanceSecurityIpGroup>>> securityIpGroups() {
        return Codegen.optional(this.securityIpGroups);
    }
    /**
     * The instance status. For more information, see the instance Status Table.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The instance status. For more information, see the instance Status Table.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The storage engine used by the instance. Valid values: `WiredTiger`.
     * 
     */
    @Export(name="storageEngine", refs={String.class}, tree="[0]")
    private Output<String> storageEngine;

    /**
     * @return The storage engine used by the instance. Valid values: `WiredTiger`.
     * 
     */
    public Output<String> storageEngine() {
        return this.storageEngine;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The ID of the VPC network.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC network.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The of the vswitch.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The of the vswitch.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The ID of the zone. Use this parameter to specify the zone created by the instance.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The ID of the zone. Use this parameter to specify the zone created by the instance.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServerlessInstance(java.lang.String name) {
        this(name, ServerlessInstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerlessInstance(java.lang.String name, ServerlessInstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerlessInstance(java.lang.String name, ServerlessInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/serverlessInstance:ServerlessInstance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServerlessInstance(java.lang.String name, Output<java.lang.String> id, @Nullable ServerlessInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/serverlessInstance:ServerlessInstance", name, state, makeResourceOptions(options, id), false);
    }

    private static ServerlessInstanceArgs makeArgs(ServerlessInstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServerlessInstanceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accountPassword"
            ))
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
    public static ServerlessInstance get(java.lang.String name, Output<java.lang.String> id, @Nullable ServerlessInstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerlessInstance(name, id, state, options);
    }
}
