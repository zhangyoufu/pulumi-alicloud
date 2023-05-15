// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceTypesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceTypesArgs Empty = new GetInstanceTypesArgs();

    /**
     * The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
     * 
     */
    @Import(name="clusterType", required=true)
    private Output<String> clusterType;

    /**
     * @return The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }

    /**
     * The destination resource of emr cluster instance
     * 
     */
    @Import(name="destinationResource", required=true)
    private Output<String> destinationResource;

    /**
     * @return The destination resource of emr cluster instance
     * 
     */
    public Output<String> destinationResource() {
        return this.destinationResource;
    }

    /**
     * Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     * 
     */
    @Import(name="instanceChargeType", required=true)
    private Output<String> instanceChargeType;

    /**
     * @return Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
     * 
     */
    public Output<String> instanceChargeType() {
        return this.instanceChargeType;
    }

    /**
     * Filter the specific ecs instance type to create emr cluster.
     * 
     */
    @Import(name="instanceType")
    private @Nullable Output<String> instanceType;

    /**
     * @return Filter the specific ecs instance type to create emr cluster.
     * 
     */
    public Optional<Output<String>> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Whether the current storage disk is local or not.
     * 
     */
    @Import(name="supportLocalStorage")
    private @Nullable Output<Boolean> supportLocalStorage;

    /**
     * @return Whether the current storage disk is local or not.
     * 
     */
    public Optional<Output<Boolean>> supportLocalStorage() {
        return Optional.ofNullable(this.supportLocalStorage);
    }

    /**
     * The specific supported node type list.
     * Possible values may be any one or combination of these: [&#34;MASTER&#34;, &#34;CORE&#34;, &#34;TASK&#34;, &#34;GATEWAY&#34;]
     * 
     */
    @Import(name="supportNodeTypes")
    private @Nullable Output<List<String>> supportNodeTypes;

    /**
     * @return The specific supported node type list.
     * Possible values may be any one or combination of these: [&#34;MASTER&#34;, &#34;CORE&#34;, &#34;TASK&#34;, &#34;GATEWAY&#34;]
     * 
     */
    public Optional<Output<List<String>>> supportNodeTypes() {
        return Optional.ofNullable(this.supportNodeTypes);
    }

    /**
     * The supported resources of specific zoneId.
     * 
     */
    @Import(name="zoneId")
    private @Nullable Output<String> zoneId;

    /**
     * @return The supported resources of specific zoneId.
     * 
     */
    public Optional<Output<String>> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    private GetInstanceTypesArgs() {}

    private GetInstanceTypesArgs(GetInstanceTypesArgs $) {
        this.clusterType = $.clusterType;
        this.destinationResource = $.destinationResource;
        this.instanceChargeType = $.instanceChargeType;
        this.instanceType = $.instanceType;
        this.outputFile = $.outputFile;
        this.supportLocalStorage = $.supportLocalStorage;
        this.supportNodeTypes = $.supportNodeTypes;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceTypesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceTypesArgs $;

        public Builder() {
            $ = new GetInstanceTypesArgs();
        }

        public Builder(GetInstanceTypesArgs defaults) {
            $ = new GetInstanceTypesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterType The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(Output<String> clusterType) {
            $.clusterType = clusterType;
            return this;
        }

        /**
         * @param clusterType The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(String clusterType) {
            return clusterType(Output.of(clusterType));
        }

        /**
         * @param destinationResource The destination resource of emr cluster instance
         * 
         * @return builder
         * 
         */
        public Builder destinationResource(Output<String> destinationResource) {
            $.destinationResource = destinationResource;
            return this;
        }

        /**
         * @param destinationResource The destination resource of emr cluster instance
         * 
         * @return builder
         * 
         */
        public Builder destinationResource(String destinationResource) {
            return destinationResource(Output.of(destinationResource));
        }

        /**
         * @param instanceChargeType Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(Output<String> instanceChargeType) {
            $.instanceChargeType = instanceChargeType;
            return this;
        }

        /**
         * @param instanceChargeType Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder instanceChargeType(String instanceChargeType) {
            return instanceChargeType(Output.of(instanceChargeType));
        }

        /**
         * @param instanceType Filter the specific ecs instance type to create emr cluster.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(@Nullable Output<String> instanceType) {
            $.instanceType = instanceType;
            return this;
        }

        /**
         * @param instanceType Filter the specific ecs instance type to create emr cluster.
         * 
         * @return builder
         * 
         */
        public Builder instanceType(String instanceType) {
            return instanceType(Output.of(instanceType));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param supportLocalStorage Whether the current storage disk is local or not.
         * 
         * @return builder
         * 
         */
        public Builder supportLocalStorage(@Nullable Output<Boolean> supportLocalStorage) {
            $.supportLocalStorage = supportLocalStorage;
            return this;
        }

        /**
         * @param supportLocalStorage Whether the current storage disk is local or not.
         * 
         * @return builder
         * 
         */
        public Builder supportLocalStorage(Boolean supportLocalStorage) {
            return supportLocalStorage(Output.of(supportLocalStorage));
        }

        /**
         * @param supportNodeTypes The specific supported node type list.
         * Possible values may be any one or combination of these: [&#34;MASTER&#34;, &#34;CORE&#34;, &#34;TASK&#34;, &#34;GATEWAY&#34;]
         * 
         * @return builder
         * 
         */
        public Builder supportNodeTypes(@Nullable Output<List<String>> supportNodeTypes) {
            $.supportNodeTypes = supportNodeTypes;
            return this;
        }

        /**
         * @param supportNodeTypes The specific supported node type list.
         * Possible values may be any one or combination of these: [&#34;MASTER&#34;, &#34;CORE&#34;, &#34;TASK&#34;, &#34;GATEWAY&#34;]
         * 
         * @return builder
         * 
         */
        public Builder supportNodeTypes(List<String> supportNodeTypes) {
            return supportNodeTypes(Output.of(supportNodeTypes));
        }

        /**
         * @param supportNodeTypes The specific supported node type list.
         * Possible values may be any one or combination of these: [&#34;MASTER&#34;, &#34;CORE&#34;, &#34;TASK&#34;, &#34;GATEWAY&#34;]
         * 
         * @return builder
         * 
         */
        public Builder supportNodeTypes(String... supportNodeTypes) {
            return supportNodeTypes(List.of(supportNodeTypes));
        }

        /**
         * @param zoneId The supported resources of specific zoneId.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(@Nullable Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The supported resources of specific zoneId.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public GetInstanceTypesArgs build() {
            $.clusterType = Objects.requireNonNull($.clusterType, "expected parameter 'clusterType' to be non-null");
            $.destinationResource = Objects.requireNonNull($.destinationResource, "expected parameter 'destinationResource' to be non-null");
            $.instanceChargeType = Objects.requireNonNull($.instanceChargeType, "expected parameter 'instanceChargeType' to be non-null");
            return $;
        }
    }

}
