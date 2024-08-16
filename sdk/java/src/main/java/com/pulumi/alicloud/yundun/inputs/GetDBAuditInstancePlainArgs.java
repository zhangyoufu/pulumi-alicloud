// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.yundun.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDBAuditInstancePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDBAuditInstancePlainArgs Empty = new GetDBAuditInstancePlainArgs();

    @Import(name="descriptionRegex")
    private @Nullable String descriptionRegex;

    public Optional<String> descriptionRegex() {
        return Optional.ofNullable(this.descriptionRegex);
    }

    @Import(name="ids")
    private @Nullable List<String> ids;

    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetDBAuditInstancePlainArgs() {}

    private GetDBAuditInstancePlainArgs(GetDBAuditInstancePlainArgs $) {
        this.descriptionRegex = $.descriptionRegex;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDBAuditInstancePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDBAuditInstancePlainArgs $;

        public Builder() {
            $ = new GetDBAuditInstancePlainArgs();
        }

        public Builder(GetDBAuditInstancePlainArgs defaults) {
            $ = new GetDBAuditInstancePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder descriptionRegex(@Nullable String descriptionRegex) {
            $.descriptionRegex = descriptionRegex;
            return this;
        }

        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        public GetDBAuditInstancePlainArgs build() {
            return $;
        }
    }

}
