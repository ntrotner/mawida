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


import * as runtime from '../runtime';
import type {
  PickupConfirmation,
  RentContract,
  ReturnProduct,
  Success,
} from '../models/index';
import {
    PickupConfirmationFromJSON,
    PickupConfirmationToJSON,
    RentContractFromJSON,
    RentContractToJSON,
    ReturnProductFromJSON,
    ReturnProductToJSON,
    SuccessFromJSON,
    SuccessToJSON,
} from '../models/index';

export interface RentalsRentContractIdCancelPostRequest {
    rentContractId: string;
}

export interface RentalsRentContractIdGetRequest {
    rentContractId: string;
}

export interface RentalsRentContractIdPickupPostRequest {
    rentContractId: string;
    pickupConfirmation?: PickupConfirmation;
}

export interface RentalsRentContractIdReturnPostRequest {
    rentContractId: string;
    returnProduct?: ReturnProduct;
}

/**
 * 
 */
export class RentalApi extends runtime.BaseAPI {

    /**
     * Retrieve all rent contracts
     */
    async rentalsGetRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<RentContract>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/rentals`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(RentContractFromJSON));
    }

    /**
     * Retrieve all rent contracts
     */
    async rentalsGet(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<RentContract>> {
        const response = await this.rentalsGetRaw(initOverrides);
        return await response.value();
    }

    /**
     * Cancel a rent contract
     */
    async rentalsRentContractIdCancelPostRaw(requestParameters: RentalsRentContractIdCancelPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Success>> {
        if (requestParameters.rentContractId === null || requestParameters.rentContractId === undefined) {
            throw new runtime.RequiredError('rentContractId','Required parameter requestParameters.rentContractId was null or undefined when calling rentalsRentContractIdCancelPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/rentals/{rentContractId}/cancel`.replace(`{${"rentContractId"}}`, encodeURIComponent(String(requestParameters.rentContractId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SuccessFromJSON(jsonValue));
    }

    /**
     * Cancel a rent contract
     */
    async rentalsRentContractIdCancelPost(requestParameters: RentalsRentContractIdCancelPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Success> {
        const response = await this.rentalsRentContractIdCancelPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve a single rent contract
     */
    async rentalsRentContractIdGetRaw(requestParameters: RentalsRentContractIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RentContract>> {
        if (requestParameters.rentContractId === null || requestParameters.rentContractId === undefined) {
            throw new runtime.RequiredError('rentContractId','Required parameter requestParameters.rentContractId was null or undefined when calling rentalsRentContractIdGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/rentals/{rentContractId}`.replace(`{${"rentContractId"}}`, encodeURIComponent(String(requestParameters.rentContractId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RentContractFromJSON(jsonValue));
    }

    /**
     * Retrieve a single rent contract
     */
    async rentalsRentContractIdGet(requestParameters: RentalsRentContractIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RentContract> {
        const response = await this.rentalsRentContractIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Confirm product pickup
     */
    async rentalsRentContractIdPickupPostRaw(requestParameters: RentalsRentContractIdPickupPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Success>> {
        if (requestParameters.rentContractId === null || requestParameters.rentContractId === undefined) {
            throw new runtime.RequiredError('rentContractId','Required parameter requestParameters.rentContractId was null or undefined when calling rentalsRentContractIdPickupPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/rentals/{rentContractId}/pickup`.replace(`{${"rentContractId"}}`, encodeURIComponent(String(requestParameters.rentContractId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PickupConfirmationToJSON(requestParameters.pickupConfirmation),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SuccessFromJSON(jsonValue));
    }

    /**
     * Confirm product pickup
     */
    async rentalsRentContractIdPickupPost(requestParameters: RentalsRentContractIdPickupPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Success> {
        const response = await this.rentalsRentContractIdPickupPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Confirm product return
     */
    async rentalsRentContractIdReturnPostRaw(requestParameters: RentalsRentContractIdReturnPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Success>> {
        if (requestParameters.rentContractId === null || requestParameters.rentContractId === undefined) {
            throw new runtime.RequiredError('rentContractId','Required parameter requestParameters.rentContractId was null or undefined when calling rentalsRentContractIdReturnPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/rentals/{rentContractId}/return`.replace(`{${"rentContractId"}}`, encodeURIComponent(String(requestParameters.rentContractId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ReturnProductToJSON(requestParameters.returnProduct),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SuccessFromJSON(jsonValue));
    }

    /**
     * Confirm product return
     */
    async rentalsRentContractIdReturnPost(requestParameters: RentalsRentContractIdReturnPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Success> {
        const response = await this.rentalsRentContractIdReturnPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
