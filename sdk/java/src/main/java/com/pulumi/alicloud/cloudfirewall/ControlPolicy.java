// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudfirewall.ControlPolicyArgs;
import com.pulumi.alicloud.cloudfirewall.inputs.ControlPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Firewall Control Policy resource.
 * 
 * For information about Cloud Firewall Control Policy and how to use it, see [What is Control Policy](https://www.alibabacloud.com/help/doc-detail/138867.htm).
 * 
 * &gt; **NOTE:** Available since v1.129.0.
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
 * import com.pulumi.alicloud.cloudfirewall.ControlPolicy;
 * import com.pulumi.alicloud.cloudfirewall.ControlPolicyArgs;
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
 *         var example = new ControlPolicy("example", ControlPolicyArgs.builder()        
 *             .applicationName("ANY")
 *             .aclAction("accept")
 *             .description("example")
 *             .destinationType("net")
 *             .destination("100.1.1.0/24")
 *             .direction("out")
 *             .proto("ANY")
 *             .source("1.2.3.0/24")
 *             .sourceType("net")
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
 * Cloud Firewall Control Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudfirewall/controlPolicy:ControlPolicy example &lt;acl_uuid&gt;:&lt;direction&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudfirewall/controlPolicy:ControlPolicy")
public class ControlPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     * 
     */
    @Export(name="aclAction", refs={String.class}, tree="[0]")
    private Output<String> aclAction;

    /**
     * @return The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     * 
     */
    public Output<String> aclAction() {
        return this.aclAction;
    }
    /**
     * (Available since v1.148.0) The unique ID of the access control policy.
     * 
     */
    @Export(name="aclUuid", refs={String.class}, tree="[0]")
    private Output<String> aclUuid;

    /**
     * @return (Available since v1.148.0) The unique ID of the access control policy.
     * 
     */
    public Output<String> aclUuid() {
        return this.aclUuid;
    }
    /**
     * The application type supported by the access control policy. Valid values: `ANY`, `HTTP`, `HTTPS`, `MQTT`, `Memcache`, `MongoDB`, `MySQL`, `RDP`, `Redis`, `SMTP`, `SMTPS`, `SSH`, `SSL`, `VNC`.
     * &gt; **NOTE:** If `proto` is set to `TCP`, you can set `application_name` to any valid value. If `proto` is set to `UDP`, `ICMP`, or `ANY`, you can only set `application_name` to `ANY`.
     * 
     */
    @Export(name="applicationName", refs={String.class}, tree="[0]")
    private Output<String> applicationName;

    /**
     * @return The application type supported by the access control policy. Valid values: `ANY`, `HTTP`, `HTTPS`, `MQTT`, `Memcache`, `MongoDB`, `MySQL`, `RDP`, `Redis`, `SMTP`, `SMTPS`, `SSH`, `SSL`, `VNC`.
     * &gt; **NOTE:** If `proto` is set to `TCP`, you can set `application_name` to any valid value. If `proto` is set to `UDP`, `ICMP`, or `ANY`, you can only set `application_name` to `ANY`.
     * 
     */
    public Output<String> applicationName() {
        return this.applicationName;
    }
    /**
     * The description of the access control policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the access control policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The destination port defined in the access control policy.
     * 
     */
    @Export(name="destPort", refs={String.class}, tree="[0]")
    private Output<String> destPort;

    /**
     * @return The destination port defined in the access control policy.
     * 
     */
    public Output<String> destPort() {
        return this.destPort;
    }
    /**
     * The destination port address book defined in the access control policy.
     * 
     */
    @Export(name="destPortGroup", refs={String.class}, tree="[0]")
    private Output<String> destPortGroup;

    /**
     * @return The destination port address book defined in the access control policy.
     * 
     */
    public Output<String> destPortGroup() {
        return this.destPortGroup;
    }
    /**
     * The destination port type defined in the access control policy. Valid values: `group`, `port`.
     * 
     */
    @Export(name="destPortType", refs={String.class}, tree="[0]")
    private Output<String> destPortType;

    /**
     * @return The destination port type defined in the access control policy. Valid values: `group`, `port`.
     * 
     */
    public Output<String> destPortType() {
        return this.destPortType;
    }
    /**
     * The destination address defined in the access control policy.
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output<String> destination;

    /**
     * @return The destination address defined in the access control policy.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * DestinationType. Valid values: If Direction is `in`, the valid values are `net`, `group`. If `direction` is `out`, the valid values are `net`, `group`, `domain`, `location`.
     * 
     */
    @Export(name="destinationType", refs={String.class}, tree="[0]")
    private Output<String> destinationType;

    /**
     * @return DestinationType. Valid values: If Direction is `in`, the valid values are `net`, `group`. If `direction` is `out`, the valid values are `net`, `group`, `domain`, `location`.
     * 
     */
    public Output<String> destinationType() {
        return this.destinationType;
    }
    /**
     * Direction. Valid values: `in`, `out`.
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output<String> direction;

    /**
     * @return Direction. Valid values: `in`, `out`.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * The ip version.
     * 
     */
    @Export(name="ipVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipVersion;

    /**
     * @return The ip version.
     * 
     */
    public Output<Optional<String>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * DestPortGroupPorts. Valid values: `en`, `zh`.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return DestPortGroupPorts. Valid values: `en`, `zh`.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * Proto. Valid values: `  TCP `, `  UDP `, `ANY`, `ICMP`.
     * 
     */
    @Export(name="proto", refs={String.class}, tree="[0]")
    private Output<String> proto;

    /**
     * @return Proto. Valid values: `  TCP `, `  UDP `, `ANY`, `ICMP`.
     * 
     */
    public Output<String> proto() {
        return this.proto;
    }
    /**
     * Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values: `true`, `false`.
     * 
     */
    @Export(name="release", refs={String.class}, tree="[0]")
    private Output<String> release;

    /**
     * @return Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values: `true`, `false`.
     * 
     */
    public Output<String> release() {
        return this.release;
    }
    /**
     * Source.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return Source.
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * The source ip.
     * 
     */
    @Export(name="sourceIp", refs={String.class}, tree="[0]")
    private Output<String> sourceIp;

    /**
     * @return The source ip.
     * 
     */
    public Output<String> sourceIp() {
        return this.sourceIp;
    }
    /**
     * SourceType. Valid values: If `direction` is `in`, the valid values are `net`, `group`, `location`. If `direction` is `out`, the valid values are `net`, `group`.
     * 
     */
    @Export(name="sourceType", refs={String.class}, tree="[0]")
    private Output<String> sourceType;

    /**
     * @return SourceType. Valid values: If `direction` is `in`, the valid values are `net`, `group`, `location`. If `direction` is `out`, the valid values are `net`, `group`.
     * 
     */
    public Output<String> sourceType() {
        return this.sourceType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ControlPolicy(String name) {
        this(name, ControlPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ControlPolicy(String name, ControlPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ControlPolicy(String name, ControlPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/controlPolicy:ControlPolicy", name, args == null ? ControlPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ControlPolicy(String name, Output<String> id, @Nullable ControlPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/controlPolicy:ControlPolicy", name, state, makeResourceOptions(options, id));
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
    public static ControlPolicy get(String name, Output<String> id, @Nullable ControlPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ControlPolicy(name, id, state, options);
    }
}
