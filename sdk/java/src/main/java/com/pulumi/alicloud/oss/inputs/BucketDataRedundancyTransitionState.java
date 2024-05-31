// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BucketDataRedundancyTransitionState extends com.pulumi.resources.ResourceArgs {

    public static final BucketDataRedundancyTransitionState Empty = new BucketDataRedundancyTransitionState();

    /**
     * Storage space name.
     * 
     */
    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    /**
     * @return Storage space name.
     * 
     */
    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    /**
     * Stores the creation time of the redundant transformation task.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Stores the creation time of the redundant transformation task.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Unique identification of the storage redundancy conversion task.
     * 
     */
    @Import(name="taskId")
    private @Nullable Output<String> taskId;

    /**
     * @return Unique identification of the storage redundancy conversion task.
     * 
     */
    public Optional<Output<String>> taskId() {
        return Optional.ofNullable(this.taskId);
    }

    private BucketDataRedundancyTransitionState() {}

    private BucketDataRedundancyTransitionState(BucketDataRedundancyTransitionState $) {
        this.bucket = $.bucket;
        this.createTime = $.createTime;
        this.status = $.status;
        this.taskId = $.taskId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BucketDataRedundancyTransitionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BucketDataRedundancyTransitionState $;

        public Builder() {
            $ = new BucketDataRedundancyTransitionState();
        }

        public Builder(BucketDataRedundancyTransitionState defaults) {
            $ = new BucketDataRedundancyTransitionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket Storage space name.
         * 
         * @return builder
         * 
         */
        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket Storage space name.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param createTime Stores the creation time of the redundant transformation task.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Stores the creation time of the redundant transformation task.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param status Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Stores the state of the redundant translation task. The values are as follows:  Queueing: in the queue.  Processing: In progress.  Finished: Finished.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param taskId Unique identification of the storage redundancy conversion task.
         * 
         * @return builder
         * 
         */
        public Builder taskId(@Nullable Output<String> taskId) {
            $.taskId = taskId;
            return this;
        }

        /**
         * @param taskId Unique identification of the storage redundancy conversion task.
         * 
         * @return builder
         * 
         */
        public Builder taskId(String taskId) {
            return taskId(Output.of(taskId));
        }

        public BucketDataRedundancyTransitionState build() {
            return $;
        }
    }

}
