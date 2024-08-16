// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TrafficMirrorSessionState extends com.pulumi.resources.ResourceArgs {

    public static final TrafficMirrorSessionState Empty = new TrafficMirrorSessionState();

    /**
     * Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return Whether to PreCheck only this request, value:
     * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
     * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * Specifies whether to enable traffic mirror sessions. default to `false`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Specifies whether to enable traffic mirror sessions. default to `false`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Maximum Transmission Unit (MTU).
     * 
     */
    @Import(name="packetLength")
    private @Nullable Output<Integer> packetLength;

    /**
     * @return Maximum Transmission Unit (MTU).
     * 
     */
    public Optional<Output<Integer>> packetLength() {
        return Optional.ofNullable(this.packetLength);
    }

    /**
     * The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags of this resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return The tags of this resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ID of the filter.
     * 
     */
    @Import(name="trafficMirrorFilterId")
    private @Nullable Output<String> trafficMirrorFilterId;

    /**
     * @return The ID of the filter.
     * 
     */
    public Optional<Output<String>> trafficMirrorFilterId() {
        return Optional.ofNullable(this.trafficMirrorFilterId);
    }

    /**
     * The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Import(name="trafficMirrorSessionDescription")
    private @Nullable Output<String> trafficMirrorSessionDescription;

    /**
     * @return The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Optional<Output<String>> trafficMirrorSessionDescription() {
        return Optional.ofNullable(this.trafficMirrorSessionDescription);
    }

    /**
     * The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     * 
     */
    @Import(name="trafficMirrorSessionName")
    private @Nullable Output<String> trafficMirrorSessionName;

    /**
     * @return The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
     * 
     */
    public Optional<Output<String>> trafficMirrorSessionName() {
        return Optional.ofNullable(this.trafficMirrorSessionName);
    }

    /**
     * The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     * 
     */
    @Import(name="trafficMirrorSourceIds")
    private @Nullable Output<List<String>> trafficMirrorSourceIds;

    /**
     * @return The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
     * 
     */
    public Optional<Output<List<String>>> trafficMirrorSourceIds() {
        return Optional.ofNullable(this.trafficMirrorSourceIds);
    }

    /**
     * The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     * 
     */
    @Import(name="trafficMirrorTargetId")
    private @Nullable Output<String> trafficMirrorTargetId;

    /**
     * @return The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
     * 
     */
    public Optional<Output<String>> trafficMirrorTargetId() {
        return Optional.ofNullable(this.trafficMirrorTargetId);
    }

    /**
     * The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     * 
     */
    @Import(name="trafficMirrorTargetType")
    private @Nullable Output<String> trafficMirrorTargetType;

    /**
     * @return The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
     * 
     */
    public Optional<Output<String>> trafficMirrorTargetType() {
        return Optional.ofNullable(this.trafficMirrorTargetType);
    }

    /**
     * The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     * 
     */
    @Import(name="virtualNetworkId")
    private @Nullable Output<Integer> virtualNetworkId;

    /**
     * @return The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
     * 
     */
    public Optional<Output<Integer>> virtualNetworkId() {
        return Optional.ofNullable(this.virtualNetworkId);
    }

    private TrafficMirrorSessionState() {}

    private TrafficMirrorSessionState(TrafficMirrorSessionState $) {
        this.dryRun = $.dryRun;
        this.enabled = $.enabled;
        this.packetLength = $.packetLength;
        this.priority = $.priority;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
        this.trafficMirrorFilterId = $.trafficMirrorFilterId;
        this.trafficMirrorSessionDescription = $.trafficMirrorSessionDescription;
        this.trafficMirrorSessionName = $.trafficMirrorSessionName;
        this.trafficMirrorSourceIds = $.trafficMirrorSourceIds;
        this.trafficMirrorTargetId = $.trafficMirrorTargetId;
        this.trafficMirrorTargetType = $.trafficMirrorTargetType;
        this.virtualNetworkId = $.virtualNetworkId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrafficMirrorSessionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrafficMirrorSessionState $;

        public Builder() {
            $ = new TrafficMirrorSessionState();
        }

        public Builder(TrafficMirrorSessionState defaults) {
            $ = new TrafficMirrorSessionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dryRun Whether to PreCheck only this request, value:
         * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
         * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun Whether to PreCheck only this request, value:
         * - **true**: sends a check request and does not create a mirror session. Check items include whether required parameters are filled in, request format, and restrictions. If the check fails, the corresponding error is returned. If the check passes, the error code &#39;DryRunOperation&#39; is returned &#39;.
         * - **false** (default): Sends a normal request and directly creates a mirror session after checking.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param enabled Specifies whether to enable traffic mirror sessions. default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether to enable traffic mirror sessions. default to `false`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param packetLength Maximum Transmission Unit (MTU).
         * 
         * @return builder
         * 
         */
        public Builder packetLength(@Nullable Output<Integer> packetLength) {
            $.packetLength = packetLength;
            return this;
        }

        /**
         * @param packetLength Maximum Transmission Unit (MTU).
         * 
         * @return builder
         * 
         */
        public Builder packetLength(Integer packetLength) {
            return packetLength(Output.of(packetLength));
        }

        /**
         * @param priority The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority The priority of the traffic mirror session. Valid values: `1` to `32766`. A smaller value indicates a higher priority. You cannot specify the same priority for traffic mirror sessions that are created in the same region with the same Alibaba Cloud account.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tags of this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param trafficMirrorFilterId The ID of the filter.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorFilterId(@Nullable Output<String> trafficMirrorFilterId) {
            $.trafficMirrorFilterId = trafficMirrorFilterId;
            return this;
        }

        /**
         * @param trafficMirrorFilterId The ID of the filter.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorFilterId(String trafficMirrorFilterId) {
            return trafficMirrorFilterId(Output.of(trafficMirrorFilterId));
        }

        /**
         * @param trafficMirrorSessionDescription The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSessionDescription(@Nullable Output<String> trafficMirrorSessionDescription) {
            $.trafficMirrorSessionDescription = trafficMirrorSessionDescription;
            return this;
        }

        /**
         * @param trafficMirrorSessionDescription The description of the traffic mirror session. The description must be `2` to `256` characters in length and cannot start with `http://` or `https://`.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSessionDescription(String trafficMirrorSessionDescription) {
            return trafficMirrorSessionDescription(Output.of(trafficMirrorSessionDescription));
        }

        /**
         * @param trafficMirrorSessionName The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSessionName(@Nullable Output<String> trafficMirrorSessionName) {
            $.trafficMirrorSessionName = trafficMirrorSessionName;
            return this;
        }

        /**
         * @param trafficMirrorSessionName The name of the traffic mirror session. The name must be `2` to `128` characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSessionName(String trafficMirrorSessionName) {
            return trafficMirrorSessionName(Output.of(trafficMirrorSessionName));
        }

        /**
         * @param trafficMirrorSourceIds The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSourceIds(@Nullable Output<List<String>> trafficMirrorSourceIds) {
            $.trafficMirrorSourceIds = trafficMirrorSourceIds;
            return this;
        }

        /**
         * @param trafficMirrorSourceIds The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSourceIds(List<String> trafficMirrorSourceIds) {
            return trafficMirrorSourceIds(Output.of(trafficMirrorSourceIds));
        }

        /**
         * @param trafficMirrorSourceIds The ID of the image source instance. Currently, the Eni is supported as the image source. The default value of N is 1, that is, only one mirror source can be added to a mirror session.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorSourceIds(String... trafficMirrorSourceIds) {
            return trafficMirrorSourceIds(List.of(trafficMirrorSourceIds));
        }

        /**
         * @param trafficMirrorTargetId The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorTargetId(@Nullable Output<String> trafficMirrorTargetId) {
            $.trafficMirrorTargetId = trafficMirrorTargetId;
            return this;
        }

        /**
         * @param trafficMirrorTargetId The ID of the mirror destination. You can specify only an ENI or a Server Load Balancer (SLB) instance as a mirror destination.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorTargetId(String trafficMirrorTargetId) {
            return trafficMirrorTargetId(Output.of(trafficMirrorTargetId));
        }

        /**
         * @param trafficMirrorTargetType The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorTargetType(@Nullable Output<String> trafficMirrorTargetType) {
            $.trafficMirrorTargetType = trafficMirrorTargetType;
            return this;
        }

        /**
         * @param trafficMirrorTargetType The type of the mirror destination. Valid values: `NetworkInterface` or `SLB`. `NetworkInterface`: an ENI. `SLB`: an internal-facing SLB instance.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorTargetType(String trafficMirrorTargetType) {
            return trafficMirrorTargetType(Output.of(trafficMirrorTargetType));
        }

        /**
         * @param virtualNetworkId The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
         * 
         * @return builder
         * 
         */
        public Builder virtualNetworkId(@Nullable Output<Integer> virtualNetworkId) {
            $.virtualNetworkId = virtualNetworkId;
            return this;
        }

        /**
         * @param virtualNetworkId The VXLAN network identifier (VNI) that is used to distinguish different mirrored traffic. Valid values: `0` to `16777215`. You can specify VNIs for the traffic mirror destination to identify mirrored traffic from different sessions. If you do not specify a VNI, the system randomly allocates a VNI. If you want the system to randomly allocate a VNI, ignore this parameter.
         * 
         * @return builder
         * 
         */
        public Builder virtualNetworkId(Integer virtualNetworkId) {
            return virtualNetworkId(Output.of(virtualNetworkId));
        }

        public TrafficMirrorSessionState build() {
            return $;
        }
    }

}
