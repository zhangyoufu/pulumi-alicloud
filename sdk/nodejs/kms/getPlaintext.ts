// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPlaintext(args: GetPlaintextArgs, opts?: pulumi.InvokeOptions): Promise<GetPlaintextResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:kms/getPlaintext:getPlaintext", {
        "ciphertextBlob": args.ciphertextBlob,
        "encryptionContext": args.encryptionContext,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlaintext.
 */
export interface GetPlaintextArgs {
    /**
     * The ciphertext to be decrypted.
     */
    readonly ciphertextBlob: string;
    /**
     * -
     * (Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).
     */
    readonly encryptionContext?: {[key: string]: string};
}

/**
 * A collection of values returned by getPlaintext.
 */
export interface GetPlaintextResult {
    readonly ciphertextBlob: string;
    readonly encryptionContext?: {[key: string]: string};
    /**
     * The globally unique ID of the CMK. It is the ID of the CMK used to decrypt ciphertext.
     */
    readonly keyId: string;
    /**
     * The decrypted plaintext.
     */
    readonly plaintext: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
