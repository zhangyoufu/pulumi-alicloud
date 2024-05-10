// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.polardb.AccountPrivilegeArgs;
import com.pulumi.alicloud.polardb.inputs.AccountPrivilegeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a PolarDB account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.
 * 
 * &gt; **NOTE:** Available in v1.67.0+.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.polardb.PolardbFunctions;
 * import com.pulumi.alicloud.polardb.inputs.GetNodeClassesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.polardb.Cluster;
 * import com.pulumi.alicloud.polardb.ClusterArgs;
 * import com.pulumi.alicloud.polardb.Account;
 * import com.pulumi.alicloud.polardb.AccountArgs;
 * import com.pulumi.alicloud.polardb.Database;
 * import com.pulumi.alicloud.polardb.DatabaseArgs;
 * import com.pulumi.alicloud.polardb.AccountPrivilege;
 * import com.pulumi.alicloud.polardb.AccountPrivilegeArgs;
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
 *         final var default = PolardbFunctions.getNodeClasses(GetNodeClassesArgs.builder()
 *             .dbType("MySQL")
 *             .dbVersion("8.0")
 *             .payType("PostPaid")
 *             .build());
 * 
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()        
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.16.0.0/16")
 *             .build());
 * 
 *         var defaultSwitch = new Switch("defaultSwitch", SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock("172.16.0.0/24")
 *             .zoneId(default_.classes()[0].zoneId())
 *             .vswitchName("terraform-example")
 *             .build());
 * 
 *         var defaultCluster = new Cluster("defaultCluster", ClusterArgs.builder()        
 *             .dbType("MySQL")
 *             .dbVersion("8.0")
 *             .dbNodeClass(default_.classes()[0].supportedEngines()[0].availableResources()[0].dbNodeClass())
 *             .payType("PostPaid")
 *             .vswitchId(defaultSwitch.id())
 *             .description("terraform-example")
 *             .build());
 * 
 *         var defaultAccount = new Account("defaultAccount", AccountArgs.builder()        
 *             .dbClusterId(defaultCluster.id())
 *             .accountName("terraform_example")
 *             .accountPassword("Example1234")
 *             .accountDescription("terraform-example")
 *             .build());
 * 
 *         var defaultDatabase = new Database("defaultDatabase", DatabaseArgs.builder()        
 *             .dbClusterId(defaultCluster.id())
 *             .dbName("terraform-example")
 *             .build());
 * 
 *         var defaultAccountPrivilege = new AccountPrivilege("defaultAccountPrivilege", AccountPrivilegeArgs.builder()        
 *             .dbClusterId(defaultCluster.id())
 *             .accountName(defaultAccount.accountName())
 *             .accountPrivilege("ReadOnly")
 *             .dbNames(defaultDatabase.dbName())
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
 * PolarDB account privilege can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:polardb/accountPrivilege:AccountPrivilege example &#34;pc-12345:tf_account:ReadOnly&#34;
 * ```
 * 
 */
@ResourceType(type="alicloud:polardb/accountPrivilege:AccountPrivilege")
public class AccountPrivilege extends com.pulumi.resources.CustomResource {
    /**
     * A specified account name.
     * 
     */
    @Export(name="accountName", refs={String.class}, tree="[0]")
    private Output<String> accountName;

    /**
     * @return A specified account name.
     * 
     */
    public Output<String> accountName() {
        return this.accountName;
    }
    /**
     * The privilege of one account access database. Valid values: [&#34;ReadOnly&#34;, &#34;ReadWrite&#34;], [&#34;DMLOnly&#34;, &#34;DDLOnly&#34;] added since version v1.101.0. Default to &#34;ReadOnly&#34;.
     * 
     */
    @Export(name="accountPrivilege", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountPrivilege;

    /**
     * @return The privilege of one account access database. Valid values: [&#34;ReadOnly&#34;, &#34;ReadWrite&#34;], [&#34;DMLOnly&#34;, &#34;DDLOnly&#34;] added since version v1.101.0. Default to &#34;ReadOnly&#34;.
     * 
     */
    public Output<Optional<String>> accountPrivilege() {
        return Codegen.optional(this.accountPrivilege);
    }
    /**
     * The Id of cluster in which account belongs.
     * 
     */
    @Export(name="dbClusterId", refs={String.class}, tree="[0]")
    private Output<String> dbClusterId;

    /**
     * @return The Id of cluster in which account belongs.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * List of specified database name.
     * 
     */
    @Export(name="dbNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> dbNames;

    /**
     * @return List of specified database name.
     * 
     */
    public Output<List<String>> dbNames() {
        return this.dbNames;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccountPrivilege(String name) {
        this(name, AccountPrivilegeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccountPrivilege(String name, AccountPrivilegeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccountPrivilege(String name, AccountPrivilegeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/accountPrivilege:AccountPrivilege", name, args == null ? AccountPrivilegeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccountPrivilege(String name, Output<String> id, @Nullable AccountPrivilegeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:polardb/accountPrivilege:AccountPrivilege", name, state, makeResourceOptions(options, id));
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
    public static AccountPrivilege get(String name, Output<String> id, @Nullable AccountPrivilegeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccountPrivilege(name, id, state, options);
    }
}
