export interface AppState {
    filter: string;
}

export interface Game {

    euRelDate?: string;
    id?: number;
    licensed?: boolean;
    loosePrice?: number;
    name: string;
    naRelDate?: string;
    publishers?: string[];
    wikipediaUrl?: string;
}

export interface PagedData<T> {
    data: T[];
    start: number;
    total: number;
}
