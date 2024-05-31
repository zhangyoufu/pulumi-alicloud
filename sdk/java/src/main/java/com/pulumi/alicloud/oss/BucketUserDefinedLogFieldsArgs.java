// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BucketUserDefinedLogFieldsArgs extends com.pulumi.resources.ResourceArgs {

    public static final BucketUserDefinedLogFieldsArgs Empty = new BucketUserDefinedLogFieldsArgs();

    /**
     * The name of the bucket.
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * Container for custom request header configuration information.
     * 
     */
    @Import(name="headerSets")
    private @Nullable Output<List<String>> headerSets;

    /**
     * @return Container for custom request header configuration information.
     * 
     */
    public Optional<Output<List<String>>> headerSets() {
        return Optional.ofNullable(this.headerSets);
    }

    /**
     * Container for custom request parameters configuration information.
     * 
     */
    @Import(name="paramSets")
    private @Nullable Output<List<String>> paramSets;

    /**
     * @return Container for custom request parameters configuration information.
     * 
     */
    public Optional<Output<List<String>>> paramSets() {
        return Optional.ofNullable(this.paramSets);
    }

    private BucketUserDefinedLogFieldsArgs() {}

    private BucketUserDefinedLogFieldsArgs(BucketUserDefinedLogFieldsArgs $) {
        this.bucket = $.bucket;
        this.headerSets = $.headerSets;
        this.paramSets = $.paramSets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BucketUserDefinedLogFieldsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BucketUserDefinedLogFieldsArgs $;

        public Builder() {
            $ = new BucketUserDefinedLogFieldsArgs();
        }

        public Builder(BucketUserDefinedLogFieldsArgs defaults) {
            $ = new BucketUserDefinedLogFieldsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param headerSets Container for custom request header configuration information.
         * 
         * @return builder
         * 
         */
        public Builder headerSets(@Nullable Output<List<String>> headerSets) {
            $.headerSets = headerSets;
            return this;
        }

        /**
         * @param headerSets Container for custom request header configuration information.
         * 
         * @return builder
         * 
         */
        public Builder headerSets(List<String> headerSets) {
            return headerSets(Output.of(headerSets));
        }

        /**
         * @param headerSets Container for custom request header configuration information.
         * 
         * @return builder
         * 
         */
        public Builder headerSets(String... headerSets) {
            return headerSets(List.of(headerSets));
        }

        /**
         * @param paramSets Container for custom request parameters configuration information.
         * 
         * @return builder
         * 
         */
        public Builder paramSets(@Nullable Output<List<String>> paramSets) {
            $.paramSets = paramSets;
            return this;
        }

        /**
         * @param paramSets Container for custom request parameters configuration information.
         * 
         * @return builder
         * 
         */
        public Builder paramSets(List<String> paramSets) {
            return paramSets(Output.of(paramSets));
        }

        /**
         * @param paramSets Container for custom request parameters configuration information.
         * 
         * @return builder
         * 
         */
        public Builder paramSets(String... paramSets) {
            return paramSets(List.of(paramSets));
        }

        public BucketUserDefinedLogFieldsArgs build() {
            if ($.bucket == null) {
                throw new MissingRequiredPropertyException("BucketUserDefinedLogFieldsArgs", "bucket");
            }
            return $;
        }
    }

}
