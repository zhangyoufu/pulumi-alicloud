// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cddc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cddc.DedicatedHostAccountArgs;
import com.pulumi.alicloud.cddc.inputs.DedicatedHostAccountState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ApsaraDB for MyBase Dedicated Host Account resource.
 * 
 * For information about ApsaraDB for MyBase Dedicated Host Account and how to use it, see [What is Dedicated Host Account](https://www.alibabacloud.com/help/en/apsaradb-for-mybase/latest/creatededicatedhostaccount).
 * 
 * &gt; **NOTE:** Available since v1.148.0.
 * 
 * &gt; **NOTE:** Each Dedicated host can have only one account. Before you create an account for a host, make sure that the existing account is deleted.
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
 * import com.pulumi.alicloud.cddc.CddcFunctions;
 * import com.pulumi.alicloud.cddc.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.cddc.DedicatedHostGroup;
 * import com.pulumi.alicloud.cddc.DedicatedHostGroupArgs;
 * import com.pulumi.alicloud.cddc.inputs.GetHostEcsLevelInfosArgs;
 * import com.pulumi.alicloud.cddc.DedicatedHost;
 * import com.pulumi.alicloud.cddc.DedicatedHostArgs;
 * import com.pulumi.alicloud.cddc.DedicatedHostAccount;
 * import com.pulumi.alicloud.cddc.DedicatedHostAccountArgs;
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
 *         final var default = CddcFunctions.getZones();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.ids()[0])
 *             .build());
 * 
 *         var defaultDedicatedHostGroup = new DedicatedHostGroup(&#34;defaultDedicatedHostGroup&#34;, DedicatedHostGroupArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .cpuAllocationRatio(101)
 *             .memAllocationRatio(50)
 *             .diskAllocationRatio(200)
 *             .allocationPolicy(&#34;Evenly&#34;)
 *             .hostReplacePolicy(&#34;Manual&#34;)
 *             .dedicatedHostGroupDesc(name)
 *             .openPermission(true)
 *             .build());
 * 
 *         final var defaultGetHostEcsLevelInfos = CddcFunctions.getHostEcsLevelInfos(GetHostEcsLevelInfosArgs.builder()
 *             .dbType(&#34;mysql&#34;)
 *             .zoneId(default_.ids()[0])
 *             .storageType(&#34;cloud_essd&#34;)
 *             .build());
 * 
 *         var defaultDedicatedHost = new DedicatedHost(&#34;defaultDedicatedHost&#34;, DedicatedHostArgs.builder()        
 *             .hostName(name)
 *             .dedicatedHostGroupId(defaultDedicatedHostGroup.id())
 *             .hostClass(defaultGetHostEcsLevelInfos.applyValue(getHostEcsLevelInfosResult -&gt; getHostEcsLevelInfosResult.infos()[0].resClassCode()))
 *             .zoneId(default_.ids()[0])
 *             .vswitchId(defaultSwitch.id())
 *             .paymentType(&#34;Subscription&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;CDDC_DEDICATED&#34;)
 *             ))
 *             .build());
 * 
 *         var defaultDedicatedHostAccount = new DedicatedHostAccount(&#34;defaultDedicatedHostAccount&#34;, DedicatedHostAccountArgs.builder()        
 *             .accountName(name)
 *             .accountPassword(&#34;Password1234&#34;)
 *             .dedicatedHostId(defaultDedicatedHost.dedicatedHostId())
 *             .accountType(&#34;Normal&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ApsaraDB for MyBase Dedicated Host Account can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount example &lt;dedicated_host_id&gt;:&lt;account_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount")
public class DedicatedHostAccount extends com.pulumi.resources.CustomResource {
    /**
     * The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     * 
     */
    @Export(name="accountName", refs={String.class}, tree="[0]")
    private Output<String> accountName;

    /**
     * @return The name of the Dedicated host account. The account name must be 2 to 16 characters in length, contain lower case letters, digits, and underscore(_). At the same time, the name must start with a letter and end with a letter or number.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&amp;*()_+-=`.
     * 
     */
    @Export(name="accountPassword", refs={String.class}, tree="[0]")
    private Output<String> accountPassword;

    /**
     * @return The password of the Dedicated host account. The account password must be 6 to 32 characters in length, and can contain letters, digits, and special characters `!@#$%^&amp;*()_+-=`.
     * 
     */
    public Output<String> accountPassword() {
        return this.accountPassword;
    }
    /**
     * The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     * 
     */
    @Export(name="accountType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountType;

    /**
     * @return The type of the Dedicated host account. Valid values: `Admin`, `Normal`.
     * 
     */
    public Output<Optional<String>> accountType() {
        return Codegen.optional(this.accountType);
    }
    /**
     * The ID of Dedicated the host.
     * 
     */
    @Export(name="dedicatedHostId", refs={String.class}, tree="[0]")
    private Output<String> dedicatedHostId;

    /**
     * @return The ID of Dedicated the host.
     * 
     */
    public Output<String> dedicatedHostId() {
        return this.dedicatedHostId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DedicatedHostAccount(String name) {
        this(name, DedicatedHostAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DedicatedHostAccount(String name, DedicatedHostAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DedicatedHostAccount(String name, DedicatedHostAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount", name, args == null ? DedicatedHostAccountArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DedicatedHostAccount(String name, Output<String> id, @Nullable DedicatedHostAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cddc/dedicatedHostAccount:DedicatedHostAccount", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static DedicatedHostAccount get(String name, Output<String> id, @Nullable DedicatedHostAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DedicatedHostAccount(name, id, state, options);
    }
}
