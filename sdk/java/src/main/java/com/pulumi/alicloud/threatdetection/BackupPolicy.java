// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.BackupPolicyArgs;
import com.pulumi.alicloud.threatdetection.inputs.BackupPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Backup Policy resource.
 * 
 * For information about Threat Detection Backup Policy and how to use it, see [What is Backup Policy](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
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
 * import com.pulumi.alicloud.threatdetection.ThreatdetectionFunctions;
 * import com.pulumi.alicloud.threatdetection.inputs.GetAssetsArgs;
 * import com.pulumi.alicloud.threatdetection.BackupPolicy;
 * import com.pulumi.alicloud.threatdetection.BackupPolicyArgs;
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
 *         final var defaultAssets = ThreatdetectionFunctions.getAssets(GetAssetsArgs.builder()
 *             .machineTypes(&#34;ecs&#34;)
 *             .build());
 * 
 *         var defaultBackupPolicy = new BackupPolicy(&#34;defaultBackupPolicy&#34;, BackupPolicyArgs.builder()        
 *             .backupPolicyName(&#34;tf-example-name&#34;)
 *             .policy(&#34;{\&#34;Exclude\&#34;:[\&#34;/bin/\&#34;,\&#34;/usr/bin/\&#34;,\&#34;/sbin/\&#34;,\&#34;/boot/\&#34;,\&#34;/proc/\&#34;,\&#34;/sys/\&#34;,\&#34;/srv/\&#34;,\&#34;/lib/\&#34;,\&#34;/selinux/\&#34;,\&#34;/usr/sbin/\&#34;,\&#34;/run/\&#34;,\&#34;/lib32/\&#34;,\&#34;/lib64/\&#34;,\&#34;/lost+found/\&#34;,\&#34;/var/lib/kubelet/\&#34;,\&#34;/var/lib/ntp/proc\&#34;,\&#34;/var/lib/container\&#34;],\&#34;ExcludeSystemPath\&#34;:true,\&#34;Include\&#34;:[],\&#34;IsDefault\&#34;:1,\&#34;Retention\&#34;:7,\&#34;Schedule\&#34;:\&#34;I|1668703620|PT24H\&#34;,\&#34;Source\&#34;:[],\&#34;SpeedLimiter\&#34;:\&#34;\&#34;,\&#34;UseVss\&#34;:true}&#34;)
 *             .policyVersion(&#34;2.0.0&#34;)
 *             .uuidLists(defaultAssets.applyValue(getAssetsResult -&gt; getAssetsResult.ids()[0]))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Threat Detection Backup Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/backupPolicy:BackupPolicy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/backupPolicy:BackupPolicy")
public class BackupPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Protection of the Name of the Policy.
     * 
     */
    @Export(name="backupPolicyName", refs={String.class}, tree="[0]")
    private Output<String> backupPolicyName;

    /**
     * @return Protection of the Name of the Policy.
     * 
     */
    public Output<String> backupPolicyName() {
        return this.backupPolicyName;
    }
    /**
     * The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The Specified Protection Policies of the Specific Configuration. see [how to use it](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createbackuppolicy).
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
     * 
     */
    @Export(name="policyRegionId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> policyRegionId;

    /**
     * @return The region ID of the non-Alibaba cloud server. You can call the [DescribeSupportRegion](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-describesupportregion) interface to view the region supported by anti-ransomware, and then select the region supported by anti-ransomware according to the region where your non-Alibaba cloud server is located.
     * 
     */
    public Output<Optional<String>> policyRegionId() {
        return Codegen.optional(this.policyRegionId);
    }
    /**
     * Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
     * 
     */
    @Export(name="policyVersion", refs={String.class}, tree="[0]")
    private Output<String> policyVersion;

    /**
     * @return Anti-Blackmail Policy Version. Valid values: `1.0.0`, `2.0.0`.
     * 
     */
    public Output<String> policyVersion() {
        return this.policyVersion;
    }
    /**
     * The status of the Backup Policy instance.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Backup Policy instance.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Specify the Protection of Server UUID List.
     * 
     */
    @Export(name="uuidLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> uuidLists;

    /**
     * @return Specify the Protection of Server UUID List.
     * 
     */
    public Output<List<String>> uuidLists() {
        return this.uuidLists;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackupPolicy(String name) {
        this(name, BackupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackupPolicy(String name, BackupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackupPolicy(String name, BackupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/backupPolicy:BackupPolicy", name, args == null ? BackupPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackupPolicy(String name, Output<String> id, @Nullable BackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/backupPolicy:BackupPolicy", name, state, makeResourceOptions(options, id));
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
    public static BackupPolicy get(String name, Output<String> id, @Nullable BackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackupPolicy(name, id, state, options);
    }
}
