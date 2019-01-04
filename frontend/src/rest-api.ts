import axios, { AxiosError, AxiosInstance, AxiosPromise, AxiosRequestConfig, AxiosResponse } from 'axios';
import {
    Game, PagedData
} from './game-collector';

export class RestApi {

    private readonly instance: AxiosInstance;

    constructor() {
        this.instance = axios.create({
            headers: {
                // Stops Spring Boot from challenging authenticated URLs with
                // "WWW-Authenticate: Basic header" (which triggers basic auth modal)
                'X-Requested-With': 'XMLHttpRequest'
            }
        });
    }

    // /**
    //  * Grabs the error response from the server, so we don't have to return an Axios-specific construct.
    //  *
    //  * @param {AxiosError} error The error received from the server.
    //  * @return The error response.
    //  */
    // private static axiosErrorToErrorResponse(error: AxiosError): ErrorResponse {
    //
    //     // AxiosError's data's payload is an ErrorResponse, but it is not a generic type
    //     // for some reason.  That's fine, we take extra care for non ErrorResponses too.
    //
    //     if (error.response) {
    //         if (error.response.data.statusCode && error.response.data.message) {
    //             return error.response.data;
    //         }
    //         return { message: error.message, statusCode: error.response.status };
    //     }
    //
    //     console.error(`No response information in error: ${JSON.stringify(error)}`);
    //     return { message: error.message, statusCode: 0 };
    // }
    //
    // checkAuthentication(): Promise<UserRep> {
    //     return this.instance.get('login')
    //         .then((response:w AxiosResponse<UserRep>) => {
    //             return response.data;
    //         });
    // }

    getGames(start: number, count: number, filter: string): Promise<PagedData<Game>> {

        const url: string = `api/games?start=${start}&count=${count}&filter=${filter}`;

        return this.instance.get(url)
            .then((response: AxiosResponse<PagedData<Game>>) => {
                return response.data;
            });
    }
}

export default new RestApi();
