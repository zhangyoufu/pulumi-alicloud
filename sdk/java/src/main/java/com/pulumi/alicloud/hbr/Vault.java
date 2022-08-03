// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.hbr.VaultArgs;
import com.pulumi.alicloud.hbr.inputs.VaultState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a HBR Backup vault resource.
 * 
 * For information about HBR Backup vault and how to use it, see [What is Backup vault](https://www.alibabacloud.com/help/doc-detail/62362.htm).
 * 
 * &gt; **NOTE:** Available in v1.129.0+.
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
 * import com.pulumi.alicloud.hbr.Vault;
 * import com.pulumi.alicloud.hbr.VaultArgs;
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
 *         var example = new Vault(&#34;example&#34;, VaultArgs.builder()        
 *             .vaultName(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * HBR Vault can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:hbr/vault:Vault example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:hbr/vault:Vault")
public class Vault extends com.pulumi.resources.CustomResource {
    /**
     * The description of Vault. Defaults to an empty string.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of Vault. Defaults to an empty string.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Source Encryption Type，It is valid only when vault_type is `STANDARD` or `OTS_BACKUP`. Valid values: `HBR_PRIVATE`,`KMS`. Defaults to `HBR_PRIVATE`.
     * - `HBR_PRIVATE`: HBR is fully hosted, uses the backup service&#39;s own encryption method.
     * - `KMS`: Use Alibaba Cloud Kms to encryption.
     * 
     */
    @Export(name="encryptType", type=String.class, parameters={})
    private Output<String> encryptType;

    /**
     * @return Source Encryption Type，It is valid only when vault_type is `STANDARD` or `OTS_BACKUP`. Valid values: `HBR_PRIVATE`,`KMS`. Defaults to `HBR_PRIVATE`.
     * - `HBR_PRIVATE`: HBR is fully hosted, uses the backup service&#39;s own encryption method.
     * - `KMS`: Use Alibaba Cloud Kms to encryption.
     * 
     */
    public Output<String> encryptType() {
        return this.encryptType;
    }
    /**
     * The key id or alias name of Alibaba Cloud Kms. It is required and valid only when encrypt_type is `KMS`.
     * 
     */
    @Export(name="kmsKeyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return The key id or alias name of Alibaba Cloud Kms. It is required and valid only when encrypt_type is `KMS`.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * The redundancy type of the vault. Valid values: `LRS`, and `ZRS`. Defaults to `LRS`.
     * - `LRS`: means locally redundant storage, data will be stored on different storage devices in the same zone.
     * - `ZRS`: means zone-redundant storage, the data will be replicated across three storage clusters in a single region. Each storage cluster is physically separated but within the same region.
     * 
     */
    @Export(name="redundancyType", type=String.class, parameters={})
    private Output<String> redundancyType;

    /**
     * @return The redundancy type of the vault. Valid values: `LRS`, and `ZRS`. Defaults to `LRS`.
     * - `LRS`: means locally redundant storage, data will be stored on different storage devices in the same zone.
     * - `ZRS`: means zone-redundant storage, the data will be replicated across three storage clusters in a single region. Each storage cluster is physically separated but within the same region.
     * 
     */
    public Output<String> redundancyType() {
        return this.redundancyType;
    }
    /**
     * The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The name of Vault.
     * 
     */
    @Export(name="vaultName", type=String.class, parameters={})
    private Output<String> vaultName;

    /**
     * @return The name of Vault.
     * 
     */
    public Output<String> vaultName() {
        return this.vaultName;
    }
    /**
     * The storage class of Vault. Valid values: `STANDARD`.
     * 
     */
    @Export(name="vaultStorageClass", type=String.class, parameters={})
    private Output<String> vaultStorageClass;

    /**
     * @return The storage class of Vault. Valid values: `STANDARD`.
     * 
     */
    public Output<String> vaultStorageClass() {
        return this.vaultStorageClass;
    }
    /**
     * The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
     * 
     */
    @Export(name="vaultType", type=String.class, parameters={})
    private Output<String> vaultType;

    /**
     * @return The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
     * 
     */
    public Output<String> vaultType() {
        return this.vaultType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vault(String name) {
        this(name, VaultArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vault(String name, VaultArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vault(String name, VaultArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/vault:Vault", name, args == null ? VaultArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vault(String name, Output<String> id, @Nullable VaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbr/vault:Vault", name, state, makeResourceOptions(options, id));
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
    public static Vault get(String name, Output<String> id, @Nullable VaultState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vault(name, id, state, options);
    }
}
