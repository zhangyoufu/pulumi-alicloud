// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.HoneypotProbeArgs;
import com.pulumi.alicloud.threatdetection.inputs.HoneypotProbeState;
import com.pulumi.alicloud.threatdetection.outputs.HoneypotProbeHoneypotBindList;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Honeypot Probe resource.
 * 
 * For information about Threat Detection Honeypot Probe and how to use it, see [What is Honeypot Probe](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-createhoneypotprobe).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
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
 * import com.pulumi.alicloud.threatdetection.HoneypotProbe;
 * import com.pulumi.alicloud.threatdetection.HoneypotProbeArgs;
 * import com.pulumi.alicloud.threatdetection.inputs.HoneypotProbeHoneypotBindListArgs;
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
 *         var default_ = new HoneypotProbe("default", HoneypotProbeArgs.builder()
 *             .uuid("032b618f-b220-4a0d-bd37-fbdc6ef58b6a")
 *             .probeType("host_probe")
 *             .controlNodeId("a44e1ab3-6945-444c-889d-5bacee7056e8")
 *             .ping(true)
 *             .honeypotBindLists(HoneypotProbeHoneypotBindListArgs.builder()
 *                 .bindPortLists(HoneypotProbeHoneypotBindListBindPortListArgs.builder()
 *                     .startPort(80)
 *                     .endPort(80)
 *                     .build())
 *                 .honeypotId("ede59ccdb1b7a2e21735d4593a6eb5ed31883af320c5ab63ab33818e94307be9")
 *                 .build())
 *             .displayName("apispec")
 *             .arp(true)
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
 * Threat Detection Honeypot Probe can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:threatdetection/honeypotProbe:HoneypotProbe example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/honeypotProbe:HoneypotProbe")
public class HoneypotProbe extends com.pulumi.resources.CustomResource {
    /**
     * ARP spoofing detection.**true**: Enable **false**: Disabled
     * 
     */
    @Export(name="arp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> arp;

    /**
     * @return ARP spoofing detection.**true**: Enable **false**: Disabled
     * 
     */
    public Output<Optional<Boolean>> arp() {
        return Codegen.optional(this.arp);
    }
    /**
     * The ID of the management node.
     * 
     */
    @Export(name="controlNodeId", refs={String.class}, tree="[0]")
    private Output<String> controlNodeId;

    /**
     * @return The ID of the management node.
     * 
     */
    public Output<String> controlNodeId() {
        return this.controlNodeId;
    }
    /**
     * Probe display name.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Probe display name.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Configure the service.See the following `Block HoneypotBindList`.
     * 
     */
    @Export(name="honeypotBindLists", refs={List.class,HoneypotProbeHoneypotBindList.class}, tree="[0,1]")
    private Output</* @Nullable */ List<HoneypotProbeHoneypotBindList>> honeypotBindLists;

    /**
     * @return Configure the service.See the following `Block HoneypotBindList`.
     * 
     */
    public Output<Optional<List<HoneypotProbeHoneypotBindList>>> honeypotBindLists() {
        return Codegen.optional(this.honeypotBindLists);
    }
    /**
     * The first ID of the resource
     * 
     */
    @Export(name="honeypotProbeId", refs={String.class}, tree="[0]")
    private Output<String> honeypotProbeId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> honeypotProbeId() {
        return this.honeypotProbeId;
    }
    /**
     * Ping scan detection. Value: **true**: Enable **false**: Disabled
     * 
     */
    @Export(name="ping", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ping;

    /**
     * @return Ping scan detection. Value: **true**: Enable **false**: Disabled
     * 
     */
    public Output<Optional<Boolean>> ping() {
        return Codegen.optional(this.ping);
    }
    /**
     * Probe type, support `host_probe` and `vpc_black_hole_probe`.
     * 
     */
    @Export(name="probeType", refs={String.class}, tree="[0]")
    private Output<String> probeType;

    /**
     * @return Probe type, support `host_probe` and `vpc_black_hole_probe`.
     * 
     */
    public Output<String> probeType() {
        return this.probeType;
    }
    /**
     * The version of the probe.
     * 
     */
    @Export(name="probeVersion", refs={String.class}, tree="[0]")
    private Output<String> probeVersion;

    /**
     * @return The version of the probe.
     * 
     */
    public Output<String> probeVersion() {
        return this.probeVersion;
    }
    /**
     * The IP address of the proxy.
     * 
     */
    @Export(name="proxyIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> proxyIp;

    /**
     * @return The IP address of the proxy.
     * 
     */
    public Output<Optional<String>> proxyIp() {
        return Codegen.optional(this.proxyIp);
    }
    /**
     * Listen to the IP address list.
     * 
     */
    @Export(name="serviceIpLists", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> serviceIpLists;

    /**
     * @return Listen to the IP address list.
     * 
     */
    public Output<List<String>> serviceIpLists() {
        return this.serviceIpLists;
    }
    /**
     * The status of the resource
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Machine uuid, **probe_type** is `host_probe`. This value cannot be empty.
     * 
     */
    @Export(name="uuid", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> uuid;

    /**
     * @return Machine uuid, **probe_type** is `host_probe`. This value cannot be empty.
     * 
     */
    public Output<Optional<String>> uuid() {
        return Codegen.optional(this.uuid);
    }
    /**
     * The ID of the VPC. **probe_type** is `vpc_black_hole_probe`. This value cannot be empty.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The ID of the VPC. **probe_type** is `vpc_black_hole_probe`. This value cannot be empty.
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HoneypotProbe(java.lang.String name) {
        this(name, HoneypotProbeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HoneypotProbe(java.lang.String name, HoneypotProbeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HoneypotProbe(java.lang.String name, HoneypotProbeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/honeypotProbe:HoneypotProbe", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HoneypotProbe(java.lang.String name, Output<java.lang.String> id, @Nullable HoneypotProbeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/honeypotProbe:HoneypotProbe", name, state, makeResourceOptions(options, id), false);
    }

    private static HoneypotProbeArgs makeArgs(HoneypotProbeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HoneypotProbeArgs.Empty : args;
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
    public static HoneypotProbe get(java.lang.String name, Output<java.lang.String> id, @Nullable HoneypotProbeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HoneypotProbe(name, id, state, options);
    }
}
