/* tslint:disable */
/* eslint-disable */
/**
 * Swagger - OpenAPI 3.0
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: nikita@ttnr.me
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ProductDocumentsInner
 */
export interface ProductDocumentsInner {
    /**
     * Unique identifier for the document
     * @type {string}
     * @memberof ProductDocumentsInner
     */
    id: string;
    /**
     * Name of the document
     * @type {string}
     * @memberof ProductDocumentsInner
     */
    name: string;
    /**
     * Base64 encoded document data
     * @type {string}
     * @memberof ProductDocumentsInner
     */
    data: string;
}

/**
 * Check if a given object implements the ProductDocumentsInner interface.
 */
export function instanceOfProductDocumentsInner(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "data" in value;

    return isInstance;
}

export function ProductDocumentsInnerFromJSON(json: any): ProductDocumentsInner {
    return ProductDocumentsInnerFromJSONTyped(json, false);
}

export function ProductDocumentsInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProductDocumentsInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'data': json['data'],
    };
}

export function ProductDocumentsInnerToJSON(value?: ProductDocumentsInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'data': value.data,
    };
}

