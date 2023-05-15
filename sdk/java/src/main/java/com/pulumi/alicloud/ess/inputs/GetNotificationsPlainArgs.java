// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNotificationsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNotificationsPlainArgs Empty = new GetNotificationsPlainArgs();

    /**
     * A list of notification ids.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of notification ids.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
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

    /**
     * Scaling group id the notifications belong to.
     * 
     */
    @Import(name="scalingGroupId", required=true)
    private String scalingGroupId;

    /**
     * @return Scaling group id the notifications belong to.
     * 
     */
    public String scalingGroupId() {
        return this.scalingGroupId;
    }

    private GetNotificationsPlainArgs() {}

    private GetNotificationsPlainArgs(GetNotificationsPlainArgs $) {
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.scalingGroupId = $.scalingGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNotificationsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNotificationsPlainArgs $;

        public Builder() {
            $ = new GetNotificationsPlainArgs();
        }

        public Builder(GetNotificationsPlainArgs defaults) {
            $ = new GetNotificationsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of notification ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of notification ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
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

        /**
         * @param scalingGroupId Scaling group id the notifications belong to.
         * 
         * @return builder
         * 
         */
        public Builder scalingGroupId(String scalingGroupId) {
            $.scalingGroupId = scalingGroupId;
            return this;
        }

        public GetNotificationsPlainArgs build() {
            $.scalingGroupId = Objects.requireNonNull($.scalingGroupId, "expected parameter 'scalingGroupId' to be non-null");
            return $;
        }
    }

}
