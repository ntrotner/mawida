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
  Location,
  Success,
} from '../models/index';
import {
    LocationFromJSON,
    LocationToJSON,
    SuccessFromJSON,
    SuccessToJSON,
} from '../models/index';

export interface LocationLocationIdGetRequest {
    locationId: string;
}

export interface LocationsPostRequest {
    location?: Location;
}

/**
 * 
 */
export class LocationApi extends runtime.BaseAPI {

    /**
     * Retrieve a single location
     */
    async locationLocationIdGetRaw(requestParameters: LocationLocationIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Location>> {
        if (requestParameters.locationId === null || requestParameters.locationId === undefined) {
            throw new runtime.RequiredError('locationId','Required parameter requestParameters.locationId was null or undefined when calling locationLocationIdGet.');
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
            path: `/location/{locationId}`.replace(`{${"locationId"}}`, encodeURIComponent(String(requestParameters.locationId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => LocationFromJSON(jsonValue));
    }

    /**
     * Retrieve a single location
     */
    async locationLocationIdGet(requestParameters: LocationLocationIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Location> {
        const response = await this.locationLocationIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve all locations
     */
    async locationsGetRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<Location>>> {
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
            path: `/locations`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(LocationFromJSON));
    }

    /**
     * Retrieve all locations
     */
    async locationsGet(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<Location>> {
        const response = await this.locationsGetRaw(initOverrides);
        return await response.value();
    }

    /**
     * Create a new location
     */
    async locationsPostRaw(requestParameters: LocationsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Success>> {
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
            path: `/locations`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: LocationToJSON(requestParameters.location),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SuccessFromJSON(jsonValue));
    }

    /**
     * Create a new location
     */
    async locationsPost(requestParameters: LocationsPostRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Success> {
        const response = await this.locationsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
