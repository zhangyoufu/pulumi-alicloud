// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.log.AuditArgs;
import com.pulumi.alicloud.log.inputs.AuditState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * SLS log audit exists in the form of log service app.
 * 
 * In addition to inheriting all SLS functions, it also enhances the real-time automatic centralized collection of audit related logs across multi cloud products under multi accounts, and provides support for storage, query and information summary required by audit. It covers actiontrail, OSS, NAS, SLB, API gateway, RDS, WAF, cloud firewall, cloud security center and other products.
 * 
 * &gt; **NOTE:** Available since v1.81.0
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.log.Audit;
 * import com.pulumi.alicloud.log.AuditArgs;
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
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Audit("example", AuditArgs.builder()
 *             .displayName("tf-audit-example")
 *             .aliuid(default_.id())
 *             .variableMap(Map.ofEntries(
 *                 Map.entry("actiontrail_enabled", "true"),
 *                 Map.entry("actiontrail_ttl", "180"),
 *                 Map.entry("oss_access_enabled", "true"),
 *                 Map.entry("oss_access_ttl", "7"),
 *                 Map.entry("oss_sync_enabled", "true"),
 *                 Map.entry("oss_sync_ttl", "180"),
 *                 Map.entry("oss_metering_enabled", "true"),
 *                 Map.entry("oss_metering_ttl", "180"),
 *                 Map.entry("rds_enabled", "true"),
 *                 Map.entry("rds_audit_collection_policy", ""),
 *                 Map.entry("rds_ttl", "180"),
 *                 Map.entry("rds_slow_enabled", "false"),
 *                 Map.entry("rds_slow_collection_policy", ""),
 *                 Map.entry("rds_slow_ttl", "180"),
 *                 Map.entry("rds_perf_enabled", "false"),
 *                 Map.entry("rds_perf_collection_policy", ""),
 *                 Map.entry("rds_perf_ttl", "180"),
 *                 Map.entry("vpc_flow_enabled", "false"),
 *                 Map.entry("vpc_flow_ttl", "7"),
 *                 Map.entry("vpc_flow_collection_policy", ""),
 *                 Map.entry("vpc_sync_enabled", "true"),
 *                 Map.entry("vpc_sync_ttl", "180"),
 *                 Map.entry("polardb_enabled", "true"),
 *                 Map.entry("polardb_audit_collection_policy", ""),
 *                 Map.entry("polardb_ttl", "180"),
 *                 Map.entry("polardb_slow_enabled", "false"),
 *                 Map.entry("polardb_slow_collection_policy", ""),
 *                 Map.entry("polardb_slow_ttl", "180"),
 *                 Map.entry("polardb_perf_enabled", "false"),
 *                 Map.entry("polardb_perf_collection_policy", ""),
 *                 Map.entry("polardb_perf_ttl", "180"),
 *                 Map.entry("drds_audit_enabled", "true"),
 *                 Map.entry("drds_audit_collection_policy", ""),
 *                 Map.entry("drds_audit_ttl", "7"),
 *                 Map.entry("drds_sync_enabled", "true"),
 *                 Map.entry("drds_sync_ttl", "180"),
 *                 Map.entry("slb_access_enabled", "true"),
 *                 Map.entry("slb_access_collection_policy", ""),
 *                 Map.entry("slb_access_ttl", "7"),
 *                 Map.entry("slb_sync_enabled", "true"),
 *                 Map.entry("slb_sync_ttl", "180"),
 *                 Map.entry("bastion_enabled", "true"),
 *                 Map.entry("bastion_ttl", "180"),
 *                 Map.entry("waf_enabled", "true"),
 *                 Map.entry("waf_ttl", "180"),
 *                 Map.entry("cloudfirewall_enabled", "true"),
 *                 Map.entry("cloudfirewall_ttl", "180"),
 *                 Map.entry("ddos_coo_access_enabled", "false"),
 *                 Map.entry("ddos_coo_access_ttl", "180"),
 *                 Map.entry("ddos_bgp_access_enabled", "false"),
 *                 Map.entry("ddos_bgp_access_ttl", "180"),
 *                 Map.entry("ddos_dip_access_enabled", "false"),
 *                 Map.entry("ddos_dip_access_ttl", "180"),
 *                 Map.entry("sas_crack_enabled", "true"),
 *                 Map.entry("sas_dns_enabled", "true"),
 *                 Map.entry("sas_http_enabled", "true"),
 *                 Map.entry("sas_local_dns_enabled", "true"),
 *                 Map.entry("sas_login_enabled", "true"),
 *                 Map.entry("sas_network_enabled", "true"),
 *                 Map.entry("sas_process_enabled", "true"),
 *                 Map.entry("sas_security_alert_enabled", "true"),
 *                 Map.entry("sas_security_hc_enabled", "true"),
 *                 Map.entry("sas_security_vul_enabled", "true"),
 *                 Map.entry("sas_session_enabled", "true"),
 *                 Map.entry("sas_snapshot_account_enabled", "true"),
 *                 Map.entry("sas_snapshot_port_enabled", "true"),
 *                 Map.entry("sas_snapshot_process_enabled", "true"),
 *                 Map.entry("sas_ttl", "180"),
 *                 Map.entry("apigateway_enabled", "true"),
 *                 Map.entry("apigateway_ttl", "180"),
 *                 Map.entry("nas_enabled", "true"),
 *                 Map.entry("nas_ttl", "180"),
 *                 Map.entry("appconnect_enabled", "false"),
 *                 Map.entry("appconnect_ttl", "180"),
 *                 Map.entry("cps_enabled", "true"),
 *                 Map.entry("cps_ttl", "180"),
 *                 Map.entry("k8s_audit_enabled", "true"),
 *                 Map.entry("k8s_audit_collection_policy", ""),
 *                 Map.entry("k8s_audit_ttl", "180"),
 *                 Map.entry("k8s_event_enabled", "true"),
 *                 Map.entry("k8s_event_collection_policy", ""),
 *                 Map.entry("k8s_event_ttl", "180"),
 *                 Map.entry("k8s_ingress_enabled", "true"),
 *                 Map.entry("k8s_ingress_collection_policy", ""),
 *                 Map.entry("k8s_ingress_ttl", "180")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * Multiple accounts Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.log.Audit;
 * import com.pulumi.alicloud.log.AuditArgs;
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
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Audit("example", AuditArgs.builder()
 *             .displayName("tf-audit-example")
 *             .aliuid(default_.id())
 *             .variableMap(Map.ofEntries(
 *                 Map.entry("actiontrail_enabled", "true"),
 *                 Map.entry("actiontrail_ttl", "180"),
 *                 Map.entry("oss_access_enabled", "true"),
 *                 Map.entry("oss_access_ttl", "180")
 *             ))
 *             .multiAccounts(            
 *                 "123456789123",
 *                 "12345678912300123")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * Resource Directory Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.log.Audit;
 * import com.pulumi.alicloud.log.AuditArgs;
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
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Audit("example", AuditArgs.builder()
 *             .displayName("tf-audit-example")
 *             .aliuid(default_.id())
 *             .variableMap(Map.ofEntries(
 *                 Map.entry("actiontrail_enabled", "true"),
 *                 Map.entry("actiontrail_ttl", "180"),
 *                 Map.entry("oss_access_enabled", "true"),
 *                 Map.entry("oss_access_ttl", "180")
 *             ))
 *             .resourceDirectoryType("all")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.log.Audit;
 * import com.pulumi.alicloud.log.AuditArgs;
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
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Audit("example", AuditArgs.builder()
 *             .displayName("tf-audit-example")
 *             .aliuid(default_.id())
 *             .variableMap(Map.ofEntries(
 *                 Map.entry("actiontrail_enabled", "true"),
 *                 Map.entry("actiontrail_ttl", "180"),
 *                 Map.entry("oss_access_enabled", "true"),
 *                 Map.entry("oss_access_ttl", "180")
 *             ))
 *             .multiAccounts()
 *             .resourceDirectoryType("custom")
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
 * Log audit can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:log/audit:Audit example tf-audit-example
 * ```
 * 
 */
@ResourceType(type="alicloud:log/audit:Audit")
public class Audit extends com.pulumi.resources.CustomResource {
    /**
     * Aliuid value of your account.
     * 
     */
    @Export(name="aliuid", refs={String.class}, tree="[0]")
    private Output<String> aliuid;

    /**
     * @return Aliuid value of your account.
     * 
     */
    public Output<String> aliuid() {
        return this.aliuid;
    }
    /**
     * Name of SLS log audit.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Name of SLS log audit.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Multi-account configuration, please fill in multiple aliuid.
     * 
     */
    @Export(name="multiAccounts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> multiAccounts;

    /**
     * @return Multi-account configuration, please fill in multiple aliuid.
     * 
     */
    public Output<Optional<List<String>>> multiAccounts() {
        return Codegen.optional(this.multiAccounts);
    }
    /**
     * Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
     * 
     */
    @Export(name="resourceDirectoryType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceDirectoryType;

    /**
     * @return Resource Directory type. Optional values are all or custom. If the value is custom, argument multi_account should be provided.
     * 
     */
    public Output<Optional<String>> resourceDirectoryType() {
        return Codegen.optional(this.resourceDirectoryType);
    }
    /**
     * Log audit detailed configuration.
     * 
     */
    @Export(name="variableMap", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> variableMap;

    /**
     * @return Log audit detailed configuration.
     * 
     */
    public Output<Optional<Map<String,String>>> variableMap() {
        return Codegen.optional(this.variableMap);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Audit(java.lang.String name) {
        this(name, AuditArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Audit(java.lang.String name, AuditArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Audit(java.lang.String name, AuditArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/audit:Audit", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Audit(java.lang.String name, Output<java.lang.String> id, @Nullable AuditState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:log/audit:Audit", name, state, makeResourceOptions(options, id), false);
    }

    private static AuditArgs makeArgs(AuditArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuditArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Audit get(java.lang.String name, Output<java.lang.String> id, @Nullable AuditState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Audit(name, id, state, options);
    }
}
