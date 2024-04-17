// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.alicloud.hbr.inputs.PolicyBindingAdvancedOptionsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyBindingState extends com.pulumi.resources.ResourceArgs {

    public static final PolicyBindingState Empty = new PolicyBindingState();

    /**
     * Backup Advanced Options. See `advanced_options` below.
     * 
     */
    @Import(name="advancedOptions")
    private @Nullable Output<PolicyBindingAdvancedOptionsArgs> advancedOptions;

    /**
     * @return Backup Advanced Options. See `advanced_options` below.
     * 
     */
    public Optional<Output<PolicyBindingAdvancedOptionsArgs>> advancedOptions() {
        return Optional.ofNullable(this.advancedOptions);
    }

    /**
     * The creation time of the resource.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time of the resource.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The data source ID.
     * 
     */
    @Import(name="dataSourceId")
    private @Nullable Output<String> dataSourceId;

    /**
     * @return The data source ID.
     * 
     */
    public Optional<Output<String>> dataSourceId() {
        return Optional.ofNullable(this.dataSourceId);
    }

    /**
     * Whether the policy is effective for the data source.
     * - true: Pause
     * - false: not paused.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Whether the policy is effective for the data source.
     * - true: Pause
     * - false: not paused.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates a file type that does not need to be backed up. All files of this type are not backed up. A maximum of 255 characters is supported.
     * 
     */
    @Import(name="exclude")
    private @Nullable Output<String> exclude;

    /**
     * @return This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates a file type that does not need to be backed up. All files of this type are not backed up. A maximum of 255 characters is supported.
     * 
     */
    public Optional<Output<String>> exclude() {
        return Optional.ofNullable(this.exclude);
    }

    /**
     * This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates the file types to be backed up, and all files of these types are backed up. A maximum of 255 characters is supported.
     * 
     */
    @Import(name="include")
    private @Nullable Output<String> include;

    /**
     * @return This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates the file types to be backed up, and all files of these types are backed up. A maximum of 255 characters is supported.
     * 
     */
    public Optional<Output<String>> include() {
        return Optional.ofNullable(this.include);
    }

    /**
     * Resource Description.
     * 
     */
    @Import(name="policyBindingDescription")
    private @Nullable Output<String> policyBindingDescription;

    /**
     * @return Resource Description.
     * 
     */
    public Optional<Output<String>> policyBindingDescription() {
        return Optional.ofNullable(this.policyBindingDescription);
    }

    /**
     * The policy ID.
     * 
     */
    @Import(name="policyId")
    private @Nullable Output<String> policyId;

    /**
     * @return The policy ID.
     * 
     */
    public Optional<Output<String>> policyId() {
        return Optional.ofNullable(this.policyId);
    }

    /**
     * When SourceType is OSS, a prefix is specified to be backed up. If it is not specified, the entire root directory of the Bucket is backed up.
     * 
     */
    @Import(name="source")
    private @Nullable Output<String> source;

    /**
     * @return When SourceType is OSS, a prefix is specified to be backed up. If it is not specified, the entire root directory of the Bucket is backed up.
     * 
     */
    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    /**
     * Data source type, value range:
     * - **UDM_ECS**: indicates the ECS instance backup.
     * - **OSS**: indicates an OSS backup.
     * - **NAS**: indicates an Alibaba Cloud NAS Backup. When you bind a file system to a policy, Cloud Backup automatically creates a mount point for the file system. If you no longer need the mount point, delete it manually.
     * - **ECS_FILE**: indicates that the ECS file is backed up.
     * - **File**: indicates a local File backup.
     * 
     */
    @Import(name="sourceType")
    private @Nullable Output<String> sourceType;

    /**
     * @return Data source type, value range:
     * - **UDM_ECS**: indicates the ECS instance backup.
     * - **OSS**: indicates an OSS backup.
     * - **NAS**: indicates an Alibaba Cloud NAS Backup. When you bind a file system to a policy, Cloud Backup automatically creates a mount point for the file system. If you no longer need the mount point, delete it manually.
     * - **ECS_FILE**: indicates that the ECS file is backed up.
     * - **File**: indicates a local File backup.
     * 
     */
    public Optional<Output<String>> sourceType() {
        return Optional.ofNullable(this.sourceType);
    }

    /**
     * This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates backup flow control. The format is {start}{end}{bandwidth}. Multiple flow control configurations use partitioning, and no overlap in configuration time is allowed. start: start hour. end: end of hour. bandwidth: limit rate, in KB/s.
     * 
     */
    @Import(name="speedLimit")
    private @Nullable Output<String> speedLimit;

    /**
     * @return This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates backup flow control. The format is {start}{end}{bandwidth}. Multiple flow control configurations use partitioning, and no overlap in configuration time is allowed. start: start hour. end: end of hour. bandwidth: limit rate, in KB/s.
     * 
     */
    public Optional<Output<String>> speedLimit() {
        return Optional.ofNullable(this.speedLimit);
    }

    private PolicyBindingState() {}

    private PolicyBindingState(PolicyBindingState $) {
        this.advancedOptions = $.advancedOptions;
        this.createTime = $.createTime;
        this.dataSourceId = $.dataSourceId;
        this.disabled = $.disabled;
        this.exclude = $.exclude;
        this.include = $.include;
        this.policyBindingDescription = $.policyBindingDescription;
        this.policyId = $.policyId;
        this.source = $.source;
        this.sourceType = $.sourceType;
        this.speedLimit = $.speedLimit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyBindingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyBindingState $;

        public Builder() {
            $ = new PolicyBindingState();
        }

        public Builder(PolicyBindingState defaults) {
            $ = new PolicyBindingState(Objects.requireNonNull(defaults));
        }

        /**
         * @param advancedOptions Backup Advanced Options. See `advanced_options` below.
         * 
         * @return builder
         * 
         */
        public Builder advancedOptions(@Nullable Output<PolicyBindingAdvancedOptionsArgs> advancedOptions) {
            $.advancedOptions = advancedOptions;
            return this;
        }

        /**
         * @param advancedOptions Backup Advanced Options. See `advanced_options` below.
         * 
         * @return builder
         * 
         */
        public Builder advancedOptions(PolicyBindingAdvancedOptionsArgs advancedOptions) {
            return advancedOptions(Output.of(advancedOptions));
        }

        /**
         * @param createTime The creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time of the resource.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param dataSourceId The data source ID.
         * 
         * @return builder
         * 
         */
        public Builder dataSourceId(@Nullable Output<String> dataSourceId) {
            $.dataSourceId = dataSourceId;
            return this;
        }

        /**
         * @param dataSourceId The data source ID.
         * 
         * @return builder
         * 
         */
        public Builder dataSourceId(String dataSourceId) {
            return dataSourceId(Output.of(dataSourceId));
        }

        /**
         * @param disabled Whether the policy is effective for the data source.
         * - true: Pause
         * - false: not paused.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Whether the policy is effective for the data source.
         * - true: Pause
         * - false: not paused.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param exclude This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates a file type that does not need to be backed up. All files of this type are not backed up. A maximum of 255 characters is supported.
         * 
         * @return builder
         * 
         */
        public Builder exclude(@Nullable Output<String> exclude) {
            $.exclude = exclude;
            return this;
        }

        /**
         * @param exclude This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates a file type that does not need to be backed up. All files of this type are not backed up. A maximum of 255 characters is supported.
         * 
         * @return builder
         * 
         */
        public Builder exclude(String exclude) {
            return exclude(Output.of(exclude));
        }

        /**
         * @param include This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates the file types to be backed up, and all files of these types are backed up. A maximum of 255 characters is supported.
         * 
         * @return builder
         * 
         */
        public Builder include(@Nullable Output<String> include) {
            $.include = include;
            return this;
        }

        /**
         * @param include This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates the file types to be backed up, and all files of these types are backed up. A maximum of 255 characters is supported.
         * 
         * @return builder
         * 
         */
        public Builder include(String include) {
            return include(Output.of(include));
        }

        /**
         * @param policyBindingDescription Resource Description.
         * 
         * @return builder
         * 
         */
        public Builder policyBindingDescription(@Nullable Output<String> policyBindingDescription) {
            $.policyBindingDescription = policyBindingDescription;
            return this;
        }

        /**
         * @param policyBindingDescription Resource Description.
         * 
         * @return builder
         * 
         */
        public Builder policyBindingDescription(String policyBindingDescription) {
            return policyBindingDescription(Output.of(policyBindingDescription));
        }

        /**
         * @param policyId The policy ID.
         * 
         * @return builder
         * 
         */
        public Builder policyId(@Nullable Output<String> policyId) {
            $.policyId = policyId;
            return this;
        }

        /**
         * @param policyId The policy ID.
         * 
         * @return builder
         * 
         */
        public Builder policyId(String policyId) {
            return policyId(Output.of(policyId));
        }

        /**
         * @param source When SourceType is OSS, a prefix is specified to be backed up. If it is not specified, the entire root directory of the Bucket is backed up.
         * 
         * @return builder
         * 
         */
        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source When SourceType is OSS, a prefix is specified to be backed up. If it is not specified, the entire root directory of the Bucket is backed up.
         * 
         * @return builder
         * 
         */
        public Builder source(String source) {
            return source(Output.of(source));
        }

        /**
         * @param sourceType Data source type, value range:
         * - **UDM_ECS**: indicates the ECS instance backup.
         * - **OSS**: indicates an OSS backup.
         * - **NAS**: indicates an Alibaba Cloud NAS Backup. When you bind a file system to a policy, Cloud Backup automatically creates a mount point for the file system. If you no longer need the mount point, delete it manually.
         * - **ECS_FILE**: indicates that the ECS file is backed up.
         * - **File**: indicates a local File backup.
         * 
         * @return builder
         * 
         */
        public Builder sourceType(@Nullable Output<String> sourceType) {
            $.sourceType = sourceType;
            return this;
        }

        /**
         * @param sourceType Data source type, value range:
         * - **UDM_ECS**: indicates the ECS instance backup.
         * - **OSS**: indicates an OSS backup.
         * - **NAS**: indicates an Alibaba Cloud NAS Backup. When you bind a file system to a policy, Cloud Backup automatically creates a mount point for the file system. If you no longer need the mount point, delete it manually.
         * - **ECS_FILE**: indicates that the ECS file is backed up.
         * - **File**: indicates a local File backup.
         * 
         * @return builder
         * 
         */
        public Builder sourceType(String sourceType) {
            return sourceType(Output.of(sourceType));
        }

        /**
         * @param speedLimit This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates backup flow control. The format is {start}{end}{bandwidth}. Multiple flow control configurations use partitioning, and no overlap in configuration time is allowed. start: start hour. end: end of hour. bandwidth: limit rate, in KB/s.
         * 
         * @return builder
         * 
         */
        public Builder speedLimit(@Nullable Output<String> speedLimit) {
            $.speedLimit = speedLimit;
            return this;
        }

        /**
         * @param speedLimit This parameter is required only when the value of SourceType is ECS_FILE or File. Indicates backup flow control. The format is {start}{end}{bandwidth}. Multiple flow control configurations use partitioning, and no overlap in configuration time is allowed. start: start hour. end: end of hour. bandwidth: limit rate, in KB/s.
         * 
         * @return builder
         * 
         */
        public Builder speedLimit(String speedLimit) {
            return speedLimit(Output.of(speedLimit));
        }

        public PolicyBindingState build() {
            return $;
        }
    }

}
