// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BucketRequestPaymentArgs extends com.pulumi.resources.ResourceArgs {

    public static final BucketRequestPaymentArgs Empty = new BucketRequestPaymentArgs();

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
     * The payer of the request and traffic fees.Valid values: BucketOwner: request and traffic fees are paid by the bucket owner. Requester: request and traffic fees are paid by the requester.
     * 
     */
    @Import(name="payer")
    private @Nullable Output<String> payer;

    /**
     * @return The payer of the request and traffic fees.Valid values: BucketOwner: request and traffic fees are paid by the bucket owner. Requester: request and traffic fees are paid by the requester.
     * 
     */
    public Optional<Output<String>> payer() {
        return Optional.ofNullable(this.payer);
    }

    private BucketRequestPaymentArgs() {}

    private BucketRequestPaymentArgs(BucketRequestPaymentArgs $) {
        this.bucket = $.bucket;
        this.payer = $.payer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BucketRequestPaymentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BucketRequestPaymentArgs $;

        public Builder() {
            $ = new BucketRequestPaymentArgs();
        }

        public Builder(BucketRequestPaymentArgs defaults) {
            $ = new BucketRequestPaymentArgs(Objects.requireNonNull(defaults));
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
         * @param payer The payer of the request and traffic fees.Valid values: BucketOwner: request and traffic fees are paid by the bucket owner. Requester: request and traffic fees are paid by the requester.
         * 
         * @return builder
         * 
         */
        public Builder payer(@Nullable Output<String> payer) {
            $.payer = payer;
            return this;
        }

        /**
         * @param payer The payer of the request and traffic fees.Valid values: BucketOwner: request and traffic fees are paid by the bucket owner. Requester: request and traffic fees are paid by the requester.
         * 
         * @return builder
         * 
         */
        public Builder payer(String payer) {
            return payer(Output.of(payer));
        }

        public BucketRequestPaymentArgs build() {
            if ($.bucket == null) {
                throw new MissingRequiredPropertyException("BucketRequestPaymentArgs", "bucket");
            }
            return $;
        }
    }

}
