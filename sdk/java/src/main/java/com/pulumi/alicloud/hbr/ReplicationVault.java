// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.hbr.ReplicationVaultArgs;
import com.pulumi.alicloud.hbr.inputs.ReplicationVaultState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Hybrid Backup Recovery (HBR) Replication Vault resource.
 * 
 * For information about Hybrid Backup Recovery (HBR) Replication Vault and how to use it, see [What is Replication Vault](https://www.alibabacloud.com/help/en/doc-detail/345603.html).
 * 
 * &gt; **NOTE:** Available in v1.152.0+.
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
 * import com.pulumi.alicloud.hbr.HbrFunctions;
 * import com.pulumi.alicloud.hbr.inputs.GetReplicationVaultRegionsArgs;
 * import com.pulumi.random.integer;
 * import com.pulumi.random.IntegerArgs;
 * import com.pulumi.alicloud.hbr.Vault;
 * import com.pulumi.alicloud.hbr.VaultArgs;
 * import com.pulumi.alicloud.hbr.ReplicationVault;
 * import com.pulumi.alicloud.hbr.ReplicationVaultArgs;
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
 *         final var sourceRegion = config.get(&#34;sourceRegion&#34;).orElse(&#34;cn-hangzhou&#34;);
 *         final var default = HbrFunctions.getReplicationVaultRegions();
 * 
 *         var defaultInteger = new Integer(&#34;defaultInteger&#34;, IntegerArgs.builder()        
 *             .min(10000)
 *             .max(99999)
 *             .build());
 * 
 *         var defaultVault = new Vault(&#34;defaultVault&#34;, VaultArgs.builder()        
 *             .vaultName(String.format(&#34;terraform-example-%s&#34;, defaultInteger.result()))
 *             .build());
 * 
 *         var defaultReplicationVault = new ReplicationVault(&#34;defaultReplicationVault&#34;, ReplicationVaultArgs.builder()        
 *             .replicationSourceRegionId(sourceRegion)
 *             .replicationSourceVaultId(defaultVault.id())
 *             .vaultName(&#34;terraform-example&#34;)
 *             .vaultStorageClass(&#34;STANDARD&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Hybrid Backup Recovery (HBR) Replication Vault can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:hbr/replicationVault:ReplicationVault example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:hbr/replicationVault:ReplicationVault")
public class ReplicationVault extends com.pulumi.resources.CustomResource {
    /**
     * The description of the backup vault. The description must be 0 to 255 characters in length.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the backup vault. The description must be 0 to 255 characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of the region where the source vault resides.
     * 
     */
    @Export(name="replicationSourceRegionId", refs={String.class}, tree="[0]")
    private Output<String> replicationSourceRegionId;

    /**
     * @return The ID of the region where the source vault resides.
     * 
     */
    public Output<String> replicationSourceRegionId() {
        return this.replicationSourceRegionId;
    }
    /**
     * The ID of the source vault.
     * 
     */
    @Export(name="replicationSourceVaultId", refs={String.class}, tree="[0]")
    private Output<String> replicationSourceVaultId;

    /**
     * @return The ID of the source vault.
     * 
     */
    public Output<String> replicationSourceVaultId() {
        return this.replicationSourceVaultId;
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The name of the backup vault. The name must be 1 to 64 characters in length.
     * 
     */
    @Export(name="vaultName", refs={String.class}, tree="[0]")
    private Output<String> vaultName;

    /**
     * @return The name of the backup vault. The name must be 1 to 64 characters in length.
     * 
     */
    public Output<String> vaultName() {
        return this.vaultName;
    }
    /**
     * The storage type of the backup vault. Valid values: `STANDARD`.
     * 
     */
    @Export(name="vaultStorageClass", refs={String.class}, tree="[0]")
    private Output<String> vaultStorageClass;

    /**
     * @return The storage type of the backup vault. Valid values: `STANDARD`.
     * 
     */
    public Output<String> vaultStorageClass() {
        return this.vaultStorageClass;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReplicationVault(String name) {
        this(name, ReplicationVaultArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReplicationVault(String name, ReplicationVaultArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReplicationVault(String name, ReplicationVaultArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/replicationVault:ReplicationVault", name, args == null ? ReplicationVaultArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ReplicationVault(String name, Output<String> id, @Nullable ReplicationVaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/replicationVault:ReplicationVault", name, state, makeResourceOptions(options, id));
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
    public static ReplicationVault get(String name, Output<String> id, @Nullable ReplicationVaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReplicationVault(name, id, state, options);
    }
}
