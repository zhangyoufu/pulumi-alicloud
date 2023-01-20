// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFunctionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFunctionsPlainArgs Empty = new GetFunctionsPlainArgs();

    /**
     * A list of functions ids.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of functions ids.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by function name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by function name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Name of the service that contains the functions to find.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return Name of the service that contains the functions to find.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetFunctionsPlainArgs() {}

    private GetFunctionsPlainArgs(GetFunctionsPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFunctionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFunctionsPlainArgs $;

        public Builder() {
            $ = new GetFunctionsPlainArgs();
        }

        public Builder(GetFunctionsPlainArgs defaults) {
            $ = new GetFunctionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of functions ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of functions ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by function name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param serviceName Name of the service that contains the functions to find.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetFunctionsPlainArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
