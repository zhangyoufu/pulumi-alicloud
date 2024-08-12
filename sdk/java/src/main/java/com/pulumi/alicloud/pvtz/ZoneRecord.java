// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.pvtz.ZoneRecordArgs;
import com.pulumi.alicloud.pvtz.inputs.ZoneRecordState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.alicloud.pvtz.Zone;
 * import com.pulumi.alicloud.pvtz.ZoneArgs;
 * import com.pulumi.alicloud.pvtz.ZoneRecord;
 * import com.pulumi.alicloud.pvtz.ZoneRecordArgs;
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
 *         var zone = new Zone("zone", ZoneArgs.builder()
 *             .name("foo.test.com")
 *             .build());
 * 
 *         var foo = new ZoneRecord("foo", ZoneRecordArgs.builder()
 *             .zoneId(zone.id())
 *             .rr("www")
 *             .type("CNAME")
 *             .value("bbb.test.com")
 *             .ttl(60)
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
 * Private Zone Record can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:pvtz/zoneRecord:ZoneRecord example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:pvtz/zoneRecord:ZoneRecord")
public class ZoneRecord extends com.pulumi.resources.CustomResource {
    /**
     * User language.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return User language.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The priority of the Private Zone Record. At present, only can &#34;MX&#34; record support it. Valid values: [1-99]. Default to 1.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return The priority of the Private Zone Record. At present, only can &#34;MX&#34; record support it. Valid values: [1-99]. Default to 1.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The Private Zone Record ID.
     * 
     */
    @Export(name="recordId", refs={String.class}, tree="[0]")
    private Output<String> recordId;

    /**
     * @return The Private Zone Record ID.
     * 
     */
    public Output<String> recordId() {
        return this.recordId;
    }
    /**
     * The remark of the Private Zone Record.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remark;

    /**
     * @return The remark of the Private Zone Record.
     * 
     */
    public Output<Optional<String>> remark() {
        return Codegen.optional(this.remark);
    }
    /**
     * The resource record of the Private Zone Record.
     * 
     * @deprecated
     * Field &#39;resource_record&#39; has been deprecated from version 1.109.0. Use &#39;rr&#39; instead.
     * 
     */
    @Deprecated /* Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead. */
    @Export(name="resourceRecord", refs={String.class}, tree="[0]")
    private Output<String> resourceRecord;

    /**
     * @return The resource record of the Private Zone Record.
     * 
     */
    public Output<String> resourceRecord() {
        return this.resourceRecord;
    }
    /**
     * The rr of the Private Zone Record.
     * 
     */
    @Export(name="rr", refs={String.class}, tree="[0]")
    private Output<String> rr;

    /**
     * @return The rr of the Private Zone Record.
     * 
     */
    public Output<String> rr() {
        return this.rr;
    }
    /**
     * Resolve record status. Value:
     * - ENABLE: enable resolution.
     * - DISABLE: pause parsing.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> status;

    /**
     * @return Resolve record status. Value:
     * - ENABLE: enable resolution.
     * - DISABLE: pause parsing.
     * 
     */
    public Output<Optional<String>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * The ttl of the Private Zone Record. Default to `60`.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return The ttl of the Private Zone Record. Default to `60`.
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    @Export(name="userClientIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userClientIp;

    public Output<Optional<String>> userClientIp() {
        return Codegen.optional(this.userClientIp);
    }
    /**
     * The value of the Private Zone Record.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The value of the Private Zone Record.
     * 
     */
    public Output<String> value() {
        return this.value;
    }
    /**
     * The name of the Private Zone Record.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The name of the Private Zone Record.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ZoneRecord(java.lang.String name) {
        this(name, ZoneRecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ZoneRecord(java.lang.String name, ZoneRecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ZoneRecord(java.lang.String name, ZoneRecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:pvtz/zoneRecord:ZoneRecord", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ZoneRecord(java.lang.String name, Output<java.lang.String> id, @Nullable ZoneRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:pvtz/zoneRecord:ZoneRecord", name, state, makeResourceOptions(options, id), false);
    }

    private static ZoneRecordArgs makeArgs(ZoneRecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ZoneRecordArgs.Empty : args;
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
    public static ZoneRecord get(java.lang.String name, Output<java.lang.String> id, @Nullable ZoneRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ZoneRecord(name, id, state, options);
    }
}
