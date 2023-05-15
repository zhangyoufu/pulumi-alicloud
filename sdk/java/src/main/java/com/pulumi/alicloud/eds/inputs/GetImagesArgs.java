// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetImagesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetImagesArgs Empty = new GetImagesArgs();

    /**
     * The desktop type of the image.
     * 
     */
    @Import(name="desktopInstanceType")
    private @Nullable Output<String> desktopInstanceType;

    /**
     * @return The desktop type of the image.
     * 
     */
    public Optional<Output<String>> desktopInstanceType() {
        return Optional.ofNullable(this.desktopInstanceType);
    }

    /**
     * A list of Image IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Image IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The image type of the image. Valid values: `SYSTEM`, `CUSTOM`.
     * 
     */
    @Import(name="imageType")
    private @Nullable Output<String> imageType;

    /**
     * @return The image type of the image. Valid values: `SYSTEM`, `CUSTOM`.
     * 
     */
    public Optional<Output<String>> imageType() {
        return Optional.ofNullable(this.imageType);
    }

    /**
     * A regex string to filter results by Image name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Image name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The os type of the image.
     * 
     */
    @Import(name="osType")
    private @Nullable Output<String> osType;

    /**
     * @return The os type of the image.
     * 
     */
    public Optional<Output<String>> osType() {
        return Optional.ofNullable(this.osType);
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
     * The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetImagesArgs() {}

    private GetImagesArgs(GetImagesArgs $) {
        this.desktopInstanceType = $.desktopInstanceType;
        this.ids = $.ids;
        this.imageType = $.imageType;
        this.nameRegex = $.nameRegex;
        this.osType = $.osType;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetImagesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetImagesArgs $;

        public Builder() {
            $ = new GetImagesArgs();
        }

        public Builder(GetImagesArgs defaults) {
            $ = new GetImagesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param desktopInstanceType The desktop type of the image.
         * 
         * @return builder
         * 
         */
        public Builder desktopInstanceType(@Nullable Output<String> desktopInstanceType) {
            $.desktopInstanceType = desktopInstanceType;
            return this;
        }

        /**
         * @param desktopInstanceType The desktop type of the image.
         * 
         * @return builder
         * 
         */
        public Builder desktopInstanceType(String desktopInstanceType) {
            return desktopInstanceType(Output.of(desktopInstanceType));
        }

        /**
         * @param ids A list of Image IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Image IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Image IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param imageType The image type of the image. Valid values: `SYSTEM`, `CUSTOM`.
         * 
         * @return builder
         * 
         */
        public Builder imageType(@Nullable Output<String> imageType) {
            $.imageType = imageType;
            return this;
        }

        /**
         * @param imageType The image type of the image. Valid values: `SYSTEM`, `CUSTOM`.
         * 
         * @return builder
         * 
         */
        public Builder imageType(String imageType) {
            return imageType(Output.of(imageType));
        }

        /**
         * @param nameRegex A regex string to filter results by Image name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Image name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param osType The os type of the image.
         * 
         * @return builder
         * 
         */
        public Builder osType(@Nullable Output<String> osType) {
            $.osType = osType;
            return this;
        }

        /**
         * @param osType The os type of the image.
         * 
         * @return builder
         * 
         */
        public Builder osType(String osType) {
            return osType(Output.of(osType));
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
         * @param status The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the image. Valid values: `Creating`, `Available`, `CreateFailed`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetImagesArgs build() {
            return $;
        }
    }

}
