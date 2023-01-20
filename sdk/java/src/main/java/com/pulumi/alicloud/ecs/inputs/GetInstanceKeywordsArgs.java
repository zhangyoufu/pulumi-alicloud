// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceKeywordsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceKeywordsArgs Empty = new GetInstanceKeywordsArgs();

    /**
     * The type of reserved keyword to query. Valid values: `account`, `database`.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The type of reserved keyword to query. Valid values: `account`, `database`.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetInstanceKeywordsArgs() {}

    private GetInstanceKeywordsArgs(GetInstanceKeywordsArgs $) {
        this.key = $.key;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceKeywordsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceKeywordsArgs $;

        public Builder() {
            $ = new GetInstanceKeywordsArgs();
        }

        public Builder(GetInstanceKeywordsArgs defaults) {
            $ = new GetInstanceKeywordsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The type of reserved keyword to query. Valid values: `account`, `database`.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The type of reserved keyword to query. Valid values: `account`, `database`.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        public GetInstanceKeywordsArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
