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


/**
 * Unique identifier for the payment method
 * @export
 */
export const PaymentMethod = {
    Cash: 'cash',
    Stripe: 'stripe'
} as const;
export type PaymentMethod = typeof PaymentMethod[keyof typeof PaymentMethod];


export function PaymentMethodFromJSON(json: any): PaymentMethod {
    return PaymentMethodFromJSONTyped(json, false);
}

export function PaymentMethodFromJSONTyped(json: any, ignoreDiscriminator: boolean): PaymentMethod {
    return json as PaymentMethod;
}

export function PaymentMethodToJSON(value?: PaymentMethod | null): any {
    return value as any;
}

