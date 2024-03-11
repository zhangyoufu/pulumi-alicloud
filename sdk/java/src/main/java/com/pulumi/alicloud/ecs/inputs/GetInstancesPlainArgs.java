// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstancesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstancesPlainArgs Empty = new GetInstancesPlainArgs();

    /**
     * Availability zone where instances are located.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable String availabilityZone;

    /**
     * @return Availability zone where instances are located.
     * 
     */
    public Optional<String> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * Default to `true`. If false, the attributes `ram_role_name` and `disk_device_mappings` will not be fetched and output.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `true`. If false, the attributes `ram_role_name` and `disk_device_mappings` will not be fetched and output.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of ECS instance IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of ECS instance IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The image ID of some ECS instance used.
     * 
     */
    @Import(name="imageId")
    private @Nullable String imageId;

    /**
     * @return The image ID of some ECS instance used.
     * 
     */
    public Optional<String> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * The name of the instance. Fuzzy search with the asterisk (*) wildcard characters is supported.
     * 
     */
    @Import(name="instanceName")
    private @Nullable String instanceName;

    /**
     * @return The name of the instance. Fuzzy search with the asterisk (*) wildcard characters is supported.
     * 
     */
    public Optional<String> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * A regex string to filter results by instance name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by instance name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The RAM role name which the instance attaches.
     * 
     */
    @Import(name="ramRoleName")
    private @Nullable String ramRoleName;

    /**
     * @return The RAM role name which the instance attaches.
     * 
     */
    public Optional<String> ramRoleName() {
        return Optional.ofNullable(this.ramRoleName);
    }

    /**
     * The ID of resource group which the instance belongs.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The ID of resource group which the instance belongs.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Instance status. Valid values: &#34;Creating&#34;, &#34;Starting&#34;, &#34;Running&#34;, &#34;Stopping&#34; and &#34;Stopped&#34;. If undefined, all statuses are considered.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return Instance status. Valid values: &#34;Creating&#34;, &#34;Starting&#34;, &#34;Running&#34;, &#34;Stopping&#34; and &#34;Stopped&#34;. If undefined, all statuses are considered.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A map of tags assigned to the ECS instances. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ecs.EcsFunctions;
     * import com.pulumi.alicloud.ecs.inputs.GetInstancesArgs;
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
     *         final var taggedInstances = EcsFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
     *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
     *             ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return A map of tags assigned to the ECS instances. It must be in the format:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ecs.EcsFunctions;
     * import com.pulumi.alicloud.ecs.inputs.GetInstancesArgs;
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
     *         final var taggedInstances = EcsFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
     *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
     *             ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * ID of the VPC linked to the instances.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return ID of the VPC linked to the instances.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * ID of the VSwitch linked to the instances.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable String vswitchId;

    /**
     * @return ID of the VSwitch linked to the instances.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetInstancesPlainArgs() {}

    private GetInstancesPlainArgs(GetInstancesPlainArgs $) {
        this.availabilityZone = $.availabilityZone;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.imageId = $.imageId;
        this.instanceName = $.instanceName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.ramRoleName = $.ramRoleName;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstancesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstancesPlainArgs $;

        public Builder() {
            $ = new GetInstancesPlainArgs();
        }

        public Builder(GetInstancesPlainArgs defaults) {
            $ = new GetInstancesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityZone Availability zone where instances are located.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable String availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param enableDetails Default to `true`. If false, the attributes `ram_role_name` and `disk_device_mappings` will not be fetched and output.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of ECS instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of ECS instance IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param imageId The image ID of some ECS instance used.
         * 
         * @return builder
         * 
         */
        public Builder imageId(@Nullable String imageId) {
            $.imageId = imageId;
            return this;
        }

        /**
         * @param instanceName The name of the instance. Fuzzy search with the asterisk (*) wildcard characters is supported.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable String instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by instance name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param ramRoleName The RAM role name which the instance attaches.
         * 
         * @return builder
         * 
         */
        public Builder ramRoleName(@Nullable String ramRoleName) {
            $.ramRoleName = ramRoleName;
            return this;
        }

        /**
         * @param resourceGroupId The ID of resource group which the instance belongs.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param status Instance status. Valid values: &#34;Creating&#34;, &#34;Starting&#34;, &#34;Running&#34;, &#34;Stopping&#34; and &#34;Stopped&#34;. If undefined, all statuses are considered.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags A map of tags assigned to the ECS instances. It must be in the format:
         * &lt;!--Start PulumiCodeChooser --&gt;
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import com.pulumi.alicloud.ecs.EcsFunctions;
         * import com.pulumi.alicloud.ecs.inputs.GetInstancesArgs;
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
         *         final var taggedInstances = EcsFunctions.getInstances(GetInstancesArgs.builder()
         *             .tags(Map.ofEntries(
         *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
         *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
         *             ))
         *             .build());
         * 
         *     }
         * }
         * ```
         * &lt;!--End PulumiCodeChooser --&gt;
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param vpcId ID of the VPC linked to the instances.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vswitchId ID of the VSwitch linked to the instances.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable String vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        public GetInstancesPlainArgs build() {
            return $;
        }
    }

}
